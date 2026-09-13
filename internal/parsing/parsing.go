package parsing


type RawRecord struct {
	Fields map[string]string
}

type Parser interface {
	Parse(data []byte) ([]RawRecord, error)
}
