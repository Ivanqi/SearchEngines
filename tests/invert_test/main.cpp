#include <stdio.h>
#include <FileStore.h>
#include <InvertWriter.h>

using namespace Store;
using namespace Index;

void test_case_1()
{
    FileStore invertListStore("./abc.ivt");
    bool ret = invertListStore.NewFileStoreWriteService();
    printf("ret: %d\n", ret);

    FileStore dictStore("./abc.dic");
    dictStore.NewFileStoreWriteService();

}

int main() {

    test_case_1();
    return 0;
}