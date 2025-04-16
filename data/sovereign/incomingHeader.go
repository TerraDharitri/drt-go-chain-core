//go:generate protoc -I=. -I=$GOPATH/src/github.com/TerraDharitri/drt-go-chain-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/TerraDharitri/protobuf/protobuf  --gogoslick_out=. incomingHeader.proto

package sovereign

import "github.com/TerraDharitri/drt-go-chain-core/data"

// GetIncomingEventHandlers returns the incoming events as an array of event handlers
func (ih *IncomingHeader) GetIncomingEventHandlers() []data.EventHandler {
	if ih == nil {
		return nil
	}

	events := ih.GetIncomingEvents()
	logHandlers := make([]data.EventHandler, len(events))

	for i := range events {
		logHandlers[i] = events[i]
	}

	return logHandlers
}

// GetHeaderHandler returns the incoming headerV2 as a header handler
func (ih *IncomingHeader) GetHeaderHandler() data.HeaderHandler {
	if ih == nil {
		return nil
	}

	return ih.GetHeader()
}

// IsInterfaceNil checks if the underlying pointer is nil
func (ih *IncomingHeader) IsInterfaceNil() bool {
	return ih == nil
}
