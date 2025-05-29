# LRU Cache (Least Recently Used)

## Problem Statement

Design and implement an **in-memory LRU (Least Recently Used) cache** that stores key-value pairs with an upper memory limit. When the cache exceeds this limit, it should evict the least recently accessed item.

---

## Requirements

### Core Features

#### Put
- Stores a key-value pair in the cache.
- If the key already exists, update its value and mark it as recently used.
- If adding a new key exceeds the cache capacity, evict the least recently used item.

#### Get
- Retrieve the value associated with a key.
- Mark the key as recently used.
- Return an error or sentinel value if the key is not found.

---

### Configuration

- Set a maximum **capacity** (number of entries).
- Capacity must be respected at all times.

---

### Internal Mechanism

- Must track usage order to quickly identify and evict the least recently used item.
- Recommended structures:
  - **Hash Map** for O(1) access to values.
  - **Doubly Linked List** to maintain usage order.

---

## Example

### Operations

```go
lru := NewLRUCache(2)       // Capacity = 2

lru.Put("a", 1)             // Cache = {a=1}
lru.Put("b", 2)             // Cache = {a=1, b=2}
lru.Get("a")                // Returns 1, usage order: {b=2, a=1}
lru.Put("c", 3)             // Evicts "b", Cache = {a=1, c=3}
lru.Get("b")                // Returns nil or -1 (not found)
```

| Feature       | Description                              |
| ------------- | ---------------------------------------- |
| Put           | Add or update a key-value pair           |
| Get           | Retrieve and mark as most recently used  |
| Capacity      | Fixed size with automatic eviction       |
| O(1) Time Ops | Get/Put operations must be constant time |
| Eviction      | Least recently accessed item is removed  |
