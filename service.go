package main

import (
	"bufio"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//test()
	production()
}

func production() {

	agent := "main"

	l := logger.NewLogger("./order_service.json", true)

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./")
	if err := v.ReadInConfig(); err != nil {
		// create config file if not exists
		if err := v.WriteConfigAs("config.json"); err != nil {
			l.Error(agent, err.Error())
			os.Exit(1)
		}
	}

	prv, err := getPrivateKey(v)
	if err != nil {
		l.Fatal(agent, err.Error())
	}

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

	ss, err := services.WrapServices(&services.Config{
		FeeServiceURL: "http://localhost:8083/fee",
		DB:            db,
		V:             v,
		L:             l,
		RC:            rc,
		PrvKey:        prv,
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

	ou := app.NewOrderUseCase(rc, s.Repo, ss.ExRepo, ss.PairConf, s.Oc, ss.Deposite, ss.Fee, l)
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

	http.NewRouter(ou, v, rc, l).Run(":8000")

	wg.Wait()

}

func test() {

	agent := "main"

	l := logger.NewLogger("./order_service.json", true)

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./")
	if err := v.ReadInConfig(); err != nil {
		// create config file if not exists
		if err := v.WriteConfigAs("config.json"); err != nil {
			l.Fatal(agent, err.Error())
		}
	}

	prv, err := getPrivateKey(v)
	if err != nil {
		l.Fatal(agent, err.Error())
	}

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

	ss, err := services.WrapServices(&services.Config{
		FeeServiceURL: "http://localhost:8083/fee",
		DB:            db,
		V:             v,
		L:             l,
		RC:            rc,
		PrvKey:        prv,
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

	ou := app.NewOrderUseCase(rc, s.Repo, ss.ExRepo, ss.PairConf, s.Oc, ss.Deposite, ss.Fee, l)
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

	http.NewRouter(ou, v, rc, l).Run(":8081")

	wg.Wait()

}

func getPrivateKey(v *viper.Viper) (*rsa.PrivateKey, error) {
	prf := v.GetString("private_key_file")
	if prf == "" {
		prf = "./private_key.pem"
	}
	privateKeyFile, err := os.Open(prf)
	if err != nil {
		return nil, err
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
