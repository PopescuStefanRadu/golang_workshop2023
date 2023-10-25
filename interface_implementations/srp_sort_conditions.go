package interface_implementations

type IntSliceSort []int

func (s IntSliceSort) Len() int {
	return len(s)
}

func (s IntSliceSort) Less(i, j int) bool {
	if s[i]%2 == 1 && s[j]%2 == 0 { // daca prima este impara, a doua para
		return true //
	}

	if s[i]%2 == s[j]%2 {
		return s[i] < s[j] // comparata egalitatea paritatilor
	}

	return false
}

func (s IntSliceSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
