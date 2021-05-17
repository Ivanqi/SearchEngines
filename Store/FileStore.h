#ifndef STORE_FILESTORE_H
#define STORE_FILESTORE_H

#include <string>
#include <fstream> 
#include <stdio.h>
#include "DicMapStruct.h"

namespace Store
{
    class FileStore
    {
        private:
            std::string name_;
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

            bool newFileStoreWriteService()
            {
                pFile_  = fopen(name_.c_str(), "w+");
                return (pFile_ != NULL ? true : false);
            }

            void appendBytes(Tools::DictMapStruct& dms)
            {
                fwrite(&dms, sizeof(Tools::DictMapStruct), 1, pFile_);
            }

            // NewFileStoreReadService(std::string name)
    };
};


#endif