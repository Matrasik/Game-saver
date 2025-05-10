package storage

type LocalStorage struct {
	Path string
}

func (l *LocalStorage) Save(gameID string, files map[string][]byte) error {
	return nil
}

func (l *LocalStorage) Load(gameID string, version int) (map[string][]byte, error) {
	return nil, nil
}
func (l *LocalStorage) ListVersions(gameID string) ([]int, error) {
	return nil, nil
}
