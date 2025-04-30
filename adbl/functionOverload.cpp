#include <iostream>

using namespace std;

float area(int l) { return l * l; }

float area(int w, int h) { return w * h; }

float area(float r) {
  const double pi = 3.14;
  return pi * r * r;
}

int main() {
  cout << "area of square" << area(4) << endl;
  cout << "area of ractangle" << area(2, 3) << endl;
  cout << "area of circle" << area(3.0f) << endl;

  return 0;
}
