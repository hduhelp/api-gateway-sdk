package svs

import (
	"github.com/samber/lo"
	"os"
)

const (
	DefaultConfigEndpoint = "https://api.hduhelp.com/config"
)

var (
	_configEndpointFromEnv = os.Getenv("CONFIG_ENDPOINT")
	ConfigEndpoint         = lo.Ternary[string](_configEndpointFromEnv == "", DefaultConfigEndpoint, _configEndpointFromEnv)
	InstanceID             = os.Getenv("INSTANCE_ID")
)

func InitConfig(conf any) {
	if InstanceID == "" {
		panic("empty instanceID")
	}
	if err := checkConfigValidate(); err != nil {
		panic(err)
	}
	for {
		getConfigFromServer(conf)
	}
}

func checkConfigValidate() error {
	return nil
}

func getConfigFromServer(conf any) {

}
