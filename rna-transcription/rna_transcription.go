package strand

func ToRNA(dna string) string {
	var m = map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}

	var res string
	for _, v := range dna {
		res += m[v]
	}
	return res
}
