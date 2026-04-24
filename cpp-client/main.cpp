// main.cpp

#include <iostream>
#include "parser.hpp"
#include "jbgolib.hpp"

int main() {
    GoLib lib;
        if (!lib.load()) {
            return -1;
        }
        const std::string host = "<hostname>";
        const std::string port = "<port>";
        const std::string username = "<username>";
        const std::string password = "<password>";
        std::string res = lib.UserLoginToEmbyServer(false, host, port, username, password);
        EmbyLogonResult r = parseLogonResult(res);
        if (r.code != 0) {
            std::cerr << r.message << "\n";
            return 1;
        }
        std::string views = lib.UserGetViews(r.baseUrl, r.userId, r.accessToken);
        std::cout << "User Views Result: " << views << "\n";
}
