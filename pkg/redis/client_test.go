package redis

import (
	"testing"
)

func TestAddAndGet(t *testing.T) {
	cli := NewClient("10.9.44.194:7001,10.9.44.194:7002,10.9.44.194:7000", "", 5, 5, 5, 8)

	if err := cli.Open(); err != nil {
		t.Fatalf("open %v", err)
	}

	key := "d:name-flow.com"
	value := 622

	if err := cli.Set(key, value, 0); err != nil {
		t.Errorf("set error, %v", err)
	}

	if data, err := cli.Get(key + "1223"); err != nil {
		t.Errorf("get error, %v", err)
	} else {
		t.Logf("get value:%s", data)
	}

	if err := cli.Close(); err != nil {
		t.Errorf("close error, %v", err)
	}
}
