package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

var initializers map[string]func(string)

func readKernelConfig() {
	cmdline, err := ioutil.ReadFile("/proc/cmdline")
	if err != nil {
		panic(fmt.Sprint("Error opening /proc/cmdline:", err.Error()))
	}
	options := strings.Split(strings.TrimSpace(string(cmdline)), " ")
	for _, option := range options {
		kv := strings.SplitN(option, "=", 2)
		if len(kv) < 2 {
			continue
		}
		if kv[0][0:4] == "pcd." {
			viper.Set(kv[0], kv[1])
		} else if kv[0] == "hostname" {
			viper.Set(kv[0], kv[1])
		}
	}
}

func saveConfig() {
	b, _ := json.MarshalIndent(viper.AllSettings(), "", "  ")
	err := ioutil.WriteFile(viper.GetString("file"), b, 0644)
	if err != nil {
		panic(fmt.Sprint("Error opening file:", err.Error()))
	}
	fmt.Println("Config saved.")

}

func handle(key string, callback func(string)) {
	if initializers == nil {
		initializers = make(map[string]func(string))
	}
	initializers[key] = callback
}

func runHandlers() {
	for _, key := range viper.AllKeys() {
		if initializers[key] != nil {
			initializers[key](viper.GetString(key))
		}
	}
}

func main() {
	viper.SetDefault("file", "/tmp/config.json")
	readKernelConfig()
	viper.SetConfigFile(viper.GetString("file"))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file: ", err.Error())
	}
	runHandlers()
	saveConfig()

	fmt.Println("Starting http server on :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)

}
