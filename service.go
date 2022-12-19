package main

import (
	"bufio"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"exchange-provider/internal/app"
	"exchange-provider/internal/delivery/http"
	"exchange-provider/internal/delivery/services"
	"exchange-provider/internal/delivery/storage"
	"exchange-provider/pkg/logger"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
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

	rc := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR")})

	if err := rc.Ping(context.Background()).Err(); err != nil {
		l.Fatal(agent, fmt.Sprintf("failed to ping redis: %s", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), "exchange-provider")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal(agent, "failed to connect to mysql")
	}
	l.Debug(agent, "connected to mysql")
	s := storage.NewStorage(db, rc, l)

	ss, err := services.WrapServices(&services.Config{
		DB:     db,
		V:      v,
		L:      l,
		RC:     rc,
		PrvKey: prv,
	})

	if err != nil {
		l.Fatal(agent, err.Error())
	}

	wg := &sync.WaitGroup{}

	ou := app.NewOrderUseCase(rc, s.Repo, ss.ExchangeRepo, ss.WalletStore,
		ss.PairConfigs, s.Oc, ss.FeeService, l)
	wg.Add(1)
	go ou.Run(wg)

	if err := http.NewRouter(ou, v, rc, l, user, pass).Run(":8000"); err != nil {
		l.Fatal(agent, err.Error())
	}

	wg.Wait()

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

	rc := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR")})

	if err := rc.Ping(context.Background()).Err(); err != nil {
		l.Fatal(agent, fmt.Sprintf("failed to ping redis: %s", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "root_L0d92Jlf0HfNmV01", "localhost:3306", "exchange-provider")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal(agent, "failed to connect to mysql")
	}
	db.Logger.LogMode(glogger.Silent)
	l.Debug(agent, "connected to mysql")
	s := storage.NewStorage(db, rc, l)

	ss, err := services.WrapServices(&services.Config{
		DB:     db,
		V:      v,
		L:      l,
		RC:     rc,
		PrvKey: prv,
	})

	if err != nil {
		l.Fatal(agent, err.Error())
	}

	wg := &sync.WaitGroup{}

	ou := app.NewOrderUseCase(rc, s.Repo, ss.ExchangeRepo, ss.WalletStore,
		ss.PairConfigs, s.Oc, ss.FeeService, l)
	wg.Add(1)
	go ou.Run(wg)

	if err := http.NewRouter(ou, v, rc, l, user, pass).Run(":8081"); err != nil {
		l.Fatal(agent, err.Error())
	}

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
