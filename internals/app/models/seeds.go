package models

var schema = `
DROP TABLE IF EXISTS user;
CREATE TABLE user (
	user_id    INTEGER PRIMARY KEY,
    first_name VARCHAR(80)  DEFAULT '',
    last_name  VARCHAR(80)  DEFAULT '',
	email      VARCHAR(250) DEFAULT '',
	password   VARCHAR(250) DEFAULT NULL
);
`

func (m *Models) Seeds() {
	m.db.MustExec(schema)
	tx := m.db.MustBegin()
	tx.MustExec("INSERT INTO user (first_name, last_name, email) VALUES ($1, $2, $3)", "Nam", "Dang", "dangnamshuy@gmail.com")
	tx.Commit()
}
