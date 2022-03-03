package src

import (
	"fmt"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRecipient(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()
	MigrateRecipient()
	Convey("Subject: GetRecipient", t, func() {
		Convey("With ID: 11", func() {
			recipId := 11
			startTime := time.Now()
			recip, err := GetRecipient(recipId)
			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			Convey("Then the Recipient should be found", func() {
				//fmt.Printf("Recipient: %s\n", spew.Sdump(recip))
				So(err, ShouldBeNil)
				So(recip.ID, ShouldEqual, recipId)
				So(recip.ContactFirstName, ShouldEqual, "Theresa")
				So(recip.ContactLastName.String, ShouldEqual, "French")
			})
		})
	})
}
