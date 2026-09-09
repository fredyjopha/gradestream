package domain

import (
	"fmt"
	"time"
)

type Learner struct {
	ID   string
	Name string
}

type Cohort struct {
	ID   string
	Name string
}

type Evaluation struct {
	ID         string
	LearnerID  string
	CohortID   string
	Score      float64
	MaxScore   float64
	RecordedAt time.Time
}

func (e Evaluation) Validate() error {
	if e.ID == "" || e.LearnerID == "" || e.CohortID == "" {
		return fmt.Errorf("Champs obligatoires manquants!")
	}
	if e.Score < 0 || e.Score > e.MaxScore {
		return fmt.Errorf("Score %.2f est hors barème (0 - %.2f)", e.Score, e.MaxScore)
	}
	return nil
}
