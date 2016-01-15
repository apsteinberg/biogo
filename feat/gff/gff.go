package gff

import (
	"math"
	"strconv"
	"strings"
)

const (
	ForwardStrand = 1
	ReverseStrand = -1
)

type Record struct {
	SeqName   string
	Source    string
	Feature   string
	Start     int
	End       int
	Score     float64
	Strand    int
	Frame     string
	Attribute string
}

func (r *Record) String() string {
	start := strconv.Itoa(r.Start)
	end := strconv.Itoa(r.End)
	score := "."
	if !math.IsNaN(r.Score) {
		score = strconv.FormatFloat(r.Score, 'g', -1, 64)
	}
	strand := "."
	switch r.Strand {
	case ForwardStrand:
		strand = "+"
		break
	case ReverseStrand:
		strand = "-"
		break
	}

	ss := []string{
		r.SeqName,
		r.Source,
		r.Feature,
		start,
		end,
		score,
		strand,
		r.Frame,
		r.Attribute,
	}

	return strings.Join(ss, "\t")
}
