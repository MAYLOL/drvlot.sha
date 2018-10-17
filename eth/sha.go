// h 20181016
//
// DrvLot.SHA

package ean

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func Sha(msg []byte) string {
	if msg == nil {
		msg = []byte("")
	}
	sha := crypto.Keccak256(msg)
	return hex.EncodeToString(sha)
}
