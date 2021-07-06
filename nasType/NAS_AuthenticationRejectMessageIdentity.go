// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// AuthenticationRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationRejectMessageIdentity struct {
	Octet uint8
}

func NewAuthenticationRejectMessageIdentity() (authenticationRejectMessageIdentity *AuthenticationRejectMessageIdentity) {
	authenticationRejectMessageIdentity = &AuthenticationRejectMessageIdentity{}
	return authenticationRejectMessageIdentity
}

// AuthenticationRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *AuthenticationRejectMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// AuthenticationRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *AuthenticationRejectMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
