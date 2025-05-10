package game

type Registrator interface {
	AddGame()
	DeleteGame()
	SearchGame()
}

type Registry struct {
	games map[string]Config
}

func NewGameRegistry() *Registry {
	return &Registry{
		games: make(map[string]Config),
	}
}

//TODO Добавление игры, удаление и поиск

func (r *Registry) AddGame() {

}

func (r *Registry) DeleteGame() {

}

func (r *Registry) SearchGame() {

}
