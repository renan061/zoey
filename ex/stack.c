#include <stdlib.h>

#include "stack.h"

typedef struct Node Node;
struct Node {
	int value;
	Node* down;
};

struct Stack {
	Node* top;
};

Stack* stack_new(void) {
	Stack* stack = (Stack*)malloc(sizeof(Stack));
	if (!stack) {
		exit(1);
	}
	stack->top = NULL;
	return stack;
}

void stack_push(Stack* stack, int n) {
	Node* node = (Node*)malloc(sizeof(Node));
	if (!node) {
		exit(1);
	}
	node->value = n;
	node->down = stack->top;
	stack->top = node;
}

int stack_pop(Stack* stack) {
	if (!stack->top) {
		return 0;
	}
	Node* temp = stack->top;
	int n = temp->value;
	stack->top = stack->top->down;
	free(temp);
	return n;
}

void stack_free(Stack* stack) {
	Node* temp = NULL;
	for (Node* node = stack->top; node;) {
		temp = node->down;
		free(node);
		node = temp;
	}
}
