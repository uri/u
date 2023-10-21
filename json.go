package u

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func JSONTuplesPretty(args ...string) string {
	w := &strings.Builder{}
	fmt.Fprintf(w, "{\n")
	for i := 0; i < len(args)-1; i += 2 {
		key := args[i]
		val := args[i+1]
		fmt.Fprintf(w, "\t%q: %q,\n", key, val)
	}
	key := args[len(args)-2]
	val := args[len(args)-1]
	fmt.Fprintf(w, "\t%q: %q\n}", key, val)

	return w.String()
}

// JSONTuples creates a marshalable JSON string by treating even-indexed args
// as keys and odd-indexed args as values.
func JSONTuples(args ...string) string {
	w := &strings.Builder{}
	fmt.Fprintf(w, "{")
	for i := 0; i < len(args)-1; i += 2 {
		key := args[i]
		val := args[i+1]
		fmt.Fprintf(w, "%q:%q,", key, val)
	}
	key := args[len(args)-2]
	val := args[len(args)-1]
	fmt.Fprintf(w, "%q:%q}", key, val)

	return w.String()
}

// PrettyJSON reads from r and writes to w, priting nicely formatted JSON.
func PrettyJSON(r io.Reader, w io.Writer) {
	var m map[string]any

	dec := json.NewDecoder(r)
	Must(dec.Decode(&m))

	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	Must(enc.Encode(m))
}
