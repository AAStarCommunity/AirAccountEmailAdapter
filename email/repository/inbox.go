package repository

import "AirAccountEmailAdapter/conf"

// MailInbox represents the recv and verify emails which ready / already process to gateway
func MailInbox() []*Mail {
	panic("Not Implemented")
}

// Save represent save the mail verified
func Save(m *Mail) error {
	tx := conf.GetDB()

	err := tx.Omit("updated_at").Create(m).Error

	return err
}
