package main

import (
	"log"

	"github.com/Md-Monowarul-Amin/go-backend/SOCIAL/internal/db"
	"github.com/Md-Monowarul-Amin/go-backend/SOCIAL/internal/env"
	"github.com/Md-Monowarul-Amin/go-backend/SOCIAL/internal/store"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:password@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db) // Replace nil with your actual database connection

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
