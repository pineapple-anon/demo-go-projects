package processor

import (
	"sync"
	"eg.com/order-processor/internal/models"
)

type OrderProcessor struct {
	seen sync.Map
}

func New() *OrderProcessor {
	return &OrderProcessor{}
}

func (p *OrderProcessor) Process(order models.Order) bool {
	// Check if the order has already been seen
	if _, ok := p.seen.Load(order.ID); ok {
		return false // Order already processed
	}

	// Mark the order as seen
	p.seen.Store(order.ID, struct{}{})

	// Process the order (e.g., store it in a database, send it to another service, etc.)
	// For demonstration purposes, we'll just print it
	println("Processing order:", order.ID)

	return true
}