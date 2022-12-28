package matches

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	Email string
}

type Photo struct {
	User  string
	Photo string
}

type Matched struct {
	Email string
	Photo string
}

// GetMatches get matches
func GetMatches(db *gorm.DB, t *User) ([]Matched, error) {
	matchedUsersQuery := `SELECT email FROM likes
				WHERE email IN (SELECT user FROM likes WHERE email = '` + t.Email + `' AND value = 1)
				AND user = '` + t.Email + `' AND value = 1`
	matchedUsers, errMatched := GetStringsFromQuery(db, matchedUsersQuery)
	if errMatched != nil {
		return nil, errMatched
	}

	var emails []string
	for _, user := range matchedUsers {
		emails = append(emails, "'"+user+"'")
	}
	emailStrings := strings.Join(emails, ", ")

	photosQuery := `SELECT user, photo FROM photos WHERE user IN (` + emailStrings + `) GROUP BY user`
	photos, errPhotos := GetPhotosFromQuery(db, photosQuery)
	if errPhotos != nil {
		return nil, errPhotos
	}

	var result []Matched
	for _, user := range matchedUsers {
		var matchedUser Matched
		for _, photo := range photos {
			if user == photo.User {
				matchedUser = Matched{Email: user, Photo: photo.Photo}
			}
		}

		result = append(result, matchedUser)
	}

	return result, nil
}

func GetStringsFromQuery(db *gorm.DB, query string) ([]string, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var array []string
	for rows.Next() {
		db.ScanRows(rows, &array)
	}

	return array, nil
}

func GetPhotosFromQuery(db *gorm.DB, query string) ([]Photo, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		db.ScanRows(rows, &photos)
	}

	return photos, nil
}
