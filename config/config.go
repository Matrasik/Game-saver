package config

import (
	"encoding/json"
	"os"
)

type GameConfigLoader interface {
	LoadGameConfigs() ([]Config, error)
}

type Manager struct {
	Path string
}

type Config struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	SavePaths   []string `json:"save_paths"`
	BackupLimit int      `json:"-"`
}

func (m *Manager) LoadGameConfigs() ([]Config, error) {
	fileData, err := os.ReadFile(m.Path)
	if err != nil {
		return nil, err
	}
	var gameCfgs []Config
	err = json.Unmarshal(fileData, &gameCfgs)
	if err != nil {
		return nil, err
	}
	return gameCfgs, nil
}
