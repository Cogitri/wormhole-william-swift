package WormholeWilliam

import (
	"context"
	"fmt"
	"io/ioutil"

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

func ReceiveText(code string) (string, error) {
	var c wormhole.Client

	ctx := context.Background()
	msg, err := c.Receive(ctx, code)
	if err != nil {
		return "", err
	}

	if msg.Type != wormhole.TransferText {
		return "", fmt.Errorf("expected a text message but got type %s", msg.Type)
	}

	msgBody, err := ioutil.ReadAll(msg)
	if err != nil {
		return "", err
	}

	return string(msgBody), nil
}
