package reader

import (
	"fmt"
	"io"
	"time"
)

const (
	CatCommand  ConsoleCommand = "cat"
	TailCommand ConsoleCommand = "tail"
	HeadCommand ConsoleCommand = "head"
	DDCommand   ConsoleCommand = "dd"
)

type ConsoleCommand string

type ConsoleReader struct {
	commandType        ConsoleCommand
	runCommandFM       RunCommandFN
	InterruptCommandFN InterruptCommandFN
	scanner            io.Reader
	path               string
	currentCursor      int
	nextCursor         int
	cursorStep         int
	isCompleted        bool
	isStream           bool
	isReadable         bool
}

var _ Reader = (*ConsoleReader)(nil)

func NewConsoleReader(path string, command ConsoleCommand) *ConsoleReader {
	r := &ConsoleReader{
		commandType: command,
		path:        path,
		isCompleted: false,
		isStream:    false,
	}

	switch r.commandType {
	case CatCommand:
		r.isCompleted = true
	case TailCommand:
	case HeadCommand:
		r.currentCursor = 1
		r.cursorStep = r.currentCursor
		r.nextCursor = r.currentCursor + r.cursorStep
	case DDCommand:
		r.cursorStep = 1
		r.currentCursor = 0
		r.nextCursor = r.currentCursor + r.cursorStep
	}

	return r
}

func (r *ConsoleReader) Read() ([]byte, error) {
	var err error

	if r.isStream {
		return io.ReadAll(r.scanner)
	}

	switch r.commandType {
	case CatCommand:
		err = r.runCommandFM(r.GetCommand())
	case TailCommand:
		//TODO: fix this
		r.isStream = true
		errChan := make(chan error)

		go func() {
			errCmd := r.runCommandFM(r.GetCommand())
			errChan <- errCmd
		}()

		select {
		case err = <-errChan:
			return nil, err
		case <-time.After(1 * time.Second):
			return io.ReadAll(r.scanner)
		}
	case HeadCommand:
		err = r.runCommandFM(r.GetCommand())
	case DDCommand:
		err = r.runCommandFM(r.GetCommand())
	}

	if err != nil {
		return nil, err
	}

	if r.isCompleted {
		return io.ReadAll(r.scanner)
	}

	r.currentCursor = r.nextCursor
	r.nextCursor = r.currentCursor + r.cursorStep

	return io.ReadAll(r.scanner)
}

func (r *ConsoleReader) GetCommand() string {
	switch r.commandType {
	case CatCommand:
		return fmt.Sprintf("%s %s", r.commandType, r.path)
	case TailCommand:
		return fmt.Sprintf("%s -f  %s", r.commandType, r.path)
	case HeadCommand:
		return fmt.Sprintf("%s -n %d %s", r.commandType, r.currentCursor, r.path)
	case DDCommand:
		return fmt.Sprintf(
			"dd if=%s ibs=1024 count=1 skip=%d",
			r.path,
			r.currentCursor,
		)
	default:
		return string(r.commandType) + " " + r.path
	}
}

func (r *ConsoleReader) Close() error {
	if r.commandType == TailCommand {
		return r.InterruptCommandFN()
	}

	return nil
}

func (r *ConsoleReader) initReader(scanner io.Reader, run RunCommandFN, interrupt InterruptCommandFN) {
	r.scanner = scanner
	r.runCommandFM = run
	r.InterruptCommandFN = interrupt
}
