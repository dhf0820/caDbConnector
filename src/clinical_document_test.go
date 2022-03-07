package src

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClinicalDocumentById(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	Convey("Subject: ClinicalDocumentById", t, func() {
		Convey("With ID: 37710", func() {
			startTime := time.Now()
			doc, err := ClinicalDocumentByID(21177954)
			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			Convey("Then the patient should be found", func() {
				//fmt.Printf("Document: %s\n", spew.Sdump(doc))
				So(err, ShouldBeNil)
				So(doc, ShouldNotBeNil)
				//So(doc.DocID, ShouldEqual, 40442)
				// So(doc.VersionID, ShouldEqual, 37710)
				// So(doc.ImageUrl, ShouldEqual, "pdf/37710")
				fmt.Printf("ClinicalDocument: %s\n", spew.Sdump(doc))
			})
		})
	})
}
