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
	Email  string
	User   string
	IsRead uint
}

type User struct {
	Email          string `json:"email"`
	Firstname      string `json:"firstname"`
	Birthday       string `json:"birthday"`
	ProfilePicture string `json:"profilePicture"`
}

type Matched struct {
	User   User `json:"user"`
	IsRead uint `json:"isRead"`
}

// GetMatches get matches
func GetMatches(db *gorm.DB, t *Email, page string) ([]Matched, error) {
	offset := helpers.GetOffset(page)

	matchedUsersQuery := `SELECT
							T1.email,
							T1.user,
							CASE WHEN T1.email = '` + t.Email + `' THEN
								T1.is_read
							ELSE
								T2.is_read
							END AS is_read
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

	formattedMatchedUsers := formatMatchedUsers(matchedUsers, t.Email)

	emailsString := getUserEmailsString(formattedMatchedUsers)
	usersQuery := `SELECT email, firstname, birthday, profile_picture FROM users WHERE email IN (` + emailsString + `)`

	users, err := GetUsersFromQuery(db, usersQuery)
	if err != nil {
		return nil, err
	}

	var result []Matched
	for _, value := range formattedMatchedUsers {
		for _, user := range users {
			if value.User == user.Email {
				result = append(result, Matched{
					User:   user,
					IsRead: value.IsRead,
				})
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

func formatMatchedUsers(matchedUsers []MatchedUser, email string) []MatchedUser {
	var result []MatchedUser
	for _, value := range matchedUsers {
		if value.Email == email {
			result = append(result, value)
		} else {
			result = append(result, MatchedUser{
				Email:  value.User,
				User:   value.Email,
				IsRead: value.IsRead,
			})
		}
	}
	return result
}

func getUserEmailsString(formattedMatchedUsers []MatchedUser) string {
	var emails []string
	for _, value := range formattedMatchedUsers {
		emails = append(emails, "'"+value.User+"'")
	}

	return strings.Join(emails, ", ")
}

func GetUsersFromQuery(db *gorm.DB, query string) ([]User, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		db.ScanRows(rows, &users)
	}

	return users, nil
}
