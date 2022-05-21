package application

import (
	"github.com/rs/zerolog"
	"sync"

	"github.com/l1f/blockornot/internal/config"
)

type Context struct {
	Config config.Config
	Logic  Logic
	Logger *zerolog.Logger
	Wg     sync.WaitGroup
}
