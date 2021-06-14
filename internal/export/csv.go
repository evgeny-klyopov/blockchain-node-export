package export

import (
	"fmt"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/node"
	"strings"
)

type csvFileExport struct {
	fileWriter
}

func newCsvFileExport(name string, outputDir string, outputFileExtension string, debug bool) Exporter {
	header := fmt.Sprintf(
		"%s;%s;%s;%s;%s;%s;%s;%s;%s;%s\n",
		"Address",
		"Category",
		"Amount",
		"Fee",
		"Confirmations",
		"Blocktime",
		"BlockDatetime",
		"Txid",
		"Time",
		"Datetime",
	)
	return &csvFileExport{
		fileWriter{
			name:                name,
			outputDir:           outputDir,
			outputFileExtension: outputFileExtension,
			header:              &header,
			bufferLength:        300,
			writeFormat: func(lines []string) string {
				return strings.Join(lines, "\n")
			},
			format: func(transaction node.Transaction) string {
				return fmt.Sprintf(
					"%s;%s;%v;%v;%d;%d;%s;%s;%d;%s",
					transaction.Address,
					transaction.Category,
					transaction.Amount,
					transaction.Fee,
					transaction.Confirmations,
					transaction.Blocktime,
					transaction.BlockDatetime,
					transaction.Txid,
					transaction.Time,
					transaction.Datetime,
				)
			},
			debug: debug,
		},
	}
}
