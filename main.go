package main

import (
	"fmt"
	"go_pass/api"
	"go_pass/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось найти env файл")
	}
	conf := config.NewConfig()
	api := api.NewApi(*conf, "https://lol.ru")
	fmt.Println(api)
}
