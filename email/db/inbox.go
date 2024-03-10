package db

const DB_FILE = "mails.db"

// MailInbox represents the recv and verify emails which ready / already process to gateway
func MailInbox() []*Mail {
	panic("Not Implemented")
}

// Save represent save the mail verified
func Save(m *Mail) error {
	panic("Not Implemented")
}
