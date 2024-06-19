// ui.go

package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createSynthUI(ae *AudioEngine) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Synthesizer")
	myWindow.Resize(fyne.NewSize(1024, 600))

	// Master volume control
	volumeControl := widget.NewSlider(0, 1)
	volumeControl.Value = ae.MasterVolume
	volumeControl.OnChanged = func(value float64) {
		ae.SetMasterVolume(value)
	}

	// Oscillator controls
	var oscWidgets []fyne.CanvasObject
	for i, osc := range ae.Oscillators {
		index := i // capture the current index for closure

		// Frequency control
		frequencyEntry := widget.NewEntry()
		frequencyEntry.SetText(strconv.FormatFloat(osc.Frequency, 'f', 2, 64))
		frequencyEntry.OnSubmitted = func(s string) {
			if freq, err := strconv.ParseFloat(s, 64); err == nil {
				ae.SetOscillatorFrequency(index, freq)
			}
		}

		// Volume control
		volumeSlider := widget.NewSlider(0, 1)
		volumeSlider.Value = osc.Volume
		volumeSlider.OnChanged = func(value float64) {
			ae.SetOscillatorVolume(index, value)
		}

		// Waveform selection
		waveformSelect := widget.NewSelect([]string{"Sine", "Square", "Sawtooth", "Triangle"}, func(s string) {
			ae.SetOscillatorWaveform(index, waveformFromString(s))
		})
		waveformSelect.SetSelectedIndex(int(osc.Waveform))

		oscForm := widget.NewForm(
			widget.NewFormItem("Frequency", frequencyEntry),
			widget.NewFormItem("Volume", volumeSlider),
			widget.NewFormItem("Waveform", waveformSelect),
		)
		oscWidgets = append(oscWidgets, oscForm)
	}

	// Organize UI components
	oscillatorContainer := container.NewVBox(oscWidgets...)
	controlContainer := container.NewVBox(volumeControl, oscillatorContainer)
	keyboard := createKeyboardUI(ae)
	content := container.NewBorder(nil, keyboard, nil, nil, controlContainer)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func createKeyboardUI(ae *AudioEngine) fyne.CanvasObject {
	octaves := 3
	whiteKeyWidth := float32(64)
	blackKeyWidth := float32(40)
	blackKeyHeight := float32(100)
	whiteKeyHeight := float32(150)
	keysContainer := container.NewWithoutLayout()

	for o := 0; o < octaves; o++ {
		for i := 0; i < 7; i++ {
			note := noteFromKey(i)
			octave := o + 4 // Adjust octave base

			whiteKey := widget.NewButton("", func(note string, octave int) func() {
				return func() {
					ae.PlayNote(fmt.Sprintf("%s%d", note, octave), 0.5)
				}
			}(note, octave))
			whiteKey.Resize(fyne.NewSize(whiteKeyWidth, whiteKeyHeight))
			whiteKey.Move(fyne.NewPos(float32((i+o*7)*int(whiteKeyWidth)), 0))
			keysContainer.Add(whiteKey)

			// Black keys, omitting for E and B notes
			if i != 2 && i != 6 {
				blackKey := widget.NewButton("", func(note string, octave int) func() {
					return func() {
						ae.PlayNote(fmt.Sprintf("%s#%d", note, octave), 0.5)
					}
				}(note, octave))
				blackKey.Resize(fyne.NewSize(blackKeyWidth, blackKeyHeight))
				blackKey.Move(fyne.NewPos(float32((i+o*7)*int(whiteKeyWidth))+whiteKeyWidth-(blackKeyWidth/2), 0))
				keysContainer.Add(blackKey)
			}
		}
	}

	return container.NewPadded(keysContainer)
}

func noteFromKey(key int) string {
	switch key {
	case 0:
		return "C"
	case 1:
		return "D"
	case 2:
		return "E"
	case 3:
		return "F"
	case 4:
		return "G"
	case 5:
		return "A"
	case 6:
		return "B"
	default:
		return ""
	}
}

func waveformFromString(waveform string) Waveform {
	switch waveform {
	case "Sine":
		return Sine
	case "Square":
		return Square
	case "Sawtooth":
		return Sawtooth
	case "Triangle":
		return Triangle
	default:
		return Sine // Default case
	}
}
