package smtp

import (
	mail "github.com/go-mail/gomail"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type GoMailInterface interface {
	DialAndSend(messages ...*mail.Message) error
}

type gomailImpl struct {
	dialer *mail.Dialer
}

func InitGoMail(cfg Config) GoMailInterface {
	result := &gomailImpl{
		dialer: mail.NewDialer(cfg.Host, int(cfg.Port), cfg.Username, cfg.Password),
	}
	return result
}

func (g *gomailImpl) DialAndSend(messages ...*mail.Message) error {
	if err := g.dialer.DialAndSend(messages...); err != nil {
		return errors.NewWithCode(codes.CodeSMTP, "failed to send messages, %v", err)
	}
	return nil
}
