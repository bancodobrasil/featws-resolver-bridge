package services

import (
	"github.com/bancodobrasil/featws-resolver-bridge/models"
	"github.com/bancodobrasil/featws-resolver-bridge/repository"
	log "github.com/sirupsen/logrus"
)

// FetchResolver ...
func FetchResolver(name string) (result models.Resolver, err error) {

	result, err = repository.GetResolversRepository().Get(name)
	if err != nil {
		log.Errorf("error occurs on fetch the resolvers repository: %v", err)
		return
	}

	return
}

// Load ...
func Load() (err error) {

	err = repository.GetResolversRepository().Load()
	if err != nil {
		log.Errorf("error occurs on load the resolvers repository: %v", err)
		return
	}

	return
}
