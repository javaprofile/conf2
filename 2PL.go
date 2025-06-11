Package main
import (
	"fmt"
	"sync"
	"time"
)
type Transaction struct {
	ID        int
	Locks     map[string]*sync.Mutex
	IsBlocked bool
}
type ResourceManager struct {
	Resources map[string]*sync.Mutex
}
func (rm *ResourceManager) LockResource(tx *Transaction, resource string) {
	if tx.IsBlocked {
		return
	}
	tx.Locks[resource].Lock()
}
func (rm *ResourceManager) UnlockResource(tx *Transaction, resource string) {
	tx.Locks[resource].Unlock()
}
func (rm *ResourceManager) BeginTransaction(tx *Transaction) {
	tx.Locks = make(map[string]*sync.Mutex)
}
func (rm *ResourceManager) CommitTransaction(tx *Transaction) {
	for resource, lock := range tx.Locks {
		rm.UnlockResource(tx, resource)
		fmt.Printf("Transaction %d committed, released lock on resource %s\n", tx.ID, resource)
	}
}
func (rm *ResourceManager) AbortTransaction(tx *Transaction) {
	for resource, lock := range tx.Locks {
		rm.UnlockResource(tx, resource)
		fmt.Printf("Transaction %d aborted, released lock on resource %s\n", tx.ID, resource)
	}
}
func main() {
	rm := &ResourceManager{
		Resources: map[string]*sync.Mutex{
			"resource1": &sync.Mutex{},
			"resource2": &sync.Mutex{},
			"resource3": &sync.Mutex{},
		},
	}
	tx1 := &Transaction{ID: 1}
	tx2 := &Transaction{ID: 2}
	rm.BeginTransaction(tx1)
	rm.BeginTransaction(tx2)
	go func() {
		rm.LockResource(tx1, "resource1")
		time.Sleep(2 * time.Second)
		rm.LockResource(tx1, "resource2")
		rm.CommitTransaction(tx1)
	}()
	go func() {
		rm.LockResource(tx2, "resource2")
		time.Sleep(2 * time.Second)
		rm.LockResource(tx2, "resource3")
		rm.CommitTransaction(tx2)
	}()
time.Sleep(5 * time.Second)
}
