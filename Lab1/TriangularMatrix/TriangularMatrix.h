#include <iostream>
#include <string>
#include <fstream>
#include <ostream>
#include <vector>

#ifndef UNTITLED_TRIANGULARMATRIX_H
#define UNTITLED_TRIANGULARMATRIX_H

template<class T>
class TriangularMatrix {
public:
    explicit TriangularMatrix(int n, int m);

    ~TriangularMatrix();

public:
    int getN() const;

    int getM() const;

    T getColumnFreeMembers() const;

    std::vector<T> methodLower();

    std::vector<T> methodHigher();

    template<typename U>
    friend std::istream &operator>>(std::istream &in, TriangularMatrix<T> &A);

    template<typename U>
    friend std::ostream &operator<<(std::ostream &out, const TriangularMatrix<T> &A);

public:
    T **A{};
    std::vector<T> b;

private:
    void allocSpace();

private:
    int n{};
    int m{};
};


#endif //UNTITLED_TRIANGULARMATRIX_H
