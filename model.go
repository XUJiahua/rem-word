package main

// Word English Word
type Word struct {
	// word itself
	Token string
	// freq
	Freq int
	// indicate if you remember
	IsKnown bool
	// chinese translation
	YDTranslate string
}
