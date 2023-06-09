package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func conn() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

type Part struct {
	id          int
	PartsNumber string
	Brand       string
	Description sql.NullString
	DateCode    sql.NullString
	LeadTime    sql.NullString
	Qty         int
	Nt          float64
	Usd         float64
	Hk          float64
	Packa       sql.NullString
	UpdateDate  *time.Time
	Telephone   string
	Contact     string
	Supplier    string
}

func getPartByPartsNumberBrandQty(partsNumber, brand string, supplier string) (*Part, error) {
	db, err := conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT id,PARTSNUMBER, BRAND, QTY, SUPPLIER FROM parts WHERE PARTSNUMBER = ? AND BRAND = ? AND SUPPLIER = ? "

	row := db.QueryRow(query, partsNumber, brand, supplier)

	var part Part
	err = row.Scan(
		&part.id,
		&part.PartsNumber,
		&part.Brand,
		&part.Qty,
		&part.Supplier,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// 記錄不存在
			return nil, nil
		}
		return nil, err
	}

	return &part, nil
}

func updateData(part Part, id int) {
	db, err := conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 執行 INSERT 陳述式
	_, err = db.Exec("UPDATE parts SET QTY = ? WHERE id = ? ",
		part.Qty, id)
	if err != nil {
		log.Fatal(err)
	}

}
func insertData(part Part) {
	db, err := conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 執行 INSERT 陳述式
	_, err = db.Exec("INSERT INTO parts (`PartsNumber`, `Brand`, `Description`, `DateCode`, `LeadTime`, `Qty`, `Nt`, `Usd`, `Hk`, `Packa`, `Update`, `Telephone`, `Contact`, `Supplier`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		part.PartsNumber, part.Brand, part.Description, part.DateCode, part.LeadTime, part.Qty, part.Nt, part.Usd, part.Hk, part.Packa, part.UpdateDate, part.Telephone, part.Contact, part.Supplier)
	if err != nil {
		fmt.Println("qweqwe", err)
		log.Fatal(err)
	}

}

// 讀
func readdb() {
	db, err := conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM parts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var partsNumber string
		var brand string
		var description string
		var dateCode string
		var leadTime int
		var qty int
		var nt float64
		var usd float64
		var hk float64
		var packa string
		var updateDate time.Time
		var telephone string
		var contact string
		var supplier string

		if err := rows.Scan(&id, &partsNumber, &brand, &description, &dateCode, &leadTime, &qty, &nt, &usd, &hk, &packa, &updateDate, &telephone, &contact, &supplier); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Parts Number: %s, Brand: %s, Description: %s, Date Code: %s\n", id, partsNumber, brand, description, dateCode)
		// 可以继续打印其他字段的值

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
