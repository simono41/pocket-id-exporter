# pocket-id-exporter

Der pocket-id-exporter ist ein Prometheus-Exporter für die Pocket ID-Datenbank. Er stellt Metriken über die Anzahl der Anmeldungen und Benutzer bereit.

## Funktionen

- Exportiert die Gesamtzahl der Anmeldungen als Metrik `pocket_id_login_count`
- Exportiert die Gesamtzahl der Benutzer als Metrik `pocket_id_user_count`
- Stellt Metriken im Prometheus-Format unter dem `/metrics`-Endpunkt bereit

## Voraussetzungen

- Docker
- Docker Compose

## Verwendung mit Docker Compose

1. Erstellen Sie eine `docker-compose.yml`-Datei im Projektverzeichnis mit folgendem Inhalt:

```yaml
version: '3'
services:
  pocket-id-exporter:
    build: .
    ports:
      - "3000:3000"
    volumes:
      - ./data:/app/data
```

2. Erstellen Sie eine `Dockerfile` im Projektverzeichnis:

```Dockerfile
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o pocket-id-exporter

EXPOSE 3000

CMD ["./pocket-id-exporter"]
```

3. Stellen Sie sicher, dass Ihre SQLite-Datenbank im Verzeichnis `./data` mit dem Namen `pocket-id.db` liegt.

4. Bauen und starten Sie den Container mit Docker Compose:

```bash
docker-compose up --build
```

Der Exporter ist nun unter `http://localhost:3000/metrics` erreichbar.

## Konfiguration von Prometheus

Fügen Sie folgende Job-Konfiguration zu Ihrer `prometheus.yml` hinzu, um die Metriken zu scrapen:

```yaml
scrape_configs:
  - job_name: 'pocket-id'
    static_configs:
      - targets: ['localhost:3000']
```

## Entwicklung

Um den Exporter lokal zu entwickeln und zu testen:

1. Installieren Sie Go (Version 1.17 oder höher)
2. Klonen Sie das Repository
3. Installieren Sie die Abhängigkeiten:

```bash
go mod download
```

4. Bauen und starten Sie den Exporter:

```bash
go build
./pocket-id-exporter
```

## Lizenz

[MIT License](LICENSE)

https://github.com/stonith404/pocket-id

https://goneuland.de/pocket-id-mit-docker-und-traefik-installieren/

https://github.com/stonith404/pocket-id/issues/56
