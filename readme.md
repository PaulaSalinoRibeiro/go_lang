# Go Lang

## Hello world

Um arquivo executável precisa seguir a assinatura do arquivo `hello.go` contendo o `package main`, `imports` e `func main`.
Para rodar o programa executa o comando `go run <nome do arquivo>`.

## Pacotes

Pacotes são um conjunto de arquivos que estão no mesmo diretório.

Quando trabalha com pacotes, precisar iniciar o modulo, para isso, basta executar o comando `go mod init <nome do modulo>`.
Para realizar o import primeiro o nome do modulo e depois o nome do pacote `ex: "module/auxiliars"`.

Note: funções com letras minusculas são visiveis apenas no prórpio pacote, para ser exportada é necessário utilizar letras maiusculas.
E é uma boa prática escrever um comentário em cima de funções que são exportadas informando o que fazem!

Para utilizar pacotes externos utiliza o comando `go get <url do pacote>`. Automaticamente será criado um arquivo `go.sum`
Para remover uma dependecia externa que não está sendo utilizada executa o comando `go mod tidy`.

## Variaveis

Em Go não é permitido criar uma variavel e não utilizá-la. Além disso, é uma linguagem fortemente tipada.

## Tipos de Dados

- int8, int16, int32, int64 suportam 8, 16, 32 e 64 bits respectivamente
- int - toma como padrão a arquitetura do computador
- uint - não aceita numeros negativos
- float32, float64
- string - sempre utiliza aspas duplas
- bool - boleano
- error - erro
#### Valor Zero - É o valor de uma variavel que não foi inicializada

- string = ""
- int = 0
- float = 0.0
- bool = false
- error = \<nil>

## Funções

Podem o não ter parametros e retorno, como também pode ter mais de um retorno e retorno nomenado

## Operadores

- Aritimeticos
  - \+
  - \- 
  - /
  - \*
  - %

- Atribuição
  - =
  - := (já utiliza a inferencia de tipos)

- Relacionais
  - \>
  - \>= 
  - ==
  - <=
  - <
  - !=

- Logicos
  - &&
  - ||
  - !

- Operadores unarios
  - ++
  - += 
  - --
  - -=
  - *= 
  - /=
  - %=

Em Go não existe operador ternario


## Structs

- É uma coleção de campos
- Semelhante a uma classe, pois no Go não existe classe

```bash

type pessoa struct {
  nome string
  sobrenome string
  idade int
}

// composição
type estudate struct {
  pessoa 
  curso string
  faculdade string
}

```

## Ponteiros

- Salva o endereços de memoria

## Array e Slices

- Array
  - Lista de valores
  - Todos os dados de um array devem ser do mesmo tipo
  - Precisa necessáriamente especificar o tamanho 

- Slices
  - Estrutura baseada no array
  - Flexivel o tamanho
  - Possui e mesma limitação do tipo de dados armazenado que o array

- Make
  - função que serve para alocar um espaço na memoria
  - recebe 3 parametros: o tipo, tamanho e capacidade(opcional quando nao passado assume o mesmo valor do tamanho)

## Maps

- É uma estrutura do tipo chave valor

## Estruturas de Controle

- caso alguma condição seja verdadeira realiza determinado código, caso contrário executa outro conjuto de codigo
- if init ( inicializa uma variavel dentro de uma condição ) só existe dentro do escopo da condição e não pode ser acessada fora da estrutura.

## Switch

- estrutura utilizada quando existe muitas condições, mais legivel que muitos if/else
- `fallthrought`: executa a próxima condição (pouco utilizada)
- não existe `break`