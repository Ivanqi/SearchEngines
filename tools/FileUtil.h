#ifndef TOOLS_FILEUTIL_H
#define TOOLS_FILEUTIL_H

#include <sys/stat.h>
#include <unistd.h>
#include <string>

namespace Tools
{
    class FileUtil
    {
        public:
            bool static exists(std::string file)
            {
                return exists(file.c_str());
            }

            bool static exists(const char * file)
            {
                struct stat buf;
                stat(file, &buf);
                return (buf.st_size >= 0 ? true : false);
            }
    };
};

#endif