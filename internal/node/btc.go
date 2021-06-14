package node

import (
	"encoding/json"
	"errors"
	bitcoinRpc "github.com/evgeny-klyopov/bitcoin-rpc"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/config"
	jsonRpcClient "github.com/evgeny-klyopov/golang-json-rpc"
	"log"
	"net/http"
	"time"
)

const btcNodeType = "BTC"

type chainApiResponse struct {
	Blocks  []Block `json:"data"`
	ErrCode int64   `json:"err_code"`
	ErrNo   int64   `json:"err_no"`
	Message string  `json:"message"`
	Status  string  `json:"status"`
}

type Block struct {
	Height           int64        `json:"height"`
	Version          int64        `json:"version"`
	MrklRoot         string       `json:"mrkl_root"`
	Timestamp        int64        `json:"timestamp"`
	Bits             int64        `json:"bits"`
	Nonce            int64        `json:"nonce"`
	Hash             string       `json:"hash"`
	PrevBlockHash    string       `json:"prev_block_hash"`
	NextBlockHash    string       `json:"next_block_hash"`
	Size             int64        `json:"size"`
	PoolDifficulty   float64      `json:"pool_difficulty"`
	Difficulty       int64        `json:"difficulty"`
	DifficultyDouble float64      `json:"difficulty_double"`
	TxCount          int64        `json:"tx_count"`
	RewardBlock      int64        `json:"reward_block"`
	RewardFees       int64        `json:"reward_fees"`
	Confirmations    int64        `json:"confirmations"`
	IsOrphan         bool         `json:"is_orphan"`
	CurrMaxTimestamp int64        `json:"curr_max_timestamp"`
	IsSwBlock        bool         `json:"is_sw_block"`
	StrippedSize     int64        `json:"stripped_size"`
	Sigops           int64        `json:"sigops"`
	Weight           int64        `json:"weight"`
	Extras           ExtrasExtras `json:"extras"`
}

type ExtrasExtras struct {
	PoolName string `json:"pool_name"`
	PoolLink string `json:"pool_link"`
}

type btcNode struct {
	client bitcoinRpc.BitcoinRpc
}

func (b btcNode) GetListSinceBlock(day string, lastBlock *string, debug bool) (*[]Transaction, *string, error) {
	if lastBlock == nil {
		startBlockHash, err := b.getStartBlockHash(day)
		if err != nil {
			return nil, nil, err
		}

		lastBlock = startBlockHash
	}

	if debug == true {
		log.Print("day = ", day)
		log.Print("startBlockHash = ", *lastBlock)
	}

	listSinceBlock, err := b.client.ListSinceBlock([]string{*lastBlock})
	if err != nil {
		return nil, nil, err
	}

	transactions := make([]Transaction, 0, len(listSinceBlock.Transactions))

	lastBlock = &listSinceBlock.Lastblock

	for _, t := range listSinceBlock.Transactions {
		transactions = append(transactions, Transaction{
			Address:       t.Address,
			Category:      t.Category,
			Amount:        t.Amount,
			Fee:           t.Fee,
			Confirmations: t.Confirmations,
			Blocktime:     t.Blocktime,
			Blockindex:    t.Blockindex,
			Blockhash:     t.Blockhash,
			BlockDatetime: time.Unix(t.Blocktime, 0).Format("02.01.2006 15:04:05"),
			Txid:          t.Txid,
			Time:          t.Time,
			Datetime:      time.Unix(t.Time, 0).Format("02.01.2006 15:04:05"),
		})
	}

	if debug == true {
		log.Print("found transactions = ", len(transactions))
	}

	return &transactions, lastBlock, nil
}

func (b *btcNode) getStartBlockHash(day string) (*string, error) {
	dt, err := time.Parse("2006-01-02", day)

	if err != nil {
		return nil, err
	}

	url := "https://chain.api.btc.com/v3/block/date/" + dt.Format("20060102")

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var responseData chainApiResponse
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	if responseData.Status != "success" ||
		responseData.Message != "success" ||
		responseData.ErrCode != 0 ||
		responseData.ErrNo != 0 {
		return nil, errors.New("chain api bad response, last block undefined")
	}

	var minDate *int64
	var hash *string
	for _, block := range responseData.Blocks {
		if minDate == nil {
			minDate = &block.Timestamp
		}
		if block.Timestamp <= *minDate {
			minDate = &block.Timestamp
			hash = &block.Hash
		}
	}

	if hash == nil {
		return nil, errors.New("last block not found, url=[" + url + "]")
	}

	return hash, nil
}

func newBtc(cfg config.RpcConfigurator) Noder {
	node := btcNode{
		client: bitcoinRpc.NewClient(jsonRpcClient.NewClient(jsonRpcClient.ClientCredential{
			Protocol: cfg.GetScheme(),
			User:     cfg.GetUser(),
			Password: cfg.GetPassword(),
			Host:     cfg.GetHost(),
			Port:     cfg.GetPort(),
		}, cfg.GetTimeout(), cfg.GetPath()), cfg.UseMock(), cfg.GetMockDir()),
	}
	return node
}
