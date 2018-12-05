package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"hex-example/internal/database/psql"
	redisdb "hex-example/internal/database/redis"
	"hex-example/internal/env"
	"hex-example/internal/middleware"
	"hex-example/internal/ticket"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/apex/gateway"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DefaultRedisUrl      = "localhost:6379"
	DefaultRedisPassword = ""
	DefaultPostgresUrl   = "postgresql://postgres@localhost/ticket?sslmode=disable"
)

var docker string

func main() {

	var server bool
	var dbType, dbURL, redisPassword string
	flag.StringVar(&dbType, "database", "redis", "database type [redis, psql]")
	flag.BoolVar(&server, "server", false, "run in server mode")
	flag.Parse()

	var ticketRepo ticket.TicketRepository

	switch dbType {
	case "psql":
		dbURL = env.EnvString("DATABASE_URL", DefaultPostgresUrl)
		pconn := postgresConnection(dbURL)
		defer pconn.Close()
		ticketRepo = psql.NewPostgresTicketRepository(pconn)
	case "redis":
		dbURL = env.EnvString("DATABASE_URL", DefaultRedisUrl)
		redisPassword = env.EnvString("REDIS_PASSWORD", DefaultRedisPassword)
		rconn := redisConnect(dbURL, redisPassword)
		defer rconn.Close()
		ticketRepo = redisdb.NewRedisTicketRepository(rconn)
	default:
		panic("Unknown database")
	}

	ticketService := ticket.NewTicketService(ticketRepo)
	ticketHandler := ticket.NewTicketHandler(ticketService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/tickets", ticketHandler.Get).Methods("GET")
	router.HandleFunc("/tickets/{id}", ticketHandler.GetById).Methods("GET")
	router.HandleFunc("/tickets", ticketHandler.Create).Methods("POST")

	http.Handle("/", accessControl(middleware.Authenticate(router)))

	errs := make(chan error, 2)
	go func() {
		if server || docker == "true" {
			logrus.Info("Listening server mode on port :3001")
			errs <- http.ListenAndServe(":3001", nil)
		} else {
			logrus.Info("Listening lambda mode on port :3000")
			errs <- gateway.ListenAndServe(":3000", nil)
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logrus.Errorf("terminated %s", <-errs)

}
func redisConnect(url string, password string) *redis.Client {

	logrus.WithField("connection", url).Info("Connecting to Redis DB")
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	err := client.Ping().Err()

	if err != nil {
		logrus.Fatal(err)
	}
	return client

}

func postgresConnection(database string) *sql.DB {
	logrus.Info("Connecting to PostgreSQL DB")
	db, err := sql.Open("postgres", database)
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}
	return db
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
