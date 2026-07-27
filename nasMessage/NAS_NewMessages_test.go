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

func TestNasTypeNewNetworkSliceSpecificAuthenticationCommand(t *testing.T) {
	a := nasMessage.NewNetworkSliceSpecificAuthenticationCommand(0)
	assert.NotNil(t, a)
}

func TestNetworkSliceSpecificAuthenticationCommandEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: NetworkSliceSpecificAuthenticationCommand---")

	a := nasMessage.NewNetworkSliceSpecificAuthenticationCommand(0)
	b := nasMessage.NewNetworkSliceSpecificAuthenticationCommand(0)
	assert.NotNil(t, a)
	assert.NotNil(t, b)

	a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(0x00)
	a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0x00)
	a.NETWORKSLICESPECIFICAUTHENTICATIONCOMMANDMessageIdentity.SetMessageType(0x50)

	// Set S-NSSAI (1 byte SST)
	a.SNSSAI.Len = 1
	a.SNSSAI.Octet[0] = 0x01

	// Set EAP message
	a.EAPMessage = nasType.EAPMessage{}
	a.EAPMessage.SetLen(4)
	copy(a.EAPMessage.Buffer, []byte{0x02, 0x01, 0x00, 0x04})

	buff := new(bytes.Buffer)
	a.EncodeNetworkSliceSpecificAuthenticationCommand(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	buff.Read(data)
	b.DecodeNetworkSliceSpecificAuthenticationCommand(&data)
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("NetworkSliceSpecificAuthenticationCommand encode/decode mismatch")
	}
}

func TestNasTypeNewNetworkSliceSpecificAuthenticationComplete(t *testing.T) {
	a := nasMessage.NewNetworkSliceSpecificAuthenticationComplete(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewNetworkSliceSpecificAuthenticationResult(t *testing.T) {
	a := nasMessage.NewNetworkSliceSpecificAuthenticationResult(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewRelayKeyRequest(t *testing.T) {
	a := nasMessage.NewRelayKeyRequest(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewRelayKeyAccept(t *testing.T) {
	a := nasMessage.NewRelayKeyAccept(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewRelayKeyReject(t *testing.T) {
	a := nasMessage.NewRelayKeyReject(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewRelayAuthenticationRequest(t *testing.T) {
	a := nasMessage.NewRelayAuthenticationRequest(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewRelayAuthenticationResponse(t *testing.T) {
	a := nasMessage.NewRelayAuthenticationResponse(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewControlPlaneServiceRequest(t *testing.T) {
	a := nasMessage.NewControlPlaneServiceRequest(0)
	assert.NotNil(t, a)
}

func TestControlPlaneServiceRequestEncodeDecode(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: ControlPlaneServiceRequest---")

	a := nasMessage.NewControlPlaneServiceRequest(0)
	b := nasMessage.NewControlPlaneServiceRequest(0)
	assert.NotNil(t, a)
	assert.NotNil(t, b)

	a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(nasMessage.Epd5GSMobilityManagementMessage)
	a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(0x00)
	a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(0x00)
	a.CONTROLPLANESERVICEREQUESTMessageIdentity.SetMessageType(0x4F)
	a.ControlPlaneServiceTypeAndNgksi.SetControlPlaneServiceType(0x01)
	a.ControlPlaneServiceTypeAndNgksi.SetNasKeySetIdentifiler(0x07)

	buff := new(bytes.Buffer)
	a.EncodeControlPlaneServiceRequest(buff)
	logger.NasMsgLog.Debugln("encode:", a)

	data := make([]byte, buff.Len())
	buff.Read(data)
	b.DecodeControlPlaneServiceRequest(&data)
	logger.NasMsgLog.Debugln("decode:", b)

	if reflect.DeepEqual(a, b) != true {
		t.Errorf("ControlPlaneServiceRequest encode/decode mismatch")
	}
}
