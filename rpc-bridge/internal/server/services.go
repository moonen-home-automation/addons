package server

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/go-hclog"
	"github.com/moonen-home-automation/addons/grpc_bridge/pkg/proto"
	hasswsclient "github.com/moonen-home-automation/hass-ws-client"
	"github.com/moonen-home-automation/hass-ws-client/pkg/services"
)

type ServicesServer struct {
	log hclog.Logger
	proto.UnimplementedServicesServer
}

func NewServicesServer(l hclog.Logger) *ServicesServer {
	return &ServicesServer{l, proto.UnimplementedServicesServer{}}
}

func (s *ServicesServer) CallService(ctx context.Context, sc *proto.ServiceCall) (*proto.ServiceResponse, error) {
	hassc := hasswsclient.GetAppInstance()

	data := make(map[string]interface{})
	_ = json.Unmarshal([]byte(sc.JsonData), &data)

	target := services.ServiceTarget{
		AreaID:   sc.GetAreaId(),
		DeviceID: sc.GetDeviceId(),
		EntityID: sc.GetEntityId(),
		LabelID:  sc.GetLabelId(),
	}

	serviceCall := services.NewServiceCall(sc.Domain, sc.Service, data, target, sc.Returns)

	resp, err := hassc.ServiceCaller.Call(serviceCall)
	if err != nil {
		return nil, err
	}

	jsonResp, _ := json.Marshal(resp.Result.Response)

	return &proto.ServiceResponse{
		JsonData: string(jsonResp),
	}, nil
}
