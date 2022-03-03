package services

import (
	"fmt"
	"sync"
)

// ResolverFunc define the Resolver Function structure
type ResolverFunc func(ResolveInput, *ResolveOutput)

var lock = &sync.Mutex{}

var resolverFunc ResolverFunc

// ResolveInput contains all input for resolver execution
type ResolveInput struct {
	Context map[string]interface{} `json:"context"`
	Load    []string               `json:"load"`
}

// ResolveOutput contais all output of resolver execution
type ResolveOutput struct {
	Context map[string]interface{} `json:"context"`
	Errors  map[string]interface{} `json:"errors"`
}

// SetupResolver to config the current resolver func
func SetupResolver(rFunc ResolverFunc) {
	lock.Lock()
	defer lock.Unlock()
	if resolverFunc == nil {
		if resolverFunc == nil {
			fmt.Println("Creating single instance now.")
			resolverFunc = rFunc
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
}

// Resolve to execute the resolver
func Resolve(input ResolveInput) (output *ResolveOutput) {
	output = &ResolveOutput{
		Context: input.Context,
		Errors:  make(map[string]interface{}),
	}
	resolverFunc(input, output)

	if len(input.Load) > 0 {
		oldContext := output.Context

		output.Context = make(map[string]interface{})

		for _, key := range input.Load {
			value, ok := oldContext[key]
			if ok {
				output.Context[key] = value
			}
		}
	}

	return
}
