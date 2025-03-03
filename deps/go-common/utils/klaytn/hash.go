package klaytn

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
	"time"

	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/crypto"
	"github.com/klaytn/klaytn/crypto/sha3"
)

func Sign(hash []byte, prv *ecdsa.PrivateKey) (r, s, v string, err error) {
	sig, err := crypto.Sign(hash, prv)
	if err != nil {
		return "", "", "", err
	}

	return DecodeSignature(sig)
}

func Recover(message string, hash string, r, s, v string, accountType string) (string, bool) {
	var hashStr []byte

	if hash != "" {
		msg := fmt.Sprintf("\x19Klaytn Signed Message:\n%v%s", len(message), message)
		hashStr1 := Keccak256([]byte(msg))
		var hashErr error
		hashStr, hashErr = hexutil.Decode(hash)

		if !bytes.Equal(hashStr1, hashStr) {
			logger.Errorf(context.TODO(), "hash and message not equal: ")
			return "", true
		}

		if hashErr != nil {
			logger.Errorf(context.TODO(), "hash error: ", hashErr)
			return "", true
		}
	} else {
		if accountType == "metamask" {
			msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%s", len(message), message)
			hashStr = Keccak256([]byte(msg))
		} else {
			msg := fmt.Sprintf("\x19Klaytn Signed Message:\n%v%s", len(message), message)
			hashStr = Keccak256([]byte(msg))
		}
	}

	rh, err1 := hexutil.Decode(r)
	sh, err2 := hexutil.Decode(s)

	vStr := v[2:]
	hexV, _ := strconv.ParseInt(vStr, 16, 16)
	hexV -= 27

	if hexV <= 9 {
		v = "0x0" + strconv.FormatInt(hexV, 10)
	} else {
		v = "0x" + strconv.FormatInt(hexV, 10)
	}

	vh, err3 := hexutil.Decode(v)

	if err1 != nil {
		logger.Errorf(context.TODO(), "r error: ", err1)
		return "", true
	}

	if err2 != nil {
		logger.Errorf(context.TODO(), "s error: ", err2)
		return "", true
	}

	if err3 != nil {
		logger.Errorf(context.TODO(), "v error: ", err3)
		return "", true
	}

	rsv := append(rh, sh...)
	rsv = append(rsv, vh...)

	public, publicErr := crypto.Ecrecover(hashStr, rsv)
	if publicErr != nil {
		logger.Errorf(context.TODO(), "public recover error: ", publicErr)
		return "", true
	}

	publicStr := hexutil.Encode(public)
	logger.Debugf(context.TODO(), "public key: ", publicStr)

	addr, addressErr := PublicKeyToAddress(publicStr)

	if addressErr != nil {
		logger.Errorf(context.TODO(), "public key error: ", addressErr)
		return "", true
	}
	logger.Debugf(context.TODO(), "address: ", addr)

	return addr, false
}

func DecodeSignature(sig []byte) (r, s, v string, err error) {
	if err != nil {
		return "", "", "", err
	}

	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}

	r = hexutil.Encode(sig[:32])
	s = hexutil.Encode(sig[32:64])
	v = hexutil.Encode([]byte{sig[len(sig)-1] + 27})

	return r, s, v, nil
}

func Keccak256(data []byte) []byte {
	h := sha3.NewKeccak256()
	h.Write(data)
	return h.Sum(nil)
}

func PublicKeyToAddress(publicKey string) (string, error) {
	pub, err := hexutil.Decode(Trim04FromPubkey(publicKey))
	if err != nil {
		return "", err
	}

	return common.BytesToAddress(Keccak256(pub)[12:]).String(), nil
}

func ToCheckSumAddress(addr string) string {
	unCheckSummed := strings.ToLower(addr[2:])
	sha := sha3.NewKeccak256()
	sha.Write([]byte(unCheckSummed))
	hash := sha.Sum(nil)

	result := []byte(unCheckSummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return "0x" + string(result)
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hexutil.Encode(bytes), nil
}

func CurrentTime() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func GetAddressFromPublicKey(pubkey string) string {
	h := &hexutil.Bytes{}
	h.UnmarshalText([]byte(ZeroPaddedPubkey(pubkey)))
	pub := []byte(*h)
	pub = append([]byte{0x04}, pub...)
	addr := common.BytesToAddress(Keccak256(pub[1:])[12:])

	return addr.String()
}

func GetXYFromPublicKey(pubkey string) (string, string) {
	if !strings.HasPrefix(pubkey, "0x") {
		return "", ""
	}
	pubkey = corePubkey(pubkey)

	x := pubkey[:64]
	y := pubkey[64:]

	trail := func(z string) string {
		for {
			if z[0] != '0' {
				break
			}
			z = z[1:]
		}

		return z
	}

	x = trail(x)
	y = trail(y)

	return "0x" + x, "0x" + y
}

func GetFirstFromArrayString(value string) string {
	return strings.Split(value, ",")[0]
}

func ZeroPaddedPubkey(pubkey string) string {
	pubkey = corePubkey(pubkey)
	return "0x" + LeftPad2Len(pubkey, "0", 128)
}

func Trim04FromPubkey(pubkey string) string {
	return "0x" + corePubkey(pubkey)
}

func IsSamePubkey(pubkey1 string, pubkey2 string) bool {
	return ZeroPaddedPubkey(pubkey1) == ZeroPaddedPubkey(pubkey2)
}

func corePubkey(pubkey string) string {
	if strings.HasPrefix(pubkey, "0x") {
		pubkey = pubkey[2:]
	}
	if len(pubkey) > 128 && strings.HasPrefix(pubkey, "04") {
		pubkey = pubkey[2:]
	}
	return pubkey
}

// LeftPad2Len reference https://github.com/git-time-metric/gtm/blob/master/util/string.go#L53-L88
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	retStr := strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
