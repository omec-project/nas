// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// ServiceRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ServiceRejectMessageIdentity struct {
	Octet uint8
}

func NewServiceRejectMessageIdentity() (serviceRejectMessageIdentity *ServiceRejectMessageIdentity) {
	serviceRejectMessageIdentity = &ServiceRejectMessageIdentity{}
	return serviceRejectMessageIdentity
}

// ServiceRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *ServiceRejectMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// ServiceRejectMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *ServiceRejectMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
