package initialize

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	g "gorm.io/gorm"
)

type connector struct {
	ctx context.Context
	sv  string
	env string
}

func NewConnector(ctx context.Context, service, env string) *connector {
	return &connector{ctx: ctx, sv: service, env: env}
}

// postgres
func (c *connector) InitPostgres(config *Postgres) *g.DB {
	uri := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", config.Username, config.Password, config.DB, "5432")

	cli, err := g.Open(postgres.New(postgres.Config{
		DSN:                  uri,
		PreferSimpleProtocol: true,
	}), &config.GormConfig)
	if err != nil {
		log.Fatalf("postgres.Open: %v", err)
		return nil
	}

	err = configureConnectionPool(cli, config.MaxIdle, config.MaxOpen, config.MaxLifetime)
	if err != nil {
		log.Fatalf("configureConnectionPool: %v", err)
		return nil
	}

	log.Println("connect postgres success")
	return cli
}

func configureConnectionPool(cli *g.DB, maxIdle, maxOpen, maxLife string) error {
	sqlDB, err := cli.DB()
	if err != nil {
		log.Fatalf("cli.DB(): %v", err)
		return err
	}

	idle, err := strconv.Atoi(maxIdle)
	if err != nil {
		log.Fatalf("Error: %s environment variable not set.\n", maxIdle)
	}

	open, err := strconv.Atoi(maxOpen)
	if err != nil {
		log.Fatalf("Error: %s environment variable not set.\n", maxOpen)
	}

	life, err := strconv.Atoi(maxLife)
	if err != nil {
		log.Fatalf("Error: %s environment variable not set.\n", maxLife)
	}

	sqlDB.SetMaxIdleConns(idle)
	sqlDB.SetMaxOpenConns(open)
	sqlDB.SetConnMaxLifetime(time.Duration(int64(life)))

	return nil
}
