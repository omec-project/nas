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

type nasTypeMessageAuthenticationCodeMACData struct {
	in  [4]uint8
	out [4]uint8
}

var nasTypeMessageAuthenticationCodeMACTable = []nasTypeMessageAuthenticationCodeMACData{
	{[4]uint8{0xff, 0xff, 0xff, 0xff}, [4]uint8{0xff, 0xff, 0xff, 0xff}},
}

func TestNasTypeNewMessageAuthenticationCode(t *testing.T) {
	a := nasType.NewMessageAuthenticationCode()
	assert.NotNil(t, a)
}

func TestNasTypeMessageAuthenticationCode(t *testing.T) {
	a := nasType.NewMessageAuthenticationCode()
	for _, table := range nasTypeMessageAuthenticationCodeMACTable {
		a.SetMAC(table.in)
		assert.Equal(t, table.out, a.GetMAC())
	}
}
