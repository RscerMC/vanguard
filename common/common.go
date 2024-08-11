package common

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/RscerMC/vanguard/config"
	"github.com/redis/go-redis/v9"
)

var (
	Version = "0.0.1"

	Redis *redis.Client
	CTX   = context.Background()

	FlagVersion bool
)

func init() {
	flag.BoolVar(&FlagVersion, "version", false, "print version and exit")
}

func Init() error {
	if !flag.Parsed() {
		flag.Parse()
	}

	if FlagVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	fmt.Printf("Running Vanguard version %s\n", Version)

	err := connectRedis()
	if err != nil {
		return err
	}
	return nil
}

func connectRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
		Password: "",
		DB:       config.DB,
	})
	_, err := Redis.Ping(CTX).Result()
	if err != nil {
		return fmt.Errorf("failed to ping Redis: %v", err)
	} else {
		fmt.Println("Connected to Redis")
	}
	return nil
}
