package pkg

// These are only so we can implement the sort interface and make it easier

// IntPair is a simple k/v pair with an int value
type IntPair struct {
	Key   string
	Value int
}

// IntPairList is a list of simple k/v int pairs
type IntPairList []IntPair

// Len returns the length of the list
func (p IntPairList) Len() int { return len(p) }

// Swap swaps two values in the list
func (p IntPairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Less checks if a given value in the list is lesser than another
func (p IntPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
