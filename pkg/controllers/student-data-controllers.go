package controllers

import (
	"encoding/json"
	//"fmt"
	"net/http"
	//"strconv"
	//"github.com/gorilla/mux"
	"github.com/kumail/student-data/pkg/models"
	"github.com/kumail/student-data/pkg/utils"
)

// func GetStudent(w http.ResponseWriter, r *http.Request){
// 	newStudents := models.GetAllStudent()
// 	res, _ := json.Marshal(newStudents)
// 	w.Header().Set("Content_Type","application/json")
// 	w.WriteHeader((http.StatusOK))
// 	w.Write(res)

// }

// func  GetStudentById(w http.ResponseWriter, r *http.Request){
// 	vars := mux.Vars(r)
// 	studentId := vars["studentId"]
// 	ID, err := strconv.ParseInt(studentId,0,0)
// 	if err != nil {
// 		fmt.Println("error while parsing")

// 	}
// 	studentDetails, _ := models.GetStudentById(ID)
// 	res, _ := json.Marshal(studentDetails)
// 	w.Header().Set("Content-type","application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	CreateStudent := &models.Student{}
	utils.ParseBody(r, CreateStudent)
	b := CreateStudent.CreateStudent()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// func DeleteStudent(w http.ResponseWriter, r *http.Request){
// 	vars := mux.Vars(r)
// 	studentId := vars["studentId"]
// 	ID, err := strconv.ParseInt(studentId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	student := models.DeleteStudent(ID)
// 	res, _ := json.Marshal(student)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateStudent(w http.ResponseWriter, r *http.Request){
// 	var updateStudent = &models.Student{}
// 	utils.ParseBody(r, updateStudent)
// 	vars := mux.Vars(r)
// 	studentId := vars["studentId"]
// 	ID, err := strconv.ParseInt(studentId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	studentDetails, db := models.GetStudentById(ID)
// 	if updateStudent.FirstName != ""{
// 		studentDetails.FirstName = updateStudent.FirstName
// 	}

// 	if updateStudent.LastName != ""{
// 		studentDetails.FirstName = updateStudent.LastName
// 	}
// 	res, _ := json.Marshal(studentDetails)
// 	db(&studentDetails)
// 	w.Header().Set("Content-Type", "applicatio/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }
