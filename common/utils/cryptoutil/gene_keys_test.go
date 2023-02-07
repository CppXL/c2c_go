package cryptoutil

import (
	"encoding/hex"
	"testing"
)

func TestGenerateAesKey(t *testing.T) {
	key, err := GenerateAesKey()
	if err != nil {
		t.Error(err)
	}
	t.Log("\n" + hex.Dump(key))
}

func TestGenerateIv(t *testing.T) {
	iv, err := GenerateIv()
	if err != nil {
		t.Error(err)
	}
	t.Log("\n" + hex.Dump(iv))
}
