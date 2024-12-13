# Pocket ID Metrics Exporter

Dieser Pocket ID Metrics Exporter ist ein Go-Programm, das Metriken aus einer Datenbank (SQLite oder PostgreSQL) ausliest und diese für Prometheus bereitstellt.

## Funktionen

- Unterstützung für SQLite und PostgreSQL Datenbanken
- Erfassung der Gesamtanzahl der Anmeldungen
- Erfassung der Gesamtanzahl der Benutzer
- Bereitstellung der Metriken über einen HTTP-Endpunkt für Prometheus

## Voraussetzungen

- Go 1.15 oder höher
- SQLite oder PostgreSQL Datenbank
- Zugriff auf die Tabellen `Audit_Logs` und `Users` in der Datenbank

## Installation

1. Klonen Sie das Repository:
   ```
   git clone https://github.com/yourusername/pocket-id-metrics-exporter.git
   cd pocket-id-metrics-exporter
   ```

2. Installieren Sie die erforderlichen Abhängigkeiten:
   ```
   go get github.com/prometheus/client_golang/prometheus
   go get github.com/prometheus/client_golang/prometheus/promhttp
   go get github.com/lib/pq
   go get github.com/mattn/go-sqlite3
   ```

## Konfiguration

Das Programm verwendet Umgebungsvariablen für die Konfiguration:

- `DB_TYPE`: Der Typ der Datenbank ("sqlite3" oder "postgres")
- `DB_CONNECTION`: Die Verbindungszeichenfolge für die Datenbank

### Beispiele:

Für SQLite:
```
export DB_TYPE=sqlite3
export DB_CONNECTION=./data/pocket-id.db
```

Für PostgreSQL:
```
export DB_TYPE=postgres
export DB_CONNECTION="host=localhost port=5432 user=yourusername dbname=yourdbname password=yourpassword sslmode=disable"
```

## Verwendung

1. Setzen Sie die Umgebungsvariablen wie oben beschrieben.

2. Starten Sie das Programm:
   ```
   go run main.go
   ```

3. Das Programm läuft nun auf `http://localhost:3000`. Die Metriken sind unter dem `/metrics` Endpunkt verfügbar.

## Metriken

- `pocket_id_login_count`: Gesamtanzahl der Anmeldungen
- `pocket_id_user_count`: Gesamtanzahl der Benutzer

## Prometheus Konfiguration

Fügen Sie folgende Job-Konfiguration zu Ihrer `prometheus.yml` hinzu:

```yaml
scrape_configs:
  - job_name: 'pocket_id_metrics'
    static_configs:
      - targets: ['localhost:3000']
```

## Fehlerbehebung

- Stellen Sie sicher, dass die Datenbank erreichbar ist und die erforderlichen Tabellen existieren.
- Überprüfen Sie die Konsolenausgabe auf Fehlermeldungen.
- Vergewissern Sie sich, dass die Umgebungsvariablen korrekt gesetzt sind.

## Beitragen

Beiträge sind willkommen! Bitte erstellen Sie einen Pull Request oder öffnen Sie ein Issue für Vorschläge und Fehlermeldungen.

## Lizenz

Dieses Projekt steht unter der MIT-Lizenz. Siehe die [LICENSE](LICENSE) Datei für Details.
