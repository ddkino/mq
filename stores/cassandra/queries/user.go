package queries

import "github.com/gocql/gocql"

func UserById(session *gocql.Session, id gocql.UUID) *gocql.Query {
	return session.Query(`
SELECT username, email, phone, password
FROM users
WHERE user_id = ?
LIMIT 1`, id)
}

func UserByUsername(session *gocql.Session, username string) *gocql.Query {
	return session.Query(`
SELECT user_id, username, email, phone, password
FROM users_by_username
WHERE username = ?
LIMIT 1 `, username)
}

func UserByEmail(session *gocql.Session, email string) *gocql.Query {
	return session.Query(`
SELECT user_id, username, email, phone, password
FROM users_by_email
WHERE email = ?
LIMIT 1 `, email)
}


func UserByLoginPassword(session *gocql.Session, username string, password string) *gocql.Query {
	return session.Query(`
SELECT user_id, username, email, phone, password
FROM users
WHERE username = ? and password = ?
LIMIT 1 ALLOW FILTERING`, username, password)
}

func UserByEmailPassword(session *gocql.Session, email string, password string) *gocql.Query {
	return session.Query(`
SELECT user_id, username, email, phone, password
FROM users
WHERE email = ? and password = ?
LIMIT 1 ALLOW FILTERING`, email, password)
}
