package bot

import (
	"errors"
	"fmt"
	"strings"
)
import pb "github.com/dialogs/stressbot/gateway"

var ErrRetry = errors.New("RETRY")

func (b *Bot) signUp(hash string) (*pb.ResponseAuth, error) {
	return b.authentication.SignUp(b.authContext, &pb.RequestSignUp{
		TransactionHash: hash,
		Name:            b.Name,
	})
}

func (b *Bot) getSelfData() error {
	fu, err := b.users.LoadFullUsers(b.authContext, &pb.RequestLoadFullUsers{
		UserPeers: []*pb.UserOutPeer{&pb.UserOutPeer{Uid: b.Self.Id, AccessHash: b.Self.AccessHash}},
	})
	if err != nil {
		return err
	}
	if len(fu.FullUsers) != 1 {
		return fmt.Errorf("Requested 1 user, got %d", len(fu.FullUsers))
	}
	for _, record := range fu.FullUsers[0].ContactInfo {
		if record.Type == pb.ContactType_CONTACTTYPE_EMAIL {
			b.Emails = append(b.Emails, record.StringValue.Value)
		}
		if record.Type == pb.ContactType_CONTACTTYPE_PHONE {
			b.Phones = append(b.Phones, record.LongValue.Value)
		}
	}
	return nil
}

func (b *Bot) AnonymousAuth() (res *pb.ResponseAuth, err error) {
	res, err = b.authentication.StartAnonymousAuth(b.authContext, &pb.RequestStartAnonymousAuth{
		Name: b.Name, AppId: b.AppId, ApiKey: b.ApiKey, DeviceTitle: b.DeviceTitle,
	})
	//if err == nil {
	//	err = b.updatesLoop()
	//}
	b.Self = res.User
	b.getSelfData()
	return
}

func (b *Bot) AuthorizeByPhone(phone int64, code string) (res *pb.ResponseAuth, err error) {
	auth, err := b.authentication.StartPhoneAuth(b.authContext, &pb.RequestStartPhoneAuth{
		PhoneNumber: phone,
		AppId:       b.AppId,
		ApiKey:      b.ApiKey,
		DeviceTitle: b.DeviceTitle,
	})
	if err != nil {
		return
	}
	res, err = b.authentication.ValidateCode(b.authContext, &pb.RequestValidateCode{
		TransactionHash: auth.GetTransactionHash(),
		Code:            code,
	})
	if err == nil {
		return nil, ErrRetry
	}
	if strings.Contains(err.Error(), "Unnoccupied phone number") {
		res, err = b.signUp(auth.GetTransactionHash())
		if err != nil {
			return
		}
	}
	b.Self = res.User
	b.getSelfData()
	//err = b.updatesLoop()
	return
}
