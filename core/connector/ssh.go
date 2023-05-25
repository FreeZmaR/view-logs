package connector

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"strconv"
	"time"
)

type SSH struct {
	connect      *SSHConnect
	config       *SSHConfig
	proxy        []SHHProxyConnect
	proxyConfigs []*SSHConfig
}

type SSHConnect struct {
	connect *ssh.Client
	session *ssh.Session
	config  *SSHConfig
}

type SHHProxyConnect struct {
	proxyConnect *ssh.Client
	config       *SSHConfig
}

type SSHConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

var _ Connector = (*SSH)(nil)

func NewSSH(main *SSHConfig, proxy ...*SSHConfig) *SSH {
	return &SSH{
		config:       main,
		proxyConfigs: proxy,
	}
}

func (connect *SSH) Connect() error {
	if err := connect.makeProxy(); err != nil {
		connect.realizeResources()

		return err
	}

	connect.connect = &SSHConnect{
		config: connect.config,
	}

	if len(connect.proxy) == 0 {
		auth := []ssh.AuthMethod{
			ssh.Password(connect.config.Password),
		}

		var err error

		connect.connect.connect, err = ssh.Dial(
			"tcp",
			connect.config.Host+":"+strconv.Itoa(connect.config.Port),
			&ssh.ClientConfig{
				User:            connect.config.User,
				Auth:            auth,
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
		)
		if err != nil {
			return err
		}

		connect.connect.session, err = connect.connect.connect.NewSession()
		if err != nil {
			connect.realizeResources()

			return err
		}

		return nil
	}

	lastProxy := connect.proxy[len(connect.proxy)-1]

	conn, err := lastProxy.proxyConnect.Dial("tcp", connect.config.Host+":"+strconv.Itoa(connect.config.Port))
	if err != nil {
		connect.realizeResources()

		return err
	}

	//agentClient := agent.NewClient(lastProxy.proxyConnect)
	auths := []ssh.AuthMethod{
		//ssh.PublicKeysCallback(agentClient.Signers),
		ssh.Password(connect.config.Password),
	}

	con, chans, req, err := ssh.NewClientConn(
		conn,
		connect.config.Host,
		&ssh.ClientConfig{
			User:            connect.config.User,
			Auth:            auths,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         30 * time.Second,
		},
	)
	if err != nil {
		connect.realizeResources()

		return err
	}

	client := ssh.NewClient(con, chans, req)
	connect.connect.connect = client

	connect.connect.session, err = client.NewSession()
	if err != nil {
		connect.realizeResources()

		return err
	}

	return nil
}

func (connect *SSH) Ping() (int, error) {
	if connect.connect == nil {
		return 0, nil
	}

	if connect.connect.session == nil {
		return 0, nil
	}

	t := time.Now()

	if err := connect.connect.session.Run("echo 1"); err != nil {
		return 0, err
	}

	return int(time.Since(t).Seconds() * 1000), nil
}

func (connect *SSH) Close() {
	connect.realizeResources()
}

func (connect *SSH) makeProxy() error {
	if len(connect.proxyConfigs) == 0 {
		return nil
	}

	connect.proxy = make([]SHHProxyConnect, len(connect.proxyConfigs))
	for i, proxyConfig := range connect.proxyConfigs {
		auth := []ssh.AuthMethod{
			ssh.Password(proxyConfig.Password),
		}

		proxyConnect, err := ssh.Dial(
			"tcp",
			connect.config.Host+":"+strconv.Itoa(connect.config.Port),
			&ssh.ClientConfig{
				User:            connect.config.User,
				Auth:            auth,
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
		)
		if err != nil {
			return err
		}

		connect.proxy[i] = SHHProxyConnect{
			proxyConnect: proxyConnect,
			config:       proxyConfig,
		}
	}

	return nil
}

func (connect *SSH) realizeResources() {
	if connect.connect != nil {
		if connect.connect.session != nil {
			if err := connect.connect.session.Close(); err != nil && !errors.Is(err, io.EOF) {
				log.Print("Error closing session: ", err)
			}
		}

		if connect.connect.connect != nil {
			if err := connect.connect.connect.Close(); err != nil {
				log.Print("Error closing connect: ", err)
			}
		}
	}

	if len(connect.proxy) != 0 {
		for i := len(connect.proxy) - 1; i >= 0; i-- {
			proxy := connect.proxy[i]

			if proxy.proxyConnect != nil {
				if err := proxy.proxyConnect.Close(); err != nil {
					log.Print("Error closing proxy client: ", err)
				}
			}
		}
	}
}
