package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Conf struct {
	DBs Databases `toml:"dbs"`
}

type Databases struct {
	Postgres      Database `toml:"postgres"`
	ClickHouse    Database `toml:"clickhouse"`
	MySQL         Database `toml:"mysql"`
	MSSQL         Database `toml:"mssql"`
	OracleService Database `toml:"oracle_service"`
	OracleSID     Database `toml:"oracle_sid"`
	Trino         Database `toml:"trino"`
	SQLite        Database `toml:"sqlite"`
}

type Database struct {
	DSN   string `toml:"dsn"`
	Table string `toml:"table"`
}

func LoadConfig() (Conf, error) {
	if _, err := os.Stat("config.toml"); err == nil {
		var cfg Conf

		if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
			return Conf{}, fmt.Errorf("toml.DecodeFile: %w", err)
		}

		return cfg, nil
	}

	return loadEnv()
}

func loadEnv() (Conf, error) {
	dbs := Databases{
		Postgres: Database{
			DSN:   os.Getenv("POSTGRES"),
			Table: os.Getenv("POSTGRES_TABLE"),
		},
		ClickHouse: Database{
			DSN:   os.Getenv("CLICKHOUSE"),
			Table: os.Getenv("CLICKHOUSE_TABLE"),
		},
		MySQL: Database{
			DSN:   os.Getenv("MYSQL"),
			Table: os.Getenv("MYSQL_TABLE"),
		},
		MSSQL: Database{
			DSN:   os.Getenv("MSSQL"),
			Table: os.Getenv("MSSQL_TABLE"),
		},
		OracleService: Database{
			DSN:   os.Getenv("ORACLE_SERVICE"),
			Table: os.Getenv("ORACLE_SERVICE_TABLE"),
		},
		OracleSID: Database{
			DSN:   os.Getenv("ORACLE_SID"),
			Table: os.Getenv("ORACLE_SID_TABLE"),
		},
		Trino: Database{
			DSN:   os.Getenv("TRINO"),
			Table: os.Getenv("TRINO_TABLE"),
		},
		SQLite: Database{
			DSN:   os.Getenv("SQLITE"),
			Table: os.Getenv("SQLITE_TABLE"),
		},
	}

	conf := Conf{
		DBs: dbs,
	}

	return conf, nil
}
