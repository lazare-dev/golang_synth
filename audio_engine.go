package main

import (
	"log"
	"math"
	"sync"

	"github.com/gordonklaus/portaudio"
)

type Waveform int

const (
	Sine Waveform = iota
	Square
	Sawtooth
	Triangle
)

type AudioEngine struct {
	Oscillators  []*Oscillator
	MasterVolume float64
	Stream       *portaudio.Stream
	mutex        sync.Mutex
	activeNotes  map[string]bool
}

func NewAudioEngine(sampleRate float64, oscillators ...*Oscillator) (*AudioEngine, error) {
	err := portaudio.Initialize()
	if err != nil {
		log.Printf("Error initializing PortAudio: %v\n", err)
		return nil, err
	}

	host, err := portaudio.DefaultHostApi()
	if err != nil {
		log.Printf("Error getting default host API: %v\n", err)
		return nil, err
	}

	deviceInfo := host.DefaultOutputDevice
	config := portaudio.LowLatencyParameters(nil, deviceInfo)
	config.Output.Channels = 1
	config.SampleRate = sampleRate
	config.FramesPerBuffer = 64

	ae := &AudioEngine{
		MasterVolume: 1.0,
		Oscillators:  oscillators,
		activeNotes:  make(map[string]bool),
	}

	streamCallback := func(out []float32) {
		ae.mutex.Lock()
		defer ae.mutex.Unlock()

		for i := range out {
			out[i] = 0
			for _, osc := range ae.Oscillators {
				if osc.Active {
					out[i] += float32(osc.NextSample() * osc.Volume)
				}
			}
			out[i] *= float32(ae.MasterVolume)
		}
	}

	stream, err := portaudio.OpenStream(config, streamCallback)
	if err != nil {
		log.Printf("Error opening audio stream: %v\n", err)
		return nil, err
	}

	ae.Stream = stream

	err = ae.Stream.Start()
	if err != nil {
		log.Printf("Error starting audio stream: %v\n", err)
		return nil, err
	}

	log.Println("Audio engine initialized successfully.")
	return ae, nil
}

func (ae *AudioEngine) SetMasterVolume(volume float64) {
	ae.mutex.Lock()
	defer ae.mutex.Unlock()
	ae.MasterVolume = math.Max(0.0, math.Min(volume, 1.0))
}

func (ae *AudioEngine) PlayNote(note string, volume float64) {
	frequency := NoteToFrequency(note)
	ae.mutex.Lock()
	defer ae.mutex.Unlock()
	if _, exists := ae.activeNotes[note]; !exists {
		for _, osc := range ae.Oscillators {
			osc.Active = true
			osc.SetFrequency(frequency)
			osc.SetVolume(volume)
			break // Assuming one note per oscillator for simplicity
		}
		ae.activeNotes[note] = true
	}
}

func (ae *AudioEngine) StopNote(note string) {
	ae.mutex.Lock()
	defer ae.mutex.Unlock()
	if _, exists := ae.activeNotes[note]; exists {
		for _, osc := range ae.Oscillators {
			if osc.Active {
				osc.SetVolume(0)
				osc.Active = false
				break // Assuming one note per oscillator for simplicity
			}
		}
		delete(ae.activeNotes, note)
	}
}

func (ae *AudioEngine) Close() error {
	err := ae.Stream.Close()
	if err != nil {
		log.Printf("Error closing audio stream: %v\n", err)
		return err
	}
	return portaudio.Terminate()
}

func NoteToFrequency(note string) float64 {
	notes := map[string]int{
		"C": 0, "C#": 1, "D": 2, "D#": 3, "E": 4, "F": 5, "F#": 6, "G": 7, "G#": 8, "A": 9, "A#": 10, "B": 11,
	}
	baseFrequency := 440.0 // A4
	a4Index := 9
	notePart, octavePart := note[:len(note)-1], note[len(note)-1]
	octave := int(octavePart - '0')
	noteIndex := notes[notePart]
	semitoneDistance := (octave-4)*12 + (noteIndex - a4Index)
	return baseFrequency * math.Pow(2.0, float64(semitoneDistance)/12.0)
}
