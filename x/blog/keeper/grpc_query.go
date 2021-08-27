package keeper

import (
	"github.com/heis-en-berg/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
