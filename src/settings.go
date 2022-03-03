package src

import (
	//"fmt"
	//"os"

	//"github.com/spf13/viper"
	//"strconv"
	//"gitlab.com/dhf0820/roi-chartarchive/internal/messaging"
	//log "github.com/sirupsen/logrus"
	//"github.com/prometheus/common/log"
)

var SettingsInitialized string

// type config struct {
// 	DeliveryPort     string
// 	DeliveryAddress  string
// 	DeliveryTLSMode  string
// 	DeliveryCertName string

// 	//MessagingClient 	*messaging.MessagingClient

// 	InQueueName        string
// 	VsAMQP             string
// 	UserId             string
// 	UserPassword       string
// 	AppName            string
// 	AppUserName        string
// 	ServerUser         string
// 	AppPassword        string
// 	URL                string
// 	Initialized        string
// 	AllScriptsCertName string
// 	Port               string
// }

// Create a new config instance.
// var (
// 	Conf *config
// )

// // Init does nothing right now. It will fill values from a yml file or config server.
// func ChartArchiveConverterSettingsInit() {
// 	//Config = fillConfig()
// 	fmt.Printf("ChartArchiveConverterSettingsInit called\n")
// 	if Conf == nil {
// 		fmt.Printf("\nInitializing AllScripts Delivery Server setting\n")
// 		//dir, _ := os.Getwd()
// 		//fmt.Printf("Currrent directory: %s\n", dir)
// 		Conf = getConf()
// 		//atlas := fmt.Sprintf(Conf.Atlas, "dhf", "Sacj0nhat1")
// 		//fmt.Printf("atlas: %s\n", atlas)
// 		//
// 		//local := fmt.Sprintf(Conf.Mongodb, "dhf", "SacjOnhat1")
// 		//fmt.Printf("local: %s\n", local)
// 		//Conf.MessagingClient = &messaging.MessagingClient{}
// 		//Conf.MessagingClient.ConnectToBroker(Conf.VsAMQP)
// 		//fmt.Printf("caDbConnector CONFIG: %s\n", spew.Sdump(Conf))
// 	} else {
// 		//fmt.Printf("caDbConnector Is configured: %s\n", spew.Sdump(Conf))
// 	}
// }

//func Init() {
//	//Config = fillConfig()
//	fmt.Printf("Intializing setting\n")
//	Conf = getConf()
//	//atlas := fmt.Sprintf(Conf.Atlas, "dhf", "Sacj0nhat1")
//	//fmt.Printf("atlas: %s\n", atlas)
//	//
//	//local := fmt.Sprintf(Conf.Mongodb, "dhf", "SacjOnhat1")
//	//fmt.Printf("local: %s\n", local)
//	fmt.Printf("CONFIG: %s\n", spew.Sdump(Conf))
//}


// func Config() *config {
// 	return Conf
// }

// var (
// 	V *viper.Viper
// )

// // Read the config file from the current directory and marshal
// // into the conf config struct.
// func getConf() *config {
// 	V = viper.New()
// 	V.SetEnvPrefix("roi")
// 	V.AutomaticEnv()

// 	V.AddConfigPath(os.Getenv("ENV_ROI"))
// 	V.AddConfigPath("./config")
// 	V.AddConfigPath("../../../config")
// 	V.SetConfigName("config.yml")
// 	V.SetConfigType("yml")
// 	fmt.Printf("Starting DocServer read config\n")
// 	err := V.ReadInConfig()
// 	fmt.Printf("End DocServer read config\n")
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	conf := &config{}
// 	err = V.Unmarshal(conf)
// 	if err != nil {
// 		fmt.Printf("unable to decode into config struct, %v", err)
// 	}
// 	test := V.Get("testing")
// 	fmt.Printf("Test: %s\n", test)
// 	conf.Initialized = "AllScripts Delivery"
// 	return conf
// }

// //Config is the easy methog to access the configurations
// func GetConfig() *config {
// 	return Conf
// }

//TODO: put certificates in a secure K/V store as actual key and Certificate
// func GetTLSCert() string {
// 	var crt = ""
// 	cert := Conf.AllScriptsCertName
// 	if cert == "" {
// 		crt = fmt.Sprintf("/etc/ssl/vscerts/%s-selfsigned.crt", cert)
// 	} else {
// 		crt = "/etc/ssl/vscerts/localhost-selfsigned.crt"
// 	}
// 	fmt.Printf("Using Certificate: %s\n", crt)
// 	return crt
// }

// func GetTLSKey() string {
// 	var key = ""
// 	cert := Conf.AllScriptsCertName
// 	if cert == "" {
// 		key = fmt.Sprintf("/etc/ssl/vsprivate/%s-selfsigned.key", cert)
// 	} else {
// 		key = "/etc/ssl/vsprivate/localhost-selfsigned.key"
// 	}

// 	fmt.Printf("Using Certificate: %s\n", cert)
// 	fmt.Printf("Using CertificateKey: %s\n", key)
// 	return key

// }
