package daosql

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"time"
	"entity"
)

func InsertAllInfos()  {
	log.Println("Dao insert")
	defer func() {
		if x := recover();x!=nil{
			log.Println("inset,err,flag")
			return
		}
	}()
	sqlString := "INSERT INTO power_meter_record (" +
			"p_code," +
			"p_time," +
			"p_zxygdn," +
			"p_zxygdn_1," +
			"p_zxygdn_2," +
			"p_zxygdn_3," +
			"p_zxygdn_4" +
			")" +
				"VALUES(" +
					"?,?,?,?,?," +
					"?,?" +
				")"
	db, err:= sql.Open("mysql", "root:Aa651830@tcp(120.27.227.95:3306)/xiuzhou?charset=utf8")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(sqlString)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for k,v :=range entity.PowerInfoMap{
		_,err :=stmt.Exec(v.PowerMeterId,
			time.Now().Format("20060102150405"),
				v.Zxygdn,
					v.Zxygdn1,
						v.Zxygdn2,
							v.Zxygdn3,
								v.Zxygdn4)
		if err != nil {
			log.Fatal(err)
		}
		delete(entity.PowerInfoMap,k)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

