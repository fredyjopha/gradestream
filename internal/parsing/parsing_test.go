package parsing

import "testing"

func TestForFilename(t *testing.T) {
	cases := []struct {
		name        string
		filename    string
		expectError bool
	}{
		{"csv reconnu", "batch.csv", false},
		{"json reconnu", "batch.json", false},
		{"format inconnu", "batch.xml", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := ForFilename(c.filename)
			if (err != nil) != c.expectError {
				t.Errorf("%s: erreur attendue=%v, obtenu err=%v", c.filename, c.expectError, err)
			}
		})
	}
}

func TestCSVParserParsesRows(t *testing.T) {
	data := []byte("id,note\n1,15\n2,18")
	records, err := CSVParser{}.Parse(data)
	if err != nil {
		t.Fatalf("erreur inattendue: %v", err)
	}
	if len(records) != 2 {
		t.Fatalf("attendu 2 enregistrements, obtenu %d", len(records))
	}
	if records[0].Fields["note"] != "15" {
		t.Errorf("attendu note=15, obtenu %s", records[0].Fields["note"])
	}
}
