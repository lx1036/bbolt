package bbolt

import (
	"k8s.io/klog/v2"
	"testing"
)

func TestName3(test *testing.T) {
	var buf [0x1000]byte
	klog.Info(len(buf))

}
