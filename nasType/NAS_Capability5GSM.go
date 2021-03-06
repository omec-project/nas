// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType

// Capability5GSM 9.11.4.1
// MH6PDU Row, sBit, len = [0, 0], 2 , 1
// RqoS Row, sBit, len = [0, 0], 1 , 1
// Spare Row, sBit, len = [1, 12], 8 , 96
type Capability5GSM struct {
	Iei   uint8
	Len   uint8
	Octet [13]uint8
}

func NewCapability5GSM(iei uint8) (capability5GSM *Capability5GSM) {
	capability5GSM = &Capability5GSM{}
	capability5GSM.SetIei(iei)
	return capability5GSM
}

// Capability5GSM 9.11.4.1
// Iei Row, sBit, len = [], 8, 8
func (a *Capability5GSM) GetIei() (iei uint8) {
	return a.Iei
}

// Capability5GSM 9.11.4.1
// Iei Row, sBit, len = [], 8, 8
func (a *Capability5GSM) SetIei(iei uint8) {
	a.Iei = iei
}

// Capability5GSM 9.11.4.1
// Len Row, sBit, len = [], 8, 8
func (a *Capability5GSM) GetLen() (len uint8) {
	return a.Len
}

// Capability5GSM 9.11.4.1
// Len Row, sBit, len = [], 8, 8
func (a *Capability5GSM) SetLen(len uint8) {
	a.Len = len
}

// Capability5GSM 9.11.4.1
// MH6PDU Row, sBit, len = [0, 0], 2 , 1
func (a *Capability5GSM) GetMH6PDU() (mH6PDU uint8) {
	return a.Octet[0] & GetBitMask(2, 1) >> (1)
}

// Capability5GSM 9.11.4.1
// MH6PDU Row, sBit, len = [0, 0], 2 , 1
func (a *Capability5GSM) SetMH6PDU(mH6PDU uint8) {
	a.Octet[0] = (a.Octet[0] & 253) + ((mH6PDU & 1) << 1)
}

// Capability5GSM 9.11.4.1
// RqoS Row, sBit, len = [0, 0], 1 , 1
func (a *Capability5GSM) GetRqoS() (rqoS uint8) {
	return a.Octet[0] & GetBitMask(1, 0)
}

// Capability5GSM 9.11.4.1
// RqoS Row, sBit, len = [0, 0], 1 , 1
func (a *Capability5GSM) SetRqoS(rqoS uint8) {
	a.Octet[0] = (a.Octet[0] & 254) + (rqoS & 1)
}

// Capability5GSM 9.11.4.1
// Spare Row, sBit, len = [1, 12], 8 , 96
func (a *Capability5GSM) GetSpare() (spare [12]uint8) {
	copy(spare[:], a.Octet[1:13])
	return spare
}

// Capability5GSM 9.11.4.1
// Spare Row, sBit, len = [1, 12], 8 , 96
func (a *Capability5GSM) SetSpare(spare [12]uint8) {
	copy(a.Octet[1:13], spare[:])
}
