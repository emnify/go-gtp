// Copyright 2019-2020 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	v1 "github.com/wmnsk/go-gtp/v1"
	"github.com/wmnsk/go-gtp/v1/ies"
	"github.com/wmnsk/go-gtp/v1/messages"
	"github.com/wmnsk/go-gtp/v1/testutils"
)

func TestCreatePDPContextRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: messages.NewCreatePDPContextRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewIMSI("123450123456789"),
				ies.NewRouteingAreaIdentity("123", "45", 0x1111, 0x22),
				ies.NewRecovery(254),
				ies.NewSelectionMode(v1.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
				ies.NewTEIDDataI(0xdeadbeef),
				ies.NewTEIDCPlane(0xdeadbeef),
				ies.NewNSAPI(5),
				ies.NewEndUserAddressIPv4(""),
				ies.NewAccessPointName("some.apn.example"),
				ies.NewProtocolConfigurationOptions(
					0, ies.NewConfigurationProtocolOption(1, []byte{0xde, 0xad, 0xbe, 0xef}),
				),
				ies.NewGSNAddress("1.1.1.1"),
				ies.NewGSNAddress("2.2.2.2"),
				ies.NewMSISDN("123412345678"),
				ies.NewQoSProfile([]byte{0xde, 0xad, 0xbe, 0xef}), // XXX - Implement!
				ies.NewCommonFlags(0, 0, 1, 0, 0, 0, 0, 0),
				ies.NewRATType(v1.RatTypeUTRAN),
				ies.NewUserLocationInformationWithSAI("123", "45", 0x1111, 0x2222),
				ies.NewMSTimeZone(0x00, 0x00),
			),
			Serialized: []byte{
				// Header
				0x32, 0x10, 0x00, 0x7f, 0x11, 0x22, 0x33, 0x44,
				0x00, 0x01, 0x00, 0x00,
				// IMSI
				0x02, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9,
				// RAI
				0x03, 0x21, 0xf3, 0x54, 0x11, 0x11, 0x22,
				// Recovery
				0x0e, 0xfe,
				// Selection Mode
				0x0f, 0xf0,
				// TEID-U
				0x10, 0xde, 0xad, 0xbe, 0xef,
				// TEID-C
				0x11, 0xde, 0xad, 0xbe, 0xef,
				// NSAPI
				0x14, 0x05,
				// End User Address
				0x80, 0x00, 0x02, 0xf1, 0x21,
				// APN
				0x83, 0x00, 0x11, 0x04, 0x73, 0x6f, 0x6d, 0x65, 0x03, 0x61, 0x70, 0x6e, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				// PCO
				0x84, 0x00, 0x08, 0x80, 0x00, 0x01, 0x04,
				0xde, 0xad, 0xbe, 0xef,
				// GSN Address
				0x85, 0x00, 0x04, 0x01, 0x01, 0x01, 0x01,
				// GSN Address
				0x85, 0x00, 0x04, 0x02, 0x02, 0x02, 0x02,
				// MSISDN
				0x86, 0x00, 0x07, 0x91, 0x21, 0x43, 0x21, 0x43,
				0x65, 0x87,
				// QoS
				0x87, 0x00, 0x04, 0xde, 0xad, 0xbe, 0xef,
				/* XXX - implement QoSProfile!
				0x87, 0x00, 0x0f, 0x02, 0x0b, 0x92, 0x1f, 0x73,
				0x96, 0xff, 0xff, 0x94, 0xf9, 0xff, 0xff, 0x00,
				0x6a, 0x00,
				*/
				// Common Flags
				0x94, 0x00, 0x01, 0x20,
				// RAT Type
				0x97, 0x00, 0x01, 0x01,
				// ULI
				0x98, 0x00, 0x08, 0x01, 0x21, 0xf3, 0x54, 0x11,
				0x11, 0x22, 0x22,
				// MS Time Zone
				0x99, 0x00, 0x02, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := messages.ParseCreatePDPContextRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
