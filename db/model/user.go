package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	id   int64
	name string
	age  int32
}

func QueryRow(db *sql.DB, id int64) (*User, error) {
	sqlQuery := "select * from user where id = ?"
	var u User
	err := db.QueryRow(sqlQuery, id).Scan(&u.id, &u.name, &u.age) // errNoRows
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &u, nil
}

func QuerMultiRows(db *sql.DB, id int64) ([]*User, error) {
	sqlStr := "select * from user where id > ?"
	users := make([]*User, 0, 8)
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err = rows.Scan(&u.id, &u.age, &u.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}

func Insert(db *sql.DB) error {
	sqlStr := "INSERT INTO user(name, age) values(?,?)"
	ret, err := db.Exec(sqlStr, "张三", 18)
	if err != nil {
		fmt.Println("insert error", err)
		return err
	}
	lastId, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("get LastInsertId error", err)
		return err
	}
	fmt.Println(lastId)

	affectRows, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("affect rows error", err)
		return err
	}
	fmt.Println(affectRows)
	return nil
}

func Update(db *sql.DB, age int32, name string) error {
	sqlStr := "UPDATE user set age = ?, name = ?"
	ret, err := db.Exec(sqlStr, age, name)
	if err != nil {
		fmt.Sprintf("update user error:%v", err)
		return err
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Sprintf("rowsAffect error:%v", err)
		return err
	}
	lastId, _ := ret.LastInsertId()
	fmt.Println(num, lastId)
	return nil
}

func Del(db *sql.DB, id int64) error {
	sqlStr := "DELETE from user where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Sprintf("del err: %v", err)
		return err
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Sprintf("rows affect err: %v", err)
		return err
	}
	fmt.Println(num)
	return nil
}
