package report

import (
	"errors"
	"fmt"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/config"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/export"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/node"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/params"
	"log"
	"strconv"
	"time"
)

type appReport struct {
	cfg       config.Configurator
	params    params.Params
	lastBlock *string
}

type Reporter interface {
	Run() error
}

func NewReport(p params.Params) Reporter {
	return &appReport{
		params: p,
	}
}

func (a *appReport) process() error {
	var err error
	a.cfg, err = config.NewConfig(a.params)

	if err != nil {
		return err
	}

	nodeClient, err := node.NewNodeClient(a.params.Blockchain, a.cfg.GetRpcConfig())
	if err != nil {
		return err
	}

	transactions, lastBlock, err := nodeClient.GetListSinceBlock(a.params.Day, a.lastBlock, a.cfg.GetDebug())
	if err != nil {
		log.Print(err)
		return nil
	}

	if lastBlock != nil {
		a.lastBlock = lastBlock

		if a.cfg.GetDebug() == true {
			log.Print("nextBlockHash = ", *a.lastBlock)
		}
	}

	if len(*transactions) == 0 {
		err = errors.New("not found transactions")
		log.Print(err)
		return nil
	}

	exportClient, err := export.NewClient(
		fmt.Sprintf("%s_%s", a.params.Blockchain, a.params.Day),
		a.params.OutputType,
		a.params.OutputFileExtension,
		a.params.OutputDirectory,
		a.cfg.GetDatabaseConfig(),
		a.cfg.GetDebug(),
	)

	if err != nil {
		return err
	}

	err = exportClient.Run(*transactions)
	if err != nil {
		return err
	}

	return err
}

func (a *appReport) Run() error {
	var err error
	err = a.process()
	if a.params.Mode != "worker" || err != nil {
		return err
	}

	period, _ := strconv.ParseInt(a.params.WorkerInterval, 10, 64)

	for _ = range time.Tick(time.Second * time.Duration(period)) {
		err = a.process()
		if err != nil {
			return err
		}
	}

	return err
}
