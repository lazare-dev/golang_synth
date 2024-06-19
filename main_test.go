// main_test.go
package main

import (
	"testing"
)

func TestAudioEngineInitialization(t *testing.T) {
	t.Log("Testing Audio Engine Initialization")

	// Initialize Audio Engine with a single oscillator
	oscillator := NewOscillator(440.0, Sine, 44100)
	audioEngine, err := NewAudioEngine(44100, oscillator)
	if err != nil {
		t.Fatalf("Failed to initialize audio engine: %v", err)
	}
	defer func() {
		if err := audioEngine.Close(); err != nil {
			t.Fatalf("Failed to close audio engine: %v", err)
		}
	}()
}

func TestSetMasterVolume(t *testing.T) {
	t.Log("Testing SetMasterVolume")

	oscillator := NewOscillator(440.0, Sine, 44100)
	audioEngine, err := NewAudioEngine(44100, oscillator)
	if err != nil {
		t.Fatalf("Failed to initialize audio engine: %v", err)
	}
	defer audioEngine.Close()

	// Set and test master volume
	volumes := []float64{0.1, 0.5, 1.0}
	for _, v := range volumes {
		audioEngine.SetMasterVolume(v)
		if audioEngine.MasterVolume != v {
			t.Errorf("Expected MasterVolume: %v, got: %v", v, audioEngine.MasterVolume)
		}
	}
}

func TestSetOscillatorFrequency(t *testing.T) {
	t.Log("Testing SetOscillatorFrequency")

	oscillator := NewOscillator(440.0, Sine, 44100)
	audioEngine, err := NewAudioEngine(44100, oscillator)
	if err != nil {
		t.Fatalf("Failed to initialize audio engine: %v", err)
	}
	defer audioEngine.Close()

	// Set and test oscillator frequency
	frequencies := []float64{220.0, 440.0, 880.0}
	for _, f := range frequencies {
		audioEngine.SetOscillatorFrequency(0, f)
		if audioEngine.Oscillators[0].Frequency != f {
			t.Errorf("Expected Frequency: %v, got: %v", f, audioEngine.Oscillators[0].Frequency)
		}
	}
}

func TestSetOscillatorWaveform(t *testing.T) {
	t.Log("Testing SetOscillatorWaveform")

	oscillator := NewOscillator(440.0, Sine, 44100)
	audioEngine, err := NewAudioEngine(44100, oscillator)
	if err != nil {
		t.Fatalf("Failed to initialize audio engine: %v", err)
	}
	defer audioEngine.Close()

	// Set and test oscillator waveform
	waveforms := []Waveform{Sine, Square, Sawtooth, Triangle}
	for _, w := range waveforms {
		audioEngine.SetOscillatorWaveform(0, w)
		if audioEngine.Oscillators[0].Waveform != w {
			t.Errorf("Expected Waveform: %v, got: %v", w, audioEngine.Oscillators[0].Waveform)
		}
	}
}

func TestSetOscillatorVolume(t *testing.T) {
	t.Log("Testing SetOscillatorVolume")

	oscillator := NewOscillator(440.0, Sine, 44100)
	audioEngine, err := NewAudioEngine(44100, oscillator)
	if err != nil {
		t.Fatalf("Failed to initialize audio engine: %v", err)
	}
	defer audioEngine.Close()

	// Set and test oscillator volume
	volumes := []float64{0.1, 0.5, 1.0}
	for _, v := range volumes {
		audioEngine.SetOscillatorVolume(0, v)
		if audioEngine.Oscillators[0].Volume != v {
			t.Errorf("Expected Volume: %v, got: %v", v, audioEngine.Oscillators[0].Volume)
		}
	}
}

// SimulatedAudioOutputTest demonstrates a placeholder for testing audio output.
// This function does not produce actual sound but simulates the process for testing.
func SimulatedAudioOutputTest(t *testing.T) {
	t.Log("Simulating Audio Output Test")

	// Initialize your audio engine and oscillator here
	oscillator := NewOscillator(440.0, Sine, 44100)
	audioEngine, err := NewAudioEngine(44100, oscillator)
	if err != nil {
		t.Fatalf("Failed to initialize audio engine for simulated output test: %v", err)
	}
	defer audioEngine.Close()

	// Simulate generating audio output, e.g., a short tone or waveform
	// This is where you would call your method to generate audio, analyze it, or send it to a mock output

	t.Log("Simulated audio output generated successfully (Note: This is a mock test and does not produce real sound)")
}
