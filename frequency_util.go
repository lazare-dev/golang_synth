// frequency_util.go

package main

import (
	"log"
	"math"
)

func NoteToFrequency(note string) float64 {
	notes := map[string]int{
		"C": 0, "C#": 1, "D": 2, "D#": 3, "E": 4, "F": 5, "F#": 6, "G": 7, "G#": 8, "A": 9, "A#": 10, "B": 11,
	}
	baseFrequency := 440.0 // Frequency of A4
	a4Index := 9           // The index of A in the fourth octave

	notePart := note[:len(note)-1]
	octavePart := note[len(note)-1] // Directly get the last character

	if octavePart < '0' || octavePart > '9' {
		log.Printf("Invalid octave part: %c", octavePart)
		return 0
	}

	octave := int(octavePart - '0') // Convert to int by subtracting '0'
	noteIndex, ok := notes[notePart]
	if !ok {
		log.Printf("Failed to find note index for: %s", notePart)
		return 0
	}

	semitoneDistance := (octave-4)*12 + (noteIndex - a4Index)
	return baseFrequency * math.Pow(2.0, float64(semitoneDistance)/12.0)
}
