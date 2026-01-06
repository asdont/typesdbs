package main

import (
	"context"
	"log"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/trinodb/trino-go-client/trino"

	"typesdbs/internal/config"
	"typesdbs/internal/types"
)

const (
	driverNamePostgres   = "postgres"
	driverNameClickHouse = "clickhouse"
	driverNameMySQL      = "mysql"
	driverNameMSSQL      = "mssql"
	driverNameOracle     = "godror"
	driverNameTrino      = "trino"
	driverNameSQLite     = "sqlite3"
)

func main() {
	ctx := context.Background()

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	targetDBs := []types.TargetDB{
		{DriverName: driverNamePostgres, DSN: conf.DBs.Postgres.DSN, TableName: conf.DBs.Postgres.Table},
		{DriverName: driverNameClickHouse, DSN: conf.DBs.ClickHouse.DSN, TableName: conf.DBs.ClickHouse.Table},
		{DriverName: driverNameMySQL, DSN: conf.DBs.MySQL.DSN, TableName: conf.DBs.MySQL.Table},
		{DriverName: driverNameMSSQL, DSN: conf.DBs.MSSQL.DSN, TableName: conf.DBs.MSSQL.Table},
		{DriverName: driverNameOracle, DSN: conf.DBs.OracleService.DSN, TableName: conf.DBs.OracleService.Table},
		{DriverName: driverNameOracle, DSN: conf.DBs.OracleSID.DSN, TableName: conf.DBs.OracleSID.Table},
		{DriverName: driverNameTrino, DSN: conf.DBs.Trino.DSN, TableName: conf.DBs.Trino.Table},
		{DriverName: driverNameSQLite, DSN: conf.DBs.SQLite.DSN, TableName: conf.DBs.SQLite.Table},
	}

	if err := types.Run(ctx, targetDBs); err != nil {
		log.Fatal(err)
	}
}
