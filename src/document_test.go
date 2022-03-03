package src

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDocumentSummaryByVersionId(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	Convey("Subject: FindDocumentByVersionId", t, func() {
		Convey("With ID: 37710", func() {
			startTime := time.Now()
			doc, err := DocumentSummaryByVersionId(37710)
			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			Convey("Then the patient should be found", func() {
				//fmt.Printf("Document: %s\n", spew.Sdump(doc))
				So(err, ShouldBeNil)
				So(doc, ShouldNotBeNil)
				So(doc.DocID, ShouldEqual, 40442)
				So(doc.VersionID, ShouldEqual, 37710)
				So(doc.ImageUrl, ShouldEqual, "pdf/37710")
				fmt.Printf("Document: %s\n", spew.Sdump(doc))
			})
		})
	})
}

func TestDocumentByVersionId(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	Convey("Subject: FindDocumentByVersionId", t, func() {
		Convey("With ID: 37710", func() {
			startTime := time.Now()
			doc, err := DocumentByVersionID(37710)
			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			Convey("Then the patient should be found", func() {
				//fmt.Printf("Document: %s\n", spew.Sdump(doc))
				So(err, ShouldBeNil)
				So(doc, ShouldNotBeNil)
				//So(doc.ID, ShouldEqual, 40442)
				//So(doc.VersionId, ShouldEqual, 37710)
				//fmt.Printf("Document: %s\n", spew.Sdump(doc))
			})
		})
	})
}
func TestDocumentById(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	Convey("Subject: FindDocumentById", t, func() {
		Convey("With ID: 37710", func() {
			startTime := time.Now()
			doc, err := DocumentByID(37710)
			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			Convey("Then the patient should be found", func() {
				//fmt.Printf("Document: %s\n", spew.Sdump(doc))
				So(err, ShouldBeNil)
				So(doc, ShouldNotBeNil)
				//So(doc.ID, ShouldEqual, 37710)
				// So(doc.Vers.ClinicalDocumentID, ShouldEqual, 92)
				//fmt.Printf("Document: %s\n", spew.Sdump(doc))
				//
				//So(doc.Imag.DocID, ShouldEqual, 92)

			})
		})
	})
}

func TestDeliveryHistoriesByRoiID(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	// 	Convey("DeliveryHistoriesByDocID", t, func() {
	// 		Convey("With ID: 366", func() {
	// 			startTime := time.Now()
	// 			dh, err := DeliveryHistoryById(1604)
	// 			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
	// 			So(err, ShouldBeNil)
	// 			So(len(dh), ShouldEqual, 2)
	// 			So(dh[0].DocumentId, ShouldEqual, 37712)
	// 			So(dh[0].ImageURL, ShouldEqual, "/pdf/37712")
	// 			fmt.Printf("Documents: %s\n", spew.Sdump(dh))
	// 		})
	// 	})
}
