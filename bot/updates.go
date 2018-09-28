package bot

import "github.com/golang/protobuf/ptypes/empty"
import pb "github.com/dialogs/stressbot/gateway"

func (b *Bot) safeUpdatesLoop(upds pb.SequenceAndUpdates_SeqUpdatesClient) {
	for {
		box, err := upds.Recv()
		if err != nil {
			return
		}
		if box.UnboxedUpdate == nil {
			continue
		}
	}
}

func (b *Bot) updatesLoop() error {
	upds, err := b.seqUpdates.SeqUpdates(b.authContext, &empty.Empty{})
	if err != nil {
		return err
	}
	go b.safeUpdatesLoop(upds)
	return nil
}
