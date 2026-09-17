package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	var magazineLetter = make(map[rune]int)
	var ransomNoteLetters = make(map[rune]int)
	for _, letter := range magazine {
		magazineLetter[letter]++
	}

	for _, letter := range ransomNote {
		ransomNoteLetters[letter]++
	}

	for letter, counter := range ransomNoteLetters {
		if counter > magazineLetter[letter] {
			return false
		}
	}

	return true
}

func main() {
	var ransomNote = "aa"
	var magazine = "bb"
	fmt.Println(canConstruct(ransomNote, magazine))
}
