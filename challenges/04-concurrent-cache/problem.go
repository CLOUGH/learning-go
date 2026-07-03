// Package concurrentcache: a Get/Set string->int cache safe for
// concurrent use by many goroutines at once.
package concurrentcache

// Cache must be safe to use from many goroutines concurrently: any mix of
// concurrent Get, Set, Delete and Len calls must never race and must
// never corrupt the underlying map.
//
// Requirements:
//   - Get(key) returns (value, true) if present, (0, false) otherwise.
//   - Set(key, value) inserts or overwrites.
//   - Delete(key) removes a key if present; a no-op if it's already absent.
//   - Len() returns the current number of entries.
//   - All of the above must be safe under `go test -race` when called
//     concurrently from many goroutines.
//   - Reads (Get, Len) should be allowed to run concurrently with each
//     other - only a write (Set, Delete) needs exclusive access.
//
// Hint: this is exactly the sync.RWMutex use case from lesson 06.
//
// TODO: add whatever fields you need and implement the methods below.
type Cache struct {
	// TODO: fields
}

func NewCache() *Cache {
	return &Cache{}
}

// TODO: implement Get.
func (c *Cache) Get(key string) (int, bool) {
	panic("TODO: implement Get")
}

// TODO: implement Set.
func (c *Cache) Set(key string, value int) {
	panic("TODO: implement Set")
}

// TODO: implement Delete.
func (c *Cache) Delete(key string) {
	panic("TODO: implement Delete")
}

// TODO: implement Len.
func (c *Cache) Len() int {
	panic("TODO: implement Len")
}
