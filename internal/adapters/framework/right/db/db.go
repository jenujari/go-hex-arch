package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	//connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failur: %v", err)
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (da *Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("Db Close error: %v", err)
	}
}

func (da *Adapter) AddToHistory(ans int32, op string) error {
	query, args, err := sq.Insert("arith_history").Columns("date", "answer", "opration").Values(time.Now(), ans, op).ToSql()
	if err != nil {
		return err
	}
	_, er := da.db.Exec(query, args...)
	return er
}
