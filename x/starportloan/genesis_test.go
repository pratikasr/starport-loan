package starportloan_test

import (
	"testing"

	keepertest "github.com/pratikasr/starport-loan/testutil/keeper"
	"github.com/pratikasr/starport-loan/x/starportloan"
	"github.com/pratikasr/starport-loan/x/starportloan/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		LoanList: []types.Loan{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LoanCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StarportloanKeeper(t)
	starportloan.InitGenesis(ctx, *k, genesisState)
	got := starportloan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.LoanList, len(genesisState.LoanList))
	require.Subset(t, genesisState.LoanList, got.LoanList)
	require.Equal(t, genesisState.LoanCount, got.LoanCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
