#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "cfunc.h"

extern void GoLog(struct msg *);
extern void GoLog2(struct msg *);
extern void GoLog3(struct msg *);

static char *default_msg = "This is a test";

void logger(char *body) {
	static struct msg msg;
	static char buf[256];

	clock_gettime(CLOCK_REALTIME, &msg.time);

	if (body == NULL) {
		msg.body = buf;
		strcpy(buf, default_msg);
	}

	GoLog(&msg);
}

void logger2(char *body) {
	struct msg msg;

	clock_gettime(CLOCK_REALTIME, &msg.time);

	if (body == NULL)
		msg.body = default_msg;

	GoLog2(&msg);
}

void logger3(char *body) {
	struct msg *msg = malloc(sizeof(struct msg));

	clock_gettime(CLOCK_REALTIME, &msg->time);

	if (body == NULL)
		msg->body = strdup(default_msg);

	GoLog3(msg);
}

void nop() {
}
