# Promise (Go)

## Sobre

Este projeto implementa uma maneira de encadear a execução de funções em Go, com ou sem estado compartilhado (contexto), parando a execução no primeiro erro.

## Instalação

```sh
go get github.com/filipewelton/promise
```

## Uso

### Promise sem contexto

```go
import "github.com/filipewelton/promise"

p := promise.NewPromise()

err := p.
  Then(func(reject promise.Reject) {
      // sua lógica
  }).
  Then(func(reject promise.Reject) {
      // outra lógica
  }).
  Catch()

fmt.Println(err) // nil se não houve erro
```

### Promise com contexto

```go
import "github.com/filipewelton/promise"

p := promise.NewPromiseWithContext[int](nil)

err := p.
  Then(func(ctx *int, reject promise.Reject) {
      *ctx = 42
  }).
  Then(func(ctx *int, reject promise.Reject) {
      *ctx += 8
  }).
  Catch()

fmt.Println(err)         // nil se não houve erro
fmt.Println(p.Context)   // 50
```

## Testes

Os testes são escritos com Ginkgo e Gomega:

```sh
ginkgo -r -v
```

## Versão

A versão atual é v2.0.0.

## Autor

Este projeto foi escrito com auxílio de Inteligência Artificial (GitHub Copilot/GPT-4.1).
