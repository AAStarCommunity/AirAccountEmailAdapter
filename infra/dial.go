package infra

import "github.com/knadh/go-pop3"

func Dial(server string, port int, user, password string) (*pop3.Conn, error) {
	pop3Client := pop3.New(pop3.Opt{
		Host:       server,
		Port:       port,
		TLSEnabled: true,
	})

	c, err := pop3Client.NewConn()
	if err != nil {
		return nil, err
	}

	// Authenticate.
	if err := c.Auth(user, password); err != nil {
		return nil, err
	}

	return c, nil
}
