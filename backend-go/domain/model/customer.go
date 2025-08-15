package model

import "time"

type Customer struct {
	CstID         int       `json:"cst_id"`
	NationalityID int       `json:"nationality_id"`
	CstName       string    `json:"cst_name"`
	CstDob        time.Time `json:"cst_dob"`
	CstPhoneNum   string    `json:"cst_phoneNum"`
	CstEmail      string    `json:"cst_email"`
}
