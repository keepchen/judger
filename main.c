#include <iostream>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[])
{
    char input[1000];
    scanf("%s", input);
    std::cout << input << std::endl;
    std::cout << argc << std::endl;
    std::cout << argv[0] << std::endl;
    std::cout << "hello world!" << std::endl;
    sleep(2);
    return 0;
}