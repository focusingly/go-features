#include "mat.h"

inline static void mat_add(const float *a, const float *b, float *result,
                           const int length) {
  int i = 0;
  for (; i + 7 < length; i += 8) {
    __m256 va = _mm256_loadu_ps(&a[i]);
    __m256 vb = _mm256_loadu_ps(&b[i]);
    __m256 vr = _mm256_add_ps(va, vb);
    _mm256_storeu_ps(&result[i], vr);
  }
  // 处理剩余元素
  for (; i < length; i++) {
    result[i] = a[i] + b[i];
  }
}

inline static void mat_sub(const float *a, const float *b, float *result,
                           const int length) {
  int i = 0;
  for (; i + 7 < length; i += 8) {
    __m256 va = _mm256_loadu_ps(&a[i]);
    __m256 vb = _mm256_loadu_ps(&b[i]);
    __m256 vr = _mm256_sub_ps(va, vb);
    _mm256_storeu_ps(&result[i], vr);
  }
  for (; i < length; i++) {
    result[i] = a[i] - b[i];
  }
}
