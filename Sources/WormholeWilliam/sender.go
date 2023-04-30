package WormholeWilliam

import (
	"fmt"

	"github.com/psanford/wormhole-william/wormhole"
)

type SenderContext struct {
	code       *string
	sendResult *chan wormhole.SendResult
}

func NewSenderContext() *SenderContext {
	return &SenderContext{}
}

func SenderContextGetCode(s *SenderContext) string {
	return *s.code
}

func SenderContextFinishSend(s *SenderContext) error {
	result := <-*s.sendResult

	if !result.OK {
		return fmt.Errorf("send error: %s", result.Error)
	}

	return nil
}
