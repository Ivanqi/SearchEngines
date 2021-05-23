#ifndef STORE_FILEMMAP_STORE_H
#define STORE_FILEMMAP_STORE_H

#include <string>
#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <sys/mman.h>

#include "StoreService.h"

namespace Store
{
    class FileMmapStore: public StoreService
    {
        private:
            std::string name_;
            int *mmapBytes_;
            long fileSize_;
        
        public:
            FileMmapStore(std::string name)
                :name_(name)
            {

            }

            ~FileMmapStore()
            {
                munmap(mmapBytes_, fileSize_);
            }

            void newSearchFileMMapStore()
            {
                int fd = open(name_.c_str(), O_RDWR);
                FILE *f = fdopen(fd, "rw");
                fseek(f, 0, SEEK_END);
                fileSize_ = ftell(f);

                // 分配的大小必须为页的整数倍
                mmapBytes_ = (int*)mmap(NULL, fileSize_, PROT_WRITE|PROT_READ, MAP_SHARED, fd, 0);
                if (mmapBytes_ == MAP_FAILED) {
                    printf("mmap error");
                }
                close(fd);
            }

            unsigned int loadHeader()
            {
                unsigned long shLen = sizeof(StoreHeader);
                // 获取头部数据
                StoreHeader sh;
                memcpy(&sh, mmapBytes_, shLen);
                unsigned int headerLen = sh.len;      // 数据长度
                return headerLen;
            }

            void readFullBytes(int64_t offset, int64_t len, std::map<std::string, Message::DictValue>* dic)
            {  
                unsigned long sbLen = sizeof(StoreBody) / ALIGNMENT;
                len /= ALIGNMENT;

                // 略过头部数据字节
                int *mmapBytes = mmapBytes_ + (sizeof(StoreHeader) / ALIGNMENT);        

                std::map<std::string, Message::DictValue> tmp;

                if (len >= sbLen) {
                    while (len >= sbLen && len >= 0) {
                        StoreBody sb;
                        memcpy(&sb, mmapBytes + offset, sbLen * ALIGNMENT);
                        len -= sbLen;
                        offset += sbLen;
                        std::string key(sb.key);
                        tmp[key] = sb.dv;
                    }
                }

                dic->swap(tmp);
            }

            int64_t appendBytes(std::map<std::string, Message::DictValue>& dic)
            {
                return 0;
            }

    };
};



#endif