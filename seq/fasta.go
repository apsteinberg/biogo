package seq

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

// A reader for reading sequences in FASTA format.
type FastaReader struct {
	DeflineParser    func(string) string            // parse the definition line to get the seq id
	AnnotationParser func(string) map[string]string // parse the defline to get the annotations.
	r                *bufio.Reader
}

// NewFastaReader returns a new Reader that reads from r.
func NewFastaReader(r io.Reader) *FastaReader {
	return &FastaReader{
		r: bufio.NewReader(r),
	}
}

// Read reads one record from r. The record is a SeqRecord.
func (r *FastaReader) Read() (seq *Sequence, err error) {
	seq, err = r.parseRecord()
	seq.Seq = bytes.Replace(seq.Seq, []byte(" "), []byte(""), -1)
	return
}

// unreadRune puts the last rune read from r back.
func (r *FastaReader) unreadRune() {
	r.r.UnreadRune()
}

// skip reads runes up to and including the rune delim or until error.
func (r *FastaReader) skip(delim rune) error {
	for {
		r1, _, err := r.r.ReadRune()
		if err != nil {
			return err
		}
		if r1 == delim {
			return nil
		}
	}
	panic("unreachable")
}

// parseRecord reads and parses a single biogo.SeqRecord from r.
func (r *FastaReader) parseRecord() (seq *Sequence, err error) {
	title2id := r.DeflineParser
	if title2id == nil {
		title2id = func(title string) string { return strings.Split(title, " ")[0] }
	}
	// Skips any text before the first record.
	var r1 rune
	for err == nil {
		r1, _, err = r.r.ReadRune()
		if err != nil {
			return nil, err
		}
		if r1 == '>' {
			break
		} else {
			err = r.skip('\n')
		}
	}

	if err != nil {
		return nil, err
	}

	var defline string
	defline, err = r.r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var id, desc string
	var annotations map[string]string
	desc = strings.TrimRightFunc(defline, unicode.IsSpace)
	id = title2id(desc)
	if r.AnnotationParser != nil {
		annotations = r.AnnotationParser(desc)
	}
	// read sequences
	seqstr := []byte{}

	for err == nil {
		r1, _, err = r.r.ReadRune()
		r.r.UnreadRune()
		if err != nil || r1 == '>' {
			break
		} else {
			var line []byte
			line, err = r.r.ReadBytes('\n')
			seqstr = append(seqstr, bytes.TrimSpace(line)...)
		}
	}

	// create SeqRecord
	seq = &Sequence{
		Id:          id,
		Name:        desc,
		Seq:         seqstr,
		Annotations: annotations,
	}

	return
}

// ReadAll reads all the remaining records from r.
// Each record is a *biogo.SeqRecord.
// A successful call return err == nil, not err == EOF.
// Because ReadAll is defined to read until EOF, it does not tread end of file as an error to be reported.
func (r *FastaReader) ReadAll() (seqs []*Sequence, err error) {
	for {
		seq, err := r.Read()
		if err == io.EOF {
			seqs = append(seqs, seq)
			return seqs, nil
		}
		if err != nil {
			return nil, err
		}
		seqs = append(seqs, seq)
	}
	panic("unreachable")
}
