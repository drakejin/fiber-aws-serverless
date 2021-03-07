package db

import (
	"database/sql"
	"fmt"
	baseLog "log"
	"os"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	gormLogger "gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/drakejin/fiber-aws-serverless/config"
	_const "github.com/drakejin/fiber-aws-serverless/const"
	"github.com/drakejin/fiber-aws-serverless/model"
)

type Client struct {
	DB       *gorm.DB
	config   *config.Config
	Migrator *gormigrate.Gormigrate
}

func New(cfg *config.Config) (*Client, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)"+
			"/%s"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.ServiceDB.Username,
		cfg.ServiceDB.Password,
		cfg.ServiceDB.Host,
		cfg.ServiceDB.Port,
		cfg.ServiceDB.Schema,
	)

	var defaultLogger gormLogger.Interface
	if cfg.Debug {
		defaultLogger = gormLogger.New(
			baseLog.New(os.Stdout, "\r\n", baseLog.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Second,     // Slow SQL threshold
				LogLevel:      gormLogger.Info, // Log level
				Colorful:      true,            // Disable color
			},
		)
	} else {
		// or not
		defaultLogger = gormLogger.Default
	}

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:        dsn,
			DriverName: cfg.ServiceDB.Dialect,
		}),
		&gorm.Config{
			Logger: defaultLogger,
		},
	)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Second)

	return &Client{
		DB:     db,
		config: cfg,
	}, nil
}

func (c *Client) InitMigrator() error {

	migrator := gormigrate.New(
		c.DB,
		gormigrate.DefaultOptions,
		nil,
	)
	migrator.InitSchema(func(m *gorm.DB) error {
		return m.AutoMigrate(model.Load()...)
	})
	c.Migrator = migrator
	return c.Migrator.Migrate()
}

func Initialize(cfg *config.DB) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	db, err := sql.Open(cfg.Dialect, dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", cfg.Schema))
	if err != nil {
		return err
	}
	return nil
}

func Clean(env string, cfg *config.DB) error {
	if env != _const.EnvAlpha {
		err := errors.New("Clean command only accept 'alpha' env")
		log.Fatal().Err(err).Send()
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	db, err := sql.Open(cfg.Dialect, dsn)
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s;", cfg.Schema))
	if err != nil {
		return err
	}
	return nil
}
