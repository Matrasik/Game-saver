package saver

import (
	"Game-saver/internal/backup"
	"Game-saver/internal/fileops"
	"Game-saver/internal/game"
	"Game-saver/internal/logger"
	"Game-saver/internal/storage"
)

type GameSaver struct {
	fileOps fileops.FileOperator
	reg     game.Registrator
	storage storage.Storage
	backup  backup.Backup
	logger  logger.Logger
}
