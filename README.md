# Pipeline (Go)

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Sobre

Este projeto implementa uma maneira simples e eficiente de encadear a execução de funções em Go, com ou sem estado compartilhado (contexto). A execução para no primeiro erro encontrado, seguindo o padrão fail-fast.

## Características

- 🔗 Encadeamento de funções com sintaxe fluente
- 🎯 Suporte a pipelines com e sem contexto
- 🛑 Execução para no primeiro erro (fail-fast)
- 🔒 Type-safe com uso de Generics do Go
- ✅ Testado com Ginkgo e Gomega

## Casos de uso

Este pacote é ideal para:

- **Validações em cadeia**: Valide dados através de múltiplas etapas
- **Processamento de dados**: Transforme dados através de várias funções
- **Workflows**: Implemente workflows complexos com estado compartilhado
- **Operações sequenciais**: Execute operações que dependem uma da outra
- **Tratamento de erros simplificado**: Evite verificações de erro repetitivas

## Instalação

```sh
go get github.com/filipewelton/pipeline/v3
```

## Uso

### Pipeline sem contexto

Ideal para executar uma sequência de operações independentes:

```go
package main

import (
    "fmt"
    "github.com/filipewelton/pipeline/v3"
)

func main() {
    p := pipeline.New()

    err := p.
        Add(func() error {
            fmt.Println("Primeira operação")
            return nil
        }).
        Add(func() error {
            fmt.Println("Segunda operação")
            return nil
        }).
        Run()

    if err != nil {
        fmt.Printf("Erro: %v\n", err)
    }
}
```

### Pipeline com contexto

Ideal quando você precisa compartilhar e modificar estado entre as operações:

```go
package main

import (
    "fmt"
    "github.com/filipewelton/pipeline/v3"
)

func main() {
    initialValue := 0
    p := pipeline.NewWithContext(&initialValue)

    result, err := p.
        Add(func(ctx *int) (int, error) {
            return *ctx + 10, nil
        }).
        Add(func(ctx *int) (int, error) {
            return *ctx * 2, nil
        }).
        Run()

    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }

    fmt.Printf("Resultado: %d\n", result) // Resultado: 20
}
```

### Tratamento de erros

O pipeline para imediatamente quando um executor retorna um erro:

```go
p := pipeline.New()

err := p.
    Add(func() error {
        return nil // Executa
    }).
    Add(func() error {
        return pipeline.ErrRejectedWithoutReason // Para aqui
    }).
    Add(func() error {
        return nil // Não executa
    }).
    Run()

fmt.Println(err) // pipeline rejected without reason
```

### Exemplo com tipos customizados

O pipeline com contexto suporta qualquer tipo através de Generics:

```go
type User struct {
    Name  string
    Email string
    Age   int
}

func main() {
    user := User{Name: "João"}
    p := pipeline.NewWithContext(&user)

    result, err := p.
        Add(func(ctx *User) (User, error) {
            ctx.Email = "joao@example.com"
            return *ctx, nil
        }).
        Add(func(ctx *User) (User, error) {
            ctx.Age = 25
            return *ctx, nil
        }).
        Run()

    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }

    fmt.Printf("Usuário: %+v\n", result)
    // Usuário: {Name:João Email:joao@example.com Age:25}
}
```

## API

### Tipos

```go
type Pipeline struct { ... }
type PipelineWithContext[T any] struct { ... }

type Executor func() error
type ExecutorWithContext[T any] func(ctx *T) (T, error)

var ErrRejectedWithoutReason = errors.New("pipeline rejected without reason")
```

### Métodos

#### Pipeline sem contexto

- `New() *Pipeline` - Cria um novo pipeline sem contexto
- `Add(executor Executor) *Pipeline` - Adiciona um executor ao pipeline
- `Run() error` - Executa todos os executores em sequência

#### Pipeline com contexto

- `NewWithContext[T any](context *T) *PipelineWithContext[T]` - Cria um pipeline com contexto tipado
- `Add(executor ExecutorWithContext[T]) *PipelineWithContext[T]` - Adiciona um executor ao pipeline
- `Run() (T, error)` - Executa todos os executores e retorna o contexto final

## Testes

Os testes são escritos com Ginkgo e Gomega:

```sh
# Executar testes
go test -v

# Com Ginkgo
ginkgo -r -v
```

## Por que usar Pipeline?

### Sem Pipeline

```go
func ProcessData(data int) (int, error) {
    result, err := Step1(data)
    if err != nil {
        return 0, err
    }

    result, err = Step2(result)
    if err != nil {
        return 0, err
    }

    result, err = Step3(result)
    if err != nil {
        return 0, err
    }

    return result, nil
}
```

### Com Pipeline

```go
func ProcessData(data int) (int, error) {
    p := pipeline.NewWithContext(&data)

    return p.
        Add(func(ctx *int) (int, error) { return Step1(*ctx) }).
        Add(func(ctx *int) (int, error) { return Step2(*ctx) }).
        Add(func(ctx *int) (int, error) { return Step3(*ctx) }).
        Run()
}
```

Mais limpo, menos repetição, e mais fácil de manter!

## Requisitos

- Go 1.25+ (para suporte a Generics)

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para:

1. Fazer fork do projeto
2. Criar uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abrir um Pull Request

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## Autor

Desenvolvido por Filipe Welton com auxílio de Inteligência Artificial (GitHub Copilot).
