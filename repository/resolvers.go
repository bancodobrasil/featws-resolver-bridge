package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bancodobrasil/featws-resolver-bridge/config"
	"github.com/bancodobrasil/featws-resolver-bridge/models"
	log "github.com/sirupsen/logrus"
)

// Resolvers ...
type Resolvers struct {
	resolvers map[string]models.Resolver
}

var instance = Resolvers{}

// GetResolversRepository ...
func GetResolversRepository() Resolvers {
	if instance.resolvers == nil || len(instance.resolvers) == 0 {
		instance.Load()
	}

	return instance
}

// Load ...
func (r Resolvers) Load() (err error) {
	config := config.GetConfig()

	jsonFile, err := os.Open(config.ResolversFile)
	if err != nil {
		log.Errorf("error occurs on open the JSON file: %v", err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &(instance.resolvers))
	if err != nil {
		log.Errorf("error occurs on unmarshal the JSON file: %v", err)
		return
	}

	return
}

// Get ...
func (r Resolvers) Get(name string) (resolver models.Resolver, err error) {
	resolver, ok := r.resolvers[name]
	if !ok {
		err = fmt.Errorf("couldn't bound resolver with '%s'", name)
		log.Error(err)
		return
	}
	resolver.Name = name
	return
}

// Find ...
func (r Resolvers) Find() (resolvers []models.Resolver, err error) {

	if r.resolvers == nil {
		err = fmt.Errorf("couldn't fetch resolvers")
		return
	}

	for name := range r.resolvers {
		resolver, err := r.Get(name)
		if err != nil {
			return nil, err
		}

		resolvers = append(resolvers, resolver)
	}

	return
}
