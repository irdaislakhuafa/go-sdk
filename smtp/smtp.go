package smtp

import (
	"context"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type (
	Interface interface {
		Send(ctx context.Context, msgs ...Message) error
	}

	MessageHeaders struct {
		From    []string
		To      []string
		Subject []string
		Extras  map[string][]string
	}

	MessageBody struct {
		ContentType string
		Template    string
		Values      map[string]any
	}

	Message struct {
		Headers MessageHeaders
		Body    MessageBody
	}

	Provider string

	Config struct {
		Username string
		Password string
		Host     string
		Port     int64
		Provider Provider
	}
)

const (
	ProviderGoMail = Provider("gomail")
)

func (c *Config) Default() {
	if c.Provider == "" {
		c.Provider = ProviderGoMail
	}
}

func Init(cfg Config) (Interface, error) {
	cfg.Default()

	if cfg.Provider == ProviderGoMail {
		return InitGoMail(cfg), nil
	} else {
		return nil, errors.NewWithCode(codes.CodeNotImplemented, "smtp provider not supported!")
	}
}
