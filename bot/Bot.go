package bot

import (
	pb "dialog-stress-bots/gateway"
	"dialog-stress-bots/utils"
	"net/http"
	"os"
	"strings"

	"io"
	"math"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func CreateBot(phone int64, conn *grpc.ClientConn) (*Bot, error) {
	reg := pb.NewRegistrationClient(conn)

	res, err := reg.RegisterDevice(context.Background(), &pb.RequestRegisterDevice{
		AppId:    125,
		AppTitle: "blabla",
	})
	if err != nil {
		log.Errorf("Can't register device: %s", err)
		return nil, err
	}
	log.Printf("Register response: %d, %s", res.GetAuthId(), res.GetToken())

	ctx := metadata.AppendToOutgoingContext(context.Background(), "x-auth-ticket", res.GetToken())

	auth := pb.NewAuthenticationClient(conn)

	resAuth, errAuth := auth.StartPhoneAuth(ctx, &pb.RequestStartPhoneAuth{
		PhoneNumber: phone,
		AppId:       1,
		ApiKey:      "api_key",
		DeviceTitle: "device_title",
	})
	if errAuth != nil {
		log.Errorf("Can't start auth: %s", errAuth)
		return nil, err
	}

	_, validateErr := auth.ValidateCode(ctx, &pb.RequestValidateCode{
		TransactionHash: resAuth.GetTransactionHash(),
		Code:            "12345",
	})
	if validateErr != nil && !strings.Contains(validateErr.Error(), "Unnoccupied phone number") {
		log.Errorf("Can't validate code: %s, %s", validateErr.Error(), conn.GetState())
		conn.Close()
		return nil, validateErr
	}

	signUpRes, signUpErr := auth.SignUp(ctx, &pb.RequestSignUp{
		TransactionHash: resAuth.GetTransactionHash(),
		Name:            "New Bot",
	})
	if signUpErr != nil {
		log.Errorf("Can't sign up: %s", signUpErr)
		return nil, signUpErr
	}
	log.Printf("Sign up res: %d, %d", signUpRes.GetUser().Id, signUpRes.GetUser().AccessHash)

	contacts := pb.NewContactsClient(conn)
	messaging := pb.NewMessagingClient(conn)
	groups := pb.NewGroupsClient(conn)
	files := pb.NewMediaAndFilesClient(conn)

	dlgs := make(map[int32]struct {
		*pb.OutPeer
		int64
	})

	seqUpdatesClient := pb.NewSequenceAndUpdatesClient(conn)
	seqUpdates, errSeqUpd := seqUpdatesClient.SeqUpdates(ctx, &empty.Empty{})
	if errSeqUpd != nil {
		log.Fatalf("Can't subscribe on seq updates: %v", errSeqUpd)
		return nil, errSeqUpd
	}

	bot := &Bot{
		Peer: pb.OutPeer{
			Type:       pb.PeerType(1),
			Id:         signUpRes.GetUser().Id,
			AccessHash: signUpRes.GetUser().AccessHash,
			StrId:      &wrappers.StringValue{Value: ""},
		},
		phone:        phone,
		user:         signUpRes.GetUser(),
		contacts:     make([]*pb.OutPeer, 0),
		dialogs:      dlgs,
		contactsSrv:  &contacts,
		messagingSrv: &messaging,
		groupsSrv:    &groups,
		seqUpdates:   &seqUpdates,
		filesSrv:     &files,
		conn:         conn,
		ctx:          &ctx,
	}
	return bot, nil
}

func CreateCertBot(conn *grpc.ClientConn) (*Bot, error) {
	reg := pb.NewRegistrationClient(conn)

	res, err := reg.RegisterDevice(context.Background(), &pb.RequestRegisterDevice{
		AppId:    125,
		AppTitle: "blabla",
	})
	if err != nil {
		log.Errorf("Can't register device: %s", err)
		return nil, err
	}
	log.Printf("Register response: %d, %s", res.GetAuthId(), res.GetToken())

	ctx := metadata.AppendToOutgoingContext(context.Background(), "x-auth-ticket", res.GetToken())

	auth := pb.NewAuthenticationClient(conn)

	responseAuth, e := auth.StartAnonymousAuth(ctx, &pb.RequestStartAnonymousAuth{
		Name: "dumb",
	})

	if e != nil {
		log.Errorf("Can't start auth: %s", e)
		return nil, err
	}

	contacts := pb.NewContactsClient(conn)
	messaging := pb.NewMessagingClient(conn)
	groups := pb.NewGroupsClient(conn)
	dlgs := make(map[int32]struct {
		*pb.OutPeer
		int64
	})

	seqUpdatesClient := pb.NewSequenceAndUpdatesClient(conn)
	seqUpdates, errSeqUpd := seqUpdatesClient.SeqUpdates(ctx, &empty.Empty{})
	if errSeqUpd != nil {
		log.Fatalf("Can't subscribe on seq updates: %v", errSeqUpd)
		return nil, errSeqUpd
	}

	bot := &Bot{
		Peer: pb.OutPeer{
			Type:       pb.PeerType(1),
			Id:         responseAuth.GetUser().Id,
			AccessHash: responseAuth.GetUser().AccessHash,
			StrId:      &wrappers.StringValue{Value: ""},
		},
		phone:        100,
		user:         responseAuth.GetUser(),
		contacts:     make([]*pb.OutPeer, 0),
		dialogs:      dlgs,
		contactsSrv:  &contacts,
		messagingSrv: &messaging,
		groupsSrv:    &groups,
		seqUpdates:   &seqUpdates,
		conn:         conn,
		ctx:          &ctx,
	}
	return bot, nil
}



func (bot *Bot) ImportContacts(importBots []int64) error {
	toImport := make([]*pb.PhoneToImport, len(importBots))
	for i, phone := range importBots {
		toImport[i] = &pb.PhoneToImport{
			PhoneNumber: phone,
			Name:        &wrappers.StringValue{Value: "Bot_" + string(phone)},
		}
	}

	res, err := (*bot.contactsSrv).ImportContacts(*bot.ctx, &pb.RequestImportContacts{
		Phones: toImport,
		Emails: make([]*pb.EmailToImport, 0),
	})
	if err != nil {
		log.Errorf("Error while import contacts: %s", err.Error())
		return err
	}

	userOutPeers := res.GetUserPeers()
	var peers []*pb.OutPeer

	bot.lock.Lock()
	for _, p := range userOutPeers {
		peer := &pb.OutPeer{
			Type:       pb.PeerType(1),
			Id:         p.Uid,
			AccessHash: p.AccessHash,
			StrId:      &wrappers.StringValue{Value: ""},
		}
		peers = append(peers, peer)

		bot.dialogs[p.Uid] = struct {
			*pb.OutPeer
			int64
		}{peer, 0}
	}
	bot.contacts = append(bot.contacts, peers...)
	bot.lock.Unlock()

	return nil
}

func (bot *Bot) FindContacts(searchString string) error {

	res, err := (*bot.contactsSrv).SearchContacts(*bot.ctx, &pb.RequestSearchContacts{
		Request: searchString,
	})

	if err != nil {
		log.Errorf("Error while FindContacts: %s", err.Error())
		return err
	}

	userOutPeers := res.GetUserPeers()
	var peers []*pb.OutPeer

	bot.lock.Lock()
	for _, p := range userOutPeers {
		peer := &pb.OutPeer{
			Type:       pb.PeerType(1),
			Id:         p.Uid,
			AccessHash: p.AccessHash,
			StrId:      &wrappers.StringValue{Value: ""},
		}
		peers = append(peers, peer)

		bot.dialogs[p.Uid] = struct {
			*pb.OutPeer
			int64
		}{peer, 0}
	}
	bot.contacts = append(bot.contacts, peers...)

	bot.lock.Unlock()
	return nil
}

func (bot *Bot) SendMessage(dest *pb.OutPeer) (*pb.ResponseSeqDate, error) {

	message := pb.MessageContent{
		Body: &pb.MessageContent_TextMessage{
			TextMessage: &pb.TextMessage{
				Text:       "Message from bot " + strconv.Itoa(int(bot.user.Id)),
				Mentions:   make([]int32, 0),
				Ext:        nil,
				Media:      make([]*pb.MessageMedia, 0),
				Extensions: make([]*pb.Any, 0),
			},
		},
	}

	sendMessageStart := time.Now()
	randomId := utils.RandomId()
	res, err := (*bot.messagingSrv).SendMessage(*bot.ctx, &pb.RequestSendMessage{
		Peer:    dest,
		Rid:     randomId,
		Message: &message,
	})
	if err != nil {
		log.Errorf("Error while sendMessage: %v, %v, %d", err, dest, randomId)
	} else {
		log.Debugf(
			"Message (%d, %s) sent in %fs from %d to %d",
			time.Since(sendMessageStart).Seconds(),
			randomId,
			res.Mid.String(),
			bot.user.Id,
			dest.Id,
		)
		bot.lock.Lock()
		bot.dialogs[dest.Id] = struct {
			*pb.OutPeer
			int64
		}{dest, res.Date}
		bot.lock.Unlock()
	}
	return res, err
}

func (bot *Bot) SendFile(filePath string) error {

	// Reading file and stats
	file, err := os.Open(filePath)

	if err != nil {
		log.Errorf("Error on open file %s with error %v", filePath, err)
		return nil
	}
	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		return nil
	}

	// Obtain file upload url
	resUrl, err := (*bot.filesSrv).GetFileUploadUrl(*bot.ctx, &pb.RequestGetFileUploadUrl{
		ExpectedSize: int32(info.Size()),
	})

	if err != nil {
		log.Errorf("Error on GetFileUploadUrl %v", err)
		return nil
	}

	// Upload the url
	res, err := http.NewRequest("PUT", resUrl.Url, file)

	if err != nil {
		log.Errorf("Error on perform PUT request %v", err)
		return nil
	}

	defer res.Body.Close()

	// Commit file
	_, commitErr := (*bot.filesSrv).CommitFileUpload(*bot.ctx, &pb.RequestCommitFileUpload{
		UploadKey: resUrl.UploadKey,
		FileName:  info.Name(),
	})

	if commitErr != nil {
		log.Errorf("Error on commit file %v", err)
		return nil
	}

	return nil
}

func (bot *Bot) CreateGroup() (*Group, error) {

	//log.Printf("Create group with %d members", len(membersMap))
	members := make([]*pb.UserOutPeer, 0)

	randomId := utils.RandomId()
	title := "GroupTitle_" + string(randomId)
	res, err := (*bot.groupsSrv).CreateGroup(*bot.ctx, &pb.RequestCreateGroup{
		Rid:           randomId,
		Title:         title,
		Users:         members,
		GroupType:     pb.GroupType(1),
		Username:      nil,
		Optimizations: make([]pb.UpdateOptimization, 0),
	})
	if err != nil {
		log.Errorf("Group creation failed: %v", err)
		return nil, err
	}

	group := res.GetGroup()

	urlRequest := &pb.RequestGetGroupInviteUrl{GroupPeer: &pb.GroupOutPeer{GroupId: group.Id, AccessHash: group.AccessHash}}
	inviteUrl, err := (*bot.groupsSrv).GetGroupInviteUrl(*bot.ctx, urlRequest)
	if err != nil {
		return nil, err
	}

	addDialog(bot, &pb.OutPeer{
		Type:       pb.PeerType(2),
		Id:         group.Id,
		AccessHash: group.AccessHash,
	})

	return &Group{InviteUrl: inviteUrl.Url, UsersCount: 1}, nil
}

func addDialog(bot *Bot, peer *pb.OutPeer) {
	bot.lock.Lock()
	bot.dialogs[peer.Id] = struct {
		*pb.OutPeer
		int64
	}{peer, 0}
	bot.lock.Unlock()
}

func (bot *Bot) ReadMessage(dest *pb.OutPeer) error {
	bot.lock.RLock()
	date := bot.dialogs[dest.Id].int64
	bot.lock.RUnlock()
	_, err := (*bot.messagingSrv).MessageRead(*bot.ctx, &pb.RequestMessageRead{
		Peer: dest,
		Date: date,
	})
	if err != nil {
		log.Errorf("Error while readMessage: %s", err.Error())
	}

	log.Debugf("Message read (%d, %d, %d)", date, bot.user.Id, dest.Id)

	return err
}

func (bot *Bot) MessageReceived(dest *pb.OutPeer, date int64) error {
	_, err := (*bot.messagingSrv).MessageReceived(*bot.ctx, &pb.RequestMessageReceived{
		Peer: dest,
		Date: date,
	})

	if err != nil {
		log.Errorf("Error while receiveMessage: %s", err.Error())
		return err
	}

	log.Debugf("Message receive (%d, %d, %d)", date, bot.user.Id, dest.Id)

	return nil
}

type Group struct {
	sync.RWMutex
	InviteUrl  string
	UsersCount int32
}

// increment users count
func (gr *Group) IncUsersCount() {
	gr.Lock()
	gr.UsersCount += 1
	gr.Unlock()
}

// get users count
func (gr *Group) GetUsersCount() int32 {
	gr.RLock()
	defer gr.RUnlock()
	return gr.UsersCount
}

// Bot state
type Bot struct {
	lock     sync.RWMutex
	Peer     pb.OutPeer
	phone    int64
	user     *pb.User
	contacts []*pb.OutPeer
	dialogs  map[int32]struct {
		*pb.OutPeer
		int64
	}

	contactsSrv  *pb.ContactsClient
	messagingSrv *pb.MessagingClient
	groupsSrv    *pb.GroupsClient
	seqUpdates   *pb.SequenceAndUpdates_SeqUpdatesClient
	filesSrv	 *pb.MediaAndFilesClient
	conn         *grpc.ClientConn
	ctx          *context.Context
}

type KillBot struct{}

func (bot *Bot) GetContacts() []*pb.OutPeer {
	return bot.contacts
}

func (bot *Bot) choosePeer() *pb.OutPeer {
	bot.lock.RLock()
	defer bot.lock.RUnlock()
	dialogsCount := len(bot.dialogs)

	if dialogsCount == 0 {
		return nil
	}

	var r int
	if dialogsCount == 1 {
		r = 0
	} else {
		r = utils.RandomInt(dialogsCount - 1)
	}

	i := 0
	for _, d := range bot.dialogs {
		if i == r {
			return d.OutPeer
		}
		i++
	}
	return nil
}
func (bot *Bot) Start(msgEachS int, readEachS int) chan<- KillBot {
	go bot.seqUpdatesReceive(bot.seqUpdates)

	msgTicker := randomTick(msgEachS)
	readTicker := randomTick(readEachS)
	//fileTick := randomTick(msgEachS)

	quit := make(chan KillBot)
	go func() {
		for {
			select {
			case <-msgTicker:
				peer := bot.choosePeer()
				if peer != nil {
					bot.SendMessage(peer)
				}
			case <-readTicker:
				peer := bot.choosePeer()
				if peer != nil {
					bot.ReadMessage(peer)
				}
			case <-quit:
				return
			}
		}
	}()

	return quit
}

//
func randomTick(tickBaseS int) <-chan struct{} {
	msgEachRandomPart := math.Floor(float64(tickBaseS*1000) / 3)
	msgRandomPartInt := int(msgEachRandomPart)
	msgEachBase := float64(tickBaseS*1000) - msgEachRandomPart
	msgTicker := make(chan struct{})
	go func() {
		for range time.NewTicker(time.Duration(msgEachBase) * time.Millisecond).C {
			sleepInS := time.Duration(utils.RandomInt(msgRandomPartInt))
			//fmt.Printf("sleep for %d/%f ms\n", sleepInS, msgEachRandomPart)
			time.Sleep(sleepInS * time.Millisecond)
			msgTicker <- struct{}{}
		}
	}()
	return msgTicker
}

func extractUpdate(b []byte) (*pb.UpdateSeqUpdate, error) {
	seqUpdate := &pb.UpdateSeqUpdate{}
	err := seqUpdate.XXX_Unmarshal(b)
	return seqUpdate, err
}

func extractFatUpdate(b []byte) (*pb.UpdateFatSeqUpdate, error) {
	seqUpdate := &pb.UpdateFatSeqUpdate{}
	err := seqUpdate.XXX_Unmarshal(b)
	return seqUpdate, err
}

func (bot *Bot) JoinGroup(group *Group) error {
	requestJoinGroupByPeer := &pb.RequestJoinGroup{Token: group.InviteUrl}
	joinGroupStart := time.Now()
	resp, err := (*bot.groupsSrv).JoinGroup(*bot.ctx, requestJoinGroupByPeer)
	log.Infof("JoinGroup %fs", time.Since(joinGroupStart).Seconds())
	if err != nil {
		log.Errorf("Failed to join group by url '%s' : %s", group.InviteUrl, err.Error())
		return err
	}
	log.Infof("Joined group by url '%s'", group.InviteUrl)
	addDialog(bot, &pb.OutPeer{
		Type:       pb.PeerType(2),
		Id:         resp.Group.Id,
		AccessHash: resp.Group.AccessHash,
		StrId:      &wrappers.StringValue{Value: ""},
	})
	group.IncUsersCount()
	return nil
}

func (bot *Bot) seqUpdatesReceive(seqUpdates *pb.SequenceAndUpdates_SeqUpdatesClient) error {
	for {
		update, err := (*seqUpdates).Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		bytes := update.GetUpdate().Value
		supd, extrErr := extractUpdate(bytes)
		if extrErr != nil {
			log.Errorf("Failed to parse SeqUpdate")
			continue
		}

		if supd.GetUpdateMessage() != nil {

			message := supd.GetUpdateMessage()
			log.Debugf("Receive UpdateMessage: %d, %d, %d", message.Peer.Id, message.SenderUid, message.Date)

			bot.lock.Lock()
			var outPeer *pb.OutPeer = nil
			if d, ok := bot.dialogs[message.Peer.Id]; ok {
				outPeer = d.OutPeer
			} else {
				res, dialogsErr := (*bot.messagingSrv).LoadDialogs(*bot.ctx, &pb.RequestLoadDialogs{
					MinDate:       0,
					Limit:         100,
					Optimizations: nil,
					Filters:       nil,
					PeersToLoad:   []*pb.Peer{message.Peer},
				})

				if dialogsErr == nil {
					for _, d := range res.Dialogs {
						for _, g := range res.GroupPeers {
							if g.GroupId == d.Peer.Id {
								outPeer = &pb.OutPeer{
									Type:       pb.PeerType(2),
									Id:         g.GroupId,
									AccessHash: g.AccessHash,
									StrId:      &wrappers.StringValue{Value: ""},
								}
								break
							}
						}

						if outPeer == nil {
							for _, u := range res.UserPeers {
								if u.Uid == d.Peer.Id {
									outPeer = &pb.OutPeer{
										Type:       pb.PeerType(1),
										Id:         u.Uid,
										AccessHash: u.AccessHash,
										StrId:      &wrappers.StringValue{Value: ""},
									}
									break
								}
							}
						}
					}
				} else {
					log.Errorf("Couldn't load dialog %d: %v", message.Peer.Id, dialogsErr)
				}
			}

			if outPeer != nil {
				bot.dialogs[message.Peer.Id] = struct {
					*pb.OutPeer
					int64
				}{outPeer, message.Date}

				go bot.MessageReceived(outPeer, message.Date)
			}

			bot.lock.Unlock()
		}
	}
}
