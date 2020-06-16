package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type RFConfig struct {
	Sentinel_Service struct {
		Hostname string `yaml:"hostname"`
		Port     string `yaml:"port"`
	} `yaml:"sentinel_service"`
	Redis_Auth               string `yaml:"redis_auth,omitempty"`
	Sentinel_Pass            string `yaml:"sentinel_password,omitempty"`
	Sentinel_Master_Name     string `yaml:"sentinel_master_name,omitempty"`
	Redis_State_Ful_Set_Name string `yaml:"redis_state_ful_set_name"`
	Service_Label_Name       string `yaml:"service_label_name"`
}

const configPath = "/etc/rfailover/rfailover.yml"

var RedisAuth, SentinelPass, SentinelMasterName, SentinelService, ServiceLabelName string

func main() {
	// ChangeService("test-redis", "pod", 3530)
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panic("Cannot read configuration file at path ", configPath)
	}

	log.Println("raw config: ", string(configData))
	var rfC RFConfig
	err = yaml.Unmarshal([]byte(configData), &rfC)
	if err != nil {
		log.Panic("Cannot validate configuration file: ", err)
	}
	concatRedisSentinelHostname := fmt.Sprintf("%s:%s", rfC.Sentinel_Service.Hostname, rfC.Sentinel_Service.Port)
	RedisAuth = rfC.Redis_Auth
	SentinelPass = rfC.Sentinel_Pass
	SentinelMasterName = rfC.Sentinel_Master_Name
	SentinelService = concatRedisSentinelHostname
	ServiceLabelName = rfC.Service_Label_Name
	log.Println("config struct: ", rfC)
	for {
		master, err := getRedisMaster(concatRedisSentinelHostname, rfC.Sentinel_Master_Name, rfC.Sentinel_Pass)
		if err != nil {
			log.Panic("F*cking error: ", err)
		}
		redisCheckEndpoint(master, "test-redis", "rfs-redisfailover", rfC.Redis_State_Ful_Set_Name)
		log.Printf("%s:%s\n", master.IP, master.PORT)

		time.Sleep(5 * time.Second)
	}

}
