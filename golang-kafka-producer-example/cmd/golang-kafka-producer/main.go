package main

import (
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/webserver"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/webserver/handlers"
)

// @title Producer Kafka - Studies
// @version 1.0
// @description API to produce a event kafka
// @termsOfService http://swagger.io/terms/

// @contact.name Diego Cordeiro
// @contact.url https://github.com/DiegoJCordeiro/golang-with-kafka-study
// @contact.email diegocordeiro.contatos@gmail.com

// @license.name Diego Cordeiro License
// @license.url  https://github.com/DiegoJCordeiro/golang-with-kafka-study/blob/main/LICENSE

// @host localhost:8080
// @BasePath /
func main() {

	webServer := webserver.NewWebServer(":8080")

	configureHandlersOnWebServer(webServer)

	err := webServer.Start()

	if err != nil {
		panic(err)
	}
}

func configureHandlersOnWebServer(webServer *webserver.WebServer) {

	queryAllPostsUseCase := NewQueryAllPostsUseCase()
	createPostsUseCase := NewCreatePostsUseCase("localhost:9092")

	postsHandler := handlers.NewPostsHandler(createPostsUseCase, queryAllPostsUseCase)

	webServer.AddHandler("POST /v1/posts", postsHandler.CreateHandler)
	webServer.AddHandler("GET /v1/posts", postsHandler.QueryAllHandler)
}
