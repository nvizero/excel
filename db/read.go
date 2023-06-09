package db

import (
	"database/sql"
	"excl/util"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func Read() {
	f, err := excelize.OpenFile("aa1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 獲取 Sheet1 上所有存儲格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, row := range rows {
		if i != 0 {
			qty, _ := strconv.Atoi(row[5])
			nt, _ := strconv.ParseFloat(row[6], 64)
			usd, _ := strconv.ParseFloat(row[7], 64)
			hk, _ := strconv.ParseFloat(row[8], 64)

			updateDate, err := parseDate(row[10])
			if err != nil {
				log.Fatal(err)
			}
			part := Part{
				PartsNumber: row[0],
				Brand:       row[1],
				Description: parseNullString(row[2]),
				DateCode:    parseNullString(row[3]),
				LeadTime:    parseNullString(row[4]),
				Qty:         qty,
				Nt:          nt,
				Usd:         usd,
				Hk:          hk,
				Packa:       parseNullString(row[9]),
				UpdateDate:  updateDate,
				Telephone:   row[11],
				Contact:     row[12],
				Supplier:    row[13],
			}
			util.Info(row[0], row[1])
			finsert, err := GetPartByPartsNumberBrandQty(row[0], row[1], row[13])
			if err != nil {
				log.Fatal(err)
			}
			if finsert == nil {
				InsertData(part)
			} else {
				if finsert.Qty != qty {
					fmt.Println("Part already exists:", finsert)
					UpdateData(part, finsert.id)
				}
			}

		}
	}
}

func parseDate(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func parseNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
