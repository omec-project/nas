// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package nasType_test

import (
	"testing"

	"github.com/free5gc/nas"
	"github.com/free5gc/nas/nasType"

	"github.com/stretchr/testify/assert"
)

func TestNasTypeNewPDUSESSIONESTABLISHMENTREJECTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREJECTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionEstablishmentReject, nas.MsgTypePDUSessionEstablishmentReject},
}

func TestNasTypeGetSetPDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREJECTMessageIdentity()
	for _, table := range nasTypePDUSESSIONESTABLISHMENTREJECTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
