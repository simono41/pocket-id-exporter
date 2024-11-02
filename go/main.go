package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath = "./data/pocket-id.db"
	port   = ":3000"
)

var (
	db *sql.DB

	loginCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pocket_id_login_count",
		Help: "Total number of sign-ins",
	})

	userCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pocket_id_user_count",
		Help: "Total number of users",
	})
)

func init() {
	prometheus.MustRegister(loginCount)
	prometheus.MustRegister(userCount)
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to the Pocket ID SQLite database.")

	http.HandleFunc("/metrics", metricsHandler)
	http.Handle("/", promhttp.Handler())

	log.Printf("Server running on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	updateMetrics()
	promhttp.Handler().ServeHTTP(w, r)
}

func updateMetrics() {
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM Audit_Logs WHERE event = 'SIGN_IN'").Scan(&count)
	if err != nil {
		log.Printf("Error querying login count: %v", err)
	} else {
		loginCount.Set(float64(count))
	}

	err = db.QueryRow("SELECT COUNT(*) FROM Users").Scan(&count)
	if err != nil {
		log.Printf("Error querying user count: %v", err)
	} else {
		userCount.Set(float64(count))
	}
}
