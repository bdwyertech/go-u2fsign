// Encoding: UTF-8

package main

import (
	"encoding/json"
	"flag"

	"github.com/bdwyertech/go-u2fsign/u2fplugin"

	"github.com/hashicorp/go-plugin"

	u2f "github.com/marshallbrekka/go-u2fhost"
)

var pluginFlag bool

func init() {
	flag.BoolVar(&pluginFlag, "plugin", false, "Run as a Go Plugin")
}

type AuthPlugin struct{}

func (*AuthPlugin) Authenticate(req string) (resp []byte, err error) {
	var authR u2f.AuthenticateRequest
	if err = json.Unmarshal([]byte(req), &authR); err != nil {
		return
	}

	var authResp *u2f.AuthenticateResponse
	authResp, err = u2fAuth(&authR, u2f.Devices())
	if err != nil {
		return
	}

	resp, err = json.Marshal(authResp)
	return
}

func RunAsPlugin() {
	var pluginMap = map[string]plugin.Plugin{
		"authenticator": &u2fplugin.AuthPlugin{Impl: &AuthPlugin{}},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: u2fplugin.Handshake,
		Plugins:         pluginMap,
	})
}
