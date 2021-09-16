package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/bdwyertech/go-u2fsign/u2fplugin"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func main() {
	var authR interface{}
	err := json.NewDecoder(os.Stdin).Decode(&authR)
	if err != nil {
		log.Fatal(err)
	}
	// Back to String
	authJson, err := json.Marshal(authR)
	if err != nil {
		log.Fatal(err)
	}

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		Cmd:             exec.Command("../../go-u2fsign", "-plugin"),
		HandshakeConfig: u2fplugin.Handshake,
		Plugins:         u2fplugin.PluginMap,
		Logger: hclog.New(&hclog.LoggerOptions{
			// Name:   "u2fsign",
			Output: os.Stdout,
			Level:  hclog.Debug,
			// Level: hclog.Info,
		}),
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("authenticator")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	auth := raw.(*u2fplugin.RPCClient)
	resp, err := auth.Authenticate(string(authJson))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
}
