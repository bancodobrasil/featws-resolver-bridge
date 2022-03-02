package services

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type ResolverService struct {
	resolverFunc ResolverFunc
}

var instance *ResolverService

type ResolveInput struct {
	Context map[string]interface{} "json:'context'"
	Load    []string               "json:'load'"
}

type ResolveOutput struct {
	Context map[string]interface{}
	Errors  map[string]interface{}
}

type ResolverFunc func(ResolveInput, *ResolveOutput)

func SetupResolver(resolverFunc ResolverFunc) {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &ResolverService{
				resolverFunc: resolverFunc,
			}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
}

func Resolve(input ResolveInput) (output *ResolveOutput) {
	output = &ResolveOutput{
		Context: input.Context,
		Errors:  make(map[string]interface{}),
	}
	instance.resolverFunc(input, output)

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
