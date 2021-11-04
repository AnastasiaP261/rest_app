package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"rest_app/internal/app/apiserver"
)

var (
	configPath string
)

// задание флагов
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse() // флаги для бинарника

	// создание и парсинг флагов в конфиг
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// создание и запуск сервера
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
