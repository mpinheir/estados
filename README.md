# Estados

O objetivo deste projeto é exemplificar a criação de uma WEB API usando a [linguagem Go](https://golang.org/).

A API retornará os dados pertinentes a cada estado da federação brasileira em formato JSON. 

Os dados retornados pela API são:

```
Nome do Estado
Area
Capital
População
```

Os dados utilizados são baseados em informações de 2017 e estão documentados em [States of Brazil](https://en.wikipedia.org/wiki/States_of_Brazil).



## Prerequisitos

[Instale a linguagem Go](https://golang.org/doc/install) na sua máquina

## Como funciona?

Compile o código

```
go build
```

Execute o programa

```
./estados
```

Teste-1

```
curl localhost:8080/RJ
```

Resultado esperado

```
{"Estado":"Rio de Janeiro","Area":43696.1,"Capital":"Rio de Janeiro","Populacao":16718956}
```

Teste-2

```
curl localhost:8080/TODOS
```

Resultado esperado

```
JSON contendo os dados de todos os estados
```

## Author

[**Marcelo Pinheiro**](https://github.com/mpinheir)

## License

Copie e use como desejado.

## Agradecimentos

* Aos criadores da linguagem Go
* Ao meu amigo [Marco Paganini](https://github.com/marcopaganini) pela paciência e incentivo

