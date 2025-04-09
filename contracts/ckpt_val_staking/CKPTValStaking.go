// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CKPTValStaking

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CKPTValStakingMetaData contains all meta data concerning the CKPTValStaking contract.
var CKPTValStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumStake\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"epochNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"powerSum\",\"type\":\"uint256\"}],\"name\":\"CheckpointSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dispatcherURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"blsPublicKey\",\"type\":\"string\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dispatcherURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"blsPublicKey\",\"type\":\"string\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CKPTRewardScalingFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPATCHER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epochNum\",\"type\":\"uint64\"}],\"name\":\"distributeCheckpointRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"distributedEpochs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"epochToCheckpoint\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"epochNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bitmap\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsMultiSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsAggrPk\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"powerSum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getTopValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"dispatcherURLs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"blsPublicKeys\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dispatcherURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsPublicKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activationTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantDispatcherRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantDistributerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantValidatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"initiateUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_dispatcherURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_blsPublicKey\",\"type\":\"string\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeDispatcherRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeDistributerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeValidatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRateByLockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_CKPTRewardScalingFactor\",\"type\":\"uint256\"}],\"name\":\"setCKPTRewardScalingFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumStake\",\"type\":\"uint256\"}],\"name\":\"setMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardRate\",\"type\":\"uint256\"}],\"name\":\"setRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeLockPeriod\",\"type\":\"uint256\"}],\"name\":\"setStakeLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockPeriod\",\"type\":\"uint256\"}],\"name\":\"setUnstakeLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeActivationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epochNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_bitmap\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsMultiSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsAggrPk\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_powerSum\",\"type\":\"uint256\"}],\"name\":\"submitCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"updateValidatorCursor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_dispatcherURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_blsPublicKey\",\"type\":\"string\"}],\"name\":\"updateValidatorInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dispatcherURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsPublicKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600455600060065562093a80600755620151806008553480156200002957600080fd5b5060405162005fb338038062005fb383398181016040528101906200004f9190620004bd565b33600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000c55760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000bc91906200052a565b60405180910390fd5b620000d681620001de60201b60201c565b5082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600381905550806005819055506200013b6000801b33620002a260201b60201c565b506200016e7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892633620002a260201b60201c565b50620001a17ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c33620002a260201b60201c565b50620001d47f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e33620002a260201b60201c565b5050505062000547565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000620002b68383620003a560201b60201c565b6200039a57600180600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620003366200041060201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600190506200039f565b600090505b92915050565b60006001600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600033905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200044a826200041d565b9050919050565b6200045c816200043d565b81146200046857600080fd5b50565b6000815190506200047c8162000451565b92915050565b6000819050919050565b620004978162000482565b8114620004a357600080fd5b50565b600081519050620004b7816200048c565b92915050565b600080600060608486031215620004d957620004d862000418565b5b6000620004e9868287016200046b565b9350506020620004fc86828701620004a6565b92505060406200050f86828701620004a6565b9150509250925092565b62000524816200043d565b82525050565b600060208201905062000541600083018462000519565b92915050565b615a5c80620005576000396000f3fe608060405234801561001057600080fd5b50600436106102a05760003560e01c80638ae647df11610167578063a8667ce1116100ce578063e69c240311610087578063e69c2403146107ee578063ec5ffac21461080c578063f10e00d61461082a578063f2fde38b14610848578063fa52c7d814610864578063ff9211a71461089b576102a0565b8063a8667ce114610744578063ae5ac92114610760578063be115d2b1461077c578063c49baebe14610798578063d547741f146107b6578063e18d4f8b146107d2576102a0565b80639e447fc6116101205780639e447fc61461066b5780639ef5aee914610687578063a17c89f9146106a5578063a217fddf146106d5578063a5f68a18146106f3578063a694fc3a14610728576102a0565b80638ae647df146105825780638da5cb5b1461059e57806391d14854146105bc57806393a5b1b6146105ec578063997453181461061f5780639dc820a51461064f576102a0565b80633d81380d1161020b5780636d5beac7116101c45780636d5beac7146104d45780637071688a146104f0578063715018a61461050e57806372f702f31461051857806374a28d61146105365780637a76646014610552576102a0565b80633d81380d1461043c578063499ab37e1461045857806351057fad1461047657806363803b23146104925780636b11b1eb1461049c5780636d0ef8d4146104b8576102a0565b80632786879f1161025d5780632786879f1461039257806329514204146103c25780632e3a9690146103de5780632f2ff15d146103fa57806336568abe14610416578063372500ab14610432576102a0565b806301ffc9a7146102a55780630e4be4c7146102d55780631904bb2e146102f1578063233e990314610328578063248a9ca31461034457806326bf1c0314610374575b600080fd5b6102bf60048036038101906102ba9190613cda565b6108b9565b6040516102cc9190613d22565b60405180910390f35b6102ef60048036038101906102ea9190613e4e565b610933565b005b61030b60048036038101906103069190613f9b565b610c28565b60405161031f989796959493929190614067565b60405180910390f35b610342600480360381019061033d91906140f3565b610ee7565b005b61035e60048036038101906103599190614120565b610ef9565b60405161036b919061415c565b60405180910390f35b61037c610f19565b6040516103899190614177565b60405180910390f35b6103ac60048036038101906103a79190613f9b565b610f1f565b6040516103b99190614177565b60405180910390f35b6103dc60048036038101906103d79190613f9b565b610f37565b005b6103f860048036038101906103f391906140f3565b610f6c565b005b610414600480360381019061040f9190614192565b610f7e565b005b610430600480360381019061042b9190614192565b610fa0565b005b61043a61101b565b005b61045660048036038101906104519190614228565b6111e8565b005b610460611520565b60405161046d9190614177565b60405180910390f35b610490600480360381019061048b91906140f3565b611526565b005b61049a6115b0565b005b6104b660048036038101906104b19190613f9b565b61183b565b005b6104d260048036038101906104cd91906140f3565b611870565b005b6104ee60048036038101906104e99190613f9b565b611882565b005b6104f86118b7565b6040516105059190614177565b60405180910390f35b6105166118c4565b005b6105206118d8565b60405161052d9190614308565b60405180910390f35b610550600480360381019061054b9190614323565b6118fe565b005b61056c60048036038101906105679190613f9b565b611c83565b6040516105799190614177565b60405180910390f35b61059c60048036038101906105979190613f9b565b611ccf565b005b6105a6611d04565b6040516105b3919061435f565b60405180910390f35b6105d660048036038101906105d19190614192565b611d2d565b6040516105e39190613d22565b60405180910390f35b610606600480360381019061060191906140f3565b611d98565b6040516106169493929190614602565b60405180910390f35b610639600480360381019061063491906140f3565b612379565b604051610646919061435f565b60405180910390f35b61066960048036038101906106649190614228565b6123b8565b005b610685600480360381019061068091906140f3565b6125a4565b005b61068f6125b6565b60405161069c919061415c565b60405180910390f35b6106bf60048036038101906106ba9190614323565b6125da565b6040516106cc9190613d22565b60405180910390f35b6106dd6125fa565b6040516106ea919061415c565b60405180910390f35b61070d60048036038101906107089190614323565b612601565b60405161071f969594939291906146c7565b60405180910390f35b610742600480360381019061073d91906140f3565b6127e9565b005b61075e60048036038101906107599190613f9b565b612ce2565b005b61077a600480360381019061077591906140f3565b612d17565b005b610796600480360381019061079191906140f3565b613090565b005b6107a06130e5565b6040516107ad919061415c565b60405180910390f35b6107d060048036038101906107cb9190614192565b613109565b005b6107ec60048036038101906107e79190613f9b565b61312b565b005b6107f6613160565b6040516108039190614177565b60405180910390f35b610814613166565b6040516108219190614177565b60405180910390f35b61083261316c565b60405161083f919061415c565b60405180910390f35b610862600480360381019061085d9190613f9b565b613190565b005b61087e60048036038101906108799190613f9b565b613216565b60405161089298979695949392919061473d565b60405180910390f35b6108a361337b565b6040516108b09190614177565b60405180910390f35b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061092c575061092b82613381565b5b9050919050565b61095d7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c33611d2d565b61099c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109939061483b565b60405180910390fd5b6000600a60008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614610a27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1e906148cd565b60405180910390fd5b6040518060c001604052808a67ffffffffffffffff16815260200189815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200182815250600a60008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550602082015181600101556040820151816002019081610b9f9190614b1e565b506060820151816003019081610bb59190614b1e565b506080820151816004019081610bcb9190614b1e565b5060a082015181600501559050508867ffffffffffffffff167f73600511933e8a49dca548b938da9bcce54dfe9820562cd901778c34f41a4f288983604051610c15929190614bf0565b60405180910390a2505050505050505050565b6000806000606080600080600080600b60008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816002015490508160060160009054906101000a900460ff168015610ca2575060008260000154115b8015610ced5750600960008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020544210155b15610d45576000826001015442610d049190614c48565b9050670de0b6b3a7640000836000015460035483610d229190614c7c565b610d2c9190614c7c565b610d369190614ced565b82610d419190614d1e565b9150505b600960008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549250816000015481836003015484600401856005018660060160009054906101000a900460ff16876007015489848054610dbe9061494b565b80601f0160208091040260200160405190810160405280929190818152602001828054610dea9061494b565b8015610e375780601f10610e0c57610100808354040283529160200191610e37565b820191906000526020600020905b815481529060010190602001808311610e1a57829003601f168201915b50505050509450838054610e4a9061494b565b80601f0160208091040260200160405190810160405280929190818152602001828054610e769061494b565b8015610ec35780601f10610e9857610100808354040283529160200191610ec3565b820191906000526020600020905b815481529060010190602001808311610ea657829003601f168201915b50505050509350995099509950995099509950995099505050919395975091939597565b610eef6133eb565b8060058190555050565b600060016000838152602001908152602001600020600101549050919050565b60075481565b60096020528060005260406000206000915090505481565b610f3f6133eb565b610f697f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892682613109565b50565b610f746133eb565b8060088190555050565b610f8782610ef9565b610f9081613472565b610f9a8383613486565b50505050565b610fa8613577565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461100c576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611016828261357f565b505050565b61102433613672565b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201549050600081116110ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a590614d9e565b60405180910390fd5b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020181905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401611153929190614dbe565b6020604051808303816000875af1158015611172573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111969190614e13565b503373ffffffffffffffffffffffffffffffffffffffff167ffc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe826040516111dd9190614177565b60405180910390a250565b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541461126d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161126490614e8c565b60405180910390fd5b6112977f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892633611d2d565b6112d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112cd90614f1e565b60405180910390fd5b60405180610100016040528060008152602001428152602001600081526020016000815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152602001600015158152602001600c80549050815250600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015560608201518160030155608082015181600401908161141f9190614f99565b5060a08201518160050190816114359190614f99565b5060c08201518160060160006101000a81548160ff02191690831515021790555060e08201518160070155905050600c339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f5196ca52b5246aa712ef709c67d65576c4f5ca8aba7d2df5d782cca62cae21d58585858560405161151294939291906150a7565b60405180910390a250505050565b60085481565b6000600c80549050905060008111611573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161156a9061512e565b60405180910390fd5b600060065490506000828410611589578261158b565b835b905082818361159a9190614d1e565b6115a4919061514e565b60068190555050505050565b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015411611635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162c906151cb565b60405180910390fd5b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301544210156116ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116b190615237565b60405180910390fd5b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015490506000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030181905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016117a6929190614dbe565b6020604051808303816000875af11580156117c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e99190614e13565b503373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75826040516118309190614177565b60405180910390a250565b6118436133eb565b61186d7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c82610f7e565b50565b6118786133eb565b8060078190555050565b61188a6133eb565b6118b47f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e82610f7e565b50565b6000600c80549050905090565b6118cc6133eb565b6118d660006137a1565b565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6119287f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e33611d2d565b611967576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195e906152c9565b60405180910390fd5b6000600a60008367ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020905060008160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff16036119f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119ee9061535b565b60405180910390fd5b600d60008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611a6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a63906153ed565b60405180910390fd5b6000611a76613865565b905060008151905060008111611ac1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ab89061547f565b60405180910390fd5b6000670de0b6b3a764000082611ad79190614c7c565b6004548560050154611ae99190614c7c565b611af39190614ced565b905060005b82811015611c3b576000848281518110611b1557611b1461549f565b5b60200260200101519050600082905061020081108015611bc75750611bc6876002018054611b429061494b565b80601f0160208091040260200160405190810160405280929190818152602001828054611b6e9061494b565b8015611bbb5780601f10611b9057610100808354040283529160200191611bbb565b820191906000526020600020905b815481529060010190602001808311611b9e57829003601f168201915b505050505082613b7f565b5b15611c265783600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002016000828254611c1e9190614d1e565b925050819055505b50508080611c33906154ce565b915050611af8565b506001600d60008767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050505050565b6000600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050919050565b611cd76133eb565b611d017f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892682610f7e565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006001600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6060806060806000600c80549050905060008103611ef057600067ffffffffffffffff811115611dcb57611dca6148ed565b5b604051908082528060200260200182016040528015611df95781602001602082028036833780820191505090505b50600067ffffffffffffffff811115611e1557611e146148ed565b5b604051908082528060200260200182016040528015611e435781602001602082028036833780820191505090505b50600067ffffffffffffffff811115611e5f57611e5e6148ed565b5b604051908082528060200260200182016040528015611e9257816020015b6060815260200190600190039081611e7d5790505b50600067ffffffffffffffff811115611eae57611ead6148ed565b5b604051908082528060200260200182016040528015611ee157816020015b6060815260200190600190039081611ecc5790505b50945094509450945050612372565b6000818710611eff5781611f01565b865b90508067ffffffffffffffff811115611f1d57611f1c6148ed565b5b604051908082528060200260200182016040528015611f4b5781602001602082028036833780820191505090505b5095508067ffffffffffffffff811115611f6857611f676148ed565b5b604051908082528060200260200182016040528015611f965781602001602082028036833780820191505090505b5094508067ffffffffffffffff811115611fb357611fb26148ed565b5b604051908082528060200260200182016040528015611fe657816020015b6060815260200190600190039081611fd15790505b5093508067ffffffffffffffff811115612003576120026148ed565b5b60405190808252806020026020018201604052801561203657816020015b60608152602001906001900390816120215790505b509250600080600654905060005b848110801561205257508383105b156123575760008582846120669190614d1e565b612070919061514e565b90506000600c82815481106120885761208761549f565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060060160009054906101000a900460ff16801561211a575060008160000154115b80156121655750600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020544210155b1561234157818c878151811061217e5761217d61549f565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600001548b87815181106121d0576121cf61549f565b5b6020026020010181815250508060040180546121eb9061494b565b80601f01602080910402602001604051908101604052809291908181526020018280546122179061494b565b80156122645780601f1061223957610100808354040283529160200191612264565b820191906000526020600020905b81548152906001019060200180831161224757829003601f168201915b50505050508a878151811061227c5761227b61549f565b5b60200260200101819052508060050180546122969061494b565b80601f01602080910402602001604051908101604052809291908181526020018280546122c29061494b565b801561230f5780601f106122e45761010080835404028352916020019161230f565b820191906000526020600020905b8154815290600101906020018083116122f257829003601f168201915b50505050508987815181106123275761232661549f565b5b6020026020010181905250858061233d906154ce565b9650505b505050808061234f906154ce565b915050612044565b508282101561236d578188528187528186528185525b505050505b9193509193565b600c818154811061238957600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541161243d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161243490615562565b60405180910390fd5b6124677f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892633611d2d565b6124a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161249d90614f1e565b60405180910390fd5b8383600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040191826124f792919061558d565b508181600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600501918261254992919061558d565b503373ffffffffffffffffffffffffffffffffffffffff167f7f35b3dd337a65b9cbaf229ba2c1c6e4aa76460db271df1248aa1a605d44ddf28585858560405161259694939291906150a7565b60405180910390a250505050565b6125ac6133eb565b8060038190555050565b7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c81565b600d6020528060005260406000206000915054906101000a900460ff1681565b6000801b81565b600a6020528060005260406000206000915090508060000160009054906101000a900467ffffffffffffffff16908060010154908060020180546126449061494b565b80601f01602080910402602001604051908101604052809291908181526020018280546126709061494b565b80156126bd5780601f10612692576101008083540402835291602001916126bd565b820191906000526020600020905b8154815290600101906020018083116126a057829003601f168201915b5050505050908060030180546126d29061494b565b80601f01602080910402602001604051908101604052809291908181526020018280546126fe9061494b565b801561274b5780601f106127205761010080835404028352916020019161274b565b820191906000526020600020905b81548152906001019060200180831161272e57829003601f168201915b5050505050908060040180546127609061494b565b80601f016020809104026020016040519081016040528092919081815260200182805461278c9061494b565b80156127d95780601f106127ae576101008083540402835291602001916127d9565b820191906000526020600020905b8154815290600101906020018083116127bc57829003601f168201915b5050505050908060050154905086565b6128137f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892633611d2d565b612852576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161284990614f1e565b60405180910390fd5b600554811015612897576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161288e906156cf565b60405180910390fd5b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541461291c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129139061573b565b60405180910390fd5b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154146129a1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612998906157a7565b60405180910390fd5b6129aa33613672565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401612a09939291906157c7565b6020604051808303816000875af1158015612a28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a4c9190614e13565b5080600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154612a9b9190614d1e565b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160009054906101000a900460ff16612c40576001600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160006101000a81548160ff021916908315150217905550600c80549050600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060070181905550600c339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60085442612c4e9190614d1e565b600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d82604051612cd79190614177565b60405180910390a250565b612cea6133eb565b612d147f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e82613109565b50565b60008111612d5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d519061584a565b60405180910390fd5b80600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541015612ddf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dd6906158b6565b60405180910390fd5b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015414612e64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e5b90615922565b60405180910390fd5b612e6d33613672565b60075442612e7b9190614d1e565b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003018190555080600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154612f0f9190614c48565b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015403612ffb576000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160006101000a81548160ff0219169083151502179055505b3373ffffffffffffffffffffffffffffffffffffffff167f9cee5eacd317086e5050733a5dad16ef5aaec185de792371a9625bfee2a3213b82600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154604051613085929190615942565b60405180910390a250565b6130986133eb565b600081116130db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130d2906159dd565b60405180910390fd5b8060048190555050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892681565b61311282610ef9565b61311b81613472565b613125838361357f565b50505050565b6131336133eb565b61315d7ffbd38eecf51668fdbc772b204dc63dd28c3a3cf32e3025f52a80aa807359f50c82613109565b50565b60045481565b60055481565b7f09630fffc1c31ed9c8dd68f6e39219ed189b07ff9a25e1efc743b828f69d555e81565b6131986133eb565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361320a5760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401613201919061435f565b60405180910390fd5b613213816137a1565b50565b600b6020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040180546132519061494b565b80601f016020809104026020016040519081016040528092919081815260200182805461327d9061494b565b80156132ca5780601f1061329f576101008083540402835291602001916132ca565b820191906000526020600020905b8154815290600101906020018083116132ad57829003601f168201915b5050505050908060050180546132df9061494b565b80601f016020809104026020016040519081016040528092919081815260200182805461330b9061494b565b80156133585780601f1061332d57610100808354040283529160200191613358565b820191906000526020600020905b81548152906001019060200180831161333b57829003601f168201915b5050505050908060060160009054906101000a900460ff16908060070154905088565b60035481565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6133f3613577565b73ffffffffffffffffffffffffffffffffffffffff16613411611d04565b73ffffffffffffffffffffffffffffffffffffffff161461347057613434613577565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401613467919061435f565b60405180910390fd5b565b6134838161347e613577565b613bee565b50565b60006134928383611d2d565b61356c57600180600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550613509613577565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050613571565b600090505b92915050565b600033905090565b600061358b8383611d2d565b156136675760006001600085815260200190815260200160002060000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550613604613577565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a46001905061366c565b600090505b92915050565b6000600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060060160009054906101000a900460ff1680156136d7575060008160000154115b80156137225750600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020544210155b1561379d5760008160010154426137399190614c48565b9050600081111561379b576000670de0b6b3a76400008360000154600354846137629190614c7c565b61376c9190614c7c565b6137769190614ced565b90508083600201546137889190614d1e565b8360020181905550428360010181905550505b505b5050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60606000600c805490509050600061020082116138825781613886565b6102005b9050600081116138cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138c29061512e565b60405180910390fd5b60008167ffffffffffffffff8111156138e7576138e66148ed565b5b6040519080825280602002602001820160405280156139155781602001602082028036833780820191505090505b50905060005b828110156139e057600084826006546139349190614d1e565b61393e919061514e565b9050600c81815481106139545761395361549f565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168383815181106139925761399161549f565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505080806139d8906154ce565b91505061391b565b5060005b6001836139f19190614c48565b811015613b755760005b60018285613a099190614c48565b613a139190614c48565b811015613b6157613a64838281518110613a3057613a2f61549f565b5b602002602001015184600184613a469190614d1e565b81518110613a5757613a5661549f565b5b6020026020010151613c3f565b15613b4e5782600182613a779190614d1e565b81518110613a8857613a8761549f565b5b6020026020010151838281518110613aa357613aa261549f565b5b6020026020010151848381518110613abe57613abd61549f565b5b6020026020010185600185613ad39190614d1e565b81518110613ae457613ae361549f565b5b602002602001018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152508273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050505b8080613b59906154ce565b9150506139fb565b508080613b6d906154ce565b9150506139e4565b5080935050505090565b600080600883613b8f9190614ced565b90506000600884613ba0919061514e565b905084518210613bb557600092505050613be8565b6000816001901b868481518110613bcf57613bce61549f565b5b602001015160f81c60f81b60f81c60ff16161415925050505b92915050565b613bf88282611d2d565b613c3b5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401613c329291906159fd565b60405180910390fd5b5050565b60008173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1611905092915050565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b613cb781613c82565b8114613cc257600080fd5b50565b600081359050613cd481613cae565b92915050565b600060208284031215613cf057613cef613c78565b5b6000613cfe84828501613cc5565b91505092915050565b60008115159050919050565b613d1c81613d07565b82525050565b6000602082019050613d376000830184613d13565b92915050565b600067ffffffffffffffff82169050919050565b613d5a81613d3d565b8114613d6557600080fd5b50565b600081359050613d7781613d51565b92915050565b6000819050919050565b613d9081613d7d565b8114613d9b57600080fd5b50565b600081359050613dad81613d87565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112613dd857613dd7613db3565b5b8235905067ffffffffffffffff811115613df557613df4613db8565b5b602083019150836001820283011115613e1157613e10613dbd565b5b9250929050565b6000819050919050565b613e2b81613e18565b8114613e3657600080fd5b50565b600081359050613e4881613e22565b92915050565b600080600080600080600080600060c08a8c031215613e7057613e6f613c78565b5b6000613e7e8c828d01613d68565b9950506020613e8f8c828d01613d9e565b98505060408a013567ffffffffffffffff811115613eb057613eaf613c7d565b5b613ebc8c828d01613dc2565b975097505060608a013567ffffffffffffffff811115613edf57613ede613c7d565b5b613eeb8c828d01613dc2565b955095505060808a013567ffffffffffffffff811115613f0e57613f0d613c7d565b5b613f1a8c828d01613dc2565b935093505060a0613f2d8c828d01613e39565b9150509295985092959850929598565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613f6882613f3d565b9050919050565b613f7881613f5d565b8114613f8357600080fd5b50565b600081359050613f9581613f6f565b92915050565b600060208284031215613fb157613fb0613c78565b5b6000613fbf84828501613f86565b91505092915050565b613fd181613e18565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015614011578082015181840152602081019050613ff6565b60008484015250505050565b6000601f19601f8301169050919050565b600061403982613fd7565b6140438185613fe2565b9350614053818560208601613ff3565b61405c8161401d565b840191505092915050565b60006101008201905061407d600083018b613fc8565b61408a602083018a613fc8565b6140976040830189613fc8565b81810360608301526140a9818861402e565b905081810360808301526140bd818761402e565b90506140cc60a0830186613d13565b6140d960c0830185613fc8565b6140e660e0830184613fc8565b9998505050505050505050565b60006020828403121561410957614108613c78565b5b600061411784828501613e39565b91505092915050565b60006020828403121561413657614135613c78565b5b600061414484828501613d9e565b91505092915050565b61415681613d7d565b82525050565b6000602082019050614171600083018461414d565b92915050565b600060208201905061418c6000830184613fc8565b92915050565b600080604083850312156141a9576141a8613c78565b5b60006141b785828601613d9e565b92505060206141c885828601613f86565b9150509250929050565b60008083601f8401126141e8576141e7613db3565b5b8235905067ffffffffffffffff81111561420557614204613db8565b5b60208301915083600182028301111561422157614220613dbd565b5b9250929050565b6000806000806040858703121561424257614241613c78565b5b600085013567ffffffffffffffff8111156142605761425f613c7d565b5b61426c878288016141d2565b9450945050602085013567ffffffffffffffff81111561428f5761428e613c7d565b5b61429b878288016141d2565b925092505092959194509250565b6000819050919050565b60006142ce6142c96142c484613f3d565b6142a9565b613f3d565b9050919050565b60006142e0826142b3565b9050919050565b60006142f2826142d5565b9050919050565b614302816142e7565b82525050565b600060208201905061431d60008301846142f9565b92915050565b60006020828403121561433957614338613c78565b5b600061434784828501613d68565b91505092915050565b61435981613f5d565b82525050565b60006020820190506143746000830184614350565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6143af81613f5d565b82525050565b60006143c183836143a6565b60208301905092915050565b6000602082019050919050565b60006143e58261437a565b6143ef8185614385565b93506143fa83614396565b8060005b8381101561442b57815161441288826143b5565b975061441d836143cd565b9250506001810190506143fe565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61446d81613e18565b82525050565b600061447f8383614464565b60208301905092915050565b6000602082019050919050565b60006144a382614438565b6144ad8185614443565b93506144b883614454565b8060005b838110156144e95781516144d08882614473565b97506144db8361448b565b9250506001810190506144bc565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061453e82613fd7565b6145488185614522565b9350614558818560208601613ff3565b6145618161401d565b840191505092915050565b60006145788383614533565b905092915050565b6000602082019050919050565b6000614598826144f6565b6145a28185614501565b9350836020820285016145b485614512565b8060005b858110156145f057848403895281516145d1858261456c565b94506145dc83614580565b925060208a019950506001810190506145b8565b50829750879550505050505092915050565b6000608082019050818103600083015261461c81876143da565b905081810360208301526146308186614498565b90508181036040830152614644818561458d565b90508181036060830152614658818461458d565b905095945050505050565b61466c81613d3d565b82525050565b600081519050919050565b600082825260208201905092915050565b600061469982614672565b6146a3818561467d565b93506146b3818560208601613ff3565b6146bc8161401d565b840191505092915050565b600060c0820190506146dc6000830189614663565b6146e9602083018861414d565b81810360408301526146fb818761468e565b9050818103606083015261470f818661468e565b90508181036080830152614723818561468e565b905061473260a0830184613fc8565b979650505050505050565b600061010082019050614753600083018b613fc8565b614760602083018a613fc8565b61476d6040830189613fc8565b61477a6060830188613fc8565b818103608083015261478c818761402e565b905081810360a08301526147a0818661402e565b90506147af60c0830185613d13565b6147bc60e0830184613fc8565b9998505050505050505050565b7f434b505456616c5374616b696e673a206d75737420686176652064697370617460008201527f6368657220726f6c6520746f207375626d697420636865636b706f696e740000602082015250565b6000614825603e83613fe2565b9150614830826147c9565b604082019050919050565b6000602082019050818103600083015261485481614818565b9050919050565b7f436865636b706f696e7420616c72656164792065786973747320666f7220746860008201527f69732065706f6368000000000000000000000000000000000000000000000000602082015250565b60006148b7602883613fe2565b91506148c28261485b565b604082019050919050565b600060208201905081810360008301526148e6816148aa565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061496357607f821691505b6020821081036149765761497561491c565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026149de7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826149a1565b6149e886836149a1565b95508019841693508086168417925050509392505050565b6000614a1b614a16614a1184613e18565b6142a9565b613e18565b9050919050565b6000819050919050565b614a3583614a00565b614a49614a4182614a22565b8484546149ae565b825550505050565b600090565b614a5e614a51565b614a69818484614a2c565b505050565b5b81811015614a8d57614a82600082614a56565b600181019050614a6f565b5050565b601f821115614ad257614aa38161497c565b614aac84614991565b81016020851015614abb578190505b614acf614ac785614991565b830182614a6e565b50505b505050565b600082821c905092915050565b6000614af560001984600802614ad7565b1980831691505092915050565b6000614b0e8383614ae4565b9150826002028217905092915050565b614b2782614672565b67ffffffffffffffff811115614b4057614b3f6148ed565b5b614b4a825461494b565b614b55828285614a91565b600060209050601f831160018114614b885760008415614b76578287015190505b614b808582614b02565b865550614be8565b601f198416614b968661497c565b60005b82811015614bbe57848901518255600182019150602085019450602081019050614b99565b86831015614bdb5784890151614bd7601f891682614ae4565b8355505b6001600288020188555050505b505050505050565b6000604082019050614c05600083018561414d565b614c126020830184613fc8565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000614c5382613e18565b9150614c5e83613e18565b9250828203905081811115614c7657614c75614c19565b5b92915050565b6000614c8782613e18565b9150614c9283613e18565b9250828202614ca081613e18565b91508282048414831517614cb757614cb6614c19565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000614cf882613e18565b9150614d0383613e18565b925082614d1357614d12614cbe565b5b828204905092915050565b6000614d2982613e18565b9150614d3483613e18565b9250828201905080821115614d4c57614d4b614c19565b5b92915050565b7f4e6f207265776172647320746f20636c61696d00000000000000000000000000600082015250565b6000614d88601383613fe2565b9150614d9382614d52565b602082019050919050565b60006020820190508181036000830152614db781614d7b565b9050919050565b6000604082019050614dd36000830185614350565b614de06020830184613fc8565b9392505050565b614df081613d07565b8114614dfb57600080fd5b50565b600081519050614e0d81614de7565b92915050565b600060208284031215614e2957614e28613c78565b5b6000614e3784828501614dfe565b91505092915050565b7f416c726561647920726567697374657265640000000000000000000000000000600082015250565b6000614e76601283613fe2565b9150614e8182614e40565b602082019050919050565b60006020820190508181036000830152614ea581614e69565b9050919050565b7f434b505456616c5374616b696e673a206d75737420686176652076616c69646160008201527f746f7220726f6c6520746f207374616b65000000000000000000000000000000602082015250565b6000614f08603183613fe2565b9150614f1382614eac565b604082019050919050565b60006020820190508181036000830152614f3781614efb565b9050919050565b60008190508160005260206000209050919050565b601f821115614f9457614f6581614f3e565b614f6e84614991565b81016020851015614f7d578190505b614f91614f8985614991565b830182614a6e565b50505b505050565b614fa282613fd7565b67ffffffffffffffff811115614fbb57614fba6148ed565b5b614fc5825461494b565b614fd0828285614f53565b600060209050601f8311600181146150035760008415614ff1578287015190505b614ffb8582614b02565b865550615063565b601f19841661501186614f3e565b60005b8281101561503957848901518255600182019150602085019450602081019050615014565b868310156150565784890151615052601f891682614ae4565b8355505b6001600288020188555050505b505050505050565b82818337600083830152505050565b60006150868385613fe2565b935061509383858461506b565b61509c8361401d565b840190509392505050565b600060408201905081810360008301526150c281868861507a565b905081810360208301526150d781848661507a565b905095945050505050565b7f4e6f2076616c696461746f727320617661696c61626c65000000000000000000600082015250565b6000615118601783613fe2565b9150615123826150e2565b602082019050919050565b600060208201905081810360008301526151478161510b565b9050919050565b600061515982613e18565b915061516483613e18565b92508261517457615173614cbe565b5b828206905092915050565b7f4e6f20756e7374616b696e6720696e2070726f67726573730000000000000000600082015250565b60006151b5601883613fe2565b91506151c08261517f565b602082019050919050565b600060208201905081810360008301526151e4816151a8565b9050919050565b7f5374696c6c20696e206c6f636b20706572696f64000000000000000000000000600082015250565b6000615221601483613fe2565b915061522c826151eb565b602082019050919050565b6000602082019050818103600083015261525081615214565b9050919050565b7f434b505456616c5374616b696e673a206d75737420686176652064697374726960008201527f627574657220726f6c6520746f20646973747269627574652072657761726473602082015250565b60006152b3604083613fe2565b91506152be82615257565b604082019050919050565b600060208201905081810360008301526152e2816152a6565b9050919050565b7f436865636b706f696e7420646f6573206e6f7420657869737420666f7220746860008201527f69732065706f6368000000000000000000000000000000000000000000000000602082015250565b6000615345602883613fe2565b9150615350826152e9565b604082019050919050565b6000602082019050818103600083015261537481615338565b9050919050565b7f5265776172647320616c726561647920646973747269627574656420666f722060008201527f746869732065706f636800000000000000000000000000000000000000000000602082015250565b60006153d7602a83613fe2565b91506153e28261537b565b604082019050919050565b60006020820190508181036000830152615406816153ca565b9050919050565b7f4e6f2076616c696461746f727320617661696c61626c6520666f72207265776160008201527f7264730000000000000000000000000000000000000000000000000000000000602082015250565b6000615469602383613fe2565b91506154748261540d565b604082019050919050565b600060208201905081810360008301526154988161545c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006154d982613e18565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361550b5761550a614c19565b5b600182019050919050565b7f4e6f7420612076616c696461746f720000000000000000000000000000000000600082015250565b600061554c600f83613fe2565b915061555782615516565b602082019050919050565b6000602082019050818103600083015261557b8161553f565b9050919050565b600082905092915050565b6155978383615582565b67ffffffffffffffff8111156155b0576155af6148ed565b5b6155ba825461494b565b6155c5828285614f53565b6000601f8311600181146155f457600084156155e2578287013590505b6155ec8582614b02565b865550615654565b601f19841661560286614f3e565b60005b8281101561562a57848901358255600182019150602085019450602081019050615605565b868310156156475784890135615643601f891682614ae4565b8355505b6001600288020188555050505b50505050505050565b7f4d757374207374616b6520657861637420626967676572207468616e206d696e60008201527f696d756d20616d6f756e74000000000000000000000000000000000000000000602082015250565b60006156b9602b83613fe2565b91506156c48261565d565b604082019050919050565b600060208201905081810360008301526156e8816156ac565b9050919050565b7f416c7265616479207374616b6564000000000000000000000000000000000000600082015250565b6000615725600e83613fe2565b9150615730826156ef565b602082019050919050565b6000602082019050818103600083015261575481615718565b9050919050565b7f556e7374616b696e6720696e2070726f67726573730000000000000000000000600082015250565b6000615791601583613fe2565b915061579c8261575b565b602082019050919050565b600060208201905081810360008301526157c081615784565b9050919050565b60006060820190506157dc6000830186614350565b6157e96020830185614350565b6157f66040830184613fc8565b949350505050565b7f43616e6e6f7420756e7374616b65203000000000000000000000000000000000600082015250565b6000615834601083613fe2565b915061583f826157fe565b602082019050919050565b6000602082019050818103600083015261586381615827565b9050919050565b7f4e6f7420656e6f756768207374616b6564000000000000000000000000000000600082015250565b60006158a0601183613fe2565b91506158ab8261586a565b602082019050919050565b600060208201905081810360008301526158cf81615893565b9050919050565b7f556e7374616b696e6720616c726561647920696e2070726f6772657373000000600082015250565b600061590c601d83613fe2565b9150615917826158d6565b602082019050919050565b6000602082019050818103600083015261593b816158ff565b9050919050565b60006040820190506159576000830185613fc8565b6159646020830184613fc8565b9392505050565b7f5363616c696e6720666163746f72206d7573742062652067726561746572207460008201527f68616e2030000000000000000000000000000000000000000000000000000000602082015250565b60006159c7602583613fe2565b91506159d28261596b565b604082019050919050565b600060208201905081810360008301526159f6816159ba565b9050919050565b6000604082019050615a126000830185614350565b615a1f602083018461414d565b939250505056fea26469706673582212204ce0021ea32b98a11c51892a661d71919b6d4776eaa25b2e1f04d18836cd25c564736f6c63430008140033",
}

// CKPTValStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use CKPTValStakingMetaData.ABI instead.
var CKPTValStakingABI = CKPTValStakingMetaData.ABI

// CKPTValStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CKPTValStakingMetaData.Bin instead.
var CKPTValStakingBin = CKPTValStakingMetaData.Bin

// DeployCKPTValStaking deploys a new Ethereum contract, binding an instance of CKPTValStaking to it.
func DeployCKPTValStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _stakingToken common.Address, _rewardRate *big.Int, _minimumStake *big.Int) (common.Address, *types.Transaction, *CKPTValStaking, error) {
	parsed, err := CKPTValStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CKPTValStakingBin), backend, _stakingToken, _rewardRate, _minimumStake)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CKPTValStaking{CKPTValStakingCaller: CKPTValStakingCaller{contract: contract}, CKPTValStakingTransactor: CKPTValStakingTransactor{contract: contract}, CKPTValStakingFilterer: CKPTValStakingFilterer{contract: contract}}, nil
}

// CKPTValStaking is an auto generated Go binding around an Ethereum contract.
type CKPTValStaking struct {
	CKPTValStakingCaller     // Read-only binding to the contract
	CKPTValStakingTransactor // Write-only binding to the contract
	CKPTValStakingFilterer   // Log filterer for contract events
}

// CKPTValStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CKPTValStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKPTValStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CKPTValStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKPTValStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CKPTValStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CKPTValStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CKPTValStakingSession struct {
	Contract     *CKPTValStaking   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CKPTValStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CKPTValStakingCallerSession struct {
	Contract *CKPTValStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CKPTValStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CKPTValStakingTransactorSession struct {
	Contract     *CKPTValStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CKPTValStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CKPTValStakingRaw struct {
	Contract *CKPTValStaking // Generic contract binding to access the raw methods on
}

// CKPTValStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CKPTValStakingCallerRaw struct {
	Contract *CKPTValStakingCaller // Generic read-only contract binding to access the raw methods on
}

// CKPTValStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CKPTValStakingTransactorRaw struct {
	Contract *CKPTValStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCKPTValStaking creates a new instance of CKPTValStaking, bound to a specific deployed contract.
func NewCKPTValStaking(address common.Address, backend bind.ContractBackend) (*CKPTValStaking, error) {
	contract, err := bindCKPTValStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CKPTValStaking{CKPTValStakingCaller: CKPTValStakingCaller{contract: contract}, CKPTValStakingTransactor: CKPTValStakingTransactor{contract: contract}, CKPTValStakingFilterer: CKPTValStakingFilterer{contract: contract}}, nil
}

// NewCKPTValStakingCaller creates a new read-only instance of CKPTValStaking, bound to a specific deployed contract.
func NewCKPTValStakingCaller(address common.Address, caller bind.ContractCaller) (*CKPTValStakingCaller, error) {
	contract, err := bindCKPTValStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingCaller{contract: contract}, nil
}

// NewCKPTValStakingTransactor creates a new write-only instance of CKPTValStaking, bound to a specific deployed contract.
func NewCKPTValStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*CKPTValStakingTransactor, error) {
	contract, err := bindCKPTValStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingTransactor{contract: contract}, nil
}

// NewCKPTValStakingFilterer creates a new log filterer instance of CKPTValStaking, bound to a specific deployed contract.
func NewCKPTValStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*CKPTValStakingFilterer, error) {
	contract, err := bindCKPTValStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingFilterer{contract: contract}, nil
}

// bindCKPTValStaking binds a generic wrapper to an already deployed contract.
func bindCKPTValStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CKPTValStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CKPTValStaking *CKPTValStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CKPTValStaking.Contract.CKPTValStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CKPTValStaking *CKPTValStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.CKPTValStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CKPTValStaking *CKPTValStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.CKPTValStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CKPTValStaking *CKPTValStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CKPTValStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CKPTValStaking *CKPTValStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CKPTValStaking *CKPTValStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.contract.Transact(opts, method, params...)
}

// CKPTRewardScalingFactor is a free data retrieval call binding the contract method 0xe69c2403.
//
// Solidity: function CKPTRewardScalingFactor() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) CKPTRewardScalingFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "CKPTRewardScalingFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CKPTRewardScalingFactor is a free data retrieval call binding the contract method 0xe69c2403.
//
// Solidity: function CKPTRewardScalingFactor() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) CKPTRewardScalingFactor() (*big.Int, error) {
	return _CKPTValStaking.Contract.CKPTRewardScalingFactor(&_CKPTValStaking.CallOpts)
}

// CKPTRewardScalingFactor is a free data retrieval call binding the contract method 0xe69c2403.
//
// Solidity: function CKPTRewardScalingFactor() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) CKPTRewardScalingFactor() (*big.Int, error) {
	return _CKPTValStaking.Contract.CKPTRewardScalingFactor(&_CKPTValStaking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DEFAULTADMINROLE(&_CKPTValStaking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DEFAULTADMINROLE(&_CKPTValStaking.CallOpts)
}

// DISPATCHERROLE is a free data retrieval call binding the contract method 0x9ef5aee9.
//
// Solidity: function DISPATCHER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) DISPATCHERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "DISPATCHER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISPATCHERROLE is a free data retrieval call binding the contract method 0x9ef5aee9.
//
// Solidity: function DISPATCHER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) DISPATCHERROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DISPATCHERROLE(&_CKPTValStaking.CallOpts)
}

// DISPATCHERROLE is a free data retrieval call binding the contract method 0x9ef5aee9.
//
// Solidity: function DISPATCHER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) DISPATCHERROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DISPATCHERROLE(&_CKPTValStaking.CallOpts)
}

// DISTRIBUTERROLE is a free data retrieval call binding the contract method 0xf10e00d6.
//
// Solidity: function DISTRIBUTER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) DISTRIBUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "DISTRIBUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISTRIBUTERROLE is a free data retrieval call binding the contract method 0xf10e00d6.
//
// Solidity: function DISTRIBUTER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) DISTRIBUTERROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DISTRIBUTERROLE(&_CKPTValStaking.CallOpts)
}

// DISTRIBUTERROLE is a free data retrieval call binding the contract method 0xf10e00d6.
//
// Solidity: function DISTRIBUTER_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) DISTRIBUTERROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.DISTRIBUTERROLE(&_CKPTValStaking.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) VALIDATORROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.VALIDATORROLE(&_CKPTValStaking.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _CKPTValStaking.Contract.VALIDATORROLE(&_CKPTValStaking.CallOpts)
}

// DistributedEpochs is a free data retrieval call binding the contract method 0xa17c89f9.
//
// Solidity: function distributedEpochs(uint64 ) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCaller) DistributedEpochs(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "distributedEpochs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DistributedEpochs is a free data retrieval call binding the contract method 0xa17c89f9.
//
// Solidity: function distributedEpochs(uint64 ) view returns(bool)
func (_CKPTValStaking *CKPTValStakingSession) DistributedEpochs(arg0 uint64) (bool, error) {
	return _CKPTValStaking.Contract.DistributedEpochs(&_CKPTValStaking.CallOpts, arg0)
}

// DistributedEpochs is a free data retrieval call binding the contract method 0xa17c89f9.
//
// Solidity: function distributedEpochs(uint64 ) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCallerSession) DistributedEpochs(arg0 uint64) (bool, error) {
	return _CKPTValStaking.Contract.DistributedEpochs(&_CKPTValStaking.CallOpts, arg0)
}

// EpochToCheckpoint is a free data retrieval call binding the contract method 0xa5f68a18.
//
// Solidity: function epochToCheckpoint(uint64 ) view returns(uint64 epochNum, bytes32 blockHash, bytes bitmap, bytes blsMultiSig, bytes blsAggrPk, uint256 powerSum)
func (_CKPTValStaking *CKPTValStakingCaller) EpochToCheckpoint(opts *bind.CallOpts, arg0 uint64) (struct {
	EpochNum    uint64
	BlockHash   [32]byte
	Bitmap      []byte
	BlsMultiSig []byte
	BlsAggrPk   []byte
	PowerSum    *big.Int
}, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "epochToCheckpoint", arg0)

	outstruct := new(struct {
		EpochNum    uint64
		BlockHash   [32]byte
		Bitmap      []byte
		BlsMultiSig []byte
		BlsAggrPk   []byte
		PowerSum    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EpochNum = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Bitmap = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlsMultiSig = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.BlsAggrPk = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.PowerSum = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochToCheckpoint is a free data retrieval call binding the contract method 0xa5f68a18.
//
// Solidity: function epochToCheckpoint(uint64 ) view returns(uint64 epochNum, bytes32 blockHash, bytes bitmap, bytes blsMultiSig, bytes blsAggrPk, uint256 powerSum)
func (_CKPTValStaking *CKPTValStakingSession) EpochToCheckpoint(arg0 uint64) (struct {
	EpochNum    uint64
	BlockHash   [32]byte
	Bitmap      []byte
	BlsMultiSig []byte
	BlsAggrPk   []byte
	PowerSum    *big.Int
}, error) {
	return _CKPTValStaking.Contract.EpochToCheckpoint(&_CKPTValStaking.CallOpts, arg0)
}

// EpochToCheckpoint is a free data retrieval call binding the contract method 0xa5f68a18.
//
// Solidity: function epochToCheckpoint(uint64 ) view returns(uint64 epochNum, bytes32 blockHash, bytes bitmap, bytes blsMultiSig, bytes blsAggrPk, uint256 powerSum)
func (_CKPTValStaking *CKPTValStakingCallerSession) EpochToCheckpoint(arg0 uint64) (struct {
	EpochNum    uint64
	BlockHash   [32]byte
	Bitmap      []byte
	BlsMultiSig []byte
	BlsAggrPk   []byte
	PowerSum    *big.Int
}, error) {
	return _CKPTValStaking.Contract.EpochToCheckpoint(&_CKPTValStaking.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CKPTValStaking.Contract.GetRoleAdmin(&_CKPTValStaking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CKPTValStaking.Contract.GetRoleAdmin(&_CKPTValStaking.CallOpts, role)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) GetStake(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getStake", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) GetStake(_validator common.Address) (*big.Int, error) {
	return _CKPTValStaking.Contract.GetStake(&_CKPTValStaking.CallOpts, _validator)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetStake(_validator common.Address) (*big.Int, error) {
	return _CKPTValStaking.Contract.GetStake(&_CKPTValStaking.CallOpts, _validator)
}

// GetTopValidators is a free data retrieval call binding the contract method 0x93a5b1b6.
//
// Solidity: function getTopValidators(uint256 _count) view returns(address[] addresses, uint256[] stakes, string[] dispatcherURLs, string[] blsPublicKeys)
func (_CKPTValStaking *CKPTValStakingCaller) GetTopValidators(opts *bind.CallOpts, _count *big.Int) (struct {
	Addresses      []common.Address
	Stakes         []*big.Int
	DispatcherURLs []string
	BlsPublicKeys  []string
}, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getTopValidators", _count)

	outstruct := new(struct {
		Addresses      []common.Address
		Stakes         []*big.Int
		DispatcherURLs []string
		BlsPublicKeys  []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.DispatcherURLs = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.BlsPublicKeys = *abi.ConvertType(out[3], new([]string)).(*[]string)

	return *outstruct, err

}

// GetTopValidators is a free data retrieval call binding the contract method 0x93a5b1b6.
//
// Solidity: function getTopValidators(uint256 _count) view returns(address[] addresses, uint256[] stakes, string[] dispatcherURLs, string[] blsPublicKeys)
func (_CKPTValStaking *CKPTValStakingSession) GetTopValidators(_count *big.Int) (struct {
	Addresses      []common.Address
	Stakes         []*big.Int
	DispatcherURLs []string
	BlsPublicKeys  []string
}, error) {
	return _CKPTValStaking.Contract.GetTopValidators(&_CKPTValStaking.CallOpts, _count)
}

// GetTopValidators is a free data retrieval call binding the contract method 0x93a5b1b6.
//
// Solidity: function getTopValidators(uint256 _count) view returns(address[] addresses, uint256[] stakes, string[] dispatcherURLs, string[] blsPublicKeys)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetTopValidators(_count *big.Int) (struct {
	Addresses      []common.Address
	Stakes         []*big.Int
	DispatcherURLs []string
	BlsPublicKeys  []string
}, error) {
	return _CKPTValStaking.Contract.GetTopValidators(&_CKPTValStaking.CallOpts, _count)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns(uint256 stakedAmount, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index, uint256 activationTime)
func (_CKPTValStaking *CKPTValStakingCaller) GetValidator(opts *bind.CallOpts, _validator common.Address) (struct {
	StakedAmount   *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
	ActivationTime *big.Int
}, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getValidator", _validator)

	outstruct := new(struct {
		StakedAmount   *big.Int
		PendingRewards *big.Int
		UnstakeTime    *big.Int
		DispatcherURL  string
		BlsPublicKey   string
		IsActive       bool
		Index          *big.Int
		ActivationTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingRewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnstakeTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DispatcherURL = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.BlsPublicKey = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Index = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ActivationTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns(uint256 stakedAmount, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index, uint256 activationTime)
func (_CKPTValStaking *CKPTValStakingSession) GetValidator(_validator common.Address) (struct {
	StakedAmount   *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
	ActivationTime *big.Int
}, error) {
	return _CKPTValStaking.Contract.GetValidator(&_CKPTValStaking.CallOpts, _validator)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns(uint256 stakedAmount, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index, uint256 activationTime)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetValidator(_validator common.Address) (struct {
	StakedAmount   *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
	ActivationTime *big.Int
}, error) {
	return _CKPTValStaking.Contract.GetValidator(&_CKPTValStaking.CallOpts, _validator)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) GetValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "getValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) GetValidatorCount() (*big.Int, error) {
	return _CKPTValStaking.Contract.GetValidatorCount(&_CKPTValStaking.CallOpts)
}

// GetValidatorCount is a free data retrieval call binding the contract method 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) GetValidatorCount() (*big.Int, error) {
	return _CKPTValStaking.Contract.GetValidatorCount(&_CKPTValStaking.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CKPTValStaking *CKPTValStakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CKPTValStaking.Contract.HasRole(&_CKPTValStaking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CKPTValStaking.Contract.HasRole(&_CKPTValStaking.CallOpts, role, account)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) MinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "minimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) MinimumStake() (*big.Int, error) {
	return _CKPTValStaking.Contract.MinimumStake(&_CKPTValStaking.CallOpts)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) MinimumStake() (*big.Int, error) {
	return _CKPTValStaking.Contract.MinimumStake(&_CKPTValStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CKPTValStaking *CKPTValStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CKPTValStaking *CKPTValStakingSession) Owner() (common.Address, error) {
	return _CKPTValStaking.Contract.Owner(&_CKPTValStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CKPTValStaking *CKPTValStakingCallerSession) Owner() (common.Address, error) {
	return _CKPTValStaking.Contract.Owner(&_CKPTValStaking.CallOpts)
}

// RewardRateByLockTime is a free data retrieval call binding the contract method 0xff9211a7.
//
// Solidity: function rewardRateByLockTime() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) RewardRateByLockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "rewardRateByLockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRateByLockTime is a free data retrieval call binding the contract method 0xff9211a7.
//
// Solidity: function rewardRateByLockTime() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) RewardRateByLockTime() (*big.Int, error) {
	return _CKPTValStaking.Contract.RewardRateByLockTime(&_CKPTValStaking.CallOpts)
}

// RewardRateByLockTime is a free data retrieval call binding the contract method 0xff9211a7.
//
// Solidity: function rewardRateByLockTime() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) RewardRateByLockTime() (*big.Int, error) {
	return _CKPTValStaking.Contract.RewardRateByLockTime(&_CKPTValStaking.CallOpts)
}

// StakeActivationTime is a free data retrieval call binding the contract method 0x2786879f.
//
// Solidity: function stakeActivationTime(address ) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) StakeActivationTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "stakeActivationTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeActivationTime is a free data retrieval call binding the contract method 0x2786879f.
//
// Solidity: function stakeActivationTime(address ) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) StakeActivationTime(arg0 common.Address) (*big.Int, error) {
	return _CKPTValStaking.Contract.StakeActivationTime(&_CKPTValStaking.CallOpts, arg0)
}

// StakeActivationTime is a free data retrieval call binding the contract method 0x2786879f.
//
// Solidity: function stakeActivationTime(address ) view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) StakeActivationTime(arg0 common.Address) (*big.Int, error) {
	return _CKPTValStaking.Contract.StakeActivationTime(&_CKPTValStaking.CallOpts, arg0)
}

// StakeLockPeriod is a free data retrieval call binding the contract method 0x499ab37e.
//
// Solidity: function stakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) StakeLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "stakeLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLockPeriod is a free data retrieval call binding the contract method 0x499ab37e.
//
// Solidity: function stakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) StakeLockPeriod() (*big.Int, error) {
	return _CKPTValStaking.Contract.StakeLockPeriod(&_CKPTValStaking.CallOpts)
}

// StakeLockPeriod is a free data retrieval call binding the contract method 0x499ab37e.
//
// Solidity: function stakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) StakeLockPeriod() (*big.Int, error) {
	return _CKPTValStaking.Contract.StakeLockPeriod(&_CKPTValStaking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CKPTValStaking *CKPTValStakingCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CKPTValStaking *CKPTValStakingSession) StakingToken() (common.Address, error) {
	return _CKPTValStaking.Contract.StakingToken(&_CKPTValStaking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CKPTValStaking *CKPTValStakingCallerSession) StakingToken() (common.Address, error) {
	return _CKPTValStaking.Contract.StakingToken(&_CKPTValStaking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CKPTValStaking *CKPTValStakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CKPTValStaking.Contract.SupportsInterface(&_CKPTValStaking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CKPTValStaking *CKPTValStakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CKPTValStaking.Contract.SupportsInterface(&_CKPTValStaking.CallOpts, interfaceId)
}

// UnstakeLockPeriod is a free data retrieval call binding the contract method 0x26bf1c03.
//
// Solidity: function unstakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCaller) UnstakeLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "unstakeLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeLockPeriod is a free data retrieval call binding the contract method 0x26bf1c03.
//
// Solidity: function unstakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingSession) UnstakeLockPeriod() (*big.Int, error) {
	return _CKPTValStaking.Contract.UnstakeLockPeriod(&_CKPTValStaking.CallOpts)
}

// UnstakeLockPeriod is a free data retrieval call binding the contract method 0x26bf1c03.
//
// Solidity: function unstakeLockPeriod() view returns(uint256)
func (_CKPTValStaking *CKPTValStakingCallerSession) UnstakeLockPeriod() (*big.Int, error) {
	return _CKPTValStaking.Contract.UnstakeLockPeriod(&_CKPTValStaking.CallOpts)
}

// ValidatorAddresses is a free data retrieval call binding the contract method 0x99745318.
//
// Solidity: function validatorAddresses(uint256 ) view returns(address)
func (_CKPTValStaking *CKPTValStakingCaller) ValidatorAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "validatorAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorAddresses is a free data retrieval call binding the contract method 0x99745318.
//
// Solidity: function validatorAddresses(uint256 ) view returns(address)
func (_CKPTValStaking *CKPTValStakingSession) ValidatorAddresses(arg0 *big.Int) (common.Address, error) {
	return _CKPTValStaking.Contract.ValidatorAddresses(&_CKPTValStaking.CallOpts, arg0)
}

// ValidatorAddresses is a free data retrieval call binding the contract method 0x99745318.
//
// Solidity: function validatorAddresses(uint256 ) view returns(address)
func (_CKPTValStaking *CKPTValStakingCallerSession) ValidatorAddresses(arg0 *big.Int) (common.Address, error) {
	return _CKPTValStaking.Contract.ValidatorAddresses(&_CKPTValStaking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stakedAmount, uint256 lastRewardTime, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index)
func (_CKPTValStaking *CKPTValStakingCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	StakedAmount   *big.Int
	LastRewardTime *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
}, error) {
	var out []interface{}
	err := _CKPTValStaking.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		StakedAmount   *big.Int
		LastRewardTime *big.Int
		PendingRewards *big.Int
		UnstakeTime    *big.Int
		DispatcherURL  string
		BlsPublicKey   string
		IsActive       bool
		Index          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DispatcherURL = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.BlsPublicKey = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.IsActive = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Index = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stakedAmount, uint256 lastRewardTime, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index)
func (_CKPTValStaking *CKPTValStakingSession) Validators(arg0 common.Address) (struct {
	StakedAmount   *big.Int
	LastRewardTime *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
}, error) {
	return _CKPTValStaking.Contract.Validators(&_CKPTValStaking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stakedAmount, uint256 lastRewardTime, uint256 pendingRewards, uint256 unstakeTime, string dispatcherURL, string blsPublicKey, bool isActive, uint256 index)
func (_CKPTValStaking *CKPTValStakingCallerSession) Validators(arg0 common.Address) (struct {
	StakedAmount   *big.Int
	LastRewardTime *big.Int
	PendingRewards *big.Int
	UnstakeTime    *big.Int
	DispatcherURL  string
	BlsPublicKey   string
	IsActive       bool
	Index          *big.Int
}, error) {
	return _CKPTValStaking.Contract.Validators(&_CKPTValStaking.CallOpts, arg0)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CKPTValStaking *CKPTValStakingTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CKPTValStaking *CKPTValStakingSession) ClaimRewards() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.ClaimRewards(&_CKPTValStaking.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.ClaimRewards(&_CKPTValStaking.TransactOpts)
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x63803b23.
//
// Solidity: function completeUnstake() returns()
func (_CKPTValStaking *CKPTValStakingTransactor) CompleteUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "completeUnstake")
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x63803b23.
//
// Solidity: function completeUnstake() returns()
func (_CKPTValStaking *CKPTValStakingSession) CompleteUnstake() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.CompleteUnstake(&_CKPTValStaking.TransactOpts)
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x63803b23.
//
// Solidity: function completeUnstake() returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) CompleteUnstake() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.CompleteUnstake(&_CKPTValStaking.TransactOpts)
}

// DistributeCheckpointRewards is a paid mutator transaction binding the contract method 0x74a28d61.
//
// Solidity: function distributeCheckpointRewards(uint64 _epochNum) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) DistributeCheckpointRewards(opts *bind.TransactOpts, _epochNum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "distributeCheckpointRewards", _epochNum)
}

// DistributeCheckpointRewards is a paid mutator transaction binding the contract method 0x74a28d61.
//
// Solidity: function distributeCheckpointRewards(uint64 _epochNum) returns()
func (_CKPTValStaking *CKPTValStakingSession) DistributeCheckpointRewards(_epochNum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.DistributeCheckpointRewards(&_CKPTValStaking.TransactOpts, _epochNum)
}

// DistributeCheckpointRewards is a paid mutator transaction binding the contract method 0x74a28d61.
//
// Solidity: function distributeCheckpointRewards(uint64 _epochNum) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) DistributeCheckpointRewards(_epochNum uint64) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.DistributeCheckpointRewards(&_CKPTValStaking.TransactOpts, _epochNum)
}

// GrantDispatcherRole is a paid mutator transaction binding the contract method 0x6b11b1eb.
//
// Solidity: function grantDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) GrantDispatcherRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "grantDispatcherRole", account)
}

// GrantDispatcherRole is a paid mutator transaction binding the contract method 0x6b11b1eb.
//
// Solidity: function grantDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) GrantDispatcherRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantDispatcherRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantDispatcherRole is a paid mutator transaction binding the contract method 0x6b11b1eb.
//
// Solidity: function grantDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) GrantDispatcherRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantDispatcherRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantDistributerRole is a paid mutator transaction binding the contract method 0x6d5beac7.
//
// Solidity: function grantDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) GrantDistributerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "grantDistributerRole", account)
}

// GrantDistributerRole is a paid mutator transaction binding the contract method 0x6d5beac7.
//
// Solidity: function grantDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) GrantDistributerRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantDistributerRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantDistributerRole is a paid mutator transaction binding the contract method 0x6d5beac7.
//
// Solidity: function grantDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) GrantDistributerRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantDistributerRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantRole(&_CKPTValStaking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantRole(&_CKPTValStaking.TransactOpts, role, account)
}

// GrantValidatorRole is a paid mutator transaction binding the contract method 0x8ae647df.
//
// Solidity: function grantValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) GrantValidatorRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "grantValidatorRole", account)
}

// GrantValidatorRole is a paid mutator transaction binding the contract method 0x8ae647df.
//
// Solidity: function grantValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) GrantValidatorRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantValidatorRole(&_CKPTValStaking.TransactOpts, account)
}

// GrantValidatorRole is a paid mutator transaction binding the contract method 0x8ae647df.
//
// Solidity: function grantValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) GrantValidatorRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.GrantValidatorRole(&_CKPTValStaking.TransactOpts, account)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) InitiateUnstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "initiateUnstake", _amount)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingSession) InitiateUnstake(_amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.InitiateUnstake(&_CKPTValStaking.TransactOpts, _amount)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) InitiateUnstake(_amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.InitiateUnstake(&_CKPTValStaking.TransactOpts, _amount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3d81380d.
//
// Solidity: function registerValidator(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RegisterValidator(opts *bind.TransactOpts, _dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "registerValidator", _dispatcherURL, _blsPublicKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3d81380d.
//
// Solidity: function registerValidator(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingSession) RegisterValidator(_dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RegisterValidator(&_CKPTValStaking.TransactOpts, _dispatcherURL, _blsPublicKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x3d81380d.
//
// Solidity: function registerValidator(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RegisterValidator(_dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RegisterValidator(&_CKPTValStaking.TransactOpts, _dispatcherURL, _blsPublicKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CKPTValStaking *CKPTValStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RenounceOwnership(&_CKPTValStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RenounceOwnership(&_CKPTValStaking.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CKPTValStaking *CKPTValStakingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RenounceRole(&_CKPTValStaking.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RenounceRole(&_CKPTValStaking.TransactOpts, role, callerConfirmation)
}

// RevokeDispatcherRole is a paid mutator transaction binding the contract method 0xe18d4f8b.
//
// Solidity: function revokeDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RevokeDispatcherRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "revokeDispatcherRole", account)
}

// RevokeDispatcherRole is a paid mutator transaction binding the contract method 0xe18d4f8b.
//
// Solidity: function revokeDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) RevokeDispatcherRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeDispatcherRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeDispatcherRole is a paid mutator transaction binding the contract method 0xe18d4f8b.
//
// Solidity: function revokeDispatcherRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RevokeDispatcherRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeDispatcherRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeDistributerRole is a paid mutator transaction binding the contract method 0xa8667ce1.
//
// Solidity: function revokeDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RevokeDistributerRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "revokeDistributerRole", account)
}

// RevokeDistributerRole is a paid mutator transaction binding the contract method 0xa8667ce1.
//
// Solidity: function revokeDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) RevokeDistributerRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeDistributerRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeDistributerRole is a paid mutator transaction binding the contract method 0xa8667ce1.
//
// Solidity: function revokeDistributerRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RevokeDistributerRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeDistributerRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeRole(&_CKPTValStaking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeRole(&_CKPTValStaking.TransactOpts, role, account)
}

// RevokeValidatorRole is a paid mutator transaction binding the contract method 0x29514204.
//
// Solidity: function revokeValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) RevokeValidatorRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "revokeValidatorRole", account)
}

// RevokeValidatorRole is a paid mutator transaction binding the contract method 0x29514204.
//
// Solidity: function revokeValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingSession) RevokeValidatorRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeValidatorRole(&_CKPTValStaking.TransactOpts, account)
}

// RevokeValidatorRole is a paid mutator transaction binding the contract method 0x29514204.
//
// Solidity: function revokeValidatorRole(address account) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) RevokeValidatorRole(account common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.RevokeValidatorRole(&_CKPTValStaking.TransactOpts, account)
}

// SetCKPTRewardScalingFactor is a paid mutator transaction binding the contract method 0xbe115d2b.
//
// Solidity: function setCKPTRewardScalingFactor(uint256 _CKPTRewardScalingFactor) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetCKPTRewardScalingFactor(opts *bind.TransactOpts, _CKPTRewardScalingFactor *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setCKPTRewardScalingFactor", _CKPTRewardScalingFactor)
}

// SetCKPTRewardScalingFactor is a paid mutator transaction binding the contract method 0xbe115d2b.
//
// Solidity: function setCKPTRewardScalingFactor(uint256 _CKPTRewardScalingFactor) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetCKPTRewardScalingFactor(_CKPTRewardScalingFactor *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetCKPTRewardScalingFactor(&_CKPTValStaking.TransactOpts, _CKPTRewardScalingFactor)
}

// SetCKPTRewardScalingFactor is a paid mutator transaction binding the contract method 0xbe115d2b.
//
// Solidity: function setCKPTRewardScalingFactor(uint256 _CKPTRewardScalingFactor) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetCKPTRewardScalingFactor(_CKPTRewardScalingFactor *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetCKPTRewardScalingFactor(&_CKPTValStaking.TransactOpts, _CKPTRewardScalingFactor)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 _minimumStake) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetMinimumStake(opts *bind.TransactOpts, _minimumStake *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setMinimumStake", _minimumStake)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 _minimumStake) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetMinimumStake(_minimumStake *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetMinimumStake(&_CKPTValStaking.TransactOpts, _minimumStake)
}

// SetMinimumStake is a paid mutator transaction binding the contract method 0x233e9903.
//
// Solidity: function setMinimumStake(uint256 _minimumStake) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetMinimumStake(_minimumStake *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetMinimumStake(&_CKPTValStaking.TransactOpts, _minimumStake)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetRewardRate(opts *bind.TransactOpts, _rewardRate *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setRewardRate", _rewardRate)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetRewardRate(_rewardRate *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetRewardRate(&_CKPTValStaking.TransactOpts, _rewardRate)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetRewardRate(_rewardRate *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetRewardRate(&_CKPTValStaking.TransactOpts, _rewardRate)
}

// SetStakeLockPeriod is a paid mutator transaction binding the contract method 0x2e3a9690.
//
// Solidity: function setStakeLockPeriod(uint256 _stakeLockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetStakeLockPeriod(opts *bind.TransactOpts, _stakeLockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setStakeLockPeriod", _stakeLockPeriod)
}

// SetStakeLockPeriod is a paid mutator transaction binding the contract method 0x2e3a9690.
//
// Solidity: function setStakeLockPeriod(uint256 _stakeLockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetStakeLockPeriod(_stakeLockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetStakeLockPeriod(&_CKPTValStaking.TransactOpts, _stakeLockPeriod)
}

// SetStakeLockPeriod is a paid mutator transaction binding the contract method 0x2e3a9690.
//
// Solidity: function setStakeLockPeriod(uint256 _stakeLockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetStakeLockPeriod(_stakeLockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetStakeLockPeriod(&_CKPTValStaking.TransactOpts, _stakeLockPeriod)
}

// SetUnstakeLockPeriod is a paid mutator transaction binding the contract method 0x6d0ef8d4.
//
// Solidity: function setUnstakeLockPeriod(uint256 _lockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SetUnstakeLockPeriod(opts *bind.TransactOpts, _lockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "setUnstakeLockPeriod", _lockPeriod)
}

// SetUnstakeLockPeriod is a paid mutator transaction binding the contract method 0x6d0ef8d4.
//
// Solidity: function setUnstakeLockPeriod(uint256 _lockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingSession) SetUnstakeLockPeriod(_lockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetUnstakeLockPeriod(&_CKPTValStaking.TransactOpts, _lockPeriod)
}

// SetUnstakeLockPeriod is a paid mutator transaction binding the contract method 0x6d0ef8d4.
//
// Solidity: function setUnstakeLockPeriod(uint256 _lockPeriod) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SetUnstakeLockPeriod(_lockPeriod *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SetUnstakeLockPeriod(&_CKPTValStaking.TransactOpts, _lockPeriod)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.Stake(&_CKPTValStaking.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.Stake(&_CKPTValStaking.TransactOpts, _amount)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x0e4be4c7.
//
// Solidity: function submitCheckpoint(uint64 _epochNum, bytes32 _blockHash, bytes _bitmap, bytes _blsMultiSig, bytes _blsAggrPk, uint256 _powerSum) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) SubmitCheckpoint(opts *bind.TransactOpts, _epochNum uint64, _blockHash [32]byte, _bitmap []byte, _blsMultiSig []byte, _blsAggrPk []byte, _powerSum *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "submitCheckpoint", _epochNum, _blockHash, _bitmap, _blsMultiSig, _blsAggrPk, _powerSum)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x0e4be4c7.
//
// Solidity: function submitCheckpoint(uint64 _epochNum, bytes32 _blockHash, bytes _bitmap, bytes _blsMultiSig, bytes _blsAggrPk, uint256 _powerSum) returns()
func (_CKPTValStaking *CKPTValStakingSession) SubmitCheckpoint(_epochNum uint64, _blockHash [32]byte, _bitmap []byte, _blsMultiSig []byte, _blsAggrPk []byte, _powerSum *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SubmitCheckpoint(&_CKPTValStaking.TransactOpts, _epochNum, _blockHash, _bitmap, _blsMultiSig, _blsAggrPk, _powerSum)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x0e4be4c7.
//
// Solidity: function submitCheckpoint(uint64 _epochNum, bytes32 _blockHash, bytes _bitmap, bytes _blsMultiSig, bytes _blsAggrPk, uint256 _powerSum) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) SubmitCheckpoint(_epochNum uint64, _blockHash [32]byte, _bitmap []byte, _blsMultiSig []byte, _blsAggrPk []byte, _powerSum *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.SubmitCheckpoint(&_CKPTValStaking.TransactOpts, _epochNum, _blockHash, _bitmap, _blsMultiSig, _blsAggrPk, _powerSum)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CKPTValStaking *CKPTValStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.TransferOwnership(&_CKPTValStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.TransferOwnership(&_CKPTValStaking.TransactOpts, newOwner)
}

// UpdateValidatorCursor is a paid mutator transaction binding the contract method 0x51057fad.
//
// Solidity: function updateValidatorCursor(uint256 _count) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) UpdateValidatorCursor(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "updateValidatorCursor", _count)
}

// UpdateValidatorCursor is a paid mutator transaction binding the contract method 0x51057fad.
//
// Solidity: function updateValidatorCursor(uint256 _count) returns()
func (_CKPTValStaking *CKPTValStakingSession) UpdateValidatorCursor(_count *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.UpdateValidatorCursor(&_CKPTValStaking.TransactOpts, _count)
}

// UpdateValidatorCursor is a paid mutator transaction binding the contract method 0x51057fad.
//
// Solidity: function updateValidatorCursor(uint256 _count) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) UpdateValidatorCursor(_count *big.Int) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.UpdateValidatorCursor(&_CKPTValStaking.TransactOpts, _count)
}

// UpdateValidatorInfo is a paid mutator transaction binding the contract method 0x9dc820a5.
//
// Solidity: function updateValidatorInfo(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingTransactor) UpdateValidatorInfo(opts *bind.TransactOpts, _dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.contract.Transact(opts, "updateValidatorInfo", _dispatcherURL, _blsPublicKey)
}

// UpdateValidatorInfo is a paid mutator transaction binding the contract method 0x9dc820a5.
//
// Solidity: function updateValidatorInfo(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingSession) UpdateValidatorInfo(_dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.UpdateValidatorInfo(&_CKPTValStaking.TransactOpts, _dispatcherURL, _blsPublicKey)
}

// UpdateValidatorInfo is a paid mutator transaction binding the contract method 0x9dc820a5.
//
// Solidity: function updateValidatorInfo(string _dispatcherURL, string _blsPublicKey) returns()
func (_CKPTValStaking *CKPTValStakingTransactorSession) UpdateValidatorInfo(_dispatcherURL string, _blsPublicKey string) (*types.Transaction, error) {
	return _CKPTValStaking.Contract.UpdateValidatorInfo(&_CKPTValStaking.TransactOpts, _dispatcherURL, _blsPublicKey)
}

// CKPTValStakingCheckpointSubmittedIterator is returned from FilterCheckpointSubmitted and is used to iterate over the raw logs and unpacked data for CheckpointSubmitted events raised by the CKPTValStaking contract.
type CKPTValStakingCheckpointSubmittedIterator struct {
	Event *CKPTValStakingCheckpointSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingCheckpointSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingCheckpointSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingCheckpointSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingCheckpointSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingCheckpointSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingCheckpointSubmitted represents a CheckpointSubmitted event raised by the CKPTValStaking contract.
type CKPTValStakingCheckpointSubmitted struct {
	EpochNum  uint64
	BlockHash [32]byte
	PowerSum  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCheckpointSubmitted is a free log retrieval operation binding the contract event 0x73600511933e8a49dca548b938da9bcce54dfe9820562cd901778c34f41a4f28.
//
// Solidity: event CheckpointSubmitted(uint64 indexed epochNum, bytes32 blockHash, uint256 powerSum)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterCheckpointSubmitted(opts *bind.FilterOpts, epochNum []uint64) (*CKPTValStakingCheckpointSubmittedIterator, error) {

	var epochNumRule []interface{}
	for _, epochNumItem := range epochNum {
		epochNumRule = append(epochNumRule, epochNumItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "CheckpointSubmitted", epochNumRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingCheckpointSubmittedIterator{contract: _CKPTValStaking.contract, event: "CheckpointSubmitted", logs: logs, sub: sub}, nil
}

// WatchCheckpointSubmitted is a free log subscription operation binding the contract event 0x73600511933e8a49dca548b938da9bcce54dfe9820562cd901778c34f41a4f28.
//
// Solidity: event CheckpointSubmitted(uint64 indexed epochNum, bytes32 blockHash, uint256 powerSum)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchCheckpointSubmitted(opts *bind.WatchOpts, sink chan<- *CKPTValStakingCheckpointSubmitted, epochNum []uint64) (event.Subscription, error) {

	var epochNumRule []interface{}
	for _, epochNumItem := range epochNum {
		epochNumRule = append(epochNumRule, epochNumItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "CheckpointSubmitted", epochNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingCheckpointSubmitted)
				if err := _CKPTValStaking.contract.UnpackLog(event, "CheckpointSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointSubmitted is a log parse operation binding the contract event 0x73600511933e8a49dca548b938da9bcce54dfe9820562cd901778c34f41a4f28.
//
// Solidity: event CheckpointSubmitted(uint64 indexed epochNum, bytes32 blockHash, uint256 powerSum)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseCheckpointSubmitted(log types.Log) (*CKPTValStakingCheckpointSubmitted, error) {
	event := new(CKPTValStakingCheckpointSubmitted)
	if err := _CKPTValStaking.contract.UnpackLog(event, "CheckpointSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CKPTValStaking contract.
type CKPTValStakingOwnershipTransferredIterator struct {
	Event *CKPTValStakingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingOwnershipTransferred represents a OwnershipTransferred event raised by the CKPTValStaking contract.
type CKPTValStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CKPTValStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingOwnershipTransferredIterator{contract: _CKPTValStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CKPTValStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingOwnershipTransferred)
				if err := _CKPTValStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseOwnershipTransferred(log types.Log) (*CKPTValStakingOwnershipTransferred, error) {
	event := new(CKPTValStakingOwnershipTransferred)
	if err := _CKPTValStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the CKPTValStaking contract.
type CKPTValStakingRewardsClaimedIterator struct {
	Event *CKPTValStakingRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingRewardsClaimed represents a RewardsClaimed event raised by the CKPTValStaking contract.
type CKPTValStakingRewardsClaimed struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingRewardsClaimedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "RewardsClaimed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingRewardsClaimedIterator{contract: _CKPTValStaking.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *CKPTValStakingRewardsClaimed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "RewardsClaimed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingRewardsClaimed)
				if err := _CKPTValStaking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsClaimed is a log parse operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseRewardsClaimed(log types.Log) (*CKPTValStakingRewardsClaimed, error) {
	event := new(CKPTValStakingRewardsClaimed)
	if err := _CKPTValStaking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CKPTValStaking contract.
type CKPTValStakingRoleAdminChangedIterator struct {
	Event *CKPTValStakingRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingRoleAdminChanged represents a RoleAdminChanged event raised by the CKPTValStaking contract.
type CKPTValStakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CKPTValStakingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingRoleAdminChangedIterator{contract: _CKPTValStaking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CKPTValStakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingRoleAdminChanged)
				if err := _CKPTValStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseRoleAdminChanged(log types.Log) (*CKPTValStakingRoleAdminChanged, error) {
	event := new(CKPTValStakingRoleAdminChanged)
	if err := _CKPTValStaking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CKPTValStaking contract.
type CKPTValStakingRoleGrantedIterator struct {
	Event *CKPTValStakingRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingRoleGranted represents a RoleGranted event raised by the CKPTValStaking contract.
type CKPTValStakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CKPTValStakingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingRoleGrantedIterator{contract: _CKPTValStaking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CKPTValStakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingRoleGranted)
				if err := _CKPTValStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseRoleGranted(log types.Log) (*CKPTValStakingRoleGranted, error) {
	event := new(CKPTValStakingRoleGranted)
	if err := _CKPTValStaking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CKPTValStaking contract.
type CKPTValStakingRoleRevokedIterator struct {
	Event *CKPTValStakingRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingRoleRevoked represents a RoleRevoked event raised by the CKPTValStaking contract.
type CKPTValStakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CKPTValStakingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingRoleRevokedIterator{contract: _CKPTValStaking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CKPTValStakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingRoleRevoked)
				if err := _CKPTValStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseRoleRevoked(log types.Log) (*CKPTValStakingRoleRevoked, error) {
	event := new(CKPTValStakingRoleRevoked)
	if err := _CKPTValStaking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the CKPTValStaking contract.
type CKPTValStakingStakedIterator struct {
	Event *CKPTValStakingStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingStaked represents a Staked event raised by the CKPTValStaking contract.
type CKPTValStakingStaked struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterStaked(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingStakedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "Staked", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingStakedIterator{contract: _CKPTValStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *CKPTValStakingStaked, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "Staked", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingStaked)
				if err := _CKPTValStaking.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseStaked(log types.Log) (*CKPTValStakingStaked, error) {
	event := new(CKPTValStakingStaked)
	if err := _CKPTValStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingUnstakeInitiatedIterator is returned from FilterUnstakeInitiated and is used to iterate over the raw logs and unpacked data for UnstakeInitiated events raised by the CKPTValStaking contract.
type CKPTValStakingUnstakeInitiatedIterator struct {
	Event *CKPTValStakingUnstakeInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingUnstakeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingUnstakeInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingUnstakeInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingUnstakeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingUnstakeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingUnstakeInitiated represents a UnstakeInitiated event raised by the CKPTValStaking contract.
type CKPTValStakingUnstakeInitiated struct {
	Validator  common.Address
	Amount     *big.Int
	UnlockTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnstakeInitiated is a free log retrieval operation binding the contract event 0x9cee5eacd317086e5050733a5dad16ef5aaec185de792371a9625bfee2a3213b.
//
// Solidity: event UnstakeInitiated(address indexed validator, uint256 amount, uint256 unlockTime)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterUnstakeInitiated(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingUnstakeInitiatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "UnstakeInitiated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingUnstakeInitiatedIterator{contract: _CKPTValStaking.contract, event: "UnstakeInitiated", logs: logs, sub: sub}, nil
}

// WatchUnstakeInitiated is a free log subscription operation binding the contract event 0x9cee5eacd317086e5050733a5dad16ef5aaec185de792371a9625bfee2a3213b.
//
// Solidity: event UnstakeInitiated(address indexed validator, uint256 amount, uint256 unlockTime)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchUnstakeInitiated(opts *bind.WatchOpts, sink chan<- *CKPTValStakingUnstakeInitiated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "UnstakeInitiated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingUnstakeInitiated)
				if err := _CKPTValStaking.contract.UnpackLog(event, "UnstakeInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstakeInitiated is a log parse operation binding the contract event 0x9cee5eacd317086e5050733a5dad16ef5aaec185de792371a9625bfee2a3213b.
//
// Solidity: event UnstakeInitiated(address indexed validator, uint256 amount, uint256 unlockTime)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseUnstakeInitiated(log types.Log) (*CKPTValStakingUnstakeInitiated, error) {
	event := new(CKPTValStakingUnstakeInitiated)
	if err := _CKPTValStaking.contract.UnpackLog(event, "UnstakeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the CKPTValStaking contract.
type CKPTValStakingUnstakedIterator struct {
	Event *CKPTValStakingUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingUnstaked represents a Unstaked event raised by the CKPTValStaking contract.
type CKPTValStakingUnstaked struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterUnstaked(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingUnstakedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "Unstaked", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingUnstakedIterator{contract: _CKPTValStaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *CKPTValStakingUnstaked, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "Unstaked", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingUnstaked)
				if err := _CKPTValStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseUnstaked(log types.Log) (*CKPTValStakingUnstaked, error) {
	event := new(CKPTValStakingUnstaked)
	if err := _CKPTValStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the CKPTValStaking contract.
type CKPTValStakingValidatorRegisteredIterator struct {
	Event *CKPTValStakingValidatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingValidatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingValidatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingValidatorRegistered represents a ValidatorRegistered event raised by the CKPTValStaking contract.
type CKPTValStakingValidatorRegistered struct {
	Validator     common.Address
	DispatcherURL string
	BlsPublicKey  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x5196ca52b5246aa712ef709c67d65576c4f5ca8aba7d2df5d782cca62cae21d5.
//
// Solidity: event ValidatorRegistered(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingValidatorRegisteredIterator{contract: _CKPTValStaking.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x5196ca52b5246aa712ef709c67d65576c4f5ca8aba7d2df5d782cca62cae21d5.
//
// Solidity: event ValidatorRegistered(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *CKPTValStakingValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingValidatorRegistered)
				if err := _CKPTValStaking.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistered is a log parse operation binding the contract event 0x5196ca52b5246aa712ef709c67d65576c4f5ca8aba7d2df5d782cca62cae21d5.
//
// Solidity: event ValidatorRegistered(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseValidatorRegistered(log types.Log) (*CKPTValStakingValidatorRegistered, error) {
	event := new(CKPTValStakingValidatorRegistered)
	if err := _CKPTValStaking.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CKPTValStakingValidatorUpdatedIterator is returned from FilterValidatorUpdated and is used to iterate over the raw logs and unpacked data for ValidatorUpdated events raised by the CKPTValStaking contract.
type CKPTValStakingValidatorUpdatedIterator struct {
	Event *CKPTValStakingValidatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CKPTValStakingValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CKPTValStakingValidatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CKPTValStakingValidatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CKPTValStakingValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CKPTValStakingValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CKPTValStakingValidatorUpdated represents a ValidatorUpdated event raised by the CKPTValStaking contract.
type CKPTValStakingValidatorUpdated struct {
	Validator     common.Address
	DispatcherURL string
	BlsPublicKey  string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorUpdated is a free log retrieval operation binding the contract event 0x7f35b3dd337a65b9cbaf229ba2c1c6e4aa76460db271df1248aa1a605d44ddf2.
//
// Solidity: event ValidatorUpdated(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) FilterValidatorUpdated(opts *bind.FilterOpts, validator []common.Address) (*CKPTValStakingValidatorUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.FilterLogs(opts, "ValidatorUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CKPTValStakingValidatorUpdatedIterator{contract: _CKPTValStaking.contract, event: "ValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorUpdated is a free log subscription operation binding the contract event 0x7f35b3dd337a65b9cbaf229ba2c1c6e4aa76460db271df1248aa1a605d44ddf2.
//
// Solidity: event ValidatorUpdated(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) WatchValidatorUpdated(opts *bind.WatchOpts, sink chan<- *CKPTValStakingValidatorUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CKPTValStaking.contract.WatchLogs(opts, "ValidatorUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CKPTValStakingValidatorUpdated)
				if err := _CKPTValStaking.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorUpdated is a log parse operation binding the contract event 0x7f35b3dd337a65b9cbaf229ba2c1c6e4aa76460db271df1248aa1a605d44ddf2.
//
// Solidity: event ValidatorUpdated(address indexed validator, string dispatcherURL, string blsPublicKey)
func (_CKPTValStaking *CKPTValStakingFilterer) ParseValidatorUpdated(log types.Log) (*CKPTValStakingValidatorUpdated, error) {
	event := new(CKPTValStakingValidatorUpdated)
	if err := _CKPTValStaking.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
