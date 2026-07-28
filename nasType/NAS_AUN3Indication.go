// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// AUN3Indication 9.11.3.104
// AUN3Indication Row, sBit, len = [0, INF], 8 , INF
type AUN3Indication struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewAUN3Indication(iei uint8) (x *AUN3Indication) {
	x = &AUN3Indication{}
	x.SetIei(iei)
	return x
}

func (a *AUN3Indication) GetIei() (iei uint8) {
	return a.Iei
}

func (a *AUN3Indication) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *AUN3Indication) GetLen() (len uint16) {
	return a.Len
}

func (a *AUN3Indication) SetLen(len uint16) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

func (a *AUN3Indication) GetAUN3Indication() (contents []uint8) {
	contents = make([]uint8, len(a.Buffer))
	copy(contents, a.Buffer)
	return contents
}

func (a *AUN3Indication) SetAUN3Indication(contents []uint8) {
	copy(a.Buffer, contents)
}
