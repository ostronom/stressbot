package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	b "stressbot/bot"
	"stressbot/utils"
	"sync"
	"sync/atomic"
	"time"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var phonePool = make([]int64, 0)

func getRandomContacts(amount int) []int64 {
	var toReturn []int64

	poolSize := len(phonePool)

	for i := 0; i < amount; i++ {
		toReturn = append(toReturn, phonePool[utils.RandomInt(poolSize-1)])
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
		fmt.Errorf("failed to open log file")
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
	serverRoots := flag.String("root_certs", "", "Server root certs")
	metricsPort := flag.String("metricsport", "8081", "prometheus metrics endpoint port")

	creationParallelism := flag.Int("cpar", 5, "creation parallelism")
	botsAmount := flag.Int("botsnum", 100, "bots amount")
	contactsAmount := flag.Int("ctsnum", 50, "contacts amount")

	groupsAmount := flag.Int("grnum", 0, "groups amount")
	membersMin := flag.Int("grmin", 5, "minimum group members")
	membersMax := flag.Int("grmax", 200, "maximum group members")

	sendMessageEachNSec := flag.Int("sendeach", 12, "send message each n seconds")
	readMessageEachNSec := flag.Int("readeach", 8, "read message each n seconds")

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

	for i := 0; i < *botsAmount; i++ {
		phonePool = append(phonePool, utils.Random(700000000, 800000000))
	}

	newGroupEvent := make(chan *b.Group, *creationParallelism)

	initBot := func(newBot b.Bot) {
		addContactStart := time.Now()
		bc.AddBot(&newBot)

		// adding phones into contacts of new bot
		newBot.ImportContacts(getRandomContacts(*contactsAmount))

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
		botsToLaunch := make(chan int64, *botsAmount)
		go func() {
			defer close(botsToLaunch)
			for _, phone := range phonePool {
				botsToLaunch <- phone
			}
		}()
		var launchedBots int32 = 0
		botCreationLock := sync.WaitGroup{}
		botCreationLock.Add(*creationParallelism)
		for w := 1; w <= *creationParallelism; w++ {
			go func() {
				defer botCreationLock.Done()
				for phone := range botsToLaunch {
					botCreationStart := time.Now()
					clientConn, err := initCliencConnection(server, serverPort, serverRoots)
					if err != nil {
						log.Errorf("client did not connect: %s", err)
					}
					bot, err := b.CreateBot(phone, clientConn)

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

func initCliencConnection(host *string, port *string, roots *string) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%s", *host, *port)

	var secureOption grpc.DialOption

	if *roots == "" {
		secureOption = grpc.WithInsecure()
	} else {
		creds, err := credentials.NewClientTLSFromFile(*roots, "")
		if err != nil {
			log.Fatal("Credentials are invalid %e", err)
		}
		secureOption = grpc.WithTransportCredentials(creds)
	}

	clientConn, err := grpc.Dial(
		address,
		secureOption,
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		//grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
		// grpc.WithInsecure(),
	)

	return clientConn, err
}

// Wait for process to be terminated.
func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	os.Exit(0)
}
