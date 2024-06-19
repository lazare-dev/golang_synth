// oscillator.go
package main

import (
	"math"
)

type Oscillator struct {
	Frequency      float64
	Waveform       Waveform
	Volume         float64
	Phase          float64
	PhaseIncrement float64
	SampleRate     float64
	Active         bool
}

func NewOscillator(frequency float64, waveform Waveform, sampleRate float64) *Oscillator {
	return &Oscillator{
		Frequency:      frequency,
		Waveform:       waveform,
		Volume:         1.0,
		Phase:          0.0,
		SampleRate:     sampleRate,
		PhaseIncrement: 2.0 * math.Pi * frequency / sampleRate,
		Active:         false,
	}
}

func (o *Oscillator) NextSample() float64 {
	if !o.Active {
		return 0
	}
	o.Phase += o.PhaseIncrement
	if o.Phase >= 2.0*math.Pi {
		o.Phase -= 2.0 * math.Pi
	}

	var sample float64
	switch o.Waveform {
	case Sine:
		sample = math.Sin(o.Phase)
	case Square:
		if o.Phase < math.Pi {
			sample = 1.0
		} else {
			sample = -1.0
		}
	case Sawtooth:
		sample = 2.0*(o.Phase/(2.0*math.Pi)) - 1.0
	case Triangle:
		sample = 1.0 - 4.0*math.Abs(math.Mod(o.Phase/(math.Pi*2)+0.75, 1.0)*2-1)
	}
	return sample * o.Volume
}

func (o *Oscillator) SetFrequency(frequency float64) {
	o.Frequency = frequency
	o.PhaseIncrement = 2.0 * math.Pi * frequency / o.SampleRate
}

func (o *Oscillator) SetWaveform(waveform Waveform) {
	o.Waveform = waveform
}

func (o *Oscillator) SetVolume(volume float64) {
	o.Volume = math.Max(0.0, math.Min(volume, 1.0))
}
