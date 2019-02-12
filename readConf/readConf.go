package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Baseurl   string
	Title     string
	Templates string
	Posts     string
	Public    string
	Admin     string
	Metadata  string
	Index     string
}

func ReadConfig() Config {
	configfile, err := ioutil.ReadFile("./conf/app.conf")
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if k, err := toml.Decode(string(configfile), &config); err != nil {
		log.Fatal(err)
	} else {
		log.Println(k)
	}
	//log.Print(config.Index)
	return config
}

func main() {
	var config = ReadConfig()
	fmt.Println(config.Baseurl)
}
