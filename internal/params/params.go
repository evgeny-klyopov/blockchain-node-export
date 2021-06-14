package params

import (
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

type Params struct {
	Mode                    string
	Blockchain              string
	Day                     string
	OutputType              string
	OutputDirectory         string
	BlockchainRpcTimeout    string
	BlockchainRpcDsn        string
	BlockchainMockDirectory string
	BlockchainDatabaseDsn   string
	path                    string
	BlockchainUseMock       string
	OutputFileExtension     string
	WorkerInterval          string
	Debug                   string
}

func (p *Params) GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "mode",
			Required:    true,
			Value:       "simple",
			Usage:       "mode in [worker, simple]",
			HasBeenSet:  true,
			Destination: &p.Mode,
		},
		&cli.StringFlag{
			Name:        "worker-interval",
			Required:    false,
			Value:       "600",
			Usage:       "worker interval in seconds",
			HasBeenSet:  true,
			Destination: &p.WorkerInterval,
		},
		&cli.StringFlag{
			Name:        "blockchain",
			Required:    true,
			Value:       "BTC",
			Usage:       "blockchain, support only BTC",
			HasBeenSet:  true,
			Destination: &p.Blockchain,
		},
		&cli.StringFlag{
			Name:        "day",
			Required:    true,
			Value:       time.Now().Format("2006-01-02"),
			HasBeenSet:  true,
			Usage:       "day export (find start block)",
			Destination: &p.Day,
		},
		&cli.StringFlag{
			Name:        "output-type",
			Required:    true,
			Value:       "file",
			HasBeenSet:  true,
			Usage:       "output type in [database, file]",
			Destination: &p.OutputType,
		},
		&cli.StringFlag{
			Name:        "output-file-extension",
			Required:    false,
			Value:       "csv",
			HasBeenSet:  true,
			Usage:       "output file extension in [csv, sql]",
			Destination: &p.OutputFileExtension,
		},
		&cli.StringFlag{
			Name:        "output-directory",
			Required:    false,
			Value:       p.path + "/output",
			HasBeenSet:  true,
			Usage:       "output directory (./output)",
			Destination: &p.OutputDirectory,
		},
		&cli.StringFlag{
			Name:        "blockchain-rpc-dsn",
			Required:    true,
			HasBeenSet:  false,
			Usage:       "blockchain rpc dsn (http://user:password@localhost:1234/wallet/wallet.dat)",
			Destination: &p.BlockchainRpcDsn,
		},
		&cli.StringFlag{
			Name:        "blockchain-rpc-timeout",
			Required:    false,
			Value:       "120",
			HasBeenSet:  true,
			Usage:       "blockchain rpc timeout",
			Destination: &p.BlockchainRpcTimeout,
		},
		&cli.StringFlag{
			Name:        "blockchain-rpc-use-mock",
			Required:    false,
			Value:       "false",
			HasBeenSet:  true,
			Usage:       "blockchain rpc use mock (dev)",
			Destination: &p.BlockchainUseMock,
		},
		&cli.StringFlag{
			Name:        "blockchain-rpc-mock-directory",
			Required:    false,
			Value:       p.path + "/mock",
			HasBeenSet:  true,
			Usage:       "blockchain rpc mock directory",
			Destination: &p.BlockchainMockDirectory,
		},
		&cli.StringFlag{
			Name:        "blockchain-database-dsn",
			Required:    false,
			HasBeenSet:  false,
			Usage:       "blockchain database dsn (postgresql://user:password@localhost:1235/dbName/tableName/?charset=utf8)",
			Destination: &p.BlockchainDatabaseDsn,
		},
		&cli.StringFlag{
			Name:        "debug",
			Required:    false,
			Value:       "false",
			HasBeenSet:  true,
			Usage:       "Debug",
			Destination: &p.Debug,
		},
	}
}

func NewParams() Params {
	path, _ := os.Getwd()
	p := Params{
		path: path,
	}

	return p
}
