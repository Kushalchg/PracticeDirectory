#include <iostream>

class Author {
private:
  std::string name;
  std::string email;

public:
  Author(std::string authorName, std::string authorEmail)
      : name(authorName), email(authorEmail) {};

  std::string getName() const { return name; }

  std::string getEmail() const { return email; }
};

class Book {
private:
  Author author;
  int price;
  std::string title;

public:
  Book(std::string authorName, std::string authorEmail, int bookPrice,
       std::string booktitle)
      : author(authorName, authorEmail), price(bookPrice), title(booktitle) {}

  void displayDetails() {
    std::cout << "Author name is ==> " << author.getName() << std::endl;
    std::cout << "Author email is ==> " << author.getEmail() << std::endl;
    std::cout << "Book Title is ==> " << title << std::endl;
    std::cout << "Price of book is  ==> " << price << std::endl;
  }
};

int main() {
  Book book("Kushal chapagain", "example@gmail.com", 200, "Atomic Havits");

  book.displayDetails();
  return 0;
}
