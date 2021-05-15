#ifndef INDEX_INVERT_WRITER_H
#define INDEX_INVERT_WRITER_H

#include <string>
#include <map>
#include <memory>

#include "DocList.h"
#include "NewMemoryDocList.h"
#include "message.pb.h"

using namespace Message;

namespace Index 
{
    class InvertWriter 
    {
        private:
            std::string name_;
            typedef std::shared_ptr<DocList> DocListPtr;
            std::map<std::string, DocListPtr> tmpInvert_;
        
        public:
            InvertWriter(std::string name)
                :name_(name)
            {

            }

            void put (std::string key, Message::DocId docid)
            {
                if (tmpInvert_.find(key) == tmpInvert_.end()) {
                    tmpInvert_[key] = DocListPtr(new NewMemoryDocList());
                }

                tmpInvert_[key]->Push(docid);
            }
    };
};

#endif