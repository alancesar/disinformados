# DISinformados
Códigos utilizados na apresentação "Go em 15 minutos"

## concorrencia-async
Executa GoRoutines de forma assíncrona

```sh
go run main.go
```

## concorrencia-sync
Executa GoRoutines sincronizadas, utilizando canais

```sh
go run main.go
```

## processamento-paralelo
Lê um arquivo de texto contendo 1000 valores numéricos e executa cálculos utilizando processamento paralelo.

**Java**
```sh
javac Main.java
java Main
```

**NodeJS**
```sh
node main.js
```

**Go**
```sh
go run main.go
```

## rest-client
Executa uma chamada a um serviço REST e converte a resposta de `JSON` para uma `Stuct`

```sh
go run main.go
```

## teste-carga
Executa um teste de carga em um servidor HTTP.

**Configurar limite de requets do SO**
```sh
ulimit -n 50000
```

**Iniciar o servidor**
```sh
node server.js
```

**Executar o script**
```sh
go run main.go
```

## web-server
Instancia um servidor Web.

```sh
go run main.go
```

Para acessar, abra o endereço `http://localhost:8080` em seu navegador
