package models

import (
	"time"

	"github.com/dungly116/CRUD/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Loan struct {
	gorm.Model
	User_id         uint64     `json:"user_id"`
	Agency_id       uint64     `json:"agency_id"`
	Code            string     `json:"code"`
	Amount          uint64     `json:"amount"`
	Tip             uint64     `json:"tip"`
	Money_paid      uint64     `json:"money_paid"`
	Money_punish    uint64     `json:"money_punish"`
	Loan_purpose    string     `json:"loan_purpose"`
	Due_date        *time.Time `json:"due_date"`
	Due_at          *time.Time `json:"due_at"`
	Message_code    string     `json:"message_code"`
	Contract_number string     `json:"contract_number"`
	Loan_consumer   uint64     `json:"loan_consumer"`
	Create_at       *time.Time `json:"create_at"`
	Err_code        *time.Time `json:"err_code"`
	Hash            string     `json:"hash"`
	Request_id      string     `json:"request_id"`
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
