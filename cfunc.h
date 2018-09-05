#ifndef __cfunc_h__
#define __cfunc_h__

#include <time.h>
#include <stdlib.h>

extern void logger(char *);
extern void logger2(char *);
extern void logger3(char *);
extern void nop();

struct msg {
	struct timespec time;
	char *body;
};

#endif
