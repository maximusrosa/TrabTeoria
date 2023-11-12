from sys import argv

def decod(num: int) -> tuple:
    a = 0
    while num % 2 == 0:
        num = num / 2
        a += 1

    b = int((num - 1) / 2)

    return a , b

def cod(a: int, b: int) -> int:
    num = (2*b + 1) * 2**a
    return num

def foo(x):
    return 2**(x + 1) * (2 * (x // 2) + 1)



def main():
    Y = foo(int(argv[1]))
    print(f'foo({argv[1]}) = {Y}')

    a, b = decod(Y)
    print(f'decod({Y}) = ({a}, {b})')



if __name__ == '__main__':
    main()