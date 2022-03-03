package src

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	//"github.com/lib/pq"
	//"github.com/davecgh/go-spew/spew"
	//"time"
)

// type ReleaseHistory struct {
// 	ID 						int
// 	RecipientID				int				`gorm:"column:recipient_id"`
// 	PatientID				int				`gorm:"column:patient_id"`
// 	ReleaseDate				*time.Time 		`gorm:"column:release_date"`
// 	Comment 				sql.NullString 	`gorm:"type:varchar()"`
// 	//ReleaseRequestID 		int				`gorm:"column:release_request_id"`
// 	Status 					sql.NullString 	`gorm:"type:varchar()"`
// 	UserID 					int				`gorm:"column:user_id"`
// 	FaxConfirmation			sql.NullString 	`gorm:"column:fax_confirmation"`
// 	DeliveryMethod			sql.NullString 	`gorm:"column:delivery_method"`
// 	DeliveryInformation 	sql.NullString 	`gorm:"column:delivery_information"`
// 	ScannedDocumentID 		sql.NullInt32   `gorm:"column:scanned_document_id"`
// 	VerbalApprovedBy 		sql.NullInt32   `gorm:"column:verbal_approved_by"`
// 	VerbalApprovedDateTime *time.Time		`gorm:"column:verbal_approved_datetime"`
// 	PrinterName 			sql.NullString 	`gorm:"column:printer_name"`
// 	PrinterFacilityName 	sql.NullString 	`gorm:"column:printer_facility_name"`
// 	ROIStatus 				sql.NullString 	`gorm:"column:roi_status"`
// 	//RepresentitiveName 		sql.NullString 	`gorm:"column:representative_name"`
// 	//RepresentativePhoneNumber sql.NullString `gorm:"column:representative_phone_num"`
// 	Redelivered 			sql.NullBool 	`gorm:"column:redelivered"`
// 	RedeliveryInfo 			sql.NullString 	`gorm:"column:redelivery_info"`
// 	RedeliveryCount 		sql.NullInt32	`gorm:"column:redelivery_count"`
// 	//ClonedFrom 				sql.NullInt32   `gorm:"column:cloned_from"`
// 	Source 					sql.NullString 	`gorm:"type:varchar()"`
// 	NoAuthReason 			sql.NullString 	`gorm:"column:no_auth_reason"`
// 	NoAuthApprovedBy 		sql.NullInt32   `gorn:"column:no_auth_approved_by"`
// 	NoAuthApprovedByName	sql.NullString 	`gorm:"column:no_auth_approved_by_name"`
// 	NoAuthApprovedDatetime	*time.Time		`gorm:"column:no_auth_approved_datetime"`
// 	CreatedAt				*time.Time		`gorm:"column:created_at"`
// 	UpdatedAt				*time.Time		`gorm:"column:updated_at"`
// 	IDSID 					sql.NullString 	`gorm:"column:ids_id"`
// 	IDSStatus 				sql.NullString 	`gorm:"column:ids_status"`
// 	DeliveryStatus 			sql.NullString  `gorm:"column:delivery_status"`
// 	PendReason 				string 			`gorm:"column:pend_reason"`
// 	DeviceID 				int				`gorm:"column:device_id"`
// }

type DeliveryInformation struct {
}

//func SetIDS(id string, status string) {
//	rel, err := GetReleaseHistoryByIDSID()
//}
//TODO: Handle a release has been create.

func GetReleaseHistoryByIDS(id string) (*ReleaseHistory, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("GetReleaseHistoryByIDS Database is not open")
	}
	hist := ReleaseHistory{}

	if err := db.Where("ids_id = ?", id).Find(&hist).Error; gorm.IsRecordNotFoundError(err) {
		// TODO: Log find ReleaseHistory Error
		log.Errorf("ReleaseHistoryByIDS %s was not found: %v", id, err)
		return nil, err
	}
	return &hist, nil
}

func ReleaseHistoryById(histid int) (*ReleaseHistory, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("GetReleaseHistory Database is not open")
	}
	hist := ReleaseHistory{}

	if err := db.Where("id = ?", histid).Find(&hist).Error; gorm.IsRecordNotFoundError(err) {
		// TODO: Log find ReleaseHitory Error
		log.Errorf("ReleaseHistory %d was not found: %s", histid, err)
		return nil, err
	}
	return &hist, nil
}

func GetNewReleaseHistories() ([]*ReleaseHistory, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("GetReleaseHistory Database is not open")
	}
	hists := []*ReleaseHistory{}
	nilstr := sql.NullString{}
	nilstr.Valid = false
	submitstr := "submit"
	//var emptyStr string

	//if err := db.Where("status = ?", submitstr).Find(&hists).Error; gorm.IsRecordNotFoundError(err)
	err = db.Where("status = ?", submitstr).Find(&hists).Error
	if err != nil {
		fmt.Printf("get new releases error: %s\n", err.Error())
		return nil, err
	}
	//if err := db.Where("ids_id <> ''").Find(&hists).Error; gorm.IsRecordNotFoundError(err) {

	// 		// TODO: Log find ReleaseHistory Error
	// 	log.Errorf("New ReleaseHistories ready were not found: %s", err)
	// 	return nil, err
	// }
	return hists, nil
}

func SetDeliveryStatus(histid int, ids_id string, status string) error {
	//fmt.Printf("$$$$ Setting DeliveryStatus %d - %s  - %s\n", histid, ids_id, status)
	db, err := CurrentDB()
	if err != nil {
		return err
	}
	//hist := &ReleaseHistory{}
	//if status != "created" {
	//	fmt.Printf(" Status: %s\n", status)
	//	hist, err = GetReleaseHistoryByIDS(ids_id)
	//}
	//if err != nil {

	hist, err := ReleaseHistoryById(histid)
	if err != nil {
		fmt.Printf("History: %d was not found\n", histid)
		return err
	}

	hist.IDSID.String = ids_id
	hist.IDSID.Valid = true

	hist.Status.String = status
	hist.Status.Valid = true
	hist.DeliveryStatus.String = status
	hist.DeliveryStatus.Valid = true
	db.Save(hist)
	//fmt.Printf("$$$ 136- After SetStatus History: %s\n", spew.Sdump(hist))
	return nil
}

func ReSetDeliveryStatus(histid int) error {
	db, err := CurrentDB()
	if err != nil {
		return err
	}
	status := "delivered"
	hist, err := ReleaseHistoryById(histid)
	//hist, err := GetReleaseHistoryByIDS(ids_id)
	if err != nil {
		fmt.Printf("History: %d was not found\n", histid)
		return err
	}
	hist.Status.String = status
	hist.Status.Valid = true
	hist.DeliveryStatus.String = status
	hist.DeliveryStatus.Valid = true
	db.Save(hist)
	//fmt.Printf("157-Found History: %s\n", spew.Sdump(hist))
	return nil
}

func GetReleaseDocuments(relId uint) ([]*Document, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("DocumentByID Database is not open")
	}
	//doc := Document{}
	//cdoc := ClinicalDocument{}
	dhDocs := []*DeliveryHistory{}
	if err := db.Where("roi_id = ?", relId).Find(&dhDocs).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("No DeliveryHistories forund for %d : %s", relId, err)
		return nil, err
	}
	//fmt.Printf("HistoryDocs: %s\n", spew.Sdump(dhDocs))

	docs := []*Document{}
	for _, rd := range dhDocs {
		doc, _ := DocumentByVersionID(rd.DocumentId) //DocumentId is actually VersionId
		//fmt.Printf("Filled By Version: %s\n", spew.Sdump(doc))
		doc.ImageUrl = os.Getenv("IMAGE_HOST")
		if doc.ImageUrl == "" {
			doc.ImageUrl = "http://docker1.ihids.com:4567/api/v1/pdf"
		}
		doc.DocUrl = os.Getenv("DOC_HOST")
		if doc.DocUrl == "" {
			doc.DocUrl = "http://docker1.ihids.com:4567/api/v1/document_details"
		}
		doc.DocUrl = fmt.Sprintf("%s/%d", doc.DocUrl, doc.DocId)
		doc.ImageUrl = fmt.Sprintf("%s/%d", doc.ImageUrl, doc.VersionId)
		doc.RelDocId = rd.ID // deliveryHistory id
		docs = append(docs, doc)
	}
	//fmt.Printf("Documents: %s\n", spew.Sdump(docs))
	return docs, nil
}

func DeliveryHistoriesById(roiID uint) ([]*DeliveryHistory, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("DocumentByID Database is not open")
	}
	//doc := Document{}
	//cdoc := ClinicalDocument{}
	dh := []*DeliveryHistory{}
	if err := db.Where("roi_id = ?", roiID).Find(&dh).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("No DeliveryHistories forund for %d : %s", roiID, err)
		return nil, err
	}

	//doc.Cdoc = &cdoc
	//vers := DocumentVersion{}
	//
	//if err := db.Where("id = ?", cdoc.CurrentVersionID).Find(&vers).Error; gorm.IsRecordNotFoundError(err) {
	//	// log it
	//	log.Errorf("Version %d was not found: %s", doc.Cdoc.CurrentVersionID, err)
	//	return nil, err
	//}
	//doc.Vers = &vers
	// imag := ArchivedImage{}

	// if err := db.Where("doc_id = ?", cdoc.ID).Find(&imag).Error; gorm.IsRecordNotFoundError(err) {
	// 	// log it
	// 	log.Errorf("Version %d was not found: %s", vers.ID, err)
	// 	return nil, err
	// }
	// doc.Imag = &imag
	return dh, nil
}
