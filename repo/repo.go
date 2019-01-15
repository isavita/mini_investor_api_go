package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mini_investor_go"
)

func main() {
	seed()
}

func seed() {
	createDatabase()

	db := dbConn()
	defer db.Close()

	createCampaignsTable(db)
	createInvestmentsTable(db)
	seedCampaigns(db)
	fmt.Println(db)
}

func createDatabase() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	if rows, err := db.Query(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbname)); rows == nil {
		_, err = db.Exec("CREATE DATABASE " + dbname)
		checkErr(err)
	}
}

func createCampaignsTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS campaigns (
		id SERIAL PRIMARY KEY,
		name VARCHAR (255) UNIQUE NOT NULL,
		target_amount_pennies BIGINT NOT NULL,
		multiplier_amount_pennies BIGINT NOT NULL DEFAULT 1,
		raised_amount_pennies BIGINT NOT NULL DEFAULT 0,
		image_url VARCHAR (255),
		sector VARCHAR (255),
		country_name VARCHAR (255),
		created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
		updated_at TIMESTAMP
	)`)
	checkErr(err)
}

func createInvestmentsTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS investments (
		id SERIAL PRIMARY KEY,
		amount_pennies BIGINT NOT NULL,
		campaign_id INTEGER REFERENCES campaigns (id) DEFERRABLE INITIALLY DEFERRED,
		created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
		updated_at TIMESTAMP
	)`)
	checkErr(err)
}

func dbConn() *sql.DB {
	dbinfo := dbInfo()
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func dbInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func seedCampaigns(db *sql.DB) {
	campaigns := []map[string]string{
		map[string]string{
			"name":                      "Starling Bank",
			"image_url":                 "https://www.bankingtech.com/files/2017/10/social-image.png",
			"target_amount_pennies":     "800000000",
			"multiplier_amount_pennies": "1250",
			"raised_amount_pennies":     "200000000",
			"sector":                    "Finance",
			"country_name":              "United Kingdom",
		},
		map[string]string{
			"name":                      "Tesla Roadster",
			"image_url":                 "https://cdn.teslarati.com/wp-content/uploads/2017/04/preisert-tesla-roadster-convertible-rendering.jpg",
			"target_amount_pennies":     "1000000000",
			"multiplier_amount_pennies": "5000",
			"raised_amount_pennies":     "810000000",
			"sector":                    "Automotive Industry",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "Uber",
			"image_url":                 "https://cached.imagescaler.hbpl.co.uk/resize/scaleWidth/614/cached.offlinehbpl.hbpl.co.uk/news/ORP/uber2-2-2016-20160202092201179.jpg",
			"target_amount_pennies":     "100000000",
			"multiplier_amount_pennies": "1100",
			"raised_amount_pennies":     "93999400",
			"sector":                    "Transportation",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "Revolut",
			"image_url":                 "https://moneycheck-9fcd.kxcdn.com/wp-content/uploads/2018/09/revolut-review-1.jpg",
			"target_amount_pennies":     "80000000",
			"multiplier_amount_pennies": "100",
			"raised_amount_pennies":     "96000000",
			"sector":                    "Finance",
			"country_name":              "France",
		},
		map[string]string{
			"name":                      "Prada",
			"image_url":                 "https://i.ebayimg.com/images/g/JzwAAOSwzaJX3z59/s-l300.jpg",
			"target_amount_pennies":     "40000000",
			"multiplier_amount_pennies": "100000",
			"raised_amount_pennies":     "17000000",
			"sector":                    "Fashion",
			"country_name":              "Italy",
		},
		map[string]string{
			"name":                      "Ferrari",
			"image_url":                 "https://static.ferrarinetwork.ferrari.com/images/GATEWAY_Home_1280x960_QJKALQhU.jpg",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "100000",
			"raised_amount_pennies":     "317000000",
			"sector":                    "Automotive Industry",
			"country_name":              "Italy",
		},
		map[string]string{
			"name":                      "BMW",
			"image_url":                 "https://hips.hearstapps.com/amv-prod-cad-assets.s3.amazonaws.com/vdat/submodels/bmw_m8-gran-coupe_bmw-concept-m8-gran-coupe_2018-1532968589970.jpg",
			"target_amount_pennies":     "500000000",
			"multiplier_amount_pennies": "500000",
			"raised_amount_pennies":     "437000000",
			"sector":                    "Automotive Industry",
			"country_name":              "Germany",
		},
		map[string]string{
			"name":                      "Mercedes",
			"image_url":                 "https://upload.wikimedia.org/wikipedia/commons/thumb/b/be/Mercedes-Benz_600_vl_silver_TCE.jpg/1200px-Mercedes-Benz_600_vl_silver_TCE.jpg",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "100000",
			"raised_amount_pennies":     "317000000",
			"sector":                    "Automotive Industry",
			"country_name":              "Germany",
		},
		map[string]string{
			"name":                      "SpaceX",
			"image_url":                 "https://www.nasaspaceflight.com/wp-content/uploads/2018/10/2018-10-22-13_27_15-Window.jpg",
			"target_amount_pennies":     "40000000",
			"multiplier_amount_pennies": "100000",
			"raised_amount_pennies":     "1700000",
			"sector":                    "Space Industry",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "Blue Origin",
			"image_url":                 "https://spacenews.com/wp-content/uploads/2017/04/blueorigin-bezos.jpg",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "1",
			"raised_amount_pennies":     "17000000",
			"sector":                    "Space Industry",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "Rocket Lab",
			"image_url":                 "https://www.rocketlabusa.com/assets/Uploads/_resampled/FillWyI3ODAiLCI0NzAiXQ/RocketLab-F3-ItsBusinessTime-Patch.jpg",
			"target_amount_pennies":     "40000000",
			"multiplier_amount_pennies": "1000000",
			"raised_amount_pennies":     "57000000",
			"sector":                    "Space Industry",
			"country_name":              "New Zealand",
		},
		map[string]string{
			"name":                      "Deliveroo",
			"image_url":                 "https://www.arabianbusiness.com/sites/default/files/styles/full_img/public/images/2018/03/15/Deliveroo%E2%80%99s-Anis-Harb.jpg",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "100",
			"raised_amount_pennies":     "270000000",
			"sector":                    "Ordering Industry",
			"country_name":              "United Kingdom",
		},
		map[string]string{
			"name":                      "Google",
			"image_url":                 "https://ai.google/static/images/share.png",
			"target_amount_pennies":     "400000",
			"multiplier_amount_pennies": "100",
			"raised_amount_pennies":     "170000",
			"sector":                    "Internet",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "Yahoo",
			"image_url":                 "https://i.ytimg.com/vi/5DpspOXs1rM/maxresdefault.jpg",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "500",
			"raised_amount_pennies":     "777000000",
			"sector":                    "Internet",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "QFood",
			"image_url":                 "https://cdn-a.william-reed.com/var/wrbm_gb_food_pharma/storage/images/1/7/6/0/5980671-3-eng-GB/Quorn-invests-150M-to-create-more-than-300-jobs_wrbm_large.jpg",
			"target_amount_pennies":     "100000000",
			"multiplier_amount_pennies": "10000",
			"raised_amount_pennies":     "110000000",
			"sector":                    "Food Industry",
			"country_name":              "United Kingdom",
		},
		map[string]string{
			"name":                      "Tesco",
			"image_url":                 "https://www.personneltoday.com/wp-content/uploads/sites/8/2018/07/tesco.jpg",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "100",
			"raised_amount_pennies":     "221000000",
			"sector":                    "Food Industry",
			"country_name":              "United Kingdom",
		},
		map[string]string{
			"name":                      "Nike",
			"image_url":                 "http://content.nike.com/content/dam/one-nike/globalAssets/social_media_images/nike_swoosh_logo_black.png",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "100000",
			"raised_amount_pennies":     "170000000",
			"sector":                    "Footwear",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "Adidas",
			"image_url":                 "https://nssdata.s3.amazonaws.com/images/galleries/12347/cover/cover-adidas-nss-magazine_v2.jpg",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "1000",
			"raised_amount_pennies":     "373000000",
			"sector":                    "Footwear",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "Shell",
			"image_url":                 "https://www.shell.com/investors/_jcr_content/pagePromo/image.img.960.jpeg/1457612146464/shell-service-station.jpeg",
			"target_amount_pennies":     "40000000",
			"multiplier_amount_pennies": "10000",
			"raised_amount_pennies":     "17000000",
			"sector":                    "Oil Industry",
			"country_name":              "USA",
		},
		map[string]string{
			"name":                      "Hollywood",
			"image_url":                 "https://www.history.com/.image/t_share/MTU3ODc5MDg2Njk4OTMxOTM1/hollywood-sign-3.jpg",
			"target_amount_pennies":     "400000000",
			"multiplier_amount_pennies": "10000",
			"raised_amount_pennies":     "379000000",
			"sector":                    "Entertainment Industry",
			"country_name":              "USA",
		},
	}

	for _, campaign := range campaigns {
		stmt, err := db.Prepare(`INSERT INTO campaigns(name,image_url,target_amount_pennies,multiplier_amount_pennies,raised_amount_pennies,sector,country_name)
			VALUES($1,$2,$3,$4,$5,$6,$7)`)
		checkErr(err)

		_, err = stmt.Exec(campaign["name"], campaign["image_url"], campaign["target_amount_pennies"],
			campaign["multiplier_amount_pennies"], campaign["raised_amount_pennies"], campaign["sector"], campaign["country_name"])
		checkErr(err)
	}

	fmt.Println(len(campaigns), "campaigns created")
}
