package foodchain

import "strings"

const testVersion = 3

var lines = map[string][]string{
	"fly": {""},
	"spider": {"It wriggled and jiggled and tickled inside her.",
		"She swallowed the spider to catch the fly."},
	"bird": {"How absurd to swallow a bird!", "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her."},
	"cat": {"Imagine that, to swallow a cat!",
		"She swallowed the cat to catch the bird."},
	"dog":  {"What a hog, to swallow a dog!", "She swallowed the dog to catch the cat."},
	"goat": {"Just opened her throat and swallowed a goat!", "She swallowed the goat to catch the dog."},
	"cow":  {"I don't know how she swallowed a cow!", "She swallowed the cow to catch the goat."},
}
var lineOrder = []string{"", "fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var firstLine = "I know an old lady who swallowed a [animal]."
var lastLineVerse = "I don't know why she swallowed the fly. Perhaps she'll die."
var lastLineSong = "She's dead, of course!"
var numOfVerse = 8

func Verse(lineNum int) (output string) {
	a := make([]string, 1)
	b := make([]string, 0)
	a[0] = strings.Replace(firstLine, "[animal]", lineOrder[lineNum], 1)
	for i := 2; i <= lineNum && lineNum != numOfVerse; i++ {
		if len(a) > 2 {
			b = append(b, a[2:]...)
			a = a[0:1]
		}
		a = append(a, lines[lineOrder[i]][0], lines[lineOrder[i]][1])
		a = append(a, b...)
		b = nil
	}
	var lastLine string
	if lineNum == numOfVerse {
		lastLine = lastLineSong
	} else {
		lastLine = lastLineVerse
	}
	a = append(a, lastLine)
	return strings.Join(a, "\n")
}
func Verses(fromLineNum int, toLineNum int) (output string) {
	for v := fromLineNum; v <= toLineNum; v++ {
		output = output + Verse(v)
		if v != toLineNum {
			output = output + "\n\n"
		}
	}
	return
}
func Song() (output string) {
	return Verses(1, numOfVerse)
}
