package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "math/rand"
  "os"
  "strings"
  "time"
)

// add a alias for []string
type Prefix []string

func (p Prefix) String() string{
  return strings.Join(p, " ")
}

func (p Prefix) Shift(word string){
  copy(p, p[1:])
  p[len(p)-1] = word
}

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
type Chain struct{
  chain map[string][]string
  prefixLen int
}

// constructor of Chain
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Build(r io.Reader){
  br := bufio.NewReader(r)
  p := make(Prefix, c.prefixLen)
  for {
    var s string
    if _, err := fmt.Fscan(br, &s); err != nil{
      break
    }
    key := p.String()
    // weird, why not append(c.chain[key], s)
    c.chain[key] = append(c.chain[key], s)
    p.Shift(s)
  }
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string{
  // a array of string of size prefixLen
  p := make(Prefix, c.prefixLen)
  var words []string
  for i := 0; i < n; i++{
    choices := c.chain[p.String()]
    if len(choices) == 0{
      break
    }
    next := choices[rand.Intn(len(choices))]
    words = append(words, next)
    p.Shift(next)
  }
  return strings.Join(words, " ")
}

// Prefix []string
func main(){
  numWords := flag.Int("words", 100, "maximum number of words to print")
  prefixLen := flag.Int("prefix", 2, "prefix length in words")

  flag.Parse()                     // Parse command-line flags.
  rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

  c := NewChain(*prefixLen)     // Initialize a new Chain.
  c.Build(os.Stdin)             // Build chains from standard input.
  text := c.Generate(*numWords) // Generate text.
  fmt.Println(text)             // Write text to standard output.
}
