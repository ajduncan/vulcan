package vulcan

import (
	"log"
	"os"
)

func Getenv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func GetWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
	   log.Fatal(err)
	}
	return wd
}
