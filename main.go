package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/C0deNe0/go-prisma/config"
	"github.com/C0deNe0/go-prisma/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load the env", err)
	}

	fmt.Println("server started at " + os.Getenv("PORT"))
	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}


	//handleing the db connection
	db,err := config.ConnectDB()
	utils.ErrorPanic(err)


	//here we disconnect to the PrismaClient in the last
	defer db.Prisma.Disconnect()
}
