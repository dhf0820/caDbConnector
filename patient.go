package cadatabase

import (
	"fmt"
	"github.com/jinzhu/gorm"
	//. "github.com/logrusorgru/aurora"
	log "github.com/sirupsen/logrus"
)

func GetPatientByMRN(mrn string) (*Patient, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("GetPatientByMRN Database is not open")
	}
	pat := Patient{}
	if err := db.Where("mrn = ?", mrn).Find(&pat).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("GetPatient %s was not found: %s", mrn, err)
		return nil, err
	}
	return &pat, nil
}

func GetPatientByID(id uint32) (*Patient, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("GetPatientByMRN Database is not open")
	}
	pat := Patient{}
	if err := db.Where("id = ?", id).Find(&pat).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("GetPatient %d was not found: %v", id, err)
		return nil, err
	}
	return &pat, nil
}


//TODO: return a proper document including Document, Version, Image after type fix

func (p *Patient) allClinicalDocuments(db *gorm.DB) ([]*ClinicalDocument, error) {
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("allClinicalDocuments Database is not open")
	}
	docs := []*ClinicalDocument{}
	if err := db.Where("patient_id = ?", p.ID).Find(&docs).Error; gorm.IsRecordNotFoundError(err) {
		// log it
		log.Errorf("allDocs for %d were not found: %s", p.ID, err)
		return docs, err
	}
	return docs, nil
}

func (patient *Patient) Find(db *gorm.DB) error {

	return nil
}

// func (p *Patient) SplitName() {
// 	//suffix := []string{"SR", "JR", "II", "III"}
// 	names := strings.Split(p.Name, ",")
// 	p.LastName = names[0]
// 	if len(names) > 1 {
// 		gnames := strings.SplitN(strings.TrimSpace(names[1]), " ", 2)
// 		p.FirstName = gnames[0]
// 		if len(gnames) > 1 {
// 			p.MiddleName = gnames[1]
// 		}
// 	}
// 	p.Name = strings.TrimSpace(fmt.Sprintf("%s, %s %s", p.LastName, p.FirstName, p.MiddleName))

// }
