package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	Path string
}

func (l *LocalStorage) Save(gameID string, files map[string][]byte) error {
	err := os.MkdirAll(gameID, 0644)
	if err != nil {
		return fmt.Errorf("невозможно создать папку для сохранения: %s", err)
	}
	for name, data := range files {
		filePath := filepath.Join(gameID, name)
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return fmt.Errorf("ошибка при сохранении файла %s: %s", name, err)
		}
	}
	return nil
}

func (l *LocalStorage) Load(gameID string, version int) (map[string][]byte, error) {
	return nil, nil
}
func (l *LocalStorage) ListVersions(gameID string) ([]int, error) {
	return nil, nil
}
