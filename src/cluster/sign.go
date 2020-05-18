package cluster

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/name5566/leaf/log"
)

const (
	key = "DAlrQaI44k9ScS_MHAPeEPDX1h_S06we"
)

func checkSignature(msg []byte) []byte {
	pkg := Package{}
	if err := json.Unmarshal(msg, &pkg); err != nil {
		log.Error("umarshal msg fail %v", err)
		return nil
	}
	sign := pkg.Sign
	data := pkg.Data
	// log.Debug("signData:%v", data)
	// log.Debug("sign:%v", signature(data))
	if signature(data) != sign {
		return nil
	}
	return []byte(data)
}

func signature(data string) string {
	h := sha1.New()
	h.Write([]byte(key + data))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

type Package struct {
	Sign string
	Data string
}
