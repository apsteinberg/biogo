package hmmer

type QueryResult struct {
	Accession   string
	Description string
	Id          string
}

type Hit struct {
	Accession   string  // hit accession
	Bias        float64 // hit-level bias
	BitScore    float64 // hit-level score
	Description string  // hit sequence description
	EValue      float64
	Score       float64
	Id          string      // target name
	Query       QueryResult // query result
}
