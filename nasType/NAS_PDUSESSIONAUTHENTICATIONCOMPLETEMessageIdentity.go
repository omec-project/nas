// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType

// PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity struct {
	Octet uint8
}

func NewPDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity() (pDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity *PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity) {
	pDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity = &PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity{}
	return pDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity
}

// PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *PDUSESSIONAUTHENTICATIONCOMPLETEMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
