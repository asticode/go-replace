package replace
import (
	"regexp"
)

func ReplaceAll(src *[]byte, rgx *regexp.Regexp, rpl []byte) {
	// Initialize
	var start, end, delta, offset, i int
	len := len(rpl)

	// Find all matches
	for _, indexes := range rgx.FindAllIndex(*src, -1) {
		// Update indexes
		start = indexes[0] + offset
		end = indexes[1] + offset
		delta = (end - start) - len
		offset -= delta

		// Update src length
		if delta < 0 {
			// Insert
			(*src) = append((*src)[:start], append(make([]byte, -delta), (*src)[start:]...)...)
		} else if delta > 0 {
			// Delete
			(*src) = append((*src)[:start], (*src)[start + delta:]...)
		}

		// Update src content
		for i = 0; i < len; i++ {
			(*src)[i + start] = rpl[i]
		}
	}
}
