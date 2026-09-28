#include <string>
#include <random>
std::string ticketId(){static std::mt19937 g{std::random_device{}()};std::uniform_int_distribution<int>d(100000,999999);return "PASS-"+std::to_string(d(g));}
