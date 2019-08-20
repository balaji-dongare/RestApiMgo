package propparser

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//Configuration struct
type Configuration struct {
	Port             string // port no
	ConnectionString string // connection string
	Database         string // database name
	Collection       string // collection
}

/*ReadConfig Reading the configs from  db.properties
 */
func ReadConfig() Configuration {
	var configfile = "config.properties"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Configuration
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
