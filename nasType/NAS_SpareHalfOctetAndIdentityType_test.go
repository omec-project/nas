// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType_test

import (
	"testing"

	"github.com/free5gc/nas/nasType"

	"github.com/stretchr/testify/assert"
)

type nasTypeIdentityTypeAndSpareHalfOctetData struct {
	in  uint8
	out uint8
}

var nasTypeIdentityTypeAndSpareHalfOctetTable = []nasTypeIdentityTypeAndSpareHalfOctetData{
	{0x07, 0x07},
}

func TestNasTypeNewSpareHalfOctetAndIdentityType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndIdentityType()
	assert.NotNil(t, a)
}

func TestNasTypeIdentityTypeAndSpareHalfOctet(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndIdentityType()
	for _, table := range nasTypeIdentityTypeAndSpareHalfOctetTable {
		a.SetTypeOfIdentity(table.in)
		assert.Equal(t, table.out, a.GetTypeOfIdentity())
	}
}
