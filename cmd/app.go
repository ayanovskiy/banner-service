package main

import (
	"flag"
	"log"

	"banner-service/internal/app"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config/config.yaml", "Путь до файла настроек")
}

func main() {
	flag.Parse()

	application, err := app.NewApp(configFile)
	if err != nil {
		log.Fatal(err)
	}

	application.Run()
}
