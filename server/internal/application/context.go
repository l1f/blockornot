package application

import (
	"github.com/l1f/blockornot/internal/config"
	"log"
	"sync"
)

type Context struct {
	Config config.Config
	Logic  Logic
	Logger *log.Logger
	Wg     sync.WaitGroup
}
