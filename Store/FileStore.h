#ifndef STORE_FILESTORE_H
#define STORE_FILESTORE_H

#include <string>
#include <fstream> 

namespace Store
{
    class FileStore
    {
        private:
            std::string name_;

        public:
            void NewFileStoreWriteService(std::string name)
            {
                name_ = name;
                std::ifstream ifs;

                ifs.open(name, std::ios::in | std::ios::out | std::ios::binary);
                ifs.close();
            }
    };
};


#endif