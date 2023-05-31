package v1

import "github.com/bancodobrasil/featws-resolver-bridge/dtos"

// HealthAll type contains a string field for live status and a Ready struct field for readiness status.
//
// Property:
//   - {string} Live: is a string that is likely used to indicate the status of a service or application. It could have values such as "running", "stopped", "error", etc.
//   - {Ready} Ready: is a nested struct that contains information about the readiness of the application. It could have properties such as `Status` (indicating whether the application is ready or not), `Message` (providing additional information about the readiness status), and `Timestamp` (indicating the timestamp when the readiness check was performed).
type HealthAll struct {
	Live  string `json:"live"`
	Ready Ready  `json:"ready"`
}

// Live type has a single field called "Status" with a JSON tag.
//
// Property:
//   - {string} Status: is a property of a struct type called Live. It is a string type and has a JSON tag of "status". The JSON tag is used to specify the name of the property when the struct is encoded to or decoded from JSON format.
type Live struct {
	Status string `json:"status"`
}

// Ready contains a map of ReadyService objects and a status string, both of which are optional in JSON.
//
// Property:
//   - Services: Services is a map of ReadyService objects, where each key represents the name of a service and the value represents the status of that service.
//   - Status: is a string that represents the overall readiness status of a system or application. It can be used to indicate whether the system is ready to receive requests or if there are any issues that need to be resolved before it can be considered fully operational.
type Ready struct {
	Services map[string]ReadyService `json:"services,omitempty"`
	Status   string                  `json:"status,omitempty"`
}

// ReadyService has a single field called Status with an optional JSON tag.
//
// Property:
//   - {string} Status: is a string field that represents the current status of a service. It is tagged with `json:"status,omitempty"` which means that when this struct is marshalled into JSON, the `Status` field will be included only if it has a non-zero value.
type ReadyService struct {
	Status string `json:"status,omitempty"`
}

// Service has two fields, "Service" and "Status", both of which are strings and have JSON tags.
//
// Property:
//   - {string} Service - Service is a property of a struct in Go programming language. It is defined as a string type and is tagged with `json:"service"`. This tag is used to specify the name of the field when the struct is serialized to JSON.
//   - {string} Status - Status is a property of the Service struct. It is of type string and is tagged with `json:"status"`. This tag indicates that when the struct is marshaled into JSON format, the property should be represented as "status".
type Service struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

// NewReady creates a new Ready object from a DTO by initializing its services and status.
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

// NewReadyService creates a new ReadyService object based on the provided DTO.
func NewReadyService(dto dtos.ReadyService) ReadyService {
	return ReadyService{
		Status: dto.Status,
	}
}

// NewHealthAll creates a new HealthAll object using data from a HealthAll DTO.
func NewHealthAll(dto dtos.HealthAll) HealthAll {
	return HealthAll{
		Live:  dto.Live,
		Ready: NewReady(dto.Ready),
	}
}

// NewLive creates a new Live object based on a Live DTO.
func NewLive(dto dtos.Live) Live {
	return Live{
		Status: dto.Status,
	}
}

// NewService creates a new Service object based on the provided DTO object.
func NewService(dto dtos.Service) Service {
	return Service{
		Service: dto.Service,
		Status:  dto.Status,
	}
}
