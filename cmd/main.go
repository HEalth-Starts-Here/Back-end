package main

import (
	"hesh/internal/app"
	"hesh/internal/pkg/utils/config"
	"hesh/internal/pkg/utils/log"
)

func main() {
	err := config.DevConfigStore.FromJson()
	if err != nil {
		log.Error(err)
	}

	err = config.ProdConfigStore.FromJson()
	if err != nil {
		log.Error(err)
	}

	app.RunServer()
}
