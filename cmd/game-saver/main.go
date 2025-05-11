package main

import (
	"Game-saver/config"
	"fmt"
	"log"
	"os"
)

func main() {
	configPath := "config" + string(os.PathSeparator) + "config.json"
	manager := config.Manager{Path: configPath}
	gameConfigs, err := manager.LoadGameConfigs()
	if err != nil {
		log.Fatalf("Error load configs: %s", err)
	}
	fmt.Printf("%#v", gameConfigs)

	// Интерфейс
	// Получение Id игры.

	// поиск игры внутри мапы и замена файла в сохранении на файл из локального хранилища(если есть)

}
