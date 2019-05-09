package main

import (
	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"github.com/zgs225/youdao"
)

type Config struct {
	AppID     string
	AppSecret string
}

var client *youdao.Client

func init() {
	var config Config
	if _, err := toml.DecodeFile("conf.toml", &config); err != nil {
		panic(err)
	}
	c := &youdao.Client{
		AppID:     config.AppID,
		AppSecret: config.AppSecret,
	}
	c.SetFrom(youdao.LEnglish)
	c.SetTo(youdao.LChinese)

	client = c
}

func translate(query string) (string, error) {
	r, err := client.Query(query)
	if err != nil {
		return "", err
	}
	return spew.Sdump(r.Translation) + "\n" + spew.Sdump(r.Basic), nil
}
