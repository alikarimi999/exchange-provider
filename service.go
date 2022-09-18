package main

import (
	"bufio"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"order_service/internal/app"
	"order_service/internal/delivery/http"
	"order_service/internal/delivery/services"
	"order_service/internal/delivery/storage"
	"order_service/pkg/logger"
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

	ou := app.NewOrderUseCase(rc, s.Repo, ss.ExRepo, ss.PairConf, s.Oc, ss.Fee, l)
	wg.Add(1)
	go ou.Run(wg)

	http.NewRouter(ou, v, rc, l).Run(":8000")

	wg.Wait()

}

func test() {

	agent := "main"

	l := logger.NewLogger("./service.log", true)

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
	db.Logger.LogMode(glogger.Silent)
	l.Debug(agent, fmt.Sprintf("connected to %s", dsn))
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

	ou := app.NewOrderUseCase(rc, s.Repo, ss.ExRepo, ss.PairConf, s.Oc, ss.Fee, l)
	wg.Add(1)
	go ou.Run(wg)

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
