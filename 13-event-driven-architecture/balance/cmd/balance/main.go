package main

import (
	"database/sql"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabiopsouza/balance/internal/core/usecase/balance"
	"github.com/fabiopsouza/balance/internal/platform/kafka"

	balanceAdapter "github.com/fabiopsouza/balance/internal/platform/adapters/balance"

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

	balanceRepository := balanceAdapter.NewMySqlClient(db)

	balanceUseCase := balance.NewUseCase(balanceRepository)

	balanceConsumer := kafka.NewConsumer(&configMap, []string{"balance"})
	balanceConsumer.Consume(balanceUseCase.Save)
}
