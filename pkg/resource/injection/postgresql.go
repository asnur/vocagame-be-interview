package injection

import (
	"fmt"

	"github.com/asnur/vocagame-be-interview/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQL struct {
	DB *gorm.DB
}

func NewPostgres(config config.PostgresConfig) (SQL, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUsername,
		config.PostgresPassword,
		config.PostgresDB,
	)

	db, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		return SQL{}, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Ping Connection to ensure it's valid
	sqlDB, err := db.DB()
	if err != nil {
		return SQL{}, fmt.Errorf("failed to get sql.DB from gorm.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return SQL{}, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return SQL{
		DB: db,
	}, nil
}
