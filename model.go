package main

// Word English Word
type Word struct {
	// TODO: add tag csv and parse
	// word itself
	Token string
	// freq
	Freq int
	// indicate if you remember
	IsKnown bool
	// chinese translation
	YDTranslate string
	// in memory
	Invalid bool
}
