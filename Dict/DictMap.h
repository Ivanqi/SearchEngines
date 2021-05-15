#ifndef DICT_DICTMAP_H
#define DICT_DICTMAP_H

#include <map>
#include <string>
#include "message.pb.h"

namespace Dict
{
    class DictMap
    {
        private:
            uint64_t storeBodyLength_;
            int64_t startOffset_;
            std::map<std::string, Message::DictValue> dic_;
        
        public:
            DictMap()
            {

            }
    };
};

#endif