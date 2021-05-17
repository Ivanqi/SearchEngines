#ifndef TOOLS_DICMAP_STRUCT_H
#define TOOLS_DICMAP_STRUCT_H

#include <string>
#include <message.pb.h>

namespace Tools
{
    struct DictMapStruct
    {
        std::string key;
        Message::DictValue dv;
    };
};

#endif