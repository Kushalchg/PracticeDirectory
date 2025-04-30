#include <iostream>
using namespace std;
//==============================================================================
// for function overloading example
void sound(std::string name, std::string sound) {
  std::cout << name << " makes " << sound << endl;
}

void sound(int count) { cout << "total" << count << "  cat" << endl; }
void sound(std::string count) { cout << "total" << count << "  cat" << endl; }

// we can call function with the same name with different number of parameter
// and type of parameter
void functionOverload() {
  sound("cat", "meow");
  sound(4);
  sound("black");
}

//==============================================================================
// example which demonstrate the default argument example
void printName(std::string myName, std::string yourName = "chapagain") {
  std::cout << "my name is " << myName << endl;
  std::cout << "your name is " << yourName << endl;
}
//==============================================================================
// Concept of Object and class and friend funciton

/*
  -- Private data of class is accessable to the class inherited from that
class(inheritance purpose).
-- Friend class can access all the data of class (private,protected and public).

*/
class Student {
  // by default the the function and veraible are private type
  std::string privateName = "private name";
  void privateShow() {
    std::cout << "You are previllage to be on private show::" << privateName
              << std::endl;
  }
  friend class Teacher;

protected:
  std::string protectedName = "Protected name";
  void protectedShow() {
    std::cout << "Don't worry this is protected show::" << protectedName
              << std::endl;
  }

public:
  int age;
  void display() {
    privateShow();
    protectedShow();
    std::cout << "name is " << privateName << " and age is " << age
              << std::endl;
  }
};

class Teacher {
public:
  void showPrivateStudent() {
    Student s1;
    std::cout << "Student Private information is leaked thourgh teacher::"
              << std::endl;
    s1.privateShow();
  }
};

//==============================================================================
// Constructor and Destructor
/*
  constructor is special function what will called when instance of a
  class(object) is created. syntax:: classname(){ logic that will execute when
  class created
  }

  Destructor is another function that will called when the instance of
  class(object) goes out of scope

*/

class Special {

public:
  std::string message;
  // constructor
  Special() { message = "this is special message created with object"; }
  // Destructor
  ~Special() {
    std::cout << "the value of message will be garbage after object goes out "
                 "of scope";
  }
};

void constructorAndDestructure() {
  Special special;
  std::cout << special.message << std::endl;
}

//==============================================================================
// Static data member and static funciton
/*
 -- stateic data member is the type of data member in class which can be access
outside of class without creating the instance of class(object)
-- Static function is the type of function which can be access outside of class
without creating class object.
-- Static function only access the static data memeber in normal way if we need
to pass the normal data member into the static functioin we need topass the
pointer of object as parameter of static function.

*/

class Principal {
public:
  std::string name;
  static int count;

  static void display() { std::cout << count << std::endl; }
};
int Principal::count = 10;

int main() {
  Principal S1;
  std::cout << S1.name << std::endl;
  std::cout << Principal::count << std::endl;
  return 0;
}
