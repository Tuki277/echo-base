package database

const (
	FlagDSN         = "DB_DSN"
	FlagDialect     = "DB_DIALECT"
	FlagConLifeTime = "DB_CONNECTION_LIFETIME"
	FlagConIdle     = "DB_CONNECTION_IDLE"
	FlagConMax      = "DB_CONNECTION_MAX"
	FlagPrefix      = "DB_PREFIX"
)

const (
	DialectMySql    = "mysql"
	DialectSqLite   = "sqlite"
	DialectPostgres = "postgres"
	DialectMsSql    = "mssql"
)

const (
	DefaultDialect     = DialectPostgres
	DefaultConIdle     = 10
	DefaultConMax      = 30
	DefaultConLifeTime = 300
	DefaultPrefix      = ""
	DefaultDsn         = ""
)
