package klaytn

import "github.com/klaytn/klaytn/crypto"

func Hash(data []byte) [32]byte {
	return crypto.Keccak256Hash(data)
}
