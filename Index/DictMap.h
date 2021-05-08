#ifndef INDEX_DICMAP_H
#define INDEX_DICMAP_H

#include <map>
#include <string>
#include "Message/message.pb.h"

namespace Index
{
    class DictMap
    {
        private:
            uint64_t storeBodyLength_;
            int64_t startOffset_;
            std::map<std::string, Message::DictValue> dic_;
        
        public:
            void put(std::string key, Message::DictValue dv)
            {
                dic_[key] = dv;
            }
    };
};
#endif