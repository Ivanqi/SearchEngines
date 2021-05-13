#ifndef INDEX_NEW_MEMORY_DOCLIST_H
#define INDEX_NEW_MEMORY_DOCLIST_H

#include "DocList.h"
#include "Message/message.pb.h"

#include <vector>
#include <stdio.h>

namespace Index
{
    class NewMemoryDocList: public DocList
    {
        private:
            std::vector<Message::DocId> docList_;
        
        public:
            NewMemoryDocList()
            {

            }

            int GetLength()
            {
                return docList_.size();
            }

            Message::DocId* GetDoc(int idx)
            {
                int docListSize = docList_.size();
                if (idx >= docListSize || idx < 0) {
                    return NULL;
                }

                return &docList_[idx];
            }

            bool Push(Message::DocId docid)
            {
                int docListSize = docList_.size();

                if (docListSize > 0 && docid.docid() <= docList_[docListSize - 1].docid()) {
                    printf("Doc Id [ %d ] is wrong,max id is : [ %d ]", docid.docid(), docList_[docListSize - 1].docid());
                    return false; 
                }

                docList_.push_back(docid);
                return true;
            }
    };
};

#endif