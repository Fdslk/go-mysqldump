package mysqldump

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBBackup() {
	// Open connection to database
	username := "root"
	password := "1qaz2wsx"
	hostname := "localhost"
	port := "3306"
	dbname := "localEmenu"

	dumpDir := "." // you should create this directory

	dumpFilenameFormat := fmt.Sprintf("%s-20060102T150405", dbname) // accepts time layout string and add .sql at the end of file

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname))
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return
	}

	// Register database with mysqldump
	dumper, err := Register(db, dumpDir, dumpFilenameFormat)
	if err != nil {
		fmt.Println("Error registering databse:", err)
		return
	}

	// Dump database to file
	resultFilename, err := dumper.Dump()
	if err != nil {
		fmt.Println("Error dumping:", err)
		return
	}
	fmt.Printf("File is saved to %s", resultFilename)

	// Close dumper and connected database
	dumper.Close()
}
