package application

import (
	"sync"

	"github.com/l1f/blockornot/internal/config"
	"github.com/l1f/blockornot/logger"
)

type Context struct {
	Config config.Config
	Logic  Logic
	Logger *logger.Logger
	Wg     sync.WaitGroup
}
