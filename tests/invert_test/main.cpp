#include <stdio.h>
#include <FileStore.h>
#include <InvertWriter.h>

using namespace Store;
using namespace Index;

void test_case_1()
{
    FileStore invertListStore("./abc.ivt");
    bool ret = invertListStore.newFileStoreWriteService();
    printf("ret: %d\n", ret);

    FileStore dictStore("./abc.dic");
    dictStore.newFileStoreWriteService();

    InvertWriter iw("abc");

    DocId dId_1, dId_2, dId_3, dId_4, dId_5, dId_6, dId_7;
    dId_1.set_docid(0);
    dId_1.set_weight(10);
    dId_2.set_docid(2);
    dId_2.set_weight(12);
    dId_3.set_docid(3);
    dId_3.set_weight(13);
    dId_4.set_docid(2);
    dId_4.set_weight(14);
    dId_5.set_docid(9);
    dId_5.set_weight(19);
    dId_6.set_docid(4);
    dId_6.set_weight(14);
    dId_7.set_docid(9);
    dId_7.set_weight(19);

    iw.put("abc", dId_1);
    iw.put("abc", dId_2);
    iw.put("abc", dId_3);
    iw.put("a", dId_4);
    iw.put("a", dId_5);
    iw.put("b", dId_6);
    iw.put("b", dId_7);
}

int main() {

    test_case_1();
    return 0;
}