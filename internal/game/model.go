package game

type Config struct {
	ID          string
	Name        string   // Человекочитаемое имя
	SavePaths   []string // Пути к файлам сохранений
	BackupLimit int      // Максимум версий бэкапов
}

type Meta struct {
	ID   string
	Name string
}
