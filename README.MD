# API Restful em GO

## Pré-requisitos

- Go
- Postgresql 
- Configurar a variável de ambiente GOPATH

## Inicializando o projeto

Executar os seguintes comandos no terminal para instalar os pacotes necessários para este projeto:
- go get github.com/gorilla/mux        -> URL router
- go get github.com/jinzhu/gorm        -> ORM library
- go get github.com/dgrijalva/jwt-go   -> Utilizado para criação e verificação de JWT tokens
- go get github.com/joho/godotenv      -> Utilizado para carregar arquivos .env para configuração de conexão de banco de dados
- go get github.com/lib/pq             -> Driver de postgres para Go
- go get golang.org/x/crypto/bcrypt    -> Utilizado para criptografia

Para este projeto o arquivo .env.example está pré-configurado, é necessário personalizá-lo de acordo com o seu ambiente para que funcione corretamente. Faça uma cópia dele renomeie o arquivo para .env e preencha as variáveis conforme necessário para realizar a sua conexão com banco de dados.

Após ter preparado todo o ambiente execute os seguintes comandos:
- go build main.go
- go run main.go

Assim que executar o comando run pela primeira vez todas as migrations serão executadas.

O arquivo Funcionarios.csv é utilizado para importação. Para realizar a importação basta enviar o arquivo pela rota "localhost:9002/api/funcionario/import"


## Cadastro de usuário

Assim que estiver rodando é possível acessar as rotas apenas se o bearer token for enviado. Caso contrário é retornado erro 403 com uma mensagem de erro.

Para cadastrar um novo usuário ou até mesmo realizar login de um usuário existente para verificar os dados do mesmo basta seguir os passos abaixo

Para criar um  usuário é necessário enviar a payload pela rota "localhost:9002/api/user/new" com os seguintes campos:
- email(string)
- password(string)
Após cadastrar um usuário é possível verificar os dados dele e o seu token realizando login pela rota "localhost:9002/api/user/login", nesta também é necessário  enviar um JSON com os mesmos campos.

## Exclusão de registros no banco de dados

A exclusão é realizada por meio de softDelete, ou seja, o registro não é de fato excluído da tabela, o que ocorre é que o campo "deleted_at" é carimbado com a data e hora de exclusão.
