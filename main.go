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
	// Lade die Konfiguration aus der JSON-Datei.
	config, err := getConfig()
	if err != nil {
		log.Fatalf("Fehler beim Laden der Konfiguration: %v", err)
		return
	}

	fs := http.FileServer(http.Dir(config.RootPath))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Printf("Listening on :%s...\n", config.Port)
	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig() (*Config, error) {
	// Ermittle den Pfad der ausführbaren Datei und das Verzeichnis, in dem sie sich befindet.
	exPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Fehler beim Ermitteln des Executable-Pfads: %v", err)
	}
	exDir := filepath.Dir(exPath)

	configFile := filepath.Join(exDir, "http2go_config.json")

	// Überprüfe, ob Config-Datei existiert
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Erstelle Config-Datei mit Standardwerten
		defaultConfig := Config{
			RootPath: filepath.Dir(exPath), // Setze den Root-Pfad auf das Verzeichnis der ausführbaren Datei.
			Port:     "8080",
		}

		data, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			log.Fatalf("Fehler beim Erzeugen der Standard-Config: %v", err)
			return nil, err
		}

		// Schreibe die Standard-Konfiguration in die Datei.
		err = os.WriteFile(configFile, data, 0644)
		if err != nil {
			log.Fatalf("Fehler beim Schreiben der Standard-Config: %v", err)
			return nil, err
		}
	}

	// Lese Config-Datei
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Fehler beim Lesen der Config-Datei: %v", err)
		return nil, err
	}

	// Parse die Konfigurationsdaten in die Config-Struktur.
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Config-Datei: %v", err)
		return nil, err
	}

	// Überprüfe, ob der angegebene Root-Pfad existiert und ein Verzeichnis ist.
	rootInfo, err := os.Stat(config.RootPath)
	if os.IsNotExist(err) {
		log.Fatalf("Der angegebene Root-Pfad existiert nicht: %v", config.RootPath)
		return nil, err
	} else if !rootInfo.IsDir() {
		log.Fatalf("Der angegebene Root-Pfad ist kein Verzeichnis: %v", config.RootPath)
		return nil, err
	}

	return &config, nil
}
