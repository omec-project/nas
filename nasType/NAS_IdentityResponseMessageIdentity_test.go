// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType_test

import (
	"testing"

	"github.com/omec-project/nas"
	"github.com/omec-project/nas/nasType"
	"github.com/stretchr/testify/assert"
)

type nasTypeIdentityResponseMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeIdentityResponseMessageIdentityTable = []nasTypeIdentityResponseMessageIdentityData{
	{nas.MsgTypeIdentityResponse, nas.MsgTypeIdentityResponse},
}

func TestNasTypeNewIdentityResponseMessageIdentity(t *testing.T) {
	a := nasType.NewIdentityResponseMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeIdentityResponseMessageIdentity(t *testing.T) {
	a := nasType.NewIdentityResponseMessageIdentity()
	for _, table := range nasTypeIdentityResponseMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
