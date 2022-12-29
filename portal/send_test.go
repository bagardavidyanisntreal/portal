package portal

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

type testStorage struct {
	lock sync.RWMutex
	data []string
}

func (s *testStorage) Add(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	fmt.Println("[test storage receiver]: stored data", str)
	s.data = append(s.data, str)
}

func (s *testStorage) Data() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.data
}

type testHandler struct {
	storage *testStorage
}

func (t *testHandler) Support(_ Message) bool {
	return true
}

func (t *testHandler) Handle(msg Message) {
	data := msg.Data()
	str, ok := data.(string)
	if !ok {
		return
	}
	t.storage.Add(str)
}

type testMsg string

func (t testMsg) Data() any {
	return string(t)
}

func TestPortal_SendAndCloseSimultaneously(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	portal := New(ctx)

	storage := &testStorage{}
	handler := &testHandler{storage: storage}
	portal.Await(ctx, handler)

	var wg sync.WaitGroup
	var mx sync.Mutex
	wg.Add(1)
	go func() {
		mx.Lock()
		defer mx.Unlock()
		defer wg.Done()
		portal.Close()
	}()

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("msg %d", i)
		portal.Send(testMsg(msg))
	}
	portal.Send(testMsg("msg 666"))

	wg.Wait()

	for _, msg := range storage.Data() {
		if msg == "msg 666" {
			t.Log("you are so lucky to see this here")
		}
	}
}
