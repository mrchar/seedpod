package test

import "github.com/mrchar/seedpod/db"

func DropTables() error {
	defaultDB := db.Default()
	for _, table := range []string{"account", "account_role", "role"} {
		if defaultDB.Migrator().HasTable(table) {
			err := defaultDB.Migrator().DropTable(table)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
