// Copyright (c) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasConvert

import "testing"

func TestLadnToModels(t *testing.T) {
	tests := []struct {
		name     string
		buf      []uint8
		expected []string
	}{
		{
			name:     "single DNN",
			buf:      []uint8{0x00, 0x03, 'a', 'b', 'c'},
			expected: []string{"abc"},
		},
		{
			// declared DNN length exceeds the remaining buffer
			name:     "overlength DNN",
			buf:      []uint8{0x00, 0x05, 'a', 'b'},
			expected: nil,
		},
		{
			// a zero-length DNN entry must not spin forever
			name:     "zero length DNN",
			buf:      []uint8{0x00, 0x00, 'a', 'b'},
			expected: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LadnToModels(tc.buf)
			if len(got) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
