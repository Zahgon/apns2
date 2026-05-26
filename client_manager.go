package apns2

import (
	"container/list"
	"crypto/sha1"
	"crypto/tls"
	"sync"
	"time"
)

type managerItem struct {
	key      [sha1.Size]byte
	client   *Client
	lastUsed time.Time
}

// ClientManager is a way to manage multiple connections to the APNs.
type ClientManager struct {
	// MaxSize is the maximum number of clients allowed in the manager. When
	// this limit is reached, the least recently used client is evicted. Set
	// zero for no limit.
	MaxSize int

	// MaxAge is the maximum age of clients in the manager. Upon retrieval, if
	// a client has remained unused in the manager for this duration or longer,
	// it is evicted and nil is returned. Set zero to disable this
	// functionality.
	MaxAge time.Duration

	// Factory is the function which constructs clients if not found in the
	// manager.
	Factory func(certificate tls.Certificate) *Client

	cache map[[sha1.Size]byte]*list.Element
	ll    *list.List
	mu    sync.Mutex
	once  sync.Once
}

// NewClientManager returns a new ClientManager for prolonged, concurrent usage
// of multiple APNs clients. ClientManager is flexible enough to work best for
// your use case. When a client is not found in the manager, Get will return
// the result of calling Factory, which can be a Client or nil.
//
// Having multiple clients per certificate in the manager is not allowed.
//
// By default, MaxSize is 64, MaxAge is 10 minutes, and Factory always returns
// a Client with default options.
func NewClientManager() *ClientManager { _ = "STUB: not implemented"; return nil }

// Add adds a Client to the manager. You can use this to individually configure
// Clients in the manager.
func (m *ClientManager) Add(client *Client) { _ = "STUB: not implemented"; return }

// Get gets a Client from the manager. If a Client is not found in the manager
// or if a Client has remained in the manager longer than MaxAge, Get will call
// the ClientManager's Factory function, store the result in the manager if
// non-nil, and return it.
func (m *ClientManager) Get(certificate tls.Certificate) *Client {
	_ = "STUB: not implemented"
	return nil
}

// Len returns the current size of the ClientManager.
func (m *ClientManager) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *ClientManager) initInternals() { _ = "STUB: not implemented"; return }

func (m *ClientManager) removeOldest() { _ = "STUB: not implemented"; return }

func (m *ClientManager) removeElement(e *list.Element) { _ = "STUB: not implemented"; return }

func cacheKey(certificate tls.Certificate) [sha1.Size]byte { _ = "STUB: not implemented"; return nil }
