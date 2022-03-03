package src

import (
	"encoding/json"
	"strings"

	//"database/sql"
	"fmt"

	"github.com/davecgh/go-spew/spew"

	//"database/sql"
	//"github.com/davecgh/go-spew/spew"
	//"github.com/go-pg/pg/v10"
	//"github.com/go-pg/pg/v10/orm"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	//"github.com/lib/pq"
	//"time"
)

type DbDevice struct {
	ID          int32  `json:"id"`
	RecipientID int    `json:"recipient_id"`
	Method      string `json:"method"`
	Name        string `json:"name"`
	IdsID       string `json:"ids_id"`
	Fields      []byte `json:"fields" gorm:"type:bytea"`
}

type IdsDevice struct {
	ID          int32   `json:"id"`
	RecipientID int     `json:"recipient_id"`
	Method      string  `json:"method"`
	Name        string  `json:"name"`
	IdsID       string  `json:"ids_id"`
	Fields      []Field `json:"fields"`
}

type Field struct {
	Name         string `json:"name"`
	Label        string `json:"label"`
	Default      string `json:"default"`
	Value        string `json:"value"`
	DisplayValue string `json:"display_value"`
	Required     string `json:"required, omitempty"`
	UserVisible  string `json:"user_visible, omitempty"`
	IsNameValue  string `json:"is_name_value"`
	Sensitive    string `json:"sensitive"`
}

type Device struct {
	ID             int32  `json:"id"`
	RecipientID    int    `json:"recipient_id" gorm:"recipient_id"`
	IdsRecipientID string `json:"ids_recipient_id" gorm:"column:ids_recipient_id"`
	Method         string `json:"method"`
	Name           string `json:"name"`
	IdsID          string `json:"ids_id" gorm:"column:ids_id"`
	Fields         []byte `json:"fields" gorm:"type:bytea"`
}

func MigrateDevice() error {
	db, err := CurrentDB()
	if err != nil {
		return fmt.Errorf("GetDeviceByRecipientAndMethod Database is not open")
	}

	db.AutoMigrate(Device{})
	return nil
}

func GetDeviceForRecipient(recipId int, method string) (*Device, error) {
	method = strings.ToLower(method)
	var err error
	var recip *Recipient
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("GetDevicByRecipientAndMethod Database is not open")
	}
	//facility := PgFacility
	device := &Device{}
	if recipId != 0 {
		recip, err = GetRecipient(recipId)
		if err != nil {
			log.Errorf("Recipient: %d was not found", recipId)
			return nil, err
		} else {
			log.Infof("Get %s device for %s", method, recip.Company.String)
		}
	}
	//recipID := fmt.Sprintf("%s-%d", facility, recipId)

	if err := db.Where("recipient_id = ? AND method = ?", recipId, method).Find(device).Error; gorm.IsRecordNotFoundError(err) {
		err = fmt.Errorf("CaDevice for Recipient %d and method: %s was not found: %v", recipId, method, err)
		return nil, err
		//device, err = CreateDeviceForRecipient(recipId, method)
		//if err != nil {
		//	return nil, err
		//}
		//return device, nil
	}

	//fmt.Printf("device:101 Found Device: %s\n", spew.Sdump(device))
	return device, err
}

func GetDeviceForId(devId int) (*Device, error) {
	var err error

	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("GetDevicrForId Database is not open")
	}

	device := &Device{}
	if err := db.Where("id = ?", devId).Find(device).Error; gorm.IsRecordNotFoundError(err) {
		log.Errorf("Device with ID %d was not found: %v", devId, err)
		return nil, err
	}
	//fmt.Printf("118-Found Device: %s\n", spew.Sdump(device))
	return device, err
}

func ToIdsDevice(d *Device) *IdsDevice {
	ids := IdsDevice{}
	json.Unmarshal(d.Fields, &ids.Fields)

	ids.Name = d.Name
	ids.ID = d.ID
	ids.Method = d.Method
	ids.RecipientID = d.RecipientID
	//fmt.Printf("IdsDevice: %s\n", spew.Sdump(ids))
	return &ids
}

func NewField(name, label, value, required string) *Field {
	f := Field{}
	f.Name = name
	f.Label = label
	f.Value = value
	f.Required = required
	return &f
}

func CreateDevice(recipId int, method string) (*Device, error) {
	method = strings.ToLower(method)
	db, err := CurrentDB()
	if err != nil {
		return nil, fmt.Errorf("CreateDevice Database is not open")
	}
	recip, err := GetRecipient(recipId)
	if err != nil {
		log.Errorf("recipient_id: %d is invalid", recipId)
		return nil, err
	}
	//facility := PgFacility
	dev, err := GetDeviceForRecipient(recipId, method)
	if err == nil {
		fmt.Printf("Device: %s for recipient: %d exists using it", method, recipId)
		return dev, nil
	}

	//fmt.Printf("\n--Creating new Device for %s\n", spew.Sdump(recip))
	d := Device{}

	d.RecipientID = recip.ID //fmt.Sprintf("%s-%d", facility, recip.ID)
	//d.Facility 			= facility
	d.Method = method
	d.IdsRecipientID = recip.IdsID

	switch strings.ToLower(d.Method) {
	case "email":
		d.Name = recip.Email.String
		fields := []Field{}
		f := Field{}
		f.Name = "to"
		f.Label = "To"
		f.Value = recip.Email.String
		f.Required = "true"
		f.UserVisible = "true"
		fields = append(fields, f)
		f = Field{}
		f.Name = "priority"
		f.Label = "Priority"
		f.Value = "3"
		f.Required = "false"
		f.UserVisible = "true"
		fields = append(fields, f)
		f = Field{}
		f.Name = "combined"
		f.Label = "Combined"
		f.Value = "true"
		f.Required = "false"
		f.UserVisible = "true"
		fields = append(fields, f)
		d.Fields, _ = json.Marshal(fields)

	case "fax":
		if recip.Fax.String != "" && recip.UsedForDelivery == true {
			fmt.Printf("Fax is used for delivery\n")
			d.Name = recip.Fax.String
			fields := []Field{}
			f := Field{}
			f.Name = "to"
			f.Label = "To"
			f.Value = recip.Fax.String
			f.Required = "true"
			f.UserVisible = "true"
			fields = append(fields, f)
			f = Field{}
			f.Name = "priority"
			f.Label = "Priority"
			f.Value = "3"
			f.Required = "false"
			f.UserVisible = "true"
			fields = append(fields, f)
			f = Field{}
			f.Name = "combined"
			f.Label = "Combined"
			f.Value = "true"
			f.Required = "false"
			f.UserVisible = "true"
			fields = append(fields, f)
			d.Fields, _ = json.Marshal(fields)
		} else {
			err := fmt.Errorf("Fax is not used for delivery")
			return nil, err
		}
	//case "Print":
	//	d.Name = rel.PrinterName.String
	//	fields := []Field{}
	//	f := Field{}
	//	f.Name 			= "to"
	//	f.Label 		= "To"
	//	f.Value 		= rel.PrinterName.String
	//	f.Required		= "true"
	//	f.UserVisible	= "true"
	//	fields = append(fields, f)
	//	f = Field{}
	//	f.Name			= "priority"
	//	f.Label			= "Priority"
	//	f.Value			= "3"
	//	f.Required		= "false"
	//	f.UserVisible	= "true"
	//	fields = append(fields, f)
	//	f = Field{}
	//	f.Name 			= "combined"
	//	f.Label			= "Combined"
	//	f.Value 		= "true"
	//	f.Required		= "false"
	//	f.UserVisible	= "true"
	//	fields = append(fields, f)
	//	f = Field{}
	//	f.Name 			= "printer"
	//	f.Label			= "PrinterHost"
	//	f.Value 		= rel.PrinterFacilityName.String
	//	f.Required		= "false"
	//	f.UserVisible	= "true"
	//	fields = append(fields, f)
	//	d.Fields, _ = json.Marshal(fields)
	//	d.RecipientID = fmt.Sprintf("%s-0",facility )
	//	//d.Address = rel.PrinterName.String
	//	//d.Facility = rel.PrinterFacilityName.String
	//	//d.Name = rel.PrinterName.String
	//case "Download":
	//	d.Name = recip.Company.String
	//	fields := []Field{}
	//	f := Field{}
	//	f.Name 			= "to"
	//	f.Label 		= "To"
	//	f.Value 		= recip.Company.String
	//	f.Required		= "true"
	//	f.UserVisible	= "true"
	//	fields = append(fields, f)
	//	f = Field{}
	//	f.Name 			= "combined"
	//	f.Label			= "Combined"
	//	f.Value 		= "true"
	//	f.Required		= "false"
	//	f.UserVisible	= "true"
	//	fields = append(fields, f)
	//	d.Fields, _ = json.Marshal(fields)
	//	d.RecipientID = fmt.Sprintf("%s-0",facility )
	//	d.Name = "download"
	case "emr":
		d.Name = recip.EmrConfigName
		fields := []Field{}
		f := Field{}
		f.Name = "to"
		f.Label = "To"
		f.Value = recip.Company.String // When we add people to recipients this will be that name
		f.Required = "true"
		f.UserVisible = "true"
		fields = append(fields, f)
		f = Field{}
		f.Name = "priority"
		f.Label = "Priority"
		f.Value = "3"
		f.Required = "false"
		f.UserVisible = "true"
		fields = append(fields, f)
		f = Field{}
		f.Name = "combined"
		f.Label = "Combined"
		f.Value = "false"
		f.Required = "false"
		f.UserVisible = "true"
		fields = append(fields, f)
		d.Fields, _ = json.Marshal(fields)

	}

	result := db.Save(&d)
	err = result.Error
	//fmt.Printf("294-new id : %v\n", result.Value)
	//fmt.Printf("Device: %s\n", spew.Sdump(d))
	return &d, nil
}

//TODO: Possibly confirm the device.IDSRecipientID is set
func SetDeviceIdsId(devId int, idsId string) error {
	fmt.Printf("Setting IDS ID %d - %s \n", devId, idsId)
	db, err := CurrentDB()
	device, err := GetDeviceForId(devId)
	if err != nil {
		fmt.Printf("SetDevice %d IdsId %s  err: %v\n", devId, idsId, err)
		return err
	}
	recip, err := GetRecipient(device.RecipientID)
	if err != nil {
		log.Errorf("recipient %d was not found for Device: %d err: %v", device.RecipientID, device.ID, err)
	}
	device.IdsRecipientID = recip.IdsID
	device.IdsID = idsId
	db.Save(device)
	fmt.Printf("Updated Device: %s\n", spew.Sdump(device))
	return nil
}

func SaveDevice(dev *Device) (*Device, error) {
	db, _ := CurrentDB()
	db.Save(dev)
	return dev, nil
}
