package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneProteins(t *testing.T) {
	trpCage := Protein{
		Name: "Trp-Cage",
		AminoAcids: []AminoAcid{
			{Name: "Asn"},
			{Name: "Gln"},
			{Name: "Leu"},
		},
	}
	p := []Protein{trpCage}

	pc := cloneProteins(p)

	assert.Equal(t, "Trp-Cage", pc[0].Name)
	assert.Equal(t, []AminoAcid{
		{Name: "Asn"},
		{Name: "Gln"},
		{Name: "Leu"},
	}, pc[0].AminoAcids)

	trpCage.AminoAcids[0].Name = "Asp"
	assert.Equal(t, "Asn", pc[0].AminoAcids[0].Name)
}

func TestCloneEmptyProtein(t *testing.T) {
	noop := Protein{Name: "Noop"}
	p := []Protein{noop}
	pc := cloneProteins(p)

	assert.Equal(t, "Noop", pc[0].Name)
	assert.Nil(t, pc[0].AminoAcids)
}
