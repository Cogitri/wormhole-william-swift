package WormholeWilliam

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/psanford/wormhole-william/wormhole"
)

type ReceiverContext struct {
	msg *wormhole.IncomingMessage
}

func NewReceiverContext() *ReceiverContext {
	return &ReceiverContext{}
}

func ReceiverContextInitReceive(r *ReceiverContext, code string) error {
	var c wormhole.Client

	ctx := context.Background()
	msg, err := c.Receive(ctx, code)
	if err != nil {
		return err
	}
	r.msg = msg
	return nil
}

func ReceiverContextGetName(r *ReceiverContext) string {
	if r.msg == nil {
		return ""
	}
	return r.msg.Name
}

func ReceiverContextGetSize(r *ReceiverContext) int64 {
	if r.msg == nil {
		return 0
	}
	return r.msg.TransferBytes64
}

func ReceiverContextGetType(r *ReceiverContext) string {
	if r.msg == nil {
		return ""
	}
	return r.msg.Type.String()
}

func ReceiverContextReceiveFile(r *ReceiverContext, path string) error {
	if r.msg.Type != wormhole.TransferFile {
		return fmt.Errorf("expected a file transfer but got type %s", r.msg.Type)
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r.msg)

	if err != nil {
		return err
	}

	return nil
}

func ReceiverContextReceiveText(r *ReceiverContext) (string, error) {
	if r.msg.Type != wormhole.TransferText {
		return "", fmt.Errorf("expected a text message but got type %s", r.msg.Type)
	}

	msgBody, err := ioutil.ReadAll(r.msg)
	if err != nil {
		return "", err
	}

	return string(msgBody), nil
}
