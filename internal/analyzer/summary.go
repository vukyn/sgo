package analyzer

import "fmt"

type SummaryStatus struct {
	Status string   `json:"status"`
	Score  int      `json:"score"`
	Notes  []string `json:"notes"`
}

func (r *AnalysisResult) SummaryStatus() {
	score := BaseScore
	notes := []string{}

	// Check project size
	if r.ProjectSize > LargeProjectSizeThreshold {
		notes = append(notes, "Large project size (> 100MB)")
		score -= DeductionLargeProjectSize
	}

	// Check TODOs
	if len(r.TODOs) > HighNumberOfTODOsThreshold {
		notes = append(notes, fmt.Sprintf("High number of unresolved TODOs (%d)", len(r.TODOs)))
		score -= DeductionHighNumberOfTODOs
	} else if len(r.TODOs) > 0 {
		notes = append(notes, fmt.Sprintf("Has %d TODO(s)", len(r.TODOs)))
		score -= DeductionMinorNumberOfTODOs
	}

	// Check secret keys
	if len(r.SecretKeys) > MultiplePotentialSecretKeysThreshold {
		notes = append(notes, fmt.Sprintf("Multiple potential secret keys found (%d)", len(r.SecretKeys)))
		score -= DeductionMultiplePotentialSecretKeys
	} else if len(r.SecretKeys) > 0 {
		notes = append(notes, fmt.Sprintf("Potential secret key found (%d)", len(r.SecretKeys)))
		score -= DeductionMultiplePotentialSecretKeys
	}

	// Check empty files
	emptyFilesCount := len(r.EmptyGoFiles) + len(r.EmptyOtherFiles)
	if emptyFilesCount > MultipleEmptyFilesThreshold {
		notes = append(notes, fmt.Sprintf("Multiple empty files found (%d)", emptyFilesCount))
		score -= DeductionMultipleEmptyFiles
	} else if emptyFilesCount > 0 {
		notes = append(notes, fmt.Sprintf("Empty files found (%d)", emptyFilesCount))
		score -= DeductionMinorEmptyFiles
	}

	// Check warnings
	if len(r.Warnings) > MultipleWarningsThreshold {
		notes = append(notes, fmt.Sprintf("Multiple warnings found (%d)", len(r.Warnings)))
		score -= DeductionMultipleWarnings
	} else if len(r.Warnings) > 0 {
		notes = append(notes, fmt.Sprintf("Warnings found (%d)", len(r.Warnings)))
		score -= DeductionMultipleWarnings
	}

	// Determine status based on score
	var status string
	switch {
	case score >= ScorePerfectThreshold:
		status = "PERFECT"
	case score >= ScoreNeedReviewThreshold:
		status = "NEED REVIEW"
	default:
		status = "BAD"
	}

	r.Summary = SummaryStatus{
		Status: status,
		Score:  score,
		Notes:  notes,
	}
}
