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
            std::ifstream ifs_;

        public:
            FileStore(std::string name)
                :name_(name)
            {

            }

            ~FileStore()
            {
                if (ifs_) {
                    ifs_.close();
                }
            }

            bool NewFileStoreWriteService()
            {

                ifs_.open(name_, std::ios::in | std::ios::out | std::ios::binary);
                return ifs_ ? true : false;
            }

            // NewFileStoreReadService(std::string name)
    };
};


#endif