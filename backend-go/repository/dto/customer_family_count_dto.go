package dto

import "time"

type CustomerWithFamilyCount struct {
	CstID       int       `json:"cst_id"`
	CstName     string    `json:"cst_name"`
	CstDob      time.Time `json:"cst_dob"`
	CstEmail    string    `json:"cst_email"`
	FamilyCount int       `json:"family_count"`
}
