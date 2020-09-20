#include <iostream>
#include "stdint.h" // we should be using  #include <cstdint>

class foo : IBar
{
    foo();
    virtual void function(int val);
    virtual bool* function2(int& val);
    std::string string = "testing \"this\" works";
    std::string stringWithPrefix= R"hello";
    char c = 'f';
    char cWithPrefix = u'z';
    int num = 2451u;
    float = .234f;
    // yeah
};