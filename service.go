package main

import (
	"bufio"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"exchange-provider/internal/app"
	"exchange-provider/internal/delivery/database"
	"exchange-provider/internal/delivery/http"
	"exchange-provider/internal/delivery/services"
	"exchange-provider/internal/delivery/services/pairconf"
	"exchange-provider/pkg/logger"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	test()
	// production()
}

func production() {

	agent := "main"

	l := logger.NewLogger("./service.log", true)

	user := os.Getenv("ADMIN_USER")
	if user == "" {
		l.Fatal(agent, "You must set ADMIN_USER environment variable")
	}
	pass := os.Getenv("ADMIN_PASS")
	if pass == "" {
		l.Fatal(agent, "You must set ADMIN_PASS environment variable")
	}

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

	uri := fmt.Sprintf("mongodb://%s/?maxPoolSize=20&w=majority", os.Getenv("MONGO_Address"))
	client, err := mongo.Connect(context.TODO(), options.Client().SetTimeout(5*time.Second).ApplyURI(uri))
	if err != nil {
		l.Fatal(agent, err.Error())
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		l.Fatal(agent, "failed to ping mongo")
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			l.Fatal(agent, err.Error())
		}
	}()
	db := client.Database("exchange-provider")
	l.Debug(agent, "connected to mongo")

	repo, err := database.NewOrderRepo(db, l)
	if err != nil {
		l.Fatal(agent, err.Error())
	}
	pairs := pairconf.NewPairRepo(l)
	ss, err := services.WrapServices(&services.Config{
		DB:     db,
		Pairs:  pairs,
		V:      v,
		L:      l,
		PrvKey: prv,
	})

	if err != nil {
		l.Fatal(agent, err.Error())
	}

	ou := app.NewOrderUseCase(pairs, repo, ss.ExchangeRepo, ss.WalletStore,
		ss.PairConfigs, ss.FeeService, l)

	go ou.Run()
	if err := http.NewRouter(ou, pairs, v, l, user, pass).Run(":8000"); err != nil {
		l.Fatal(agent, err.Error())
	}
}

func test() {

	agent := "main"

	l := logger.NewLogger("./service.log", true)

	user := os.Getenv("ADMIN_USER")
	if user == "" {
		l.Fatal(agent, "You must set ADMIN_USER environment variable")
	}
	pass := os.Getenv("ADMIN_PASS")
	if pass == "" {
		l.Fatal(agent, "You must set ADMIN_PASS environment variable")
	}

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

	uri := "mongodb://root:123@127.0.0.1:27017/?maxPoolSize=20&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().SetTimeout(5*time.Second).ApplyURI(uri))
	if err != nil {
		l.Fatal(agent, err.Error())
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		l.Fatal(agent, "failed to ping mongo")
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			l.Fatal(agent, err.Error())
		}
	}()
	db := client.Database("exchange-provider")
	l.Debug(agent, "connected to mongo")

	repo, err := database.NewOrderRepo(db, l)
	if err != nil {
		l.Fatal(agent, err.Error())
	}
	pairs := pairconf.NewPairRepo(l)
	ss, err := services.WrapServices(&services.Config{
		DB:     db,
		Pairs:  pairs,
		V:      v,
		L:      l,
		PrvKey: prv,
	})

	if err != nil {
		l.Fatal(agent, err.Error())
	}

	ou := app.NewOrderUseCase(pairs, repo, ss.ExchangeRepo, ss.WalletStore,
		ss.PairConfigs, ss.FeeService, l)

	go ou.Run()
	if err := http.NewRouter(ou, pairs, v, l, user, pass).Run(":8081"); err != nil {
		l.Fatal(agent, err.Error())
	}
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

	pemfileinfo, err := privateKeyFile.Stat()
	if err != nil {
		return nil, err
	}
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		return nil, err
	}

	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
