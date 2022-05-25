package dtos

type Ready struct {
	Services map[string]ReadyService
	Status   string
}

type ReadyService struct {
	Status string
}

type HealthAll struct {
	Live  string
	Ready Ready
}

type Live struct {
	Status string
}

type Service struct {
	Service string
	Status  string
}
