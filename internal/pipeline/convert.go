package pipeline

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fredyjopha/gradestream/internal/domain"
	"github.com/fredyjopha/gradestream/internal/parsing"
)

func convert(r parsing.RawRecord) (domain.Evaluation, error) {
	score, err := strconv.ParseFloat(r.Fields["score"], 64)
	if err != nil {
		return domain.Evaluation{}, fmt.Errorf("score invalide: %w", err)
	}
	maxScore, err := strconv.ParseFloat(r.Fields["max_score"], 64)
	if err != nil {
		return domain.Evaluation{}, fmt.Errorf("max_score invalide: %w", err)
	}
	recordedAt, err := time.Parse(time.RFC3339, r.Fields["recorded_at"])
	if err != nil {
		return domain.Evaluation{}, fmt.Errorf("recorded_at invalide: %w", err)
	}
	return domain.Evaluation{
		ID: r.Fields["id"], LearnerID: r.Fields["learner_id"], CohortID: r.Fields["cohort_id"],
		Score: score, MaxScore: maxScore, RecordedAt: recordedAt,
	}, nil
}
