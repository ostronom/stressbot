package main

import (
	"bufio"
	"crypto/tls"
	b "dialog-stress-bots/bot"
	"dialog-stress-bots/utils"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var emailsPool = make([]string, 0)

func getRandomContacts(amount int) []string {
	var toReturn []string

	poolSize := len(emailsPool)

	for i := 0; i < amount; i++ {
		toReturn = append(toReturn, emailsPool[utils.RandomInt(poolSize-1)])
	}
	return toReturn
}

func init() {

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	logFile, err := os.OpenFile("./log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Error(fmt.Errorf("failed to open log file"))
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	// enable client rpc time metrics
	grpc_prometheus.EnableClientHandlingTimeHistogram()
}

var bc = b.NewBotsCollection()
var gc = b.NewGroupsCollection()

func main() {
	server := flag.String("server", "127.0.0.1", "server host")
	serverPort := flag.String("serverport", "8080", "server port")
	metricsPort := flag.String("metricsport", "8081", "prometheus metrics endpoint port")

	creationParallelism := flag.Int("cpar", 5, "creation parallelism")
	botsAmount := flag.Int("botsnum", 100, "bots amount")
	contactsAmount := flag.Int("ctsnum", 50, "contacts amount")

	groupsAmount := flag.Int("grnum", 0, "groups amount")
	membersMin := flag.Int("grmin", 5, "minimum group members")
	membersMax := flag.Int("grmax", 200, "maximum group members")

	sendMessageEachNSec := flag.Int("sendeach", 12, "send message each n seconds")
	readMessageEachNSec := flag.Int("readeach", 8, "read message each n seconds")

	certsDirectory := "/Users/ademin/tmp/fuuuu"//flag.String("certs", "certs", "directory of certificates")
	usersList := "/Users/ademin/tmp/fuuuu/accounts.txt" //flag.String("users", "users.txt", "file with list of users ")

	flag.Parse()

	log.Infof("server: %s", *server)
	log.Infof("serverport: %s", *serverPort)
	log.Infof("metricsPort: %s", *metricsPort)
	log.Infof("cpar: %d", *creationParallelism)
	log.Infof("botsnum: %d", *botsAmount)
	log.Infof("ctsnum: %d", *contactsAmount)
	log.Infof("grnum: %d", *groupsAmount)
	log.Infof("grmax: %d", *membersMax)
	log.Infof("sendeach: %d", *sendMessageEachNSec)
	log.Infof("readeach: %d", *readMessageEachNSec)

	users, _ := readLines(usersList)
	emailsPool = append(emailsPool, users...)

	//print(users)
	//
	//return
	//for i := 0; i < *botsAmount; i++ {
	//	phonePool = append(phonePool, utils.Random(700000000, 800000000))
	//}

	newGroupEvent := make(chan *b.Group, *creationParallelism)

	initBot := func(newBot b.Bot) {
		addContactStart := time.Now()
		bc.AddBot(&newBot)

		// adding phones into contacts of new bot
		//newBot.ImportContacts(getRandomContacts(*contactsAmount))

		// joining existing groups
		existingGroups := gc.Groups()
		for _, g := range existingGroups {
			if utils.IFeelLucky(*membersMax, *botsAmount) {
				go newBot.JoinGroup(g)
			}
		}

		log.Infof("Contacts imported in: %fs", time.Since(addContactStart).Seconds())

		// create new group
		go func() {
			if utils.IFeelLucky(*groupsAmount, *botsAmount) {
				groupCreationStarted := time.Now()
				log.Infof("Bot %d create group", newBot.Peer.Id)
				group, err := newBot.CreateGroup()
				if err == nil {
					newGroupEvent <- group
					gc.AddGroup(group)
				}
				log.Infof("Group created in: %fs", time.Since(groupCreationStarted).Seconds())
			}
		}()
	}

	go func() {
		for newGroup := range newGroupEvent {
			currentBots := bc.Bots()
			joinGroupChan := make(chan *b.Bot)
			go func() {
				defer close(joinGroupChan)
				for _, bot := range currentBots {
					membersAmount := int(utils.Random(int64(*membersMin), int64(*membersMax)))
					if utils.IFeelLucky(membersAmount, *botsAmount) {
						joinGroupChan <- bot
					}
				}
			}()

			for bot := range joinGroupChan {
				bot.JoinGroup(newGroup)
			}
		}
	}()

	go func() {
		botsToLaunch := make(chan string, *botsAmount)
		go func() {
			defer close(botsToLaunch)
			for _, mail := range emailsPool {
				botsToLaunch <- mail
			}
		}()
		var launchedBots int32 = 0
		botCreationLock := sync.WaitGroup{}
		botCreationLock.Add(*creationParallelism)
		for w := 1; w <= *creationParallelism; w++ {
			go func() {
				defer botCreationLock.Done()
				for email := range botsToLaunch {
					botCreationStart := time.Now()
					clientConn, err := initCertConnection(server, serverPort, email, certsDirectory)
					if err != nil {
						log.Errorf("client did not connect: %s", err)
					}
					bot, err := b.CreateCertBot(clientConn)

					log.Infof("Bot created in %fs", time.Since(botCreationStart).Seconds())

					if err != nil {
						log.Errorf("failed to create bot: %s", err)
					} else {
						botStart := time.Now()
						bot.Start(*sendMessageEachNSec, *readMessageEachNSec)
						initBot(*bot)
						log.Infof("Bot started in: %fs", time.Since(botStart).Seconds())
						atomic.AddInt32(&launchedBots, 1)
					}

					time.Sleep(200 * time.Millisecond)
				}
			}()
		}
		botCreationLock.Wait()
		if launchedBots == 0 {
			log.Fatal("No bots was launched")
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *metricsPort), nil))
	waitForExit()
}

func initClientConnection(host *string, port *string) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%s", *host, *port)
	clientConn, err := grpc.Dial(
		address,
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		//grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		grpc.WithInsecure(),
	)
	return clientConn, err
}

func initCertConnection(host *string, port *string, user string, certsFolder string) (*grpc.ClientConn, error) {
	// load user cert/key
	cert, err := tls.LoadX509KeyPair(certsFolder+"/"+user+".pem", certsFolder+"/"+user+".key")
	if err != nil {
		panic(err)
	}

	// load CA cert
	//caCert, err := ioutil.ReadFile("certs/ca2.pem")
	//if err != nil {
	//	panic(err)
	//}

	//caCertPool := x509.NewCertPool()
	//if !caCertPool.AppendCertsFromPEM(caCert) {
	//	panic("Failed to append caCert to certs pool")
	//}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	config.BuildNameToCertificate()
	creds := credentials.NewTLS(config)
	address := "sbrf.transmit.im:8443"//fmt.Sprintf("%s:%s", *host, *port)
	//address := fmt.Sprintf("%s:%s", *host, *port)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds), grpc.WithTimeout(15*time.Second), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	return conn, err
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Wait for process to be terminated.
func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	os.Exit(0)
}
