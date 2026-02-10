# Promise (Go)

## Sobre

Este projeto implementa uma maneira de encadear a execução de funções com um estado compartilhado, e parando a execução no primeiro erro.

## Instalação

```
go get github.com/filipewelton/promise
```

## Uso

```go
p := NewPromise[int]()

err := p.
  Then(func(ctx *int, reject Reject) {
      *ctx = 42
  })
  Then(func(ctx *int, reject Reject) {
      fmt.Println(*ctx)
  }).
  Catch()

fmt.Println(err)
```

## Testes

Os testes são escritos com Ginkgo e Gomega:

```
ginkgo -r -v
```

## Versão

A versão atual é v1.0.0.

## Autor

Este projeto foi escrito com auxílio de Inteligência Artificial (GitHub Copilot/GPT-4.1).
