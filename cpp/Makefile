# File: Makefile

# Variables
CXX = g++
CXXFLAGS = -Wall -g
TARGET = main
SRC = main.cpp

# Default target
all: build run

# Build target
build:
	@$(CXX) $(CXXFLAGS) -o $(TARGET) $(SRC)

# Run target
run:
	@./$(TARGET)

# Clean target
clean:
	rm -f $(TARGET)

