#include <assert.h>

#include "stack.h"

int main(void) {
	// stack_new
	Stack* stack = stack_new();
	assert(stack);

	// stack_push
	for (int i = 0; i < 10; i++) {
		stack_push(stack, i + 1);
	}

	// stack_push
	for (int i = 0; i < 10; i++) {
		assert(stack_pop(stack) == 10 - i);
	}

	return 0;
}
