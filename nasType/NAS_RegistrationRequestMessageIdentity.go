// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// RegistrationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type RegistrationRequestMessageIdentity struct {
	Octet uint8
}

func NewRegistrationRequestMessageIdentity() (registrationRequestMessageIdentity *RegistrationRequestMessageIdentity) {
	registrationRequestMessageIdentity = &RegistrationRequestMessageIdentity{}
	return registrationRequestMessageIdentity
}

// RegistrationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RegistrationRequestMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// RegistrationRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *RegistrationRequestMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
