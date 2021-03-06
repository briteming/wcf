package lb

import (
	"testing"
	"time"
)

func TestLoad(t *testing.T) {
	lb := New(3, 30 * time.Second)
	lb.Add("127.0.0.1", 10)
	lb.Add("127.0.0.2", 50)
	lb.Add("127.0.0.3", 30)
	lb.Add("127.0.0.4", 10)
	cnt := make(map[string]int)
	total := 200
	now := time.Now()
	for i := 0; i < total; i++ {
		addr, err := lb.Get()
		if err != nil {
			t.Log(err)
			time.Sleep(1 * time.Second)
			continue
		}
		t.Logf("addr:%s, err:%v, curr:%d, base:%d, lastfail:%v, maxerr:%d", addr, err, lb.mp[addr].Current, lb.mp[addr].Base, lb.mp[addr].LastFail, lb.mp[addr].Errtime)
		cnt[addr]++
		if now.Add(30 * time.Second).Before(time.Now()) {
			lb.Update(addr, true)
		} else {
			lb.Update(addr, false)
		}
	}
	for k, v := range cnt {
		t.Logf("addr:%d, cnt:%d, per:%f, info:%+v", k, v, float64(v) / float64(total), lb.mp[k])
	}
}

