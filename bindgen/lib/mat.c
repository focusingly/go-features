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

inline static void mat_mul(const float *a, const float *b, float *result,
                           const int length) {
  for (int i = 0; i < length; i++) {
    for (int j = 0; j < length; j++) {
      __m256 sum = _mm256_setzero_ps();
      int k = 0;
      for (; k + 7 < length; k += 8) {
        __m256 va = _mm256_loadu_ps(&a[i * length + k]);
        __m256 vb = _mm256_loadu_ps(&b[j + k * length]);
        __m256 prod = _mm256_mul_ps(va, vb);
        sum = _mm256_add_ps(sum, prod);
      }
      // 水平求和 AVX2 加速部分
      float temp[8];
      _mm256_storeu_ps(temp, sum);
      float s = 0;
      for (int t = 0; t < 8; t++) {
        s += temp[t];
      }
      // 处理尾部标量部分
      for (; k < length; k++) {
        s += a[i * length + k] * b[j + k * length];
      }
      result[i * length + j] = s;
    }
  }
}
