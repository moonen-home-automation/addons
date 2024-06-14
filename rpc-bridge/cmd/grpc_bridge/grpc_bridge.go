package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/moonen-home-automation/addons/grpc_bridge/internal/server"
	"github.com/moonen-home-automation/addons/grpc_bridge/pkg/proto"
	hass_client "github.com/moonen-home-automation/hass-ws-client"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	_, err := hass_client.InitializeAppInstance(hass_client.InitializeAppRequest{URL: "ws://supervisor/core/websocket", Secure: false, HAAuthToken: os.Getenv("SUPERVISOR_TOKEN")})
	if err != nil {
		log.Error("Hass client error", "error", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()

	es := server.NewEventsServer(log)
	proto.RegisterEventsServer(gs, es)

	ss := server.NewServicesServer(log)
	proto.RegisterServicesServer(gs, ss)

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
