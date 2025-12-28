
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
