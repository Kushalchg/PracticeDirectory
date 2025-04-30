#include <iostream>

struct Node {
  int data;
  Node *next;

  // constructor to initialize the node

  /*Node(int data){*/
  /*  this->data=data;*/
  /*  this->next=nullptr;*/
  /*}*/

  // another way of assiging with constructor is
  Node(int nodeData) : data(nodeData), next(nullptr) {};
};

int CountLength(Node *head) {
  int count = 0;
  Node *curr = head;

  while (head->next != nullptr) {
    count++;
    curr = curr->next;
  }

  return count;
}

int main() {
  std::cout << "hello from cpp world";

  return 0;
}
