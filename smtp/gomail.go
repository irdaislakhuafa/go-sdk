package smtp

import (
	"context"

	mail "github.com/go-mail/gomail"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/strformat"
)

type gomail struct {
	dialer *mail.Dialer
}

func InitGoMail(cfg Config) Interface {
	result := &gomail{
		dialer: mail.NewDialer(cfg.Host, int(cfg.Port), cfg.Username, cfg.Password),
	}
	return result
}

func (g *gomail) Send(ctx context.Context, msgs ...Message) error {
	messages := []*mail.Message{}
	for _, msg := range msgs {
		m := mail.NewMessage()
		m.SetHeaders(map[string][]string{
			"From":    msg.Headers.From,
			"To":      msg.Headers.To,
			"Subject": msg.Headers.Subject,
		})
		for hk, hv := range msg.Headers.Extras {
			m.SetHeader(hk, hv...)
		}

		body, err := strformat.T(msg.Body.Template, msg.Body.Values)
		if err != nil {
			return errors.NewWithCode(codes.CodeStrTemplateExecuteErr, err.Error())
		}

		m.SetBody(msg.Body.ContentType, body)
		messages = append(messages, m)
	}

	if err := g.dialer.DialAndSend(messages...); err != nil {
		return errors.NewWithCode(codes.CodeSMTP, "failed to send messages, %v", err)
	}
	return nil
}
