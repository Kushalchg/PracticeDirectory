#include <iostream>
using namespace std;

class Stack {
private:
  int arr[6];
  int top;

public:
  Stack() : top(-1) {};
  void Push(int value) {
    if (top >= 5) {
      return;
    }
    arr[++top] = value;
    cout << value << "::popped";
  }

  void Pop() {
    if (top < 0) {
      cout << "Stack is empty" << endl;
      ;
      return;
    }
    int val = arr[top--];
    cout << val << "::popped";
  }
};

int main() {
  Stack s;
  s.Push(8);
  s.Pop();

  return 0;
}
