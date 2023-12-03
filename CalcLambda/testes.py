from sys import argv
from math import floor

def invFrac(a, b):
    if a == 0:
        return floor(a / b)
    else:
        return floor(b / a)

def multFrac(p1, p2):
    a, b = p1
    c, d = p2
    return (a * c, b * d)

def main ():
    p1 = int(argv[1]), int(argv[2])
    p2 = int(argv[3]), int(argv[4])

    print(multFrac(p1, p2))

main()