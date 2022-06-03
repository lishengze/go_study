package BaseData

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// func init() {
// 	db, _ = sql.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/test?charset=utf8")
// 	db.SetMaxOpenConns(2000)
// 	db.SetMaxIdleConns(1000)
// 	db.SetConnMaxLifetime(time.Minute * 60)
// 	db.Ping()
// 	createTable()
// 	insert()
// }

// func main() {
// 	startHttpServer()
// }

func createTable() {
	db, err := sql.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	table := `CREATE TABLE IF NOT EXISTS test.user (
 user_id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
 user_name VARCHAR(45) NOT NULL COMMENT '用户名称',
 user_age TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户年龄',
 user_sex TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
 PRIMARY KEY (user_id))
 ENGINE = InnoDB
 AUTO_INCREMENT = 1
 DEFAULT CHARACTER SET = utf8
 COLLATE = utf8_general_ci
 COMMENT = '用户表'`
	if _, err := db.Exec(table); err != nil {
		checkErr(err)
	}
}

func insert() {
	stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec("tony", 20, 1)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func queryToMap() []map[string]string {
	var records []map[string]string
	rows, err := db.Query("SELECT * FROM user")
	defer rows.Close()
	checkErr(err)
	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		records = append(records, record)
	}
	return records
}

func startHttpServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func pool(w http.ResponseWriter, r *http.Request) {
	records := queryToMap()
	fmt.Println(records)
	fmt.Fprintln(w, "finish")
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// "host" : "tcp://127.0.0.1",
// "port" : 3306,
// "usr" : "bcts",
// "pwd" : "bcts",
// "schema": "market"
func TestTables() {

	db, err := sql.Open("mysql", "bcts:bcts@tcp(127.0.0.1:3306)/market")

	if err != nil {
		fmt.Printf("err: %+v", err)
		return
	}

	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(time.Minute * 60)

	// query_str := "select * from kline_FTX_BTC_USDT where time=1648253087247782804;"
	query_str := "show tables;"

	rows, err := db.Query(query_str)

	fmt.Printf("rows: %+v\nerr: %+v\n", rows, err)

	columns, _ := rows.Columns()
	fmt.Printf("rows.Columns: %+v \n", columns)

	var records []map[string]string
	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)

		fmt.Printf("scanArgs : %+v \n", scanArgs)

		record := make(map[string]string)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		records = append(records, record)
	}

	fmt.Printf("records: %+v \n", records)
}
