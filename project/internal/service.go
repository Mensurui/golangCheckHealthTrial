package internal

import (
	"context"
	"time"

	protos "github.com/Mensurui/golangCheckHealthTrial/protos/golang"
	"github.com/hashicorp/go-hclog"
)

type Service struct {
	protos.UnimplementedServiceServer
	hcl hclog.Logger
}

func NewService(hcl hclog.Logger) *Service {
	return &Service{
		hcl: hcl,
	}
}

func (service *Service) Check(ctx context.Context, req *protos.HealthCheckRequest) (*protos.HealthCheckResponse, error) {
	status := service.determineStatus(req.Service)
	return &protos.HealthCheckResponse{
		Status: status,
	}, nil
}

func (service *Service) Wait(req *protos.HealthCheckRequest, stream protos.Service_WaitServer) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			status := service.determineStatus(req.Service)
			resp := &protos.HealthCheckResponse{Status: status}

			if err := stream.Send(resp); err != nil {
				service.hcl.Error("Failed to send health check update", "error", err)
				return err
			}

		case <-stream.Context().Done():
			service.hcl.Info("Client stopped health check streaming")
			return nil
		}
	}
}

func (service *Service) GetUsername(ctx context.Context, req *protos.GetUsernameRequest) (*protos.GetUsernameResponse, error) {
	_ = req.GetId()

	return &protos.GetUsernameResponse{
		Firstname: "Mensur",
		Lastname:  "Khalid",
	}, nil
}

func (service *Service) State(ctx context.Context, req *protos.StateRequest) (*protos.StateResponse, error) {
	_ = req.GetTemprature()
	return &protos.StateResponse{
		Status: "Solid",
	}, nil
}

func (service *Service) determineStatus(serviceName string) protos.HealthCheckResponse_ServingStatus {
	switch serviceName {
	case "":
		return protos.HealthCheckResponse_SERVING
	case "grpc.health.v1.Service":
		return protos.HealthCheckResponse_SERVING
	case "Service":
		return protos.HealthCheckResponse_SERVING
	default:
		service.hcl.Warn("Unknown service requested", "service", serviceName)
		return protos.HealthCheckResponse_SERVICE_UNKNOWN
	}
}
