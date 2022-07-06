package types

import (
	"github.com/babylonchain/babylon/crypto/bls12381"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "checkpointing"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkpointing"
)

var (
	BlsSigsPrefix      = []byte{0x0} // reserve this namespace for BLS sigs
	CheckpointsPrefix  = []byte{0x1} // reserve this namespace for checkpoints
	RegistrationPrefix = []byte{0x2} // reserve this namespace for BLS keys

	BlsSigsObjectPrefix      = append(BlsSigsPrefix, 0x0) // where we save the concrete BLS sig bytes
	BlsSigsHashToEpochPrefix = append(BlsSigsPrefix, 0x1) // where we map hash to epoch

	CkptsObjectPrefix = append(CheckpointsPrefix, 0x0) // where we save the concrete BLS sig bytes

	AddrToBlsKeyPrefix = append(RegistrationPrefix, 0x0) // where we save the concrete BLS public keys
	BlsKeyToAddrPrefix = append(RegistrationPrefix, 0x1) // where we save BLS key set
)

// BlsSigsObjectKey defines epoch + hash
func BlsSigsObjectKey(epoch uint64, hash BlsSigHash) []byte {
	ee := sdk.Uint64ToBigEndian(epoch)
	epochPrefix := append(BlsSigsObjectPrefix, ee...)
	return append(epochPrefix, hash...)
}

// BlsSigsEpochKey defines BLS sig hash
func BlsSigsEpochKey(hash BlsSigHash) []byte {
	return append(BlsSigsHashToEpochPrefix, hash...)
}

// CkptsObjectKey defines epoch
func CkptsObjectKey(epoch uint64) []byte {
	return append(CkptsObjectPrefix, sdk.Uint64ToBigEndian(epoch)...)
}

// AddrToBlsKeyKey defines validator address
func AddrToBlsKeyKey(valAddr ValidatorAddress) []byte {
	return append(AddrToBlsKeyPrefix, []byte(valAddr)...)
}

// BlsKeyToAddrKey defines BLS public key
func BlsKeyToAddrKey(pk bls12381.PublicKey) []byte {
	return append(BlsKeyToAddrPrefix, pk...)
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}