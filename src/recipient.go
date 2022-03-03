package src

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

//type Recipient struct{
//	ID 					int 			`gorm:"column:id"`
//	ContactFirstName	sql.NullString 	`gorm:"type:varchar() column:contact_first_name"`
//	ContactLastName 	sql.NullString 	`gorm:"type:varchar() column:contact_last_name"`
//	Company 			sql.NullString 	`gorm:"type:varchar()"`
//	Address1 			sql.NullString 	`gorm:"type:varchar()"`
//	Address2 			sql.NullString 	`gorm:"type:varchar()"`
//	City 				sql.NullString 	`gorm:"type:varchar()"`
//	State 				sql.NullString 	`gorm:"type:varchar()"`
//	Zip 				sql.NullString 	`gorm:"type:varchar()"`
//	Phone 				sql.NullString 	`gorm:"type:varchar()"`
//	Email 				sql.NullString 	`gorm:"type:varchar()"`
//	Fax 				sql.NullString 	`gorm:"type:varchar()"`
//	ConfirmedDate 		*time.Time 		`gorm:"column:confirmed_date"`
//	ConfirmedBy			sql.NullInt32	`gorm:"column:confirmed_by"`
//	ConfirmCode 		sql.NullString 	`gorm:"type:varchar()"`
//	UsedForDelivery 	sql.NullBool	`gorm:"column:used_for_delivery"`
//	ROIFaxStatus 		sql.NullString 	`gorm:"type:varchar() column:roi_fax_status"`
//	ApprovedManually 	sql.NullBool	`gorm:"colume:approved_manually"`
//}

func MigrateRecipient() error {
	db, err := CurrentDB()
	if err != nil {
		return fmt.Errorf("MigrateRecipient Database is not open")
	}

	db.AutoMigrate(Recipient{})
	return nil
}

func GetRecipient(recipId int) (*Recipient, error) {
	var err error
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("GetRecipient Database is not open")
	}
	recip := Recipient{}

	if err := db.Where("id = ?", recipId).Find(&recip).Error; gorm.IsRecordNotFoundError(err) {
		// TODO: Log find ReleseHitory Error
		log.Errorf("Recipient %d was not found: %s", recipId, err)
		return nil, err
	}
	return &recip, nil
}

func SaveRecipient(cr *Recipient) (*Recipient, error) {
	db, err := CurrentDB()
	if err != nil {
		return nil, err
	}
	db.Save(cr)
	return cr, nil
}
