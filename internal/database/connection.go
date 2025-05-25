package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *Config) (*Database, error) {
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  config.LogLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	gormConfig := &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	db, err := gorm.Open(postgres.Open(config.GetDSN()), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to connect to the database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	return &Database{db}, nil
}

// Transaction executa uma função dentro de uma transação
func (db *Database) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return db.WithContext(ctx).Transaction(fn)
}

// Execute executa uma operação de banco de dados com retry em caso de erro
func (db *Database) Execute(ctx context.Context, operation func(tx *gorm.DB) error) error {
	maxRetries := 3
	var err error

	for i := 0; i < maxRetries; i++ {
		err = operation(db.WithContext(ctx))
		if err == nil {
			return nil
		}

		if !db.shouldRetry(err) {
			return err
		}

		time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
	}

	return fmt.Errorf("❌ maximum attempts exceeded: %w", err)
}

// shouldRetry verifica se deve tentar novamente baseado no erro
func (db *Database) shouldRetry(_ error) bool {
	// Adicione aqui condições para retry baseadas nos erros
	// Por exemplo, deadlocks, timeouts, etc.
	return false
}

// Batch processa registros em lotes
func (db *Database) Batch(ctx context.Context, batchSize int, model interface{}, fn func(tx *gorm.DB, batch []interface{}) error) error {
	var offset int
	for {
		var batch []interface{}
		result := db.WithContext(ctx).
			Model(model).
			Offset(offset).
			Limit(batchSize).
			Find(&batch)

		if result.Error != nil {
			return result.Error
		}

		if len(batch) == 0 {
			break
		}

		if err := fn(db.DB, batch); err != nil {
			return err
		}

		offset += len(batch)
	}
	return nil
}

// SafeDelete realiza uma deleção segura (soft delete)
func (db *Database) SafeDelete(ctx context.Context, model interface{}, conditions ...interface{}) error {
	return db.WithContext(ctx).Delete(model, conditions...).Error
}

// Health verifica a saúde da conexão com o banco
func (db *Database) Health(ctx context.Context) error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("❌ Failed to get connection: %w", err)
	}
	return sqlDB.PingContext(ctx)
}

// Close fecha a conexão com o banco
func (db *Database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *Database) AutoMigrate(models ...interface{}) error {
	return db.DB.AutoMigrate(models...)
}
