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

func TestNasTypeNewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionAuthenticationResult, nas.MsgTypePDUSessionAuthenticationResult},
}

func TestNasTypeGetSetPDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity()
	for _, table := range nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
