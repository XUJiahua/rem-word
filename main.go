package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	term "github.com/nsf/termbox-go"
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

var filename string

func init() {
	const (
		defaultFilename = "word_freq_sample.csv"
	)
	flag.StringVar(&filename, "csv_file", defaultFilename, "rem-word -csv_file="+defaultFilename)
	flag.Parse()
}

const usage = `


	------ shortcuts ----------

	space: next word (or skip)
	k: mark it known (space)
	d: delete it (space)
	t: translate it
	s: statistics
	CTRL-C: quit (save)
`

func reset() {
	term.Sync() // cosmestic purpose
}

func main() {
	words, err := Read(filename)
	if err != nil {
		panic(err)
	}

	err = term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	wordChan := nextUnknownWord(words)

	for word := range wordChan {
		reset()
		fmt.Printf("word:      %s (%d)\n", word.Token, word.Freq)
		fmt.Println(usage)

	keyPressListenerLoop:
		for {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {
				case term.KeyCtrlC:
					Dump(filename, words)
					total, known := stat(words)
					fmt.Printf("%d (known) / %d (total)\n", known, total)
					os.Exit(0)
				case term.KeySpace:
					break keyPressListenerLoop
				default:
					switch ev.Ch {
					case 'k':
						word.IsKnown = true
						break keyPressListenerLoop
					case 'd':
						word.Invalid = true
						break keyPressListenerLoop
					case 't':
						if word.YDTranslate == "" {
						// TODO: youdaoapi
						}
						reset()
						fmt.Printf("word:      %s (%d)\n", word.Token, word.Freq)
						fmt.Printf("translate: %s \n", word.YDTranslate)
						fmt.Println(usage)
				
					case 's':
						total, known := stat(words)
						reset()
						fmt.Printf("%d (known) / %d (total)\n", known, total)
						fmt.Println(usage)
					}
				}
			case term.EventError:
				panic(ev.Err)
			}
		}
	}
}

func stat(words []*Word) (total, known int) {
	for _, word :=range words {
		if word.IsKnown {
			known++
		}
	}
	
	total = len(words)
	
	return
}