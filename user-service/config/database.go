package config

import (
	"fmt"
	"user-service/database/seeds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/rs/zerolog/log"
)

type PostgresDB struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*PostgresDB, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PsqlDB.User,
		cfg.PsqlDB.Password,
		cfg.PsqlDB.Host,
		cfg.PsqlDB.Port,
		cfg.PsqlDB.DBName)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-1] Failed to connect to database" + cfg.PsqlDB.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-2] Failed to get sql db")
		return nil, err
	}

	seeds.SeedRole(db)

	sqlDB.SetMaxOpenConns(cfg.PsqlDB.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.PsqlDB.DBMaxIdle)

	return &PostgresDB{DB: db}, nil
}
