// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// PDUSESSIONESTABLISHMENTACCEPTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONESTABLISHMENTACCEPTMessageIdentity struct {
	Octet uint8
}

func NewPDUSESSIONESTABLISHMENTACCEPTMessageIdentity() (pDUSESSIONESTABLISHMENTACCEPTMessageIdentity *PDUSESSIONESTABLISHMENTACCEPTMessageIdentity) {
	pDUSESSIONESTABLISHMENTACCEPTMessageIdentity = &PDUSESSIONESTABLISHMENTACCEPTMessageIdentity{}
	return pDUSESSIONESTABLISHMENTACCEPTMessageIdentity
}

// PDUSESSIONESTABLISHMENTACCEPTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONESTABLISHMENTACCEPTMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// PDUSESSIONESTABLISHMENTACCEPTMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONESTABLISHMENTACCEPTMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
