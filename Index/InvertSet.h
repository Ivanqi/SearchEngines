#ifndef INDEX_INVERT_SET_H
#define INDEX_INVERT_SET_H

#include <string>
#include <map>

namespace Index
{
    class FieldInfo 
    {
        public:
            std::string name_;
            uint32_t type_;
            int64_t offset_;
    };


    class InvertSet
    {
        public:
            std::string name_;
            std::string path_;
            uint32_t mode_;
            std::map<std::string, FieldInfo> fieldInformations_;

            // 字符串倒排写服务
            std::map<std::string, std::string> stringInvertWriteServices_;

            // 字符串倒排读服务
            std::map<std::string, std::string> stringInvertReadServices_;

            std::string invertListStoreReader_;

            std::string dictStoreReader_;

        
        public:
            InvertSet(std::string name, std::string path)
                :name_(name), path_(path)
            {

            }

            void NewInvertSet()
            {
                std::string metaFile = path_ + "/" + name_ + ".mt";
                std::string ivtFile = path_ + "/" + name_ + ".ivt";
                std::string dicFile = path_ + "/" + name_ + ".dic";

                // 载入相关文件
            }
    };
};

#endif