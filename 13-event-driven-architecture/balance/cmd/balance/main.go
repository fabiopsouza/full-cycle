package main

import (
	"database/sql"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	balanceUseCase "github.com/fabiopsouza/balance/internal/core/usecase/balance"
	"github.com/fabiopsouza/balance/internal/platform/adapters/inbound/balance"
	balanceOutboundAdapter "github.com/fabiopsouza/balance/internal/platform/adapters/outbound/balance"
	"github.com/fabiopsouza/balance/internal/platform/kafka"
	webserver "github.com/fabiopsouza/balance/internal/platform/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "4000", "balances"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "fcutils",
	}

	balanceRepository := balanceOutboundAdapter.NewMySqlClient(db)

	balanceUseCase := balanceUseCase.NewUseCase(balanceRepository)

	balanceConsumer := kafka.NewConsumer(&configMap, []string{"balance"})
	balanceConsumer.Consume(balanceUseCase.Save)

	balanceHandler := balance.NewHandler(balanceUseCase)

	webserver := webserver.NewWebServer(":3003")
	webserver.AddHandler("/balances/:accountID", balanceHandler.List)

	fmt.Println("Server is running")
	webserver.Start()
}
