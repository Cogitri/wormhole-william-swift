package WormholeWilliam

import (
	"context"

	"github.com/psanford/wormhole-william/wormhole"
)

func PrepareSendText(s *SenderContext, msg string) error {
	var c wormhole.Client

	ctx := context.Background()

	code, status, err := c.SendText(ctx, msg)
	if err != nil {
		return err
	}

	s.sendResult = &status
	s.code = &code

	return nil
}
