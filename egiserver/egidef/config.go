package egidef

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Endpoint 				string `json:"endpoint"`
	ResourceTpl 			string `json:"resourceTpl"`
	OsTpl 					string `json:"osTpl"`
	PublicKey 				string `json:"publicKey"`
	Contextualisation 		string `json:"contextualisation"`
	ProxyPath 				string `json:"proxyPath"`
	Auth 					string `json:"auth"`
	Vo 						string `json:"vo"`
	VomsDir 				string `json:"vomsDir"`
	TrustedCertificatesPath string `json:"trustedCertificatesPath"`
	PublicIP 				string `json:"publicIP"`
	State 					string `json:"vmState"`
}

// ReadConfigFile reads a configuration file
func ReadConfigFile(configFilepath string) (Configuration, error) {
	var config Configuration

	file, err := os.Open(configFilepath)
	if err != nil {
		return config, Err(err, "Cannot open config file %s", configFilepath)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	
	err = decoder.Decode(&config)
	if err != nil {
		return config, Err(err, "Cannot read config file %s", configFilepath)
	}

	return config, nil
}
