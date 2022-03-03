package src
import (
	"os"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	router   *mux.Router
	port           string
	//env            *Environment
	source         string
	//mrnID          string
	appURL         string
	appName        string
	mode           string
	serverURL      string
	imageURL       string
	recordLimit    string
	logLevel       log.Level
}

var config Config
func InitializeAll(mongo, ca string)*Config {
	config = Config{}
	config.port = os.Getenv("DBPORT")
	return &config

}