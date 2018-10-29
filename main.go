//go:generate go run assets-generator.go

package main

import (
	fp "path/filepath"

	"github.com/viknash/shiori/cmd"
	dt "github.com/viknash/shiori/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var dataDir = "."

func main() {
	// Open database
	dbPath := fp.Join(dataDir, "shiori.db")
	sqliteDB, err := dt.OpenSQLiteDatabase(dbPath)
	checkError(err)

	// Start cmd
	shioriCmd := cmd.NewShioriCmd(sqliteDB, dataDir)
	if err := shioriCmd.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
