package storage

type Storage interface {
	Save(gameID string, files map[string][]byte) error
	Load(gameID string, version int) (map[string][]byte, error)
	ListVersions(gameID string) ([]int, error)
}
