package export

import (
	"fmt"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/config"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/node"
	each "github.com/evgeny-klyopov/each"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const databaseExportType = "database"

type databaseExport struct {
	bufferLength int
	cfg          config.DatabaseConfigurator
	db           *sqlx.DB
	debug        bool
	format       formatSql
}

func newDatabaseExport(cfg config.DatabaseConfigurator, debug bool) Exporter {
	return &databaseExport{
		bufferLength: 300,
		cfg:          cfg,
		debug:        debug,
	}
}

func (d *databaseExport) insert(sql string) error {
	_, err := d.db.Exec(sql)

	return err
}
func (d *databaseExport) Run(transactions []node.Transaction) error {
	var err error
	d.db, err = sqlx.Connect(
		d.cfg.GetDatabaseType(),
		fmt.Sprintf(
			"user=%s dbname=%s  password=%s host=%s port=%s sslmode=disable",
			d.cfg.GetUser(),
			d.cfg.GetDbname(),
			d.cfg.GetPassword(),
			d.cfg.GetHost(),
			d.cfg.GetPort(),
		),
	)
	if err != nil {
		return err
	}

	d.format = newFormat(d.cfg.GetDatabaseType(), d.cfg.GetTableName())
	e := each.NewEach(d.bufferLength, d.callback)

	for _, t := range transactions {
		hasError := e.Add(d.format.add(t))
		if true == hasError {
			break
		}
	}
	e.Close()

	return e.GetError()
}

func (d *databaseExport) callback(lines []string, hasError bool) error {
	if hasError == true {
		return nil
	}

	if d.debug == true {
		log.Print("insert database = ", len(lines))
	}

	return d.insert(d.format.insert(lines))
}
