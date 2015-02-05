package hmmer

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ParseTblReport(rd io.Reader) (hits []Hit) {
	records := parseTable(rd)
	for _, cols := range records {
		hits = append(hits, parseHit(cols))
	}
	return
}

func parseTable(rd io.Reader) (records [][]string) {
	reader := bufio.NewReader(rd)
	line, err := reader.ReadString('\n')
	for err == nil {
		if line[0] != '#' {
			fields := strings.Split(strings.TrimSpace(line), " ")
			terms := []string{}
			for _, f := range fields {
				if f != "" {
					terms = append(terms, f)
				}
			}
			records = append(records, terms)
		}
		line, err = reader.ReadString('\n')
	}

	if err != nil && err != io.EOF {
		panic(err)
	}

	return records
}

func parseHit(r []string) (h Hit) {
	h.Id = r[0]
	h.Accession = r[1]
	q := QueryResult{}
	q.Id = r[2]
	q.Accession = r[3]
	h.Query = q
	h.EValue = parseFloat(r[4])
	h.Score = parseFloat(r[5])
	h.Bias = parseFloat(r[6])

	return
}

func parseFloat(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return v
}
