#ifndef STACK_H
#define STACK_H

typedef struct Stack Stack;

extern Stack* stack_new(void);
extern void stack_push(Stack*, int);
extern int stack_pop(Stack*);
extern void stack_free(Stack*);

#endif