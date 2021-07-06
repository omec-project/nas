// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// PDUSESSIONMODIFICATIONREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONMODIFICATIONREJECTMessageIdentity struct {
	Octet uint8
}

func NewPDUSESSIONMODIFICATIONREJECTMessageIdentity() (pDUSESSIONMODIFICATIONREJECTMessageIdentity *PDUSESSIONMODIFICATIONREJECTMessageIdentity) {
	pDUSESSIONMODIFICATIONREJECTMessageIdentity = &PDUSESSIONMODIFICATIONREJECTMessageIdentity{}
	return pDUSESSIONMODIFICATIONREJECTMessageIdentity
}

// PDUSESSIONMODIFICATIONREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONMODIFICATIONREJECTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// PDUSESSIONMODIFICATIONREJECTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONMODIFICATIONREJECTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
