package services

import (
	"github.com/bancodobrasil/featws-resolver-bridge/models"
	"github.com/bancodobrasil/featws-resolver-bridge/repository"
)

// FetchResolver ...
func FetchResolver(name string) (result models.Resolver, err error) {

	result, err = repository.GetResolversRepository().Get(name)
	if err != nil {
		return
	}

	return
}

func Load() (err error) {

	err = repository.GetResolversRepository().Load()
	if err != nil {
		return
	}

	return
}
