package serviceQueue

import (
	"github.com/atom-providers/queue"
	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"
)

type Queue struct {
	dig.In

	Service  *queue.Queue
	Handlers []queue.QueueHandler `group:"queue_handler"`
}

func Serve() error {
	defer container.Close()

	return container.Container.Invoke(func(queue Queue) error {
		queue.Service.Handle(queue.Handlers...)
		return queue.Service.Serve()
	})
}
