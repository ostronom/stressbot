package bot

import (
	pb "github.com/dialogs/stressbot/gateway"
	"math/rand"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Bot struct {
	AppId       int32
	ApiKey      string
	AppTitle    string
	DeviceTitle string
	Name        string
	Emails      []string
	Phones      []int64

	parentCtx   context.Context
	authContext context.Context
	conn        *grpc.ClientConn
	Self        *pb.User
	token       string
	authId      int64

	// gRPC clients
	registration   pb.RegistrationClient
	authentication pb.AuthenticationClient
	users          pb.UsersClient
	profile        pb.ProfileClient
	contacts       pb.ContactsClient
	messaging      pb.MessagingClient
	groups         pb.GroupsClient
	files          pb.MediaAndFilesClient
	seqUpdates     pb.SequenceAndUpdatesClient
	weaks          pb.ObsoleteClient
}

func New(ctx context.Context, conn *grpc.ClientConn, appId int32, appTitle, apiKey, deviceTile, name string) (*Bot, error) {
	b := &Bot{parentCtx: ctx, conn: conn, AppId: appId, AppTitle: appTitle, ApiKey: apiKey, DeviceTitle: deviceTile,
		Name: name, Emails: make([]string, 0), Phones: make([]int64, 0)}
	b.registration = pb.NewRegistrationClient(conn)
	b.authentication = pb.NewAuthenticationClient(conn)
	b.contacts = pb.NewContactsClient(conn)
	b.users = pb.NewUsersClient(conn)
	b.profile = pb.NewProfileClient(conn)
	b.messaging = pb.NewMessagingClient(conn)
	b.groups = pb.NewGroupsClient(conn)
	b.files = pb.NewMediaAndFilesClient(conn)
	b.seqUpdates = pb.NewSequenceAndUpdatesClient(conn)
	b.weaks = pb.NewObsoleteClient(conn)

	reg, err := b.registration.RegisterDevice(ctx, &pb.RequestRegisterDevice{
		AppId: appId, AppTitle: b.AppTitle, DeviceTitle: deviceTile,
	})
	if err != nil {
		return nil, err
	}

	b.authContext = metadata.AppendToOutgoingContext(b.parentCtx, "x-auth-ticket", reg.Token)
	b.token = reg.Token
	b.authId = reg.AuthId

	return b, nil
}
