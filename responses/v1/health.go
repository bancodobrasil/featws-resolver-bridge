package v1

import "github.com/bancodobrasil/featws-resolver-bridge/dtos"

type HealthAll struct {
	Live  string `json:"live"`
	Ready Ready  `json:"ready"`
}

type Live struct {
	Status string `json:"status"`
}

type Ready struct {
	Services map[string]ReadyService `json:"services,omitempty"`
	Status   string                  `json:"status,omitempty"`
}

type ReadyService struct {
	Status string `json:"status,omitempty"`
}

type Service struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

// NewReady ...
func NewReady(dto dtos.Ready) Ready {
	services := make(map[string]ReadyService)
	for k, v := range dto.Services {
		services[k] = NewReadyService(v)
	}
	return Ready{
		Services: services,
		Status:   dto.Status,
	}
}

func NewReadyService(dto dtos.ReadyService) ReadyService {
	return ReadyService{
		Status: dto.Status,
	}
}

func NewHealthAll(dto dtos.HealthAll) HealthAll {
	return HealthAll{
		Live:  dto.Live,
		Ready: NewReady(dto.Ready),
	}
}

func NewLive(dto dtos.Live) Live {
	return Live{
		Status: dto.Status,
	}
}

func NewService(dto dtos.Service) Service {
	return Service{
		Service: dto.Service,
		Status:  dto.Status,
	}
}
