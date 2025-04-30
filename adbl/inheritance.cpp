#include <iostream>
using namespace std;

class Animal {
private:
  void privateSkill() { cout << "thie is priveate skill of animal" << endl; }

public:
  // calling  constructor for animal class
  Animal() { cout << "Constructor called" << endl; }
  void eat() { cout << "Animal is eating" << endl; }

  void walking() { cout << "Animal is walking" << endl; }

  void run() { cout << "Animal is running" << endl; }

  ~Animal() { cout << "Animal Distructor called" << endl; }
};

class Dog : private Animal {

public:
  void eat() {

  };

  void sound() { cout << "bark" << endl; }
};

int main() {
  Dog dog;
  dog.sound();
  return 0;
}
