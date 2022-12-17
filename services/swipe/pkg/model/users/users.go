package users

import (
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	Email              string
	Firstname          string
	Birthday           string
	About              string
	Photos             pq.StringArray
	Tags               pq.StringArray
	Gender             int
	DistancePreference int
	AgePreference      string
	ShowMe             int
	FilterByTags       int
}

type Photo struct {
	User  string
	Photo string
}

type Tag struct {
	User string
	Tag  string
}

type UserData struct {
	Email     string   `json:"email"`
	Firstname string   `json:"firstname"`
	Birthday  string   `json:"birthday"`
	About     string   `json:"about"`
	Photos    []string `json:"photos"`
	Tags      []string `json:"tags"`
}

// Get users from DB for swiper
func GetUsers(db *gorm.DB, t *User) ([]UserData, error) {
	query := `SELECT * FROM users WHERE email != '` + t.Email +
		`' AND email NOT IN (SELECT user FROM likes WHERE email = '` + t.Email +
		`') AND distance_preference <= ` + strconv.Itoa(t.DistancePreference)

	if t.ShowMe != 2 {
		query += ` AND gender = ` + strconv.Itoa(t.ShowMe)
	}

	minDate, maxDate := GetAgePreferences(t.AgePreference)
	query += ` AND birthday > '` + minDate + `' AND birthday <= '` + maxDate + `'`

	query += ` ORDER BY last_active DESC`

	users, err := GetUsersFromQuery(db, query)
	if err != nil {
		return nil, err
	}

	users, err = getPhotosByUsers(db, users)
	if err != nil {
		return nil, err
	}

	users, err = getTagsByUsers(db, users)
	if err != nil {
		return nil, err
	}

	if t.FilterByTags == 1 {
		users = FilterUsersByTags(t, users)
	}

	return users, nil
}

func GetAgePreferences(agePreference string) (minDate string, maxDate string) {
	agePreference1, _ := strconv.Atoi(agePreference)
	agePreference2, _ := strconv.Atoi(agePreference)

	if strings.Contains(agePreference, "-") {
		agePreference1, _ = strconv.Atoi(agePreference[0:2])
		agePreference2, _ = strconv.Atoi(agePreference[3:5])
	}

	t := time.Now()
	today := t.Format("2006-01-02")
	year, _ := strconv.Atoi(today[0:4])

	minYear := year - agePreference2 - 1
	maxYear := year - agePreference1

	minDate = strings.Replace(today, today[0:4], strconv.Itoa(minYear), -1)
	maxDate = strings.Replace(today, today[0:4], strconv.Itoa(maxYear), -1)

	return minDate, maxDate
}

func GetUsersFromQuery(db *gorm.DB, query string) ([]UserData, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []UserData
	for rows.Next() {
		db.ScanRows(rows, &users)
	}

	return users, nil
}

func getPhotosByUsers(db *gorm.DB, users []UserData) ([]UserData, error) {
	var emails []string
	for _, user := range users {
		emails = append(emails, "'"+user.Email+"'")
	}
	emailStrings := strings.Join(emails, ", ")

	query := "SELECT * FROM photos WHERE user IN (" + emailStrings + ")"

	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		db.ScanRows(rows, &photos)
	}

	for i, user := range users {
		var arr []string

		for _, photo := range photos {
			if user.Email == photo.User {
				arr = append(arr, photo.Photo)
			}
		}

		users[i].Photos = arr
	}

	return users, nil
}

func getTagsByUsers(db *gorm.DB, users []UserData) ([]UserData, error) {
	var emails []string
	for _, user := range users {
		emails = append(emails, "'"+user.Email+"'")
	}
	emailStrings := strings.Join(emails, ", ")

	query := "SELECT * FROM tags WHERE user IN (" + emailStrings + ")"

	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		db.ScanRows(rows, &tags)
	}

	for i, user := range users {
		var arr []string

		for _, tag := range tags {
			if user.Email == tag.User {
				arr = append(arr, tag.Tag)
			}
		}

		users[i].Tags = arr
	}

	return users, nil
}

func FilterUsersByTags(t *User, users []UserData) []UserData {
	var arr []UserData
	for _, user := range users {
		containsTags := contains(t.Tags, user.Tags)

		if containsTags {
			arr = append(arr, user)
		}
	}

	return arr
}

func contains(s []string, e []string) bool {
	for _, a := range s {
		for _, b := range e {
			if a == b {
				return true
			}
		}
	}
	return false
}
