package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	RootPath string `json:"rootPath"`
	Port     string `json:"port"`
}

// TODO: Ich muss den Code echt mal aufräumen und mehr Debugging einbauen

func main() {
	config := loadConfig()

	fs := http.FileServer(http.Dir(config.RootPath))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Printf("Listening on :%s...\n", config.Port)
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() Config {
	configFilePath := "config.json"
	defaultRootPath, _ := os.Getwd()

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		config := Config{RootPath: defaultRootPath, Port: "8080"}
		configJson, _ := json.Marshal(config)
		_ = os.WriteFile(configFilePath, configJson, 0644)
		return config
	} else {
		configJson, _ := os.ReadFile(configFilePath)
		var config Config
		_ = json.Unmarshal(configJson, &config)

		if !filepath.IsAbs(config.RootPath) {
			config.RootPath = filepath.Join(defaultRootPath, config.RootPath)
		}

		if config.Port == "" {
			config.Port = "8080"
		}

		return config
	}
}
