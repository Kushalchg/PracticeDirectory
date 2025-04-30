#include <iostream>
using namespace std;

class Abstraction {
private:
  string secrateWord = "suhhh";
  void SecrateMessage() { cout << "this is secrate" << endl; }

public:
  int x, y;
  string commonWord = "broadcast";
  void visible() {
    cout << "secrate message from same class: " << secrateWord << endl;
  }

  void accessSecreate() {
    cout << "secreate message is here" << "\n" << endl;
    SecrateMessage();
  }

  Abstraction(int x, int y) {
    this->x = x;
    this->y = y;
  }
};

int main() {
  Abstraction a(3, 4);

  a.visible();
  /*a.SecrateMessage();*/ // this is not allowed because the secrateMessage is
  // private funciton of Abstraction class

  a.accessSecreate();

  return 0;
}
