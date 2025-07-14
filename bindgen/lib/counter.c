#include "counter.h"

static inline __uint64_t next_id() {
    static atomic_uint_fast64_t counter = ATOMIC_VAR_INIT(0);

    return atomic_fetch_add(&counter, 1);
}