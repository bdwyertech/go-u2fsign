// Package shared contains shared data between the host and plugins.
package u2fplugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "u2fplugin",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"authenticator": &AuthPlugin{},
}

// KV is the interface that we're exposing as a plugin.
type AuthInterface interface {
	Authenticate(req string) ([]byte, error)
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
type AuthPlugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl AuthInterface
}

func (p *AuthPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (*AuthPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}
