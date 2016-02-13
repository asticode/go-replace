package replace
import (
	"testing"
	"regexp"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestReplaceAll(t *testing.T) {
	// Initialize
	src := []byte("/test/{m1}/test/{ma2}/test/{match3}")
	rgx := regexp.MustCompile("{[A-Za-z0-9_]+}")
	rpl := []byte("value")

	// Replace all
	ReplaceAll(&src, rgx, rpl)

	// Assert
	assert.Equal(t, fmt.Sprintf("/test/%s/test/%s/test/%s", string(rpl), string(rpl), string(rpl)), string(src))
}

func TestReplaceAllBis(t *testing.T) {
	// Initialize
	src := []byte("/test/{m1}/test/{ma2}/test/{match3}")
	rgx := regexp.MustCompile("{[A-Za-z0-9_]+}")
	rpl := []byte("valuevaluevaluevaluevaluevaluevalue")

	// Replace all
	ReplaceAll(&src, rgx, rpl)

	// Assert
	assert.Equal(t, fmt.Sprintf("/test/%s/test/%s/test/%s", string(rpl), string(rpl), string(rpl)), string(src))
}
