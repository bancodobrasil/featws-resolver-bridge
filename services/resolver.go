package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bancodobrasil/featws-resolver-bridge/dtos"
	"github.com/bancodobrasil/featws-resolver-bridge/models"
	log "github.com/sirupsen/logrus"
)

// Resolve ...
func Resolve(ctx context.Context, resolverName string, dto *dtos.ResolveContext) (err error) {
	log.Debugf("Resolving with '%s': %s", resolverName, dto)

	resolvers, err := FetchResolvers(ctx, models.Resolver{
		Name: resolverName,
	})
	if err != nil {
		return
	}

	if len(resolvers) != 1 {
		err = fmt.Errorf("couldn't bound resolver with '%s'", resolverName)
		return
	}

	resolver := resolvers[0]

	if resolver.Type == "http" {
		err = resolveHTTP(ctx, resolver, dto)
		if err != nil {
			return
		}
	}

	return
}

type resolveInputV1 struct {
	Context map[string]interface{} `json:"context"`
	Load    []string               `json:"load"`
}

type resolveOutputV1 struct {
	Context map[string]interface{} `json:"context"`
	Errors  map[string]interface{} `json:"errors"`
}

func resolveHTTP(ctx context.Context, resolver models.Resolver, dto *dtos.ResolveContext) (err error) {

	url := fmt.Sprintf("%s/api/v1/resolve", resolver.Options["url"])

	url = strings.ReplaceAll(url, "//api/v1", "/api/v1")

	input := resolveInputV1{
		Context: dto.Context,
		Load:    dto.Load,
	}

	log.Debugf("Resolving with '%s': %v", url, input)

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(input)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return
	}

	if resolver.Options["headers"] != nil {
		headers := resolver.Options["headers"].(map[string][]string)
		if len(headers) > 0 {
			req.Header = headers
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Debugf("Resolving with '%s': %v > %s", url, input, string(data))

	output := resolveOutputV1{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		return
	}

	dto.Context = output.Context
	dto.Errors = output.Errors

	return
}
