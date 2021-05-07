#ifndef INDEX_INVERT_WRITER_H
#define INDEX_INVERT_WRITER_H

#include <string>
#include <map>
#include "Message/DocId.h"

using Message::DocId;

namespace Index 
{
    class InvertWriter 
    {
        private:
            std::string name_;
            std::map<DocId> tmpInvert_;
        
        public:
            InvertWriter(std::string name)
                :name_(name)
            {

            }

            void put (std::string key, DocId docid)
            {
                if (tmpInvert_.find(key) == tmpInvert_.end()) {
                    tmpInvert_[key] = docid;
                }
            }

    };
};

#endif