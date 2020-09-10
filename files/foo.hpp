#include <iostream>

class foo : IBar
{
    foo();
    virtual void function(int val);
    std::string test = R"hello";
    char c = u'z';
};