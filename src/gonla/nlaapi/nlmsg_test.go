// -*- coding: utf-8 -*-

// Copyright (C) 2017 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nlaapi

import (
	"bytes"
	"gonla/nlamsg"
	"syscall"
	"testing"
)

func TestNlMsghdrFromApi(t *testing.T) {
	h := NlMsghdr{
		Len:   0x01020304,
		Type:  0x11121314,
		Flags: 0x21222324,
		Seq:   0x31323334,
		Pid:   0x41424344,
	}

	v := h.ToNative()

	if v == nil {
		t.Errorf("NlMsghdrFromApi nil. %v", v)
	}

	vv := syscall.NlMsghdr{
		Len:   0x01020304,
		Type:  0x1314,
		Flags: 0x2324,
		Seq:   0x31323334,
		Pid:   0x41424344,
	}

	if *v != vv {
		t.Errorf("NlMsghdrFromApi unmatch. %v", v)
	}
}

func TestNlMsghdrToApi(t *testing.T) {
	h := syscall.NlMsghdr{
		Len:   0x01020304,
		Type:  0x1314,
		Flags: 0x2324,
		Seq:   0x31323334,
		Pid:   0x41424344,
	}

	v := NewNlMsghdrFromNative(&h)

	if v == nil {
		t.Errorf("NlMsghdrToApi nil. %v", v)
	}

	vv := NlMsghdr{
		Len:   0x01020304,
		Type:  0x00001314,
		Flags: 0x00002324,
		Seq:   0x31323334,
		Pid:   0x41424344,
	}

	if *v != vv {
		t.Errorf("NlMsghdrToApi unmatch. %v", v)
	}
}

func TestNetinkMessageFromApi(t *testing.T) {
	header := &NlMsghdr{
		Len:   0x01020304,
		Type:  0x11121314,
		Flags: 0x21222324,
		Seq:   0x31323334,
		Pid:   0x41424344,
	}
	data := []byte{0x51, 0x52, 0x53}
	m := NetlinkMessage{
		Header: header,
		Data:   data,
	}

	v := m.ToNative()

	if v == nil {
		t.Errorf("NetinkMessageFromApi nil. %v", v)
	}

	hh := syscall.NlMsghdr{
		Len:   0x01020304,
		Type:  0x1314,
		Flags: 0x2324,
		Seq:   0x31323334,
		Pid:   0x41424344,
	}

	if v.Header != hh {
		t.Errorf("NetinkMessageFromApi Header unmatch. %v", v)
	}

	if bytes.Compare(data, v.Data) != 0 {
		t.Errorf("NetinkMessageFromApi Data unmatch. %v", v.Data)
	}
}

func TestNetlinkMessageToApi(t *testing.T) {
	h := syscall.NlMsghdr{
		Len:   0x01020304,
		Type:  0x1314,
		Flags: 0x2324,
		Seq:   0x31323334,
		Pid:   0x41424344,
	}
	data := []byte{0x51, 0x52, 0x53}
	m := nlamsg.NetlinkMessage{
		NId: 0,
		NetlinkMessage: syscall.NetlinkMessage{
			Header: h,
			Data:   data,
		},
	}

	v := NewNetlinkMessageFromNative(&m)

	if v == nil {
		t.Errorf("NetlinkMessageToApi nil")
	}

	hh := NlMsghdr{
		Len:   0x01020304,
		Type:  0x00001314,
		Flags: 0x00002324,
		Seq:   0x31323334,
		Pid:   0x41424344,
	}

	if *v.Header != hh {
		t.Errorf("NetlinkMessageToApi Header unmatch. %v", v)
	}

	if bytes.Compare(v.Data, data) != 0 {
		t.Errorf("NetlinkMessageToApi Data unmatch. %v", v)
	}
}
