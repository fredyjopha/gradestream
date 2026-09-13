package parsing

import "encoding/json"

type JSONParser struct{}

func (JSONParser) Parse(data []byte) ([]RawRecord, error) {
	var raw []map[string]string
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	records := make([]RawRecord, len(raw))
	for i, fields := range raw {
		records[i] = RawRecord{Fields: fields}
	}
	return records, nil
}
