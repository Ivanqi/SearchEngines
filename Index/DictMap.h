#ifndef INDEX_DICMAP_H
#define INDEX_DICMAP_H

#include <map>
#include <string>
#include <iostream>
#include "message.pb.h"
#include "StoreService.h"

using namespace Store;

namespace Index
{
    class DictMap
    {
        private:
            int64_t storeBodyLength_;
            int64_t startOffset_;
            std::map<std::string, Message::DictValue> dic_;
        
        public:
            void put(std::string key, Message::DictValue dv)
            {
                dic_[key] = dv;
            }

            void persistence(StoreService *storeService)
            {
                startOffset_ = storeService->appendBytes(dic_);
            }

            void LoadDic(StoreService *storeService)
            {
                std::map<std::string, Message::DictValue> dic;
                storeBodyLength_ = storeService->loadHeader();
                
                storeService->readFullBytes(startOffset_, storeBodyLength_, &dic);
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