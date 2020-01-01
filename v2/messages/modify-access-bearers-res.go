// Copyright 2019-2020 go-gtp authors. All rights reservev.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages

import "github.com/wmnsk/go-gtp/v2/ies"

// ModifyAccessBearersResponse is a ModifyAccessBearersResponse Header and its IEs above.
type ModifyAccessBearersResponse struct {
	*Header
	Cause                          *ies.IE
	BearerContextsModified         *ies.IE
	BearerContextsMarkedForRemoval *ies.IE
	Recovery                       *ies.IE
	IndicationFlags                *ies.IE
	SGWNodeLoadControlInformation  *ies.IE
	PrivateExtension               *ies.IE
	AdditionalIEs                  []*ies.IE
}

// NewModifyAccessBearersResponse creates a new ModifyAccessBearersResponse.
func NewModifyAccessBearersResponse(teid, seq uint32, ie ...*ies.IE) *ModifyAccessBearersResponse {
	m := &ModifyAccessBearersResponse{
		Header: NewHeader(
			NewHeaderFlags(2, 0, 1),
			MsgTypeModifyAccessBearersResponse, teid, seq, nil,
		),
	}

	for _, i := range ie {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			m.Cause = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsModified = i
			case 1:
				m.BearerContextsMarkedForRemoval = i
			}
		case ies.Recovery:
			m.Recovery = i
		case ies.Indication:
			m.IndicationFlags = i
		case ies.LoadControlInformation:
			m.SGWNodeLoadControlInformation = i
		case ies.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	m.SetLength()
	return m
}

// Marshal serializes ModifyAccessBearersResponse into bytes.
func (m *ModifyAccessBearersResponse) Marshal() ([]byte, error) {
	b := make([]byte, m.MarshalLen())
	if err := m.MarshalTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// MarshalTo serializes ModifyAccessBearersResponse into bytes.
func (m *ModifyAccessBearersResponse) MarshalTo(b []byte) error {
	if m.Header.Payload != nil {
		m.Header.Payload = nil
	}
	m.Header.Payload = make([]byte, m.MarshalLen()-m.Header.MarshalLen())

	offset := 0
	if ie := m.Cause; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.BearerContextsModified; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.BearerContextsMarkedForRemoval; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.Recovery; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.IndicationFlags; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.SGWNodeLoadControlInformation; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}
	if ie := m.PrivateExtension; ie != nil {
		if err := ie.MarshalTo(m.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	for _, ie := range m.AdditionalIEs {
		if ie == nil {
			continue
		}
		if err := ie.MarshalTo(m.Header.Payload[offset:]); err != nil {
			return err
		}
		offset += ie.MarshalLen()
	}

	m.Header.SetLength()
	return m.Header.MarshalTo(b)
}

// ParseModifyAccessBearersResponse decodes given bytes as ModifyAccessBearersResponse.
func ParseModifyAccessBearersResponse(b []byte) (*ModifyAccessBearersResponse, error) {
	m := &ModifyAccessBearersResponse{}
	if err := m.UnmarshalBinary(b); err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalBinary decodes given bytes as ModifyAccessBearersResponse.
func (m *ModifyAccessBearersResponse) UnmarshalBinary(b []byte) error {
	var err error
	m.Header, err = ParseHeader(b)
	if err != nil {
		return err
	}
	if len(m.Header.Payload) < 2 {
		return nil
	}

	decodedIEs, err := ies.ParseMultiIEs(m.Header.Payload)
	if err != nil {
		return err
	}
	for _, i := range decodedIEs {
		if i == nil {
			continue
		}
		switch i.Type {
		case ies.Cause:
			m.Cause = i
		case ies.BearerContext:
			switch i.Instance() {
			case 0:
				m.BearerContextsModified = i
			case 1:
				m.BearerContextsMarkedForRemoval = i
			}
		case ies.Recovery:
			m.Recovery = i
		case ies.Indication:
			m.IndicationFlags = i
		case ies.LoadControlInformation:
			m.SGWNodeLoadControlInformation = i
		case ies.PrivateExtension:
			m.PrivateExtension = i
		default:
			m.AdditionalIEs = append(m.AdditionalIEs, i)
		}
	}

	return nil
}

// MarshalLen returns the serial length in int.
func (m *ModifyAccessBearersResponse) MarshalLen() int {
	l := m.Header.MarshalLen() - len(m.Header.Payload)
	if ie := m.Cause; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.BearerContextsModified; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.BearerContextsMarkedForRemoval; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.Recovery; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.IndicationFlags; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.SGWNodeLoadControlInformation; ie != nil {
		l += ie.MarshalLen()
	}
	if ie := m.PrivateExtension; ie != nil {
		l += ie.MarshalLen()
	}

	for _, ie := range m.AdditionalIEs {
		if ie == nil {
			continue
		}
		l += ie.MarshalLen()
	}
	return l
}

// SetLength sets the length in Length field.
func (m *ModifyAccessBearersResponse) SetLength() {
	m.Header.Length = uint16(m.MarshalLen() - 4)
}

// MessageTypeName returns the name of protocol.
func (m *ModifyAccessBearersResponse) MessageTypeName() string {
	return "Modify Access Bearers Response"
}

// TEID returns the TEID in uint32.
func (m *ModifyAccessBearersResponse) TEID() uint32 {
	return m.Header.teid()
}
