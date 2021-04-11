package aladin

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Snapshot 配置快照
// 当调用 Watcher.Next 后，可能更新当前快照
type Snapshot struct {
	Data      []byte
	Path      string
	Source    string
	Checksum  string
	Timestamp time.Time
}

func (s *Snapshot) Sum() string {
	h := sha256.New()
	h.Write(s.Data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
