package keeper

import (
	"github.com/heis-en-berg/planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}
