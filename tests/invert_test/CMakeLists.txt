  
cmake_minimum_required(VERSION 2.8)
project(invert_test CXX)

set(CMAKE_CXX_COMPILER "c++")
set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_STANDARD_REQUIRED True)
set(CMAKE_CXX_FLAGS -g)

message(STATUS "PROJECT_SOURCE_DIR ${PROJECT_SOURCE_DIR}")

INCLUDE_DIRECTORIES(
    ../../Index
    ../../Store
    ../../Message
)

set(src_dir
    ../../Message/message.pb.cc
)

add_executable(invert_test  ${src_dir} main.cpp)