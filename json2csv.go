package main 

import (
  "encoding/json"
  "encoding/csv"
  "bufio"
  "fmt"
  "os"
  "sort"
)


type Pair struct  {
  Key string
  Val string
}

type Pairs []*Pair

func (s Pairs) Len() int { return len(s) }
func (s Pairs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Pairs) Keys() []string { 
  keys := []string{}
  for _, o := range s { 
    keys = append(keys,o.Key)
  }
  return keys
}
func (s Pairs) Vals() []string {
  vals := []string{}
  for _, o := range s {
    vals = append(vals,o.Val)
  }
  return vals
}

type ByKey struct { Pairs }

func (s ByKey) Less(i, j int) bool { return s.Pairs[i].Key < s.Pairs[j].Key }

type ByVal struct { Pairs }

func (s ByVal) Less(i, j int) bool { return s.Pairs[i].Val < s.Pairs[j].Val }

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  writer := csv.NewWriter(os.Stdout)
  for scanner.Scan() {
    var i interface{}
    lineBytes := []byte(scanner.Text())
    err := json.Unmarshal(lineBytes, &i)
    if err != nil { panic(err) }
    m := i.(map[string]interface{})
    s := Pairs{}
    for k, v := range m { 
      p := Pair{k, fmt.Sprintf("%+v", v)}
      s = append(s,&p)
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
