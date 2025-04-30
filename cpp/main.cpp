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
  Student S1;
  std::cout << S1.name << std::endl;
  std::cout << Student::count << std::endl;
  return 0;
}
