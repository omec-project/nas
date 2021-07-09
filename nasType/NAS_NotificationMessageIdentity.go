// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// NotificationMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
type NotificationMessageIdentity struct {
	Octet uint8
}

func NewNotificationMessageIdentity() (notificationMessageIdentity *NotificationMessageIdentity) {
	notificationMessageIdentity = &NotificationMessageIdentity{}
	return notificationMessageIdentity
}

// NotificationMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *NotificationMessageIdentity) GetMessageType() (messageType uint8) {
	return a.Octet
}

// NotificationMessageIdentity 9.7
// MessageType Row, sBit, len = [0, 0], 8 , 8
func (a *NotificationMessageIdentity) SetMessageType(messageType uint8) {
	a.Octet = messageType
}
