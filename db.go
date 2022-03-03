package cadatabase

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var (
	_db          *gorm.DB
	PgServer     string
	PgUser       string
	PgPassword   string
	PgDatabase   string
	PgFacility 	 string
	searchSchema string
)

func SetDbConfig(pgServer, pgUser, pgPassword, pgDatabase, pgFacility string) {
	PgServer = pgServer
	PgUser = pgUser
	PgPassword = pgPassword
	PgDatabase = pgDatabase
	PgFacility = pgFacility
	fmt.Printf("cadatabase.setDbConfig: %s:%s facility: %s\n", PgUser, PgPassword, pgFacility)

}

//func Open(pgServer, pgUser, pgPassword, pgDatabase string) (*gorm.DB, error){
func Open() (*gorm.DB, error) {
	var err error
	fmt.Printf("Open CADB host=%s port=5432 user=%s password=%s dbname=%s  facility=%s sslmode=disable",
		PgServer, PgUser, PgPassword, PgDatabase, PgFacility)
	//pgServer := "db.vertisoft.com"
	//pgUser := "dhf"
	//pgPassword := "Sacj0nhat!"
	//pgDatabase := "chartarchive_dev"
	//pgURL :=    fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", pgServer, pgUser, pgPassword, pgDatabase)
	if PgServer != "" && PgUser != "" && PgPassword != "" && PgDatabase != "" {
		pgValues := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
			PgServer, PgUser, PgPassword, PgDatabase)
		log.Debugf("PGDB: %s", pgValues)
		_db, err = gorm.Open("postgres", pgValues)
		if err != nil {
			log.Errorf("Could not access the db: %s", err.Error())
			return nil, err
		}
		searchSchema := fmt.Sprintf("set search_path=%s, public", PgFacility)
		fmt.Printf("Setting inital search path: %s\n", searchSchema)
		_db.Exec(searchSchema)
	}
	return _db, err
}

func SetFacility(facility string) (*gorm.DB, error) {
	if _db == nil {
		return nil, fmt.Errorf("SetFacility Database is not open")
	}
	PgFacility = facility
	searchSchema = fmt.Sprintf("set search_path=%s, public", PgFacility)
	fmt.Printf("Setting new search path: %s\n", searchSchema)
	_db.Exec(searchSchema)
	return _db, nil
}

func CurrentDB() (*gorm.DB, error) {
	if _db == nil {
		if PgServer != "" && PgUser != "" && PgPassword != "" && PgDatabase != "" {
			_, err := Open()
			if err != nil {
				return nil, fmt.Errorf("CurrentDB Database is not open and could not be opened: %v", err)
			}
		} else {
			return nil, fmt.Errorf("Database is not configured")
		}
	}
	return _db, nil
}

//ToNullString invalidates a sql.NullString if empty, validates if not empty
func ToNullString(s string) sql.NullString {
	return sql.NullString{String : s, Valid : s != ""}
}

//ToNullInt64 validates a sql.NullInt64 if incoming string evaluates to an integer, invalidates if it does not
func ToNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64 : int64(i), Valid : err == nil}
}

//ToNullInt32 validates a sql.NullInt64 if incoming string evaluates to an integer, invalidates if it does not
func ToNullInt32(s string) sql.NullInt32 {
	i, err := strconv.Atoi(s)
	return sql.NullInt32{Int32 : int32(i), Valid : err == nil}
}

func MigrateAll() {
	MigrateDevice()
	MigrateRecipient()
}