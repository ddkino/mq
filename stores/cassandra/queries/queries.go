package queries

import "github.com/gocql/gocql"
import (
	"time"
)

//var securityAccess := []string{"public", "public_only_cha", "public_strict_cha"}

func ChaList(session gocql.Session, security string) (*gocql.Iter) {
	return session.Query(`
SELECT id, c, h, c_alt, h_alt, a_id, datetime_created_at, datetime_updated_at
FROM cha_by_id_and_security
WHERE security_access = ?
LIMIT ?`, security, "2").Iter()
}

func ChaLastList(session gocql.Session, created time.Time, security string) *gocql.Iter {
	if security == "" {
		security = "public"
	}
	const cqlLimit string = " LIMIT ?"
	var cqlCreated string
	if &created != nil {
		cqlCreated = " AND datetime_created_at >= '" + created.Format(time.RFC3339) + "'"
	} else {
		cqlCreated = " AND datetime_created_at >= '" + time.Now().AddDate(0, 0, -1).Format(time.RFC3339) + "'"
	}
	cqlQuery := `
	SELECT id, c, h, c_alt, h_alt, a_id, datetime_created_at, datetime_updated_at
	FROM cha_by_id_and_security
	WHERE security_access = ? `
	return session.Query(cqlQuery+cqlCreated+cqlLimit+" ALLOW FILTERING", security, "2").Iter()
}