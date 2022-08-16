package messages

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/MonikaCat/njuno/database"
	"github.com/MonikaCat/njuno/modules"
)

var _ modules.Module = &Module{}

// Module represents the module allowing to store messages inside database
type Module struct {
	parser MessageAddressesParser
	cdc    codec.Marshaler
	db     database.Database
}

func NewModule(parser MessageAddressesParser, cdc codec.Marshaler, db database.Database) *Module {
	return &Module{
		parser: parser,
		cdc:    cdc,
		db:     db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "messages"
}
