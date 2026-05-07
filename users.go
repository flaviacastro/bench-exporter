package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// getUserCountonSite returns the count for allusers, activeusers and system managers
// uses direct SQL queries.
func getUserCountonSite(siteName string) (float64, float64, float64) {
	allUserQuery := `SELECT COUNT(name) FROM tabUser WHERE enabled=1 AND user_type != 'Website User' AND name NOT IN ("Administrator","Guest");`
	activeUserQuery := `SELECT COUNT(*) FROM tabUser WHERE enabled=1 AND user_type != 'Website User' AND name NOT IN ("Administrator","Guest") AND hour(timediff(now(), last_active)) < 72;`
	systemManagerQuery := "SELECT DISTINCT COUNT(name) FROM `tabUser` AS p WHERE enabled=1 AND docstatus<2 AND name NOT IN (\"Administrator\",\"Guest\") AND EXISTS(SELECT * FROM `tabHas Role` AS ur WHERE ur.parent=p.name AND ur.role=\"System Manager\");"
	db, err := sql.Open("mysql", generateDbURI(siteName))

	if err != nil {
		log.Println("Database Connection Error ", err)
		return 0, 0, 0
	}

	defer db.Close()

	allUserQuerystmt, err := db.Prepare(allUserQuery)
	if err != nil {
		log.Println("Preparing allUserQuery failed ", err)
		return 0, 0, 0
	}

	defer allUserQuerystmt.Close()
	activeUserQuerystmt, err := db.Prepare(activeUserQuery)

	if err != nil {
		log.Println("Preparing activeUserQuery failed ", err)
		return 0, 0, 0
	}

	defer activeUserQuerystmt.Close()
	systemManagerQuerystmt, err := db.Prepare(systemManagerQuery)

	if err != nil {
		log.Println("Preparing systemManagerQuery failed ", err)
		return 0, 0, 0
	}

	defer systemManagerQuerystmt.Close()
	var activeUsers, systemManagers, allUsers float64

	if err := allUserQuerystmt.QueryRow().Scan(&allUsers); err != nil {
		log.Println("AllUserQuery failed ", err)
	}

	if err := activeUserQuerystmt.QueryRow().Scan(&activeUsers); err != nil {
		log.Println("activeUserQuery failed", err)
	}

	if err := systemManagerQuerystmt.QueryRow().Scan(&systemManagers); err != nil {
		log.Println("systemManagerQuery failed ", err)
	}

	return allUsers, activeUsers, systemManagers
}
