#include <iostream>
using namespace std;

void ErrorHandling() {
  try {
    int val = 100;
    int z = 0;
    int res = val / z;
    cout << "the value is::" << res << endl;

  } catch (...) {
    cerr << "unknown error" << endl;
  }
}

int main() { ErrorHandling(); }
