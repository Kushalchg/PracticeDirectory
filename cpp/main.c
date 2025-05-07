#include <stdio.h>

int main() {
  int *ptr;
  int num = 10;

  ptr = &num;
  // by assigning the ptr 100 it will overwirte the content of address of ptr
  //(i.e num address)
  *ptr = 100;

  printf("the value of a is::%d\n", num);
  printf("the value of ptr is::%d\n", *ptr);
  return 0;
}
