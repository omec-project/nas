// SPDX-FileCopyrightText: 2025 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasConvert

import (
	"testing"
)

func TestPeiToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte // Hexadecimal representation of the input
		expected string // Expected output
	}{
		{
			name:     "Valid IMEI",
			input:    []byte{0x3b, 0x65, 0x39, 0x08, 0x53, 0x46, 0x83, 0x85}, // IMEI: 356938035643859
			expected: "imei-356938035643858",
		},
		{
			name:     "Valid IMEISV",
			input:    []byte{0x35, 0x55, 0x12, 0x18, 0x01, 0x31, 0x66, 0x44, 0xF6}, // IMEISV: 3549774394518016
			expected: "imeisv-3552181101366446",
		},
		{
			name:     "Invalid IMEI length",
			input:    []byte{0x3b, 0x65, 0x39, 0x08, 0x53, 0x46, 0x83}, // IMEI: 13 digits
			expected: "",
		},
		{
			name:     "Invalid IMEISV length",
			input:    []byte{0x35, 0x65, 0x39, 0x08, 0x53, 0x46, 0x83, 0x95, 0x01, 0x01}, // IMEISV: 18 digits
			expected: "",
		},
		{
			name:     "Invalid character in IMEI",
			input:    []byte{0x9b, 0x09, 0x0a, 0x80, 0x26, 0x74, 0x81, 0x35}, // invalid character 'a'
			expected: "",
		},
		{
			name:     "Invalid character in IMEISV",
			input:    []byte{0x35, 0x45, 0x79, 0x47, 0xd3, 0x54, 0x81, 0x10, 0xF6}, // invalid character 'd'
			expected: "",
		},
		{
			name:     "Invalid TAC/SNR (Luhn check fails)",
			input:    []byte{0x3b, 0x65, 0x39, 0x58, 0x53, 0x46, 0x83, 0x15}, // IMEI: Fails Luhn check
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PeiToString(tt.input)
			if result != tt.expected {
				t.Errorf("PeiToString() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
