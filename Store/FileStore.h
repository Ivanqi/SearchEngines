#ifndef STORE_FILESTORE_H
#define STORE_FILESTORE_H

#include <string>
#include <fstream> 
#include <stdio.h>
#include <iostream>
#include <map>
#include <message.pb.h>
#include "StoreService.h"

namespace Store
{
    class FileStore: public StoreService
    {
        private:
            std::string name_;
            int64_t end_;

        public:
            FileStore(std::string name)
                :name_(name), end_(0)
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
                StoreBody dms;
                dms.key = key.c_str();
                dms.dv = dv;

                FILE *pFile  = fopen(name_.c_str(), "w+");
                fwrite(&dms, sizeof(StoreBody), 1, pFile);
                fclose(pFile);
            }

            int64_t appendBytes(std::map<std::string, Message::DictValue>& dic)
            {
                FILE *pFile  = fopen(name_.c_str(), "wb+");
                if (!pFile) {
                    printf("打开文件失败!");
                    return -1;
                }

                StoreBody sb;
                StoreHeader sh;

                unsigned long sbLen = sizeof(StoreBody);
                int64_t ret = end_;
                int64_t len = 0;

                // 写入头部信息
                int dictSize = dic.size();
                sh.len = dictSize * sbLen;
                fwrite(&sh, sizeof(StoreHeader), 1, pFile);

                // 写入内容
                for (auto d: dic) {
                    sb.key = d.first.c_str();
                    sb.dv = d.second;
                    fwrite(&sb, sbLen, 1, pFile);
                }

                fclose(pFile);
                return ret;
            }

            unsigned int loadHeader()
            {
                FILE *pFile  = fopen(name_.c_str(), "rb+");
                if (!pFile) {
                    printf("打开文件失败!");
                    return -1;
                }

                 // 获取头部数据
                StoreHeader sh;
                fread(&sh, sizeof(StoreHeader), 1, pFile);
                fclose(pFile);
                unsigned int shLen = sh.len;      // 数据长度
                return shLen;
            }

            void readFullBytes(int64_t offset, int64_t len, std::map<std::string, Message::DictValue>* dic)
            {
                FILE *pFile  = fopen(name_.c_str(), "rb+");
                if (!pFile) {
                    printf("打开文件失败!");
                    return;
                }

                // 略过头部数据字节
                fseek(pFile, sizeof(StoreHeader), SEEK_SET);

                int i = 0;
                StoreBody sb;
                std::map<std::string, Message::DictValue> tmp;
                unsigned long sbLen = sizeof(StoreBody);

                while (i < len / sbLen) {
                    fread(&sb, sbLen, 1, pFile);
                    i++;
                    std::string key(sb.key);
                    tmp[key] = sb.dv;
                }

                dic->swap(tmp);
                fclose(pFile);
            }
    };
};


#endif