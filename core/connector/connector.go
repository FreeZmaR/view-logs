package connector

type Connector interface {
	Connect() error
	Ping() (int, error)
	Close()
}
