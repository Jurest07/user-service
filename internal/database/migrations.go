package database

import (
	"fmt"

	"github.com/Jurest07/user-service/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg *config.Config) error {
	strConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.SSLMode)
	m, err := migrate.New("file://db/migrations", strConn)
	if err != nil {
		err = fmt.Errorf("Ошибка при подключении к файлу миграций: %w", err)
		return err
	}
	defer func() {
		if _, err := m.Close(); err != nil {
			fmt.Printf("Ошибка при закрытии migrate: %v", err)
		}
	}()
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("Никаких изменений не обнаружено")
		} else {
			err = fmt.Errorf("Ошибка при применении миграций: %w", err)
			return err
		}
	}
	return nil
}
