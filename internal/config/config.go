package config

import (
	"github.com/evgeny-klyopov/blockchain-node-export/internal/params"
	"net"
	"net/url"
	"regexp"
	"strconv"
)

type config struct {
	rpc      RpcConfigurator
	database DatabaseConfigurator
	debug    bool
}
type Configurator interface {
	GetRpcConfig() RpcConfigurator
	GetDatabaseConfig() DatabaseConfigurator
	GetDebug() bool
}

func NewConfig(p params.Params) (Configurator, error) {
	blockchainRpc, err := url.Parse(p.BlockchainRpcDsn)
	if err != nil {
		return nil, err
	}
	blockchainRpcHost, blockchainRpcPort, err := net.SplitHostPort(blockchainRpc.Host)
	if err != nil {
		return nil, err
	}

	database, err := url.Parse(p.BlockchainDatabaseDsn)
	if err != nil {
		return nil, err
	}

	databaseHost, databasePort, err := net.SplitHostPort(database.Host)
	if err != nil {
		return nil, err
	}

	databasePassword, _ := database.User.Password()
	rpcPassword, _ := blockchainRpc.User.Password()

	charset, _ := url.ParseQuery(database.RawQuery)
	match := regexp.MustCompile(`(?m)([^\/]+)`).FindAllString(database.Path, -1)

	dbName := match[0]
	tableName := match[1]

	debug, _ := strconv.ParseBool(p.Debug)
	c := config{
		debug: debug,
		rpc: &rpcConfig{
			scheme:   blockchainRpc.Scheme,
			user:     blockchainRpc.User.Username(),
			password: rpcPassword,
			host:     blockchainRpcHost,
			port:     blockchainRpcPort,
			timeout:  p.BlockchainRpcTimeout,
			path:     blockchainRpc.Path,
			useMock:  p.BlockchainUseMock,
			mockDir:  p.BlockchainMockDirectory,
		},

		database: &databaseConfig{
			databaseType: database.Scheme,
			user:         database.User.Username(),
			password:     databasePassword,
			dbname:       dbName,
			host:         databaseHost,
			port:         databasePort,
			tableName:    tableName,
			charset:      charset.Get("charset"),
		},
	}
	return &c, nil
}

func (c *config) GetRpcConfig() RpcConfigurator {
	return c.rpc
}

func (c *config) GetDatabaseConfig() DatabaseConfigurator {
	return c.database
}
func (c *config) GetDebug() bool {
	return c.debug
}
