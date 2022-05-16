package bosagora

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

func RewardCommonsBudget (blockNumber *big.Int, commonsBudget common.Address, lastCommonsBudgetRewardBlock big.Int, commonsBudgetReward big.Int, statedb *state.StateDB) {
	if blockNumber.Cmp(&lastCommonsBudgetRewardBlock) >= 0 {
		return
	}
	statedb.AddBalance(commonsBudget, &commonsBudgetReward)
}
