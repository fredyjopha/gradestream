package parsing

import (
	"fmt"
	"strings"
)

func ForFilename(name string) (Parser, error) {
	switch {
	case strings.HasSuffix(name, ".csv"):
		return CSVParser{}, nil
	case strings.HasSuffix(name, ".json"):
		return JSONParser{}, nil
	default:
		return nil, fmt.Errorf("format non supporté: %s", name)
	}
}