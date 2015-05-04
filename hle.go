package hle

/*
#include <stdlib.h>
#include <stdbool.h>
#include <assert.h>
#include <immintrin.h>

enum lock_states {unlocked, locked};

struct locker {
	enum lock_states lock;
};

struct locker *locker_init()
{
	struct locker *l = malloc(sizeof(struct locker));
	assert(l != NULL);

	l->lock = (enum lock_states) malloc(sizeof(unlocked));
	l->lock = unlocked;

	return l;
}

void acquire(struct locker *l)
{
	int ll, ul;
	ll = locked;
	ul = unlocked;
	int *ulp = (int*)&ul;
	do {
		ul = unlocked;
		ulp = (int*)&ul;
		_mm_pause();
	} while (__atomic_compare_exchange_n(&l->lock, ulp, (int*)&ll, true, __ATOMIC_ACQUIRE|__ATOMIC_HLE_ACQUIRE,  __ATOMIC_ACQUIRE|__ATOMIC_HLE_ACQUIRE) != true);
}

void release(struct locker *l) {
	__atomic_store_n(&l->lock, unlocked, __ATOMIC_RELEASE|__ATOMIC_HLE_RELEASE);
}

#cgo CFLAGS: -I./include
*/
import "C"

type Locker struct {
	l *C.struct_locker
}

func New() *Locker {
	return &Locker{
		l: C.locker_init(),
	}
}

func (l *Locker) Lock() {
	C.acquire(l.l)
}

func (l *Locker) Unlock() {
	C.release(l.l)
}
