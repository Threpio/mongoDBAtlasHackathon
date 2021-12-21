package db

const (
	collectionNames = ""
)

func (db *DB) FirstStartDBMigrations() error {
	db.Client.Database(db.DatabaseName).CreateCollection(db.BasicContext, "")
	return nil
}
