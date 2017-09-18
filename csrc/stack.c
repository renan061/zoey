#include <stdlib.h>

#include "stack.h"

typedef struct StackNode StackNode;
struct StackNode {
	int value;
	StackNode* next;
};

struct Stack {
	StackNode* first;
};

Stack* stack_new(void) {
	Stack* stack = (Stack*)malloc(sizeof(Stack));
	stack->first = NULL;
	return stack;
}
