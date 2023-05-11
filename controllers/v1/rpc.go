package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/bancodobrasil/featws-resolver-bridge/dtos"
	payloads "github.com/bancodobrasil/featws-resolver-bridge/payloads/v1"
	responses "github.com/bancodobrasil/featws-resolver-bridge/responses/v1"
	"github.com/bancodobrasil/featws-resolver-bridge/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ResolveHandler godoc
// @Summary 		Execute the Resolve resolutions
// @Description
// @Description	O Resolver Bridge pode ser usado para buscar dados em outros resolvers que fazem parte do FeatWS, podendo retornar diversos parâmetros, como os seguintes:
// @Description
// @Description	- account (conta)
// @Description	- accountType (tipo de conta)
// @Description	- age (idade)
// @Description	- agenciaDet (tipo de conta)
// @Description	- branch (agencia)
// @Description	- branchState (Estado da agencia)
// @Description	- customerBase (agencia)
// @Description	- dataNascimento (data de nascimento)
// @Description	- employeeDependency (Dependecia do empregado do banco - só trará um retorno se a pessoa for funci do banco)
// @Description	- employeeKey (se é empregado do banco)
// @Description	- enterpriseKey (chave empresarial)
// @Description	- gender (sexo)
// @Description	- holder (titularidade da conta)
// @Description	- holderState (estado do titularidade da conta)
// @Description	- mci
// @Description	- mcipj
// @Description	- wallet (carteira)
// @Description
// @Description No geral, o Resolver Adapter Project é uma biblioteca útil que pode simplificar o desenvolvimento de APIs GraphQL fornecendo resolvers pré-construídos que podem ser facilmente integrados em outros projetos.
// @Description
// @Description	Para conseguir utilizar o endpoint é necessário colocar no body, dentro do *context*, contexto, da requisição a agencia do cliente, *branch*, como também, a conta do mesmo, *account*, como o exemplo a seguir:
// @Description
// @Description	```
// @Description	{
// @Description	"context": {
// @Description		"account": “7894”,
// @Description		"branch": “4024”,
// @Description	},
// @Description	"load":[]
// @Description	}
// @Description	```
// @Description
// @Description	Com esse input o body de resposta trará todos os parâmetros da conta. O *load* é opcional na requisição. Nele é possível passar parâmetros que você deseja buscar no banco de dados ao invés de buscar todos os parâmetros, podendo buscar quantos sejam necessários, como o exemplo a seguir:
// @Description
// @Description	```
// @Description	{
// @Description	"context": {
// @Description		"account": "7894",
// @Description		"branch": "4024",
// @Description	},
// @Description	"load":["age","gender","holder"]
// @Description	}
// @Description	```
// @Description	Com esse input o body de resposta será trago dentro do parâmetro *context*, contexto.
// @Description
// @Description	```
// @Description	{
// @Description	"context": {
// @Description	"age": "36",
// @Description	"gender": "M",
// @Description	"holder": "1"
// @Description	},
// @Description	"errors": {}
// @Description	}
// @Description	```
// @Tags 			resolve
// @Accept  		json
// @Produce  		json
// @Param           resolver path string false "resolver"
// @Param           parameters body payloads.Resolve true "Parameters"
// @Success 		200 {object} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Security 		Authentication Api Key
// @Router 			/resolve/{resolver} [post]
func ResolveHandler(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	var input payloads.Resolve
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Errorf("error occurs on binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, responses.Error{
			Error: err.Error(),
		})
		return
	}

	resolveContext := dtos.NewResolveV1(input)

	resolver := input.Resolver

	resolverName, exists := c.Params.Get("resolver")

	if exists {
		resolver = resolverName
	}

	err := services.Resolve(ctx, resolver, &resolveContext)
	if err != nil {
		log.Errorf("error occurs on resolve the context: %v", err)
		c.JSON(http.StatusInternalServerError, responses.Error{
			Error: err.Error(),
		})
		return
	}

	resolverOutput := responses.NewResolve(resolveContext)

	c.JSON(http.StatusOK, resolverOutput)
}

// LoadHandler 		godoc
// @Summary 		Load Resolvers
// @Description 	Load Resolvers description
// @Tags 			load
// @Accept  		json
// @Produce  		json
// @Success 		200 {object} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Security 		Authentication Api Key
// @Router 			/load [get]
func LoadHandler(c *gin.Context) {

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := services.Load()
	if err != nil {
		log.Errorf("error occurs on load: %v", err)
		c.JSON(http.StatusInternalServerError, responses.Error{
			Error: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
