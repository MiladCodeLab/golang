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


