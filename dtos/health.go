package dtos

// Ready ...
type Ready struct {
	Services map[string]ReadyService
	Status   string
}

// ReadyService ...
type ReadyService struct {
	Status string
}

// HealthAll ...
type HealthAll struct {
	Live  string
	Ready Ready
}

// Live ...
type Live struct {
	Status string
}

// Service ...
type Service struct {
	Service string
	Status  string
}
