package parsing

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
)

type CSVParser struct{}

func (CSVParser) Parse(data []byte) ([]RawRecord, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	header, err := reader.Read()

	if err != nil {
		return nil, fmt.Errorf("lecture en-tête csv: %w", err)
	}

	var records []RawRecord
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("lecture ligne CSV: %w", err)
		}
		records = append(records, rowToRecord(header, row))
	}
	return records, nil
}

func rowToRecord(header, row []string) RawRecord {
	fields := make(map[string]string, len(header))
	for i, col := range header {
		if i < len(row) {
			fields[col] = row[i]
		}
	}
	return RawRecord{Fields: fields}
}