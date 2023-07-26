package email

import (
	"crypto/tls"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/zhaixinlong/go-utils/zip"
)

type EmailConfig struct {
	Sender            string
	AuthorizationCode string
	SmtpServer        string
	SendAddr          string
}

type SendEmailInfo struct {
	ReplyTo []string
	From    string
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    string
	HTML    string
	Files   []string
}

type EmailSender struct {
	sender            string
	authorizationCode string
	smtpServer        string
	sendAddr          string
}

func NewEmail(ec EmailConfig) *EmailSender {
	return &EmailSender{
		sender:            ec.Sender,
		sendAddr:          ec.SendAddr,
		smtpServer:        ec.SmtpServer,
		authorizationCode: ec.AuthorizationCode,
	}
}

func (es *EmailSender) SendMail(info SendEmailInfo) error {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = info.From
	// 设置接收方的邮箱
	e.To = info.To
	//设置主题
	e.Subject = info.Subject

	for _, v := range info.Files {
		zipFileName := zip.ZipFile(v)
		if _, err := e.AttachFile(zipFileName); err != nil {
			log.Printf("send email AttachFile err, file:%s \n", v)
			return err
		}
	}

	//设置文件发送的内容
	e.HTML = []byte(info.HTML)
	log.Printf("send email begin: %+v \n", info.To)

	//设置服务器相关的配置
	err := e.SendWithTLS(es.sendAddr, smtp.PlainAuth("", es.sender, es.authorizationCode, es.smtpServer), &tls.Config{ServerName: es.smtpServer})
	if err != nil {
		//发送失败
		log.Printf("send email error to: %+v, err: %+v\n", info.To, err)
		return err
	}
	// 发送成功
	log.Printf("send email to: %+v success \n", info.To)
	return nil
}
