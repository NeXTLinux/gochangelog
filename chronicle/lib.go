package chronicle

import (
	"github.com/wagoodman/go-partybus"

	"github.com/nextlinux/chronicle/internal/bus"
	"github.com/nextlinux/chronicle/internal/log"
	"github.com/nextlinux/go-logger"
)

// SetLogger sets the logger object used for all logging calls.
func SetLogger(logger logger.Logger) {
	log.Log = logger
}

// SetBus sets the event bus for all published events onto (in-library subscriptions are not allowed).
func SetBus(b *partybus.Bus) {
	bus.SetPublisher(b)
}
