#include <stdio.h>
#include <FileStore.h>

using namespace Store;

void test_case_1()
{
    FileStore fs("./data/abc.ivt");
    bool ret = fs.NewFileStoreWriteService();
    printf("ret: %d\n", ret);
}

int main() {

    test_case_1();
    return 0;
}