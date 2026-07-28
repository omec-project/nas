// Copyright (C) 2026 Intel Corporation
// SPDX-License-Identifier: Apache-2.0

package nasType

// FeatureAuthorizationIndication 9.11.3.x
type FeatureAuthorizationIndication struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewFeatureAuthorizationIndication(iei uint8) (x *FeatureAuthorizationIndication) {
	x = &FeatureAuthorizationIndication{}
	x.SetIei(iei)
	return x
}

func (a *FeatureAuthorizationIndication) GetIei() (iei uint8) { return a.Iei }
func (a *FeatureAuthorizationIndication) SetIei(iei uint8)    { a.Iei = iei }
func (a *FeatureAuthorizationIndication) GetLen() (len uint8) { return a.Len }
func (a *FeatureAuthorizationIndication) SetLen(len uint8) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

func (a *FeatureAuthorizationIndication) GetFeatureAuthorizationIndication() []uint8 {
	v := make([]uint8, len(a.Buffer))
	copy(v, a.Buffer)
	return v
}

func (a *FeatureAuthorizationIndication) SetFeatureAuthorizationIndication(v []uint8) {
	copy(a.Buffer, v)
}
