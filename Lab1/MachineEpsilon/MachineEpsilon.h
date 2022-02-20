#ifndef MACHINEEPSILON_MACHINEEPSILON_H
#define MACHINEEPSILON_MACHINEEPSILON_H

template <class T>
class MachineEpsilon {
public:
    explicit MachineEpsilon(T epsilon) {
        this->epsilon = epsilon;
    }

public:
    void calculateMachineEpsilon();
    T getEpsilon() const;

private:
    T epsilon;
private:
    static T divideEpsByTwo(T epsilon);
};


#endif //MACHINEEPSILON_MACHINEEPSILON_H
