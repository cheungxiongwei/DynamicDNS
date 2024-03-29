package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"time"
)

type CMDParam struct {
	Host     string `json:"host"`
	Domain   string `json:"domain"`
	Password string `json:"password"`
	TimeOut  int    `json:"time_out"`
}

var cmd CMDParam

var ConfigFile string

func init() {
	// set cmd line param
	flag.StringVar(&cmd.Host, "host", "", "host")
	flag.StringVar(&cmd.Domain, "domain", "", "domain_name")
	flag.StringVar(&cmd.Password, "password", "", "ddns_password")
	flag.IntVar(&cmd.TimeOut, "timeout", 15*60, "auto update time.default value 15 min")

	flag.StringVar(&ConfigFile, "c", "", "config file")
}

func main() {
	flag.Parse()
	if !flag.Parsed() {
		println("please input *.exe -h get help.")
		os.Exit(0)
	}

	println("ddns runing.")

	if len(cmd.Password) > 0 && len(cmd.Domain) > 0 && len(cmd.Host) > 0 {
		ticker := time.NewTicker(time.Duration(cmd.TimeOut) * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			updateRemoteIp(cmd)
		}
	} else {

		if len(ConfigFile) == 0 {
			ConfigFile = "config.json"
		}

		file, _ := ioutil.ReadFile(ConfigFile)

		err := json.Unmarshal([]byte(file), &cmd)
		if err != nil {
			println("not find config.json file.")
			return
		}

		ticker := time.NewTicker(time.Duration(cmd.TimeOut) * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			UpdateTencentIp(cmd)
			// updateRemoteIp(cmd)
			// ipv4, _ := GetLocalHostAddress()
			// println(ipv4)
		}
	}
}
