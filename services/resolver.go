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
	telemetry "github.com/bancodobrasil/gin-telemetry"
	log "github.com/sirupsen/logrus"
)

// Resolve ...
func Resolve(ctx context.Context, resolverName string, dto *dtos.ResolveContext) (err error) {
	log.Debugf("Resolving with '%s': %s", resolverName, dto)

	resolver, err := FetchResolver(resolverName)
	if err != nil {
		log.Errorf("error occurs on fetch the resolver: %v", err)
		return
	}

	if resolver.Type == "http" {
		err = resolveHTTP(ctx, resolver, dto)
		if err != nil {
			log.Errorf("error occurs on resolve HTTP: %v", err)
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

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(input)
	if err != nil {
		log.Errorf("error occurs on encoder the JSON: %v", err)
		return
	}

	log.Debugf("Resolving with '%s' Encoded: %v", url, buf.String())

	req, err := http.NewRequestWithContext(ctx, "POST", url, &buf)
	if err != nil {
		log.Errorf("error occurs on create a request: %v", err)
		return
	}

	if resolver.Options["headers"] != nil {
		headers := resolver.Options["headers"].(map[string][]string)
		if len(headers) > 0 {
			req.Header = headers
		}
	}
	client := &http.Client{}

	if !telemetry.MiddlewareDisabled {
		telemetry.Inject(ctx, req.Header)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.WithContext(ctx).Errorf("error occurs on initializate a HTTP client: %v", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Errorf("error occurs on read the reponse body: %v", err)
		return
	}

	log.Debugf("Resolving with '%s': %v > %s", url, input, string(data))

	output := resolveOutputV1{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		log.WithContext(ctx).Errorf("error occurs on unmarshal the data into output: %v", err)
		return
	}

	dto.Context = output.Context
	dto.Errors = output.Errors

	return
}
