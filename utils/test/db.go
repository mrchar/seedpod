package test

import "github.com/mrchar/seedpod/common/db"

func DropTables() error {
	defaultDB := db.Default()
	for _, table := range []string{
		"account",
		"role",
		"account_role",
		"profile",
		"mobile_phone",
		"email",
		"user",
		"application",
		"application_account",
	} {
		if defaultDB.Migrator().HasTable(table) {
			err := defaultDB.Migrator().DropTable(table)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func AutoMigrate(models ...interface{}) error {
	return db.Default().AutoMigrate(models...)
}
