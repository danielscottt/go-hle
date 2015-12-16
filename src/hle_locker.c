#include <stdlib.h>
#include <immintrin.h>

#include <hle-emulation.h>

#include <hle.h>

void acquire_hle(unsigned int* lock)
{
        while (__hle_acquire_test_and_set4(lock) == 1) {
                while (*lock == 0)
                        _mm_pause();
        }
}

void release_hle(unsigned int* lock)
{
        __hle_release_clear4(lock);
}
