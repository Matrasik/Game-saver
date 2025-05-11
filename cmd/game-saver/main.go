package main

import (
	"Game-saver/config"
	"fmt"
	"log"
	"os"
)

func main() {
	configPath := "config" + string(os.PathSeparator) + "config.json"
	manager := config.Manager{configPath}
	gameConfigs, err := manager.LoadGameConfigs()
	if err != nil {
		log.Fatalf("Error load configs: %s", err)
	}
	fmt.Printf("%#v", gameConfigs)

}
