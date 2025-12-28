# Go Struct Passing Strategy: Pointer vs Value

## Context

We evaluated two approaches for passing a **large domain struct (~200 bytes)** across multiple application layers (controller → service → repository → database):

- **Value passing** (`Conference`)
- **Pointer passing** (`*Conference`)

The struct includes multiple `time.Time`, `uuid.UUID`, strings, and pointer fields, making it a **non-trivial, heap-influencing object**.

---

## Struct Characteristics

- Approximate size: **~200 bytes**
- Contains:
  - Heap-backed fields (`string`, pointers)
  - Multiple `time.Time` values
  - UUIDs
- Passed through **4 layers**
- High iteration count, GC-sensitive workload

---

## Test Setup

- Runtime duration: **1 minute**
- Identical logic and workload
- Metrics captured via `runtime.MemStats`
- Go runtime defaults (GC enabled)

---

## Results

### Pointer Passing (`*Conference`)

```

Iterations: 224,000,000
TotalAlloc: 14,526 MB
HeapAlloc:  3 MB
StackInuse: 192 KB
Mallocs:   1,120,000,403
Frees:     1,119,781,350
GC Count:  4,014

```

### Value Passing (`Conference`)

```

Iterations: 220,000,000
TotalAlloc: 14,267 MB
HeapAlloc:  1 MB
StackInuse: 192 KB
Mallocs:   1,100,000,449
Frees:     1,099,898,049
GC Count:  3,945

```

---

## Comparative Analysis

| Metric | Pointer Passing | Value Passing | Insight |
|-----|-----------------|---------------|--------|
| Throughput | **Higher** | Lower | Fewer struct copies |
| Struct Copies | **0** | 4 per request | Scales poorly with size |
| TotalAlloc | Slightly higher | Slightly lower | Cumulative metric only |
| Live Heap | 3 MB | 1 MB | Both stable |
| Stack Usage | Identical | Identical | No stack pressure difference |
| GC Count | Slightly higher | Slightly lower | GC cycles are cheap due to low live heap |

---

## Interpretation

- **Throughput is higher with pointer passing**, despite marginally more allocations.
- **Live heap remains flat** in both cases, indicating healthy GC behavior.
- GC frequency alone is **not a performance KPI**.
- The cost of copying a ~200-byte struct across multiple layers outweighs the cost of short-lived heap allocations.

---

## Architectural Implications

Pointer passing:

- Eliminates copy amplification
- Scales better as the struct evolves
- Reduces CPU overhead from memory copying
- Produces predictable memory behavior
- Aligns with Go best practices for large aggregates

Value passing:

- Slightly fewer allocations
- Lower GC count
- But incurs repeated large memory copies
- Becomes increasingly expensive as complexity grows

---

## Final Verdict

**Passing the struct by pointer (`*Conference`) is the correct and more efficient approach.**

For a ~200-byte struct passed across multiple layers:
- Pointer passing delivers higher throughput
- Maintains stable heap usage
- Avoids unnecessary memory copying
- Fits production-grade Go backend architecture

**Recommendation:**  
Standardize on pointer passing for large domain models and aggregates.
