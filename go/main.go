package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	port = ":3000"
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
	dbType := os.Getenv("DB_TYPE")
	dbConnection := os.Getenv("DB_CONNECTION")

	if dbType == "" || dbConnection == "" {
		log.Fatal("DB_TYPE and DB_CONNECTION environment variables must be set")
	}

	var err error
	db, err = sql.Open(dbType, dbConnection)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Printf("Connected to the %s database.", dbType)

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

	query := "SELECT COUNT(*) FROM Audit_Logs WHERE event = 'SIGN_IN'"
	if os.Getenv("DB_TYPE") == "postgres" {
		query = "SELECT COUNT(*) FROM audit_logs WHERE event = 'SIGN_IN'"
	}

	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Printf("Error querying login count: %v", err)
	} else {
		loginCount.Set(float64(count))
	}

	query = "SELECT COUNT(*) FROM Users"
	if os.Getenv("DB_TYPE") == "postgres" {
		query = "SELECT COUNT(*) FROM users"
	}

	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Printf("Error querying user count: %v", err)
	} else {
		userCount.Set(float64(count))
	}
}
