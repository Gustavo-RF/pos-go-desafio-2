# Desafio 2 pos Go
Esse desafio consiste em fazer duas chamadas distintas para duas apis diferentes e aceitar a resposta mais rápida entre elas.

Vamos utilizar conceitos de apis e multithreads para tratar esses casos.

Vamos fazer a chamada para
- https://brasilapi.com.br/api/cep/v1/ + cep
- https://viacep.com.br/ws/ + cep + /json/

Vamos utilizar go routines nas duas chamadas, que acontecerá ao mesmo tempo.

O retorno mais rápido entre as duas será tratado por um select que receberá o primeiro canal que retornou uma resposta.

A resposta será mostrada no terminal, bem como a origem, tempo gasto e endereço recuperado

## Rodando o projeto

  > go mod tidy 
  
  > go run main.go cep

O cep deverá ser enviado via atributo. Deverá ser somente um cep e também deverá ser um cep válido.

### Casos de erro
Caso não envie o cep ou caso envie um cep inválido, o programa é interrompido.
