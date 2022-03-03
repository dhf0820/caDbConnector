package cadatabase

import (
	"fmt"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

// func TestGetReleaseHistory(t *testing.T) {
// 	db, err := setupTest()
// 	if err != nil {
// 		log.Errorf("setupTest failed: %s", err)
// 	}
// 	defer db.Close()

// 	Convey("Subject: GetReleaseHistory", t, func() {
// 		Convey("With ID: 1603", func() {
// 			histId := 1603
// 			startTime := time.Now()
// 			hist, err := GetReleaseHistory(histId)
// 			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
// 			Convey("Then the ReleaseHistory should be found", func() {
// 				fmt.Printf("History: %s\n", spew.Sdump(hist))
// 				So(err, ShouldBeNil)
// 				So(hist.ID, ShouldEqual, histId)
// 				So(hist.RecipientID, ShouldEqual, 10)
// 			})
// 		})
// 	})
// }

func TestGetReleaseHistoryByIDS(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	Convey("Subject: GetReleaseHistoryByIDS", t, func() {
		Convey("With IDS_ID: 6035758b0c71ef62a2332a63", func() {
			Id := "6035758b0c71ef62a2332a63"
			//startTime := time.Now()
			hist, err := GetReleaseHistoryByIDS(Id)
			//fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			//Convey("Then the ReleaseHistory should be found", func() {
				//fmt.Printf("History: %s\n", spew.Sdump(hist))
				So(err, ShouldBeNil)
				So(hist.IDSID.String, ShouldEqual, Id)
				So(hist.RecipientID, ShouldEqual, 10)
			//})
		})
	})
}


func TestSetDeliveryStatus(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()

	Convey("Subject: SetDeliveryStatus", t, func() {
		Convey("With IDS_ID: 6035758b0c71ef62a2332a63", func() {
			Id := "6035758b0c71ef62a2332a63"
			//Id := "test123"
			startTime := time.Now()
			err := SetDeliveryStatus(1612, Id, "delivered")
			//hist, err := GetReleaseHistoryByIDS(Id)
			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			//Convey("Then the ReleaseHistory should be found", func() {
			//fmt.Printf("History: %s\n", spew.Sdump(hist))
			So(err, ShouldBeNil)
			//So(hist.IDSID.String, ShouldEqual, Id)
			//So(hist.RecipientID, ShouldEqual, 9)
			//})
		})
	})
}

func TestGetReleaseDocuments(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()
	Convey("Subject: GetReleaseDocuments", t, func() {
		Convey("With ID: 1603", func() {

			startTime := time.Now()
			docs, err := GetReleaseDocuments(1603)
			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			So(err, ShouldBeNil)
			So(docs, ShouldNotBeNil)
			//fmt.Printf("Documents Returned: %s\n", spew.Sdump(docs))
		})
	})
}

func TestGetNewReleaseHistories(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()
	Convey("Subject: GetNewReleaseHistories", t, func() {
		Convey("With ID: 261", func() {

			startTime := time.Now()
			hists, err := GetNewReleaseHistories()
			fmt.Printf("Elapse time for Document: %s\n", time.Since(startTime))
			Convey("Then the ReleaseHistory should be found", func() {
				//fmt.Printf("Test History: %s\n", spew.Sdump(hists))
				So(err, ShouldBeNil)
				So(hists, ShouldNotBeNil)
				//fmt.Printf("Pending releases: %s\n", spew.Sdump(hists))
				//for _, hist := range hists {
				//	ReSetDeliveryStatus(hist.ID)
				//}
			})
		})
	})
}

