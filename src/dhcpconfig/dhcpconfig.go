/*
	For dhcpconfig package, it read configuration from
	an json file.
	For configuration file, it include server infomation,
	instance num and so on
*/

package dhcpconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type tServer struct {
	ServerName       string
	ServerListenIp   string
	ServerListenPort uint32
}

type tAllCfg struct {
	Server tServer
}

type TdhcpCfg struct {
	configFileName string
	configContent  tAllCfg
}

func (this *TdhcpCfg) Init(cfgname string) error {
	var data []byte
	var err error
	data, err = ioutil.ReadFile(cfgname)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	//parse data to json struct
	this.configFileName = cfgname
	err = json.Unmarshal(data, &this.configContent)
	return err
}
