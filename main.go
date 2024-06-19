// main.go

package main

import (
	"log"
)

func main() {
	log.Println("Creating oscillators with default settings.")

	// Initialize oscillators with default settings.
	oscillator1 := NewOscillator(440.0, Sine, 44100)     // A4, Sine wave
	oscillator2 := NewOscillator(440.0, Square, 44100)   // A4, Square wave
	oscillator3 := NewOscillator(440.0, Sawtooth, 44100) // A4, Sawtooth wave

	// Initialize the audio engine with the oscillators.
	ae, err := NewAudioEngine(44100, oscillator1, oscillator2, oscillator3)
	if err != nil {
		log.Fatalf("Failed to initialize audio engine: %v", err)
	}
	defer func() {
		if err := ae.Close(); err != nil {
			log.Fatalf("Failed to close audio engine: %v", err)
		}
	}()

	// Create the synthesizer UI without starting any tone.
	createSynthUI(ae)
}
