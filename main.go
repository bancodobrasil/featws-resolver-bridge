package main

import (
	"os"

	"github.com/bancodobrasil/featws-resolver-bridge/config"
	_ "github.com/bancodobrasil/featws-resolver-bridge/docs"
	"github.com/bancodobrasil/featws-resolver-bridge/routes"
	ginMonitor "github.com/bancodobrasil/gin-monitor"
	"github.com/bancodobrasil/goauth"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func setupLog() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)
}

// @title			FeatWS Resolver Bridge
// @version 		1.0
// @Description     O Bridge é um padrão de design estrutural que ajuda a dividir uma classe grande ou um conjunto de classes intimamente relacionadas em duas hierarquias distintas, para saber mais [**clique aqui**](https://refactoring.guru/pt-br/design-patterns/bridge). Essas duas hierarquias podem ser desenvolvidas independentemente uma da outra, permitindo maior flexibilidade e escalabilidade no desenvolvimento de software. Basicamente, o padrão Bridge é uma solução para o problema de como separar a interface de uma classe de sua implementação, permitindo que ambas evoluam independentemente.
// @Description
// @Description    	O projeto FeatWS Resolver Bridge tem como objetivo estabelecer a ligação entre o [**motor de regras**](https://github.com/bancodobrasil/featws-ruller) do projeto e os demais resolvers, ou possíveis resolvers, de forma a permitir uma maior flexibilidade na estruturação e desenvolvimento de software. De modo mais detalhado, essa abordagem de design estrutural proporciona a separação da lógica de implementação da abstração do código, permitindo que ambas evoluam independentemente uma da outra. Isso possibilita a utilização de diferentes tipos de resolvers no projeto, de forma a atender às necessidades específicas do negócio.
// @termsOfService 	http://swagger.io/terms/

// @contact.name 	API Support
// @contact.url 	http://www.swagger.io/support
// @contact.email 	support@swagger.io

// @license.name 	Apache 2.0
// @license.url 	http://www.apache.org/licenses/LICENSE-2.0.html

// @host 			localhost:9000
// @BasePath 		/api/v1

// @securityDefinitions.apikey 	Authentication Api Key
// @in 							header
// @name X-API-Key

// @x-extension-openapi {"example": "value on a json format"}

func main() {

	setupLog()

	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações: %s\n", err)
		return
	}

	cfg := config.GetConfig()
	if cfg == nil {
		log.Fatalf("Não foi carregada configuracão!\n")
		return
	}

	monitor, err := ginMonitor.New("v1.0.0", ginMonitor.DefaultErrorMessageKey, ginMonitor.DefaultBuckets)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = log.StandardLogger().WriterLevel(log.DebugLevel)
	gin.DefaultErrorWriter = log.StandardLogger().WriterLevel(log.ErrorLevel)

	router := gin.New()

	goauth.BootstrapMiddleware()

	router.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())
	router.Use(monitor.Prometheus())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	routes.SetupRoutes(router)
	routes.APIRoutes(router)

	port := cfg.Port

	router.Run(":" + port)
}
