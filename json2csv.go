package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
    p StringPair
		for k, v := range m {
      switch v := v.(type) {
        default:
          fmt.Printf("unexpected type %T", v)
        case map:
          if len(v) == 1 {
            //Do some logic here to extract the lone value in the map
            p = StringPair(k, fmt.Sprintf("%v", v))
          } else {
            p = StringPair(k, fmt.Sprintf("%+v", v)
          }
      }
			s = append(s, &p)
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
