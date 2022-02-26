#include <iostream>
#include <string>
#include <fstream>
#include <ostream>
#include "TriangularMatrix.h"

template<class T>
TriangularMatrix<T>::TriangularMatrix(int n, int m) : n(n), m(m) {
    allocSpace();
    this->b = std::vector<T>(n);
}

template<class T>
void TriangularMatrix<T>::allocSpace() {
    A = new T *[getN()];
    for (int i = 0; i < getN(); i++) {
        A[i] = new T[getM()];
    }
}

template<class T>
TriangularMatrix<T>::~TriangularMatrix() {
    for (int i = 0; i < getN(); ++i) {
        delete[] A[i];
    }
    delete[] A;
}

template<class T>
int TriangularMatrix<T>::getN() const {
    return this->n;
}

template<class T>
int TriangularMatrix<T>::getM() const {
    return this->m;
}

template<class T>
std::istream &operator>>(std::istream &in, TriangularMatrix<T> &A) {
    for (int i = 0; i < A.getN(); i++) {
        for (int j = 0; j < A.getM(); j++) {
            in >> A.A[i][j];
        }
        in >> A.b[i];
    }

    return in;
}

template<class T>
std::ostream &operator<<(std::ostream &out, const TriangularMatrix<T> &A) {
    for (int i = 0; i < A.getN(); i++) {
        for (int j = 0; j < A.getM(); j++) {
            out << A.A[i][j] << " ";
        }
        out << "| " << A.b[i] << std::endl;
    }

    return out;
}

template<class T>
T TriangularMatrix<T>::getColumnFreeMembers() const {
    return this->b.front();
}

template<class T>
std::vector<T> TriangularMatrix<T>::methodLower() {
    std::vector<T> x(n);

    x[0] = b[0] / A[0][0];

    for (int i = 1; i < n; i++) {
        auto sum = 0;
        for (int j = 0; j <= i - 1; j++)
            sum += A[i][j] * x[j];
        x[i] = (b[i] - sum) / A[i][i];
    }

    return x;
}

template<class T>
std::vector<T> TriangularMatrix<T>::methodHigher() {
    std::vector<T> x(n);

    x[n - 1] = b[n - 1] / A[n - 1][n - 1];

    for (int i = n - 2; i >= 0; i--) {
        auto sum = 0;
        for (int j = i + 1; j <= n - 1; j++)
            sum += A[i][j] * x[j];
        x[i] = (b[i] - sum) / A[i][i];
    }

    return x;
}