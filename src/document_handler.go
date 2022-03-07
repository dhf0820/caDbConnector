package src

import (
	//"bytes"
	//b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	//"strconv"

	//"time"

	"github.com/davecgh/go-spew/spew"

	log "github.com/sirupsen/logrus"

	//"net/url"
	//"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	//"go.mongodb.org/mongo-driver/bson"
)

//"gopkg.in/mgo.v2/bson"

//"strconv"
//"strings"
//"time"
//"github.com/oleiade/reflections"

type DocumentSummaryResponse struct {
	StatusCode int               	`json:"status_code"`
	Message     string            `json:"message"`
	Count      	int               `json:"count"`
	Document		DocumentSummary		`json:"document"`
	Documents  	[]DocumentSummary `json:"documents"`
}

type ClinicalDocumentResponse struct {
	StatusCode 	int 								`json:"status_code"`
	Message 		string							`json:"message"`
	Document 		ClinicalDocument		`json:"document"`
	Documents 	[]ClinicalDocument	`json:"documents"`
}
// type DocumentImageResponse struct {
// 	StatusCode int             `json:"status_code"`
// 	Status     string          `json:"status"`
// 	Data       m.DocumentImage `json:"data"`
// }

type DocumentResponse struct {
	StatusCode   int              `json:"status_code"`
	Message      string           `json:"message"`
	CacheStatus  string           `json:"cache_status"`
	Total        int64            `json:"total_documents"`
	PagesInCache int64            `json:"pages_in_cache"`
	NumberInPage int64            `json:"num_in_page"`
	Page         int64            `json:"page"`
	Encounter    string           `json:"visit_num"`
	SessionId    string           `json:"session_id"`
	// Documents    []*fhir.Document `json:"documents"`
	// Document     *fhir.Document   `json:"document"`
}

type DocumentFilter struct {
	Skip          int64    `schema:"skip"`
	Page          int64    `schema:"page"`
	Limit         int64    `schema:"limit"`
	SortBy        []string `schema:"sort"`
	Column        string   `schema:"column"`
	Order         string   `schema:"order"`
	Count         string   `schema:"count"`  //Header
	Source string
	//EncounterID string `schema:"visit_num"``
	MRN          string `schema:"mrn"`
	ID           string `schema:"id"`
	DocID        string `schema:"doc_id"`
	ReptDatetime string `schema:"rept_datetime"`
	Category     string `schema:"category"`
	SourceValues []string `schema:"source_values"`
	BeginDate    string   `schema:"begin_date"`
	EndDate      string   `schema:"end_date"`
	TabID        string   `schema:"tab_id"`

}

func WriteDocumentResponse(w http.ResponseWriter, resp *DocumentResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err

	}
	return nil
}

func WriteCADocumentResponse(w http.ResponseWriter, resp *DocumentResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func SessionDocuments(w http.ResponseWriter, r *http.Request) {
}

////////////////////////////////////////  Handlers //////////////////////////////////////////////////

func QueryDocuments(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("#####################################QueryDocuments ################################\n")
	log.Infof("###Raw QueryDocument query: %s", r.URL.RawQuery)
	docFilter, err := SetupDocumentFilter(r)
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Printf("QueryDocuments:114 -- %s\n", spew.Sdump(docFilter))
	// // log.Debugf("QueryDocuments handler")
	// // if r.Header.Get("AUTHORIZATION") == "" {
	// // 	resp := CADocumentResponse{}
	// // 	resp.StatusCode = 400
	// // 	resp.Message = "No Authorization"
	// // 	WriteCADocumentResponse(w, resp)
	// // }

	// // as := m.AuthSession{Token: r.Header.Get("AUTHORIZATION")}
	// // _, err := m.ValidateAuth(as.Token)
	// // if err != nil {
	// // 	err = as.CreateSession()
	// // 	if err != nil {
	// // 		resp := CADocumentResponse{}
	// // 		resp.StatusCode = 400
	// // 		resp.Message = "No Authorization"
	// // 		WriteCADocumentResponse(w, resp)
	// // 	}
	// // }
	// // as.CreateSession()

	// // if err != nil {
	// // 	resp:= CADocumentResponse{}
	// // 	resp.StatusCode = 400
	// // 	resp.Message = "No Authorization"
	// // 	WriteCADocumentResponse(w, resp)
	// // }
	// //docFilter, err := SetupDocumentFilter(r)

	// if err != nil {
	// 	log.Errorf("SetupDocumentFilter:120 -- failed: %s", err.Error())
	// 	// as := m.AuthSession{Token: r.Header.Get("AUTHORIZATION")}
	// 	// as.CreateSession()
	// 	// docFilter, err = SetupDocumentFilter(r)
	// 	// if err != nil {
	// 	resp := DocumentResponse{}
	// 	//resp.CacheStatus = docFilter.Session.GetDocumentStatus()
	// 	resp.StatusCode = 400
	// 	resp.Message = err.Error()
	// 	WriteCADocumentResponse(w, &resp)
	// 	return

	// }
	// if docFilter.PatientGPI == "" {
	// 	if docFilter.PatientID == "" {
	// 		resp := DocumentResponse{}
	// 		//resp.CacheStatus = docFilter.Session.GetDocumentStatus()
	// 		resp.StatusCode = 400
	// 		resp.Message = "no patient specified(patient_gpi)"
	// 		WriteCADocumentResponse(w, &resp)
	// 		return
	// 	}
	// 	docFilter.PatientGPI = docFilter.PatientID
	// }
	// if docFilter.Cache == "stats" { // Retrieve from cache only
	// 	log.Debugf("DocumentHandler:148 -- querying stats")
	// 	cacheStatus, pagesInCache, totalInCache, _ := docFilter.DocumentCacheStats()
	// 	total := totalInCache //pf.CountCachedCaPatients()
	// 	resp := DocumentResponse{}
	// 	resp.CacheStatus = cacheStatus
	// 	resp.Total = total
	// 	resp.PagesInCache = pagesInCache
	// 	resp.Page = docFilter.Page
	// 	resp.StatusCode = 200
	// 	resp.Message = "Ok"
	// 	WriteDocumentResponse(w, &resp)
	// 	return
	// }
	// // if docFilter.Cache == "clear" {
	// // 	//log.Info("Clearing Cache")
	// // 	m.DeleteDocuments(docFilter.PatientGPI)
	// // 	inCache, err := docFilter.DocumentsInCache()
	// // 	if err != nil {
	// // 		resp := DocumentResponse{}
	// // 		//resp.CacheStatus = docFilter.Session.GetDocumentStatus()
	// // 		resp.StatusCode = 400
	// // 		resp.Message = fmt.Sprintf("DocumentsInCache:168 -- err: %s", err.Error())
	// // 		WriteCADocumentResponse(w, &resp)
	// // 		return
	// // 	}
	// // 	log.Infof("QueryDocuments:172 -- %d in cache", inCache)
	// // }
	// if docFilter.Page == 0 {
	// 	docFilter.Page = 1
	// }
	// inCache, err := docFilter.DocumentsInCache()
	// if err != nil {
	// 	resp := DocumentResponse{}
	// 	//resp.CacheStatus = docFilter.Session.GetDocumentStatus()
	// 	resp.StatusCode = 400
	// 	resp.Message = fmt.Sprintf("DocumentsInCache err: %s", err.Error())
	// 	WriteCADocumentResponse(w, &resp)
	// 	return
	// }
	// log.Infof("QueryDocuments:176 -- %d in cache", inCache)
	// //log.Debugf("QueryDocuments:177 -- DocFilter: %s\n", spew.Sdump(docFilter))
	// docFilter.SearchReports()
	// if docFilter.Page == 0 {
	// 	docFilter.Page = 1
	// }
	// startTime := time.Now()
	// fhirDocs, cacheStatus, numberInPage, pagesInCache, totalInCache, err := docFilter.GetFhirDocumentPage()
	// elapsedTime := time.Since(startTime)
	// if err != nil {
	// 	resp := DocumentResponse{}
	// 	//resp.CacheStatus = docFilter.Session.GetDocumentStatus()
	// 	resp.StatusCode = 400
	// 	resp.Message = err.Error()
	// 	WriteCADocumentResponse(w, &resp)
	// 	return
	// }
	// //fmt.Printf("\nFhirDocuments: %s\n\n", spew.Sdump(fhirDocs))
	// if docFilter.ResultFormat == "ca" {
	// 	log.Infof("QueryDocuments:181 -- CacheStatus: %s", cacheStatus)
	// 	log.Infof("QueryDocuments:182 -- TotalInCache: %d", totalInCache)
	// 	log.Infof("QueryDocuments:183 -- PagesInCache: %d", pagesInCache)
	// 	log.Infof("QueryDocuments:187 -- ElapsedTime: %f seconds", elapsedTime.Seconds())
	// 	ca.FhirDocumentsToCA(w, totalInCache, pagesInCache, numberInPage, docFilter.Page, cacheStatus, fhirDocs)
	// 	// //fmt.Printf("\n\n\n   ### Requesting cached page %d\n", docFilter.Page)
	// 	// // fhirDocs, cacheStatus, numberInPage, pagesInCache, totalInCache, err := docFilter.GetFhirDocumentPage()
	// 	// // if err != nil {
	// 	// // 	resp.Message = err.Error()
	// 	// // 	resp.StatusCode = 400
	// 	// // 	WriteCADocumentResponse(w, &resp)
	// 	// // 	return
	// 	// // }
	// 	// fmt.Printf("number of documents in array: %d\n", len(fhirDocs))
	// 	// resp.Message = "Ok"
	// 	// resp.StatusCode = 200
	// 	// resp.CacheStatus = cacheStatus
	// 	// resp.Documents = fhirDocs
	// 	// resp.NumberInPage = numberInPage
	// 	// resp.PagesInCache = pagesInCache
	// 	// resp.Total = totalInCache
	// 	// resp.Documents = fhirDocs
	// 	// WriteCADocumentResponse(w, &resp)
	// 	return
	// }

	//resp := DocumentResponse{}
	// fmt.Printf("fhirDocuments: %s\n", spew.Sdump(fhirDocs))

	// if docFilter.ResultFormat != "" {
	// 	//Do nothing it is set via url
	// } else if r.Header.Get("ResultFormat") != "" {
	// 	docFilter.ResultFormat = r.Header.Get("ResultFormat")
	// } else if r.Header.Get("RESULTFORMAT") != "" {
	// 	docFilter.ResultFormat = r.Header.Get("RESULTFORMAT")
	// } else if docFilter.ResultFormat == "fhir4" {
	// 	docFilter.ResultFormat = "fhir4"
	// } else if docFilter.ResultFormat == "fhir2" {
	// 	docFilter.ResultFormat = "fhir2" // Default to fhir2
	// } else {
	// 	resp := DocumentResponse{}
	// 	resp.Message = "ResultFormat/RESULTFORMAT header is required"
	// 	resp.StatusCode = 400
	// 	WriteCADocumentResponse(w, &resp)
	// 	return
	// }

	// switch docFilter.ResultFormat {
	// case "ca":
	// 	resp := DocumentResponse{}
	// 	if docFilter.Cache == "stats" { // Retrieve from cache only
	// 		log.Debugf("PatientHandler:202 -- querying stats")
	// 		fillStatus, pagesInCache, totalInCache, _ := docFilter.DocumentCacheStats()
	// 		//SetTokenCookie(w, pf.JWTokenStr)
	// 		//tc, err := UpdateTokenCookie(pf.TokenCookie)

	// 		//caResp.SessionId = pf.SessionId
	// 		//caResp.NumberInPage = numberInPage // len(patients)
	// 		total := totalInCache //pf.CountCachedCaPatients()
	// 		resp.CacheStatus =fillStatus
	// 		resp.Total = total
	// 		resp.PagesInCache = pagesInCache
	// 		resp.Page = docFilter.Page
	// 		resp.StatusCode = 200
	// 		resp.Message = "Ok"
	// 		WriteCADocumentResponse(w, &resp)
	// 		return
	// 	}
	// 	// if docFilter.Cache == "reset" || docFilter.Page == 0 { // Reset the cache based upon the cullent filters
	// 	// 	fmt.Printf("\nPatientHandler:219 - Testing reset early\n\n")
	// 	// }

	// 	_, cacheStatus ,inPage, pagesInCache, docTotal, err := docFilter.SearchReports()
	// 	if err != nil {
	// 		resp.Message = err.Error()
	// 		resp.StatusCode = 400
	// 		WriteCADocumentResponse(w, &resp)
	// 		return
	// 	}
	// 	fhirDocs, cacheStatus, inPage, pagesInCache, docTotal, err := docFilter.GetFhirDocumentPage()
	// 	if err != nil {
	// 		resp.Message = err.Error()
	// 		resp.StatusCode = 400
	// 		WriteCADocumentResponse(w, &resp)
	// 		return
	// 	}
	// 	log.Debugf("QueryDocuments:259 -- returned %d documents", len(fhirDocs))
	// 	ca.FhirDocumentsToCA(w, docTotal, pagesInCache, inPage, docFilter.Page, cacheStatus, fhirDocs)
	// 	resp.CacheStatus = cacheStatus
	// 	resp.StatusCode = 200
	// 	resp.Message = "Ok"
	// 	resp.Total = docTotal
	// 	resp.PagesInCache = pagesInCache
	// 	resp.Page = docFilter.Page
	// 	resp.NumberInPage = inPage
	// 	resp.Documents = fhirDocs
	// 	WriteCADocumentResponse(w, &resp)
	// }
	// documents, total, err := documentFilter.Search()

	// if err != nil {
	// 	HandleFhirError("PatientDocuments", w, err)
	// 	return
	// }

	// log.Debugf("Handler has %d documents out of %d\n", len(documents), total)
	//HandleDocumentResponse(w, documentFilter, documents, total)

}

func SetupDocumentFilter(r *http.Request) (*DocumentFilter, error) {

	//config := m.ActiveConfig()
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	docFilter := DocumentFilter{}
	err := decoder.Decode(&docFilter, r.URL.Query())
	if err != nil {
		log.Printf("Filter:238 - parameters :%s\n ", err)
	}

	sessionId := r.Header.Get("SESSION")
	if sessionId == "" {
		log.Errorf("No SESSION Header")
		return nil, errors.New("please log in")
	}

	return &docFilter, nil
}

// func HandleDocumentResponse(w http.ResponseWriter, df *m.DocumentFilter, documents []*m.DocumentSummary, total int64) {
// 	var statusCode int
// 	if df.ResultFormat == "ca" {
// 		//m.FhirDiagDocToCA()
// 		//cadocuments := m.ConvertDocumentsToCA(documents)
// 		// l := len(cadocuments)
// 		// log.Debugf("Convert to ca l = %T, %d\n", l, l)
// 		// var resp CADocumentResponse
// 		// //resp.Total = total
// 		// resp.Page = df.Page
// 		// log.Debugf("Limit = %d\n", df.Limit)
// 		// if df.Limit == 0 {
// 		// 	resp.Pages = 1
// 		// } else {
// 		// 	pages, extra := resp.Total/df.Limit, resp.Total%df.Limit
// 		// 	if extra > 0 {
// 		// 		pages++
// 		// 	}
// 		// 	resp.Pages = pages
// 		// }
// 		// if df.Encounter == "" {
// 		// 	resp.Encounter = "all"
// 		// } else {
// 		// 	resp.Encounter = df.Encounter
// 		// }
// 		// resp.NumberInPage = len(cadocuments)
// 		// resp.Documents = cadocuments

// 		// // log.Debugf("l = %T, %d\n", l, l)
// 		// if len(cadocuments) == 0 {
// 		// 	resp.StatusCode = 404
// 		// 	//resp.Status = fmt.Sprintf("No documents found matching [%s]\n", df.QueryFilter)
// 		// } else {
// 		// 	resp.StatusCode = 200
// 		// }
// 		// WriteCADocumentResponse(w, resp)
// 	} else {
// 		//m.ConvertDocumentsToVS(documents)
// 		if len(documents) == 0 {
// 			statusCode = 404
// 		} else {
// 			statusCode = 200
// 		}
// 		WriteDocumentResponse(w, statusCode, documents)
// 	}
// 	return
// }

type DocumentImageResponse []byte

func WriteDocumentImageResponse(w http.ResponseWriter, statusCode int, data *[]byte) error {
	w.Header().Set("Content-Type", "application/pdf")

	resp := *data

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

//GetDocumentImage returns the image from the url passed in url=
// func GetDocumentImage(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	log.Debugf("GetDocumentImage Params: %v\n", params)
// 	log.Debugf("sid: %v\n", params["id"])
// 	sid := params["id"]

// 	log.Debugf("Handler sid: %s\n", sid)
// 	id, err := strconv.ParseUint(sid, 10, 64)
// 	doc := DocumentSummary{PatientID:uint32(id)}
// 	// err := doc.GetDocumentImage()
// 	// if err != nil {
// 	// 	log.Errorf("GetDocumentImage handler-379: Err: %v\n", err)
// 	// 	HandleFhirError("GetDocumentImage-Handler", w, err)

// 	// 	return
// 	// }

// 	pdfBytes, err := b64.StdEncoding.DecodeString(doc.Image)
// 	if err != nil {
// 		log.Errorf("GetDocumentImage:308 error: %s", err.Error())
// 	}
// 	b := bytes.NewBuffer(pdfBytes)

// 	if _, err := b.WriteTo(w); err != nil {
// 		fmt.Fprintf(w, "%s", err)
// 	}
// }
