package export

import (
	"errors"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/config"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/node"
)

type Exporter interface {
	Run([]node.Transaction) error
}

func NewClient(
	name string,
	outputType string,
	outputFileExtension string,
	outputDir string,
	cfg config.DatabaseConfigurator,
	debug bool,
) (Exporter, error) {
	if outputType == fileExportType {
		return newFileExport(name, outputFileExtension, outputDir, cfg.GetDatabaseType(), cfg.GetTableName(), debug)
	} else if outputType == databaseExportType {
		return newDatabaseExport(cfg, debug), nil
	}

	return nil, errors.New("not support type export")
}
