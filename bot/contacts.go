package bot

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"strconv"
)
import pb "github.com/dialogs/stressbot/gateway"

func (b *Bot) ImportPhoneContacts(phones []int64) ([]*pb.UserOutPeer, error) {
	toImport := make([]*pb.PhoneToImport, len(phones))
	for i, phone := range phones {
		toImport[i] = &pb.PhoneToImport{
			PhoneNumber: phone,
			Name:        &wrappers.StringValue{Value: string(strconv.AppendInt([]byte("Contact_"), phone, 10))},
		}
	}

	res, err := b.contacts.ImportContacts(b.authContext, &pb.RequestImportContacts{Phones: toImport})
	if err != nil {
		return nil, err
	}
	return res.GetUserPeers(), nil
}

func (b *Bot) ImportEmailContacts(emails []string) ([]*pb.UserOutPeer, error) {
	toImport := make([]*pb.EmailToImport, len(emails))
	for i, email := range emails {
		toImport[i] = &pb.EmailToImport{
			Email: email,
			Name:  &wrappers.StringValue{Value: "Contact_" + email},
		}
	}

	res, err := b.contacts.ImportContacts(b.authContext, &pb.RequestImportContacts{Emails: toImport})
	if err != nil {
		return nil, err
	}
	return res.GetUserPeers(), nil
}

func (b *Bot) SearchContact(query string) (*pb.ResponseSearchContacts, error) {
	return b.contacts.SearchContacts(b.authContext, &pb.RequestSearchContacts{Request: query})
}

func (b *Bot) AddContact(peer *pb.UserOutPeer) (*pb.ResponseSeq, error) {
	return b.contacts.AddContact(b.authContext, &pb.RequestAddContact{Uid: peer.Uid, AccessHash: peer.AccessHash})
}
