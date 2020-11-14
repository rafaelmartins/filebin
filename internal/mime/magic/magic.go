package magic

import (
	"errors"
	"os/exec"
	"strings"
)

var (
	ErrNotSupported = errors.New("magic: not supported")
	fileBin         string
	fileErr         error
)

func init() {
	fileBin, fileErr = exec.LookPath("file")
}

func Detect(data []byte) (string, error) {
	if fileErr != nil {
		return "", ErrNotSupported
	}

	cmd := exec.Command(fileBin, "--mime-type", "--brief", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		stdin.Write(data)
	}()

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
