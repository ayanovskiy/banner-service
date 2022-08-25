package main

import (
	"flag"

	"banner-service/internal/app"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config/config.yaml", "Путь до файла настроек")
}

func main() {
	flag.Parse()

	application, _ := app.NewApp(configFile)

	application.Run()
}
