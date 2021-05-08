#ifndef INDEX_DOCLIST_H
#define INDEX_DOCLIST_H

#include "Message/message.pb.h"

namespace Index
{
    // 倒排链
    class DocList
    {
        public:
            virtual int GetLength() = 0;
            virtual Message::DocId* GetDoc(int idx) = 0;
            virtual bool Push(Message::DocId docid) = 0;
    };
};


#endif