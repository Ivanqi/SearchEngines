#ifndef STORE_FILESTORE_H
#define STORE_FILESTORE_H

#include <string>
#include <fstream> 
#include <stdio.h>
#include <iostream>
#include <map>
#include <message.pb.h>
#include "DicMapStruct.h"

namespace Store
{
    class FileStore
    {
        private:
            std::string name_;

        public:
            FileStore(std::string name)
                :name_(name)
            {

            }

            ~FileStore()
            {
            }

            bool newFileStoreWriteService()
            {
                FILE *pFile  = fopen(name_.c_str(), "wb+");
                if (!pFile) {
                    return false;
                }
                fclose(pFile);
                return (pFile != NULL ? true : false);
            }

            void appendBytes(std::string key, Message::DictValue dv)
            {
                Tools::DictMapStruct dms;
                dms.key = key.c_str();
                dms.dv = dv;

                FILE *pFile  = fopen(name_.c_str(), "w+");
                fwrite(&dms, sizeof(Tools::DictMapStruct), 1, pFile);
                fclose(pFile);
            }

            void appendBytes(std::map<std::string, Message::DictValue>& dic)
            {
                Tools::DictMapStruct dms;
                FILE *pFile  = fopen(name_.c_str(), "wb+");
                if (!pFile) {
                    printf("打开文件失败!");
                    return;
                }

                for (auto d: dic) {
                    dms.key = d.first.c_str();
                    dms.dv = d.second;
                    fwrite(&dms, sizeof(Tools::DictMapStruct), 1, pFile);
                }

                fclose(pFile);
            }

            void readFullBytes(std::map<std::string, Message::DictValue>* dic)
            {
                FILE *pFile  = fopen(name_.c_str(), "rb+");
                if (!pFile) {
                    printf("打开文件失败!");
                    return;
                }

                fseek(pFile, 0L, SEEK_END);    // 定位到文件尾
                int len = ftell(pFile);        // 获取到文件大小 字节数
                rewind(pFile);                 // 文件指针复位 到文件头

                int i = 0;
                Tools::DictMapStruct dms;
                std::map<std::string, Message::DictValue> tmp;

                while (i < len / sizeof(Tools::DictMapStruct)) {
                    fread(&dms, sizeof(Tools::DictMapStruct), 1, pFile);
                    i++;
                    std::string key(dms.key);
                    tmp[key] = dms.dv;
                }

                dic->swap(tmp);
            }

            // NewFileStoreReadService(std::string name)
    };
};


#endif