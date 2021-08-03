package mntp

import (
	"encoding/json"

	"github.com/segmentio/ksuid"
)

func UID() string {
	return ksuid.New().String()
}

func pack(pkt NtpPackage) []byte {
	b, _ := json.Marshal(pkt)
	return b
}

func unpack(b []byte) NtpPackage {
	var p NtpPackage
	json.Unmarshal(b, &p)
	return p
}
