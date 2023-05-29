package connector

import "github.com/FreeZmaR/view-logs/core/reader"

type Connector interface {
	Connect() error
	Ping() (int, error)
	SetReader(reader reader.Reader) error
	Close()
}
