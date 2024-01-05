package postgres

import (
	"exporia/internal/domain/entity"
	"exporia/platform/postgres/seed"
	"exporia/platform/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase(dsn string) (*gorm.DB, error) {
	db := ConnectToDb(dsn)
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
	)
	if err != nil {
		return nil, err
	}
	seed.UserSeed(db)
	seed.RolSeed(db)

	return db, nil
}
func ConnectToDb(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.Logger.Fatalf("Failed to connect the database %s", err)
	}
	return db
}
