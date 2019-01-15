package main

import (
	"database/sql"
	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	_ "strconv"
)

type Campaign struct {
	Id                      int    `json:"id"`
	Name                    string `json:"name"`
	TargetAmountPennies     int    `json:"targetAmount"`
	MultiplierAmountPennies int    `json:"multiplierAmount"`
	RaisedAmountPennies     int    `json:"raisedAmount"`
	ImageUrl                string `json:"imageUrl"`
	Sector                  string `json:"sector"`
	CountryName             string `json:"countryName"`
}

type Campaigns struct {
	Campaigns []Campaign `json:"campaigns"`
}

var (
	campaigns = map[int]*Campaign{}
	seq       = 1
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mini_investor_demo"
)

func getCampaigns(c echo.Context) error {
	return c.JSON(http.StatusOK, "It works!!!")
}

/*
func getCampaign(c echo.Context) error {
}
*/

func createCampaignsTable(db DB) {
	db.QueryRow(`CREATE TABLE IF NOT EXISTS campaigns(
		id int(11) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
		name varchar(255) NOT NULL COMMENT 'employee name',
		target_amount_pennies double NOT NULL COMMENT 'employee salary',
		employee_age int(11) NOT NULL COMMENT 'employee age',
		PRIMARY KEY (id)
	) ENGINE=InnoDB  DEFAULT CHARSET=latin1 COMMENT='datatable demo table' AUTO_INCREMENT=158
	`)
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	createCampaignsTable(db)

	e := echo.New()

	e.GET("/api/campaigns", getCampaigns)
	//	e.GET("/api/campaigns/:id", getCampaign)
	e.Logger.Fatal(e.Start(":4000"))
}
