package dingtalk

import (
	dk "github.com/blinkbean/dingtalk"
)

type MsgInfo struct {
	Title        string `json:"title"`
	Message      string `json:"message"`
	NotWithAtAll bool   `json:"not_with_at_all"`
}

type Config struct {
	AccessKey    string
	AccessSecret string
}

type DingTalkClient struct {
	dingTalk *dk.DingTalk
	cfg      *Config
}

func NewDingTalkClientMgr(cfg *Config) (*DingTalkClient, error) {
	dingTalk := dk.InitDingTalkWithSecret(cfg.AccessKey, cfg.AccessSecret)
	return &DingTalkClient{
		dingTalk: dingTalk,
		cfg:      cfg,
	}, nil
}

func (d *DingTalkClient) SendMarkDown(msg *MsgInfo) error {
	if msg.NotWithAtAll {
		return d.dingTalk.SendMarkDownMessage(msg.Title, msg.Message)
	}
	return d.dingTalk.SendMarkDownMessage(msg.Title, msg.Message, dk.WithAtAll())
}
