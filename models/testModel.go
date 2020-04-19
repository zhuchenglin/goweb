package models

import "fmt"

type TestModel struct {
	BaseModel
}

func (test TestModel) Test() {
	db := test.getDbConnect()
	if db == nil {
		fmt.Println("db 连接失败")
	}

	rows, err := db.Raw("select * from goods where id = ?", 158).Rows() // (*sql.Rows, error)

	if err != nil {
		fmt.Println(err)
		return
	}
	var a, b, c, d, e, f string
	for rows.Next() {
		rows.Scan(&a, &b, &c, &d, &e, &f)
		fmt.Println(a, b, c, d, e, f)
	}

	defer rows.Close()
	defer test.close()

}
