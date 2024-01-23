# http2go

`http2go` ist ein vielseitiger Webserver, präsentiert von [SomeCatCode](https://github.com/SomeCatCode). Dieses Tool ermöglicht es Ihnen, mit minimalem Aufwand einen voll funktionsfähigen Webserver aufzusetzen. Die Konfiguration des Root-Pfads und des Ports erfolgt durch eine einfache JSON-Datei (config.json).

## Hauptmerkmale

- **Flexible Konfiguration**: Der Root-Pfad und der Port des Webservers werden durch die `config.json` Datei bestimmt.
- **Automatische Konfigurationsinitialisierung**: Bei Abwesenheit der `config.json` wird sie automatisch mit Standardwerten für den aktuellen Ordner und den Port 8080 erstellt.
- **Benutzerfreundlich**: Einfach zu starten und zu verwalten, ideal für Entwicklungsumgebungen oder kleinere Produktionsumgebungen.

## Voraussetzungen

Installieren Sie [Go](https://golang.org/) auf Ihrem System, um `http2go` zu kompilieren und auszuführen.

## Installation

Klonen Sie das Repository und navigieren Sie in das Projektverzeichnis:

```bash
git clone https://github.com/SomeCatCode/http2go.git
cd http2go
```

Kompilieren und starten Sie den Server:

```bash
go build
./http2go
```

## Konfiguration
http2go nutzt eine config.json Datei im Hauptverzeichnis für die Konfiguration. Die Datei sollte folgendermaßen aufgebaut sein:

```json
{
    "rootPath": "Pfad/zum/Verzeichnis",
    "port": "Portnummer"
}
```
Falls die config.json nicht existiert, wird sie beim Start des Programms erstellt. Der Root-Pfad wird auf das aktuelle Verzeichnis gesetzt und der Port auf 8080, falls nicht anders angegeben.

## Mitwirken
Ihr Beitrag ist willkommen und wertvoll. Wenn Sie zur Verbesserung beitragen möchten, erstellen Sie bitte einen Fork des Repositories und senden Sie einen Pull Request. Sie können auch ein Issue eröffnen, wenn Sie Fragen haben oder einen Fehler melden möchten.

## Lizenz
http2go wird unter der MIT-Lizenz veröffentlicht. Weitere Informationen finden Sie in der LICENSE-Datei.

Viel Spaß und erfolgreiche Entwicklung! :cat: :computer: - SomeCatCode