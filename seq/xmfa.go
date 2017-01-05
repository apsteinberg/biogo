package seq

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// XMFAReader for extended Multi-FASTA format
type XMFAReader struct {
	r *bufio.Reader
}

// NewXMFAReader return a XMFAReader.
func NewXMFAReader(rd io.Reader) *XMFAReader {
	xmfa := XMFAReader{}
	xmfa.r = bufio.NewReader(rd)
	return &xmfa
}

// Read returns a alignment (block).
func (r XMFAReader) Read() (alignment []Sequence, err error) {
	var seqID, seqBytes []byte
	for {
		var line []byte
		line, err = r.r.ReadBytes('\n')
		if err != nil {
			return
		}

		line = bytes.TrimSpace(line)
		if line[0] == '=' {
			break
		} else if line[0] == '#' {
			continue
		} else if line[0] == '>' {
			if len(seqBytes) > 0 {
				ss := Sequence{}
				ss.Id = string(seqID)
				ss.Seq = seqBytes
				alignment = append(alignment, ss)
			}
			seqID = line[1:]
			seqBytes = []byte{}
		} else {
			seqBytes = append(seqBytes, line...)
		}

	}

	if len(seqBytes) > 0 {
		ss := Sequence{}
		ss.Id = string(seqID)
		ss.Seq = seqBytes
		alignment = append(alignment, ss)
	}

	return
}

// ReadXMFA read sequences from XMFA file.
func ReadXMFA(filename string) [][]*Sequence {
	seqGroups := [][]*Sequence{}
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var b bytes.Buffer

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		if line[0] != '=' {
			b.Write(line)
		} else {
			sequences := readFasta(&b)
			seqGroups = append(seqGroups, sequences)
			b = *new(bytes.Buffer)
		}
	}

	return seqGroups
}

func readFasta(b io.Reader) []*Sequence {
	fastaReader := NewFastaReader(b)
	sequences, err := fastaReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return sequences
}
