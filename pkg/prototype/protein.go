package prototype

type Protein struct {
	Name       string
	AminoAcids []AminoAcid
}

type AminoAcid struct {
	Name string
}

func cloneProteins(source []Protein) []Protein {
	var p []Protein

	for _, protein := range source {
		pc := Protein{Name: protein.Name}
		if protein.AminoAcids != nil {
			pc.AminoAcids = make([]AminoAcid, len(protein.AminoAcids))
			copy(pc.AminoAcids, protein.AminoAcids)
		}
		p = append(p, pc)
	}

	return p
}
