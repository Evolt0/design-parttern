package adapter

import (
	"testing"
)

func TestAliyunAdaptor_Create(t *testing.T) {
	adaptor := &AliyunAdapter{}
	adaptor.Create("4Core", "8G")
	t.Log(adaptor)
}

func TestAwsAdapter_Create(t *testing.T) {
	adaptor := &AwsAdapter{}
	adaptor.Create("2Core", "4G")
	t.Log(adaptor)
}
