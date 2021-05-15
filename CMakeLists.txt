cmake_minimum_required(VERSION 2.8)

project(searchEngines C CXX)

enable_testing()

message("CMAKE_BUILD_BITS: " ${CMAKE_BUILD_BITS})
message("PROJECT_BINARY_DIR: " ${PROJECT_BINARY_DIR})
message("PROJECT_SOURCE_DIR:" ${PROJECT_SOURCE_DIR})
message("CMAKE_PROJECT_NAME:" ${CMAKE_PROJECT_NAME})


set(CXX_FLAGS
    -g
    -std=c++11
    -Werror
    -Wall
)

if(CMAKE_BUILD_BITS EQUAL 32)
    list(APPEND CXX_FLAGS "-m32")
endif()

string(REPLACE ";" " " CMAKE_CXX_FLAGS "${CXX_FLAGS}")

set(CMAKE_CXX_FLAGS_DEBUG "-O0")
set(CMAKE_CXX_FLAGS_RELEASE "-O2 -DNDEBUG")

include_directories(${PROJECT_SOURCE_DIR}/Index)
include_directories(${PROJECT_SOURCE_DIR}/Store)
# include_directories(./Index)

# add_subdirectory 指令
# 这个指令用于向当前工程添加存放源文件的子目录，并可以指定中间二进制和目标二进制存放的位置
# add_subdirectory(Index)

add_executable(searchEngines ./main.cpp)
# target_link_libraries(searchEngines Index)