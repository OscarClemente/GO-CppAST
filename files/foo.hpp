#include <iostream>
#include "stdint"

class foo : IBar
{
    foo();
    virtual void function(int val);
    int var = 01234;
    float test = .234;
    int wuu = 0x1234u;
};