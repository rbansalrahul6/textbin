package mysql

import (
	"database/sql"

	"github.com/rbansalrahul6/textbin/pkg/models"
)

// Think of SnippetModel as DAO object
type SnippetModel struct {
	DB *sql.DB
}

/* DB methods for our models*/

// Insert new snippet into DB-> returns id
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(query, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Get snippet by ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(query, id)
	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// Get 10 recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}
	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
