#ifndef STORE_FILEMMAP_STORE_H
#define STORE_FILEMMAP_STORE_H

#include <string>
#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <sys/mman.h>

namespace Store
{
    class FileMmapStore
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

            void NewSearchFileMMapStore()
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

    };
};



#endif