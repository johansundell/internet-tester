package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type appSettings struct {
	Debug      bool   `json:"debug"`
	SecToWait  int    `json:"sec_to_wait"`
	UrlToCheck string `json:"url_to_test"`
}

const filenameSettings = "settings.json"

var settings appSettings

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dir, _ := filepath.Split(ex)
	dat, err := ioutil.ReadFile(dir + filenameSettings)
	if err != nil {
		data, _ := json.Marshal(settings)
		ioutil.WriteFile(dir+filenameSettings, data, 0664)
		log.Fatal("settings.json missing, " + err.Error())
	}

	if err := json.Unmarshal(dat, &settings); err != nil {
		log.Fatal(err)
	}
}
