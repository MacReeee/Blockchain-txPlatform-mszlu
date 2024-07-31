package tools

import "time"

func ISO(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
