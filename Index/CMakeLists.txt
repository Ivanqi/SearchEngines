# add_library(Index)

# # 添加install操作
# install(TARGETS Index DESTINATION lib)

# # GLOB 会产生一个由所有匹配globbing表达式的文件组成的列表，并将其保存到变量中
# # 以下命令会把后面匹配到的所有文件交给GLOB子命令，由后者生成一个文件列表，并将列表赋值给HEADERS变量
# file(GLOB HEADERS "*.h")
# # 安装HEADERS, 把HEADERS内容添加到networker/base中
# install(FILES ${HEADERS} DESTINATION include/Index)

cmake_minimum_required(VERSION 2.8)

# 在目录中查找所有源文件
# aux_source_directory(<dir> <variable>) 收集指定目录中所有源文件的名称，并将列表存储在提供的<variable>中
aux_source_directory(. DIR_INDEX_SRCS)
# 添加一个Index的库
add_library(Index ${DIR_INDEX_SRCS})