# Go Concurrency Tool Selection Cheatsheet

Welcome to the ultimate guide for choosing the right concurrency primitive in Go. Writing concurrent programs in Go is easy, but choosing the *wrong* tool can introduce deadlocks, data races, or performance bottlenecks. 

Use this repository/document as a quick-reference guide to decide exactly when to use **Goroutines**, **Channels**, **Mutexes**, and **Atomics**.

---

## 🧠 Quick Reference Decision Tree

1. **Do you need to run an independent task in the background?**
   * 👉 Use a **Goroutine**.
2. **Are you passing data or a stream of events from one goroutine to another?**
   * 👉 Use a **Channel**.
3. **Are you updating a single shared number or a boolean flag?**
   * 👉 Use **Atomic**.
4. **Are you protecting a shared Map, Slice, or Struct from concurrent access?**
   * 👉 Use a **Mutex**.

---

## 🚀 1. Goroutines (The Engine)
Goroutines are lightweight, independent execution threads managed entirely by the Go runtime. They form the bedrock of Go's concurrency model.

### 🟢 Use When:
* **Handling I/O-bound tasks:** Waiting for external APIs, database queries, file system operations, or network sockets.
* **Background Tasks (Fire-and-Forget):** Offloading non-blocking processing like sending an email, writing logs, or handling file uploads.
* **High-Throughput Servers:** Spawning routines to manage individual incoming HTTP requests or WebSocket connections independently.
* **CPU-Heavy Parallelism:** Splitting massive computational algorithms (e.g., image processing, data chunking) across multiple CPU cores.

### 🔴 Do NOT Use When:
* **Strict sequential logic is required:** If Task B cannot start until Task A finishes, concurrency adds pointless synchronization overhead.
* **The task is trivial:** Spawning a goroutine has a tiny overhead, but performing basic operations (like `a + b`) inside one actually slows your program down.
* **Uncontrolled Lifecycles:** Spawning infinite goroutines without bounds (like a worker pool or a semaphore) can leak memory and crash your system.

---

## ✉️ 2. Channels (The Communicator)
Channels are natively thread-safe pipes that connect concurrent goroutines, allowing them to pass data and synchronize status without explicit locking.

### 🟢 Use When:
* **Passing Ownership:** Safely passing data from Worker A to Worker B so that Worker B becomes the sole owner responsible for modifying it.
* **Work Pipelines:** Setting up assembly lines where one goroutine fetches data, another processes it, and a third stores it.
* **Worker Pools:** Distributing a single queue of tasks across a fixed number of idling worker goroutines.
* **Event Signaling:** Utilizing an empty struct channel (`chan struct{}`) to broadcast cancellation or execution signals across multiple threads.

### 🔴 Do NOT Use When:
* **Managing internal struct state:** If you just need to update an internal property safely, channels are clumsy and over-engineered.
* **Microsecond-critical performance:** Channels have internal synchronization logic. For raw performance over simple data shares, Mutex or Atomic is faster.
* **Complex Multi-entity States:** Coordinating interlocking states over multiple channels heavily exposes your codebase to deadlocks.

---

## 🔒 3. Mutex / RWMutex (The Guard)
A Mutex (`sync.Mutex`) provides mutual exclusion, ensuring only one goroutine can touch a specific block of memory at a time. A Read-Write Mutex (`sync.RWMutex`) permits multiple readers or a single writer.

### 🟢 Use When:
* **Protecting Shared State:** Guarding non-thread-safe data types like **Slices, Maps, or Structs** from concurrent access.
* **Thread-safe Caches:** Managing in-memory map caches where multiple threads look up values or apply mutations concurrently.
* **Read-Heavy Operations (`RWMutex`):** When thousands of goroutines need to safely read a configuration variable that only changes once in a while.
* **Check-Then-Act Transactions:** Verifying a state condition and executing a resulting action as one unbroken, locked block (e.g., *"If account balance > \$10, subtract \$10"*).

### 🔴 Do NOT Use When:
* **Task Orchestration:** If you are passing results back and forth between functions, use channels. Mutexes protect raw memory; they do not route data.
* **Long-Running I/O:** Never hold a lock while waiting for a slow disk read or network API. This will freeze up your entire system.
* **Simple Counter Increments:** For updating isolated numerical variables, a Mutex introduces too much software overhead. Use Atomics.

---

## ⚡ 4. Atomics (The Precision Strike)
Atomic operations (`sync/atomic`) use low-level CPU hardware instructions to perform lock-free operations directly on single, basic numeric variables.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var requestCounter int64
	var wg sync.WaitGroup

	// Spawning 1000 concurrent routines making an increment
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&requestCounter, 1) // Pure hardware execution loop
		}()
	}

	wg.Wait()
	fmt.Println("Total requests handled safely:", atomic.LoadInt64(&requestCounter))
}
```

### 🟢 Use When:
* **High-Frequency Metrics:** Keeping real-time score of system analytics, such as tracking total requests, active connections, or bytes processed per second.
* **Global State Flags:** Managing binary toggles (e.g., `isShuttingDown` or `isServerHealthy`) that thousands of routines check constantly.
* **Throttling/Rate Limiters:** Tracking integer window allocations across highly concurrent networking threads.

### 🔴 Do NOT Use When:
* **Modifying Compound Types:** You cannot execute atomic actions on maps, slices, strings, or multi-field structs.
* **Multi-Variable Operations:** You cannot group multiple atomic updates into one transaction. If you modify $X$ atomically and $Y$ atomically, another thread can still read the system mid-way and see broken/mismatched state data.
* **Readability Over Performance:** Atomic code can be highly unintuitive to debug. If your logic isn't running millions of times per second, choose a `sync.Mutex` for cleaner code maintenance.

---
