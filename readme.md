# Go Lang

## Hello world

- Um arquivo executável precisa seguir a assinatura do arquivo `hello.go` contendo o `package main`, `imports` e `func main`.
- Para rodar o programa executa o comando `go run <nome do arquivo>`.

## Pacotes

- Pacotes são um conjunto de arquivos que estão no mesmo diretório.
- Quando trabalha com pacotes, precisar iniciar o modulo, para isso, basta executar o comando `go mod init <nome do modulo>`.
- Para realizar o import primeiro vem o nome do modulo e depois o nome do pacote `ex: "module/auxiliars"`.
- Note: Funções com letras minusculas são visiveis apenas dentro do prórpio pacote, para ser exportada é necessário utilizar letras maiusculas.
- E é uma boa prática escrever um comentário em cima de funções que são exportadas informando o que fazem!
- Para utilizar pacotes externos utiliza o comando `go get <url do pacote>`. Automaticamente será criado um arquivo `go.sum`
- Para remover uma dependecia externa que não está sendo utilizada executa o comando `go mod tidy`.

## Variaveis

- Em Go não é permitido criar uma variavel e não utilizá-la. 
- Além disso, é uma linguagem fortemente tipada.

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

- Podem o não ter parametros e retorno, como também pode ter mais de um retorno e retorno nomenados.
- Funções variádicas são funções que recebem n parametros, sendo passado como o último parametro da função e só podendo existir um.

- Funções recursivas são funcões que dependem da sua propria execusão. É muito importante que tenha uma condição de parada para que não ocorra o estouro de pilha,e não é recomendada em caso que precise de muitas execuções.

- `defer` clausula utilizada para adiar a execução de uma função. Muito utilizada para encerrar conexões com banco de dados.

- `panic` encerra o programa, mas antes de encerrar o programa executa tudo que estiver em `defer`.

- `recover` permite recuperar um programa que está em `panic`.

- `Closure` funções que refereciam variáveis que estão fora do seu escopo.

- `inti()` função que é executada antes da função `main` podendo ter uma por arquivo.
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
  - \--
  - -=
  - *= 
  - /=
  - %=

- Em Go não existe operador ternário

## Structs

- É uma coleção de campos
- Semelhante a uma classe, em Go não existe classe!

```bash

type pessoa struct {
  nome string
  sobrenome string
  idade int
}

// composição ou embeding
type estudate struct {
  pessoa 
  curso string
  faculdade string
}

```

## Ponteiros

- Salva o endereços de memória

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
  - Função que serve para alocar um espaço na memoria
  - Recebe 3 parametros: o tipo, tamanho e capacidade(opcional; quando nao passado assume o mesmo valor do tamanho)

## Maps

- É uma estrutura do tipo chave valor
- Mais rigidas em relação aos tipos que a Struct

## Estruturas de Controle

- Caso alguma condição seja verdadeira realiza determinado código, caso contrário executa outro conjuto de codigo
- if init ( inicializa uma variavel dentro de uma condição ) só existe dentro do escopo da condição e não pode ser acessada fora da estrutura.

## Switch

- Estrutura utilizada quando existe muitas condições, mais legivel que muitos if/else
- `fallthrought` executa a próxima condição (pouco utilizada)
- Não existe clausula `break`

## Loops

- São estruturas de repetição
- Só existe o `for`
- É possivel interar sob `maps`, mas não é permitido interar em `structs`
- Para sair de um loop pode-se utilizar a clasula `break`

## Métodos

- Permite "ações" as `struct`.

## Concorrência

- Paralelismo != Concorrência.
- Paralelismo processos que ocorrem simultanemanente em processadores que mais de um nucleo.
- Concorrencia podem ou não ocorre de forma simultanea em processadores que tem mais de um nucleo, mas em processadores que so existem um nucleo ocorre o revesamento dos processos, deste modo um processo não precisa experar o outro acabar para iniciar.

## Goroutine

- Colocar a clausula `go` antes de uma função ou metodo.

## Wait Group

- Metodo para sincronizar goroutines.
- waitGroup.Add(x) informa quantas goroutines serão executadas
- waitGroup.Done() informa que a goroutine foi finalizada e remove um de waitGroup.Add()
- waitGroup.Wait() indica que existem goroutines para ser executadas antes de encerrar o programa.

## Canais

- Forma mais utilizada para sincronizar as goroutines
- É um canal de comunicação que pode tanto enviar como receber dados
- Canais bloqueiam a execução do programa


## Canais com Buffer

- É especificado uma capacidade
- Nesse caso quando uma capacidade é espeficicada não ocorre o bloqueio `deadlock`

## Select

- Muito parecida com o switch mas utilizada para concorrência.