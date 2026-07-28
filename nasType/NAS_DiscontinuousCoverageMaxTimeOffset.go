// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// DiscontinuousCoverageMaxTimeOffset 9.11.x
type DiscontinuousCoverageMaxTimeOffset struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewDiscontinuousCoverageMaxTimeOffset(iei uint8) (x *DiscontinuousCoverageMaxTimeOffset) {
	x = &DiscontinuousCoverageMaxTimeOffset{}
	x.SetIei(iei)
	return x
}

func (a *DiscontinuousCoverageMaxTimeOffset) GetIei() (iei uint8) { return a.Iei }
func (a *DiscontinuousCoverageMaxTimeOffset) SetIei(iei uint8)    { a.Iei = iei }
func (a *DiscontinuousCoverageMaxTimeOffset) GetLen() (len uint8) { return a.Len }
func (a *DiscontinuousCoverageMaxTimeOffset) SetLen(len uint8) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

func (a *DiscontinuousCoverageMaxTimeOffset) GetDiscontinuousCoverageMaxTimeOffset() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}

func (a *DiscontinuousCoverageMaxTimeOffset) SetDiscontinuousCoverageMaxTimeOffset(v []uint8) {
	copy(a.Buffer, v)
}
