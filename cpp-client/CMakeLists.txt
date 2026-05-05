cmake_minimum_required(VERSION 3.16)

project(EmbyAPIDemo)
set(CMAKE_CXX_STANDARD 20)

# C++ Executable
add_executable(EmbyAPIDemo main.cpp jbgolib.hpp jbparser.hpp json.hpp)
# Copy dll
add_custom_command(TARGET EmbyAPIDemo POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${CMAKE_SOURCE_DIR}/../export/jbembyapi.dll"
        "$<TARGET_FILE_DIR:EmbyAPIDemo>"
)
