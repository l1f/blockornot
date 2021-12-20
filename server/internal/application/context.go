package application

import (
	"log"
	"sync"

	"github.com/l1f/blockornot/internal/config"
)

type Context struct {
	Config config.Config
	Logic  Logic
	Logger *log.Logger
	Wg     sync.WaitGroup
}
