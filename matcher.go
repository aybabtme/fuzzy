package fuzzy

type Matcher interface {
	Push(r rune)
	Pop()
	Clear()
	Matches() []string

	// like matches, but return the total number of comparison
	// that the matcher did since inception.
	matchesCost() ([]string, int)
}

func computeMatches(lines []string, fuzzy []rune) ([]string, int) {
	comparisons := 0
	// very naive implementation, I haven't looked for
	// proper algorithms
	var matches []string
nextLine:
	for _, line := range lines {
		i := 0
		if len(line) < len(fuzzy) {
			// impossible to match
			continue nextLine
		}

	nextFuzzyLetter:
		for j, l := range fuzzy {
			rest := []rune(line)[i:]

			fuzzyLeft := len(fuzzy) - j
			letterLeft := len(rest)
			for _, r := range rest {
				if fuzzyLeft > letterLeft {
					// not enough letters to match
					continue nextLine
				}

				i++

				comparisons++
				if r == l {
					continue nextFuzzyLetter
				}
				letterLeft--
			}
			continue nextLine
		}
		// all letters were found
		matches = append(matches, line)
	}

	return matches, comparisons
}
