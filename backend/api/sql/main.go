// following: https://go.dev/doc/tutorial/web-service-gin
// following: https://go.dev/doc/tutorial/database-access
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"strconv"
)

// jobapp represents data about a job application.
type JobApp struct {
	ID      string `json:"id"`
	Status 	string `json:"status"`
	Title   string `json:"title"`
	Company string `json:"company"`
	Url     string `json:"url"`
	Source  string `json:"source"`
}

var db *sql.DB

func main() {
	initDbConnection()
	db.Ping()
	router := gin.Default()
	router.GET("/jobapps", getJobAppsApi)
	router.POST("/jobapps/create", createJobApp)
	router.POST("/jobapps/update", updateJobApp)
	router.POST("/jobapps/delete", deleteJobApp)

	router.Run("localhost:8080")
}

func initDbConnection () {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "careers",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
			log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
			log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

// I'm not sure if this is the best way to go about doing this,
// but I thought there was a possibility, in this specific case,
// where we'd want to call the SQL portion of this separately
// from the API portion.
func getJobAppsSql() ([]JobApp, error) {
	var jobApps []JobApp
	rows, err := db.Query("SELECT * FROM jobapps")
	if err != nil {
			return nil, fmt.Errorf("getJobAppsSql: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
			var ja JobApp
			if err := rows.Scan(&ja.ID, &ja.Status, &ja.Title, &ja.Company, &ja.Url, &ja.Source); err != nil {
					return nil, fmt.Errorf("getJobAppsSql: %v", err)
			}
			jobApps = append(jobApps, ja)
	}
	if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("getJobAppsSql %v", err)
	}
	return jobApps, nil
}

func getJobAppsApi(c *gin.Context) {
	jobApps, err := getJobAppsSql()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, fmt.Errorf("getJobAppsSql %v", err))
	}
	c.IndentedJSON(http.StatusOK, jobApps)
}

func createJobApp(c *gin.Context) {
	var newJobApp JobApp
	if err := c.BindJSON(&newJobApp); err != nil {
		return
	}
	result, err := db.Exec("INSERT INTO jobapps (status, title, company, url, source) VALUES (?, ?, ?, ?, ?)", newJobApp.Status, newJobApp.Title, newJobApp.Company, newJobApp.Url, newJobApp.Source)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	newJobApp.ID = strconv.FormatInt(int64(id), 10)
	c.IndentedJSON(http.StatusCreated, newJobApp)
}

func updateJobApp(c *gin.Context) {
	var targetJobApp JobApp
	if err := c.BindJSON(&targetJobApp); err != nil {
		return
	}
	_, err := db.Exec("UPDATE jobapps SET status = ?, title = ?, company = ?, url = ?, source = ? WHERE ID = ?", targetJobApp.Status, targetJobApp.Title, targetJobApp.Company, targetJobApp.Url, targetJobApp.Source, targetJobApp.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusAccepted, targetJobApp)
}

func deleteJobApp(c *gin.Context) {
	var deleteJobApp JobApp
	if err := c.BindJSON(&deleteJobApp); err != nil {
		return
	}
	_, err := db.Exec("DELETE FROM jobapps WHERE ID = ?", deleteJobApp.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusAccepted, deleteJobApp)
}