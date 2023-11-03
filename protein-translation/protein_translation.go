package protein

import "errors"

var ErrStop error = errors.New("")
var ErrInvalidBase error = errors.New("")
var Proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var result []string
	for nt := 0; nt <= len(rna)-3; nt += 3 {
		aminoacid, ok := FromCodon(rna[nt : nt+3])
		if ok == ErrInvalidBase {
			return result, ErrInvalidBase
		} else if ok == ErrStop {
			return result, nil
		}
		result = append(result, aminoacid)
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	protein := Proteins[codon]
	if protein == "STOP" {
		return "", ErrStop
	}
	if protein == "" {
		return "", ErrInvalidBase
	}
	return protein, nil
}
