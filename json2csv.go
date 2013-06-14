//json2csv converts homogenous arrays of json objects into comma-delimited text.
//It takes output from stdin and writes to stdout
package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
  "github.com/thethirdwheel/stringmap"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := csv.NewWriter(os.Stdout)
	header := true
	for scanner.Scan() {
		var i interface{}
		lineBytes := []byte(scanner.Text())
		err := json.Unmarshal(lineBytes, &i)
		if err != nil {
			panic(err)
		}
		m := i.(map[string]interface{})
		s := stringmap.StringPairs{}
		for k, v := range m {
			switch v := v.(type) {
			case map[string]interface{}:
				s = append(s, &stringmap.StringPair{k, fmt.Sprint(stringmap.StringMap(v))})
			default:
				s = append(s, &stringmap.StringPair{k, fmt.Sprintf("%v", v)})
			}
		}
		sort.Sort(stringmap.ByKey{s})
		if header {
			header = false
			writer.Write(s.Keys())
		}
		writer.Write(s.Vals())
		writer.Flush()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
