#include <iostream>
using namespace std;

struct Vector2 {
  float x, y;

  // Vector2(float x,float y):x(x),y(y){}

  // constructior to initialize the struct same as we use constructor to
  // initialize the class

  // this is equivalent to the
  Vector2(float x, float y) {
    this->x = x;
    this->y = y;
  }

  Vector2 Add(const Vector2 &others) const {
    return Vector2(x + others.x, y + others.y);
  }
  // for the operator overloading part

  Vector2 operator+(const Vector2 &others) const { return Add(others); };
};

int main() {
  Vector2 pos1(3.12f, 1.1f);
  Vector2 pos2(3.12f, 1.1f);

  Vector2 res = pos1.Add(pos2);
  Vector2 res1 = pos1 + pos2;

  // without overloading
  cout << "the x position is: " << res.x << endl;
  cout << "the y position is: " << res.y << endl;

  // with overloading
  cout << "the x position is(overloading): " << res1.x << endl;
  cout << "the y position is(overloading): " << res1.y << endl;

  return 0;
}
