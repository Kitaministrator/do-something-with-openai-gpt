package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	openai "github.com/sashabaranov/go-openai"
)

type config struct {
	AuthToken          string `json:"authtoken"`
	BaseURL            string `json:"baseurl"`
	OrgID              string `json:"orgid"`
	EmptyMessagesLimit uint   `json:"emptymessageslimit"`
}

var cfg config
var ClientCfg openai.ClientConfig

func loadConfig() error {
	// Read local config file first if exists, else read from env
	file, err := os.Open("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			// config file not found, try to read from env
			loadConfigFromEnv()
			return nil
		}
		return err
	}
	defer file.Close()

	// Decode config.json
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	setValue()

	return nil
}

func loadConfigFromEnv() {
	// Read config from env
	cfg.AuthToken = os.Getenv("OPENAI_AUTH_TOKEN")
	cfg.BaseURL = os.Getenv("OPENAI_BASE_URL")
	cfg.OrgID = os.Getenv("OPENAI_ORG_ID")
	limit, err := strconv.ParseUint(os.Getenv("OPENAI_MTY_MSG_LIM"), 10, 64)
	if err != nil {
		cfg.EmptyMessagesLimit = 300
		log.Printf("Failed to parse OPENAI_MTY_MSG_LIM: %s. Using default value of %d.\n", err, cfg.EmptyMessagesLimit)
	} else {
		cfg.EmptyMessagesLimit = uint(limit)
	}

	setValue()

}

func setValue() {
	ClientCfg = openai.DefaultConfig(cfg.AuthToken)
	ClientCfg.BaseURL = cfg.BaseURL
	ClientCfg.OrgID = cfg.OrgID
	ClientCfg.EmptyMessagesLimit = cfg.EmptyMessagesLimit
}

/// external file config.json
/// path: ./config.json
/// structure:
/*
	{
		"authtoken": "your-token",
		"baseurl": "your-private-domain-api",
		"orgid": "your-organization-id",
		"emptymessageslimit": 300
	}
*/

/// list of env variables
/*
	OPENAI_AUTH_TOKEN
	OPENAI_BASE_URL
	OPENAI_ORG_ID
	OPENAI_MTY_MSG_LIM
*/
