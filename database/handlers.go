package database

import (
	"awesome/settings"
	"awesome/types"
	"database/sql"
)

func GetAllApiKeysFromDatabase() ([]types.ApiKeys, error) {
	db, err := sql.Open("mysql", settings.Database)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(getAllApiKeysQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apikeys []types.ApiKeys
	for rows.Next() {
		var ak types.ApiKeys
		err := rows.Scan(&ak.ApiKey)
		if err != nil {
			return nil, err
		}
		apikeys = append(apikeys, ak)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return apikeys, nil
}

func GetProfilesFromDatabase() ([]types.UserProfile, error) {
	db, err := sql.Open("mysql", settings.Database)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(GetAllProfilesQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []types.UserProfile
	for rows.Next() {
		var up types.UserProfile
		err := rows.Scan(&up.Username, &up.ID, &up.FirstName, &up.LastName, &up.City, &up.School)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, up)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}

func GetProfileFromDatabase(username string) (*types.UserProfile, error) {
	db, err := sql.Open("mysql", settings.Database)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var profile types.UserProfile
	err = db.QueryRow(GetProfileQuery(username)).Scan(
		&profile.Username,
		&profile.ID,
		&profile.FirstName,
		&profile.LastName,
		&profile.City,
		&profile.School,
	)

	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, fmt.Errorf("record not found")
		//}
		return nil, err
	}

	return &profile, nil
}
