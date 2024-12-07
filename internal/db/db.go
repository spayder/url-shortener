package db

import "database/sql"

func CreateTable(sqlite *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	short_url VARCHAR(255) UNIQUE NOT NULL,
		origin_url VARCHAR(500) NOT NULL
    )`

	_, err := sqlite.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CreateURL(db *sql.DB, shortURL string, originURL string) error {
	query := `INSERT INTO urls (short_url, origin_url) VALUES (?, ?)`

	_, err := db.Exec(query, shortURL, originURL)
	if err != nil {
		return err
	}

	return nil
}

func GetOriginURL(db *sql.DB, shortURL string) (string, error) {
	var originURL string
	query := `SELECT origin_url FROM urls WHERE short_url = ? LIMIT 1`
	err := db.QueryRow(query, shortURL).Scan(&originURL)
	if err != nil {
		return "", err
	}
	return originURL, nil
}
