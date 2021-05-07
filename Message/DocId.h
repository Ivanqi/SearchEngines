#ifndef MESSAGE_DOCID_H
#define MESSAGE_DOCID_H

#include <string>

namespace Message
{
    class DocId
    {
        private:
            uint32_t docID_;     // 文档ID
            uint32_t weight_;    // 权重

        public:
            DocId(uint32_t DocID, uint32_t Weight)
                :docID_(DocID), weight_(Weight)
            {

            }

            std::string String(DocId *dId) 
            {
                return "";
            }
    };
};

#endif