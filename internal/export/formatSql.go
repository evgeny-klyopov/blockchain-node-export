package export

import (
	"fmt"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/node"
	"strings"
	"time"
)

type formatSql struct {
	insert func(lines []string) string
	add    func(transaction node.Transaction) string
}

func newFormat(databaseType string, tableName string) formatSql {
	insert := map[string]func(lines []string) string{
		"mysql": func(lines []string) string {
			return "INSERT IGNORE INTO " + tableName + " " +
				"(`address`, `category`, `amount`, `fee`, `confirmations`, `blocktime`, `block_datetime`, `tx_id`, `time`, `datetime`) VALUES " +
				"\n" +
				strings.Join(lines, ",\n") + ";"
		},
		"postgres": func(lines []string) string {
			return "INSERT INTO " + tableName + "  " +
				"(address, category, amount, fee, confirmations, blocktime, block_datetime, tx_id, time, datetime, block_index, block_hash) VALUES " +
				"\n" +
				strings.Join(lines, ",\n") + " ON CONFLICT DO NOTHING;"
		},
	}

	f := formatSql{
		insert: insert[databaseType],
		add: func(transaction node.Transaction) string {
			blockDatetime, _ := time.Parse("02.01.2006 15:04:05", transaction.BlockDatetime)
			datetime, _ := time.Parse("02.01.2006 15:04:05", transaction.Datetime)

			return fmt.Sprintf(
				"('%s', '%s', %v, %v, %d, %d, '%s', '%s', %d, '%s', %v,'%s')",
				transaction.Address,
				transaction.Category,
				transaction.Amount,
				transaction.Fee,
				transaction.Confirmations,
				transaction.Blocktime,
				blockDatetime.Format("2006-01-02 15:04:05"),
				transaction.Txid,
				transaction.Time,
				datetime.Format("2006-01-02 15:04:05"),
				transaction.Blockindex,
				transaction.Blockhash,
			)
		},
	}

	return f
}
