package fuzzy

var _ Matcher = &LazyMatcher{}

type LazyMatcher struct {
	given       []string
	letters     []rune
	lastMatches []string
	totalCost   int
}

func LazyMatch(lines []string) *LazyMatcher {
	return &LazyMatcher{
		given:       lines,
		lastMatches: lines,
	}
}

func (m *LazyMatcher) Push(r rune) {
	m.letters = append(m.letters, r)
}

func (m *LazyMatcher) Pop() {
	m.letters = m.letters[:len(m.letters)-1]
	m.lastMatches = m.given
}

func (m *LazyMatcher) Len() int {
	return len(m.letters)
}

func (m *LazyMatcher) Clear() {
	m.letters = m.letters[:0]
	m.lastMatches = m.given
}

func (m *LazyMatcher) matchesCost() ([]string, int) {
	matches, cost := computeMatches(m.lastMatches, m.letters)
	m.totalCost += cost
	m.lastMatches = matches // save the result, faster next time
	return matches, m.totalCost
}

// Matches computes the current matches. O(n) where n is the number of
// runes in the input, if all lines were concatenated.
//
// Worst case is when all lines match at their very end.
// Best case is when no lines match with fuzzy query as long
// as the longest line.
func (m *LazyMatcher) Matches() []string {
	matches, _ := m.matchesCost()
	return matches
}
