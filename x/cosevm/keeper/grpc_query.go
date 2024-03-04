package keeper

import (
	"cosevm/x/cosevm/types"
)

var _ types.QueryServer = Keeper{}
