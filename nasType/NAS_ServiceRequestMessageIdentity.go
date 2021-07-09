// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// ServiceRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type ServiceRequestMessageIdentity struct {
	Octet uint8
}

func NewServiceRequestMessageIdentity() (serviceRequestMessageIdentity *ServiceRequestMessageIdentity) {
	serviceRequestMessageIdentity = &ServiceRequestMessageIdentity{}
	return serviceRequestMessageIdentity
}

// ServiceRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *ServiceRequestMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// ServiceRequestMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *ServiceRequestMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
