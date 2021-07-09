// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// AuthenticationFailureMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type AuthenticationFailureMessageIdentity struct {
	Octet uint8
}

func NewAuthenticationFailureMessageIdentity() (authenticationFailureMessageIdentity *AuthenticationFailureMessageIdentity) {
	authenticationFailureMessageIdentity = &AuthenticationFailureMessageIdentity{}
	return authenticationFailureMessageIdentity
}

// AuthenticationFailureMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *AuthenticationFailureMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// AuthenticationFailureMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *AuthenticationFailureMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
