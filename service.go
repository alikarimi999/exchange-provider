package main

import (
	"context"
	"fmt"
	"order_service/internal/app"
	"order_service/internal/delivery/event"
	"order_service/internal/delivery/http"
	"order_service/internal/delivery/services"
	"order_service/internal/delivery/storage"
	"order_service/pkg/logger"
	"order_service/pkg/queue"
	"os"
	"sync"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	test()
	// production()
}

func production() {

	agent := "main"

	l := logger.NewLogger("./order_service.json", true)

	rc := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR")})

	if err := rc.Ping(context.Background()).Err(); err != nil {
		l.Fatal(agent, fmt.Sprintf("failed to ping redis: %s", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), "order_service")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal(agent, fmt.Sprintf("failed to connect to %s", dsn))
	}
	l.Debug(agent, fmt.Sprintf("connected to %s", dsn))
	s := storage.NewStorage(db, rc, l)

	ss := services.WrapServices(&services.Config{
		DepositeServiceURL: os.Getenv("DEPOSITE_SERVICE_URL"),
		FeeServiceURL:      "http://localhost:8083/fee",
	})

	// kucoin := kucoin.NewKucoinExchange(&kucoin.Configs{
	// 	ApiKey:        os.Getenv("KUCOIN_API_KEY"),
	// 	ApiSecret:     os.Getenv("KUCOIN_API_SECRET"),
	// 	ApiPassphrase: os.Getenv("KUCOIN_API_PASSPHRASE"),
	// 	ApiVersion:    "2",
	// 	ApiUrl:        "https://api.kucoin.com",
	// }, r, l)

	wg := &sync.WaitGroup{}

	// wg.Add(1)
	// go kucoin.Run(wg)

	// exs := make(map[string]entity.Exchange)
	// exs["kucoin"] = kucoin

	ou := app.NewOrderUseCase(rc, s.Repo, s.Oc, s.Wc, ss.Deposite, ss.Fee, l)
	wg.Add(1)
	go ou.Run(wg)

	//
	sub := queue.NewSubscriber(fmt.Sprintf("amqp://%s:%s@%s/", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), os.Getenv("RABBITMQ_HOST")), "order_service",
		[]*queue.Topic{
			{Exchange: "deposites", RoutingKey: "deposite.confirmed"},
		})

	wg.Add(1)
	go sub.Run(wg)

	wg.Add(1)
	go event.NewConsumer(ou, sub, l, []string{"deposite.confirmed"}).Run(wg)

	http.NewRouter(ou, l).Run(":8000")

	wg.Wait()

}

func test() {

	agent := "main"

	l := logger.NewLogger("./order_service.json", true)

	rc := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR")})

	if err := rc.Ping(context.Background()).Err(); err != nil {
		l.Fatal(agent, fmt.Sprintf("failed to ping redis: %s", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123", "localhost:3306", "order_service")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal(agent, fmt.Sprintf("failed to connect to %s", dsn))
	}
	l.Debug(agent, fmt.Sprintf("connected to %s", dsn))
	s := storage.NewStorage(db, rc, l)

	ss := services.WrapServices(&services.Config{
		DepositeServiceURL: "http://localhost:8080",
		FeeServiceURL:      "http://localhost:8083/fee",
	})

	// kucoin := kucoin.NewKucoinExchange(&kucoin.Configs{
	// 	ApiKey:        "62b9e1232b968a0001539730",
	// 	ApiSecret:     "dfcbb3c0-c417-498f-a139-c5961e912426",
	// 	ApiPassphrase: "77103121",
	// 	ApiVersion:    "2",
	// 	ApiUrl:        "https://openapi-sandbox.kucoin.com",
	// }, r, l)

	wg := &sync.WaitGroup{}

	// wg.Add(1)
	// go kucoin.Run(wg)

	// exs := make(map[string]entity.Exchange)
	// exs["kucoin"] = kucoin

	ou := app.NewOrderUseCase(rc, s.Repo, s.Oc, s.Wc, ss.Deposite, ss.Fee, l)
	wg.Add(1)
	go ou.Run(wg)

	//
	sub := queue.NewSubscriber(fmt.Sprintf("amqp://%s:%s@%s/", "user_0",
		"user_0a3l3b3qmqnb83rj", "localhost:5672"), "order_service",
		[]*queue.Topic{
			{Exchange: "deposites", RoutingKey: "deposite.confirmed"},
		})

	wg.Add(1)
	go sub.Run(wg)

	wg.Add(1)
	go event.NewConsumer(ou, sub, l, []string{"deposite.confirmed"}).Run(wg)

	http.NewRouter(ou, l).Run(":8081")

	wg.Wait()

}
