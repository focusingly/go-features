#include <stdatomic.h>
#include <stdio.h>

static inline __uint64_t next_id() {
  static atomic_uint_fast64_t counter = ATOMIC_VAR_INIT(0);

  return atomic_fetch_add(&counter, 1);
}

int worker(void *arg) {
  for (size_t i = 0; i < 10; i++) {
    __uint64_t id = next_id();
    printf("Thread %d got ID: %llu\n", *(int *)arg, id);
  }

  return 0;
}

int main() {
  __uint64_t threads[10];
  int ids[10];

  for (int i = 0; i < 10; ++i) {
    ids[i] = i;
    thrd_create(&threads[i], worker, &ids[i]);
  }

  for (int i = 0; i < 10; ++i) {
    thrd_join(threads[i], NULL);
  }

  return 0;
}