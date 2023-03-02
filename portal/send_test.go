package portal

import (
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

func (t *testHandler) Handle(msg any) {
	str, ok := msg.(string)
	if !ok {
		return
	}
	t.storage.Add(str)
	time.Sleep(200 * time.Millisecond)
}

func TestPortal_SendAndCloseSimultaneously(t *testing.T) {
	t.Parallel()

	count := 25

	portal := New()

	storage := &testStorage{}
	handler := &testHandler{storage: storage}
	portal.Subscribe(handler)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		portal.Close()
	}()

	for i := 0; i < count; i++ {
		msg := fmt.Sprintf("msg %d", i)
		portal.Send(msg)
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
