package WormholeWilliam

import (
	"context"
	"errors"
	"os"
	"path"

	"github.com/psanford/wormhole-william/wormhole"
)

func PrepareSendFile(s *SenderContext, filePath string) error {
	var base = path.Base(filePath)

	if base == "." || base == "/" {
		return errors.New("invalid file path")
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	var c wormhole.Client
	ctx := context.Background()

	code, status, err := c.SendFile(ctx, base, f)
	if err != nil {
		return err
	}

	s.sendResult = &status
	s.code = &code

	return nil
}
