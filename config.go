package main

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigData struct {
	SourceSubUrl   string `json:"sourceSubUrl"`
	ClientType     string `json:"clientType"`
	CustomBackend  string `json:"customBackend"`
	RemoteConfig   string `json:"remoteConfig"`
	ExcludeRemarks string `json:"excludeRemarks"`
	IncludeRemarks string `json:"includeRemarks"`
	Filename       string `json:"filename"`
	Emoji          bool   `json:"emoji"`
	NodeList       bool   `json:"nodeList"`
	Extraset       bool   `json:"extraset"`
	Sort           bool   `json:"sort"`
	Udp            bool   `json:"udp"`
	Tfo            bool   `json:"tfo"`
	Scv            bool   `json:"scv"`
	Fdn            bool   `json:"fdn"`
	AppendType     bool   `json:"appendType"`
	Insert         bool   `json:"insert"`
	NewName        bool   `json:"new_name"`
	Tpl            Tpl    `json:"tpl"`
	SubUrl         string `json:"subUrl"`
	SubUrlShort    string `json:"subUrlShort"`
}

type Tpl struct {
	Surge TplConfig `json:"surge"`
	Clash TplConfig `json:"clash"`
}

type TplConfig struct {
	Doh bool `json:"doh"`
}

func Config() *ConfigData {
	// Check if config file exists
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		// Create new config file
		config := ConfigData{
			// Initialize default config values here
		}
		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("config.json", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created new config file")
	} else {
		log.Println("Config file already exists")
	}

	// Load config file
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config ConfigData
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
	// Use config values here
}

func UpdateConfig(config *ConfigData) {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("config.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}