#include "TriangularMatrix.h"
#include "TriangularMatrix.cpp"
#include <iostream>

int main() {
    std::string filePath = "../matrix.txt";
    std::ifstream file;
    file.open(filePath);

    int n, m;
    file >> n >> m;

    if (!file.is_open()) {
        std::cout << "The file could not be opened";
        return 0;
    }

    auto *A = new TriangularMatrix<float>(n, m);

    std::cout << "ROWS: " << n << std::endl << "COLUMNS: " << m << std::endl;

    file >> *A;
    file.close();

    std::cout << *A << std::endl;

    auto resultMethodHigher = A->methodHigher();
    auto iter = resultMethodHigher.begin();

    std::cout << "RESULT:" << std::endl;
    for (; iter != resultMethodHigher.end(); ++iter)
        std::cout << *iter << " ";

    return 0;
}
