package routes

import (
	"github.com/gorilla/mux"
	"github.com/kumail/student-data/pkg/controllers"
)

var RegisterStudentDataRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")
	// router.HandleFunc("/student/{studentId}/", controllers.GetStudentById).Methods("GET")
	// router.HandleFunc("/students/", controllers.GetAllStudent).Methods("GET")
	// router.HandleFunc("/students/{studentId}/", controllers.UpdateStudent).Methods("PUT")
	// router.HandleFunc("/students/{studentId}/", controllers.DeleteStudent).Methods("DELETE")
}
