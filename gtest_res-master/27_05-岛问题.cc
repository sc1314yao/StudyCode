#include <iostream>
#include <initializer_list>
#include <vector>
#include <gtest/gtest.h>

using namespace std;

class IslandProblem {
public:
    using Matrix = vector<vector<char>>;
    IslandProblem(const initializer_list<vector<char>> list) {
        _islands.assign(list);
    }

    int Do() {
        int num = 0;
        for (int row = 0; row < (int)_islands.size(); row++) {
            for (int col = 0; col < (int)_islands[row].size(); col++) {
                if (canUnion(row, col)) {
                    num++;
                    unionIsland(row, col);
                }
            }
        }
        return num;
    }

protected:
    bool canUnion(int row, int col) {
        if (row < 0 || row >= (int)_islands.size())
            return false;
        if (col < 0 || col >= (int)_islands[row].size())
            return false;
        if (_islands[row][col] != 1)
            return false;
        return true;
    }
    void unionIsland(int row, int col) {
        _islands[row][col] = 2;
        // up
        if (canUnion(row-1, col)) unionIsland(row-1, col);
        // left
        if (canUnion(row, col-1)) unionIsland(row, col-1);
        // down
        if (canUnion(row+1, col)) unionIsland(row+1, col);
        // right
        if (canUnion(row, col+1)) unionIsland(row, col+1);
    }

private:
    Matrix _islands;
};

TEST(IslandProblem, logic) {
    IslandProblem ip1{
        {1,1,1,1},
        {1,0,1,1},
        {0,0,0,0},
        {1,0,1,0}
    };
    EXPECT_EQ(ip1.Do(), 3);

    IslandProblem ip2{
        {1,0,1,1},
        {1,0,1,1},
        {0,0,0,0},
        {1,0,1,0}
    };
    EXPECT_EQ(ip2.Do(), 4);
}

TEST(IslandProblem, boundary) {
    IslandProblem ip1{
        {1,1,1,1},
        {1,0,0,1},
        {1,0,0,1},
        {1,1,1,1}
    };
    EXPECT_EQ(ip1.Do(), 1);
    IslandProblem ip2{
    };
    EXPECT_EQ(ip2.Do(), 0);
}

TEST(IslandProblem, exception) {
    IslandProblem ip1{
        {-1,1,1,1},
        {1,0,0,1},
        {1,0,0,1},
        {1,1,1,1}
    };
    EXPECT_EQ(ip1.Do(), 1);
}