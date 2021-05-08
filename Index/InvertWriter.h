#ifndef INDEX_INVERT_WRITER_H
#define INDEX_INVERT_WRITER_H

#include <string>
#include <map>

#include "DocList.h"
#include "NewMemoryDocList.h"
#include "Message/message.pb.h"

namespace Index 
{
    class InvertWriter 
    {
        private:
            std::string name_;
            std::map<std::string, DocList> tmpInvert_;
        
        public:
            InvertWriter(std::string name)
                :name_(name)
            {

            }

            void put (std::string key, Message::DocId docid)
            {
                if (tmpInvert_.find(key) == tmpInvert_.end()) {
                    tmpInvert_[key] = NewMemoryDocList();
                }

                tmpInvert_[key].Push(docid);
            }
    };
};

#endif