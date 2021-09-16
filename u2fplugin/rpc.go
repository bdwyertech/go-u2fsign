package u2fplugin

import (
	"net/rpc"
)

// RPCClient is an implementation of Authenticator that talks over RPC.
type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) Authenticate(key string) ([]byte, error) {
	var resp []byte
	err := m.client.Call("Plugin.Authenticate", key, &resp)
	return resp, err
}

// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl AuthInterface
}

func (m *RPCServer) Authenticate(req string, resp *[]byte) error {
	v, err := m.Impl.Authenticate(req)
	*resp = v
	return err
}
