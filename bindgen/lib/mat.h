#include <immintrin.h>
#include <stdio.h>
#include <stdlib.h>

inline static void mat_add(const float *a, const float *b, float *result,
                           const int length);
inline static void mat_sub(const float *a, const float *b, float *result,
                           const int length);
