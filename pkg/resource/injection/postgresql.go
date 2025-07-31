package injection

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/asnur/vocagame-be-interview/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
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
