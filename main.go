/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/joho/godotenv"
	"log"
	"my-cli-test/cmd"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
