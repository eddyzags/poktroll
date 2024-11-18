package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/testutil/sample"
	tokenomicstypes "github.com/pokt-network/poktroll/x/tokenomics/types"
)

func TestMsgUpdateParams(t *testing.T) {
	tokenomicsKeeper, srv, ctx := setupMsgServer(t)
	require.NoError(t, tokenomicsKeeper.SetParams(ctx, tokenomicstypes.DefaultParams()))

	validParams := tokenomicstypes.DefaultParams()
	validParams.DaoRewardAddress = sample.AccAddress()

	tests := []struct {
		desc string

		req *tokenomicstypes.MsgUpdateParams

		shouldError    bool
		expectedErrMsg string
	}{
		{
			desc: "invalid: malformed authority address",

			req: &tokenomicstypes.MsgUpdateParams{
				Authority: "invalid",
				Params:    tokenomicstypes.DefaultParams(),
			},

			shouldError:    true,
			expectedErrMsg: "invalid authority",
		},
		{
			desc: "invalid: incorrect authority address",

			req: &tokenomicstypes.MsgUpdateParams{
				Authority: sample.AccAddress(),
				Params:    validParams,
			},

			shouldError:    true,
			expectedErrMsg: "the provided authority address does not match the on-chain governance address",
		},
		{
			desc: "invalid: dao reward address missing",

			req: &tokenomicstypes.MsgUpdateParams{
				Authority: tokenomicsKeeper.GetAuthority(),
				Params: tokenomicstypes.Params{
					// DaoRewardAddress MUST NOT be empty string
					// when MintAllocationDao is greater than 0.
					MintAllocationDao: 0,
					DaoRewardAddress:  "",

					// MintAllocationXXX params MUST sum to 1.
					MintAllocationProposer:    0.1,
					MintAllocationSupplier:    0.1,
					MintAllocationSourceOwner: 0.1,
					MintAllocationApplication: 0.7,
				},
			},

			shouldError:    true,
			expectedErrMsg: "empty address string is not allowed",
		},
		{
			desc: "valid: successful param update",

			req: &tokenomicstypes.MsgUpdateParams{
				Authority: tokenomicsKeeper.GetAuthority(),
				Params: tokenomicstypes.Params{
					MintAllocationDao:         0.1,
					MintAllocationProposer:    0.1,
					MintAllocationSupplier:    0.1,
					MintAllocationSourceOwner: 0.1,
					MintAllocationApplication: 0.6,
					DaoRewardAddress:          sample.AccAddress(),
				},
			},

			shouldError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			_, err := srv.UpdateParams(ctx, test.req)
			if test.shouldError {
				require.Error(t, err)
				require.ErrorContains(t, err, test.expectedErrMsg)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
