package models

import (
	//"database/sql"

	//"github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	"github.com/kumail/student-data/pkg/config"

	//	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type Student struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	//Gender    string `gorm:""json:"gender"`
}

// func init() {
// 	db := config.Connect()
// }
var db = config.Connect()

func (b *Student) CreateStudent() *Student {

	myuuid := uuid.NewV4().String()
	//INSERT INTO table_name (column1, column2, column3, ...)
	db.Exec("INSERT INTO student (iD, FirstName, LastName ) values($1,$2,$3)", myuuid, b.FirstName, b.LastName)
	return b
}

// func GetAllStudent() []Student {
// 	var Students []Student
// 	db.Find(&Students)
// 	return Students

// }

// func GetStudentById(Id int64) (*Student, *sql.DB) {
// 	var getStudent Student
// 	db := db.Where("ID=?", Id).Find(&getStudent)
// 	return &getStudent, db
// }

// func DeleteStudent(Id int64) Student {
// 	var student Student
// 	db.Where("ID=?", Id).Find(&student)
// 	return student
// }
