// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/omec-project/nas/logger"
	"github.com/omec-project/nas/nasMessage"
	"github.com/omec-project/nas/nasType"
	"github.com/omec-project/openapi/models"
)

func ModelsToSessionAMBR(ambr *models.Ambr) (sessAmbr nasType.SessionAMBR) {
	fmt.Println(ambr)

	uplink := strings.Split(ambr.Uplink, " ")
	if bitRate, err := strconv.ParseUint(uplink[0], 10, 16); err != nil {
		logger.ConvertLog.Warnf("uplink AMBR parse failed: %+v", err)
	} else {
		var bitRateBytes [2]byte
		binary.BigEndian.PutUint16(bitRateBytes[:], uint16(bitRate))
		sessAmbr.SetSessionAMBRForUplink(bitRateBytes)
	}
	sessAmbr.SetUnitForSessionAMBRForUplink(strToAMBRUnit(uplink[1]))

	downlink := strings.Split(ambr.Downlink, " ")
	if bitRate, err := strconv.ParseUint(downlink[0], 10, 16); err != nil {
		logger.ConvertLog.Warnf("downlink AMBR parse failed: %+v", err)
	} else {
		var bitRateBytes [2]byte
		binary.BigEndian.PutUint16(bitRateBytes[:], uint16(bitRate))
		sessAmbr.SetSessionAMBRForDownlink(bitRateBytes)
	}
	sessAmbr.SetUnitForSessionAMBRForDownlink(strToAMBRUnit(downlink[1]))
	return
}

func strToAMBRUnit(unit string) uint8 {
	switch unit {
	case "bps":
		return nasMessage.SessionAMBRUnitNotUsed
	case "Kbps":
		return nasMessage.SessionAMBRUnit1Kbps
	case "Mbps":
		return nasMessage.SessionAMBRUnit1Mbps
	case "Gbps":
		return nasMessage.SessionAMBRUnit1Gbps
	case "Tbps":
		return nasMessage.SessionAMBRUnit1Tbps
	case "Pbps":
		return nasMessage.SessionAMBRUnit1Pbps
	}
	return nasMessage.SessionAMBRUnitNotUsed
}
