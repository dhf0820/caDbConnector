package cadatabase

import (
	"fmt"

	//"github.com/davecgh/go-spew/spew"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	//. "github.com/logrusorgru/aurora"
	//log "github.com/sirupsen/logrus"
	"os"
)


func DocumentSummaryByVersionId(versionId uint32) (*DocumentSummary, error) {
	var err error
	db, err := CurrentDB()
	ds := DocumentSummary{}
	vers := DocumentVersion{}

	if err = db.Where("id = ?", versionId).Find(&vers).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("Version:23 %d was not found: %s", versionId, err)
		return nil, err
	}
	fmt.Printf("Version: %s\n", spew.Sdump(vers))
	ds.VersionID = versionId
	ds.DocID = uint32(vers.ClinicalDocumentID)
	ds.SubTitle = vers.SubTitle
	ds.Pages = uint(vers.Pages)
	ds.RecvDateTime = vers.RecvDateTime
	ds.ReptDateTime = vers.ReptDateTime
	ds.ImageID = uint(vers.ImageID)
	ds.Description = vers.Description.String
	if vers.ImageType == "" {
		vers.ImageType = "pdf"
	}
	ds.ImageType = vers.ImageType
	ds.UpdatedAt = vers.UpdatedAt
	ds.CreatedAt = vers.CreatedAt
	ds.DeletedAt = vers.DeletedAt
	ds.Repository = vers.Repository
	ds.ImageUrl = fmt.Sprintf("pdf/%d", ds.VersionID)
	doc := ClinicalDocument{}

	if err = db.Where("id = ?", ds.DocID).Find(&doc).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("Document:35 %d was not found: %s", ds.DocID, err)
		return nil, err
	}

	ds.PatientID = uint32(doc.PatientID)
	ds.VisitID = uint32(doc.VisitID)
	ds.Source = doc.Source
	ds.Facility = doc.Facility


	dt := DocumentType{}
	if err = db.Where("id = ?", vers.DocumentTypeID).Find(&dt).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("DocumentType:44 %d was not found: %s", vers.DocumentTypeID, err)
		return nil, err
	}
	ds.Description = dt.Description
	ds.DocType = dt.Code

	return &ds, nil
}

func DocumentByVersionID(versionId uint32) (*Document, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("DocumentByID Database is not open")
	}
	vers := DocumentVersion{}
	//doc := Document{}
	//caDoc := CaDocument{}

	//cdoc := ClinicalDocument{}
	if err := db.Where("id = ?", versionId).Find(&vers).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("Version:29 %d was not found: %s", versionId, err)
		return nil, err
	}
	docId := vers.ClinicalDocumentID

	doc, err := DocumentByID(docId)
	return doc, err
}


//fullDocument returns the set of records that make up the old QC document ClinDoc, Version, Image
func DocumentByID(cdocid uint) (*Document, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("DocumentByID Database is not open")
	}
	doc := Document{}

	cdoc := ClinicalDocument{}
	if err := db.Where("id = ?", cdocid).Find(&cdoc).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("ClinicalDocument %d was not found: %s", cdocid, err)
		return nil, err
	}
	doc.DocId = uint32(cdoc.ID)
	doc.VersionId = uint32(cdoc.CurrentVersionID)
	doc.PatientId = cdoc.PatientID
  doc.VisitId = cdoc.VisitID
	
	// caDoc.ID = cdoc.ID
	// caDoc.VersionId = cdoc.CurrentVersionID

	vers := DocumentVersion{}

	if err := db.Where("id = ?", doc.VersionId).Find(&vers).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("Version:66 %d was not found: %s", doc.VersionId, err)
		return nil, err
	}

	// doc.Vers = &vers
	doc.DateOfService = *vers.ReptDateTime
	doc.Description = vers.Description.String
	//caDoc.DateOfService = *vers.ReptDateTime

	dType := DocumentType{}

	if err := db.Where("id = ?", vers.DocumentTypeID).Find(&dType).Error; gorm.IsRecordNotFoundError(err) {
		//log it
		log.Errorf("DocumentType: %d was not found: %s", vers.DocumentTypeID,  err.Error())
		return nil, err
	}

	doc.DocTypeId = dType.ID
	doc.TypeClass = dType.Code
	imageHost := os.Getenv("IMAGE_HOST")
	if imageHost == "" {
		imageHost = "http://docker1.ihids.com:4567/api/v1/pdf"
	}
	imageUrl := fmt.Sprintf("%s/%d", imageHost, doc.VersionId)	
	//fmt.Printf("imageURL: %s\n", imageUrl)
  doc.ImageUrl = imageUrl	
	docHost := os.Getenv("DOC_HOST")
	if docHost == "" {
		docHost = "http://docker1.ihids.com:4567/api/v1/document_details"
	}
	docUrl := fmt.Sprintf("%s/%d", docHost, doc.DocId)
	//fmt.Printf("docURL: %s\n", docUrl)
  doc.DocUrl = docUrl
	//fmt.Printf("Final Document from database: %s\n", spew.Sdump(doc))
	return &doc, nil
}

//func DeliveryHistoriesByRoiID(roiID int) ([]DeliveryHistories, error) {
//	var err error s
//	db, err := CurrentDB()
//	if err != nil {
//		return nil, fmt.Errorf("DocumentByID Database is not open")
//	}
//	//doc := Document{}
//	//cdoc := ClinicalDocument{}
//	dh := []DeliveryHistories{}
//	if err := db.Where("roi_id = ?", roiID).Find(&dh).Error; gorm.IsRecordNotFoundError(err) {
//		// log it
//		log.Errorf("No DeliveryHistories forund for %d : %s", roiID, err)
//		return nil, err
//	}
//
//
//	//doc.Cdoc = &cdoc
//	//vers := DocumentVersion{}
//	//
//	//if err := db.Where("id = ?", cdoc.CurrentVersionID).Find(&vers).Error; gorm.IsRecordNotFoundError(err) {
//	//	// log it
//	//	log.Errorf("Version %d was not found: %s", doc.Cdoc.CurrentVersionID, err)
//	//	return nil, err
//	//}
//	//doc.Vers = &vers
//	// imag := ArchivedImage{}
//
//	// if err := db.Where("doc_id = ?", cdoc.ID).Find(&imag).Error; gorm.IsRecordNotFoundError(err) {
//	// 	// log it
//	// 	log.Errorf("Version %d was not found: %s", vers.ID, err)
//	// 	return nil, err
//	// }
//	// doc.Imag = &imag
//	return dh, nil
//}