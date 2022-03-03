package src

import (
	"fmt"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFindPatientByMRN(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	Convey("Subject: Find Patient", t, func() {
		Convey("With mrn 422173", func() {
			pat, err := GetPatientByMRN("422173")
			Convey("Then the patient should be found", func() {
				So(err, ShouldBeNil)
				So(pat.MRN, ShouldEqual, "422173")
			})
		})
	})
}

func TestFindPatientByID(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	Convey("Subject: Find Patient", t, func() {
		Convey("With id: 1", func() {
			pat, err := GetPatientByID(1)
			Convey("Then the patient should be found", func() {
				So(err, ShouldBeNil)
				So(pat.MRN, ShouldEqual, "1681300")
			})
		})
	})
}

func TestAllPatientClinicalDocuments(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()
	//var pat *Patient{}
	//var err error
	Convey("Subject: AllDocuments for Patient", t, func() {
		Convey("For patient with MRN: '422173'", func() {
			pat, err := GetPatientByMRN("422173")
			if err != nil {
				log.Fatalf("PATIENT NOT FOUND")
			}
			startTime := time.Now()
			docs, err := pat.allClinicalDocuments(db)
			fmt.Printf("Elapsed time to fetch patient: %s\n", time.Since(startTime))
			So(err, ShouldBeNil)

			So(len(docs), ShouldBeGreaterThan, 0)
			fmt.Printf("First DocumentID: %d\n", docs[0].ID)
		})
	})
}
