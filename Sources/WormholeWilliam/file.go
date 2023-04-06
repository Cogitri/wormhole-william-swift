package WormholeWilliam

import (
	"context"
	"errors"
	"fmt"
	"io"
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

func ReceiveFile(code string, targetFilePath string) error {
	var c wormhole.Client

	ctx := context.Background()
	fileInfo, err := c.Receive(ctx, code)
	if err != nil {
		return err
	}

	if fileInfo.Type != wormhole.TransferFile {
		return fmt.Errorf("expected a file transfer but got type %s", fileInfo.Type)
	}

	f, err := os.Open(targetFilePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, fileInfo)

	if err != nil {
		return err
	}

	return nil
}
