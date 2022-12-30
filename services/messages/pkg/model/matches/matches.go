package matches

import (
	"strings"

	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/model/helpers"
	"gorm.io/gorm"
)

type Email struct {
	Email string
}

type Photo struct {
	User  string
	Photo string
}

type MatchedUser struct {
	Email string
	User  string
}

type Matched struct {
	Email          string `json:"email"`
	Firstname      string `json:"firstname"`
	Birthday       string `json:"birthday"`
	ProfilePicture string `json:"profilePicture"`
}

// GetMatches get matches
func GetMatches(db *gorm.DB, t *Email, page string) ([]Matched, error) {
	offset := helpers.GetOffset(page)

	matchedUsersQuery := `SELECT
							T1.email,
							T1.user
						FROM
							likes T1
							INNER JOIN likes T2 ON (T1.email = '` + t.Email + `'
									OR T1.user = '` + t.Email + `')
								AND((T1.email = T2.user)
								AND(T2.email = T1.user))
								AND T1.value = T2.value
								AND T1.value = 1
								AND T1.id > T2.id
							ORDER BY
								T1.id DESC
								LIMIT 10 OFFSET ` + offset

	matchedUsers, errMatched := GetMatchedUsersFromQuery(db, matchedUsersQuery)
	if errMatched != nil {
		return nil, errMatched
	}

	emailStrings := getEmailStringsArray(matchedUsers, t.Email)

	userEmails := getUserEmailsString(emailStrings)
	usersQuery := `SELECT email, firstname, birthday, profile_picture FROM users WHERE email IN (` + userEmails + `)`

	users, err := GetUsersFromQuery(db, usersQuery)
	if err != nil {
		return nil, err
	}

	var result []Matched
	for _, email := range emailStrings {
		for _, user := range users {
			if email == user.Email {
				result = append(result, user)
			}
		}
	}

	return result, nil
}

func GetMatchedUsersFromQuery(db *gorm.DB, query string) ([]MatchedUser, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var array []MatchedUser
	for rows.Next() {
		db.ScanRows(rows, &array)
	}

	return array, nil
}

func getEmailStringsArray(matchedUsers []MatchedUser, email string) []string {
	var emailStrings []string
	for _, value := range matchedUsers {
		if value.Email == email {
			emailStrings = append(emailStrings, value.User)
		} else {
			emailStrings = append(emailStrings, value.Email)
		}
	}
	return emailStrings
}

func getUserEmailsString(emailStrings []string) string {
	var emails []string
	for _, user := range emailStrings {
		emails = append(emails, "'"+user+"'")
	}

	return strings.Join(emails, ", ")
}

func GetUsersFromQuery(db *gorm.DB, query string) ([]Matched, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []Matched
	for rows.Next() {
		db.ScanRows(rows, &users)
	}

	return users, nil
}
