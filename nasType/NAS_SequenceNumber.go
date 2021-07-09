// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// SequenceNumber 9.10
// SQN Row, sBit, len = [0, 0], 8 , 8
type SequenceNumber struct {
	Octet uint8
}

func NewSequenceNumber() (sequenceNumber *SequenceNumber) {
	sequenceNumber = &SequenceNumber{}
	return sequenceNumber
}

// SequenceNumber 9.10
// SQN Row, sBit, len = [0, 0], 8 , 8
func (a *SequenceNumber) GetSQN() (sQN uint8) {
	return a.Octet
}

// SequenceNumber 9.10
// SQN Row, sBit, len = [0, 0], 8 , 8
func (a *SequenceNumber) SetSQN(sQN uint8) {
	a.Octet = sQN
}
