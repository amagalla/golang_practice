package appointments

import "golang_practice/pckg/db"

func InsertPhysiciansData(firstName, lastName string) error {
	db := db.GetDB()

	resp, err := db.Prepare("INSERT INTO physicians (physician_first_name, physician_last_name) VALUES (?, ?)")

	if err != nil {
		return err
	}

	defer resp.Close()

	_, err = resp.Exec(firstName, lastName)

	if err != nil {
		return err
	}

	return nil
}

// type PhysicianListResponse struct {
// 	PhysicianList []Physicians `json:"results"`
// }

type Physicians struct {
	PhysicianID   uint64 `json:"physician_id"`
	FirstName string `json:"physician_first_name"`
	LastName string `json:"physician_last_name"`
}

func GetPhysiciansList() ([]Physicians, error) {
	db := db.GetDB()

	rows, err := db.Query("SELECT * FROM physicians")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var physiciansList []Physicians

	for rows.Next() {
		var physician Physicians
		if err := rows.Scan(&physician.PhysicianID, &physician.FirstName, &physician.LastName); err != nil {
			return nil, err
		}
		physiciansList = append(physiciansList, physician)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return physiciansList, nil
}
