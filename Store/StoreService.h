#ifndef STORE_STORE_H
#define STORE_STORE_H

#include <string>
#include "message.pb.h"

#define ALIGNMENT 4
#pragma pack(4)


namespace Store
{
    class StoreService
    {
        public:
            struct StoreHeader {
                unsigned int len;
            };

            struct StoreBody {
                const char* key;
                Message::DictValue dv;
            };

        public:

            virtual void readFullBytes(int64_t offset, int64_t len, std::map<std::string, Message::DictValue>* dic) = 0;

            virtual int64_t appendBytes(std::map<std::string, Message::DictValue>& dic) = 0;

            virtual unsigned int loadHeader() = 0;
    };
};

#endif