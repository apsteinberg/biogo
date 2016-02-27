package pileup

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

type Reader struct {
	r *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	rd := Reader{
		r: bufio.NewReader(r),
	}
	return &rd
}

func (r *Reader) Read() (*SNP, error) {
	line, err := r.r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	snp := parseLine(line)
	return snp, nil
}

func parseLine(line string) *SNP {
	line = strings.TrimSpace(line)
	snp := SNP{}
	terms := strings.Split(line, "\t")
	snp.Reference = terms[0]
	snp.Position = stringToInt(terms[1])
	snp.RefBase = terms[2][0]
	num := stringToInt(terms[3])
	if num > 0 {
		snp.Bases = parseBases([]byte(terms[4]), snp.RefBase)
		snp.Quals = []byte(terms[5])
	}
	return &snp
}

func parseBases(s []byte, refBase byte) []byte {
	bases := []byte{}
	i := 0
	for i < len(s) {
		b := s[i]
		if b == '^' {
			i += 2
		} else if b == '$' {
			i++
		} else if b == '+' || b == '-' {
			k := i + 1
			v := []byte{}
			for s[k] >= '0' && s[k] <= '9' {
				v = append(v, s[k])
				k++
			}
			n := bytesToInt(v)
			i = i + (k - i) + n
		} else {
			if b == '.' || b == ',' {
				b = refBase
			}
			bases = append(bases, b)
			i++
		}
	}
	return bytes.ToUpper(bases)
}

func bytesToInt(v []byte) int {
	s := string(v)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
