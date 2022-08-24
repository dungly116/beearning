package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dungly116/CRUD/pkg/models"
	"github.com/dungly116/CRUD/pkg/utils"
	"github.com/gorilla/mux"
)

func GetLoan(w http.ResponseWriter, r *http.Request) {
	newLoan := models.GetAllLoans()
	res, _ := json.Marshal(newLoan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLoanByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanId := vars["loanId"]
	ID, err := strconv.ParseInt(loanId, 0, 0)
	if err != nil {
		fmt.Println("error parsing")
	}
	loanDetail, _ := models.GetLoanByID(ID)
	res, _ := json.Marshal(loanDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateLoan(w http.ResponseWriter, r *http.Request) {
	CreateLoan := &models.Loan{}
	utils.ParseBody(r, CreateLoan)
	b := CreateLoan.CreateLoan()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteLoanByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanId := vars["loanId"]
	ID, err := strconv.ParseInt(loanId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	loan := models.DeleteLoanByID(ID)
	res, _ := json.Marshal(loan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateLoanByID(w http.ResponseWriter, r *http.Request) {
	var updateLoan = &models.Loan{}
	utils.ParseBody(r, updateLoan)
	vars := mux.Vars(r)
	updateLoanId := vars["loanId"]
	ID, err := strconv.ParseInt(updateLoanId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	loanDetail, db := models.GetLoanByID(ID)

	if updateLoan.Message_code != "" {
		loanDetail.Message_code = updateLoan.Message_code
	}

	db.Save(&loanDetail)
	res, _ := json.Marshal(loanDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
