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

// Insert new snippet into DB
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get snippet by ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Get 10 recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
