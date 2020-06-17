package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

// RFConfig configuration file structure
type RFConfig struct {
	Sentinel_Service struct {
		Hostname string `yaml:"hostname"`
		Port     string `yaml:"port"`
	} `yaml:"sentinel_service"`
	Redis_Auth               string `yaml:"redis_auth,omitempty"`
	Sentinel_Master_Name     string `yaml:"sentinel_master_name,omitempty"`
	Redis_State_Ful_Set_Name string `yaml:"redis_state_ful_set_name"`
	Service_Label_Name       string `yaml:"service_label_name"`
	Check_Timeout            uint64 `yaml:"check_timeout"`
}

const configPath = "/etc/rfailover/rfailover.yml"

var (
	// RedisAuth variables from configuration file
	RedisAuth string
	// SentinelMasterName var from configuration file
	SentinelMasterName string
	// SentinelService var from configuration file
	SentinelService string
	// ServiceLabelName var from configuration file
	ServiceLabelName string
	// CheckTimeout var from configuration file
	CheckTimeout uint64
)

func main() {
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panic("Cannot read configuration file at path ", configPath)
	}

	var rfC RFConfig
	err = yaml.Unmarshal([]byte(configData), &rfC)
	if err != nil {
		log.Panic("Cannot validate configuration file: ", err)
	}
	concatRedisSentinelHostname := fmt.Sprintf("%s:%s", rfC.Sentinel_Service.Hostname, rfC.Sentinel_Service.Port)
	RedisAuth = rfC.Redis_Auth
	SentinelMasterName = rfC.Sentinel_Master_Name
	SentinelService = concatRedisSentinelHostname
	ServiceLabelName = rfC.Service_Label_Name
	CheckTimeout = rfC.Check_Timeout

	for {
		master, err := getRedisMaster(concatRedisSentinelHostname, rfC.Sentinel_Master_Name)
		if err != nil {
			log.Panic("Cannot getting address of Redis Master Pod: ", err)
		}
		redisCheckEndpoint(master, "test-redis", "rfs-redisfailover", rfC.Redis_State_Ful_Set_Name)
		log.Printf("Getting Redis master pod: %s:%s\n", master.IP, master.PORT)
		// Check Timeout cannot be less then 0
		if CheckTimeout < 1 {
			CheckTimeout = 1
		}
		time.Sleep(time.Duration(CheckTimeout) * time.Second)
	}
}
