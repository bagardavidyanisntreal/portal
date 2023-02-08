package portal

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
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
	time.Sleep(200 * time.Millisecond)
}

type testMsg string

func (t testMsg) Data() any {
	return string(t)
}

func TestPortal_SendAndCloseSimultaneously(t *testing.T) {
	t.Parallel()

	count := 25

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

	for i := 0; i < count; i++ {
		msg := fmt.Sprintf("msg %d", i)
		portal.Send(testMsg(msg))
	}

	wg.Wait()

	rand.Seed(time.Now().UnixNano())
	randId := rand.Intn(count)

	wantMsg := fmt.Sprintf("msg %d", randId)

	var fails int
	for _, msg := range storage.Data() {
		if msg != wantMsg {
			fails++
			continue
		}
		t.Logf("you are so lucky, wantMsg '%s' found!", wantMsg)
	}

	if fails == count {
		t.Errorf("got fails eq to count, wantMsg '%s' not found", wantMsg)
	}
}
