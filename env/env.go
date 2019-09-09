package env

import (
	"fmt"
	"os"
)

type SaftEnv struct {
	APP_NAME          string
	PORT              string
	POSTGRES_URL      string
	DEFAULT_DATA_PATH string
}

func GetAppEnv() SaftEnv {
	return SaftEnv{
		APP_NAME:          getLocalEnv("APP_NAME"),
		PORT:              getLocalEnv("PORT"),
		POSTGRES_URL:      getLocalEnv("POSTGRES_URL"),
		DEFAULT_DATA_PATH: getLocalEnv("DEFAULT_DATA_PATH"),
	}
}

func getLocalEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Errorf("Environment var not provided: %s", key))
	}
	return val
}
