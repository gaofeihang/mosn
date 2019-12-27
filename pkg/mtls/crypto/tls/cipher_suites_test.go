/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tls

import (
	"bytes"
	"crypto/cipher"
	"testing"
)

func TestSM3(t *testing.T) {
	key := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	res := []byte{0xf6, 0x86, 0x59, 0xac, 0x70, 0x88, 0xe2, 0x7e, 0x9c, 0xac, 0xc0, 0xb6, 0x8d, 0xf8, 0x45, 0xff,
		0x83, 0x20, 0x37, 0x18, 0x96, 0xa5, 0xfa, 0xe0, 0x29, 0x9f, 0xdc, 0x68, 0xab, 0xbe, 0x5d, 0x91}

	sm3 := macSM3(VersionTLS11, key)
	if sm3.Size() != 32 {
		t.Error("sm3.Size() must be 32")
	}

	buf := make([]byte, 16)
	buf = sm3.MAC(buf, nil, nil, data, nil)

	if !bytes.Equal(res, buf) {
		t.Error("sm3 mac error")
	}
}

func TestSM4(t *testing.T) {
	key := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}
	iv := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

	encryptor := cipherSM4(key, iv, false).(cipher.BlockMode)
	decryptor := cipherSM4(key, iv, true).(cipher.BlockMode)

	buf1 := make([]byte, 16)
	buf2 := make([]byte, 16)

	encryptor.CryptBlocks(buf1, data)
	decryptor.CryptBlocks(buf2, buf1)

	if !bytes.Equal(data, buf2) {
		t.Error("data != decrypt(encrypt(data))")
	}
}