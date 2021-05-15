cmake_minimum_required(VERSION 2.8)

# 在目录中查找所有源文件
# aux_source_directory(<dir> <variable>) 收集指定目录中所有源文件的名称，并将列表存储在提供的<variable>中
aux_source_directory(. DIR_MESSAGE_SRCS)
# 添加一个Index的库
add_library(Index ${DIR_MESSAGE_SRCS})