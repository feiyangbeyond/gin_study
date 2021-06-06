package util

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"strconv"
)

func init() {
	err := getConfigIni("app.ini")
	if err != nil {
		LogError("Load app.ini error. Message is %s", err)
		panic(err)
	}
	LogInfo("Load app.ini success")
}

var cfg *goconfig.ConfigFile

// getConfigIni Read config.
func getConfigIni(filepath string) (err error) {
	config, err := goconfig.LoadConfigFile(filepath)
	if err != nil {
		fmt.Println("Config file not found.", err)
		return err
	}
	cfg = config
	return nil
}

func GetServer() (httpPort int, wsPort int, err error) {
	httpPortTemp, err := cfg.GetValue("server", "HttpPort")
	if err != nil {
		LogError("Property [port] in server selection not found. %s", err)
		return httpPort, wsPort, err
	}

	httpPort, err = strconv.Atoi(httpPortTemp)
	if err != nil {
		LogError("Type error. %s", err)
		return httpPort, wsPort, err
	}

	wsPortTemp, err := cfg.GetValue("server", "WsPort")
	if err != nil {
		LogError("Property [WsPort] in server selection not found. %s", err)
		return httpPort, wsPort, err
	}

	wsPort, err = strconv.Atoi(wsPortTemp)
	if err != nil {
		LogError("Type error. %s", err)
		return httpPort, wsPort, err
	}

	return httpPort, wsPort, nil
}
