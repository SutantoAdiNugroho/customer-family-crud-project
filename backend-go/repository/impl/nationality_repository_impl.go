package impl

import (
	"customer-family-crud-backend/domain/model"
	"customer-family-crud-backend/repository"
	"database/sql"
	"fmt"
)

type NationalityRepositoryImpl struct {
	DB *sql.DB
}

func NewNationalityRepositoryImpl(db *sql.DB) repository.NationalityRepository {
	return &NationalityRepositoryImpl{DB: db}
}

func (r *NationalityRepositoryImpl) GetAllNationalities() ([]*model.Nationality, error) {
	nationalities := []*model.Nationality{}
	query := `select nationality_id, nationality_name, nationality_code from nationality order by nationality_id ASC`

	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Printf("failed to fetch nationalities: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := &model.Nationality{}
		err := rows.Scan(&n.NationalityID, &n.NationalityName, &n.NationalityCode)
		if err != nil {
			fmt.Printf("failed to scan nationalities: %v", err)
			return nil, err
		}
		nationalities = append(nationalities, n)
	}

	return nationalities, nil
}
