// Knuth Morris Pratt string matching algorithm.
// This code is a direct translation of the code in the textbook
// "T.Cormen Introduction to Algorithms 3rd Edition" Pages 1004 to 1006

package kmp

type KMPArg struct {
    Text string
    Pattern string
}

func (km *KMPArg) Match (occurances []int) {
    if km == nil {
        return
    }

    T, P := km.T, km.P
	n := len(T)
	m := len(P)

	pi := computePrefixFunction(P)

	q := 0
	for i := 0; i < n; i++ {
		for q > 0 && P[q] != T[i] {
			q = pi[q]
		}

		if P[q] == T[i] {
			q += 1
		}

		if q == m {
			occurances = append(occurances, i-m)
			q = pi[q] // look for the next match
		}
	}

	return
}

func computePrefixFunction(P string) []int {
	m := len(P)
	pi := make([]int, m)

	pi[0] = 0
	k := 0

	opCount := 1
	for q := 1; q < m; q++ {
		for k > 0 && P[k] != P[q] {
			opCount += 1
			k = pi[k]
		}

		opCount += 1
		if P[k] == P[q] {
			k += 1
		}

		pi[q] = k
	}

	return pi
}
