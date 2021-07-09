// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// ULNASTRANSPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ULNASTRANSPORTMessageIdentity struct {
	Octet uint8
}

func NewULNASTRANSPORTMessageIdentity() (uLNASTRANSPORTMessageIdentity *ULNASTRANSPORTMessageIdentity) {
	uLNASTRANSPORTMessageIdentity = &ULNASTRANSPORTMessageIdentity{}
	return uLNASTRANSPORTMessageIdentity
}

// ULNASTRANSPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *ULNASTRANSPORTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// ULNASTRANSPORTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *ULNASTRANSPORTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
