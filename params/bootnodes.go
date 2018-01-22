// Copyright 2015 The go-icube Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://0c115f7088c197ef402f359489a9b52242611dfa8d0fd9799c3f8a4413ac2bc93413a2b18bfc8c98b4a11dce987e00248d57573af65604ec4afe1eb38bc3d064@47.89.193.53:30821",  //US_West
	"enode://ce1242239d4c1c71e72efbbff0115dd641dfd5e6a7378d8c7dcb13f404b0d9294ea63c8782ff80cf1e49a8e9f7b81ac47b0dd6e49e41587420962a69883de4bd@47.90.205.122:30821", //US_East
	"enode://7c17bf64cd58cb823d600aeb91c17c1ebf225e4bae73100ac8790bff23dfdc889d4cc1201ecce4abbfae4ad5b9fb4d14ff9343a9124de3fd127a80a6a2dcf721@47.74.5.209:30821",   //Japan
	"enode://78402a6c73a7c81334d306bdea83519ab94d74fa824ee19f69e3f6e9e811e9f7a1cd9b6a3e2ee05684fc7964a0013203e8eff1aa221cb77ede146c385151e036@47.254.131.15:30821", //Germany
	"enode://0fe1d51d113be35a88cc8f06fe172eae3b491fcd8f564840b01cb327d77e7b2cadd380e3afd4953ad32f742196fda4ce58cca6f7a410ea6430bc4c8f4a4f4def@47.88.223.157:30821", //Singapore
	"enode://639503f6c45118efa9e74b80c116382d692ab62f0d8fbf33dd1f5dc868c86455ac34abdc1ce35c6ade8189909c7519dc41baa0666b80c1e64b860ef201618f67@47.91.47.212:30821",  //Sydney
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://b80cd4e508249022cd15add06f2dfb15f99cc3bbf6bff90770886ef38832964f6dc9b3642b8ce06079b10090e8b90b2b713b895445dc465973bbac2f70b5b88b@47.74.10.13:30821",   //Japan
	"enode://caa23b8dfb2a4845281077dd390fc5b57cd2181ad24530cd4e098311bfc67c4a437b2acf0383e328c44f6bec3066c3bf84642198af51e10fa299c71400a1a2b0@47.74.239.171:30821", //Singapore
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
//var RinkebyBootnodes = []string{}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
//var RinkebyV5Bootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
