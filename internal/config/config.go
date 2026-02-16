package config

import (
	"encoding/json"
	"os"
)

/* Using encoding/decoding is important because it will read the JSON-encoded (or decoded data) as a stream and allow for buffering both in terms of reading and writing
while marshalling and unmarshalling works, in this case we know how and where the data is coming from so it is ok to read it as a stream because we KNOW the json
format thanks to struct titles */

// define const that provides filename
const configFileName = ".gatorconfig.json"

// struct object that holds the config data; use struct titles for efficient encoding/decoding
type Config struct {
	Dburl string `json:"db_url"`
	Username string `json:"current_user_name"`
}

// returns file path of home directory
func getConfigFilePath() (string, error) {
	// for use on own system
	_, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// for use in codespaces
	github_path := "/workspaces/bloggator/" + configFileName
	return github_path, nil
}

// updates a config struct with new values
func write(cfg Config) error{
	// get the file path that needs to be written to and gracefully handle errors
	dir, err := getConfigFilePath()
	if err != nil {
		return err
	}
	// create a file path
	file, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	// create empty struct
	config := Config{}
	// get path of file and handle any errors
	dir, err := getConfigFilePath()
	if err != nil {
		return config, err
	}
	// opens file for reading and handles any errors
	jsonFile, err := os.Open(dir)
	if err != nil {
		return config, err
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (cfg *Config) SetUser(user string) error {
	updateCfg := Config {
		Dburl: cfg.Dburl,
		Username: user,
	}
	write(updateCfg)
	return nil
}