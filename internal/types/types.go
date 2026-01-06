package types

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/lib/pq"
)

const defaultTimeout = 5

type TargetDB struct {
	DriverName string
	DSN        string
	TableName  string
}

func Run(
	ctx context.Context,
	targetDBs []TargetDB,
) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout*time.Second)
	defer cancel()

	for _, targetDB := range targetDBs {
		conn, err := sql.Open(targetDB.DriverName, targetDB.DSN)
		if err != nil {
			return fmt.Errorf("sql.Open: %s: %w", targetDB.DriverName, err)
		}

		if err := conn.PingContext(ctx); err != nil {
			return fmt.Errorf("conn.Ping: %s: %w", targetDB.DriverName, err)
		}

		if err := printTypes(ctx, conn, targetDB); err != nil {
			return fmt.Errorf("printTypes: %w", err)
		}
	}

	return nil
}

func printTypes(
	ctx context.Context,
	conn *sql.DB,
	connData TargetDB,
) error {
	//nolint:gosec
	rows, err := conn.QueryContext(ctx, "SELECT * FROM "+pq.QuoteIdentifier(connData.TableName))
	if err != nil {
		return fmt.Errorf("conn.QueryContext: %s: %w", connData.DriverName, err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close: %s: %v", connData.DriverName, err)
		}
	}()

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return fmt.Errorf("rows.ColumnTypes: %s: %w", connData.DriverName, err)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows.Err: %s: %w", connData.DriverName, err)
	}

	//nolint:forbidigo,mnd
	fmt.Printf(
		"\n%s\n   %s\n%s\n",
		strings.Repeat("-", 70), strings.ToUpper(connData.DriverName), strings.Repeat("-", 70),
	)

	for _, colType := range columnTypes {
		if _, err := fmt.Fprintf(os.Stdout,
			"| %-35s | %-35s | %-35s |\n",
			colType.Name(), colType.ScanType(), colType.DatabaseTypeName(),
		); err != nil {
			return fmt.Errorf("fmt.Fprintf: %s: %w", connData.DriverName, err)
		}
	}

	return nil
}
