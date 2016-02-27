package pileup

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBases(t *testing.T) {
	refBase := 'G'
	s := ".A.$,,+1a.,,,,-2gt.,.,$,,,.,a,,..,..,,,.,,.,.,,,,,,,..,,.,.,.....,"
	bases := parseBases([]byte(s), byte(refBase))
	assert.Equal(t, 57, len(bases), "Number of bases should be equal.")
	assert.Equal(t, byte('A'), bases[1], "Base should be equal.")
	assert.Equal(t, byte('G'), bases[0], "Base should be equal.")
}

func TestParseLine(t *testing.T) {
	line := "GCA_000155815.1_ASM15581v1\t355\tG\t57\t...$,,+1a.,,,,.,.,$,,,.,.,,..,..,,,.,,.,.,,,,,,,..,,.,.,.....,\tfkDG>fGGF@iG:AGGGDGEGFfHGgGFEGEEDGGEBGDDEEBFEEEGCG>GGGGG<"
	snp := parseLine(line)
	assert.Equal(t, "GCA_000155815.1_ASM15581v1", snp.Reference, "Reference should be the same.")
	assert.Equal(t, 355, snp.Position, "Position should be the same.")
	assert.Equal(t, byte('G'), snp.RefBase, "Reference base should be the same.")
	assert.Equal(t, len(snp.Bases), len(snp.Quals), "Bases and quals should be equal in length.")
	assert.Equal(t, "fkDG>fGGF@iG:AGGGDGEGFfHGgGFEGEEDGGEBGDDEEBFEEEGCG>GGGGG<", string(snp.Quals), "Quals should be equal.")
}

func TestPi(t *testing.T) {

}
