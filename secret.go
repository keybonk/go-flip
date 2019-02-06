package flip

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func (s *Secret) XOR(t Secret) *Secret {
	for i, b := range t {
		s[i] = b ^ s[i]
	}
	return s
}

func (s Secret) IsNil() bool {
	for _, b := range s[:] {
		if b != byte(0) {
			return false
		}
	}
	return true
}

func (s Secret) Hash() Secret {
	h := sha256.New()
	h.Write(s[:])
	tmp := h.Sum(nil)
	var ret Secret
	copy(ret[:], tmp[:])
	return ret
}

func (s Secret) Eq(t Secret) bool {
	return hmac.Equal(s[:], t[:])
}

func (s Secret) computeCommitment(cp CommitmentPayload) (Commitment, error) {
	var ret Commitment
	hm := hmac.New(sha256.New, s[:])
	raw, err := msgpackEncode(cp)
	if err != nil {
		return ret, err
	}
	fmt.Printf("%#x\n", raw)
	hm.Write(raw)
	tmp := hm.Sum(nil)
	copy(ret[:], tmp)
	return ret, nil
}

func (c Commitment) Eq(d Commitment) bool {
	return hmac.Equal(c[:], d[:])
}

func GenerateSecret() Secret {
	var ret Secret
	n, err := rand.Read(ret[:])
	if n != len(ret) {
		panic("failed to generate secret; short random read")
	}
	if err != nil {
		panic(fmt.Sprintf("failed to generated secret: %s", err.Error()))
	}
	return ret
}
