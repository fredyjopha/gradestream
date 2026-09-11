package postgres

import (
	"context"

	"github.com/fredyjopha/gradestream/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresEvaluationRepo struct {
	pool *pgxpool.Pool
}

func NewEvaluationRepo(pool *pgxpool.Pool) *PostgresEvaluationRepo {
	return &PostgresEvaluationRepo{pool: pool}
}

const insertEvaluationSQL = `INSERT INTO evaluations (id, learner_id, cohort_id, score, max_score, recorded_at) VALUES ($1, $2, $3, $4, $5, $6)`

func (r *PostgresEvaluationRepo) Save(ctx context.Context, e domain.Evaluation) error {
	_, err := r.pool.Exec(ctx, insertEvaluationSQL, e.ID, e.LearnerID, e.CohortID, e.Score, e.MaxScore, e.RecordedAt)
	return err
}

const selectByCohortSQL = `SELECT id, learner_id, cohort_id, score, max_score, recorded_at FROM evaluations WHERE cohort_id = $1`

func (r *PostgresEvaluationRepo) FindByCohort(ctx context.Context, CohortID string) ([]domain.Evaluation, error) {
	rows, err := r.pool.Query(ctx, selectByCohortSQL, CohortID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.Evaluation

	for rows.Next() {
		var e domain.Evaluation
		if err := rows.Scan(&e.ID, &e.LearnerID, &e.CohortID, &e.Score, &e.MaxScore, &e.RecordedAt); err != nil {
			return nil, err
		}
		results = append(results, e)
	}
	return results, rows.Err()
}
