package main

import (
	"context"
	"fmt"
	"log"

	"github.com/elangreza14/schat/config"
	migrations "github.com/elangreza14/schat/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	env, err := config.SetupEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	conn, err := pgxpool.New(ctx, getDbURL(env))
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	migrator, err := migrations.NewMigrator(ctx, getDbURL(env))
	if err != nil {
		log.Fatal(err)
	}

	if err = migrator.SetLatest(ctx); err != nil {
		log.Fatal(err)
	}

	if err = migrator.Status(ctx); err != nil {
		log.Fatal(err)
	}

	// entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	// ms := ulid.Timestamp(time.Now())
	// fmt.Println(ulid.New(ms, entropy))

	// newID := internal.NewID()
	// var a GenID = newID
}

func getDbURL(env *config.Env) string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresHostname,
		env.PostgresPort,
		env.PostgresDB,
		env.PostgresSsl)
}
