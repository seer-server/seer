package config

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/tree-server/trees/errors"
	"github.com/tree-server/trees/log"
)

type Database struct {
	Host string `toml:"host"`
	//Auth bool
	//User string
	//Pass string
}

type Config struct {
	DB       Database `toml:"database"`
	RootNode string   `toml:"root_node_id"`
}

const configFileName = "Trees.toml"

var (
	loadedConfig *Config
	logger       *log.Logger
)

func init() {
	logger = log.Make("config", ":stdout:", log.LogDebug)
}

func LoadOrCreate() {
	err := load()

	if os.IsNotExist(err) {
		config := withDefaults()
		f, err := os.Create(configFileName)
		if err != nil {
			logger.Fatal(err, errors.ErrCreateConfigFailed)
		}
		defer f.Close()

		buf := new(bytes.Buffer)
		if err = toml.NewEncoder(buf).Encode(config); err != nil {
			logger.Fatal(err, errors.ErrCreateConfigFailed)
		}
		_, err = f.Write(buf.Bytes())
		if err != nil {
			logger.Fatal(err, errors.ErrCreateConfigFailed)
		}
	}
}

func load() error {
	loadedConfig = withDefaults()
	f, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(b, &loadedConfig)
	if err != nil {
		loadedConfig = withDefaults()
		return err
	}

	return nil
}

func Get() *Config {
	if loadedConfig == nil {
		load()
	}

	return loadedConfig
}

func withDefaults() *Config {
	return &Config{
		DB: Database{
			Host: "localhost:7474",
			//Auth: false,
			//User: "",
			//Pass: "",
		},
		RootNode: "",
	}
}
