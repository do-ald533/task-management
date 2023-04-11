# Boilerplate - Go API

## Quick start

1. Copie `.env.example` para `.env` e preencha com os valores das suas variaveis de ambiente.

2. Gerar documentação do swagger:

> Para Instalar `swag`:

```bash
go install github.com/swaggo/swag/cmd/swag@latest && swag init  --output swagger/
```

3. Para utilizar um live reload na aplicação, baixe o pacote `air`.

> Para Instalar `air`:

```bash
go install github.com/cosmtrek/air@latest
```

4. Para executar o live reload:

```bash
air run .
```

5. Go to API Docs page (Swagger): [localhost:5000/docs/index.html](http://localhost:5000/docs/index.html)

# Setup SNS e SQS utilizando localstack

Pré-requisitos: `aws cli`, `docker`, `docker-compose`

Execute o `docker-compose` e crie um perfil no arquivo de credenciais da AWS na sua máquina.

No Linux, por exemplo, fica em `~/.aws/credentials`.

Abra esse arquivo e adicione o seguinte trecho de código:

```
[localstack]
aws_access_key_id = fake-access-key
aws_secret_access_key = fake-secret-key
```

Agora, sempre que for utilizar os comandos da AWS CLI, lembre-se de colocar a flag `--profile=localstack`.

## Criar um tópico SNS:

```bash
aws --profile=localstack --endpoint-url=http://localhost:4566 sns create-topic --name nome-do-seu-topico
```

Isso retornará o ARN do tópico criado.

Vai retornar o arn do tópico criado.

## Criar uma fila SQS:

```bash
aws --profile=localstack --endpoint-url=http://localhost:4566 sqs create-queue --queue-name nome-da-sua-fila
```

Isso retornará o endereço HTTP da sua fila. Você precisará dele para descobrir o endereço ARN da sua fila para realizar a assinatura dela em um tópico SNS.

## Criar a assinatura de uma fila no tópico SNS:

```bash
aws --profile=localstack --endpoint-url=http://localhost:4566 sqs get-queue-attributes --queue-url http://localhost:4566/000000000000/nome-da-sua-fila --attribute-names QueueArn
```

Com o arn da sua fila em mãos basta criar a assinatura.

```bash
aws --profile=localstack --endpoint-url=http://localhost:4566 sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:nome-do-seu-topico --protocol sqs --notification-endpoint arn:aws:sqs:us-east-1:000000000000:nome-da-sua-fila
```

Essas instruções detalham como configurar e executar sua aplicação localmente, incluindo a configuração do SNS e SQS usando o LocalStack. Seguindo as etapas acima, você pode executar sua aplicação com live reload e configurar os serviços da AWS necessários para testar suas funcionalidades localmente.

# Como criar novos módulos

Siga os passos a seguir para criar novos módulos, mantendo a organização e estrutura do projeto.

## Estrutura de diretórios e arquivos

Ao criar novos módulos, siga a estrutura de diretórios e arquivos sugerida abaixo para manter a organização do projeto:

```bash
├── internal
│   ├── modules
│   │   └── books
│   │       ├── books.go
│   │       ├── controllers
│   │       │   ├── controllers.go
│   │       │   ├── book_store_controller.go
│   │       │   └── ...
│   │       ├── dtos
│   │       │   ├── books_requests.go
│   │       │   └── book_response.go
│   │       ├── entities
│   │       │   └── books.go
│   │       ├── repositories
│   │       │   ├── repositories.go
│   │       │   ├── store_book_repository.go
│   │       │   └── ...
│   │       └── services
│   │           ├── services.go
│   │           ├── store_books_service.go
│   │           └── ...
├── test
│   ├── mocks
│   └── units
│       └── books
│           ├── controllers
│           │   ├── controllers_test.go
│           │   ├── books_controller_store_test.go
│           │   └── ...
│           └── mocks
│               ├── books_data_mock.go
│               ├── books_repository_mock.go
│               └── books_service_mock.go
```

- **internal/modules**: A pasta modules contém os módulos da aplicação. Adicione uma nova pasta para cada módulo, como books neste exemplo.
- **controllers**: Aqui estão os controladores do módulo, responsáveis por gerenciar as requisições e respostas HTTP. Inclua um arquivo para cada controlador específico, como book_store_controller.go.
- **dtos**: A pasta dtos contém os objetos de transferência de dados (Data Transfer Objects) para as requisições e respostas. Adicione arquivos para cada objeto específico, como books_requests.go e book_response.go.
- **entities**: Aqui estão as entidades do módulo, que representam as estruturas de dados do domínio. Neste exemplo, temos apenas uma entidade, books.go.
- **repositories**: A pasta repositories contém os repositórios do módulo, responsáveis por realizar operações no banco de dados. Adicione arquivos para cada repositório específico, como store_book_repository.go.
- **services**: Aqui estão os serviços do módulo, que contêm a lógica de negócio e a conexão entre os controladores e repositórios. Inclua um arquivo para cada serviço específico, como store_books_service.go.

- **test**: A pasta test contém os testes e mocks do projeto. Ela é dividida em duas subpastas, mocks e units.

  - **mocks**: A pasta mocks armazena arquivos de mocks globais que podem ser usados em diferentes módulos do projeto.

  - **units**: A pasta units contém os testes unitários organizados por módulo. Neste exemplo, temos a pasta books representando o módulo de livros. Dentro dela, você encontrará as pastas e arquivos relacionados aos testes unitários do módulo.

## DTOs (Data Transfer Objects)

Os DTOs são objetos de transferência de dados que permitem a validação e a transformação dos dados de entrada e saída entre os diferentes componentes da aplicação. Ao usar o pacote [go-playground/validator](https://github.com/go-playground/validator), é possível garantir que os dados de entrada sejam validados antes de serem processados.

Lembre-se de criar um DTO de entrada e outro de saída para cada controlador em seu módulo.

### Exemplo de DTO de entrada

O exemplo abaixo demonstra como criar um DTO de entrada para armazenar um livro. Este DTO inclui validações usando as tags de estrutura do pacote **`validator`**.

```go
// /internal/modules/books/dtos/books_requests.go
package dtos

type BookStoreRequest struct {
	Title     string    `json:"title" validate:"required,lte=255"`
	Author    string    `json:"author" validate:"required,lte=255"`
}
```

### Exemplo de DTO de saída

O exemplo abaixo demonstra como criar um DTO de saída para exibir informações de um livro após o armazenamento. Observe que este DTO não inclui validações, pois é usado apenas para exibir os dados processados.

```go
// /internal/modules/books/dtos/books_responses.go
package dtos

type BookStoreResponse struct {
	ID         *uuid.UUID `json:"id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Title      string     `json:"title"`
	Author     string     `json:"author"`
}
```

## Entities (opcional)

As entities são usadas para representar o esquema de sua entidade no banco de dados. É importante declarar as tags para os campos, tanto para `bson` quanto para `json`, facilitando a serialização e desserialização dos dados.

Ao utilizar DTOs de entrada e saída, você garante a separação de responsabilidades e a validação adequada dos dados, além de facilitar a manutenção e a evolução do código.

```go
// /internal/modules/books/entities/book.go
package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID         *uuid.UUID `bson:"_id" json:"_id"`
	CreatedAt  *time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `bson:"updated_at" json:"updated_at"`
	Title      string     `bson:"title" json:"title"`
	Author     string     `bson:"author" json:"author"`
}

func (b Book) Value() primitive.M {
	byte, _ := bson.Marshal(b)

	var updated bson.M
	bson.Unmarshal(byte, &updated)

	return updated
}

func (b Book) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return bson.Unmarshal(j, &b)
}
```

## Repositories (opcional)

A camada de repositórios é responsável pela comunicação direta com o banco de dados. Embora não seja obrigatória, sua utilização é recomendada quando a lógica de acesso ao banco de dados é complexa, e pode ser separada da camada de services.

Ao utilizar a camada de repositórios, você promove uma separação de responsabilidades e facilita a manutenção e a evolução do código.

> Obrigatório!!

- Sempre crie uma interface declarando o contrato dos métodos presentes.
- Struct para injetar as dependências de conexão com o banco de dados.
- Construtor para instanciar a struct do repositório.

> Ao criar qualquer método, lembre-se de sempre definir um contexto controlando, no mínimo, um timeout de conexão.

```go
// /internal/modules/books/repositories/store_book_repository.go
package repositories

import (
	"context"
	"goapi/internal/modules/books/entities"
	"goapi/pkg/errors"
	"time"
)

func (repo *BookRepository) Store(b *entities.Book) (*entities.Book, error) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	value := b.Value()

	_, err := repo.BooksCollection.InsertOne(ctxTimeout, value)

	if err != nil {
		return nil, errors.BadRequest(errors.Message{
			"error": true,
			"msg":   "Can't create this book!",
		})
	}

	return b, nil
}
```

```go
// /internal/modules/books/repositories/repositories.go
package repositories

import (
	"goapi/infrastructure/database"
	"goapi/internal/modules/books/entities"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepositoryImpl interface {
	Store(b *entities.Book) (*entities.Book, error)
}

type BookRepository struct {
	BooksCollection *mongo.Collection
}

func NewBookRepository() (BookRepositoryImpl, error) {
	db, err := database.OpenDBConnection("mongodb")
	if err != nil {
    // Retorna os erros de conexão com o banco de dados.
    // Já está configurado para cancelar o processo e retornar
    // com o status code e mensagem corretos.
		return nil, err
	}

	database := db.Mongo.Database("app")
	booksCollection := database.Collection("books")

	return &BookRepository{booksCollection}, nil
}
```

## Services

A camada de services, também conhecida como usecases, tem a função de orquestrar cada um dos recursos e disponibilizar um resultado.

A camada de services é responsável por gerenciar a lógica do negócio e a comunicação entre as diferentes partes do sistema, como a camada de repositórios e as APIs externas. Ao utilizar essa camada, você garante uma melhor organização do código e facilita a manutenção e evolução do projeto.

> Obrigatório!!

- Sempre crie uma interface declarando o contrato dos métodos presentes.
- Struct para injetar as dependências dos recursos que planeja utilizar.
- Construtor para instanciar a struct da service.

```go
// /internal/modules/books/services/services.go
package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/repositories"
	"goapi/pkg/aws/sns"

	"github.com/google/uuid"
)

type BookServicesImpl interface {
	Store(userId uuid.UUID, body *dto.BookStoreRequest) (*dto.BookStoreResponse, error)
}

type BookService struct {
	BookRepository repositories.BookRepositoryImpl
	SnsService     sns.SNSService
}

func NewBookService(
	bookRepository repositories.BookRepositoryImpl,
	snsService sns.SNSService,
) BookServicesImpl {

	return &BookService{bookRepository, snsService}
}
```

```go
// /internal/modules/books/services/store_book_service.go
package services

import (
	"goapi/internal/modules/books/dto"
	"goapi/internal/modules/books/entities"
	"goapi/pkg/aws/sns/snsactions"
	"goapi/pkg/convert"
	"os"
	"time"

	"github.com/google/uuid"
)

func (serv BookService) Store(userId uuid.UUID, body *dto.BookStoreRequest) (*dto.BookStoreResponse, error) {
	book := entities.Book{}
	convert.ToStruct(body, &book)

	id := uuid.New()
	status := 1
	now := time.Now()

	book.ID = &id
	book.CreatedAt = &now
	book.BookStatus = &status
	book.UserID = &userId

	newBook, err := serv.BookRepository.Store(&book)
	if err != nil {
		return nil, err

	}

	response := dto.BookStoreResponse{}
	convert.ToStruct(newBook, &response)

	serv.SnsService.Setup(os.Getenv("SNS_CREATED_BOOKS"))
	serv.SnsService.Publish(&response, snsactions.CREATION)

	return &response, nil
}
```

## Controllers

A camada do controller é responsável por receber e tratar as requisições HTTP, fazer a comunicação com a camada de serviços e retornar as respostas adequadas. Ao utilizar esta camada, garantimos uma melhor separação das responsabilidades, facilitando a manutenção e a evolução do projeto. Além disso, o uso de documentação automática, como o Swagger, torna mais fácil a compreensão e a integração com outras partes do sistema.

No controller, assim como nas outras camadas, vamos criar um construtor que receberá algumas dependências e tratará informações básicas, como: extrair dados do token, fazer o parse do corpo da requisição ou outros parâmetros para uma struct, executar a validação desses dados e chamar a service (usecases).

Também declaramos alguns parâmetros em comentários no método para criar a documentação no Swagger automaticamente.

```go
// /internal/modules/books/controllers/controllers.go
package controllers

import (
	"goapi/internal/modules/books/services"
)

type BooksController struct {
	bookServices services.BookServicesImpl
}

func NewBooksController(
	bookServices services.BookServicesImpl,
) *BooksController {
	return &BooksController{bookServices}
}
```

```go
// /internal/modules/books/controllers/book_store_controller.go
package controllers

import (
	"goapi/internal/modules/books/dto"
	"goapi/pkg/errors"
	"goapi/pkg/jwt"
	"goapi/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// @Description Create a new book.
// @Summary create a new book
// @Tags Books
// @Accept json
// @Produce json
// @Param request body dto.BookStoreRequest true "Request Body"
// @Success 200 {object} dto.BookStoreResponse
// @Security ApiKeyAuth
// @Router /books [post]
func (ctrl *BooksController) Store(req *fiber.Ctx) error {
	claims, err := jwt.ExtractTokenMetadata(req)
	if err != nil {
		return errors.ErrorResponse(req, err)
	}

	book := dto.BookStoreRequest{}
	req.BodyParser(&book)

	validate := utils.NewValidator()

	if err := validate.Struct(book); err != nil {
		return req.Status(fiber.StatusUnprocessableEntity).JSON(utils.ValidatorErrors(err))
	}

	response, err := ctrl.bookServices.Store(claims.UserID, &book)
	if err != nil {
		return errors.ErrorResponse(req, err)
	}

	return req.Status(fiber.StatusCreated).JSON(response)
}
```

## Rotas

A camada de rotas tem a função de mapear as requisições HTTP para os controladores e métodos específicos, além de aplicar middlewares para garantir a segurança e a consistência das requisições. Ao utilizar essa camada, é possível organizar melhor o projeto e facilitar a manutenção e evolução das funcionalidades.

No arquivo de rotas, você pode declarar duas funções, PrivateRoutes e PublicRoutes, para facilitar a visualização de quais rotas são protegidas e quais são públicas. Em seguida, chame essas funções no declarador de rotas geral do projeto

```go
// /internal/modules/books/books.go

package users

import (
	"goapi/internal/middleware"
	"goapi/internal/modules/books/controllers"
	"goapi/internal/modules/books/repositories"
	"goapi/internal/modules/books/services"
	"goapi/pkg/permissions"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(route fiber.Router) {
	booksRepositories, _ := repositories.NewBooksRepositories()
	booksServices := services.NewBooksServices(booksRepositories)

	controllers := controllers.NewBooksControllers(
		booksRepositories,
		booksServices,
	)

	route.Post("/books", middleware.Credentials(permissions.BookCreateCredential), controllers.Store)
}
```

```go
// /internal/routes/private_routes.go
package routes

import (
	"goapi/internal/middleware"
	"goapi/internal/modules/books"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	api := app.Group("", middleware.JWTProtected())

	books.PrivateRoutes(api)
  // ...
}
```

## Testes

O teste mais básico que pode ser feito é o teste unitário dos controladores.

Para criar um teste de um controlador, precisamos criar algumas estruturas iniciais para facilitar o processo e reutilizar alguns dados.

Criamos uma struct com um contexto do Fiber para conseguir injetar certos dados para nosso controlador, como body, headers de autenticação, etc.

### Configuração inicial para testar um controlador

```go
// /test/units/books/controllers/controllers_tests.go
package controllers

import (
	"goapi/internal/modules/books/controllers"
	"goapi/test/units/books/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
)

type BooksControllerTestSuite struct {
	ctx        *fiber.Ctx
	controller *controllers.BooksController
	service    *mocks.BookServiceMock
	repository *mocks.BookRepositoryMock
	suite.Suite
}

func (suite *BooksControllerTestSuite) SetupTest() {
	suite.service = &mocks.BookServiceMock{}
	suite.repository = &mocks.BookRepositoryMock{}
	suite.controller = controllers.NewBooksController(suite.service)

	app := fiber.New()

	suite.ctx = app.AcquireCtx(&fasthttp.RequestCtx{})
}

func TestBooksControllerSuite(t *testing.T) {
	suite.Run(t, new(BooksControllerTestSuite))
	// ...
}
```

Neste exemplo, utilizamos o pacote `github.com/stretchr/testify/suite` para organizar os testes em suítes e facilitar a configuração e execução dos testes. A estrutura `BooksControllerTestSuite` é criada para manter o contexto do Fiber, o controlador e as instâncias de serviço e repositório mockados. Isso permite que você possa injetar dependências mockadas nos controladores e verificar o comportamento correto dos métodos.

O método `SetupTest` é responsável por inicializar as instâncias e configurar o contexto do Fiber antes de cada teste. A função `TestBooksControllerSuite` é utilizada para executar a suíte de testes, garantindo que todos os testes sejam executados corretamente.

Com essa configuração inicial, você pode começar a escrever testes unitários específicos para cada método do controlador, garantindo que eles funcionem corretamente e atendam às expectativas.

### Criar os mocks

Vamos criar mocks para injetar nas dependências do nosso controlador, neste caso, os métodos de serviço e também os DTOs de entrada e saída do nosso controlador.

```go
// /test/units/books/mocks/books_data.go
package mocks

import (
	"goapi/internal/modules/books/dto"
	"goapi/pkg/jwt"

	"github.com/google/uuid"
)

var (
	idMock = uuid.New()
	rating = 8
)

var UserId = uuid.New()
var UserLogged, _ = jwt.GenerateNewTokens(UserId.String(), "admin", []string{"book:create", "book:update", "book:delete"})

var StoreBookMock = dto.BookStoreRequest{
	Title:  "Title",
	Author: "Jaum",
}

var BookMock = dto.BookStoreResponse{
	ID:     &idMock,
	Title:  "Title",
	Author: "Jaum",
}
```

Agora vamos criar a implementação do mock do nosso serviço.

```go
// /test/units/books/mocks/books_service_mock.go
package mocks

import (
	"goapi/internal/modules/books/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BookServiceMock struct {
	mock.Mock
}

func (m *BookServiceMock) Store(userId uuid.UUID, body *dto.BookStoreRequest) (*dto.BookResponse, error) {
	args := m.Called(userId, body)
	return args.Get(0).(*dto.BookResponse), args.Error(1)
}

func (m *BookServiceMock) Update(id *uuid.UUID, body *dto.BookUpdateRequest) (*dto.BookResponse, error) {
	args := m.Called(id, body)
	return args.Get(0).(*dto.BookResponse), args.Error(1)
}

func (m *BookServiceMock) Show(id uuid.UUID) (*dto.BookResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.BookResponse), args.Error(1)
}

func (m *BookServiceMock) Index(limit, page int64) (*dto.BookIndexResponse, error) {
	args := m.Called(limit, page)
	return args.Get(0).(*dto.BookIndexResponse), args.Error(1)
}

func (m *BookServiceMock) Delete(id *uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
```

Os mocks acima são criados para substituir a implementação real do serviço e dos DTOs durante a execução dos testes. Eles ajudam a isolar o comportamento dos controladores e garantir que os testes sejam independentes da implementação do serviço.

O pacote `github.com/stretchr/testify/mock` é usado para criar e gerenciar os mocks. Isso permite que você defina expectativas e retornos para os métodos chamados, facilitando a criação de cenários de teste específicos.

### Criar os testes

Agora que temos a estrutura e os mocks configurados, basta criar os testes como métodos para a struct criada `BooksControllerTestSuite`.

```go
// /test/units/books/controllers/books_controller_store_test.go
package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/test/units/books/mocks"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func (suite *BooksControllerTestSuite) TestSuccessStoreBook() {
	suite.ctx.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", mocks.UserLogged.Access))
	suite.ctx.Request().Header.SetContentType(fiber.MIMEApplicationJSON)

	body, _ := json.Marshal(mocks.StoreBookMock)
	suite.ctx.Request().SetBody(body)

	suite.service.On("Store", mocks.UserId, &mocks.StoreBookMock).Return(&mocks.BookMock, nil)
	err := suite.controller.Store(suite.ctx)

	expectedResponse, _ := json.Marshal(mocks.BookMock)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusCreated, suite.ctx.Response().StatusCode())
	assert.JSONEq(suite.T(), string(expectedResponse), string(suite.ctx.Response().Body()))
}
```

No exemplo acima, criamos um teste para verificar se o método `Store` do controlador `BooksController` retorna a resposta esperada e cria um livro com sucesso. Configuramos o cabeçalho de autorização e o tipo de conteúdo do contexto da requisição, bem como o corpo da requisição. Além disso, definimos a expectativa para o método `Store` do serviço mock e verificamos se a resposta é igual à resposta esperada.

Ao executar os testes, você poderá verificar se o controlador está funcionando conforme o esperado, independentemente da implementação do serviço e das interações com o banco de dados ou outras dependências. Isso facilita a manutenção e a garantia da qualidade do código.

# Como criar novos consumers

Ao criar novos consumers em Golang, você deve executar três etapas:

- Configuração inicial da estrutura do consumer.
- Registro do consumer no Consumer Manager.
- Adição da lógica para processar a mensagem.

## Estrutura de diretórios e arquivos

```bash
├── consumers
│   ├── bookscreatedconsumers
│   │   ├── books_created_process.go
│   │   └── books_created_consumer.go
│   ├── ...
│   ├── consumers.go
│   └── manager
│       └── consumer_manager.go
```

## Configuração padrão

Dentro do diretório `consumers`, adicione seu consumer customizado. Crie um diretório com um nome representativo para o seu consumer e, dentro dele, crie os seguintes arquivos:

```go
// /consumers/bookscreatedconsumers/books_created_consumer.go
package bookscreatedconsumers

import (
	"goapi/consumers"
	sqsconsumer "goapi/pkg/aws/sqs"
	"goapi/pkg/logging"
	"os"
)

type BooksCreatedConsumer struct {
	sqsConsumer *sqsconsumer.SqsConsumer
	stop        chan bool
}

func NewBooksCreatedConsumer() consumers.Consumer {
	var (
		queueURL        = os.Getenv("SQS_CREATED_BOOKS")
		batchSize int64 = 1
		waitTime  int64 = 20
	)

	baseConsumer := sqsconsumer.NewSqsConsumer(queueURL, batchSize, waitTime)
	return &BooksCreatedConsumer{
		sqsConsumer: baseConsumer,
		stop:        make(chan bool),
	}
}

func (c *BooksCreatedConsumer) Consume() error {
	logging.Info("Starting BooksCreatedConsumer")
	getMessage := func() (interface{}, error) {
		return c.sqsConsumer.GetMessage()
	}

	return consumers.ConsumeMessages(c, getMessage, c.stop)
}

func (c *BooksCreatedConsumer) Stop() {
	logging.Info("Stopping BooksCreatedConsumer")
	c.stop <- true
}
```

Cada consumer seguirá essa mesma estrutura. Portanto, altere o nome da struct, do construtor e a URL da fila, conforme necessário.

## Processamento da mensagem

Para processar a mensagem, adicione um método `ProcessMessage` à struct do consumer. Esse método será responsável por processar as mensagens da fila.

Você pode adicionar esse método no mesmo arquivo ou em um arquivo separado. No exemplo a seguir, o método é adicionado em um arquivo separado para facilitar a leitura:

```go
// /consumers/bookscreatedconsumers/books_created_process.go
package bookscreatedconsumers

import (
	"goapi/pkg/logging"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func (c *BooksCreatedConsumer) ProcessMessage(msg interface{}) error {
	message := msg.(*sqs.Message)

	logging.Info("Processing SQS Message:", *message.MessageId)
	logging.Info("SQS Body:", *message.Body)

	// A lógica de processamento da mensagem relacionada vai aqui
	// ...

	// Após processar a mensagem, delegue a remoção da mensagem ao consumer base
	return c.sqsConsumer.DeleteMessage(message.ReceiptHandle)
}
```

## Registro do consumer no Consumer Manager

Para registrar seu consumer no Consumer Manager, você deve adicionar o consumer ao método `RegisterConsumers`. Isso garante que o processo de inicialização e encerramento do consumer seja gerenciado junto com os demais consumers.

Acesse o arquivo `consumers/manager/consumer_manager.go` e adicione seu consumer no método `RegisterConsumers`.

```go
package consumers

import (
	"goapi/consumers/bookscreatedconsumers"
  // ...
)
// ...

func (cm *ConsumerManager) RegisterConsumers() {
	booksConsumer := bookscreatedconsumers.NewBooksCreatedConsumer()
	cm.AddConsumer(booksConsumer)

	// Adicione outros consumers aqui
}

// ...
```

Com essa etapa concluída, seu consumer está registrado no Consumer Manager e pronto para ser gerenciado junto com os demais consumers. Isso facilita o controle dos processos de inicialização e encerramento de cada consumer.

# 📦 Features Roadmap

- [x] Deployment to aws
- [x] Http request implementation
- [x] Logs implementation
- [x] Implementation AWS stuffs, like: SNS, SQS
- [x] Documentation - how create new modules
- [ ] Documentation - how connect to database
- [ ] Documentation - how private route by credentials
- [ ] Documentation - how create a new token
- [ ] Documentation - how generate swagger documentation
- [x] JWT Token generator
- [x] Middleware to valid credentials or roles (provided by JWT token)
- [x] Support to MongoDb, MySql, Postgres, Redis
- [x] Pagination with MongoDb
- [x] Swagger generator
- [x] Struct validator who handle path and validation tag on response
- [x] Multi routes implementation, not found, swagger, public, private
- [x] Dockerfile to execute binary file
