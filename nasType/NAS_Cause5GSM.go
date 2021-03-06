// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType

// Cause5GSM 9.11.4.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
type Cause5GSM struct {
	Iei   uint8
	Octet uint8
}

func NewCause5GSM(iei uint8) (cause5GSM *Cause5GSM) {
	cause5GSM = &Cause5GSM{}
	cause5GSM.SetIei(iei)
	return cause5GSM
}

// Cause5GSM 9.11.4.2
// Iei Row, sBit, len = [], 8, 8
func (a *Cause5GSM) GetIei() (iei uint8) {
	return a.Iei
}

// Cause5GSM 9.11.4.2
// Iei Row, sBit, len = [], 8, 8
func (a *Cause5GSM) SetIei(iei uint8) {
	a.Iei = iei
}

// Cause5GSM 9.11.4.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
func (a *Cause5GSM) GetCauseValue() (causeValue uint8) {
	return a.Octet
}

// Cause5GSM 9.11.4.2
// CauseValue Row, sBit, len = [0, 0], 8 , 8
func (a *Cause5GSM) SetCauseValue(causeValue uint8) {
	a.Octet = causeValue
}
