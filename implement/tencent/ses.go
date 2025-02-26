//go:build tencent && ses

package tencent

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/SliverHornTrident/components/config"
	"io"
	"net"
	"net/smtp"
	"net/url"
)

type Ses struct {
	Auth   smtp.Auth
	Config config.TencentSes
}

func NewSes(config config.TencentSes) *Ses {
	return &Ses{Auth: config.Auth(), Config: config}
}

type SesMessage struct {
	Body    string
	Subject string
	Tos     []string
	Header  url.Values
}

func NewSesMessage(body string, subject string, tos []string) *SesMessage {
	return &SesMessage{Body: body, Subject: subject, Tos: tos}
}

func (s *SesMessage) Messages() []byte {
	builder := new(bytes.Buffer)
	for key, value := range s.Header {
		builder.WriteString(key)
		builder.WriteString(": ")
		if len(value) >= 1 {
			builder.WriteString(value[0])
		}
		builder.WriteString("\r\n")
	}
	builder.WriteString("\r\n")
	builder.WriteString(s.Body)
	return builder.Bytes()
}

// DialAndSend send email
func (s *Ses) DialAndSend(ctx context.Context, message *SesMessage) error {
	var client *smtp.Client
	if s.Config.TLS {
		// ssl_protocols    TLSv1 TLSv1.1 TLSv1.2;
		// ssl_ciphers      AES128-SHA:AES256-SHA:RC4-SHA:DES-CBC3-SHA:RC4-MD5;
		conn, err := tls.Dial("tcp", s.Config.Address(), &tls.Config{
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_RC4_128_SHA,
				tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
				// tls.TLS_RSA_WITH_RC4_128_MD5,
			},
		})
		if err != nil {
			return err
		}
		client, err = smtp.NewClient(conn, s.Config.Host)
		if err != nil {
			return err
		}
	} else {
		conn, err := net.Dial("tcp", s.Config.Address())
		if err != nil {
			return err
		}
		client, err = smtp.NewClient(conn, s.Config.Host)
		if err != nil {
			return err
		}
	}
	defer func() {
		_ = client.Close()
	}()
	if s.Auth != nil {
		ok, _ := client.Extension("AUTH")
		if ok {
			err := client.Auth(s.Auth)
			if err != nil {
				return err
			}
		}
	}
	message.Header = s.Config.Header(message.Subject, message.Tos)
	err := client.Mail(s.Config.Username)
	if err != nil {
		return err
	}
	for i := 0; i < len(message.Tos); i++ {
		err = client.Rcpt(message.Tos[i])
		if err != nil {
			return err
		}
	}
	var writeCloser io.WriteCloser
	writeCloser, err = client.Data()
	if err != nil {
		return err
	}
	messages := message.Messages()
	_, err = writeCloser.Write(messages)
	if err != nil {
		return err
	}
	err = writeCloser.Close()
	if err != nil {
		return err
	}
	return client.Quit()
}
