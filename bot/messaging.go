package bot

import (
	"math/rand"
	"strconv"
)
import pb "github.com/dialogs/stressbot/gateway"

func (b *Bot) SendMessage(dest *pb.OutPeer) (*pb.ResponseSeqDate, error) {
	text := string(strconv.AppendInt([]byte("Message from bot "), int64(b.Self.Id), 10))
	message := &pb.MessageContent{Body: &pb.MessageContent_TextMessage{TextMessage: &pb.TextMessage{Text: text}}}
	return b.messaging.SendMessage(b.authContext, &pb.RequestSendMessage{
		Peer: dest, Rid: rand.Int63(), Message: message,
	})
}

func (b *Bot) ReadMessage(dest *pb.OutPeer, date int64) error {
	_, err := b.messaging.MessageRead(b.authContext, &pb.RequestMessageRead{Peer: dest, Date: date})
	return err
}

func (b *Bot) MessageReceived(dest *pb.OutPeer, date int64) error {
	_, err := b.messaging.MessageReceived(b.authContext, &pb.RequestMessageReceived{Peer: dest, Date: date})
	return err
}