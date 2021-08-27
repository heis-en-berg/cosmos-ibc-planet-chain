package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/heis-en-berg/planet/x/blog/keeper"
	"github.com/heis-en-berg/planet/x/blog/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the post
	for _, elem := range genState.PostList {
		k.SetPost(ctx, *elem)
	}

	// Set post count
	k.SetPostCount(ctx, genState.PostCount)

	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all post
	postList := k.GetAllPost(ctx)
	for _, elem := range postList {
		elem := elem
		genesis.PostList = append(genesis.PostList, &elem)
	}

	// Set the current count
	genesis.PostCount = k.GetPostCount(ctx)

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
