package bot

import (
	pb "github.com/dialogs/stressbot/gateway"
	"math/rand"
)

func (b *Bot) CreateGroup(title string, members []*pb.UserOutPeer) (*pb.ResponseCreateGroup, error) {
	randomId := rand.Int63()
	res, err := b.groups.CreateGroup(b.authContext, &pb.RequestCreateGroup{
		Rid:       randomId,
		Title:     title,
		GroupType: pb.GroupType_GROUPTYPE_GROUP,
		Users:     members,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *Bot) GetGroupInviteUrl(peer *pb.GroupOutPeer) (*pb.ResponseInviteUrl, error) {
	return b.groups.GetGroupInviteUrl(b.authContext, &pb.RequestGetGroupInviteUrl{GroupPeer: peer})
}

func (b *Bot) JoinGroupByUrl(url string) (*pb.ResponseJoinGroup, error) {
	return b.groups.JoinGroup(b.authContext, &pb.RequestJoinGroup{Token: url})
}