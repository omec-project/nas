// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// EthernetHeaderCompressionConfiguration 9.11.4.28
// Iei Row, sBit, len = [], 8, 8
// Len Row, sBit, len = [], 8, 8
// Value Row, sBit, len = [0, 0], 8, 8
type EthernetHeaderCompressionConfiguration struct {
	Iei   uint8
	Len   uint8
	Octet uint8
}

func NewEthernetHeaderCompressionConfiguration(iei uint8) (x *EthernetHeaderCompressionConfiguration) {
	x = &EthernetHeaderCompressionConfiguration{}
	x.SetIei(iei)
	return x
}

func (a *EthernetHeaderCompressionConfiguration) GetIei() (iei uint8) { return a.Iei }
func (a *EthernetHeaderCompressionConfiguration) SetIei(iei uint8)    { a.Iei = iei }
func (a *EthernetHeaderCompressionConfiguration) GetLen() (len uint8) { return a.Len }
func (a *EthernetHeaderCompressionConfiguration) SetLen(len uint8)    { a.Len = len }
func (a *EthernetHeaderCompressionConfiguration) GetEHCI() (ehci uint8) {
	return a.Octet & GetBitMask(2, 1) >> 1
}

func (a *EthernetHeaderCompressionConfiguration) SetEHCI(ehci uint8) {
	a.Octet = (a.Octet & 253) + ((ehci & 1) << 1)
}
