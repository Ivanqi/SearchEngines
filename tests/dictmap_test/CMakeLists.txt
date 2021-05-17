  
cmake_minimum_required(VERSION 2.8)
project(dictmap_test CXX)

set(CMAKE_CXX_COMPILER "c++")
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)
set(CMAKE_CXX_FLAGS -g)
set(CXX_FLAGS -D_GLIBCXX_USE_CXX11_ABI=0)

message(STATUS "PROJECT_SOURCE_DIR ${PROJECT_SOURCE_DIR}")

INCLUDE_DIRECTORIES(
    ../../Index
    ../../Store
    ../../Message
    ../../Dict
    ../../tools
)

set(src_dir
    ../../Message/message.pb.cc
)

add_executable(dictmap_test ${src_dir} main.cpp)

target_link_libraries(dictmap_test protobuf)