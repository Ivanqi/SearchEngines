#ifndef STORE_FILESTORE_H
#define STORE_FILESTORE_H

#include <string>
#include <fstream> 
#include <stdio.h>

namespace Store
{
    class FileStore
    {
        private:
            std::string name_;
            std::ifstream ifs_;
            FILE *pFile_;

        public:
            FileStore(std::string name)
                :name_(name)
            {

            }

            ~FileStore()
            {
                if (pFile_ != NULL) {
                    fclose(pFile_);
                }
            }

            bool NewFileStoreWriteService()
            {
                pFile_  = fopen(name_.c_str(), "w+");
                return (pFile_ != NULL ? true : false);
            }

            // NewFileStoreReadService(std::string name)
    };
};


#endif