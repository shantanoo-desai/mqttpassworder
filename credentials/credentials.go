// mqttpassworder: Generate Encrypted Passwords for MQTT Mosquitto Broker
// Copyright (C) 2022  Shantanoo "Shan" Desai <sdes.softdev@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package credentials

import (
	"crypto/rand"
	"crypto/sha512"
	b64 "encoding/base64"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

const (
	SaltLength = 12  // 12 Random Bytes are required to generate Salt
	Iterations = 101 // Default Iterations when using PBDFK2_SHA512
)

// Credentials Structure
type Credentials struct {
	Username string
	Password string
}

// Generate Random Bytes of dedicated Length (here 12 Bytes)
// to generate the Salt for Encryption
func generateRandomBytes() (blk []byte) {
	// allocate SaltLength in Memory
	blk = make([]byte, SaltLength)
	// Read Bytes of SaltLength into Memory
	_, err := rand.Read(blk)
	if err != nil {
		fmt.Println("[Error]: generateRandomBytes: cannot generating random bytes for Salt")
		panic(err)
	}
	return
}

// Generate Salt (Base64E Encoded) as well as Base64 decoded
// to generate the desired Encryption
func (cred *Credentials) generateSalt() (string, string) {
	blk := generateRandomBytes()
	// Get the Base64 Encoded String from the 12 bytes array
	saltB64String := b64.StdEncoding.EncodeToString(blk)
	salt, err := b64.StdEncoding.DecodeString(saltB64String)
	if err != nil {
		fmt.Println("[Error]: generateSalt: Cannot perform Bas64 Encoding of generated Salt")
		panic(err)
	}
	return saltB64String, string(salt)
}

// Generate the Bas64 Encoded SHA512 Digest for the Password + Salt Combination
func (cred Credentials) GenerateSHA512() (saltB64, digestB64 string) {
	saltB64, salt := cred.generateSalt()
	combinedPassword := cred.Password + salt
	// generate Hash for the combined Password string
	h := sha512.New()
	h.Write([]byte(combinedPassword))
	chksum := h.Sum(nil)
	digestB64 = b64.StdEncoding.EncodeToString([]byte(chksum))
	return
}

// Generate PBKDF2 SHA512 Encryption for Generated Salt
func (cred Credentials) GeneratePBKDF2() (saltB64, digestB64 string) {
	saltB64, salt := cred.generateSalt()
	pbkdf2Key := pbkdf2.Key([]byte(cred.Password), []byte(salt), Iterations, 64, sha512.New)
	digestB64 = b64.StdEncoding.EncodeToString(pbkdf2Key)
	return
}
