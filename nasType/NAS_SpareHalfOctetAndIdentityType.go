// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// SpareHalfOctetAndIdentityType 9.11.3.3 9.5
// TypeOfIdentity Row, sBit, len = [0, 0], 3 , 3
type SpareHalfOctetAndIdentityType struct {
	Octet uint8
}

func NewSpareHalfOctetAndIdentityType() (spareHalfOctetAndIdentityType *SpareHalfOctetAndIdentityType) {
	spareHalfOctetAndIdentityType = &SpareHalfOctetAndIdentityType{}
	return spareHalfOctetAndIdentityType
}

// SpareHalfOctetAndIdentityType 9.11.3.3 9.5
// TypeOfIdentity Row, sBit, len = [0, 0], 3 , 3
func (a *SpareHalfOctetAndIdentityType) GetTypeOfIdentity() (typeOfIdentity uint8) {
	return a.Octet & GetBitMask(3, 0)
}

// SpareHalfOctetAndIdentityType 9.11.3.3 9.5
// TypeOfIdentity Row, sBit, len = [0, 0], 3 , 3
func (a *SpareHalfOctetAndIdentityType) SetTypeOfIdentity(typeOfIdentity uint8) {
	a.Octet = (a.Octet & 248) + (typeOfIdentity & 7)
}
