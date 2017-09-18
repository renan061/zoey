#include <assert.h>
#include <stdio.h>

#include "stack.h"

int main(void) {
	printf("Starting...\n");

	// stack_new
	Stack* stack = stack_new();
	// "stack_new == NULL"
	assert(stack);

	return 0;
}