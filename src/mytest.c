#include "mytest.h"
#include <stdio.h>

void display(MyTest* test) {
    printf("{%d, %d}\n", test ? test->a : 0, test ? test->b : 0);
}