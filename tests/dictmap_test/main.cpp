#include <stdio.h>
#include <DictMap.h>

using namespace Index;

void test_case_1() {

    DictMap mapWriter;
    
    Message::DictValue dv_1, dv_2, dv_3;
    dv_1.set_offset(100);
    dv_1.set_length(200);
    dv_2.set_offset(200);
    dv_2.set_length(202);
    dv_3.set_offset(300);
    dv_3.set_length(203);

    FileStore fmapWriter("./map.dic");
    mapWriter.persistence(fmapWriter);


}

int main() {

    test_case_1();
    return 0;
}