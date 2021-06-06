package util

import (
	"encoding/json"
	"fmt"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"sync"
)

var jsonConfigPath string

func init() {
	flag.StringVar(&jsonConfigPath, "filepath", "config.army.modal.json", "json config file location.")
}

type GlobalConfig struct {
	Database DatabaseConfig `json:"database"`
	Self     SelfConfig     `json:"self"`
}

type DatabaseConfig struct {
	Types  string `json:"types"`
	Local  string `json:"local"`
	Online string `json:"online"`
}

type SelfConfig struct {
	Port string `json:"port"`
	Flag int    `json:"flag"`
	Tag  int    `json:"tag"`
}

var (
	globalConfig *GlobalConfig
	configMux    sync.RWMutex
)

func Config() *GlobalConfig {
	return globalConfig
}

func InitConfigJson() error {
	var config GlobalConfig
	file, err := ioutil.ReadFile(jsonConfigPath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}

	if err = json.Unmarshal(file, &config); err != nil {
		fmt.Println("配置文件读取失败", err)
		return err
	}

	configMux.Lock()
	globalConfig = &config
	configMux.Unlock()
	return nil
}
