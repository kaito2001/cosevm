package cosevm_test

import (
	"testing"

	keepertest "cosevm/testutil/keeper"
	"cosevm/testutil/nullify"
	"cosevm/x/cosevm"
	"cosevm/x/cosevm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosevmKeeper(t)
	cosevm.InitGenesis(ctx, *k, genesisState)
	got := cosevm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
