// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// IdentityRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type IdentityRequestMessageIdentity struct {
	Octet uint8
}

func NewIdentityRequestMessageIdentity() (identityRequestMessageIdentity *IdentityRequestMessageIdentity) {
	identityRequestMessageIdentity = &IdentityRequestMessageIdentity{}
	return identityRequestMessageIdentity
}

// IdentityRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *IdentityRequestMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// IdentityRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *IdentityRequestMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
