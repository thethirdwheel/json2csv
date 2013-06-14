package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type StringPair struct {
	Key string
	Val string
}

type StringPairs []*StringPair

func (s StringPairs) Len() int      { return len(s) }
func (s StringPairs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s StringPairs) Keys() []string {
	keys := []string{}
	for _, o := range s {
		keys = append(keys, o.Key)
	}
	return keys
}
func (s StringPairs) Vals() []string {
	vals := []string{}
	for _, o := range s {
		vals = append(vals, o.Val)
	}
	return vals
}

type StringMap map[string]interface{}

func (m StringMap) String() string {
	vals := []string{}
	for k, v := range m {
		if len(m) == 1 {
			return fmt.Sprintf("%v", v)
		} else {
			vals = append(vals, fmt.Sprintf("%+v:%+v", k, v))
		}
	}
	return strings.Join(vals, " ")
}

type ByKey struct{ StringPairs }

func (s ByKey) Less(i, j int) bool { return s.StringPairs[i].Key < s.StringPairs[j].Key }

type ByVal struct{ StringPairs }

func (s ByVal) Less(i, j int) bool { return s.StringPairs[i].Val < s.StringPairs[j].Val }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := csv.NewWriter(os.Stdout)
	for scanner.Scan() {
		var i interface{}
		lineBytes := []byte(scanner.Text())
		err := json.Unmarshal(lineBytes, &i)
		if err != nil {
			panic(err)
		}
		m := i.(map[string]interface{})
		s := StringPairs{}
		for k, v := range m {
			switch v := v.(type) {
			case map[string]interface{}:
				s = append(s, &StringPair{k, fmt.Sprint(StringMap(v))})
			default:
				s = append(s, &StringPair{k, fmt.Sprintf("%v", v)})
			}
		}
		sort.Sort(ByKey{s})
		csvseq := s.Vals()
		writer.Write(csvseq)
		writer.Flush()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
