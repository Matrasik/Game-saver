package game

type Config struct {
	Meta
	SavePaths []string `json:"savePaths"` // Пути к файлам сохранений
}

type Meta struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
