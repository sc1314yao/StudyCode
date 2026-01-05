cmake_minimum_required(VERSION 2.8)

project (4-addressbook)

ADD_DEFINITIONS(-W -Wall -D_REENTRANT -D_FILE_OFFSET_BITS=64 -DAC_HAS_INFO
-DAC_HAS_WARNING -DAC_HAS_ERROR -DAC_HAS_CRITICAL -DTIXML_USE_STL
-DAC_HAS_DEBUG -DLINUX_DAEMON -std=c++11)

set(pb_srcs addressbook.pb.cc)

# 指定protobuf链接库路径
link_directories(/usr/local/lib)
add_executable(add_person add_person.cc ${pb_srcs})
add_executable(list_people list_people.cc ${pb_srcs})
TARGET_LINK_LIBRARIES(add_person protobuf pthread)
TARGET_LINK_LIBRARIES(list_people protobuf pthread)





