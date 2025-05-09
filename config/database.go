package config

import (
	"github.com/C0deNe0/go-prisma/prisma/db"
	"github.com/rs/zerolog/log"
)	

func ConnectDB() (*db.PrismaClient, error) {
	client := db.NewClient()

	//here we connect to the prisma clied with Prisma.Connect()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	log.Info().Msg("connected to database!")
	return client, nil
}
