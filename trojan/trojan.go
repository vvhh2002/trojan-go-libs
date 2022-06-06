package trojan

import (
	"github.com/p4gefau1t/trojan-go/constant"
	"io/ioutil"
	"strings"

	"github.com/p4gefau1t/trojan-go/common"
	"github.com/p4gefau1t/trojan-go/log"
	_ "github.com/p4gefau1t/trojan-go/log/simplelog"
	"github.com/p4gefau1t/trojan-go/proxy"
	_ "github.com/p4gefau1t/trojan-go/proxy/client"
	//_ "github.com/p4gefau1t/trojan-go/stat/memory"
)

var client common.Runnable

func RunClient(filename string) {
	log.Info("Trojan-Go core version", constant.Version)
	if client != nil {
		log.Error("Client is already running")
		return
	}
	log.Info("Running client, config file:", filename)
	configBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("failed to read file", err)
	}
	log.Info("config:", string(configBytes))
	//config, err := parseJSON(configBytes)
	//if err != nil {
	//	log.Error("error", err)
	//	return
	//}
	//client, err = proxy.NewProxy(config)
	client, err = proxy.NewProxyFromConfigData(configBytes, true)
	if err != nil {
		log.Error("error", err)
		return
	}
	go client.Run()
	log.Info("trojan launched")
}

func StopClient() {
	log.Info("Stopping client")
	if client != nil {
		client.Close()
		client = nil
	}
	log.Info("Stopped")
}

// SetLoglevel set trojan log level
// possible input: debug/info/warn/error/none
func SetLoglevel(logLevel string) {
	// Set log level.
	switch strings.ToLower(logLevel) {
	case "debug":
		log.SetLogLevel(log.AllLevel)
	case "info":
		log.SetLogLevel(log.InfoLevel)
	case "warn":
		log.SetLogLevel(log.WarnLevel)
	case "error":
		log.SetLogLevel(log.ErrorLevel)
	case "none":
		log.SetLogLevel(log.OffLevel)
	default:
		panic("unsupport logging level")
	}
}
