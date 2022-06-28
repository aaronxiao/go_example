package tsync

import (
	"fmt"
	"sync"
	"testing"
)

func TestLoadOrStore(t *testing.T)  {
	var m sync.Map
	key := fmt.Sprintf("111")
	val := true

	m.LoadOrStore(key, val)
	fmt.Println( m.LoadAndDelete(key) )

	fmt.Println( m.Load(key) )



}
