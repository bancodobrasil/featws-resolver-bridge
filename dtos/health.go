package dtos

// Ready contains a map of ReadyService objects and a status string.
//
// Property
//   - Services: is a map that contains information about the readiness status of various services. The keys of the map are the names of the services, and the values are of type ReadyService, which contains information about the readiness status of a single service.
//   - {string} Status - Status is a string property that represents the overall status of the system. It could be used to indicate whether the system is ready or not.
type Ready struct {
	Services map[string]ReadyService
	Status   string
}

// ReadyService has a single string field called Status.
//
// Property
//   - {string} Statu: is a string type property that belongs to the "ReadyService" struct. It is used to represent the current status of the service. The value of this property can be set to different strings such as "ready", "not ready", "in progress", etc, depending on the specific state or condition of the service. The "Status" property provides information about the readiness of the service and can be used to determine its operational state.
type ReadyService struct {
	Status string
}

// HealthAll type contains a string field for live status and a Ready struct field for readiness
// status.
//
// Property:
//   - {string} Live: is a string that is likely used to indicate the health status of a service or application. It could have a value of "true" or "false" to indicate whether the service is currently running or not.
//   - {Ready} Ready: is a struct type that contains properties related to the readiness of a service or application. It may include information such as the status of database connections, availability of required resources, and other factors that determine whether the service is ready to handle requests.
type HealthAll struct {
	Live  string
	Ready Ready
}

// Live  type has a single field called "Status" of type string.
//
// Property:
//   - {string} Status: is a string type property that represents the current status of a "Live" object. It could be used to indicate whether the object is currently active, inactive, or any other relevant status.
type Live struct {
	Status string
}

// Service has two fields, "Service" and "Status", both of which are strings.
//
// Property:
//   - {string} Service: is a string that represents the name or identifier of a particular service.
//   - {string} Status: is a string that represents the current status of a service. It could be used to indicate whether a service is currently running, stopped, or experiencing an error, for example.
type Service struct {
	Service string
	Status  string
}
