package services

import (
	"net/http"
	"os"
	"time"

	"github.com/bancodobrasil/featws-resolver-bridge/config"
	"github.com/bancodobrasil/featws-resolver-bridge/dtos"
	"github.com/sirupsen/logrus"
)

// Client type in Go represents an HTTP client with customizable transport, redirect handling,
// cookie jar, and timeout.
//
// Property:
//   - Transport: is a property of the Client struct that represents the mechanism used to make HTTP requests. It is an interface that defines the behavior of an HTTP client. The default transport used by the Client is the http.DefaultTransport, which is an implementation of the Transport interface provided by the Go standard library.
//   - CheckRedirect: is a function that specifies the policy for handling redirects. It takes in the original request and a slice of requests that were made in the redirect chain. It returns an error if the redirect should not be followed, or nil if it should be followed.
//   - Jar: The `Jar` property is of type `http.CookieJar` and is used to store and manage cookies for HTTP requests. Cookies are small pieces of data that are sent from a website to a user's web browser and are used to remember user preferences, login information, and other data.
//   - Timeout: Timeout is a property of the Client struct that specifies the maximum amount of time a request can take before it is cancelled. It is of type time.Duration, which allows for specifying the timeout in units such as seconds, milliseconds, or nanoseconds.
type Client struct {
	Transport     http.RoundTripper
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
	Timeout       time.Duration
}

// CheckHealthReady checks the health readiness of services by fetching resolvers and making HTTP requests
// to their URLs.
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

// CheckHealthLive returns a DTO object indicating that the application is live.
func CheckHealthLive() dtos.Live {
	return dtos.Live{
		Status: "Application is live!!!",
	}
}

// CheckHealthAll checks the health status of all services and returns a response with their live and
// ready status.
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
