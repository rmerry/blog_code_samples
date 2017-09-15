// A code sample that reads a document and generates a word frequency summary
// from it. This was creates as a demonstration for my blog article on bitsociety.uk: `A Quick
// Delve into Goroutines and Channels'
//
// Author: Richard Merry
// Date:  14/09/2017

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	out := make(chan string)
	go getWords(out)

	fl := make(map[string]int)
	for word := range out {
		_, ok := fl[word]
		if !ok {
			fl[word] = 1
		} else {
			fl[word] = fl[word] + 1
		}
	}

	// invert the `fl' map in order to sort it by frequency
	inverted := make(map[int]string)
	indicies := []int{}
	for k, v := range fl {
		inverted[v] = k
		indicies = append(indicies, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(indicies)))

	for i := 0; i < 50; i++ {
		fmt.Printf("%s\t\t%d\n", inverted[indicies[i]], indicies[i])
	}
}

func getWords(out chan<- string) {
	file, err := ioutil.ReadFile("doc.txt")
	if err != nil {
		panic(1)
	}

	for _, word := range strings.Fields(string(file)) {
		out <- word
	}
	close(out)
}
