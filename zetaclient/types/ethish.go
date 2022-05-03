package types

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeta-chain/zetacore/common"
	"math/big"
)

type ChainETHish struct {
	//TSSSigner          zetaclient.TSSSigner
	MPIABI             abi.ABI
	ChainID            *big.Int
	MPIContractAddress string
	Client             *ethclient.Client
	Name               common.Chain
	Topics             [][]ethcommon.Hash
}

func BytesToEthHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func HashToAddress(h ethcommon.Hash) ethcommon.Address {
	return ethcommon.BytesToAddress(h.Bytes()[12:32])
}