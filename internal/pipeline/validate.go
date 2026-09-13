package pipeline

import (
	"context"
	"fmt"

	"github.com/fredyjopha/gradestream/internal/domain"
	"github.com/fredyjopha/gradestream/internal/parsing"
)

type ValidationResult struct {
	Evaluation domain.Evaluation
	Err        error
	Raw        parsing.RawRecord
}

func Validate(ctx context.Context, in <-chan parsing.RawRecord) <-chan ValidationResult {
	out := make(chan ValidationResult)
	seen := make(map[string]bool)

	go func() {
		defer close(out)
		for r := range in {
			out <- validateOne(r, seen, ctx)
		}
	}()
	return out
}

func validateOne(r parsing.RawRecord, seen map[string]bool, ctx context.Context) ValidationResult {
	e, err := convert(r)
	if err != nil {
		return ValidationResult{Err: err, Raw: r}
	}
	if err := e.Validate(); err != nil {
		return ValidationResult{Err: err, Raw: r}
	}
	key := e.LearnerID + "|" + e.CohortID + "|" + e.RecordedAt.String()
	if seen[key] {
		return ValidationResult{Err: fmt.Errorf("doublon détecté: %s", key), Raw: r}
	}
	seen[key] = true
	return ValidationResult{Evaluation: e}
}
