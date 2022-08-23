package routes

import (
	"github.com/dungly116/CRUD/pkg/controllers"
	"github.com/gorilla/mux"
)

var AddLoanRoutes = func(router *mux.Router) {
	router.HandleFunc("/loan/", controllers.CreateLoan).Methods("POST")
	router.HandleFunc("/loan/", controllers.GetLoan).Methods("GET")
	router.HandleFunc("/loan/{loanID}", controllers.GetLoanByID).Methods("GET")
	router.HandleFunc("loan/{loanID}", controllers.UpdateLoanByID).Methods("PUT")
	router.HandleFunc("loan/{loanID}", controllers.DeleteLoanByID).Methods("DELETE")
}
