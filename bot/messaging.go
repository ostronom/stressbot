package bot

import (
	"fmt"
	"github.com/dialogs/stressbot/utils"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
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

func (bot *Bot) SendFile(dest *pb.OutPeer, filePath string)  error {

	// Reading file and stats
	file, err := os.Open(filePath)

	if err != nil {
		//log.Errorf("Error on open file %s with error %v", filePath, err)
		return err
	}
	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		return err
	}

	if info.IsDir() {
		// Send random file from dir
		infos, err := ioutil.ReadDir(filePath)
		if err != nil {
			//log.Errorf("Error on when reading dir infos %v", err)
			return err
		}

		n := rand.Int() % len(infos)
		return bot.SendFile(dest, fmt.Sprintf("%s/%s", filePath, infos[n].Name()))
	}

	size := int32(info.Size())
	// Obtain file upload url
	resUrl, err := (bot.files).GetFileUploadUrl(bot.parentCtx, &pb.RequestGetFileUploadUrl{
		ExpectedSize: size,
	})

	if err != nil {
		//log.Errorf("Error on GetFileUploadUrl %v", err)
		return err
	}

	// Upload the url
	res, err := http.NewRequest("PUT", resUrl.Url, file)

	if err != nil {
		//log.Errorf("Error on perform PUT request %v", err)
		return err
	}

	defer res.Body.Close()

	// Commit file
	commitRes, commitErr := (bot.files).CommitFileUpload(bot.parentCtx, &pb.RequestCommitFileUpload{
		UploadKey: resUrl.UploadKey,
		FileName:  info.Name(),
	})

	if commitErr != nil {
		//log.Errorf("Error on commit file %v", err)
		return err
	}

	message := pb.MessageContent{
		Body: &pb.MessageContent_DocumentMessage{
			DocumentMessage: &pb.DocumentMessage{
				FileId:     commitRes.UploadedFileLocation.FileId,
				AccessHash: commitRes.UploadedFileLocation.AccessHash,
				FileSize:   size,
				Name:       info.Name(),
				MimeType:   "application/octet-stream",
			},
		},
	}

	_, err = PerformSend(bot, dest, message)
	return err
}

func PerformSend(bot *Bot, dest *pb.OutPeer, message pb.MessageContent) (*pb.ResponseSeqDate, error) {

	//sendMessageStart := time.Now()
	randomId := utils.RandomId()
	res, err := bot.messaging.SendMessage(bot.parentCtx, &pb.RequestSendMessage{
		Peer:    dest,
		Rid:     randomId,
		Message: &message,
	})
	if err != nil {
		//log.Errorf("Error while sendMessage: %v, %v, %d", err, dest, randomId)
	} else {
		//log.Debugf(
		//	"Message (%d, %s) sent in %fs from %d to %d",
		//	time.Since(sendMessageStart).Seconds(),
		//	randomId,
		//	res.Mid.String(),
		//	bot.user.Id,
		//	dest.Id,
		//)
		//bot.lock.Lock()
		//bot.dialogs[dest.Id] = struct {
		//	*pb.OutPeer
		//	int64
		//}{dest, res.Date}
		//bot.lock.Unlock()
	}
	return res, err
}