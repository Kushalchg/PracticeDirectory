# Makefile

# C++ Variables
CXX = g++
CXXFLAGS = -Wall -g
CPPTARGET = main
CPPSRC = main.cpp

# C Variables
CC = gcc
CFLAGS = -Wall -g
CTARGET = cmain
CSRC = main.c

# Default targets
cpp: cppbuild cpprun
c: cbuild crun

# C++ build and run
cppbuild:
	@$(CXX) $(CXXFLAGS) -o $(CPPTARGET) $(CPPSRC)

cpprun:
	@./$(CPPTARGET)

cppclean:
	rm -f $(CPPTARGET)

# C build and run
cbuild:
	@$(CC) $(CFLAGS) -o $(CTARGET) $(CSRC)

crun:
	@./$(CTARGET)

cclean:
	rm -f $(CTARGET)

# Clean all
clean: cppclean cclean
