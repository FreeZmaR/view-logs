package reader

import (
	"io"
)

type Reader interface {
	Read() ([]byte, error)
	GetCommand() string
	Close() error
	initReader(reader io.Reader, run RunCommandFN, interrupt InterruptCommandFN)
}

type (
	RunCommandFN       func(command string) error
	InterruptCommandFN func() error
)

func InitReader(r Reader, reader io.Reader, run RunCommandFN, interrupt InterruptCommandFN) {
	r.initReader(reader, run, interrupt)
}
