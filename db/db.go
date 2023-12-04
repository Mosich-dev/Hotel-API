package db

const DBNAME = "hotel-api"

func IsObjedID(id string) bool {
	if len(id) == 24 {
		return true
	} else {
		return false
	}
}
