#include <stdio.h>
#include <DictMap.h>
#include "FileStore.h"
#include "FileMmapStore.h"

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

    mapWriter.put("hello", dv_1);
    mapWriter.put("hello2", dv_2);
    mapWriter.put("hello3", dv_3);

    FileStore *fmapWriter = new FileStore("./map.dic");


    for (int i = 4; i < 100; i++) {
        std::string key = "hello" + std::to_string(i);
        Message::DictValue dv;
        dv.set_offset(100);
        dv.set_length(200);
        mapWriter.put(key, dv);
    }

    mapWriter.persistence(fmapWriter);

    FileMmapStore *fms = new FileMmapStore("./map.dic");
    fms->newSearchFileMMapStore();

    mapWriter.LoadDic(fms);
    mapWriter.toString();
}

int main() {

    test_case_1();
    return 0;
}