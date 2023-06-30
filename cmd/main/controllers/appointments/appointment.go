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