#ifndef STORE_STORE_H
#define STORE_STORE_H

namespace Store
{
    class StoreService
    {
        public:
            virtual void readFullBytes(int64_t offset, int64_t len) = 0;
    };
};

#endif