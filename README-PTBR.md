
[![Go Reference](https://pkg.go.dev/badge/github.com/abu-lang/goabu.svg)](https://pkg.go.dev/github.com/bancodobrasil/featws-resolver-bridge)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/bancodobrasil/featws-resolver-bridge/blob/develop/LICENSE)

# Featws Reslver Bridge [![About_en](https://github.com/yammadev/flag-icons/blob/master/png/US.png?raw=true)](https://github.com/bancodobrasil/featws-resolver-bridge/blob/develop/README.md)


## Como executar

Para executar este projeto, você precisa ter certos pré-requisitos configurados em sua máquina. Esses pré-requisitos incluem:

- [Golang](https://go.dev/doc/install)
 - [Swaggo](https://github.com/swaggo/swag/blob/master/README_pt.md#come%C3%A7ando)

Para executar o projeto, siga estes passos:

- Abra o terminal no diretório do projeto e execute o comando `go mod tidy` para garantir que todas as dependências necessárias estejam instaladas.

- Em seguida, execute o comando `swag init` para inicializar o Swagger e gerar a documentação da API necessária.

- Finalmente, execute o comando `make run` para iniciar o projeto.

O projeto será executado em `localhost:9000`. Para acessar a documentação do Swagger, [clique aqui](http://localhost:9000/swagger/index.html#/).

Seguindo estes passos, o projeto estará em execução e você poderá acessar a documentação da API através do Swagger.

## GoDoc

Para acessar a documentação do GoDoc, primeiro instale o GoDoc na sua máquina. Abra um terminal e digite:

````
go get golang.org/x/tools/cmd/godoc
````

Em seguida rode no terminal do repositório o comando a seguir:

````
godoc -http=:6060
````

O GoDoc será executado em `localhost:6060`. Para acessar a documentação do GoDoc, basta [clicar aqui](http://localhost:6060/pkg/).