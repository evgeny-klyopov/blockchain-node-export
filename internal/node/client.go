package node

import (
	"errors"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/config"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	Address       string
	Category      string
	Amount        decimal.Decimal
	Fee           decimal.Decimal
	Confirmations int
	Blocktime     int64
	BlockDatetime string
	Txid          string
	Time          int64
	Datetime      string
	Blockindex    int64
	Blockhash     string
}

type Noder interface {
	GetListSinceBlock(day string, lastBlock *string, debug bool) (*[]Transaction, *string, error)
}

func NewNodeClient(name string, rpc config.RpcConfigurator) (Noder, error) {
	if name == btcNodeType {
		return newBtc(rpc), nil
	}

	return nil, errors.New("not support type node")
}
