package config

import (
	"Game-saver/internal/game"
	"encoding/json"
	"os"
)

type GameConfigLoader interface {
	LoadGameConfigs() ([]game.Config, error)
	AddGame(config game.Config) error
}

type Manager struct {
	Path string
}

func (m *Manager) LoadGameConfigs() ([]game.Config, error) {
	fileData, err := os.ReadFile(m.Path)
	if err != nil {
		return nil, err
	}
	var gameCfgs []game.Config
	err = json.Unmarshal(fileData, &gameCfgs)
	if err != nil {
		return nil, err
	}
	return gameCfgs, nil
}

// Добавляет игру в конфиг файл
func (m *Manager) AddGame(game game.Config) error {
	games, err := m.LoadGameConfigs()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	games = append(games, game)

	data, err := json.MarshalIndent(games, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(m.Path, data, 0644)
}
