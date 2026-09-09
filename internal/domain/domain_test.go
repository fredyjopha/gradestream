package domain

import "testing"

func TestEvaluationValidate(t *testing.T) {
	valid := Evaluation{ID: "e1", LearnerID: "l1", CohortID: "c1", Score: 15, MaxScore: 20}
	if err := valid.Validate(); err != nil {
		t.Errorf("Attendu valide, obtenu erreur: %v", err)
	}

	invalid := Evaluation{ID: "e2", LearnerID: "l1", CohortID: "c1", Score: 25, MaxScore: 20}
	if err := invalid.Validate(); err == nil {
		t.Error("attendu une erreur pour score hors barème, obtenu nil")
	}
}