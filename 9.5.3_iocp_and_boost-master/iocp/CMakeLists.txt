cmake_minimum_required(VERSION 3.26)

project(iocp-test)

add_executable(iocp iocp.cpp)

target_link_libraries(iocp WS2_32 mswsock)