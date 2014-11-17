package fuzzy

type Matcher struct {
	given   []string
	letters []rune
}

func Match(lines []string) *Matcher {
	return &Matcher{
		given: lines,
	}
}

func (m *Matcher) Push(r rune) {
	m.letters = append(m.letters, r)
}

func (m *Matcher) Pop() {
	m.letters = m.letters[:len(m.letters)-1]
}

func (m *Matcher) Len() int {
	return len(m.letters)
}

func (m *Matcher) Clear() {
	m.letters = m.letters[:0]
}

func (m *Matcher) Matches() []string {
	// very naive implementation, I haven't looked for
	// proper algorithms
	var matches []string
nextLine:
	for _, line := range m.given {
		i := 0
	nextFuzzyLetter:
		for _, l := range m.letters {
			for _, r := range []rune(line[i:]) {
				if r == l {
					continue nextFuzzyLetter
				}
				i++
			}
			continue nextLine
		}
		// all letters were found
		matches = append(matches, line)
	}
	return matches
}
