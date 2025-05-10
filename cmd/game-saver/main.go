package main

import (
	"Game-saver/config"
	"fmt"
	"os"
)

func main() {
	configPath := "config" + string(os.PathSeparator) + "config.json"
	manager := config.Manager{configPath}
	gameConfigs, err := manager.LoadGameConfigs()
	if err != nil {
		return // TODO Логировать
	}
	fmt.Printf("%#v", gameConfigs)
}
