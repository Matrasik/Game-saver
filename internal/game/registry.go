package game

import (
	"fmt"
	"sync"
)

type Registrator interface {
	AddGame(game Config) error
	DeleteGame(id string) error
	SearchGame(id string) (*Config, error)
}

type Registry struct {
	games sync.Map
}

func NewGameRegistry() *Registry {
	return &Registry{
		games: sync.Map{},
	}
}

// TODO Добавление игры, удаление и поиск
// Добавляет игру себе в мапу для последующего обрашения по ней
func (r *Registry) AddGame(game Config) error {
	r.games.Store(game.ID, game)
	return nil
}

// Удаляем игру из мапы
func (r *Registry) DeleteGame(id string) error {
	_, loaded := r.games.Load(id)
	if !loaded {
		return fmt.Errorf("игра с ID '%s' не найдена", id)
	}
	r.games.Delete(id)
	return nil
}

// Ищем игру по мапе
func (r *Registry) SearchGame(id string) (*Config, error) {
	value, ok := r.games.Load(id)
	if !ok {
		return nil, fmt.Errorf("игра с ID '%s' не найдена", id)
	}
	game, ok := value.(Config)
	if !ok {
		return nil, fmt.Errorf("неверный тип данных для игры с ID '%s'", id)
	}
	return &game, nil
}
