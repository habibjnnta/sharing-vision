package connection

import (
	"context"
	"database/sql"
	"fmt"
	"sharing-vision/database/ent"
	"sharing-vision/database/ent/migrate"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func ConnectionDB(ctx context.Context, log *logrus.Entry, dbuser string, dbpass string, dbhost string, dbport string, dbname string) *ent.Client {
	dsnDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		dbuser,
		dbpass,
		dbhost,
		dbport,
	)

	db, err := sql.Open("mysql", dsnDB)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer db.Close()

	// Cek apakah database ada
	query := fmt.Sprintf("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '%s'", dbname)
	var exists string
	err = db.QueryRow(query).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			if _, err := db.Exec("CREATE DATABASE " + dbname); err != nil {
				log.Fatalf("failed creating database: %v", err)
				return nil
			}
		} else {
			log.Fatalf("Error checking database: %v", err)
			return nil
		}
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		dbuser,
		dbpass,
		dbhost,
		dbport,
		dbname,
	)

	// Inisialisasi schema Ent
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi Ent: %v", err)
		return nil
	}

	// Migrasi schema ke database
	err = client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				return next.Create(ctx, tables...)
			})
		}),
	)
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi schema: %v", err)
		return nil
	}

	return client
}

func DeleteDatabase(dbuser string, dbpass string, dbhost string, dbport string, dbname string) bool {
	dsnDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		dbuser,
		dbpass,
		dbhost,
		dbport,
	)

	db, err := sql.Open("mysql", dsnDB)
	if err != nil {
		return false
	}
	defer db.Close()

	// Cek apakah database ada
	query := fmt.Sprintf("DROP DATABASE %s;", dbname)
	err = db.QueryRow(query).Err()
	if err != nil {
		return false
	} else {
		return true
	}
}
