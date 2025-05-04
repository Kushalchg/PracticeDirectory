#include <iostream>

// for concept of class and objects
class Student {
public:
  std::string name;
  static int count;

  static void display() { std::cout << count << std::endl; }
};
int Student::count = 10;

int main() {
  for (int i = 0; i < 10; i++) {
    std::cout << i << std::endl;
  }
  return 0;
}
