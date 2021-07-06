// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType_test

import (
	"testing"

	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/nas/nasType"

	"github.com/stretchr/testify/assert"
)

type nasTypeRequestMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeRequestMessageIdentityTable = []nasTypeRequestMessageIdentityData{
	{nasMessage.AuthenticationRequestEAPMessageType, nasMessage.AuthenticationRequestEAPMessageType},
}

func TestNasTypeNewAuthenticationRequestMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationRequestMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetAuthenticationRequestMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationRequestMessageIdentity()
	for _, table := range nasTypeRequestMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
