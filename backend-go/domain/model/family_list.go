package model

type FamilyList struct {
	FlID       int    `json:"fl_id"`
	CstID      int    `json:"cst_id"`
	FlRelation string `json:"fl_relation"`
	FlName     string `json:"fl_name"`
	FlDob      string `json:"fl_dob"`
}
