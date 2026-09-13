package pipeline

import (
	"context"
	"testing"
	"time"

	"github.com/fredyjopha/gradestream/internal/parsing"
)

func makeRecord(id, learner, cohort, score, max, at string) parsing.RawRecord {
	return parsing.RawRecord{Fields: map[string]string{
		"id": id, "learner_id": learner, "cohort_id": cohort,
		"score": score, "max_score": max, "recorded_at": at,
	}}
}

func TestValidateRejectsAndAccepts(t *testing.T) {
	now := time.Now().Format(time.RFC3339)
	cases := []struct {
		name      string
		record    parsing.RawRecord
		expectErr bool
	}{
		{"valide", makeRecord("e1", "l1", "c1", "15", "20", now), false},
		{"score hors barème", makeRecord("e2", "l1", "c1", "25", "20", now), true},
		{"score non numérique", makeRecord("e3", "l1", "c1", "abc", "20", now), true},
	}

	in := make(chan parsing.RawRecord, len(cases))
	for _, c := range cases {
		in <- c.record
	}
	close(in)

	out := Validate(context.Background(), in)
	for _, c := range cases {
		res := <-out
		if (res.Err != nil) != c.expectErr {
			t.Errorf("%s: erreur attendue=%v, obtenu err=%v", c.name, c.expectErr, res.Err)
		}
	}
}
