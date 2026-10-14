// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasConvert

import (
	"encoding/hex"

	"github.com/omec-project/nas/v2/logger"
)

func PlmnIDToString(nasBuf []byte) string {
	if len(nasBuf) < 3 {
		logger.ConvertLog.Errorf("invalid PLMN ID buffer length: %d", len(nasBuf))
		return ""
	}

	mccDigit1 := nasBuf[0] & 0x0f
	mccDigit2 := (nasBuf[0] & 0xf0) >> 4
	mccDigit3 := (nasBuf[1] & 0x0f)

	mncDigit1 := (nasBuf[2] & 0x0f)
	mncDigit2 := (nasBuf[2] & 0xf0) >> 4
	mncDigit3 := (nasBuf[1] & 0xf0) >> 4

	tmpBytes := []byte{(mccDigit1 << 4) | mccDigit2, (mccDigit3 << 4) | mncDigit1, (mncDigit2 << 4) | mncDigit3}

	plmnID := hex.EncodeToString(tmpBytes)
	if plmnID[5] == 'f' {
		plmnID = plmnID[:5] // get plmnID[0~4]
	}
	return plmnID
}
