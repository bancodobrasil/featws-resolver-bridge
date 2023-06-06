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

// Resolvers is a struct that contains a map of string keys to models.Resolver values.
//
// Property
//   - resolvers: is a property of the `Resolvers` struct. It is a map that stores `models.Resolver` objects. The keys of the map are strings that represent the names of the resolvers.
type Resolvers struct {
	resolvers map[string]models.Resolver
}

var instance = Resolvers{}

// GetResolversRepository returns the Resolvers instance and loads it if it hasn't been loaded yet.
func GetResolversRepository() Resolvers {
	if instance.resolvers == nil || len(instance.resolvers) == 0 {
		instance.Load()
	}

	return instance
}

// Load is a method of the `Resolvers` struct that loads the resolvers from a JSON file. It first calls
// the `FetchFromFile` method to read the JSON file and returns an error if it fails. Then, it
// unmarshals the byte array into the `instance.resolvers` map using the `json.Unmarshal` function. If
// the unmarshaling fails, it returns an error. Finally, it returns `nil` if there are no errors.
func (r Resolvers) Load() (err error) {
	byteValue, err := r.FetchFromFile()
	if err != nil {
		log.Errorf("error occurs on fetch the JSON file: %v", err)
		return
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &(instance.resolvers))
	if err != nil {
		log.Errorf("error occurs on unmarshal the JSON file: %v", err)
		return
	}

	return
}

// Get is a method of the `Resolvers` struct that takes a `name` string as input and returns a
// `models.Resolver` object and an error. It first tries to retrieve the resolver with the given name
// from the `r.resolvers` map. If the resolver is not found, it returns an error. If the resolver is
// found, it sets the `Name` property of the resolver to the given `name` and returns the resolver
// object and `nil` error.
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

// Find is a method of the `Resolvers` struct that returns a slice of `models.Resolver` objects and an error. It first checks if the
// `r.resolvers` map is nil, and if it is, it returns an error. Then, it iterates over the keys of the
// `r.resolvers` map and calls the `Get` method to retrieve the resolver with the given name. If the
// `Get` method returns an error, it returns `nil` and the error. Otherwise, it appends the retrieved
// resolver to the `resolvers` slice and continues iterating. Finally, it returns the `resolvers` slice
// and `nil` error. This method is used to retrieve all the resolvers stored in the `r.resolvers` map.
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

// FetchFromFile is a method of the `Resolvers` struct that reads the content of a JSON file and returns it as a byte array. It first gets the configuration
// using the `config.GetConfig()` function, which returns a `config.Config` object. Then, it opens the
// JSON file using the `os.Open` function and returns an error if it fails. If the file is opened
// successfully, it reads the content of the file using the `ioutil.ReadAll` function and returns the
// byte array and `nil` error. Finally, it closes the file using the `defer` statement to ensure that
// the file is closed after the function returns.
func (r Resolvers) FetchFromFile() ([]byte, error) {
	config := config.GetConfig()
	jsonFile, err := os.Open(config.ResolversFile)
	if err != nil {
		log.Errorf("error occurs on open the JSON file: %v", err)
		return nil, err
	}
	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}
