// Copyright 2016 The go-icube Authors
// This file is part of the go-icube library.
//
// The go-icube library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-icube library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-icube library. If not, see <http://www.gnu.org/licenses/>.

// +build none

// This program generates contract/code.go, which contains the chequebook code
// after deployment.
package main

import (
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/icube/go-icube/accounts/abi/bind"
	"github.com/icube/go-icube/accounts/abi/bind/backends"
	"github.com/icube/go-icube/contracts/chequebook/contract"
	"github.com/icube/go-icube/core"
	"github.com/icube/go-icube/crypto"
)

var (
	testKey, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAccount = core.GenesisAccount{
		Address: crypto.PubkeyToAddress(testKey.PublicKey),
		Balance: big.NewInt(500000000000),
	}
)

func main() {
	backend := backends.NewSimulatedBackend(testAccount)
	auth := bind.NewKeyedTransactor(testKey)

	// Deploy the contract, get the code.
	addr, _, _, err := contract.DeployChequebook(auth, backend)
	if err != nil {
		panic(err)
	}
	backend.Commit()
	code, err := backend.CodeAt(nil, addr, nil)
	if err != nil {
		panic(err)
	}
	if len(code) == 0 {
		panic("empty code")
	}

	// Write the output file.
	content := fmt.Sprintf(`package contract

// ContractDeployedCode is used to detect suicides. This constant needs to be
// updated when the contract code is changed.
const ContractDeployedCode = "%#x"
`, code)
	if err := ioutil.WriteFile("contract/code.go", []byte(content), 0644); err != nil {
		panic(err)
	}
}
