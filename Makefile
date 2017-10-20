# 
# Scripts
# 

CC := gcc-7
CFLAGS := -std=c99 -O0 -Wall

main: clean

run: clean setup
	@- go run zoey.go

ctest: clean setup
	@- echo "-------------------------"
	@- $(CC) $(CFLAGS) -fsyntax-only -c csrc/stack.c
	@- echo "-------------------------"

test: setup
	@- go test ./...

ex: setup
	@- gcc -c ex/stack.c -o ex/stack.o
	@- gcc ex/stack.o ex/assert.c -o ex/assert.out
	@- ./ex/assert.out

	@- gcc ex/output.c -o ex/output.out
	@- ./ex/output.out &> ex/output.result 

	@- rm -f ex/*.out ex/*.o ex/*.result

factorial: main
	@- gcc -c

setup:
	@- go fmt ./...

clean:
	@- rm -f main
	@- rm -f *.o
	@- rm -f *.out


# # TODO: gnu11 / c11
# # TODO: llvm include
# CC := gcc-7 
# CFLAGS := -I/usr/local/opt/llvm/include -std=gnu11 -Wall

# CPPFLAGS := `llvm-config --cxxflags --ldflags --system-libs`
# CPPFLAGS += `llvm-config --libs analysis bitwriter core executionengine `
# CPPFLAGS += `llvm-config --libs target mcjit native`

# main: objs
#   @- $(CC) $(CFLAGS) -c src/aria.c -o obj/aria.o

#   @- clang++ $(CPPFLAGS) obj/errs.o obj/vector.o  \
#   obj/scanner.o obj/parser.o obj/ast.o            \
#   obj/symtable.o obj/sem.o obj/backend.o          \
#   obj/aria.o -o bin/aria

# objs: errs vector parser scanner ast sem backend

# # ==================================================
# # 
# # Modules
# # 
# # ==================================================

# errs:
#   @- $(CC) $(CFLAGS) -c src/errs.c -o obj/errs.o

# vector:
#   @- $(CC) $(CFLAGS) -c src/vector.c -o obj/vector.o

# scanner:
#   @- flex src/scanner.l
#   @- mv lex.yy.c src/flex.c
#   @- $(CC) $(CFLAGS) -c src/flex.c -o obj/scanner.o -Isrc/

# parser:
#   @- bison -v --defines=src/bison.h src/parser.y
#   @- mv parser.tab.c src/bison.c
#   @- mkdir -p temp
#   @- mv parser.output temp/bison.output
#   @- $(CC) $(CFLAGS) -c src/bison.c -o obj/parser.o

# ast:
#   @- $(CC) $(CFLAGS) -c src/ast.c -o obj/ast.o

# sem:
#   @- $(CC) $(CFLAGS) -c src/symtable.c -o obj/symtable.o
#   @- $(CC) $(CFLAGS) -c src/sem.c -o obj/sem.o

# backend:
#   @- $(CC) $(CFLAGS) -c src/backend.c -o obj/backend.o

# # ==================================================
# # 
# # Tests
# # 
# # ==================================================

# vector_test: errs vector
#   @- $(CC) $(CFLAGS) -o bin/vectortest    \
#   obj/errs.o obj/vector.o                 \
#   tests/src/vector_test.c -Isrc/
	
#   @- ./bin/vectortest

# scanner_test: errs vector parser scanner ast
#   @- $(CC) $(CFLAGS) -o bin/scannertest                           \
#   obj/errs.o obj/vector.o obj/scanner.o obj/parser.o obj/ast.o    \
#   tests/src/scanner_test.c -Isrc/

#   @- sh tests/test.sh scanner

# parser_test: errs vector parser scanner ast
#   @- $(CC) $(CFLAGS) -o bin/parsertest                            \
#   obj/errs.o obj/vector.o obj/scanner.o obj/parser.o obj/ast.o    \
#   tests/src/parser_test.c -Isrc/

#   @- sh tests/test.sh parser

# ast_test: errs vector parser scanner ast
#   @- $(CC) $(CFLAGS) -o bin/asttest                               \
#   obj/errs.o obj/vector.o obj/scanner.o obj/parser.o obj/ast.o    \
#   tests/src/ast_test.c -Isrc/

#   @- sh tests/test.sh ast

# sem_test: errs vector parser scanner ast sem
#   @- $(CC) $(CFLAGS) -o bin/semtest                               \
#   obj/errs.o obj/vector.o obj/scanner.o obj/parser.o obj/ast.o    \
#   obj/symtable.o obj/sem.o                                        \
#   tests/src/sem_test.c -Isrc/

#   @- sh tests/test.sh sem

# backend_test: main
#   @- mv bin/aria bin/backendtest
#   @- sh tests/test.sh backend

# test: clean vector_test scanner_test parser_test ast_test sem_test backend_test

# # ==================================================
# # 
# # Misc
# # 
# # ==================================================

# clean:
#   @- rm -f src/flex.c
#   @- rm -f src/bison.*
#   @- rm -rf temp
#   @- rm -f obj/*
#   @- rm -f bin/*
#   @- rm -f *.bc
#   @- rm -f tests/backend/*.bc
#   @- rm -f *.ll
