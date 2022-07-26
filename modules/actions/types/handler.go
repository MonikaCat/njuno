package types

import "github.com/MonikaCat/njuno/node"

// Context contains the data about a Hasura actions worker execution
type Context struct {
	Node node.Node
}

// NewContext returns a new Context instance
func NewContext(node node.Node) *Context {
	return &Context{
		Node: node,
	}
}

// ActionHandler represents a Hasura action request handler.
// It returns an interface to be returned to the called, or an error if something is wrong
type ActionHandler = func(context *Context, payload *Payload) (interface{}, error)
