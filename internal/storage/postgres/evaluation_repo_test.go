package postgres

import (
	"strings"
	"testing"
)

func TestInsertEvaluationSQLHasSixParams(t *testing.T) {
	count := strings.Count(insertEvaluationSQL, "$")
	if count != 6 {
		t.Errorf("attendu 6 paramètres, obtenu %d", count)
	}
}