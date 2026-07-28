// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// PayloadContainerInformation 9.11.3.106
// Iei Row, sBit, len = [0, 0], 8 , 4
// Value Row, sBit, len = [0, 0], 4 , 4
type PayloadContainerInformation struct {
	Octet uint8
}

func NewPayloadContainerInformation(iei uint8) (x *PayloadContainerInformation) {
	x = &PayloadContainerInformation{}
	x.SetIei(iei)
	return x
}

func (a *PayloadContainerInformation) GetIei() (iei uint8) {
	return a.Octet & GetBitMask(8, 4) >> 4
}

func (a *PayloadContainerInformation) SetIei(iei uint8) {
	a.Octet = (a.Octet & 0x0F) | ((iei & 0x0F) << 4)
}

func (a *PayloadContainerInformation) GetValue() uint8 {
	return a.Octet & GetBitMask(4, 0)
}

func (a *PayloadContainerInformation) SetValue(v uint8) {
	a.Octet = (a.Octet & 0xF0) | (v & 0x0F)
}
