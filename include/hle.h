#ifndef _HLE_LOCKER
#define _HLE_LOCKER

void acquire_hle(unsigned int* lock);
void release_hle(unsigned int* lock);

#endif
