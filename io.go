package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// Dump to csv file
func Dump(filename string, words []*Word) error {
	backup(filename)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	for _, word := range words {
		if word.Invalid {
			continue
		}
		if err := w.Write([]string{
			word.Token,
			fmt.Sprintf("%d", word.Freq),
			fmt.Sprintf("%v", word.IsKnown),
			// word.YDTranslate,
			"", // NOTE: don't dump translate currently
		}); err != nil {
			return err
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

// backup previous version, add timestamp
func backup(filename string) {
	dir, file := filepath.Split(filename)
	backupFilename := filepath.Join(dir, time.Now().Format("20060102150405")+"_"+file)
	err := os.Rename(filename, backupFilename)
	if err != nil {
		logrus.Error(err)
	}
}

// Read filename
func Read(filename string) ([]*Word, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var words []*Word

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) < 2 {
			continue
		}

		word := &Word{
			Token: record[0],
		}

		if freq, err := strconv.ParseInt(record[1], 10, 32); err == nil {
			word.Freq = int(freq)
		}

		if len(record) >= 3 {
			if k, err := strconv.ParseBool(record[2]); err == nil {
				word.IsKnown = k
			}
		}

		if len(record) >= 4 {
			word.YDTranslate = record[3]
		}

		words = append(words, word)
	}

	return words, nil
}
