#include<iostream>

/*
g++ -o cpp.o -c cpp.cpp
g++ -o ./bin/cppprogram cpp.o
./bin/cppprogram
*/
class Program
{
private:
    std::string greeting;
public:
    Program() {
        greeting = "cpp, hello world";
    }
    int print() {
        std::cout<<greeting<<std::endl;
        return 1;
    }
};

int main() {
    Program p;
    p.print();
}