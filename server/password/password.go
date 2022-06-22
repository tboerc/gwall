package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/tboerc/gwall/server/messages"
	"golang.org/x/crypto/argon2"
)

const (
	MEMORY      = 64 * 1024
	ITERATIONS  = 3
	PARALLELISM = 2
	SALT_LENGTH = 16
	KEY_LENGTH  = 32
)

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func decodeHash(encodedHash []byte) (salt, hash []byte, err error) {
	vals := strings.Split(string(encodedHash[:]), "$")
	if len(vals) != 6 {
		return nil, nil, messages.ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return
	}
	if version != argon2.Version {
		return nil, nil, messages.ErrIncompatibleVersion
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return
	}

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return
	}

	return salt, hash, nil
}

func Hash(p []byte) (encodedHash []byte, err error) {
	salt, err := generateRandomBytes(SALT_LENGTH)
	if err != nil {
		return nil, err
	}

	hash := argon2.IDKey(p, salt, ITERATIONS, MEMORY, PARALLELISM, KEY_LENGTH)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = []byte(fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, MEMORY, ITERATIONS, PARALLELISM, b64Salt, b64Hash))

	return encodedHash, nil
}

func Compare(password, encodedHash []byte) (match bool, err error) {
	salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, ITERATIONS, MEMORY, PARALLELISM, KEY_LENGTH)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}
