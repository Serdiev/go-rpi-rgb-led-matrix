//go:build !linux
// +build !linux

package rgbmatrix

import "C"
import "fmt"

type uint32_t uint32

// NewRGBLedMatrix returns a new matrix using the given size and config
func NewRGBLedMatrix(config *HardwareConfig) (c Matrix, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("error creating matrix: %v", r)
			}
		}
	}()

	if isJulienEmulator() {
		return buildJulienMatrix(config), nil
	}
	if isMatrixEmulator() {
		return buildMatrixEmulator(config), nil
	}
	if isTerminalMatrixEmulator() {
		return buildTerminalMatrixEmulator(config), nil
	}

	panic("No emulator found")
}
