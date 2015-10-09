package main

import (
	"flag"
	"fmt"
	"os"

	"leaguelog/Godeps/_workspace/src/github.com/mattes/migrate/migrate"
)

var direction := flag.String("direction", "", "The db migration direction")

func main() {
	db := os.Getenv("DATABASE_URL")
	if db == "" {
		fmt.Println("Unable to determine the database.")
		os.Exit(1)
	}

	path := os.Getenv("MIGRATION_PATH")
	if db == "" {
		fmt.Println("Unable to determine the migration path.")
		os.Exit(1)
	}
	
	if *direction == "up" {
		fmt.Println("Migrating up!")
		allErrors, ok := migrate.UpSync(db, path)
		if !ok {
			fmt.Println("Error migrating up.")
			printErrors(allErrors)
		}
	} else {
		if *direction == "down" {
			allErrors, ok := migrate.UpSync(db, path)
			if !ok {
				fmt.Println("Error migrating up.")
				printErrors(allErrors)
			}
		} else {
			fmt.Printf("Invalid direction given: %s\n", *direction)
			os.Exit(1)
		}
	}
}

func printErrors(errors []error) {
	for i, e := range errors {
		fmt.Printf("Error %d: %v\n", i, e)
	}
}
