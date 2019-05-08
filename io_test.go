package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestRead(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Word
		wantErr bool
	}{
		{
			name: "read sample",
			args: args{
				filename: "word_freq_sample.csv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			spew.Dump(got)
		})
	}
}

func TestDump(t *testing.T) {
	words, err := Read("word_freq_sample.csv")
	if err != nil {
		panic(err)
	}

	type args struct {
		filename string
		words    []*Word
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "rename filename",
			args: args{
				filename: "word_freq_sample.csv",
				words:    words,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Dump(tt.args.filename, tt.args.words); (err != nil) != tt.wantErr {
				t.Errorf("Dump() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
