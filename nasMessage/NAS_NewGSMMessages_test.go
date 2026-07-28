// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasMessage_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/omec-project/nas/v2/logger"
	"github.com/omec-project/nas/v2/nasMessage"
	"github.com/omec-project/nas/v2/nasType"
	"github.com/stretchr/testify/assert"
)

func TestNasTypeNewServiceLevelAuthenticationCommand(t *testing.T) {
	a := nasMessage.NewServiceLevelAuthenticationCommand(0)
	assert.NotNil(t, a)
}

func TestServiceLevelAuthenticationCommandEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: ServiceLevelAuthenticationCommand---")

	a := nasMessage.NewServiceLevelAuthenticationCommand(0)
	b := nasMessage.NewServiceLevelAuthenticationCommand(0)
	assert.NotNil(t, a)
	assert.NotNil(t, b)

	a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	a.PDUSessionID.SetPDUSessionID(0x01)
	a.PTI.SetPTI(0x01)
	a.SERVICELEVELAUTHENTICATIONCOMMANDMessageIdentity.SetMessageType(0xD8)

	// Set ServiceLevelAAContainer
	a.ServiceLevelAAContainer = nasType.ServiceLevelAAContainer{}
	a.ServiceLevelAAContainer.SetLen(4)
	copy(a.ServiceLevelAAContainer.Buffer, []byte{0x01, 0x02, 0x03, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeServiceLevelAuthenticationCommand(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	buff.Read(data)
	b.DecodeServiceLevelAuthenticationCommand(&data)
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("ServiceLevelAuthenticationCommand encode/decode mismatch")
	}
}

func TestNasTypeNewServiceLevelAuthenticationComplete(t *testing.T) {
	a := nasMessage.NewServiceLevelAuthenticationComplete(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewRemoteUEReport(t *testing.T) {
	a := nasMessage.NewRemoteUEReport(0)
	assert.NotNil(t, a)
}

func TestRemoteUEReportEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RemoteUEReport---")

	a := nasMessage.NewRemoteUEReport(0)
	b := nasMessage.NewRemoteUEReport(0)
	assert.NotNil(t, a)
	assert.NotNil(t, b)

	a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	a.PDUSessionID.SetPDUSessionID(0x01)
	a.PTI.SetPTI(0x02)
	a.REMOTEUEREPORTMessageIdentity.SetMessageType(0xDA)

	buff := new(bytes.Buffer)
	a.EncodeRemoteUEReport(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	buff.Read(data)
	b.DecodeRemoteUEReport(&data)
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RemoteUEReport encode/decode mismatch")
	}
}

func TestNasTypeNewRemoteUEReportResponse(t *testing.T) {
	a := nasMessage.NewRemoteUEReportResponse(0)
	assert.NotNil(t, a)
}

func TestRemoteUEReportResponseEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: RemoteUEReportResponse---")

	a := nasMessage.NewRemoteUEReportResponse(0)
	b := nasMessage.NewRemoteUEReportResponse(0)
	assert.NotNil(t, a)
	assert.NotNil(t, b)

	a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSSessionManagementMessage)
	a.PDUSessionID.SetPDUSessionID(0x01)
	a.PTI.SetPTI(0x02)
	a.REMOTEUEREPORTRESPONSEMessageIdentity.SetMessageType(0xDB)

	buff := new(bytes.Buffer)
	a.EncodeRemoteUEReportResponse(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	buff.Read(data)
	b.DecodeRemoteUEReportResponse(&data)
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("RemoteUEReportResponse encode/decode mismatch")
	}
}
