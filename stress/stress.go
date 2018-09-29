package stress

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/VividCortex/multitick"
	"github.com/dialogs/stressbot/bot"
	pb "github.com/dialogs/stressbot/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type wrappedBot struct {
	bot     *bot.Bot
	reads   map[pb.OutPeer]int64
	dialogs []pb.OutPeer
}

func (w *wrappedBot) choosePeer() (res pb.OutPeer, ok bool) {
	ok = false
	l := len(w.dialogs)
	if l == 0 {
		return
	}
	ok = true
	res = w.dialogs[rand.Intn(l)]
	return
}

func (w *wrappedBot) appendDialog(peer pb.OutPeer) {
	w.dialogs = append(w.dialogs, peer)
	w.reads[peer] = time.Now().Unix()
}

func (w *wrappedBot) befriend(bots []*wrappedBot, minQty int64) {
	for _, friend := range bots {
		if friend.bot == w.bot {
			continue
		}
		//outpeers, err := w.bot.SearchContact(friend.bot.Self.Nick.Value)
		if len(friend.bot.Phones) == 0 {
			continue
		}
		outpeers, err := w.bot.SearchContact(strconv.Itoa(int(friend.bot.Phones[0])))
		if err != nil {
			fmt.Printf("Failed to import contacts: %s\n", err.Error())
			os.Exit(1)
		}
		if len(outpeers.UserPeers) == 0 {
			fmt.Printf("User with name `%s` not found\n", friend.bot.Self.Nick.Value)
			os.Exit(1)
		}
		for _, peer := range outpeers.UserPeers {
			if minQty <= 0 {
				return
			}
			w.bot.AddContact(peer)
			w.appendDialog(pb.OutPeer{Type: pb.PeerType_PEERTYPE_PRIVATE, Id: peer.Uid, AccessHash: peer.AccessHash})
			minQty--
		}
	}
}

type stress struct {
	ctx         context.Context
	bots        []*wrappedBot
	serverUrl   string
	serverPort  int64
	dialTimeout time.Duration
}

func (s *stress) chooseBot() *wrappedBot {
	l := len(s.bots)
	return s.bots[rand.Intn(l)]
}

func (s *stress) createConn(cert tls.Certificate) *grpc.ClientConn {
	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	})
	conn, err := grpc.Dial(s.serverUrl, grpc.WithTransportCredentials(creds), grpc.WithTimeout(s.dialTimeout), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		PermitWithoutStream: true,
	}))
	if err != nil {
		fmt.Printf("Failed to establish connection: %s\n", err.Error())
		os.Exit(1)
	}
	return conn
}

func genPhone() int64 {
	pow := int64(10)
	result := int64(0)
	var rnd int64
	for pos := 0; pos < 11; pos++ {
		switch pos {
		case 0:
			rnd = 7
		case 1:
			rnd = 9
		case 2:
			rnd = 2
		case 3:
			rnd = 6
		default:
			rnd = int64(1 + rand.Intn(8))
		}
		result += rnd * pow
		pow = pow * 10
	}
	return result
}

// TODO: add retry logic
func (s *stress) createLoop(createCh chan string, output chan *bot.Bot, doneCh chan struct{}, certsPath string) {
	for name := range createCh {
		//botFile := certsPath + "/" + name
		//cert, err := tls.LoadX509KeyPair(botFile+".crt", botFile+".key")
		//if err != nil {
		//	fmt.Printf("Unable to load X509 keypair: %s\n", err.Error())
		//	os.Exit(1)
		//}
		//
		//conn := s.createConn(cert)
		creds := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, s.serverUrl))

		address := fmt.Sprintf("%s:%d", s.serverUrl, s.serverPort)
		//fmt.Printf("Failed to instantiate a new bot: %s\n", err.Error())

		conn, err := grpc.Dial(address, grpc.WithBlock(), creds)
		if err != nil {
			fmt.Printf("Failed to instantiate a new bot:")
		}

		retries := 0
		var b *bot.Bot
		for {
			b, err = bot.New(s.ctx, conn, 1, "bot", "bot", "bot", name)
			if err == nil {
				retries = 0
				break
			}
			if err != nil && retries > 10 {
				fmt.Printf("Failed to instantiate a new bot: %s\n", err.Error())
				os.Exit(1)
			}
			if err != nil {
				retries++
			}
			//if err != nil {
			//	fmt.Printf("Failed to instantiate a new bot: %s\n", err.Error())
			//	os.Exit(1)
			//}
		}
		fmt.Printf("Created bot: %s\n", b.Name)
		//_, err = b.AnonymousAuth()
		for {
			_, err = b.AuthorizeByPhone(genPhone(), "12345")
			if err == nil {
				break
			}
		}
		//if err != nil {
		//	fmt.Printf("Failed to authenticate a new bot: %s\n", err.Error())
		//	os.Exit(1)
		//}
		output <- b
		<- time.After(2 * time.Second)
	}
	doneCh <- struct{}{}
}

func (s *stress) spawnBots(users []string, certsPath string, totalBots, creationParallelism int64) {
	botsCh := make(chan *bot.Bot, totalBots)
	createCh := make(chan string, creationParallelism)
	doneCh := make(chan struct{})
	for i := 0; i < int(creationParallelism); i++ {
		go s.createLoop(createCh, botsCh, doneCh, certsPath)
	}
	go func() {
		for _, user := range users {
			fmt.Printf("Scheduling creation of bot for %s. Left: %d\n", user, totalBots)
			if totalBots == 0 {
				break
			}
			createCh <- user
			totalBots--
		}
		close(createCh)
	}()

	wg := int(creationParallelism)
	for {
		select {
		case b, ok := <-botsCh:
			if !ok {
				goto stop
			}
			s.bots = append(s.bots, &wrappedBot{bot: b, dialogs: make([]pb.OutPeer, 0), reads: make(map[pb.OutPeer]int64)})
		case <-doneCh:
			wg--
			if wg == 0 {
				close(botsCh)
			}
		}
	}
stop:
	fmt.Printf("Created %d bots\n ", len(s.bots))
}

func randIntN(r rand.Source, n int64) int64 {
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int63() & (n - 1)
	}
	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > max {
		v = r.Int63()
	}
	return v % n

}

func (s *stress) botLoop(b *wrappedBot, ticker <-chan time.Time, sFreq, rFreq int64) {
	rs := rand.NewSource(time.Now().Unix())
	sendTick := int64(0)
	readTick := int64(0)
	nextSendTick := sFreq
	nextReadTick := rFreq
	for range ticker {
		sendTick++
		readTick++
		if sendTick >= nextSendTick {
			sendTick = 0
			nextSendTick = randIntN(rs, rFreq)
			peer, ok := b.choosePeer()
			if !ok {
				continue
			}
			if _, err := b.bot.SendMessage(&peer); err != nil {
				//panic(err) // inc error count
				fmt.Printf("Bot %s failed to send message to %d:%s: %s\n", b.bot.Name, peer.Id, peer.Type.String(), err.Error())
			}
			fmt.Printf("Bot %s sent message to %d:%s\n", b.bot.Name, peer.Id, peer.Type.String())
		}
		if readTick >= nextReadTick {
			readTick = 0
			nextReadTick = randIntN(rs, rFreq)
			peer, ok := b.choosePeer()
			if !ok {
				continue
			}
			dt := b.reads[peer]
			if err := b.bot.ReadMessage(&peer, dt); err != nil {
				//panic(err) // inc error count
				fmt.Printf("Bot %s failed to read messages of %d:%s: %s\n", b.bot.Name, peer.Id, peer.Type.String(), err.Error())
				continue
			}
			b.reads[peer] = time.Now().Unix()
			fmt.Printf("Bot %s read messages of %d:%s\n", b.bot.Name, peer.Id, peer.Type.String())
		}
	}
}

func Stress(serverUrl, certsPath, usersFile string, serverPort, cPar, bQty, grQty, grMin, grMax, sFreq, rFreq int64) {

	//clientDeadline := time.Now().Add(time.Duration(100) * time.Hour)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(500) * time.Second)

	s := &stress{
		ctx:         ctx,
		serverUrl:   serverUrl,
		serverPort:  serverPort,
		bots:        make([]*wrappedBot, 0),
		dialTimeout: 55 * time.Second,
	}
	//uf, err := ioutil.ReadFile(usersFile)
	//if err != nil {
	//	fmt.Printf("Failed to read uesrs file: %s\n", err.Error())
	//	os.Exit(1)
	//}
	users := make([]string, 0)
	//for _, email := range strings.Split(string(uf), "\n") {
	//	users = append(users, strings.Split(email, "@")[0])
	//}
	for i := 0; i < int(bQty); i++ {
		users = append(users, "BOTUS_"+strconv.Itoa(i))
	}
	s.spawnBots(users, certsPath, bQty, cPar)

	for i := 0; i < int(grQty); i++ {
		b := s.chooseBot()
		g, err := b.bot.CreateGroup("group_"+strconv.Itoa(i), []*pb.UserOutPeer{})
		if err != nil {
			fmt.Printf("Bot failed to create group: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Created group: %s [id=%d]\n", g.Group.Title, g.Group.Id)

		peer := &pb.GroupOutPeer{GroupId: g.Group.Id, AccessHash: g.Group.AccessHash}
		url, err := b.bot.GetGroupInviteUrl(peer)
		if err != nil {
			fmt.Printf("Bot failed to get group invite url: %s\n", err.Error())
			os.Exit(1)
		}

		members := make([]*wrappedBot, 0)
		for _, m := range s.bots {
			if m != b {
				members = append(members, m)
			}
		}

		for i := 0; i < int(grMin)+rand.Intn(int(grMax)); i++ {
			l := len(members)
			if l == 0 {
				fmt.Println("Not enough bots to meet group-min/group-max criteria\n")
				os.Exit(1)
			}
			randomBot := members[rand.Intn(l)]
			jg, err := randomBot.bot.JoinGroupByUrl(url.Url)
			if err != nil {
				fmt.Printf("Bot failed to join group by url: %s\n", err.Error())
				os.Exit(1)
			}
			outPeer := pb.OutPeer{Type: pb.PeerType_PEERTYPE_GROUP, Id: jg.Group.Id, AccessHash: jg.Group.AccessHash}
			b.appendDialog(outPeer)
			fmt.Printf("Bot %s joined group %s [id=%d]\n", randomBot.bot.Name, jg.Group.Title, jg.Group.Id)
		}
	}

	for _, b := range s.bots {
		b.befriend(s.bots, grMin)
	}
	fmt.Printf("Starting stress test")
	ticker := multitick.NewTicker(time.Second, 0)
	for _, b := range s.bots {
		<-time.After(time.Duration(rand.Intn(int(sFreq))) * time.Second)
		go s.botLoop(b, ticker.Subscribe(), sFreq, rFreq)
	}
	stopS := make(chan os.Signal, 1)
	signal.Notify(stopS, os.Interrupt, os.Kill)
	<-stopS
}
