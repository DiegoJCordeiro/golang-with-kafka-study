//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/database/repository"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/database/sqlc"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/kafka_client"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/usecase"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

var (
	setDatabaseDependency = wire.NewSet(
		NewDatabaseConnectionSQLite,             // This provides *sql.DB
		wire.Bind(new(sqlc.DBTX), new(*sql.DB)), // Bind *sql.DB to DBTX
	)

	setPostsRepositoryDependency = wire.NewSet(
		repository.NewPostRepository,
		wire.Bind(new(repository.PostRepository), new(*repository.PostRepositoryImpl)),
	)

	setKafkaClientDependency = wire.NewSet(
		kafka_client.NewKafkaClient,
		wire.Bind(new(kafka_client.KafkaClient), new(*kafka_client.KafkaClientImpl)),
	)

	setCreatePostsUseCaseDependency = wire.NewSet(
		wire.Bind(new(usecase.CreatePostsUseCase), new(usecase.CreatePostsUseCaseImpl)),
	)

	setQueryAllPostsUseCaseDependency = wire.NewSet(
		wire.Bind(new(usecase.QueryAllPostsUseCase), new(usecase.QueryAllPostsUseCaseImpl)),
	)
)

func NewDatabaseConnectionSQLite() *sql.DB {
	conn, _ := sql.Open("sqlite3", "../golang-kafka-db-example.db")

	return conn
}

func NewCreatePostsUseCase(bootstrapServer string) usecase.CreatePostsUseCase {

	wire.Build(
		setKafkaClientDependency,
		usecase.NewCreatePostsUseCase,
	)

	return &usecase.CreatePostsUseCaseImpl{}
}

func NewQueryAllPostsUseCase() usecase.QueryAllPostsUseCase {

	wire.Build(
		setDatabaseDependency,
		setPostsRepositoryDependency,
		usecase.NewQueryAllPostsUseCase,
	)

	return &usecase.QueryAllPostsUseCaseImpl{}
}
