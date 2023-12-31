package migrate

import (
	"context"
	"embed"
	"fmt"
	"io/fs"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
)

//go:embed data/*.sql
var migrationFiles embed.FS

type (
	internalMigrator struct {
		m *migrate.Migrator
	}

	Migrator interface {
		Status(ctx context.Context) error
		SetVersion(ctx context.Context, version int32) error
		SetLatest(ctx context.Context) error
	}
)

func NewMigrator(ctx context.Context, dbURL string) (Migrator, error) {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if conn.Ping(ctx); err != nil {
		return nil, err
	}
	migrator, err := migrate.NewMigrator(ctx, conn, "db_version")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	migrationRoot, _ := fs.Sub(migrationFiles, "data")
	err = migrator.LoadMigrations(migrationRoot)
	if err != nil {
		fmt.Println(" a", err)
		return nil, err
	}

	return &internalMigrator{
		m: migrator,
	}, nil
}

func (m internalMigrator) Status(ctx context.Context) error {
	currentVersion, err := m.m.GetCurrentVersion(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	logo := "current -->"
	logoZero := ""
	if currentVersion == 0 {
		logoZero = logo
	}
	fmt.Printf("%11s %3d %s\n", logoZero, 0, "start")

	for _, migration := range m.m.Migrations {
		indicator := "   "
		if currentVersion == migration.Sequence {
			indicator = logo
		}
		fmt.Printf(
			"%11s %3d %s\n",
			indicator,
			migration.Sequence, migration.Name)
	}

	return nil
}

func (m internalMigrator) SetVersion(ctx context.Context, version int32) error {
	return m.m.MigrateTo(ctx, version)
}

func (m internalMigrator) SetLatest(ctx context.Context) error {
	return m.m.Migrate(ctx)
}
