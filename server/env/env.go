package env

import (
	"fmt"
	"os"
)

type SaftEnv struct {
	APP_NAME        string
	PORT            string
	POSTGRES_URL    string
	POSTGRES_DBNAME string
	USER            string
	PASS            string
	DEBUG           string
}

func GetAppEnv() SaftEnv {
	return SaftEnv{
		APP_NAME: getLocalEnv("APP_NAME"),
		PORT:     getLocalEnv("PORT"),
		// POSTGRES_URL:    getLocalEnv("POSTGRES_URL"),
		// POSTGRES_DBNAME: getLocalEnv("POSTGRES_DBNAME"),
		// USER:  getLocalEnv("USER"),
		// PASS:  getLocalEnv("PASS"),
		DEBUG: getLocalEnv("DEBUG"),
	}
}

func getLocalEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Errorf("Environment var not provided: %s", key))
	}
	return val
}
