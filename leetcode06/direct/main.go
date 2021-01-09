package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1i := 0
	w2i := 0
	w1s := 0
	w2s := 0
	for {
		for w1i < len(word1) && w1s >= len(word1[w1i]) {
			w1s = 0
			w1i++
		}
		for w2i < len(word2) && w2s >= len(word2[w2i]) {
			w2s = 0
			w2i++
		}
		if w1i >= len(word1) || w2i >= len(word2) {
			return w1i >= len(word1) && w2i >= len(word2)
		}
		if word1[w1i][w1s] != word2[w2i][w2s] {
			return false
		}
		w1s++
		w2s++
	}
}

func main() {
	arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"})
}
