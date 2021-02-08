package graph

import (
	"github.com/ip75/media.server.go/graph/generated"
	"github.com/ip75/media.server.go/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is resolver :)
type Resolver struct {
	root  generated.ResolverRoot
	input *model.MediaInput
}
