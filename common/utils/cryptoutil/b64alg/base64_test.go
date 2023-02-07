package b64alg

import (
	"testing"
)

func TestEncode(t *testing.T) {
	data := "abc123!?$*&()'-=@~"
	sEnc := StdEncoding.EncodeToString([]byte(data))
	t.Log(sEnc)
	dst, _ := StdEncoding.DecodeString(sEnc)
	t.Log(string(dst))

}
