package serviceQueue

import (
	"github.com/atom-providers/log"
	"github.com/atom-providers/queue"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		queue.DefaultProvider(),
	}, providers...)
}
