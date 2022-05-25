package services

import (
	"net/http"
	"os"
	"time"

	"github.com/bancodobrasil/featws-resolver-bridge/config"
	"github.com/bancodobrasil/featws-resolver-bridge/dtos"
	"github.com/sirupsen/logrus"
)

// Client ...
type Client struct {
	Transport     http.RoundTripper
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
	Timeout       time.Duration
}

// CheckHealthReady ...
func CheckHealthReady() dtos.Ready {
	response := dtos.Ready{
		Services: make(map[string]dtos.ReadyService),
		Status:   "OK",
	}

	resolvers, err := FetchResolvers()
	if err != nil {
		panic(err)
	}

	for _, resolver := range resolvers {
		cfg := config.GetConfig()
		url := resolver.Options["url"].(string)
		client := &http.Client{
			Timeout: time.Duration(cfg.ReadyTimeout) * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			if os.IsTimeout(err) {
				logrus.Warnf("timeout on %s", url)
			}
			response.Services[resolver.Name] = dtos.ReadyService{
				Status: err.Error(),
			}
			response.Status = "Error"
			continue
		}
		response.Services[resolver.Name] = dtos.ReadyService{
			Status: resp.Status,
		}
	}
	return response
}

// CheckHealthLive ...
func CheckHealthLive() dtos.Live {
	return dtos.Live{
		Status: "Application is live!!!",
	}
}

// CheckHealthAll ...
func CheckHealthAll() dtos.HealthAll {

	response := dtos.HealthAll{
		Live:  CheckHealthLive().Status,
		Ready: CheckHealthReady(),
	}
	for k, v := range response.Ready.Services {
		if v.Status != "200 OK" {
			v.Status = "Service is not ready!!!!"
			response.Ready.Services[k] = v
			continue
		}
		v.Status = "Service is ready!!!!"
		response.Ready.Services[k] = v
	}

	return response
}
