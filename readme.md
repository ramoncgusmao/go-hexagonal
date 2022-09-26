
pra rodar o projeto na sua maquina precisa só ter o docker instalado e rodar o comando

```
docker-compose up -d
```

para entrar dentro do container você usa o comando 

```
docker exec -it appproduct bash
```

dentro do container para executar os testes voce roda o comando
```
go test ./...
```

caso queira ligar o servidor web utiliza o comando cli do cobra pra ligar o servidor

```
go run main.go http
```

para criar um novo comando no cli use o comando

```
cobra-cli add [comando]
```

libs utilizadas

- **govalidator** - para validação - github.com/asaskevich/govalidator 
- **negroni** - para logs - github.com/codegangsta/negroni 
- **mock** - para mock - github.com/golang/mock 
- **mux** - para servidor http - github.com/gorilla/mux 
- **go-sqlite3** - para banco dedados em memoria - github.com/mattn/go-sqlite3 
- **go.uuid** - para gerar uuid - github.com/satori/go.uuid v1.2.0
- **cobra** - para cli(command line interface) - github.com/spf13/cobra v1.5.0
- **testify** - para testes - github.com/stretchr/testify v1.8.0
