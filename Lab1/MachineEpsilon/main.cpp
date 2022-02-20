#include "MachineEpsilon.h"
#include "MachineEpsilon.cpp"

//  Calculating the machine epsilon
template <class T>
void calcMachineEpsilon(T epsilon) {
    auto *machineEpsilon = new MachineEpsilon<T>(epsilon);
    machineEpsilon->calculateMachineEpsilon();

    delete machineEpsilon;
}

int main() {
    float epsilonFloat = 1.0;
    double epsilonDouble = 1.0;

    calcMachineEpsilon<float>(epsilonFloat);
    calcMachineEpsilon<double>(epsilonDouble);

    return 0;
}
