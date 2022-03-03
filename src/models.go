package src

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	//pat "github.com/vsoftcorp/caDbConnector/patient"
	// vis "github.com/vsoftcorp/caDbConnectorvisit"
)

type ArchivedImage struct {
	ID                uint
	DocumentVersionID uint
	Image             []byte
	RepHeader         sql.NullString
	DocID             uint
	MedRecNum         string
	Location          sql.NullString
	URL               sql.NullString
	RepositoryID      sql.NullString `gorm:"column:repository_id"`
	VersionNumber     int
	PatientID         uint
	HpfDocUrls        pq.StringArray `gorm:"type:varchar(100)[]"` //[]string `gorm:"column: hpt_doc_urls"`
	HpfDocHost        string         `gorm:"column:hpf_doc_host"`
	ImageType         string         `gorm:"column:image_type"`
	DeletedAt         *time.Time
}

type CaDocument struct {
	ID            uint //ClinicalDocument ID
	PatientId     uint //
	VersionId     uint
	VisitId       uint
	TypeId        uint
	Description   string
	TypeCode      string
	DateOfService time.Time
}

type ClinicalDocument struct {
	ID                 uint
	PatientID          uint
	VisitID            uint
	CurrentVersionID   uint
	TypeID             uint
	Facility           string
	SendingApp         string
	AssigningAuthority sql.NullString //Not Used
	AccessionNumber    sql.NullString //Not used
	AccessionQualifier sql.NullString // Not used
	DocID              sql.NullInt64  // old QC docID
	DocRef             sql.NullString
	Hidden             bool
	Description        string
	Class              string
	ClassCoding        string
	Category           string
	CategoryCoding     string
	DocumentId         string
	DocumentUrl        string
	Version            string
	VersionId          uint
	//HpfDocURLS         pq.StringArray `gorm:"column:hpf_doc_urls type:varchar(50)[]"` //[]string `gorm:"column:hpf_doc_urls"`
	// string `gorm:"column:hpf_doc_urls"`
	//HpfDocHost sql.NullString
	DUID      string `gorm:"column:duid"` //Current not filled in this will
	Source    string
	OldType   uint //Not used
	UpdatedAt *time.Time
	CreatedAt *time.Time
	DeletedAt *time.Time
}

type DocumentSummary struct {
	VersionID    uint32 //Since everything in CA is based of the version number this is ID
	DocID        uint32
	PatientID    uint32
	VisitID      uint32
	DocType      string
	Description  string
	SubTitle     string
	Pages        uint
	Source       string
	Facility     string
	RecvDateTime *time.Time
	ReptDateTime *time.Time
	ImageID      uint
	ImageType    string
	ImageUrl     string // pdf/versionID
	Repository   string
	UpdatedAt    *time.Time
	CreatedAt    *time.Time
	DeletedAt    *time.Time // When soft delete

}

//DeliveryHistory
type DeliveryHistory struct {
	ID                 uint32     `gorm:"column:id"`
	DocumentId         uint32     `gorm:"column:document_id"`
	PatientId          uint32     `gorm:"column:patient_id"`
	RoiId              uint32     `gorm:"column:roi_id"`
	Source             string     `gorm:"column:source"`
	ImageURL           string     `gorm:"column:image_url"`
	DocDescription     string     `gorm:"column:doc_description"`
	DocTypeId          uint32     `gorm:"column:doc_class_id"`
	DocTypeCode        string     `gorm:"column:`
	DocTypeDescription string     `gorm:"column:doc_class_description"`
	DocReptDatetime    *time.Time `gorm:"column:doc_rept_datetime"`
}

//Document is a combination of the three types that make up an Old QC document
// type Document struct {
// 	Cdoc *ClinicalDocument
// 	Vers *DocumentVersion
// 	Dtype *DocType
// 	Imag *ArchivedImage
// }
type Document struct {
	RelDocId      uint32
	DocId         uint32
	RelId         uint32
	VersionId     uint32 // Version ID used to request the image NOT DOC ID
	PatientId     uint   // ID of the patient to get patient details
	VisitId       uint   // ID of the visit if needed
	DocTypeId     uint   // DocumentType ID
	Description   string
	TypeClass     string
	DateOfService time.Time // Version rept_dateTime
	ImageUrl      string
	DocUrl        string
	ImageType     string
}

type DocumentType struct {
	ID          uint
	Code        string
	Description string
}

type DocumentVersion struct {
	ID                 uint
	ClinicalDocumentID uint
	VersionNumber      int
	AssigningAuthority sql.NullString `gorm:"type:varchar(16)"`
	SendingApp         string
	UserID             sql.NullString
	DocumentTypeID     uint           `gorm:"column:document_type_id"`
	Original           sql.NullString `gorm:"type:varchar(1)"`
	Status             sql.NullString `gorm:"type:varchar(2)"`
	PriorStatus        sql.NullString `gorm:"type:varchar(2)"`
	WorkFlowStatus     sql.NullString `gorm:"type:varchar(2)"`
	Confidentiality    sql.NullString `gorm:"type:varchar(1)"`
	DescriptionID      sql.NullInt64
	Description        sql.NullString `gorm:"type:varchar()"`
	Pages              int
	RecvDateTime       *time.Time `gorm:"column:recv_datetime"`
	ReptDateTime       *time.Time `gorm:"column:rept_datetime"`
	Originator         sql.NullString
	OrigDateTime       *time.Time `gorm:"column:orig_datetime"`
	Transcriptionist   sql.NullString
	TransDateTime      *time.Time `gorm:"column:trans_datetime"`
	Editor             sql.NullString
	EditDateTime       *time.Time `gorm:"column:edit_datetime"`
	EditReason         sql.NullString
	Repository         string
	ImageID            uint

	SubTitle  string `gorm:"column:subtitle"`
	ImageType string `gorm:"column:image_type"`
	OldType   sql.NullInt64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type DocType struct {
	ID          uint
	Code        string
	Description string
}

type Patient struct {
	ID               uint
	Facility         string     `json:"facility"`
	MRN              string     `json:"mrn"`
	SSN              string     `json:"ssn"`
	Name             string     `json:"name"`
	FirstName        string     `json:"first_name"`
	MiddleName       string     `json:"middle_name"`
	LastName         string     `json:"last_name"`
	Sex              string     `json:"Sex"`
	Race             string     `json:"Race"`
	Religion         string     `json:"religion"`
	MaritalStatus    string     `json:"marital_status"`
	BirthDate        *time.Time `json:"birth_date"`
	Address1         string     `json:"address_1"`
	Address2         string     `json:"address_2"`
	Address3         string     `json:"address_3"`
	City             string     `json:"city"`
	State            string     `json:"state"`
	PostalCode       string     `json:"postal_code"`
	HomePhone        string     `json:"home_phone"`
	CellPhone        string     `json:"cell_phone"`
	WorkPhone        string     `json:"work_phone"`
	EmailAddress     string     `json:"email_address"`
	BirthPlace       string     `json:"birth_place"`
	Church           string     `json:"church"`
	EmergencyContact string     `json:"emergency_contact"`
	GlobalID         string     `json:"global_id"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Source           string `json:"source"`
	PatientId        string
}

type Recipient struct {
	ID               int            `gorm:"column:id"`
	ContactFirstName string         `gorm:"type:varchar() column:contact_first_name"`
	ContactLastName  sql.NullString `gorm:"type:varchar() column:contact_last_name"`
	Company          sql.NullString `gorm:"type:varchar()"`
	Address1         sql.NullString `gorm:"type:varchar()"`
	Address2         string         `gorm:"type:varchar()"`
	City             sql.NullString `gorm:"type:varchar()"`
	State            sql.NullString `gorm:"type:varchar()"`
	Zip              sql.NullString `gorm:"type:varchar()"`
	Phone            sql.NullString `gorm:"type:varchar()"`
	Email            sql.NullString `gorm:"type:varchar()"`
	Fax              sql.NullString `gorm:"type:varchar()"`
	ConfirmedDate    *time.Time     `gorm:"column:confirmed_date"`
	ConfirmedBy      sql.NullInt32  `gorm:"column:confirmed_by"`
	ConfirmCode      sql.NullString `gorm:"type:varchar()"`
	UsedForDelivery  bool           `gorm:"column:used_for_delivery"`
	ROIFaxStatus     sql.NullString `gorm:"type:varchar() column:roi_fax_status"`
	ApprovedManually bool           `gorm:"column:approved_manually"`
	IdsID            string         `gorm:"column:ids_id"`
	EmrConfigName    string         `gorm:"column:emr_config_name"` // The name of the config for this emr. Should be same for
	// all physicians in an EMR.
	EmrBrand string `gorm:"column:emr_brand"` // used for survey to find out who has one brand
	EmrModel string `gorm:"column:emr_model"`
}

type ReleaseHistory struct {
	ID          int
	RecipientID int            `gorm:"column:recipient_id"`
	PatientID   int            `gorm:"column:patient_id"`
	ReleaseDate *time.Time     `gorm:"column:release_date"`
	Comment     sql.NullString `gorm:"type:varchar()"`
	//ReleaseRequestID 		int				`gorm:"column:release_request_id"`
	Status                 sql.NullString `gorm:"type:varchar()"`
	UserID                 int            `gorm:"column:user_id"`
	FaxConfirmation        sql.NullString `gorm:"column:fax_confirmation"`
	DeliveryMethod         sql.NullString `gorm:"column:delivery_method"`
	DeliveryInformation    sql.NullString `gorm:"column:delivery_information"`
	ScannedDocumentID      sql.NullInt32  `gorm:"column:scanned_document_id"`
	VerbalApprovedBy       sql.NullInt32  `gorm:"column:verbal_approved_by"`
	VerbalApprovedDateTime *time.Time     `gorm:"column:verbal_approved_datetime"`
	PrinterName            sql.NullString `gorm:"column:printer_name"`
	PrinterFacilityName    sql.NullString `gorm:"column:printer_facility_name"`
	ROIStatus              sql.NullString `gorm:"column:roi_status"`
	//RepresentitiveName 		sql.NullString 	`gorm:"column:representative_name"`
	//RepresentativePhoneNumber sql.NullString `gorm:"column:representative_phone_num"`
	Redelivered     sql.NullBool   `gorm:"column:redelivered"`
	RedeliveryInfo  sql.NullString `gorm:"column:redelivery_info"`
	RedeliveryCount sql.NullInt32  `gorm:"column:redelivery_count"`
	//ClonedFrom 				sql.NullInt32   `gorm:"column:cloned_from"`
	Source                 sql.NullString `gorm:"type:varchar()"`
	NoAuthReason           sql.NullString `gorm:"column:no_auth_reason"`
	NoAuthApprovedBy       sql.NullInt32  `gorn:"column:no_auth_approved_by"`
	NoAuthApprovedByName   sql.NullString `gorm:"column:no_auth_approved_by_name"`
	NoAuthApprovedDatetime *time.Time     `gorm:"column:no_auth_approved_datetime"`
	CreatedAt              *time.Time     `gorm:"column:created_at"`
	UpdatedAt              *time.Time     `gorm:"column:updated_at"`
	IDSID                  sql.NullString `gorm:"column:ids_id"`
	IDSStatus              sql.NullString `gorm:"column:ids_status"`
	DeliveryStatus         sql.NullString `gorm:"column:delivery_status"`
	PendReason             string         `gorm:"column:pend_reason"`
	DeviceID               int            `gorm:"column:device_id"`
}

type Visit struct {
	ID              uint
	PatientID       uint
	VisitNum        string     `json:"visit_num"`
	MRN             string     `json:"mrn"`
	AdmitDate       *time.Time `json:"admit_date"`
	DischargeDate   *time.Time `json:"discharge_date"`
	Facility        string     `json:"facility"`
	Clinic          string     `json:"clinic"`
	PatientType     string     `json:"patient_type"`
	HospitalService string     `json:"hospital_service"`
	PayorCode       string     `json:"payor_code"`
	FinancialClass  string     `json:"financial_class"`
	AdmitSource     string     `json:"admit_source"`
	Comment         string     `json:"comment"`
	Origin          string     `json:"origin"`
}
