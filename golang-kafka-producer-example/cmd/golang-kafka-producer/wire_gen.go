// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/database/repository"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/database/sqlc"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/kafka_client"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/usecase"
	"github.com/google/wire"
)

import (
	_ "github.com/mattn/go-sqlite3"
)

// Injectors from wire.go:

func NewCreatePostsUseCase(bootstrapServer string) usecase.CreatePostsUseCase {
	kafkaClientImpl := kafka_client.NewKafkaClient(bootstrapServer)
	createPostsUseCase := usecase.NewCreatePostsUseCase(kafkaClientImpl)
	return createPostsUseCase
}

func NewQueryAllPostsUseCase() usecase.QueryAllPostsUseCase {
	db := NewDatabaseConnectionSQLite()
	postRepositoryImpl := repository.NewPostRepository(db)
	queryAllPostsUseCase := usecase.NewQueryAllPostsUseCase(postRepositoryImpl)
	return queryAllPostsUseCase
}

// wire.go:

var (
	setDatabaseDependency = wire.NewSet(
		NewDatabaseConnectionSQLite, wire.Bind(new(sqlc.DBTX), new(*sql.DB)),
	)

	setPostsRepositoryDependency = wire.NewSet(repository.NewPostRepository, wire.Bind(new(repository.PostRepository), new(*repository.PostRepositoryImpl)))

	setKafkaClientDependency = wire.NewSet(kafka_client.NewKafkaClient, wire.Bind(new(kafka_client.KafkaClient), new(*kafka_client.KafkaClientImpl)))

	setCreatePostsUseCaseDependency = wire.NewSet(wire.Bind(new(usecase.CreatePostsUseCase), new(usecase.CreatePostsUseCaseImpl)))

	setQueryAllPostsUseCaseDependency = wire.NewSet(wire.Bind(new(usecase.QueryAllPostsUseCase), new(usecase.QueryAllPostsUseCaseImpl)))
)

func NewDatabaseConnectionSQLite() *sql.DB {
	conn, _ := sql.Open("sqlite3", "../golang-kafka-db-example.db")

	return conn
}
