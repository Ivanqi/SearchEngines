#ifndef INDEX_DICMAP_H
#define INDEX_DICMAP_H

#include <map>
#include <string>
#include "message.pb.h"
#include "FileStore.h"
#include "DicMapStruct.h"

using namespace Store;
using namespace Tools;

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

            void persistence(FileStore& storeService)
            {
                DictMapStruct dms; 
                for (auto dic: dic_) {
                    dms.key = dic.first;
                    dms.dv = dic.second;
                    storeService.appendBytes(dms);
                }
            }
    };
};
#endif