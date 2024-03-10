package migrations

import (
	"AirAccountEmailAdapter/email/repository"
	"gorm.io/gorm"
)

type Migration20240310 struct {
}

func (m *Migration20240310) Up(db *gorm.DB) error {

	if db.Migrator().HasTable(&repository.Mail{}) {
		return nil
	}

	if err := db.AutoMigrate(&repository.Mail{}); err != nil {
		return err
	}
	return nil
}

func (m *Migration20240310) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&repository.Mail{},
	)
}
