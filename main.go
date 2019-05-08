package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
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

func dumpPeriodically(words []*Word, done chan bool) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		var gracefulStop = make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				_ = t
				// TODO: dump words if has change
			case <-gracefulStop:
				// fmt.Println("Catch CTRL-C")
				Dump(filename, words)
				os.Exit(0)
			}
		}
	}()
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
	space: next word
	k: i know it
	i: invalid word
	t: translate it
	CTRL-C: quit
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

	done := make(chan bool)
	dumpPeriodically(words, done)

	wordChan := nextUnknownWord(words)

	// TODO: in a endless loop
	for word := range wordChan {
		reset()
		fmt.Println(word.Token)
		fmt.Println(usage)

	keyPressListenerLoop:
		for {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {
				case term.KeyCtrlC:
					// TODO: handle
					os.Exit(0)
				case term.KeySpace:
					break keyPressListenerLoop
				default:
					// we only want to read a single character or one key pressed event
					// fmt.Println("ASCII : ", ev.Ch)
					switch ev.Ch {
					case 'k':
						fmt.Println("mark it known")
					case 'd':
						fmt.Println("delete it")
					case 't':
						fmt.Println("translate it")
					}
				}
			case term.EventError:
				panic(ev.Err)
			}
		}
	}

	done <- true
}
