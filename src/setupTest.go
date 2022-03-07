package src

import (
	"github.com/jinzhu/gorm"
	//. "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
)

func SetupTest() (*gorm.DB, error) {
	//pgServer := "db.vertisoft.com"
	pgServer := "chartarchive.chw.edu"
	pgUser := "dhf"
	pgPassword := "Sacj0nhat1"
	pgDatabase := "chartarchive"
	pgFacility := "mmr"

	// pgServer := "docker-2.ihids.com"
	// pgUser := "chartarchive"
	// pgPassword := "Sacj0nhat!"
	// pgDatabase := "chartarchive_dev"
	// pgFacility := "demo"
	SetDbConfig(pgServer, pgUser, pgPassword, pgDatabase, pgFacility)
	_, err := Open()
	//pgURL := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", pgServer, pgUser, pgPassword, pgDatabase)
	//pgValues := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", pgServer, pgUser, pgPassword, pgDatabase)
	//log.Debugf("PGDB: %s", pgURL)
	//db, err := gorm.Open("postgres", pgValues)
	if err != nil {
		log.Fatalf("Could not access the db: %s", err.Error())
	}
	_, err = SetFacility("mmr")
	//searchSchema := fmt.Sprintf("set search_path=%s, public", "demo")
	//db.Exec(searchSchema)
	return CurrentDB()
	//return db, err
}
