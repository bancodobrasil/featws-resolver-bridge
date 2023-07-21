package services

import (
	"github.com/bancodobrasil/featws-resolver-bridge/models"
	"github.com/bancodobrasil/featws-resolver-bridge/repository"
	log "github.com/sirupsen/logrus"
)

// FetchResolver fetches a resolver from a repository by name.
func FetchResolver(name string) (result models.Resolver, err error) {

	result, err = repository.GetResolversRepository().Get(name)
	if err != nil {
		log.Errorf("error occurs on fetch the resolvers repository: %v", err)
		return
	}

	return
}

// FetchResolvers retrieves a list of resolvers from a repository.
func FetchResolvers() (result []models.Resolver, err error) {

	result, err = repository.GetResolversRepository().Find()
	if err != nil {
		return
	}

	return
}

// Load function loads the resolvers repository and logs any errors that occur.
func Load() (err error) {

	err = repository.GetResolversRepository().Load()
	if err != nil {
		log.Errorf("error occurs on load the resolvers repository: %v", err)
		return
	}

	return
}

// FetchFromFile fetches data from a resolvers repository file and logs any errors that occur.
func FetchFromFile() (result []byte, err error) {
	result, err = repository.GetResolversRepository().FetchFromFile()
	if err != nil {
		log.Errorf("error occurs on fetch the resolvers repository: %v", err)
		return
	}

	return
}
