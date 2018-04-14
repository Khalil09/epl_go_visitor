# Implementação do EPL utilizando linguagem GO e padrão *Visitor*

### Resumo

Esse breve relatório é direcionado a diciplina técnicas de programação 2 lecionada pelo professor Bonifácio na Universidade de Brasília.
Aqui apresentaremos uma implementação básica do EPL utilizando o padrão visitor em cima da linguagem GO.

Além disso apresentaremos limitações da linguagem em relação o uso de outros padrões de projeto.

### Estrutura dos arquivos

Os arquivos foram dividos baseado nos nomes de cada entidade que no caso de *Go* representaria uma *struct*. Todos pertencem a um mesmo *package* sendo compilados como se fossem um só arquivo.

O arquivo principal do pacote eve possuir o nome do pacote e o arquivo de teste deve possuir *NomeDoPacote_teste.go* assim somente execute o comando abaixo para executar o compilado de testes.

```sh
go test
```

### Implementando o EPL com GO e Visitor

Utilizando o projeto em *scala* usando padrão *visitor* disponibilizado pelo professor, tentamos reproduzir o mesmo com a linguagem *Go*.

Usando interfaces e a capacidade de herença de métodos por parte das struct conseguimos criar a inteface *Expression* da qual *Add, Multi, Sub e Literal* herdam os métodos *print()* e *accept()*.

Após isso criamos a interface *visitor*. No caso de *Go* não existe sobrecarga para métodos. Se tentar o código abaixo receberá a seguinte mensagem:

```go
type Sobrecarga interface {
    sobrecarga(value int)
    sobrecarga(value string)
}
```

resultado no terminal:

```sh
./main.go:5: duplicate method sobrecarga
```

Baseado nisso fomos obrigados a criar um método visitor para cada tipo de *struct* que requisitava um visitor:

```go
type Visitor interface {
    visitLiteral(lit *Literal)
    visitAdd(add *Add)
    visitSub(sub *Sub)
    visitMulti(multi *Multi)
}
```

Com isso fomos capazes e fazer um impleentação fo visitor utilizando *Go*.

### Outros padrões tentados

Foi tentado também utilizar métodos *mixin* como utilizado no exemplo de *scala* feito pelo professor. O próprio conceito de *mixin* não se encaixa nas possibilidades que *Go* fornece, ou seja, não podemos utilizar métodos de outras classes sem que a classe tenha parentesco com essa outras classes.

### Possibilidades para Go

Encontramos um repositório do *Github* onde existe várias categorias de padrão de projeto e quais a linguagem *GO* tem possibilidade de emplementar ou não.

repositório: https://github.com/tmrts/go-patterns

### Execução do projeto atual

Para executar esse projeto incialmente é necessário que tenha o *Go lang* instalado no seu computador. Acompanhe o *link* abaixo para melhores explicações:

Getting started GO: https://golang.org/doc/install

A estrutura do seu projeto deve ser parecida com a que está abaixo:

```
pasta <- GOPATH aponta para cá. Utilize o camando <go install>
│
└───src
│   │
│   └───epl_go_visitor
│       │   eval.go
│       │   main.go
│       │   ...
│   
└───bin <- pasta criada automaticamente
│    │   epl_go_visitor <- executável gerado
│    
└───pkg <- Alguma finalidade que ainda não sei
```