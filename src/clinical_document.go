package src

import (
	"fmt"

	"github.com/jinzhu/gorm"

	//. "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
)

//find looks at all non nil fields and creates a query equal to that. if id  is specified, that is the only one used
func ClinicalDocumentByID(id int) (*ClinicalDocument, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("ClinicalDocumentByID Database is not open")
	}
	cdoc := ClinicalDocument{}

//if err = db.Debug().Where("id = ?", id).Find(&cdoc).Error; gorm.IsRecordNotFoundError(err) {
	if err = db.Where("id = ?", id).Find(&cdoc).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("ClinicalDocument %d was not found: %s", cdoc.ID, err)
		return nil, err
	}
	return &cdoc, err
}

// func (vs *ValidateSet) Validate(db *gorm.DB, showOK bool, caFac string) {
// 	searchSchema := fmt.Sprintf("set search_path=%s, public", caFac)
// 	db.Exec(searchSchema)
// 	visit, err := vis.GetVisit(db, vs.VisitNum)
// 	if err != nil {
// 		log.Errorf("Visit: %s  for mrn: %s   err: %s", vs.VisitNum, vs.Pat.MRN, err)
// 		return
// 	}
// 	vs.VisitNum = visit.VisitNum
// 	vs.Vis = visit
// 	vs.Doc.VisitID = visit.ID
// 	ok := vs.ValidateClinical(db, showOK, caFac)
// 	if ok == false {
// 		return //Skip rest
// 	}
// 	ok = vs.ValidateVersion(db, showOK, caFac)
// 	if ok == false {
// 		return /// Skip rest
// 	}
// 	vs.ValidateImage(db, showOK, caFac)
// }

// func (vs *ValidateSet) ValidateClinical(db *gorm.DB, showOK bool, caFac string) bool {
// 	searchSchema := fmt.Sprintf("set search_path=%s, public", caFac)
// 	db.Exec(searchSchema)
// 	origDoc := ClinicalDocument{}
// 	result := "" //fmt.Sprintf("%s  ClinicalDocID: %d  -  %s", Green("Ok"), vs.Doc.ID, vs.Doc.DUID)
// 	log.Debugf("Looking for DUID: %s", vs.Doc.DUID)
// 	if err := db.Where("duid = ? ", vs.Doc.DUID).Find(&origDoc).Error; gorm.IsRecordNotFoundError(err) {
// 		res := db.Create(&vs.Doc)
// 		if res.Error != nil {
// 			aud.LogDocEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Doc.DUID, "CreateFailed", res.Error.Error())
// 			//log.Errorf("MRN: %s UUID: %s failed: %s", vs.Pat.MRN, vs.Pat.UUID, res.Error )
// 			result = fmt.Sprintf("%s  MRN: %s  UUID: %s  - err: %s", Red("CreateFailed"), vs.Pat.MRN, vs.Doc.DUID, err)
// 			fmt.Printf("%s\n", result)
// 			return false
// 		}
// 		aud.LogDocEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Doc.DUID, "added", "")
// 		result = fmt.Sprintf("%s  MRN: %s  ClinicalDocID: %d  -  %s", Red("Added"), vs.Pat.MRN, vs.Doc.ID, vs.Doc.DUID)
// 		fmt.Printf("%s\n", result)
// 	} else {
// 		vs.Doc = &origDoc
// 		//aud.LogDocEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Doc.DUID,"Ok","")
// 		result = fmt.Sprintf("%s  MRN: %s  ClinicalDocID: %d  -  %s", Green("Ok"), vs.Pat.MRN, vs.Doc.ID, vs.Doc.DUID)

// 	}

// 	vs.Vers.ClinicalDocumentID = vs.Doc.ID
// 	// log.Debugf("validate urls: %s  - %d", vs.Doc.HpfDocURLS, vs.Doc.ID)
// 	// log.Debugf("validate old : %s  - %d", origDoc.HpfDocURLS, origDoc.ID)
// 	if showOK {
// 		fmt.Printf("%s\n", result)
// 	}
// 	return true
// }

// func (vs *ValidateSet) ValidateVersion(db *gorm.DB, showOK bool, caFac string) bool {
// 	searchSchema := fmt.Sprintf("set search_path=%s, public", caFac)
// 	db.Exec(searchSchema)
// 	origVers := DocumentVersion{}
// 	var docType *dt.DocumentType
// 	log.Debugf("cDocID: %d  ver docid: %d", vs.Doc.ID, vs.Vers.ClinicalDocumentID)
// 	result := "" //""mt.Sprintf("%s  VersionDocID: %d", Green("Ok"), vs.Doc.ID)
// 	if err := db.Where("clinical_document_id = ? ", vs.Doc.ID).Find(&origVers).Error; gorm.IsRecordNotFoundError(err) {
// 		docType = dt.GetDocumentType(db, vs.Vers.Description, vs.Pat.Facility)
// 		vs.Vers.DocumentTypeID = docType.ID //Create it with the proper Type ID
// 		// if vs.Vers.Description == "FACE SHEET(PREVIOUS)" { // Handle our more descriptive FaceSheets
// 		// 	docType = dt.GetDocumentType(db, "FACE SHEET", caFac)
// 		// } else {
// 		// 	docType = dt.GetDocumentType(db, vs.Vers.Description, caFac)
// 		// }
// 		res := db.Create(&vs.Vers)
// 		result = fmt.Sprintf("%s  MRN: %s  Version for DocID: %d  Version.ID: %d", Red("Added"), vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID)
// 		if res.Error != nil {
// 			aud.LogVersEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID, "CreateFailed", res.Error.Error())
// 			result = fmt.Sprintf("%s  MRN: %s  DocID: %d  Err: %s", Red("FAILED INSERT"), vs.Pat.MRN, vs.Doc.ID, res.Error)
// 			fmt.Printf("%s\n", result)
// 			return false
// 		}
// 		fmt.Printf("Saved new Version: %d\n", vs.Vers.ID)
// 		vs.Doc.TypeID = docType.ID
// 		vs.Doc.CurrentVersionID = vs.Vers.ID
// 		db.Save(&vs.Doc)
// 		aud.LogVersEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID, "added", "")

// 	} else {
// 		//aud.LoVersEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID,"Ok","")
// 		result = fmt.Sprintf("%s  MRN: %s  VersionDocID: %d  Version.ID: %d", Green("Ok"), vs.Pat.MRN, vs.Doc.ID, origVers.ID)
// 		vs.Vers.ID = origVers.ID
// 		if showOK {
// 			fmt.Printf("%s\n", result)
// 		}
// 	}
// 	// if vs.Vers.ID == 0 {
// 	// 	vs.Vers = &origVers
// 	// }

// 	// if vs.Vers.Description == "FACE SHEET(PREVIOUS)" { // Handle our more descriptive FaceSheets
// 	// 	docType = dt.GetDocumentType(db, "FACE SHEET", caFac)
// 	// } else {
// 	// 	docType = dt.GetDocumentType(db, vs.Vers.Description, caFac)
// 	// }
// 	// oldTypeID := vs.Vers.DocumentTypeID
// 	// if vs.Vers.DocumentTypeID != docType.ID { // DocumentType was changed, update it
// 	// 	//vs.Doc.TypeID = docType.ID
// 	// 	vs.Vers.DocumentTypeID = docType.ID
// 	// 	result = fmt.Sprintf("%s  MRN: %s  VersionDocID: %d  Type Changed from %d to Type.ID: %d", Yellow("Update"), vs.Pat.MRN, vs.Doc.ID, oldTypeID, vs.Doc.TypeID)
// 	// 	fmt.Printf("%s\n", result)
// 	// 	db.Save(&vs.Vers)
// 	// 	// comment := fmt.Sprintf("From: %d to %d", vs.Doc.TypeID, docType.ID)
// 	// 	// aud.LogTypeEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, "updated", comment)
// 	// }

// 	// Setup Image with links to parent
// 	vs.Image.DocID = vs.Doc.ID
// 	vs.Image.VersionNumber = vs.Vers.VersionNumber
// 	vs.Image.DocumentVersionID = vs.Vers.ID
// 	vs.Image.PatientID = vs.Pat.ID
// 	vs.Image.MedRecNum = vs.Pat.MRN
// 	//vs.Doc.CurrentVersionID = vs.Vers.ID
// 	if vs.Doc.CurrentVersionID == 0 {
// 		log.Errorf("Doc.Current Version is set to 0 for doc:%d", vs.Doc.ID)
// 	}
// 	//db.Save(&vs.Doc)
// 	// if showOK {
// 	// 	fmt.Printf("%s\n", result)
// 	// }
// 	return true
// }

// func (vs *ValidateSet) ValidateImage(db *gorm.DB, showOK bool, caFac string) {
// 	searchSchema := fmt.Sprintf("set search_path=%s, public", caFac)
// 	db.Exec(searchSchema)
// 	var origImage ArchivedImage
// 	image := &origImage
// 	count := 0
// 	foundImage := false
// 	log.Debugf("ValidateImage  cDocID: %d  verid: %d", vs.Doc.ID, vs.Vers.ID)
// 	log.Debugf("Looking for Image for ClinDocID: %d", vs.Doc.ID)
// 	db.Model(&origImage).Where("doc_id = ?", vs.Doc.ID).Count(&count)
// 	//fmt.Printf("   Found %d images\n\n", count)
// 	switch count {
// 	case 0:
// 		//log.Errorf("No images found for %d", vs.Doc.ID)
// 		image = vs.createImage(db, caFac)
// 		if image == nil {
// 			// Create failed
// 			return
// 		}
// 	case 1:
// 		//log.Errorf("One image found for DocID: %d", vs.Doc.ID)
// 		image = vs.getImage(db, caFac)
// 		foundImage = true
// 	default:
// 		//log.Errorf("Found %d images for DocID: %d", count, vs.Doc.ID)
// 		image = vs.deleteExtraImages(db, caFac) // Deletes newest image
// 	}
// 	if image == nil {
// 		aud.LogImageEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID, 0, "invalid Image", "Image returned nil")
// 		return
// 	}
// 	if foundImage { // the image was not created so validate the URLS
// 		// fmt.Printf("Len vs: %d  Len image: %d\n", len(vs.Image.HpfDocUrls), len(image.HpfDocUrls))
// 		// fmt.Printf("vs: %s\n", vs.Image.HpfDocUrls[0])
// 		// fmt.Printf("im: %s\n", image.HpfDocUrls[0])
// 		//fmt.Printf("foundImage: Setting vs.Image.ID to %d\n", image.ID)
// 		vs.Image.ID = image.ID
// 		//fmt.Printf("Len vs: %d  len img: %d\n", len(vs.Image.HpfDocUrls), len(image.HpfDocUrls))
// 		if len(vs.Image.HpfDocUrls) == len(image.HpfDocUrls) {
// 			for i, url := range vs.Image.HpfDocUrls {
// 				//fmt.Printf("Current index: %d\n", i)
// 				log.Debugf("vs: %s == orig: %s", url, image.HpfDocUrls[i])
// 				if url != image.HpfDocUrls[i] {
// 					//log.Errorf("    Urls do not match")
// 					aud.LogImageEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID, image.ID, "update", "URLS")
// 					result := fmt.Sprintf("%s  MRN: %s  Image for DocID: %d Vers.ID: %d urls do not match", Yellow("Update"), vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID)
// 					image.HpfDocUrls = vs.Image.HpfDocUrls
// 					//fmt.Printf("Saving Image: %s\n", spew.Sdump(vs.Image))
// 					db.Save(&vs.Image)
// 					fmt.Printf("%s\n", result)
// 				}
// 			}
// 		} else {
// 			vs.Image.ID = image.ID
// 			//fmt.Printf("SLength match Saving: %s\n", spew.Sdump(vs.Image))
// 			db.Save(&vs.Image)
// 			aud.LogImageEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID, image.ID, "update", "URLS")

// 			result := fmt.Sprintf("%s  MRN: %s Image for DocID: %d Vers.ID: %d urls do not match", Yellow("Update"), vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID)
// 			fmt.Printf("%s\n", result)
// 		}
// 		//vs.Image = image // incase we need to do more validateion
// 	} // vs.Image has the new image
// 	if showOK {
// 		result := fmt.Sprintf("%s  MRN: %s Image for DocID: %d", Green("Ok"), vs.Pat.MRN, vs.Doc.ID)
// 		fmt.Printf("%s\n", result)
// 	}

// 	// vs.Vers.ImageID = vs.Image.ID
// 	// db.Save(&vs.Vers)
// 	return
// }

// func (vs *ValidateSet) createImage(db *gorm.DB, caFac string) *ArchivedImage {
// 	res := db.Create(&vs.Image)
// 	if res.Error != nil {
// 		aud.LogImageEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID, vs.Image.ID, "create Failed", res.Error.Error())

// 		result := fmt.Sprintf("%s  MRN: %s  DocID: %d  Err: %s", Red("FAILED INSERT"), vs.Pat.MRN, vs.Doc.ID, res.Error.Error())
// 		fmt.Printf("%s\n", result)
// 		return nil
// 	}
// 	aud.LogImageEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID, vs.Image.ID, "added", "")
// 	result := fmt.Sprintf("%s  MRN: %s  Image for DocID: %d", Red("Added"), vs.Pat.MRN, vs.Doc.ID)
// 	fmt.Printf("%s\n", result)
// 	vs.Vers.ImageID = vs.Image.ID // update the new version with this image
// 	res = db.Save(&vs.Vers)
// 	if res.Error != nil {
// 		aud.LogImageEntry(db, caFac, vs.Pat.MRN, vs.Doc.ID, vs.Vers.ID, vs.Image.ID, "failed", res.Error.Error())
// 		//fmt.Printf("error: %s\n", spew.Sdump(vs.Vers))
// 	}

// 	return vs.Image
// }

// func (vs *ValidateSet) deleteExtraImages(db *gorm.DB, caFac string) *ArchivedImage {
// 	images := []ArchivedImage{}
// 	//fmt.Printf("Deleting images for doc_id: %d\n", vs.Doc.ID)
// 	res := db.Where("doc_id = ? ", vs.Doc.ID).Order("id asc").Find(&images)
// 	if res.Error != nil {
// 		// Something is wrong. We found  more than one Image
// 	}
// 	image := images[0] //save the oldest one
// 	//fmt.Printf("   Save olest image: %d\n", image.ID)
// 	vs.Image.ID = image.ID
// 	numImages := len(images)
// 	//fmt.Printf("   Number of images: %d\n", numImages)
// 	for i := 1; i < numImages; i++ {
// 		if images[i].ID == 0 {
// 			continue
// 		}
// 		//fmt.Printf("   Deleting Image %d ID: %d\n", i, images[i].ID)
// 		db.Unscoped().Delete(&images[i])
// 	}
// 	//fmt.Printf("Delete returning Image ID: %d  vs: %d\n", image.ID, vs.Image.ID)
// 	//fmt.Printf("Delete calling get image %d - %d\n", vs.Image.ID, image.ID)
// 	//fmt.Printf("\n\n@@@ Getting image for mrn: %s document: %d\n", vs.Pat.MRN, vs.Doc.ID)
// 	//return vs.getImage(db, caFac)
// 	return &image
// }

// func (vs *ValidateSet) getImage(db *gorm.DB, caFac string) *ArchivedImage {

// 	image := ArchivedImage{}
// 	//fmt.Printf("\n\n@@@ Getting image for mrn: %s document: %d\n", vs.Pat.MRN, vs.Doc.ID)
// 	//spew.Sdump(vs.Doc))
// 	if vs.Doc.ID == 0 {
// 		fmt.Printf("!!! Vs.Doc.ID is 0\n")
// 		return nil
// 	}

// 	if err := db.Where("doc_id = ? ", vs.Doc.ID).Find(&image).Error; gorm.IsRecordNotFoundError(err) {
// 		//log.Errorf("Image for doc_id %d was not found after getting count: %s", vs.Doc.ID, err)
// 		//fmt.Printf("Creating Image\n")
// 		return vs.createImage(db, caFac)
// 	}
// 	//fmt.Printf("getImage returning Image  for MRN: %s:   ID: %d  vs: %d\n", vs.Pat.MRN, image.ID, vs.Image.ID)
// 	return &image
// }
