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
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	creds "github.com/shantanoo-desai/mqttpassworder/credentials"
)

var (
	sha512Ptr bool
	credsPtr  string
)

func init() {
	flag.BoolVar(&sha512Ptr, "sha512", false, "SHA512 Ecryption. Default is PBKDF2-SHA512")
	flag.StringVar(&credsPtr, "creds", "", "Credentials in <username>:<password> Format")
}

func main() {
	flag.Parse()

	if credsPtr == "" {
		fmt.Println("Credentials cannot be empty. Exiting...")
		os.Exit(-1)
	}
	credentials := &creds.Credentials{
		Username: strings.Split(credsPtr, ":")[0],
		Password: strings.Split(credsPtr, ":")[1],
	}
	if sha512Ptr {
		salt, encryptedPass := credentials.GenerateSHA512()
		fmt.Printf("%s:$6$%s$%s", credentials.Username, salt, encryptedPass)
		fmt.Println()
		os.Exit(0)
	}
	salt, encryptedPass := credentials.GeneratePBKDF2()
	fmt.Printf("%s:$7$%d$%s$%s\n", credentials.Username, creds.Iterations, salt, encryptedPass)
	fmt.Println()
}
