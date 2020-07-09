// Package ordering defines the interface of the ordering service. The
// high-level purpose of this service is to order the transactions from the
// pool.
//
// Depending on the implementation, the service can be composed of multiple
// sub-components. For instance, an ordering service using CoSiPBFT will need to
// elect a leader every round but one running PoW will only do an ordering
// locally and creates a block with the proof of work.
package ordering

import "context"

// Block is the result of an ordering round.
type Block interface {
	GetIndex() uint64
}

type Event struct {
	Index uint64
}

// Service is the interface of an ordering service. It provides the primitives
// to order transactions from a pool.
type Service interface {
	// Listen opens the endpoints of the service so that other participants can
	// contact the node.
	Listen() error

	GetBlock(index uint64) (Block, error)

	Watch(ctx context.Context) <-chan Event
}
