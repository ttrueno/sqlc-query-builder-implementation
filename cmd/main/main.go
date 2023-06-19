package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ttrueno/auth-service/config"
	"github.com/ttrueno/auth-service/internal/storage/dog/psql"
	"github.com/ttrueno/auth-service/internal/storage/dog/psql/db"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, config.DB.DSN)
	if err != nil {
		log.Fatalf("can't connect: %v", err)
	}
	defer conn.Close(context.Background())

	b := &psql.Builder{}
	b.Where("name = $1", "Aqtos")

	query := db.New(psql.Wrap(conn, b))
	doggos, err := query.ListDoggos(context.Background())
	if err != nil {
		log.Fatalf("can't get doggos: %v", err)
	}

	log.Println(doggos)

}
