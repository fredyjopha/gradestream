CREATE TABLE evaluations (
    id          TEXT PRIMARY KEY,
    learner_id  TEXT NOT NULL REFERENCES learners(id),
    cohort_id   TEXT NOT NULL REFERENCES cohorts(id),
    score       DOUBLE PRECISION NOT NULL,
    max_score   DOUBLE PRECISION NOT NULL,
    recorded_at TIMESTAMPTZ NOT NULL,
    UNIQUE (learner_id, cohort_id, recorded_at)
);
