package cmd

import (
	"database/sql"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the psql stress test",
	Long:  `run the psql stress test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"%s dbname=%s sslmode=disable",
			cfg.dbConfig.host, cfg.dbConfig.port, cfg.dbConfig.user, cfg.dbConfig.pass, cfg.dbConfig.name)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Printf("Failed to connect using connection string: %s\n", psqlInfo)
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			fmt.Println("Ping Failed")
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}

		// Create Table
		sqlStatement := `
			CREATE TABLE "testTable" (
  			id SERIAL PRIMARY KEY,
  			age INT,
  			first_name TEXT,
  			last_name TEXT
  			);`
		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}

		var insertCount int
		var updateCount int
		var clearCondition bool
		for !clearCondition {
			switch rand.Intn(2) {
			case 0:
				//INSERT TEST
				sqlStatement := fmt.Sprintf(`
					INSERT INTO "testTable" (age, first_name, last_name)
					VALUES (%v, '%s', '%s')`,
					rand.Intn(50),
					randomdata.FirstName(1),
					randomdata.LastName(),
				)
				fmt.Printf("Executing INSERT STATEMENT \n %s \n", sqlStatement)
				_, err = db.Exec(sqlStatement)
				if err != nil {
					panic(err)
				}
				insertCount++

			case 1:
				//UPDATE TEST
				updateCount++
			}

			// Clear Condition Checker
			if updateCount >= cfg.minUPDATES && insertCount >= cfg.minINSERTS {
				clearCondition = true
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
