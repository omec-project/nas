// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType

// SpareHalfOctetAndPayloadContainerType 9.11.3.40 9.5
// PayloadContainerType Row, sBit, len = [0, 0], 4 , 4
type SpareHalfOctetAndPayloadContainerType struct {
	Octet uint8
}

func NewSpareHalfOctetAndPayloadContainerType() (spareHalfOctetAndPayloadContainerType *SpareHalfOctetAndPayloadContainerType) {
	spareHalfOctetAndPayloadContainerType = &SpareHalfOctetAndPayloadContainerType{}
	return spareHalfOctetAndPayloadContainerType
}

// SpareHalfOctetAndPayloadContainerType 9.11.3.40 9.5
// PayloadContainerType Row, sBit, len = [0, 0], 4 , 4
func (a *SpareHalfOctetAndPayloadContainerType) GetPayloadContainerType() (payloadContainerType uint8) {
	return a.Octet & GetBitMask(4, 0)
}

// SpareHalfOctetAndPayloadContainerType 9.11.3.40 9.5
// PayloadContainerType Row, sBit, len = [0, 0], 4 , 4
func (a *SpareHalfOctetAndPayloadContainerType) SetPayloadContainerType(payloadContainerType uint8) {
	a.Octet = (a.Octet & 240) + (payloadContainerType & 15)
}
