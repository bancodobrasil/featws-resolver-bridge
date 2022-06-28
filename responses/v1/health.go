package v1

import "github.com/bancodobrasil/featws-resolver-bridge/dtos"

// HealthAll ...
type HealthAll struct {
	Live  string `json:"live"`
	Ready Ready  `json:"ready"`
}

// Live ...
type Live struct {
	Status string `json:"status"`
}

// Ready ...
type Ready struct {
	Services map[string]ReadyService `json:"services,omitempty"`
	Status   string                  `json:"status,omitempty"`
}

// ReadyService ...
type ReadyService struct {
	Status string `json:"status,omitempty"`
}

// Service ...
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

// NewReadyService ...
func NewReadyService(dto dtos.ReadyService) ReadyService {
	return ReadyService{
		Status: dto.Status,
	}
}

// NewHealthAll ...
func NewHealthAll(dto dtos.HealthAll) HealthAll {
	return HealthAll{
		Live:  dto.Live,
		Ready: NewReady(dto.Ready),
	}
}

// NewLive ...
func NewLive(dto dtos.Live) Live {
	return Live{
		Status: dto.Status,
	}
}

// NewService ...
func NewService(dto dtos.Service) Service {
	return Service{
		Service: dto.Service,
		Status:  dto.Status,
	}
}
