package fuzzy

var _ Matcher = &EagerMatcher{}

type EagerMatcher struct {
	given     []string
	matches   []string
	letters   []rune
	matching  int
	totalCost int
}

func EagerMatch(lines []string) *EagerMatcher {
	return &EagerMatcher{
		given:   lines,
		matches: lines,
	}
}

func (m *EagerMatcher) Push(r rune) {
	m.letters = append(m.letters, r)
	var cost int
	m.matches, cost = computeMatches(m.matches, m.letters)
	m.totalCost += cost
}

func (m *EagerMatcher) Pop() {
	m.letters = m.letters[:len(m.letters)-1]
	var cost int
	m.matches, cost = computeMatches(m.given, m.letters)
	m.totalCost += cost
}

func (m *EagerMatcher) Len() int {
	return len(m.letters)
}

func (m *EagerMatcher) Clear() {
	m.letters = m.letters[:0]
	m.matches = m.given
}

func (m *EagerMatcher) matchesCost() ([]string, int) {
	return m.matches, m.totalCost
}

func (m *EagerMatcher) Matches() []string {
	matches, _ := m.matchesCost()
	return matches
}
