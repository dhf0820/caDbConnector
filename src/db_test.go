package src

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// func TestSetDeviceIdsId(t *testing.T) {
// 	db, err := setupTest()
// 	if err != nil {
// 		log.Errorf("setupTest failed: %s", err)
// 	}

// 	//defer db.Close()
// 	MigrateDevice()
// 	Convey("SetDevice IDS", t, func() {
// 		Convey("6", func() {
// 			So(db, ShouldNotBeNil)
// 			err := SetDeviceIdsId(25, "test_id")
// 			So(err, ShouldBeNil)
// 			device, err := GetDeviceForId(25)
// 			So(err, ShouldBeNil)
// 			So(device.IdsID, ShouldEqual, "test_id")
// 			err = SetDeviceIdsId(25, "")
// 			So(err, ShouldBeNil)
// 		})
// 	})
// }

// func TestGetDeviceByRecipient(t *testing.T) {
// 	db, err := setupTest()
// 	if err != nil {
// 		log.Errorf("setupTest failed: %s", err)
// 	}
// 	defer db.Close()
// 	MigrateDevice()
// 	Convey("GetDeviceForRecipient", t, func() {
// 		Convey("With ReleaseID: 85", func() {
// 			//rel, err := GetReleaseHistory(362)
// 			//So(err, ShouldBeNil)

// 			recipId := 11
// 			//recip, err := GetRecipient(recipId)
// 			//fmt.Printf("Recipient: %s\n", spew.Sdump(recip))
// 			//So(err, ShouldBeNil)
// 			Convey("FindDevice", func() {
// 				method := "Email"
// 				//fmt.Printf("%s\n",spew.Sdump(rel))
// 				d, _ := GetDeviceForRecipient(recipId, method)
// 				//device, err := GetDeviceByRecipientAndMethod(recip, rel.DeliveryMethod.String, "demo", rel)
// 				So(err, ShouldBeNil)
// 				So(d, ShouldNotBeNil)
// 				fmt.Printf("\n\n---Returned: %s\n", spew.Sdump(d))
// 				//d
// 				//fields := []Field{}
// 				//
// 				////fmt.Printf("as []byte: %s\n", spew.Sdump([]byte(d.Fields)))
// 				//json.Unmarshal([]byte(d.Fields), &fields)
// 				//fmt.Printf("Fields: %s\n", spew.Sdump(fields))
// 			})
// 		})
// 	})
// }

// func TestCreateDevice(t *testing.T) {
// 	db, err := setupTest()
// 	if err != nil {
// 		log.Errorf("setupTest failed: %s", err)
// 	}
// 	defer db.Close()
// 	MigrateDevice()
// 	Convey("CreateFaxForRecipient", t, func() {
// 		Convey("With RecipientId = 16 and set for delivery", func() {
// 			//rel, err := GetReleaseHistory(362)
// 			//So(err, ShouldBeNil)

// 			recipId := 16
// 			//recip, err := GetRecipient(recipId)
// 			//fmt.Printf("Recipient: %s\n", spew.Sdump(recip))
// 			//So(err, ShouldBeNil)
// 			Convey("CreateDevice", func() {
// 				method := "Fax"
// 				//fmt.Printf("%s\n",spew.Sdump(rel))
// 				d, err := CreateDevice(recipId, method)
// 				//device, err := GetDeviceByRecipientAndMethod(recip, rel.DeliveryMethod.String, "demo", rel)
// 				So(err, ShouldBeNil)
// 				So(d, ShouldNotBeNil)
// 				//fmt.Printf("\n\n---Returned: %s\n", spew.Sdump(d))
// 				//d
// 				//fields := []Field{}
// 				//
// 				////fmt.Printf("as []byte: %s\n", spew.Sdump([]byte(d.Fields)))
// 				//json.Unmarshal([]byte(d.Fields), &fields)
// 				//fmt.Printf("Fields: %s\n", spew.Sdump(fields))
// 			})
// 		})
// 		Convey("With RecipientId = 11 and not set for delivery", func() {
// 			//rel, err := GetReleaseHistory(362)
// 			//So(err, ShouldBeNil)

// 			recipId := 11
// 			//recip, err := GetRecipient(recipId)
// 			//fmt.Printf("Recipient: %s\n", spew.Sdump(recip))
// 			//So(err, ShouldBeNil)
// 			Convey("CreateDevice", func() {
// 				method := "Fax"
// 				d, err := CreateDevice(recipId, method)
// 				So(err.Error(), ShouldEqual, "Fax is not used for delivery")

// 				So(d, ShouldBeNil)
// 				fmt.Printf("err = %s\n", err.Error())
// 			})
// 		})
// 	})
// }

func TestReleaseById(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()
	MigrateDevice()

	_, err = CurrentDB()
	if err != nil {
		t.Errorf("TesGetRelease Database is not open")
	}
	Convey("GetDevice", t, func() {
		Convey("With ReleaseID: 1603", func() {
			rel, err := ReleaseHistoryById(1612)
			//rel, err := GetReleaseHistory(362)
			So(err, ShouldBeNil)
			So(rel, ShouldNotBeNil)
			fmt.Printf("1612.ids: %s\n", rel.IDSID.String)
			//recipId := 11
			//recip, err := GetRecipient(rel.RecipientID)
			//fmt.Printf("Recipient: %s\n", spew.Sdump(recip))
			//So(err, ShouldBeNil)
			// Convey("Add the Device", func() {
			// 	fmt.Printf("%s\n",spew.Sdump(rel))
			// 	d, _ := GetDevice(rel, PgFacility)
			// 	//device, err := GetDeviceByRecipientAndMethod(recip, rel.DeliveryMethod.String, "demo", rel)
			// 	So(err, ShouldBeNil)
			// 	So(d, ShouldNotBeNil)
			// 	fmt.Printf("\n\n---Returned: %s\n", spew.Sdump(d))
			// 	//d
			// 	//fields := []Field{}
			// 	//
			// 	////fmt.Printf("as []byte: %s\n", spew.Sdump([]byte(d.Fields)))
			// 	//json.Unmarshal([]byte(d.Fields), &fields)
			// 	//fmt.Printf("Fields: %s\n", spew.Sdump(fields))
			// })
		})
	})
}
func TestGetDocument(t *testing.T) {
	db, err := SetupTest()
	if err != nil {
		log.Errorf("setupTest failed: %s", err)
	}
	defer db.Close()
	MigrateDevice()
	// Convey("GetDevice", t, func() {
	// 	Convey("With ReleaseID: 85", func() {
	// 		rel, err := ReleaseByID(db, 362)
	// 		//rel, err := GetReleaseHistory(362)
	// 		So(err, ShouldBeNil)
	// 		So(rel, ShouldNotBeNil)
	// 		fmt.Printf("Release: %s\n", spew.Sdump(rel))
	// 		//recipId := 11
	// 		//recip, err := GetRecipient(rel.RecipientID)
	// 		//fmt.Printf("Recipient: %s\n", spew.Sdump(recip))
	// 		//So(err, ShouldBeNil)
	// 		// Convey("Add the Device", func() {
	// 		// 	fmt.Printf("%s\n",spew.Sdump(rel))
	// 		// 	d, _ := GetDevice(rel, PgFacility)
	// 		// 	//device, err := GetDeviceByRecipientAndMethod(recip, rel.DeliveryMethod.String, "demo", rel)
	// 		// 	So(err, ShouldBeNil)
	// 		// 	So(d, ShouldNotBeNil)
	// 		// 	fmt.Printf("\n\n---Returned: %s\n", spew.Sdump(d))
	// 		// 	//d
	// 		// 	//fields := []Field{}
	// 		// 	//
	// 		// 	////fmt.Printf("as []byte: %s\n", spew.Sdump([]byte(d.Fields)))
	// 		// 	//json.Unmarshal([]byte(d.Fields), &fields)
	// 		// 	//fmt.Printf("Fields: %s\n", spew.Sdump(fields))
	// 		// })
	// 	})
	// })
}

// func TestGetDevice(t *testing.T) {
// 	db, err := setupTest()
// 	if err != nil {
// 		log.Errorf("setupTest failed: %s", err)
// 	}
// 	defer db.Close()
// 	MigrateDevice()
// 	Convey("GetDevice", t, func() {
// 		Convey("With ReleaseID: 85", func() {
// 			rel, err := GetReleaseHistory(362)
// 			So(err, ShouldBeNil)

// 			//recipId := 11
// 			//recip, err := GetRecipient(rel.RecipientID)
// 			//fmt.Printf("Recipient: %s\n", spew.Sdump(recip))
// 			//So(err, ShouldBeNil)
// 			Convey("Add the Device", func() {
// 				fmt.Printf("%s\n",spew.Sdump(rel))
// 				d, _ := GetDevice(rel, PgFacility)
// 				//device, err := GetDeviceByRecipientAndMethod(recip, rel.DeliveryMethod.String, "demo", rel)
// 				So(err, ShouldBeNil)
// 				So(d, ShouldNotBeNil)
// 				fmt.Printf("\n\n---Returned: %s\n", spew.Sdump(d))
// 				//d
// 				//fields := []Field{}
// 				//
// 				////fmt.Printf("as []byte: %s\n", spew.Sdump([]byte(d.Fields)))
// 				//json.Unmarshal([]byte(d.Fields), &fields)
// 				//fmt.Printf("Fields: %s\n", spew.Sdump(fields))
// 			})
// 		})
// 	})
// }
