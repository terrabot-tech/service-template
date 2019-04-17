package provider

import (
	"database/sql"
	"fmt"

	"github.com/terrabot-tech/service-template/models"

	//postgres
	_ "github.com/lib/pq"
)

// Provider interface
type Provider interface {
	Open() error
	GetConn() (*sql.DB, error)
}

type provider struct {
	db       *sql.DB
	cs       string
	poolSize int
}

// NewProvider new provider
func NewProvider(db models.SQLDataBase) Provider {
	info := fmt.Sprintf(`host=%s port=%d dbname=%s user=%s password=%s sslmode=disable application_name=%s`,
		db.Server, db.Port, db.Database, "postgres", "postgres", db.ApplicationName)
	return &provider{
		cs:       info,
		poolSize: db.PoolSize,
	}
}

// Open connection
func (p *provider) Open() error {
	var err error
	p.db, err = sql.Open("postgres", p.cs)
	if err != nil {
		return err
	}
	//p.db.SetMaxIdleConns(p.poolSize)
	//p.db.SetMaxOpenConns(p.poolSize)
	return nil
}

// GetConn return sql database
func (p *provider) GetConn() (*sql.DB, error) {
	return p.db, nil
}
