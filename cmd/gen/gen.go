package main

import (
	"github.com/alendavid/talk-about-backend/internal/database"
	"github.com/alendavid/talk-about-backend/pkg/plants"
	"gorm.io/gen"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "plants"
	password = "password"
	name     = "plants"
)

func main() {
	db, err := database.Run(host, port, user, password, name)
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/persistence/plants",
		// TODO: remove these Modes
		Mode:           gen.WithDefaultQuery | gen.WithQueryInterface,
		WithUnitTest:   true,
		FieldCoverable: true,
	})

	g.UseDB(&db.DB)

	g.ApplyBasic(plants.Plant{})

	if err := g.Revise(); err != nil {
		panic(err)
	}

	g.Execute()
}
