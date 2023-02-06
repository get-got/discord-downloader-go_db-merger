package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/fatih/color"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(color.Output)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var databasePaths []string

EnterDmNum:
	log.Print(color.HiYellowString("How many databases are you merging? (please enter a number): "))
	inputNum, _ := reader.ReadString('\n')
	inputNum = strings.ReplaceAll(inputNum, "\n", "")
	inputNum = strings.ReplaceAll(inputNum, "\r", "")
	if numDbs, err := strconv.Atoi(inputNum); err != nil {
		log.Println(color.HiRedString("That's not a number dork"))
		goto EnterDmNum
	} else {
		for i := 1; i <= numDbs; i++ {
		EnterNewDb:
			log.Print(color.HiYellowString("[DATABASE %d] Enter path to database: ", i))
			inputDb, _ := reader.ReadString('\n')
			inputDb = strings.ReplaceAll(inputDb, "\n", "")
			inputDb = strings.ReplaceAll(inputDb, "\r", "")
			test, err := db.OpenDB(inputDb)
			if err != nil {
				log.Println(color.HiRedString("[DATABASE %d] Error connecting to database: %s", i, err))
				goto EnterNewDb
			} else {
				test.Close()
				databasePaths = append(databasePaths, inputDb)
				log.Println(color.HiGreenString("[DATABASE %d] Connected to database, adding to merge list...", i))
			}
		}

		var newDB *db.DB

		type merger struct {
			path string
			db   *db.DB
		}

		if len(databasePaths) == 0 {
			log.Println(color.HiRedString("No valid databases to merge..."))
		} else {
			//create
			log.Println(color.YellowString("Creating database, please wait..."))
			if err := newDB.Create("Downloads"); err != nil {
				log.Println(color.HiRedString("Error while trying to create database: %s", err))
				return
			}
			log.Println(color.HiYellowString("Created new database..."))
			log.Println(color.YellowString("Indexing database, please wait..."))
			if err := newDB.Use("Downloads").Index([]string{"URL"}); err != nil {
				log.Println(color.HiRedString("Unable to create database index for URL: %s", err))
				return
			}
			if err := newDB.Use("Downloads").Index([]string{"ChannelID"}); err != nil {
				log.Println(color.HiRedString("Unable to create database index for ChannelID: %s", err))
				return
			}
			if err := newDB.Use("Downloads").Index([]string{"UserID"}); err != nil {
				log.Println(color.HiRedString("Unable to create database index for UserID: %s", err))
				return
			}
			log.Println(color.HiYellowString("Created new indexes..."))

			//
			log.Println(color.HiGreenString("Merging..."))
			for _, dbPath := range databasePaths {

			}
		}
		newDB.Close()
	}
	log.Println(color.HiGreenString("Done!"))
}
