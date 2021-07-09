// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// PDUSESSIONMODIFICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONMODIFICATIONCOMMANDMessageIdentity struct {
	Octet uint8
}

func NewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity() (pDUSESSIONMODIFICATIONCOMMANDMessageIdentity *PDUSESSIONMODIFICATIONCOMMANDMessageIdentity) {
	pDUSESSIONMODIFICATIONCOMMANDMessageIdentity = &PDUSESSIONMODIFICATIONCOMMANDMessageIdentity{}
	return pDUSESSIONMODIFICATIONCOMMANDMessageIdentity
}

// PDUSESSIONMODIFICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONMODIFICATIONCOMMANDMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// PDUSESSIONMODIFICATIONCOMMANDMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONMODIFICATIONCOMMANDMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
