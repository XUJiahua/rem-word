package main

import (
	"fmt"
	"math/rand"
	"time"
)

func nextUnknownWord(words []*Word) chan *Word {
	wordChan := make(chan *Word)

	go func() {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for _, i := range r.Perm(len(words)) {
			if words[i].IsKnown {
				continue
			}
			wordChan <- words[i]
		}
		close(wordChan)
	}()

	return wordChan
}

func main() {
	// space: next word
	// k: i know it
	// i: invalid word

	// as a param
	filename := "word_freq_sample.csv"
	words, err := Read(filename)
	if err != nil {
		panic(err)
	}

	wordChan := nextUnknownWord(words)

	for word := range wordChan {
		fmt.Println(word.Token)
	}

	Dump(filename, words)
}
