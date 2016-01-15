package gff

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

const WrongNumberColumns = "Wrong number of columns"

type Reader struct {
	r *bufio.Reader
}

func NewReader(rd io.Reader) *Reader {
	return &Reader{bufio.NewReader(rd)}
}

func (r *Reader) Read() (*Record, error) {
	var l string
	var err error
	for {
		l, err = r.r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		l = strings.TrimSpace(l)
		if l[0] != '#' && len(l) != 0 {
			break
		}
	}
	record, err := decord(l)
	return record, err
}

func (r *Reader) ReadAll() ([]*Record, error) {
	records := []*Record{}
	for {
		record, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		records = append(records, record)
	}

	return records, nil
}

func decord(line string) (*Record, error) {
	l := strings.TrimSpace(line)
	terms := strings.Split(l, "\t")
	if len(terms) != 9 {
		return nil, errors.New(WrongNumberColumns + fmt.Sprintf(": %s", l))
	}
	r := Record{}
	r.SeqName = terms[0]
	r.Source = terms[1]
	r.Feature = terms[2]
	r.Start = atoi(terms[3])
	r.End = atoi(terms[4])
	if terms[5] == "." {
		r.Score = math.NaN()
	} else {
		r.Score = atof(terms[5])
	}
	switch terms[6] {
	case "+":
		r.Strand = ForwardStrand
		break
	case "-":
		r.Strand = ReverseStrand
		break
	default:
		r.Strand = 0
	}
	r.Frame = terms[7]

	r.Attribute = terms[8]
	return &r, nil
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}
