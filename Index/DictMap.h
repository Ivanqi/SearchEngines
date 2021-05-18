#ifndef INDEX_DICMAP_H
#define INDEX_DICMAP_H

#include <map>
#include <string>
#include <iostream>
#include "message.pb.h"
#include "FileStore.h"

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

            void persistence(FileStore storeService)
            {
                storeService.appendBytes(dic_);
            }

            void LoadDic(FileStore storeService)
            {
                std::map<std::string, Message::DictValue> dic;
                storeService.readFullBytes(&dic);
                dic_.swap(dic);
            }

            void toString()
            {
                for (auto dic: dic_) {
                    std::cout << "key:" << dic.first << " | offset: " << dic.second.offset() << ", length: " << dic.second.length() <<std::endl;
                }
            }
    };
};
#endif