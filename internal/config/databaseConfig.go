package config

type databaseConfig struct {
	databaseType string
	user         string
	password     string
	dbname       string
	host         string
	port         string
	charset      string
	tableName    string
}
type DatabaseConfigurator interface {
	GetDatabaseType() string
	GetTableName() string
	GetUser() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetCharset() string
	GetDbname() string
}

func (d *databaseConfig) GetDatabaseType() string {
	return d.databaseType
}

func (d *databaseConfig) GetUser() string {
	return d.user
}

func (d *databaseConfig) GetPassword() string {
	return d.password
}

func (d *databaseConfig) GetHost() string {
	return d.host
}

func (d *databaseConfig) GetPort() string {
	return d.port
}

func (d *databaseConfig) GetCharset() string {
	return d.charset
}

func (d *databaseConfig) GetDbname() string {
	return d.dbname
}

func (d *databaseConfig) GetTableName() string {
	return d.tableName
}
