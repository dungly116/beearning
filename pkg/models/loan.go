package models

import (
	"github.com/dungly116/CRUD/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Loan struct {
	gorm.Model
	user_id   int64  `gorm:""json:"user_id"`
	agency_id int64  `json:"agency_id`
	code      string `json:"code"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Loan{})
}

func (loan *Loan) CreateLoan() *Loan {
	db.NewRecord(loan)
	db.Create(&loan)
	return loan
}

func GetAllLoans() []Loan {
	var Loans []Loan
	db.Find(&Loans)
	return Loans
}

func GetLoanByID(Id int64) (*Loan, *gorm.DB) {
	var getLoan Loan
	db := db.Where("ID=?", Id).Find(&getLoan)
	return &getLoan, db
}

func DeleteLoanByID(Id int64) Loan {
	var loan Loan
	db.Where("ID=?", Id).Delete(loan)
	return loan
}
