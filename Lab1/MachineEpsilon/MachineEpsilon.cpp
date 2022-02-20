#include "MachineEpsilon.h"
#include <iostream>

template<class T>
void MachineEpsilon<T>::calculateMachineEpsilon() {
    T eps = this->getEpsilon();

    do {
        eps = this->divideEpsByTwo(eps);
    } while (1 + eps > 1);

    std::cout << "Machine epsilon <" << typeid(eps).name() << "> = " << eps << std::endl;
}

template<class T>
T MachineEpsilon<T>::divideEpsByTwo(T epsilon) {
    return epsilon / 2;
}

template<class T>
T MachineEpsilon<T>::getEpsilon() const {
    return this->epsilon;
}