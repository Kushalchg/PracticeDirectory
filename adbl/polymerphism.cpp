#include <iostream>
using namespace std;

class Person {

private:
  string name;
  int age;

public:
  // this is constructor function with the same name a s the class
  Person(string personName, int personAge) {
    name = personName;
    age = personAge;
  }

  void displayPersonDetail() {
    cout << "name is: " << name << endl;
    cout << "Age is: " << age << endl;
  }
};

int main() {
  cout << "class example" << "\n" << endl;
  Person firstPerson("kushal", 11);

  firstPerson.displayPersonDetail();

  return 0;
};
