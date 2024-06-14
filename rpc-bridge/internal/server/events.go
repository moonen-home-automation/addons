package server

import (
	"github.com/hashicorp/go-hclog"
	"github.com/moonen-home-automation/addons/grpc_bridge/pkg/proto"
	hassclient "github.com/moonen-home-automation/hass-ws-client"
	"google.golang.org/protobuf/types/known/anypb"
)

type EventsServer struct {
	log hclog.Logger
	proto.UnimplementedEventsServer
}

func NewEventsServer(l hclog.Logger) *EventsServer {
	return &EventsServer{l, proto.UnimplementedEventsServer{}}
}

func (e *EventsServer) Subscribe(rr *proto.EventSubscribeRequest, src proto.Events_SubscribeServer) error {
	hassc := hassclient.GetAppInstance()
	channel := make(chan hassclient.EventData, 10)
	eventListener := hassclient.EventListener{EventType: "zha_event"}
	hassc.RegisterEventListener(eventListener)
	go hassc.ListenForEvents(eventListener, channel)
	e.log.Info("Registered event listener", "event", eventListener.EventType)
	for {
		data, ok := <-channel
		if !ok {
			break
		}

		err := src.Send(&proto.Event{
			Type: data.Type,
			Data: &anypb.Any{Value: data.RawEventJSON},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
