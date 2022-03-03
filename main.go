package main

import (
	"os"

	"github.com/bancodobrasil/featws-resolver-bridge/config"
	"github.com/bancodobrasil/featws-resolver-bridge/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func setupLog() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)
}

// Config contains all settings of module
var Config = config.Config{}

// Run start the resolver server with resolverFunc
func main() {

	setupLog()

	err := config.LoadConfig(&Config)
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações: %s\n", err)
	}

	router := gin.New()

	routes.SetupRoutes(router)

	port := Config.Port

	router.Run(":" + port)
}
