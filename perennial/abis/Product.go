// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package perennial

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
)

// Accumulator is an auto generated low-level Go binding around an user-defined struct.
type Accumulator struct {
	Maker *big.Int
	Taker *big.Int
}

// IOracleProviderOracleVersion is an auto generated low-level Go binding around an user-defined struct.
type IOracleProviderOracleVersion struct {
	Version   *big.Int
	Timestamp *big.Int
	Price     *big.Int
}

// IProductProductInfo is an auto generated low-level Go binding around an user-defined struct.
type IProductProductInfo struct {
	Name             string
	Symbol           string
	PayoffDefinition PayoffDefinition
	Oracle           common.Address
	Maintenance      *big.Int
	FundingFee       *big.Int
	MakerFee         *big.Int
	TakerFee         *big.Int
	PositionFee      *big.Int
	MakerLimit       *big.Int
	UtilizationCurve JumpRateUtilizationCurve
}

// JumpRateUtilizationCurve is an auto generated low-level Go binding around an user-defined struct.
type JumpRateUtilizationCurve struct {
	MinRate           *big.Int
	MaxRate           *big.Int
	TargetRate        *big.Int
	TargetUtilization *big.Int
}

// PayoffDefinition is an auto generated low-level Go binding around an user-defined struct.
type PayoffDefinition struct {
	PayoffType      uint8
	PayoffDirection uint8
	Data            [30]byte
}

// PendingFeeUpdates is an auto generated low-level Go binding around an user-defined struct.
type PendingFeeUpdates struct {
	MakerFeeUpdated    bool
	PendingMakerFee    uint64
	TakerFeeUpdated    bool
	PendingTakerFee    uint64
	PositionFeeUpdated bool
	PendingPositionFee uint64
}

// Position is an auto generated low-level Go binding around an user-defined struct.
type Position struct {
	Maker *big.Int
	Taker *big.Int
}

// PrePosition is an auto generated low-level Go binding around an user-defined struct.
type PrePosition struct {
	OracleVersion *big.Int
	OpenPosition  Position
	ClosePosition Position
}

// ProductMetaData contains all meta data concerning the Product contract.
var ProductMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CurveMathOutOfBoundsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fixed18OverflowError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"Fixed18PackingOverflowError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"Fixed18PackingUnderflowError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidControllerError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"NotAccountOrMultiInvokerError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCollateralError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"coordinatorId\",\"type\":\"uint256\"}],\"name\":\"NotOwnerError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIProduct\",\"name\":\"product\",\"type\":\"address\"}],\"name\":\"NotProductError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamProviderInvalidFundingFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamProviderInvalidMakerFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamProviderInvalidPositionFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamProviderInvalidTakerFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumPayoffDefinitionLib.PayoffType\",\"name\":\"payoffType\",\"type\":\"uint8\"},{\"internalType\":\"bytes30\",\"name\":\"data\",\"type\":\"bytes30\"}],\"name\":\"PayoffDefinitionNotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumPayoffDefinitionLib.PayoffType\",\"name\":\"payoffType\",\"type\":\"uint8\"},{\"internalType\":\"enumPayoffDefinitionLib.PayoffDirection\",\"name\":\"payoffDirection\",\"type\":\"uint8\"}],\"name\":\"PayoffDefinitionUnsupportedTransform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayoffProviderInvalidOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayoffProviderInvalidPayoffDefinitionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PendingFeeUpdatesUnsupportedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductClosedError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductDoubleSidedError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductInLiquidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductInsufficientCollateralError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"socializationFactor\",\"type\":\"uint256\"}],\"name\":\"ProductInsufficientLiquidityError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductInvalidOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductMakerOverLimitError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductNotOwnerError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductOracleBootstrappingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProductOverClosedError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UFixed18PackingOverflowError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"UInitializableAlreadyInitializedError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UInitializableNotInitializingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UInitializableZeroVersionError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UReentrancyGuardReentrantCallError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toVersion\",\"type\":\"uint256\"}],\"name\":\"AccountSettle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"newClosed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ClosedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newFundingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"FundingFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"PackedFixed18\",\"name\":\"minRate\",\"type\":\"int128\"},{\"internalType\":\"PackedFixed18\",\"name\":\"maxRate\",\"type\":\"int128\"},{\"internalType\":\"PackedFixed18\",\"name\":\"targetRate\",\"type\":\"int128\"},{\"internalType\":\"PackedUFixed18\",\"name\":\"targetUtilization\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structJumpRateUtilizationCurve\",\"name\":\"\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"JumpRateUtilizationCurveUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newMaintenance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"MaintenanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MakeClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MakeOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newMakerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"MakerFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newMakerLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"MakerLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newMakerFee\",\"type\":\"uint256\"}],\"name\":\"PendingMakerFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newPositionFee\",\"type\":\"uint256\"}],\"name\":\"PendingPositionFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newTakerFee\",\"type\":\"uint256\"}],\"name\":\"PendingTakerFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newPositionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"PositionFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toVersion\",\"type\":\"uint256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TakeClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TakeOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"UFixed18\",\"name\":\"newTakerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"TakerFeeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oracleVersion\",\"type\":\"uint256\"}],\"name\":\"atVersion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"Fixed18\",\"name\":\"price\",\"type\":\"int256\"}],\"internalType\":\"structIOracleProvider.OracleVersion\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"closeMake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"closeMakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"closeTake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"closeTakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"Fixed18\",\"name\":\"price\",\"type\":\"int256\"}],\"internalType\":\"structIOracleProvider.OracleVersion\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingFee\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumPayoffDefinitionLib.PayoffType\",\"name\":\"payoffType\",\"type\":\"uint8\"},{\"internalType\":\"enumPayoffDefinitionLib.PayoffDirection\",\"name\":\"payoffDirection\",\"type\":\"uint8\"},{\"internalType\":\"bytes30\",\"name\":\"data\",\"type\":\"bytes30\"}],\"internalType\":\"structPayoffDefinition\",\"name\":\"payoffDefinition\",\"type\":\"tuple\"},{\"internalType\":\"contractIOracleProvider\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"maintenance\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"fundingFee\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"positionFee\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"makerLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"PackedFixed18\",\"name\":\"minRate\",\"type\":\"int128\"},{\"internalType\":\"PackedFixed18\",\"name\":\"maxRate\",\"type\":\"int128\"},{\"internalType\":\"PackedFixed18\",\"name\":\"targetRate\",\"type\":\"int128\"},{\"internalType\":\"PackedUFixed18\",\"name\":\"targetUtilization\",\"type\":\"uint128\"}],\"internalType\":\"structJumpRateUtilizationCurve\",\"name\":\"utilizationCurve\",\"type\":\"tuple\"}],\"internalType\":\"structIProduct.ProductInfo\",\"name\":\"productInfo_\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isLiquidating\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"latestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maintenance\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maintenance\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maintenanceNext\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"makerFee\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"makerLimit\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"openMake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"openMakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"openTake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"UFixed18\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"openTakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracleProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoffDefinition\",\"outputs\":[{\"components\":[{\"internalType\":\"enumPayoffDefinitionLib.PayoffType\",\"name\":\"payoffType\",\"type\":\"uint8\"},{\"internalType\":\"enumPayoffDefinitionLib.PayoffDirection\",\"name\":\"payoffDirection\",\"type\":\"uint8\"},{\"internalType\":\"bytes30\",\"name\":\"data\",\"type\":\"bytes30\"}],\"internalType\":\"structPayoffDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingFeeUpdates\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"makerFeeUpdated\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"pendingMakerFee\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"takerFeeUpdated\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"pendingTakerFee\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"positionFeeUpdated\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"pendingPositionFee\",\"type\":\"uint64\"}],\"internalType\":\"structPendingFeeUpdates\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"position\",\"outputs\":[{\"components\":[{\"internalType\":\"UFixed18\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"taker\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oracleVersion\",\"type\":\"uint256\"}],\"name\":\"positionAtVersion\",\"outputs\":[{\"components\":[{\"internalType\":\"UFixed18\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"taker\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionFee\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pre\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oracleVersion\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"UFixed18\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"taker\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"openPosition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"UFixed18\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"taker\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"closePosition\",\"type\":\"tuple\"}],\"internalType\":\"structPrePosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pre\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oracleVersion\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"UFixed18\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"taker\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"openPosition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"UFixed18\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"taker\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"closePosition\",\"type\":\"tuple\"}],\"internalType\":\"structPrePosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"UFixed18\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"UFixed18\",\"name\":\"taker\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"position_\",\"type\":\"tuple\"}],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"Fixed18\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"settleAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oracleVersion\",\"type\":\"uint256\"}],\"name\":\"shareAtVersion\",\"outputs\":[{\"components\":[{\"internalType\":\"Fixed18\",\"name\":\"maker\",\"type\":\"int256\"},{\"internalType\":\"Fixed18\",\"name\":\"taker\",\"type\":\"int256\"}],\"internalType\":\"structAccumulator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"takerFee\",\"outputs\":[{\"internalType\":\"UFixed18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newClosed\",\"type\":\"bool\"}],\"name\":\"updateClosed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"newFundingFee\",\"type\":\"uint256\"}],\"name\":\"updateFundingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"newMaintenance\",\"type\":\"uint256\"}],\"name\":\"updateMaintenance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"newMakerFee\",\"type\":\"uint256\"}],\"name\":\"updateMakerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"newMakerLimit\",\"type\":\"uint256\"}],\"name\":\"updateMakerLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"newPositionFee\",\"type\":\"uint256\"}],\"name\":\"updatePositionFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"UFixed18\",\"name\":\"newTakerFee\",\"type\":\"uint256\"}],\"name\":\"updateTakerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"PackedFixed18\",\"name\":\"minRate\",\"type\":\"int128\"},{\"internalType\":\"PackedFixed18\",\"name\":\"maxRate\",\"type\":\"int128\"},{\"internalType\":\"PackedFixed18\",\"name\":\"targetRate\",\"type\":\"int128\"},{\"internalType\":\"PackedUFixed18\",\"name\":\"targetUtilization\",\"type\":\"uint128\"}],\"internalType\":\"structJumpRateUtilizationCurve\",\"name\":\"newUtilizationCurve\",\"type\":\"tuple\"}],\"name\":\"updateUtilizationCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"utilizationCurve\",\"outputs\":[{\"components\":[{\"internalType\":\"PackedFixed18\",\"name\":\"minRate\",\"type\":\"int128\"},{\"internalType\":\"PackedFixed18\",\"name\":\"maxRate\",\"type\":\"int128\"},{\"internalType\":\"PackedFixed18\",\"name\":\"targetRate\",\"type\":\"int128\"},{\"internalType\":\"PackedUFixed18\",\"name\":\"targetUtilization\",\"type\":\"uint128\"}],\"internalType\":\"structJumpRateUtilizationCurve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oracleVersion\",\"type\":\"uint256\"}],\"name\":\"valueAtVersion\",\"outputs\":[{\"components\":[{\"internalType\":\"Fixed18\",\"name\":\"maker\",\"type\":\"int256\"},{\"internalType\":\"Fixed18\",\"name\":\"taker\",\"type\":\"int256\"}],\"internalType\":\"structAccumulator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProductABI is the input ABI used to generate the binding from.
// Deprecated: Use ProductMetaData.ABI instead.
var ProductABI = ProductMetaData.ABI

// Product is an auto generated Go binding around an Ethereum contract.
type Product struct {
	ProductCaller     // Read-only binding to the contract
	ProductTransactor // Write-only binding to the contract
	ProductFilterer   // Log filterer for contract events
}

// ProductCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProductCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProductTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProductFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProductSession struct {
	Contract     *Product          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProductCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProductCallerSession struct {
	Contract *ProductCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProductTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProductTransactorSession struct {
	Contract     *ProductTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProductRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProductRaw struct {
	Contract *Product // Generic contract binding to access the raw methods on
}

// ProductCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProductCallerRaw struct {
	Contract *ProductCaller // Generic read-only contract binding to access the raw methods on
}

// ProductTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProductTransactorRaw struct {
	Contract *ProductTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProduct creates a new instance of Product, bound to a specific deployed contract.
func NewProduct(address common.Address, backend bind.ContractBackend) (*Product, error) {
	contract, err := bindProduct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Product{ProductCaller: ProductCaller{contract: contract}, ProductTransactor: ProductTransactor{contract: contract}, ProductFilterer: ProductFilterer{contract: contract}}, nil
}

// NewProductCaller creates a new read-only instance of Product, bound to a specific deployed contract.
func NewProductCaller(address common.Address, caller bind.ContractCaller) (*ProductCaller, error) {
	contract, err := bindProduct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProductCaller{contract: contract}, nil
}

// NewProductTransactor creates a new write-only instance of Product, bound to a specific deployed contract.
func NewProductTransactor(address common.Address, transactor bind.ContractTransactor) (*ProductTransactor, error) {
	contract, err := bindProduct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProductTransactor{contract: contract}, nil
}

// NewProductFilterer creates a new log filterer instance of Product, bound to a specific deployed contract.
func NewProductFilterer(address common.Address, filterer bind.ContractFilterer) (*ProductFilterer, error) {
	contract, err := bindProduct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProductFilterer{contract: contract}, nil
}

// bindProduct binds a generic wrapper to an already deployed contract.
func bindProduct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProductABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Product *ProductRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Product.Contract.ProductCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Product *ProductRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Product.Contract.ProductTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Product *ProductRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Product.Contract.ProductTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Product *ProductCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Product.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Product *ProductTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Product.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Product *ProductTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Product.Contract.contract.Transact(opts, method, params...)
}

// AtVersion is a free data retrieval call binding the contract method 0x7ece075d.
//
// Solidity: function atVersion(uint256 oracleVersion) view returns((uint256,uint256,int256))
func (_Product *ProductCaller) AtVersion(opts *bind.CallOpts, oracleVersion *big.Int) (IOracleProviderOracleVersion, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "atVersion", oracleVersion)

	if err != nil {
		return *new(IOracleProviderOracleVersion), err
	}

	out0 := *abi.ConvertType(out[0], new(IOracleProviderOracleVersion)).(*IOracleProviderOracleVersion)

	return out0, err

}

// AtVersion is a free data retrieval call binding the contract method 0x7ece075d.
//
// Solidity: function atVersion(uint256 oracleVersion) view returns((uint256,uint256,int256))
func (_Product *ProductSession) AtVersion(oracleVersion *big.Int) (IOracleProviderOracleVersion, error) {
	return _Product.Contract.AtVersion(&_Product.CallOpts, oracleVersion)
}

// AtVersion is a free data retrieval call binding the contract method 0x7ece075d.
//
// Solidity: function atVersion(uint256 oracleVersion) view returns((uint256,uint256,int256))
func (_Product *ProductCallerSession) AtVersion(oracleVersion *big.Int) (IOracleProviderOracleVersion, error) {
	return _Product.Contract.AtVersion(&_Product.CallOpts, oracleVersion)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_Product *ProductCaller) Closed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "closed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_Product *ProductSession) Closed() (bool, error) {
	return _Product.Contract.Closed(&_Product.CallOpts)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_Product *ProductCallerSession) Closed() (bool, error) {
	return _Product.Contract.Closed(&_Product.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Product *ProductCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Product *ProductSession) Controller() (common.Address, error) {
	return _Product.Contract.Controller(&_Product.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Product *ProductCallerSession) Controller() (common.Address, error) {
	return _Product.Contract.Controller(&_Product.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns((uint256,uint256,int256))
func (_Product *ProductCaller) CurrentVersion(opts *bind.CallOpts) (IOracleProviderOracleVersion, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "currentVersion")

	if err != nil {
		return *new(IOracleProviderOracleVersion), err
	}

	out0 := *abi.ConvertType(out[0], new(IOracleProviderOracleVersion)).(*IOracleProviderOracleVersion)

	return out0, err

}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns((uint256,uint256,int256))
func (_Product *ProductSession) CurrentVersion() (IOracleProviderOracleVersion, error) {
	return _Product.Contract.CurrentVersion(&_Product.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns((uint256,uint256,int256))
func (_Product *ProductCallerSession) CurrentVersion() (IOracleProviderOracleVersion, error) {
	return _Product.Contract.CurrentVersion(&_Product.CallOpts)
}

// FundingFee is a free data retrieval call binding the contract method 0x5d16e120.
//
// Solidity: function fundingFee() view returns(uint256)
func (_Product *ProductCaller) FundingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "fundingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingFee is a free data retrieval call binding the contract method 0x5d16e120.
//
// Solidity: function fundingFee() view returns(uint256)
func (_Product *ProductSession) FundingFee() (*big.Int, error) {
	return _Product.Contract.FundingFee(&_Product.CallOpts)
}

// FundingFee is a free data retrieval call binding the contract method 0x5d16e120.
//
// Solidity: function fundingFee() view returns(uint256)
func (_Product *ProductCallerSession) FundingFee() (*big.Int, error) {
	return _Product.Contract.FundingFee(&_Product.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0x6943b017.
//
// Solidity: function isClosed(address account) view returns(bool)
func (_Product *ProductCaller) IsClosed(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "isClosed", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClosed is a free data retrieval call binding the contract method 0x6943b017.
//
// Solidity: function isClosed(address account) view returns(bool)
func (_Product *ProductSession) IsClosed(account common.Address) (bool, error) {
	return _Product.Contract.IsClosed(&_Product.CallOpts, account)
}

// IsClosed is a free data retrieval call binding the contract method 0x6943b017.
//
// Solidity: function isClosed(address account) view returns(bool)
func (_Product *ProductCallerSession) IsClosed(account common.Address) (bool, error) {
	return _Product.Contract.IsClosed(&_Product.CallOpts, account)
}

// IsLiquidating is a free data retrieval call binding the contract method 0x58ca6f98.
//
// Solidity: function isLiquidating(address account) view returns(bool)
func (_Product *ProductCaller) IsLiquidating(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "isLiquidating", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidating is a free data retrieval call binding the contract method 0x58ca6f98.
//
// Solidity: function isLiquidating(address account) view returns(bool)
func (_Product *ProductSession) IsLiquidating(account common.Address) (bool, error) {
	return _Product.Contract.IsLiquidating(&_Product.CallOpts, account)
}

// IsLiquidating is a free data retrieval call binding the contract method 0x58ca6f98.
//
// Solidity: function isLiquidating(address account) view returns(bool)
func (_Product *ProductCallerSession) IsLiquidating(account common.Address) (bool, error) {
	return _Product.Contract.IsLiquidating(&_Product.CallOpts, account)
}

// LatestVersion is a free data retrieval call binding the contract method 0x8e480b20.
//
// Solidity: function latestVersion(address account) view returns(uint256)
func (_Product *ProductCaller) LatestVersion(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "latestVersion", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersion is a free data retrieval call binding the contract method 0x8e480b20.
//
// Solidity: function latestVersion(address account) view returns(uint256)
func (_Product *ProductSession) LatestVersion(account common.Address) (*big.Int, error) {
	return _Product.Contract.LatestVersion(&_Product.CallOpts, account)
}

// LatestVersion is a free data retrieval call binding the contract method 0x8e480b20.
//
// Solidity: function latestVersion(address account) view returns(uint256)
func (_Product *ProductCallerSession) LatestVersion(account common.Address) (*big.Int, error) {
	return _Product.Contract.LatestVersion(&_Product.CallOpts, account)
}

// LatestVersion0 is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_Product *ProductCaller) LatestVersion0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "latestVersion0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersion0 is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_Product *ProductSession) LatestVersion0() (*big.Int, error) {
	return _Product.Contract.LatestVersion0(&_Product.CallOpts)
}

// LatestVersion0 is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_Product *ProductCallerSession) LatestVersion0() (*big.Int, error) {
	return _Product.Contract.LatestVersion0(&_Product.CallOpts)
}

// Maintenance is a free data retrieval call binding the contract method 0x6c376cc5.
//
// Solidity: function maintenance() view returns(uint256)
func (_Product *ProductCaller) Maintenance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "maintenance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maintenance is a free data retrieval call binding the contract method 0x6c376cc5.
//
// Solidity: function maintenance() view returns(uint256)
func (_Product *ProductSession) Maintenance() (*big.Int, error) {
	return _Product.Contract.Maintenance(&_Product.CallOpts)
}

// Maintenance is a free data retrieval call binding the contract method 0x6c376cc5.
//
// Solidity: function maintenance() view returns(uint256)
func (_Product *ProductCallerSession) Maintenance() (*big.Int, error) {
	return _Product.Contract.Maintenance(&_Product.CallOpts)
}

// Maintenance0 is a free data retrieval call binding the contract method 0x91689024.
//
// Solidity: function maintenance(address account) view returns(uint256)
func (_Product *ProductCaller) Maintenance0(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "maintenance0", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maintenance0 is a free data retrieval call binding the contract method 0x91689024.
//
// Solidity: function maintenance(address account) view returns(uint256)
func (_Product *ProductSession) Maintenance0(account common.Address) (*big.Int, error) {
	return _Product.Contract.Maintenance0(&_Product.CallOpts, account)
}

// Maintenance0 is a free data retrieval call binding the contract method 0x91689024.
//
// Solidity: function maintenance(address account) view returns(uint256)
func (_Product *ProductCallerSession) Maintenance0(account common.Address) (*big.Int, error) {
	return _Product.Contract.Maintenance0(&_Product.CallOpts, account)
}

// MaintenanceNext is a free data retrieval call binding the contract method 0xab582f29.
//
// Solidity: function maintenanceNext(address account) view returns(uint256)
func (_Product *ProductCaller) MaintenanceNext(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "maintenanceNext", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaintenanceNext is a free data retrieval call binding the contract method 0xab582f29.
//
// Solidity: function maintenanceNext(address account) view returns(uint256)
func (_Product *ProductSession) MaintenanceNext(account common.Address) (*big.Int, error) {
	return _Product.Contract.MaintenanceNext(&_Product.CallOpts, account)
}

// MaintenanceNext is a free data retrieval call binding the contract method 0xab582f29.
//
// Solidity: function maintenanceNext(address account) view returns(uint256)
func (_Product *ProductCallerSession) MaintenanceNext(account common.Address) (*big.Int, error) {
	return _Product.Contract.MaintenanceNext(&_Product.CallOpts, account)
}

// MakerFee is a free data retrieval call binding the contract method 0xfc741c7c.
//
// Solidity: function makerFee() view returns(uint256)
func (_Product *ProductCaller) MakerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "makerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MakerFee is a free data retrieval call binding the contract method 0xfc741c7c.
//
// Solidity: function makerFee() view returns(uint256)
func (_Product *ProductSession) MakerFee() (*big.Int, error) {
	return _Product.Contract.MakerFee(&_Product.CallOpts)
}

// MakerFee is a free data retrieval call binding the contract method 0xfc741c7c.
//
// Solidity: function makerFee() view returns(uint256)
func (_Product *ProductCallerSession) MakerFee() (*big.Int, error) {
	return _Product.Contract.MakerFee(&_Product.CallOpts)
}

// MakerLimit is a free data retrieval call binding the contract method 0x19377567.
//
// Solidity: function makerLimit() view returns(uint256)
func (_Product *ProductCaller) MakerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "makerLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MakerLimit is a free data retrieval call binding the contract method 0x19377567.
//
// Solidity: function makerLimit() view returns(uint256)
func (_Product *ProductSession) MakerLimit() (*big.Int, error) {
	return _Product.Contract.MakerLimit(&_Product.CallOpts)
}

// MakerLimit is a free data retrieval call binding the contract method 0x19377567.
//
// Solidity: function makerLimit() view returns(uint256)
func (_Product *ProductCallerSession) MakerLimit() (*big.Int, error) {
	return _Product.Contract.MakerLimit(&_Product.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Product *ProductCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Product *ProductSession) Name() (string, error) {
	return _Product.Contract.Name(&_Product.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Product *ProductCallerSession) Name() (string, error) {
	return _Product.Contract.Name(&_Product.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Product *ProductCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Product *ProductSession) Oracle() (common.Address, error) {
	return _Product.Contract.Oracle(&_Product.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Product *ProductCallerSession) Oracle() (common.Address, error) {
	return _Product.Contract.Oracle(&_Product.CallOpts)
}

// PayoffDefinition is a free data retrieval call binding the contract method 0x05d5c1cb.
//
// Solidity: function payoffDefinition() view returns((uint8,uint8,bytes30))
func (_Product *ProductCaller) PayoffDefinition(opts *bind.CallOpts) (PayoffDefinition, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "payoffDefinition")

	if err != nil {
		return *new(PayoffDefinition), err
	}

	out0 := *abi.ConvertType(out[0], new(PayoffDefinition)).(*PayoffDefinition)

	return out0, err

}

// PayoffDefinition is a free data retrieval call binding the contract method 0x05d5c1cb.
//
// Solidity: function payoffDefinition() view returns((uint8,uint8,bytes30))
func (_Product *ProductSession) PayoffDefinition() (PayoffDefinition, error) {
	return _Product.Contract.PayoffDefinition(&_Product.CallOpts)
}

// PayoffDefinition is a free data retrieval call binding the contract method 0x05d5c1cb.
//
// Solidity: function payoffDefinition() view returns((uint8,uint8,bytes30))
func (_Product *ProductCallerSession) PayoffDefinition() (PayoffDefinition, error) {
	return _Product.Contract.PayoffDefinition(&_Product.CallOpts)
}

// PendingFeeUpdates is a free data retrieval call binding the contract method 0x3a25c111.
//
// Solidity: function pendingFeeUpdates() view returns((bool,uint64,bool,uint64,bool,uint64))
func (_Product *ProductCaller) PendingFeeUpdates(opts *bind.CallOpts) (PendingFeeUpdates, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "pendingFeeUpdates")

	if err != nil {
		return *new(PendingFeeUpdates), err
	}

	out0 := *abi.ConvertType(out[0], new(PendingFeeUpdates)).(*PendingFeeUpdates)

	return out0, err

}

// PendingFeeUpdates is a free data retrieval call binding the contract method 0x3a25c111.
//
// Solidity: function pendingFeeUpdates() view returns((bool,uint64,bool,uint64,bool,uint64))
func (_Product *ProductSession) PendingFeeUpdates() (PendingFeeUpdates, error) {
	return _Product.Contract.PendingFeeUpdates(&_Product.CallOpts)
}

// PendingFeeUpdates is a free data retrieval call binding the contract method 0x3a25c111.
//
// Solidity: function pendingFeeUpdates() view returns((bool,uint64,bool,uint64,bool,uint64))
func (_Product *ProductCallerSession) PendingFeeUpdates() (PendingFeeUpdates, error) {
	return _Product.Contract.PendingFeeUpdates(&_Product.CallOpts)
}

// Position is a free data retrieval call binding the contract method 0xb7648fb9.
//
// Solidity: function position(address account) view returns((uint256,uint256))
func (_Product *ProductCaller) Position(opts *bind.CallOpts, account common.Address) (Position, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "position", account)

	if err != nil {
		return *new(Position), err
	}

	out0 := *abi.ConvertType(out[0], new(Position)).(*Position)

	return out0, err

}

// Position is a free data retrieval call binding the contract method 0xb7648fb9.
//
// Solidity: function position(address account) view returns((uint256,uint256))
func (_Product *ProductSession) Position(account common.Address) (Position, error) {
	return _Product.Contract.Position(&_Product.CallOpts, account)
}

// Position is a free data retrieval call binding the contract method 0xb7648fb9.
//
// Solidity: function position(address account) view returns((uint256,uint256))
func (_Product *ProductCallerSession) Position(account common.Address) (Position, error) {
	return _Product.Contract.Position(&_Product.CallOpts, account)
}

// PositionAtVersion is a free data retrieval call binding the contract method 0x9a427d03.
//
// Solidity: function positionAtVersion(uint256 oracleVersion) view returns((uint256,uint256))
func (_Product *ProductCaller) PositionAtVersion(opts *bind.CallOpts, oracleVersion *big.Int) (Position, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "positionAtVersion", oracleVersion)

	if err != nil {
		return *new(Position), err
	}

	out0 := *abi.ConvertType(out[0], new(Position)).(*Position)

	return out0, err

}

// PositionAtVersion is a free data retrieval call binding the contract method 0x9a427d03.
//
// Solidity: function positionAtVersion(uint256 oracleVersion) view returns((uint256,uint256))
func (_Product *ProductSession) PositionAtVersion(oracleVersion *big.Int) (Position, error) {
	return _Product.Contract.PositionAtVersion(&_Product.CallOpts, oracleVersion)
}

// PositionAtVersion is a free data retrieval call binding the contract method 0x9a427d03.
//
// Solidity: function positionAtVersion(uint256 oracleVersion) view returns((uint256,uint256))
func (_Product *ProductCallerSession) PositionAtVersion(oracleVersion *big.Int) (Position, error) {
	return _Product.Contract.PositionAtVersion(&_Product.CallOpts, oracleVersion)
}

// PositionFee is a free data retrieval call binding the contract method 0x0dbc9ce1.
//
// Solidity: function positionFee() view returns(uint256)
func (_Product *ProductCaller) PositionFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "positionFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionFee is a free data retrieval call binding the contract method 0x0dbc9ce1.
//
// Solidity: function positionFee() view returns(uint256)
func (_Product *ProductSession) PositionFee() (*big.Int, error) {
	return _Product.Contract.PositionFee(&_Product.CallOpts)
}

// PositionFee is a free data retrieval call binding the contract method 0x0dbc9ce1.
//
// Solidity: function positionFee() view returns(uint256)
func (_Product *ProductCallerSession) PositionFee() (*big.Int, error) {
	return _Product.Contract.PositionFee(&_Product.CallOpts)
}

// Pre is a free data retrieval call binding the contract method 0x1e0c6fb9.
//
// Solidity: function pre(address account) view returns((uint256,(uint256,uint256),(uint256,uint256)))
func (_Product *ProductCaller) Pre(opts *bind.CallOpts, account common.Address) (PrePosition, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "pre", account)

	if err != nil {
		return *new(PrePosition), err
	}

	out0 := *abi.ConvertType(out[0], new(PrePosition)).(*PrePosition)

	return out0, err

}

// Pre is a free data retrieval call binding the contract method 0x1e0c6fb9.
//
// Solidity: function pre(address account) view returns((uint256,(uint256,uint256),(uint256,uint256)))
func (_Product *ProductSession) Pre(account common.Address) (PrePosition, error) {
	return _Product.Contract.Pre(&_Product.CallOpts, account)
}

// Pre is a free data retrieval call binding the contract method 0x1e0c6fb9.
//
// Solidity: function pre(address account) view returns((uint256,(uint256,uint256),(uint256,uint256)))
func (_Product *ProductCallerSession) Pre(account common.Address) (PrePosition, error) {
	return _Product.Contract.Pre(&_Product.CallOpts, account)
}

// Pre0 is a free data retrieval call binding the contract method 0x59ea287d.
//
// Solidity: function pre() view returns((uint256,(uint256,uint256),(uint256,uint256)))
func (_Product *ProductCaller) Pre0(opts *bind.CallOpts) (PrePosition, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "pre0")

	if err != nil {
		return *new(PrePosition), err
	}

	out0 := *abi.ConvertType(out[0], new(PrePosition)).(*PrePosition)

	return out0, err

}

// Pre0 is a free data retrieval call binding the contract method 0x59ea287d.
//
// Solidity: function pre() view returns((uint256,(uint256,uint256),(uint256,uint256)))
func (_Product *ProductSession) Pre0() (PrePosition, error) {
	return _Product.Contract.Pre0(&_Product.CallOpts)
}

// Pre0 is a free data retrieval call binding the contract method 0x59ea287d.
//
// Solidity: function pre() view returns((uint256,(uint256,uint256),(uint256,uint256)))
func (_Product *ProductCallerSession) Pre0() (PrePosition, error) {
	return _Product.Contract.Pre0(&_Product.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x60ec91d6.
//
// Solidity: function rate((uint256,uint256) position_) view returns(int256)
func (_Product *ProductCaller) Rate(opts *bind.CallOpts, position_ Position) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "rate", position_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x60ec91d6.
//
// Solidity: function rate((uint256,uint256) position_) view returns(int256)
func (_Product *ProductSession) Rate(position_ Position) (*big.Int, error) {
	return _Product.Contract.Rate(&_Product.CallOpts, position_)
}

// Rate is a free data retrieval call binding the contract method 0x60ec91d6.
//
// Solidity: function rate((uint256,uint256) position_) view returns(int256)
func (_Product *ProductCallerSession) Rate(position_ Position) (*big.Int, error) {
	return _Product.Contract.Rate(&_Product.CallOpts, position_)
}

// ShareAtVersion is a free data retrieval call binding the contract method 0x476fa96d.
//
// Solidity: function shareAtVersion(uint256 oracleVersion) view returns((int256,int256))
func (_Product *ProductCaller) ShareAtVersion(opts *bind.CallOpts, oracleVersion *big.Int) (Accumulator, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "shareAtVersion", oracleVersion)

	if err != nil {
		return *new(Accumulator), err
	}

	out0 := *abi.ConvertType(out[0], new(Accumulator)).(*Accumulator)

	return out0, err

}

// ShareAtVersion is a free data retrieval call binding the contract method 0x476fa96d.
//
// Solidity: function shareAtVersion(uint256 oracleVersion) view returns((int256,int256))
func (_Product *ProductSession) ShareAtVersion(oracleVersion *big.Int) (Accumulator, error) {
	return _Product.Contract.ShareAtVersion(&_Product.CallOpts, oracleVersion)
}

// ShareAtVersion is a free data retrieval call binding the contract method 0x476fa96d.
//
// Solidity: function shareAtVersion(uint256 oracleVersion) view returns((int256,int256))
func (_Product *ProductCallerSession) ShareAtVersion(oracleVersion *big.Int) (Accumulator, error) {
	return _Product.Contract.ShareAtVersion(&_Product.CallOpts, oracleVersion)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Product *ProductCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Product *ProductSession) Symbol() (string, error) {
	return _Product.Contract.Symbol(&_Product.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Product *ProductCallerSession) Symbol() (string, error) {
	return _Product.Contract.Symbol(&_Product.CallOpts)
}

// TakerFee is a free data retrieval call binding the contract method 0x43f0179b.
//
// Solidity: function takerFee() view returns(uint256)
func (_Product *ProductCaller) TakerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "takerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TakerFee is a free data retrieval call binding the contract method 0x43f0179b.
//
// Solidity: function takerFee() view returns(uint256)
func (_Product *ProductSession) TakerFee() (*big.Int, error) {
	return _Product.Contract.TakerFee(&_Product.CallOpts)
}

// TakerFee is a free data retrieval call binding the contract method 0x43f0179b.
//
// Solidity: function takerFee() view returns(uint256)
func (_Product *ProductCallerSession) TakerFee() (*big.Int, error) {
	return _Product.Contract.TakerFee(&_Product.CallOpts)
}

// UtilizationCurve is a free data retrieval call binding the contract method 0xa12e1b33.
//
// Solidity: function utilizationCurve() view returns((int128,int128,int128,uint128))
func (_Product *ProductCaller) UtilizationCurve(opts *bind.CallOpts) (JumpRateUtilizationCurve, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "utilizationCurve")

	if err != nil {
		return *new(JumpRateUtilizationCurve), err
	}

	out0 := *abi.ConvertType(out[0], new(JumpRateUtilizationCurve)).(*JumpRateUtilizationCurve)

	return out0, err

}

// UtilizationCurve is a free data retrieval call binding the contract method 0xa12e1b33.
//
// Solidity: function utilizationCurve() view returns((int128,int128,int128,uint128))
func (_Product *ProductSession) UtilizationCurve() (JumpRateUtilizationCurve, error) {
	return _Product.Contract.UtilizationCurve(&_Product.CallOpts)
}

// UtilizationCurve is a free data retrieval call binding the contract method 0xa12e1b33.
//
// Solidity: function utilizationCurve() view returns((int128,int128,int128,uint128))
func (_Product *ProductCallerSession) UtilizationCurve() (JumpRateUtilizationCurve, error) {
	return _Product.Contract.UtilizationCurve(&_Product.CallOpts)
}

// ValueAtVersion is a free data retrieval call binding the contract method 0x20fe9c3c.
//
// Solidity: function valueAtVersion(uint256 oracleVersion) view returns((int256,int256))
func (_Product *ProductCaller) ValueAtVersion(opts *bind.CallOpts, oracleVersion *big.Int) (Accumulator, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "valueAtVersion", oracleVersion)

	if err != nil {
		return *new(Accumulator), err
	}

	out0 := *abi.ConvertType(out[0], new(Accumulator)).(*Accumulator)

	return out0, err

}

// ValueAtVersion is a free data retrieval call binding the contract method 0x20fe9c3c.
//
// Solidity: function valueAtVersion(uint256 oracleVersion) view returns((int256,int256))
func (_Product *ProductSession) ValueAtVersion(oracleVersion *big.Int) (Accumulator, error) {
	return _Product.Contract.ValueAtVersion(&_Product.CallOpts, oracleVersion)
}

// ValueAtVersion is a free data retrieval call binding the contract method 0x20fe9c3c.
//
// Solidity: function valueAtVersion(uint256 oracleVersion) view returns((int256,int256))
func (_Product *ProductCallerSession) ValueAtVersion(oracleVersion *big.Int) (Accumulator, error) {
	return _Product.Contract.ValueAtVersion(&_Product.CallOpts, oracleVersion)
}

// CloseAll is a paid mutator transaction binding the contract method 0xf6b32008.
//
// Solidity: function closeAll(address account) returns()
func (_Product *ProductTransactor) CloseAll(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "closeAll", account)
}

// CloseAll is a paid mutator transaction binding the contract method 0xf6b32008.
//
// Solidity: function closeAll(address account) returns()
func (_Product *ProductSession) CloseAll(account common.Address) (*types.Transaction, error) {
	return _Product.Contract.CloseAll(&_Product.TransactOpts, account)
}

// CloseAll is a paid mutator transaction binding the contract method 0xf6b32008.
//
// Solidity: function closeAll(address account) returns()
func (_Product *ProductTransactorSession) CloseAll(account common.Address) (*types.Transaction, error) {
	return _Product.Contract.CloseAll(&_Product.TransactOpts, account)
}

// CloseMake is a paid mutator transaction binding the contract method 0x59218fe9.
//
// Solidity: function closeMake(uint256 amount) returns()
func (_Product *ProductTransactor) CloseMake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "closeMake", amount)
}

// CloseMake is a paid mutator transaction binding the contract method 0x59218fe9.
//
// Solidity: function closeMake(uint256 amount) returns()
func (_Product *ProductSession) CloseMake(amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CloseMake(&_Product.TransactOpts, amount)
}

// CloseMake is a paid mutator transaction binding the contract method 0x59218fe9.
//
// Solidity: function closeMake(uint256 amount) returns()
func (_Product *ProductTransactorSession) CloseMake(amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CloseMake(&_Product.TransactOpts, amount)
}

// CloseMakeFor is a paid mutator transaction binding the contract method 0xe503b007.
//
// Solidity: function closeMakeFor(address account, uint256 amount) returns()
func (_Product *ProductTransactor) CloseMakeFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "closeMakeFor", account, amount)
}

// CloseMakeFor is a paid mutator transaction binding the contract method 0xe503b007.
//
// Solidity: function closeMakeFor(address account, uint256 amount) returns()
func (_Product *ProductSession) CloseMakeFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CloseMakeFor(&_Product.TransactOpts, account, amount)
}

// CloseMakeFor is a paid mutator transaction binding the contract method 0xe503b007.
//
// Solidity: function closeMakeFor(address account, uint256 amount) returns()
func (_Product *ProductTransactorSession) CloseMakeFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CloseMakeFor(&_Product.TransactOpts, account, amount)
}

// CloseTake is a paid mutator transaction binding the contract method 0x76f37001.
//
// Solidity: function closeTake(uint256 amount) returns()
func (_Product *ProductTransactor) CloseTake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "closeTake", amount)
}

// CloseTake is a paid mutator transaction binding the contract method 0x76f37001.
//
// Solidity: function closeTake(uint256 amount) returns()
func (_Product *ProductSession) CloseTake(amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CloseTake(&_Product.TransactOpts, amount)
}

// CloseTake is a paid mutator transaction binding the contract method 0x76f37001.
//
// Solidity: function closeTake(uint256 amount) returns()
func (_Product *ProductTransactorSession) CloseTake(amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CloseTake(&_Product.TransactOpts, amount)
}

// CloseTakeFor is a paid mutator transaction binding the contract method 0x2131ea4a.
//
// Solidity: function closeTakeFor(address account, uint256 amount) returns()
func (_Product *ProductTransactor) CloseTakeFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "closeTakeFor", account, amount)
}

// CloseTakeFor is a paid mutator transaction binding the contract method 0x2131ea4a.
//
// Solidity: function closeTakeFor(address account, uint256 amount) returns()
func (_Product *ProductSession) CloseTakeFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CloseTakeFor(&_Product.TransactOpts, account, amount)
}

// CloseTakeFor is a paid mutator transaction binding the contract method 0x2131ea4a.
//
// Solidity: function closeTakeFor(address account, uint256 amount) returns()
func (_Product *ProductTransactorSession) CloseTakeFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CloseTakeFor(&_Product.TransactOpts, account, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecffff9.
//
// Solidity: function initialize((string,string,(uint8,uint8,bytes30),address,uint256,uint256,uint256,uint256,uint256,uint256,(int128,int128,int128,uint128)) productInfo_) returns()
func (_Product *ProductTransactor) Initialize(opts *bind.TransactOpts, productInfo_ IProductProductInfo) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "initialize", productInfo_)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecffff9.
//
// Solidity: function initialize((string,string,(uint8,uint8,bytes30),address,uint256,uint256,uint256,uint256,uint256,uint256,(int128,int128,int128,uint128)) productInfo_) returns()
func (_Product *ProductSession) Initialize(productInfo_ IProductProductInfo) (*types.Transaction, error) {
	return _Product.Contract.Initialize(&_Product.TransactOpts, productInfo_)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecffff9.
//
// Solidity: function initialize((string,string,(uint8,uint8,bytes30),address,uint256,uint256,uint256,uint256,uint256,uint256,(int128,int128,int128,uint128)) productInfo_) returns()
func (_Product *ProductTransactorSession) Initialize(productInfo_ IProductProductInfo) (*types.Transaction, error) {
	return _Product.Contract.Initialize(&_Product.TransactOpts, productInfo_)
}

// OpenMake is a paid mutator transaction binding the contract method 0xd7d7d6b8.
//
// Solidity: function openMake(uint256 amount) returns()
func (_Product *ProductTransactor) OpenMake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "openMake", amount)
}

// OpenMake is a paid mutator transaction binding the contract method 0xd7d7d6b8.
//
// Solidity: function openMake(uint256 amount) returns()
func (_Product *ProductSession) OpenMake(amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.OpenMake(&_Product.TransactOpts, amount)
}

// OpenMake is a paid mutator transaction binding the contract method 0xd7d7d6b8.
//
// Solidity: function openMake(uint256 amount) returns()
func (_Product *ProductTransactorSession) OpenMake(amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.OpenMake(&_Product.TransactOpts, amount)
}

// OpenMakeFor is a paid mutator transaction binding the contract method 0x2d2e52be.
//
// Solidity: function openMakeFor(address account, uint256 amount) returns()
func (_Product *ProductTransactor) OpenMakeFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "openMakeFor", account, amount)
}

// OpenMakeFor is a paid mutator transaction binding the contract method 0x2d2e52be.
//
// Solidity: function openMakeFor(address account, uint256 amount) returns()
func (_Product *ProductSession) OpenMakeFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.OpenMakeFor(&_Product.TransactOpts, account, amount)
}

// OpenMakeFor is a paid mutator transaction binding the contract method 0x2d2e52be.
//
// Solidity: function openMakeFor(address account, uint256 amount) returns()
func (_Product *ProductTransactorSession) OpenMakeFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.OpenMakeFor(&_Product.TransactOpts, account, amount)
}

// OpenTake is a paid mutator transaction binding the contract method 0x73b88f3b.
//
// Solidity: function openTake(uint256 amount) returns()
func (_Product *ProductTransactor) OpenTake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "openTake", amount)
}

// OpenTake is a paid mutator transaction binding the contract method 0x73b88f3b.
//
// Solidity: function openTake(uint256 amount) returns()
func (_Product *ProductSession) OpenTake(amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.OpenTake(&_Product.TransactOpts, amount)
}

// OpenTake is a paid mutator transaction binding the contract method 0x73b88f3b.
//
// Solidity: function openTake(uint256 amount) returns()
func (_Product *ProductTransactorSession) OpenTake(amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.OpenTake(&_Product.TransactOpts, amount)
}

// OpenTakeFor is a paid mutator transaction binding the contract method 0x9378bf7b.
//
// Solidity: function openTakeFor(address account, uint256 amount) returns()
func (_Product *ProductTransactor) OpenTakeFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "openTakeFor", account, amount)
}

// OpenTakeFor is a paid mutator transaction binding the contract method 0x9378bf7b.
//
// Solidity: function openTakeFor(address account, uint256 amount) returns()
func (_Product *ProductSession) OpenTakeFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.OpenTakeFor(&_Product.TransactOpts, account, amount)
}

// OpenTakeFor is a paid mutator transaction binding the contract method 0x9378bf7b.
//
// Solidity: function openTakeFor(address account, uint256 amount) returns()
func (_Product *ProductTransactorSession) OpenTakeFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Product.Contract.OpenTakeFor(&_Product.TransactOpts, account, amount)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_Product *ProductTransactor) Settle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "settle")
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_Product *ProductSession) Settle() (*types.Transaction, error) {
	return _Product.Contract.Settle(&_Product.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() returns()
func (_Product *ProductTransactorSession) Settle() (*types.Transaction, error) {
	return _Product.Contract.Settle(&_Product.TransactOpts)
}

// SettleAccount is a paid mutator transaction binding the contract method 0xf667f897.
//
// Solidity: function settleAccount(address account) returns()
func (_Product *ProductTransactor) SettleAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "settleAccount", account)
}

// SettleAccount is a paid mutator transaction binding the contract method 0xf667f897.
//
// Solidity: function settleAccount(address account) returns()
func (_Product *ProductSession) SettleAccount(account common.Address) (*types.Transaction, error) {
	return _Product.Contract.SettleAccount(&_Product.TransactOpts, account)
}

// SettleAccount is a paid mutator transaction binding the contract method 0xf667f897.
//
// Solidity: function settleAccount(address account) returns()
func (_Product *ProductTransactorSession) SettleAccount(account common.Address) (*types.Transaction, error) {
	return _Product.Contract.SettleAccount(&_Product.TransactOpts, account)
}

// UpdateClosed is a paid mutator transaction binding the contract method 0x212e0ad3.
//
// Solidity: function updateClosed(bool newClosed) returns()
func (_Product *ProductTransactor) UpdateClosed(opts *bind.TransactOpts, newClosed bool) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "updateClosed", newClosed)
}

// UpdateClosed is a paid mutator transaction binding the contract method 0x212e0ad3.
//
// Solidity: function updateClosed(bool newClosed) returns()
func (_Product *ProductSession) UpdateClosed(newClosed bool) (*types.Transaction, error) {
	return _Product.Contract.UpdateClosed(&_Product.TransactOpts, newClosed)
}

// UpdateClosed is a paid mutator transaction binding the contract method 0x212e0ad3.
//
// Solidity: function updateClosed(bool newClosed) returns()
func (_Product *ProductTransactorSession) UpdateClosed(newClosed bool) (*types.Transaction, error) {
	return _Product.Contract.UpdateClosed(&_Product.TransactOpts, newClosed)
}

// UpdateFundingFee is a paid mutator transaction binding the contract method 0xe1c3c98d.
//
// Solidity: function updateFundingFee(uint256 newFundingFee) returns()
func (_Product *ProductTransactor) UpdateFundingFee(opts *bind.TransactOpts, newFundingFee *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "updateFundingFee", newFundingFee)
}

// UpdateFundingFee is a paid mutator transaction binding the contract method 0xe1c3c98d.
//
// Solidity: function updateFundingFee(uint256 newFundingFee) returns()
func (_Product *ProductSession) UpdateFundingFee(newFundingFee *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateFundingFee(&_Product.TransactOpts, newFundingFee)
}

// UpdateFundingFee is a paid mutator transaction binding the contract method 0xe1c3c98d.
//
// Solidity: function updateFundingFee(uint256 newFundingFee) returns()
func (_Product *ProductTransactorSession) UpdateFundingFee(newFundingFee *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateFundingFee(&_Product.TransactOpts, newFundingFee)
}

// UpdateMaintenance is a paid mutator transaction binding the contract method 0x56bc1ad4.
//
// Solidity: function updateMaintenance(uint256 newMaintenance) returns()
func (_Product *ProductTransactor) UpdateMaintenance(opts *bind.TransactOpts, newMaintenance *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "updateMaintenance", newMaintenance)
}

// UpdateMaintenance is a paid mutator transaction binding the contract method 0x56bc1ad4.
//
// Solidity: function updateMaintenance(uint256 newMaintenance) returns()
func (_Product *ProductSession) UpdateMaintenance(newMaintenance *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateMaintenance(&_Product.TransactOpts, newMaintenance)
}

// UpdateMaintenance is a paid mutator transaction binding the contract method 0x56bc1ad4.
//
// Solidity: function updateMaintenance(uint256 newMaintenance) returns()
func (_Product *ProductTransactorSession) UpdateMaintenance(newMaintenance *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateMaintenance(&_Product.TransactOpts, newMaintenance)
}

// UpdateMakerFee is a paid mutator transaction binding the contract method 0x3e17b8c7.
//
// Solidity: function updateMakerFee(uint256 newMakerFee) returns()
func (_Product *ProductTransactor) UpdateMakerFee(opts *bind.TransactOpts, newMakerFee *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "updateMakerFee", newMakerFee)
}

// UpdateMakerFee is a paid mutator transaction binding the contract method 0x3e17b8c7.
//
// Solidity: function updateMakerFee(uint256 newMakerFee) returns()
func (_Product *ProductSession) UpdateMakerFee(newMakerFee *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateMakerFee(&_Product.TransactOpts, newMakerFee)
}

// UpdateMakerFee is a paid mutator transaction binding the contract method 0x3e17b8c7.
//
// Solidity: function updateMakerFee(uint256 newMakerFee) returns()
func (_Product *ProductTransactorSession) UpdateMakerFee(newMakerFee *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateMakerFee(&_Product.TransactOpts, newMakerFee)
}

// UpdateMakerLimit is a paid mutator transaction binding the contract method 0x8c94b48f.
//
// Solidity: function updateMakerLimit(uint256 newMakerLimit) returns()
func (_Product *ProductTransactor) UpdateMakerLimit(opts *bind.TransactOpts, newMakerLimit *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "updateMakerLimit", newMakerLimit)
}

// UpdateMakerLimit is a paid mutator transaction binding the contract method 0x8c94b48f.
//
// Solidity: function updateMakerLimit(uint256 newMakerLimit) returns()
func (_Product *ProductSession) UpdateMakerLimit(newMakerLimit *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateMakerLimit(&_Product.TransactOpts, newMakerLimit)
}

// UpdateMakerLimit is a paid mutator transaction binding the contract method 0x8c94b48f.
//
// Solidity: function updateMakerLimit(uint256 newMakerLimit) returns()
func (_Product *ProductTransactorSession) UpdateMakerLimit(newMakerLimit *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateMakerLimit(&_Product.TransactOpts, newMakerLimit)
}

// UpdatePositionFee is a paid mutator transaction binding the contract method 0x54fb3921.
//
// Solidity: function updatePositionFee(uint256 newPositionFee) returns()
func (_Product *ProductTransactor) UpdatePositionFee(opts *bind.TransactOpts, newPositionFee *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "updatePositionFee", newPositionFee)
}

// UpdatePositionFee is a paid mutator transaction binding the contract method 0x54fb3921.
//
// Solidity: function updatePositionFee(uint256 newPositionFee) returns()
func (_Product *ProductSession) UpdatePositionFee(newPositionFee *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdatePositionFee(&_Product.TransactOpts, newPositionFee)
}

// UpdatePositionFee is a paid mutator transaction binding the contract method 0x54fb3921.
//
// Solidity: function updatePositionFee(uint256 newPositionFee) returns()
func (_Product *ProductTransactorSession) UpdatePositionFee(newPositionFee *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdatePositionFee(&_Product.TransactOpts, newPositionFee)
}

// UpdateTakerFee is a paid mutator transaction binding the contract method 0x611c71b4.
//
// Solidity: function updateTakerFee(uint256 newTakerFee) returns()
func (_Product *ProductTransactor) UpdateTakerFee(opts *bind.TransactOpts, newTakerFee *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "updateTakerFee", newTakerFee)
}

// UpdateTakerFee is a paid mutator transaction binding the contract method 0x611c71b4.
//
// Solidity: function updateTakerFee(uint256 newTakerFee) returns()
func (_Product *ProductSession) UpdateTakerFee(newTakerFee *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateTakerFee(&_Product.TransactOpts, newTakerFee)
}

// UpdateTakerFee is a paid mutator transaction binding the contract method 0x611c71b4.
//
// Solidity: function updateTakerFee(uint256 newTakerFee) returns()
func (_Product *ProductTransactorSession) UpdateTakerFee(newTakerFee *big.Int) (*types.Transaction, error) {
	return _Product.Contract.UpdateTakerFee(&_Product.TransactOpts, newTakerFee)
}

// UpdateUtilizationCurve is a paid mutator transaction binding the contract method 0x153261e5.
//
// Solidity: function updateUtilizationCurve((int128,int128,int128,uint128) newUtilizationCurve) returns()
func (_Product *ProductTransactor) UpdateUtilizationCurve(opts *bind.TransactOpts, newUtilizationCurve JumpRateUtilizationCurve) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "updateUtilizationCurve", newUtilizationCurve)
}

// UpdateUtilizationCurve is a paid mutator transaction binding the contract method 0x153261e5.
//
// Solidity: function updateUtilizationCurve((int128,int128,int128,uint128) newUtilizationCurve) returns()
func (_Product *ProductSession) UpdateUtilizationCurve(newUtilizationCurve JumpRateUtilizationCurve) (*types.Transaction, error) {
	return _Product.Contract.UpdateUtilizationCurve(&_Product.TransactOpts, newUtilizationCurve)
}

// UpdateUtilizationCurve is a paid mutator transaction binding the contract method 0x153261e5.
//
// Solidity: function updateUtilizationCurve((int128,int128,int128,uint128) newUtilizationCurve) returns()
func (_Product *ProductTransactorSession) UpdateUtilizationCurve(newUtilizationCurve JumpRateUtilizationCurve) (*types.Transaction, error) {
	return _Product.Contract.UpdateUtilizationCurve(&_Product.TransactOpts, newUtilizationCurve)
}

// ProductAccountSettleIterator is returned from FilterAccountSettle and is used to iterate over the raw logs and unpacked data for AccountSettle events raised by the Product contract.
type ProductAccountSettleIterator struct {
	Event *ProductAccountSettle // Event containing the contract specifics and raw log

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
func (it *ProductAccountSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductAccountSettle)
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
		it.Event = new(ProductAccountSettle)
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
func (it *ProductAccountSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductAccountSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductAccountSettle represents a AccountSettle event raised by the Product contract.
type ProductAccountSettle struct {
	Account    common.Address
	PreVersion *big.Int
	ToVersion  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountSettle is a free log retrieval operation binding the contract event 0x9d7055d24918d8c2fd08660a27bf31d4086fa71a51cd07874276470223aa480f.
//
// Solidity: event AccountSettle(address indexed account, uint256 preVersion, uint256 toVersion)
func (_Product *ProductFilterer) FilterAccountSettle(opts *bind.FilterOpts, account []common.Address) (*ProductAccountSettleIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.FilterLogs(opts, "AccountSettle", accountRule)
	if err != nil {
		return nil, err
	}
	return &ProductAccountSettleIterator{contract: _Product.contract, event: "AccountSettle", logs: logs, sub: sub}, nil
}

// WatchAccountSettle is a free log subscription operation binding the contract event 0x9d7055d24918d8c2fd08660a27bf31d4086fa71a51cd07874276470223aa480f.
//
// Solidity: event AccountSettle(address indexed account, uint256 preVersion, uint256 toVersion)
func (_Product *ProductFilterer) WatchAccountSettle(opts *bind.WatchOpts, sink chan<- *ProductAccountSettle, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.WatchLogs(opts, "AccountSettle", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductAccountSettle)
				if err := _Product.contract.UnpackLog(event, "AccountSettle", log); err != nil {
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

// ParseAccountSettle is a log parse operation binding the contract event 0x9d7055d24918d8c2fd08660a27bf31d4086fa71a51cd07874276470223aa480f.
//
// Solidity: event AccountSettle(address indexed account, uint256 preVersion, uint256 toVersion)
func (_Product *ProductFilterer) ParseAccountSettle(log types.Log) (*ProductAccountSettle, error) {
	event := new(ProductAccountSettle)
	if err := _Product.contract.UnpackLog(event, "AccountSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductClosedUpdatedIterator is returned from FilterClosedUpdated and is used to iterate over the raw logs and unpacked data for ClosedUpdated events raised by the Product contract.
type ProductClosedUpdatedIterator struct {
	Event *ProductClosedUpdated // Event containing the contract specifics and raw log

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
func (it *ProductClosedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductClosedUpdated)
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
		it.Event = new(ProductClosedUpdated)
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
func (it *ProductClosedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductClosedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductClosedUpdated represents a ClosedUpdated event raised by the Product contract.
type ProductClosedUpdated struct {
	NewClosed bool
	Version   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClosedUpdated is a free log retrieval operation binding the contract event 0x482763b305ee10cd21c16c0cbeed259f2e4fcacdc9767cef16dc1fbe483d3488.
//
// Solidity: event ClosedUpdated(bool indexed newClosed, uint256 version)
func (_Product *ProductFilterer) FilterClosedUpdated(opts *bind.FilterOpts, newClosed []bool) (*ProductClosedUpdatedIterator, error) {

	var newClosedRule []interface{}
	for _, newClosedItem := range newClosed {
		newClosedRule = append(newClosedRule, newClosedItem)
	}

	logs, sub, err := _Product.contract.FilterLogs(opts, "ClosedUpdated", newClosedRule)
	if err != nil {
		return nil, err
	}
	return &ProductClosedUpdatedIterator{contract: _Product.contract, event: "ClosedUpdated", logs: logs, sub: sub}, nil
}

// WatchClosedUpdated is a free log subscription operation binding the contract event 0x482763b305ee10cd21c16c0cbeed259f2e4fcacdc9767cef16dc1fbe483d3488.
//
// Solidity: event ClosedUpdated(bool indexed newClosed, uint256 version)
func (_Product *ProductFilterer) WatchClosedUpdated(opts *bind.WatchOpts, sink chan<- *ProductClosedUpdated, newClosed []bool) (event.Subscription, error) {

	var newClosedRule []interface{}
	for _, newClosedItem := range newClosed {
		newClosedRule = append(newClosedRule, newClosedItem)
	}

	logs, sub, err := _Product.contract.WatchLogs(opts, "ClosedUpdated", newClosedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductClosedUpdated)
				if err := _Product.contract.UnpackLog(event, "ClosedUpdated", log); err != nil {
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

// ParseClosedUpdated is a log parse operation binding the contract event 0x482763b305ee10cd21c16c0cbeed259f2e4fcacdc9767cef16dc1fbe483d3488.
//
// Solidity: event ClosedUpdated(bool indexed newClosed, uint256 version)
func (_Product *ProductFilterer) ParseClosedUpdated(log types.Log) (*ProductClosedUpdated, error) {
	event := new(ProductClosedUpdated)
	if err := _Product.contract.UnpackLog(event, "ClosedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductFundingFeeUpdatedIterator is returned from FilterFundingFeeUpdated and is used to iterate over the raw logs and unpacked data for FundingFeeUpdated events raised by the Product contract.
type ProductFundingFeeUpdatedIterator struct {
	Event *ProductFundingFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ProductFundingFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductFundingFeeUpdated)
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
		it.Event = new(ProductFundingFeeUpdated)
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
func (it *ProductFundingFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductFundingFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductFundingFeeUpdated represents a FundingFeeUpdated event raised by the Product contract.
type ProductFundingFeeUpdated struct {
	NewFundingFee *big.Int
	Version       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFundingFeeUpdated is a free log retrieval operation binding the contract event 0xb2743902575ffdab882dfa43a3501f82fbbefa0d2c637e393aed1e80e2cb840b.
//
// Solidity: event FundingFeeUpdated(uint256 newFundingFee, uint256 version)
func (_Product *ProductFilterer) FilterFundingFeeUpdated(opts *bind.FilterOpts) (*ProductFundingFeeUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "FundingFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductFundingFeeUpdatedIterator{contract: _Product.contract, event: "FundingFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFundingFeeUpdated is a free log subscription operation binding the contract event 0xb2743902575ffdab882dfa43a3501f82fbbefa0d2c637e393aed1e80e2cb840b.
//
// Solidity: event FundingFeeUpdated(uint256 newFundingFee, uint256 version)
func (_Product *ProductFilterer) WatchFundingFeeUpdated(opts *bind.WatchOpts, sink chan<- *ProductFundingFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "FundingFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductFundingFeeUpdated)
				if err := _Product.contract.UnpackLog(event, "FundingFeeUpdated", log); err != nil {
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

// ParseFundingFeeUpdated is a log parse operation binding the contract event 0xb2743902575ffdab882dfa43a3501f82fbbefa0d2c637e393aed1e80e2cb840b.
//
// Solidity: event FundingFeeUpdated(uint256 newFundingFee, uint256 version)
func (_Product *ProductFilterer) ParseFundingFeeUpdated(log types.Log) (*ProductFundingFeeUpdated, error) {
	event := new(ProductFundingFeeUpdated)
	if err := _Product.contract.UnpackLog(event, "FundingFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Product contract.
type ProductInitializedIterator struct {
	Event *ProductInitialized // Event containing the contract specifics and raw log

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
func (it *ProductInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductInitialized)
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
		it.Event = new(ProductInitialized)
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
func (it *ProductInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductInitialized represents a Initialized event raised by the Product contract.
type ProductInitialized struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xbe9b076dc5b65990cca9dd9d7366682482e7817a6f6bc7f4faf4dc32af497f32.
//
// Solidity: event Initialized(uint256 version)
func (_Product *ProductFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProductInitializedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProductInitializedIterator{contract: _Product.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xbe9b076dc5b65990cca9dd9d7366682482e7817a6f6bc7f4faf4dc32af497f32.
//
// Solidity: event Initialized(uint256 version)
func (_Product *ProductFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProductInitialized) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductInitialized)
				if err := _Product.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xbe9b076dc5b65990cca9dd9d7366682482e7817a6f6bc7f4faf4dc32af497f32.
//
// Solidity: event Initialized(uint256 version)
func (_Product *ProductFilterer) ParseInitialized(log types.Log) (*ProductInitialized, error) {
	event := new(ProductInitialized)
	if err := _Product.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductJumpRateUtilizationCurveUpdatedIterator is returned from FilterJumpRateUtilizationCurveUpdated and is used to iterate over the raw logs and unpacked data for JumpRateUtilizationCurveUpdated events raised by the Product contract.
type ProductJumpRateUtilizationCurveUpdatedIterator struct {
	Event *ProductJumpRateUtilizationCurveUpdated // Event containing the contract specifics and raw log

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
func (it *ProductJumpRateUtilizationCurveUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductJumpRateUtilizationCurveUpdated)
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
		it.Event = new(ProductJumpRateUtilizationCurveUpdated)
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
func (it *ProductJumpRateUtilizationCurveUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductJumpRateUtilizationCurveUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductJumpRateUtilizationCurveUpdated represents a JumpRateUtilizationCurveUpdated event raised by the Product contract.
type ProductJumpRateUtilizationCurveUpdated struct {
	Arg0    JumpRateUtilizationCurve
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterJumpRateUtilizationCurveUpdated is a free log retrieval operation binding the contract event 0xb08b24c7f4e8b392d39bbaf1f4aa9bbb733ba9a467559eb8d9e2ffb634658e90.
//
// Solidity: event JumpRateUtilizationCurveUpdated((int128,int128,int128,uint128) arg0, uint256 version)
func (_Product *ProductFilterer) FilterJumpRateUtilizationCurveUpdated(opts *bind.FilterOpts) (*ProductJumpRateUtilizationCurveUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "JumpRateUtilizationCurveUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductJumpRateUtilizationCurveUpdatedIterator{contract: _Product.contract, event: "JumpRateUtilizationCurveUpdated", logs: logs, sub: sub}, nil
}

// WatchJumpRateUtilizationCurveUpdated is a free log subscription operation binding the contract event 0xb08b24c7f4e8b392d39bbaf1f4aa9bbb733ba9a467559eb8d9e2ffb634658e90.
//
// Solidity: event JumpRateUtilizationCurveUpdated((int128,int128,int128,uint128) arg0, uint256 version)
func (_Product *ProductFilterer) WatchJumpRateUtilizationCurveUpdated(opts *bind.WatchOpts, sink chan<- *ProductJumpRateUtilizationCurveUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "JumpRateUtilizationCurveUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductJumpRateUtilizationCurveUpdated)
				if err := _Product.contract.UnpackLog(event, "JumpRateUtilizationCurveUpdated", log); err != nil {
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

// ParseJumpRateUtilizationCurveUpdated is a log parse operation binding the contract event 0xb08b24c7f4e8b392d39bbaf1f4aa9bbb733ba9a467559eb8d9e2ffb634658e90.
//
// Solidity: event JumpRateUtilizationCurveUpdated((int128,int128,int128,uint128) arg0, uint256 version)
func (_Product *ProductFilterer) ParseJumpRateUtilizationCurveUpdated(log types.Log) (*ProductJumpRateUtilizationCurveUpdated, error) {
	event := new(ProductJumpRateUtilizationCurveUpdated)
	if err := _Product.contract.UnpackLog(event, "JumpRateUtilizationCurveUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductMaintenanceUpdatedIterator is returned from FilterMaintenanceUpdated and is used to iterate over the raw logs and unpacked data for MaintenanceUpdated events raised by the Product contract.
type ProductMaintenanceUpdatedIterator struct {
	Event *ProductMaintenanceUpdated // Event containing the contract specifics and raw log

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
func (it *ProductMaintenanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductMaintenanceUpdated)
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
		it.Event = new(ProductMaintenanceUpdated)
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
func (it *ProductMaintenanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductMaintenanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductMaintenanceUpdated represents a MaintenanceUpdated event raised by the Product contract.
type ProductMaintenanceUpdated struct {
	NewMaintenance *big.Int
	Version        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceUpdated is a free log retrieval operation binding the contract event 0x18eb97ecd7366a8be310b7e675d67335c7604d5f47446a54a222473625a27e22.
//
// Solidity: event MaintenanceUpdated(uint256 newMaintenance, uint256 version)
func (_Product *ProductFilterer) FilterMaintenanceUpdated(opts *bind.FilterOpts) (*ProductMaintenanceUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "MaintenanceUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductMaintenanceUpdatedIterator{contract: _Product.contract, event: "MaintenanceUpdated", logs: logs, sub: sub}, nil
}

// WatchMaintenanceUpdated is a free log subscription operation binding the contract event 0x18eb97ecd7366a8be310b7e675d67335c7604d5f47446a54a222473625a27e22.
//
// Solidity: event MaintenanceUpdated(uint256 newMaintenance, uint256 version)
func (_Product *ProductFilterer) WatchMaintenanceUpdated(opts *bind.WatchOpts, sink chan<- *ProductMaintenanceUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "MaintenanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductMaintenanceUpdated)
				if err := _Product.contract.UnpackLog(event, "MaintenanceUpdated", log); err != nil {
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

// ParseMaintenanceUpdated is a log parse operation binding the contract event 0x18eb97ecd7366a8be310b7e675d67335c7604d5f47446a54a222473625a27e22.
//
// Solidity: event MaintenanceUpdated(uint256 newMaintenance, uint256 version)
func (_Product *ProductFilterer) ParseMaintenanceUpdated(log types.Log) (*ProductMaintenanceUpdated, error) {
	event := new(ProductMaintenanceUpdated)
	if err := _Product.contract.UnpackLog(event, "MaintenanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductMakeClosedIterator is returned from FilterMakeClosed and is used to iterate over the raw logs and unpacked data for MakeClosed events raised by the Product contract.
type ProductMakeClosedIterator struct {
	Event *ProductMakeClosed // Event containing the contract specifics and raw log

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
func (it *ProductMakeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductMakeClosed)
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
		it.Event = new(ProductMakeClosed)
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
func (it *ProductMakeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductMakeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductMakeClosed represents a MakeClosed event raised by the Product contract.
type ProductMakeClosed struct {
	Account common.Address
	Version *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMakeClosed is a free log retrieval operation binding the contract event 0x39854479080fac0b5e7c0ecedb0fb02308a72a43cd102c6b9f918653d3400367.
//
// Solidity: event MakeClosed(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) FilterMakeClosed(opts *bind.FilterOpts, account []common.Address) (*ProductMakeClosedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.FilterLogs(opts, "MakeClosed", accountRule)
	if err != nil {
		return nil, err
	}
	return &ProductMakeClosedIterator{contract: _Product.contract, event: "MakeClosed", logs: logs, sub: sub}, nil
}

// WatchMakeClosed is a free log subscription operation binding the contract event 0x39854479080fac0b5e7c0ecedb0fb02308a72a43cd102c6b9f918653d3400367.
//
// Solidity: event MakeClosed(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) WatchMakeClosed(opts *bind.WatchOpts, sink chan<- *ProductMakeClosed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.WatchLogs(opts, "MakeClosed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductMakeClosed)
				if err := _Product.contract.UnpackLog(event, "MakeClosed", log); err != nil {
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

// ParseMakeClosed is a log parse operation binding the contract event 0x39854479080fac0b5e7c0ecedb0fb02308a72a43cd102c6b9f918653d3400367.
//
// Solidity: event MakeClosed(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) ParseMakeClosed(log types.Log) (*ProductMakeClosed, error) {
	event := new(ProductMakeClosed)
	if err := _Product.contract.UnpackLog(event, "MakeClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductMakeOpenedIterator is returned from FilterMakeOpened and is used to iterate over the raw logs and unpacked data for MakeOpened events raised by the Product contract.
type ProductMakeOpenedIterator struct {
	Event *ProductMakeOpened // Event containing the contract specifics and raw log

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
func (it *ProductMakeOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductMakeOpened)
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
		it.Event = new(ProductMakeOpened)
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
func (it *ProductMakeOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductMakeOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductMakeOpened represents a MakeOpened event raised by the Product contract.
type ProductMakeOpened struct {
	Account common.Address
	Version *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMakeOpened is a free log retrieval operation binding the contract event 0xf98b31465ac12e92b5cb136ade913276c267463c4395bb1a3999bc88fb837806.
//
// Solidity: event MakeOpened(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) FilterMakeOpened(opts *bind.FilterOpts, account []common.Address) (*ProductMakeOpenedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.FilterLogs(opts, "MakeOpened", accountRule)
	if err != nil {
		return nil, err
	}
	return &ProductMakeOpenedIterator{contract: _Product.contract, event: "MakeOpened", logs: logs, sub: sub}, nil
}

// WatchMakeOpened is a free log subscription operation binding the contract event 0xf98b31465ac12e92b5cb136ade913276c267463c4395bb1a3999bc88fb837806.
//
// Solidity: event MakeOpened(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) WatchMakeOpened(opts *bind.WatchOpts, sink chan<- *ProductMakeOpened, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.WatchLogs(opts, "MakeOpened", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductMakeOpened)
				if err := _Product.contract.UnpackLog(event, "MakeOpened", log); err != nil {
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

// ParseMakeOpened is a log parse operation binding the contract event 0xf98b31465ac12e92b5cb136ade913276c267463c4395bb1a3999bc88fb837806.
//
// Solidity: event MakeOpened(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) ParseMakeOpened(log types.Log) (*ProductMakeOpened, error) {
	event := new(ProductMakeOpened)
	if err := _Product.contract.UnpackLog(event, "MakeOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductMakerFeeUpdatedIterator is returned from FilterMakerFeeUpdated and is used to iterate over the raw logs and unpacked data for MakerFeeUpdated events raised by the Product contract.
type ProductMakerFeeUpdatedIterator struct {
	Event *ProductMakerFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ProductMakerFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductMakerFeeUpdated)
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
		it.Event = new(ProductMakerFeeUpdated)
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
func (it *ProductMakerFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductMakerFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductMakerFeeUpdated represents a MakerFeeUpdated event raised by the Product contract.
type ProductMakerFeeUpdated struct {
	NewMakerFee *big.Int
	Version     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMakerFeeUpdated is a free log retrieval operation binding the contract event 0x6bf5f38f4e94f47ae7a4d6e56dbf388c856f890b5f0f2808b66914710b0266ca.
//
// Solidity: event MakerFeeUpdated(uint256 newMakerFee, uint256 version)
func (_Product *ProductFilterer) FilterMakerFeeUpdated(opts *bind.FilterOpts) (*ProductMakerFeeUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "MakerFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductMakerFeeUpdatedIterator{contract: _Product.contract, event: "MakerFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchMakerFeeUpdated is a free log subscription operation binding the contract event 0x6bf5f38f4e94f47ae7a4d6e56dbf388c856f890b5f0f2808b66914710b0266ca.
//
// Solidity: event MakerFeeUpdated(uint256 newMakerFee, uint256 version)
func (_Product *ProductFilterer) WatchMakerFeeUpdated(opts *bind.WatchOpts, sink chan<- *ProductMakerFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "MakerFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductMakerFeeUpdated)
				if err := _Product.contract.UnpackLog(event, "MakerFeeUpdated", log); err != nil {
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

// ParseMakerFeeUpdated is a log parse operation binding the contract event 0x6bf5f38f4e94f47ae7a4d6e56dbf388c856f890b5f0f2808b66914710b0266ca.
//
// Solidity: event MakerFeeUpdated(uint256 newMakerFee, uint256 version)
func (_Product *ProductFilterer) ParseMakerFeeUpdated(log types.Log) (*ProductMakerFeeUpdated, error) {
	event := new(ProductMakerFeeUpdated)
	if err := _Product.contract.UnpackLog(event, "MakerFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductMakerLimitUpdatedIterator is returned from FilterMakerLimitUpdated and is used to iterate over the raw logs and unpacked data for MakerLimitUpdated events raised by the Product contract.
type ProductMakerLimitUpdatedIterator struct {
	Event *ProductMakerLimitUpdated // Event containing the contract specifics and raw log

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
func (it *ProductMakerLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductMakerLimitUpdated)
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
		it.Event = new(ProductMakerLimitUpdated)
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
func (it *ProductMakerLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductMakerLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductMakerLimitUpdated represents a MakerLimitUpdated event raised by the Product contract.
type ProductMakerLimitUpdated struct {
	NewMakerLimit *big.Int
	Version       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMakerLimitUpdated is a free log retrieval operation binding the contract event 0xec4e0c10e4e0d362960c0114996546aa9c4d6cc297b00a7577e127b720791beb.
//
// Solidity: event MakerLimitUpdated(uint256 newMakerLimit, uint256 version)
func (_Product *ProductFilterer) FilterMakerLimitUpdated(opts *bind.FilterOpts) (*ProductMakerLimitUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "MakerLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductMakerLimitUpdatedIterator{contract: _Product.contract, event: "MakerLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchMakerLimitUpdated is a free log subscription operation binding the contract event 0xec4e0c10e4e0d362960c0114996546aa9c4d6cc297b00a7577e127b720791beb.
//
// Solidity: event MakerLimitUpdated(uint256 newMakerLimit, uint256 version)
func (_Product *ProductFilterer) WatchMakerLimitUpdated(opts *bind.WatchOpts, sink chan<- *ProductMakerLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "MakerLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductMakerLimitUpdated)
				if err := _Product.contract.UnpackLog(event, "MakerLimitUpdated", log); err != nil {
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

// ParseMakerLimitUpdated is a log parse operation binding the contract event 0xec4e0c10e4e0d362960c0114996546aa9c4d6cc297b00a7577e127b720791beb.
//
// Solidity: event MakerLimitUpdated(uint256 newMakerLimit, uint256 version)
func (_Product *ProductFilterer) ParseMakerLimitUpdated(log types.Log) (*ProductMakerLimitUpdated, error) {
	event := new(ProductMakerLimitUpdated)
	if err := _Product.contract.UnpackLog(event, "MakerLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductPendingMakerFeeUpdatedIterator is returned from FilterPendingMakerFeeUpdated and is used to iterate over the raw logs and unpacked data for PendingMakerFeeUpdated events raised by the Product contract.
type ProductPendingMakerFeeUpdatedIterator struct {
	Event *ProductPendingMakerFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ProductPendingMakerFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductPendingMakerFeeUpdated)
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
		it.Event = new(ProductPendingMakerFeeUpdated)
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
func (it *ProductPendingMakerFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductPendingMakerFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductPendingMakerFeeUpdated represents a PendingMakerFeeUpdated event raised by the Product contract.
type ProductPendingMakerFeeUpdated struct {
	NewMakerFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingMakerFeeUpdated is a free log retrieval operation binding the contract event 0x5dcd1837ba774ef7f1f06e50ad3f64097201da74055346f26a4c7dea9af9084e.
//
// Solidity: event PendingMakerFeeUpdated(uint256 newMakerFee)
func (_Product *ProductFilterer) FilterPendingMakerFeeUpdated(opts *bind.FilterOpts) (*ProductPendingMakerFeeUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "PendingMakerFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductPendingMakerFeeUpdatedIterator{contract: _Product.contract, event: "PendingMakerFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPendingMakerFeeUpdated is a free log subscription operation binding the contract event 0x5dcd1837ba774ef7f1f06e50ad3f64097201da74055346f26a4c7dea9af9084e.
//
// Solidity: event PendingMakerFeeUpdated(uint256 newMakerFee)
func (_Product *ProductFilterer) WatchPendingMakerFeeUpdated(opts *bind.WatchOpts, sink chan<- *ProductPendingMakerFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "PendingMakerFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductPendingMakerFeeUpdated)
				if err := _Product.contract.UnpackLog(event, "PendingMakerFeeUpdated", log); err != nil {
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

// ParsePendingMakerFeeUpdated is a log parse operation binding the contract event 0x5dcd1837ba774ef7f1f06e50ad3f64097201da74055346f26a4c7dea9af9084e.
//
// Solidity: event PendingMakerFeeUpdated(uint256 newMakerFee)
func (_Product *ProductFilterer) ParsePendingMakerFeeUpdated(log types.Log) (*ProductPendingMakerFeeUpdated, error) {
	event := new(ProductPendingMakerFeeUpdated)
	if err := _Product.contract.UnpackLog(event, "PendingMakerFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductPendingPositionFeeUpdatedIterator is returned from FilterPendingPositionFeeUpdated and is used to iterate over the raw logs and unpacked data for PendingPositionFeeUpdated events raised by the Product contract.
type ProductPendingPositionFeeUpdatedIterator struct {
	Event *ProductPendingPositionFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ProductPendingPositionFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductPendingPositionFeeUpdated)
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
		it.Event = new(ProductPendingPositionFeeUpdated)
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
func (it *ProductPendingPositionFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductPendingPositionFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductPendingPositionFeeUpdated represents a PendingPositionFeeUpdated event raised by the Product contract.
type ProductPendingPositionFeeUpdated struct {
	NewPositionFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPendingPositionFeeUpdated is a free log retrieval operation binding the contract event 0x75d77d82706bf877275b9bc542ae4f30aca347d73558e05b56b35a1452f86132.
//
// Solidity: event PendingPositionFeeUpdated(uint256 newPositionFee)
func (_Product *ProductFilterer) FilterPendingPositionFeeUpdated(opts *bind.FilterOpts) (*ProductPendingPositionFeeUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "PendingPositionFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductPendingPositionFeeUpdatedIterator{contract: _Product.contract, event: "PendingPositionFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPendingPositionFeeUpdated is a free log subscription operation binding the contract event 0x75d77d82706bf877275b9bc542ae4f30aca347d73558e05b56b35a1452f86132.
//
// Solidity: event PendingPositionFeeUpdated(uint256 newPositionFee)
func (_Product *ProductFilterer) WatchPendingPositionFeeUpdated(opts *bind.WatchOpts, sink chan<- *ProductPendingPositionFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "PendingPositionFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductPendingPositionFeeUpdated)
				if err := _Product.contract.UnpackLog(event, "PendingPositionFeeUpdated", log); err != nil {
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

// ParsePendingPositionFeeUpdated is a log parse operation binding the contract event 0x75d77d82706bf877275b9bc542ae4f30aca347d73558e05b56b35a1452f86132.
//
// Solidity: event PendingPositionFeeUpdated(uint256 newPositionFee)
func (_Product *ProductFilterer) ParsePendingPositionFeeUpdated(log types.Log) (*ProductPendingPositionFeeUpdated, error) {
	event := new(ProductPendingPositionFeeUpdated)
	if err := _Product.contract.UnpackLog(event, "PendingPositionFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductPendingTakerFeeUpdatedIterator is returned from FilterPendingTakerFeeUpdated and is used to iterate over the raw logs and unpacked data for PendingTakerFeeUpdated events raised by the Product contract.
type ProductPendingTakerFeeUpdatedIterator struct {
	Event *ProductPendingTakerFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ProductPendingTakerFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductPendingTakerFeeUpdated)
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
		it.Event = new(ProductPendingTakerFeeUpdated)
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
func (it *ProductPendingTakerFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductPendingTakerFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductPendingTakerFeeUpdated represents a PendingTakerFeeUpdated event raised by the Product contract.
type ProductPendingTakerFeeUpdated struct {
	NewTakerFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingTakerFeeUpdated is a free log retrieval operation binding the contract event 0x4024ee9b51b90807c93e46f7e500580ef5b0b8491e1105188a44f62eed26b129.
//
// Solidity: event PendingTakerFeeUpdated(uint256 newTakerFee)
func (_Product *ProductFilterer) FilterPendingTakerFeeUpdated(opts *bind.FilterOpts) (*ProductPendingTakerFeeUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "PendingTakerFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductPendingTakerFeeUpdatedIterator{contract: _Product.contract, event: "PendingTakerFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPendingTakerFeeUpdated is a free log subscription operation binding the contract event 0x4024ee9b51b90807c93e46f7e500580ef5b0b8491e1105188a44f62eed26b129.
//
// Solidity: event PendingTakerFeeUpdated(uint256 newTakerFee)
func (_Product *ProductFilterer) WatchPendingTakerFeeUpdated(opts *bind.WatchOpts, sink chan<- *ProductPendingTakerFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "PendingTakerFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductPendingTakerFeeUpdated)
				if err := _Product.contract.UnpackLog(event, "PendingTakerFeeUpdated", log); err != nil {
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

// ParsePendingTakerFeeUpdated is a log parse operation binding the contract event 0x4024ee9b51b90807c93e46f7e500580ef5b0b8491e1105188a44f62eed26b129.
//
// Solidity: event PendingTakerFeeUpdated(uint256 newTakerFee)
func (_Product *ProductFilterer) ParsePendingTakerFeeUpdated(log types.Log) (*ProductPendingTakerFeeUpdated, error) {
	event := new(ProductPendingTakerFeeUpdated)
	if err := _Product.contract.UnpackLog(event, "PendingTakerFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductPositionFeeUpdatedIterator is returned from FilterPositionFeeUpdated and is used to iterate over the raw logs and unpacked data for PositionFeeUpdated events raised by the Product contract.
type ProductPositionFeeUpdatedIterator struct {
	Event *ProductPositionFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ProductPositionFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductPositionFeeUpdated)
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
		it.Event = new(ProductPositionFeeUpdated)
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
func (it *ProductPositionFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductPositionFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductPositionFeeUpdated represents a PositionFeeUpdated event raised by the Product contract.
type ProductPositionFeeUpdated struct {
	NewPositionFee *big.Int
	Version        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPositionFeeUpdated is a free log retrieval operation binding the contract event 0xd07cf37924efa8fca893c372ddfb1614e6fd8c02cfcc9205262535680b77b934.
//
// Solidity: event PositionFeeUpdated(uint256 newPositionFee, uint256 version)
func (_Product *ProductFilterer) FilterPositionFeeUpdated(opts *bind.FilterOpts) (*ProductPositionFeeUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "PositionFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductPositionFeeUpdatedIterator{contract: _Product.contract, event: "PositionFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPositionFeeUpdated is a free log subscription operation binding the contract event 0xd07cf37924efa8fca893c372ddfb1614e6fd8c02cfcc9205262535680b77b934.
//
// Solidity: event PositionFeeUpdated(uint256 newPositionFee, uint256 version)
func (_Product *ProductFilterer) WatchPositionFeeUpdated(opts *bind.WatchOpts, sink chan<- *ProductPositionFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "PositionFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductPositionFeeUpdated)
				if err := _Product.contract.UnpackLog(event, "PositionFeeUpdated", log); err != nil {
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

// ParsePositionFeeUpdated is a log parse operation binding the contract event 0xd07cf37924efa8fca893c372ddfb1614e6fd8c02cfcc9205262535680b77b934.
//
// Solidity: event PositionFeeUpdated(uint256 newPositionFee, uint256 version)
func (_Product *ProductFilterer) ParsePositionFeeUpdated(log types.Log) (*ProductPositionFeeUpdated, error) {
	event := new(ProductPositionFeeUpdated)
	if err := _Product.contract.UnpackLog(event, "PositionFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the Product contract.
type ProductSettleIterator struct {
	Event *ProductSettle // Event containing the contract specifics and raw log

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
func (it *ProductSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductSettle)
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
		it.Event = new(ProductSettle)
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
func (it *ProductSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductSettle represents a Settle event raised by the Product contract.
type ProductSettle struct {
	PreVersion *big.Int
	ToVersion  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0x88a84ea6dd274b386afd27dbbe11b6192b25017f5e60bb8c4053dfddb45c294d.
//
// Solidity: event Settle(uint256 preVersion, uint256 toVersion)
func (_Product *ProductFilterer) FilterSettle(opts *bind.FilterOpts) (*ProductSettleIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "Settle")
	if err != nil {
		return nil, err
	}
	return &ProductSettleIterator{contract: _Product.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0x88a84ea6dd274b386afd27dbbe11b6192b25017f5e60bb8c4053dfddb45c294d.
//
// Solidity: event Settle(uint256 preVersion, uint256 toVersion)
func (_Product *ProductFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *ProductSettle) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "Settle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductSettle)
				if err := _Product.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0x88a84ea6dd274b386afd27dbbe11b6192b25017f5e60bb8c4053dfddb45c294d.
//
// Solidity: event Settle(uint256 preVersion, uint256 toVersion)
func (_Product *ProductFilterer) ParseSettle(log types.Log) (*ProductSettle, error) {
	event := new(ProductSettle)
	if err := _Product.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductTakeClosedIterator is returned from FilterTakeClosed and is used to iterate over the raw logs and unpacked data for TakeClosed events raised by the Product contract.
type ProductTakeClosedIterator struct {
	Event *ProductTakeClosed // Event containing the contract specifics and raw log

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
func (it *ProductTakeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductTakeClosed)
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
		it.Event = new(ProductTakeClosed)
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
func (it *ProductTakeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductTakeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductTakeClosed represents a TakeClosed event raised by the Product contract.
type ProductTakeClosed struct {
	Account common.Address
	Version *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTakeClosed is a free log retrieval operation binding the contract event 0x63625b85818a29587ee919ee6a968ee0b32f3513f2884b3968001062ba49eb6b.
//
// Solidity: event TakeClosed(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) FilterTakeClosed(opts *bind.FilterOpts, account []common.Address) (*ProductTakeClosedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.FilterLogs(opts, "TakeClosed", accountRule)
	if err != nil {
		return nil, err
	}
	return &ProductTakeClosedIterator{contract: _Product.contract, event: "TakeClosed", logs: logs, sub: sub}, nil
}

// WatchTakeClosed is a free log subscription operation binding the contract event 0x63625b85818a29587ee919ee6a968ee0b32f3513f2884b3968001062ba49eb6b.
//
// Solidity: event TakeClosed(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) WatchTakeClosed(opts *bind.WatchOpts, sink chan<- *ProductTakeClosed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.WatchLogs(opts, "TakeClosed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductTakeClosed)
				if err := _Product.contract.UnpackLog(event, "TakeClosed", log); err != nil {
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

// ParseTakeClosed is a log parse operation binding the contract event 0x63625b85818a29587ee919ee6a968ee0b32f3513f2884b3968001062ba49eb6b.
//
// Solidity: event TakeClosed(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) ParseTakeClosed(log types.Log) (*ProductTakeClosed, error) {
	event := new(ProductTakeClosed)
	if err := _Product.contract.UnpackLog(event, "TakeClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductTakeOpenedIterator is returned from FilterTakeOpened and is used to iterate over the raw logs and unpacked data for TakeOpened events raised by the Product contract.
type ProductTakeOpenedIterator struct {
	Event *ProductTakeOpened // Event containing the contract specifics and raw log

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
func (it *ProductTakeOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductTakeOpened)
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
		it.Event = new(ProductTakeOpened)
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
func (it *ProductTakeOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductTakeOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductTakeOpened represents a TakeOpened event raised by the Product contract.
type ProductTakeOpened struct {
	Account common.Address
	Version *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTakeOpened is a free log retrieval operation binding the contract event 0xb9726781b72c53f23217f424d70445b222951f008aeac7eece8139caed71ed2d.
//
// Solidity: event TakeOpened(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) FilterTakeOpened(opts *bind.FilterOpts, account []common.Address) (*ProductTakeOpenedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.FilterLogs(opts, "TakeOpened", accountRule)
	if err != nil {
		return nil, err
	}
	return &ProductTakeOpenedIterator{contract: _Product.contract, event: "TakeOpened", logs: logs, sub: sub}, nil
}

// WatchTakeOpened is a free log subscription operation binding the contract event 0xb9726781b72c53f23217f424d70445b222951f008aeac7eece8139caed71ed2d.
//
// Solidity: event TakeOpened(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) WatchTakeOpened(opts *bind.WatchOpts, sink chan<- *ProductTakeOpened, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Product.contract.WatchLogs(opts, "TakeOpened", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductTakeOpened)
				if err := _Product.contract.UnpackLog(event, "TakeOpened", log); err != nil {
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

// ParseTakeOpened is a log parse operation binding the contract event 0xb9726781b72c53f23217f424d70445b222951f008aeac7eece8139caed71ed2d.
//
// Solidity: event TakeOpened(address indexed account, uint256 version, uint256 amount)
func (_Product *ProductFilterer) ParseTakeOpened(log types.Log) (*ProductTakeOpened, error) {
	event := new(ProductTakeOpened)
	if err := _Product.contract.UnpackLog(event, "TakeOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProductTakerFeeUpdatedIterator is returned from FilterTakerFeeUpdated and is used to iterate over the raw logs and unpacked data for TakerFeeUpdated events raised by the Product contract.
type ProductTakerFeeUpdatedIterator struct {
	Event *ProductTakerFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ProductTakerFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductTakerFeeUpdated)
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
		it.Event = new(ProductTakerFeeUpdated)
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
func (it *ProductTakerFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductTakerFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductTakerFeeUpdated represents a TakerFeeUpdated event raised by the Product contract.
type ProductTakerFeeUpdated struct {
	NewTakerFee *big.Int
	Version     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTakerFeeUpdated is a free log retrieval operation binding the contract event 0x8e37539629486a88c6a6083eb90030fb61cb856d56d96f53a326fda0d399ef32.
//
// Solidity: event TakerFeeUpdated(uint256 newTakerFee, uint256 version)
func (_Product *ProductFilterer) FilterTakerFeeUpdated(opts *bind.FilterOpts) (*ProductTakerFeeUpdatedIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "TakerFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ProductTakerFeeUpdatedIterator{contract: _Product.contract, event: "TakerFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchTakerFeeUpdated is a free log subscription operation binding the contract event 0x8e37539629486a88c6a6083eb90030fb61cb856d56d96f53a326fda0d399ef32.
//
// Solidity: event TakerFeeUpdated(uint256 newTakerFee, uint256 version)
func (_Product *ProductFilterer) WatchTakerFeeUpdated(opts *bind.WatchOpts, sink chan<- *ProductTakerFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "TakerFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductTakerFeeUpdated)
				if err := _Product.contract.UnpackLog(event, "TakerFeeUpdated", log); err != nil {
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

// ParseTakerFeeUpdated is a log parse operation binding the contract event 0x8e37539629486a88c6a6083eb90030fb61cb856d56d96f53a326fda0d399ef32.
//
// Solidity: event TakerFeeUpdated(uint256 newTakerFee, uint256 version)
func (_Product *ProductFilterer) ParseTakerFeeUpdated(log types.Log) (*ProductTakerFeeUpdated, error) {
	event := new(ProductTakerFeeUpdated)
	if err := _Product.contract.UnpackLog(event, "TakerFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
