def fibo(max):
    n, a, b = 0, 0, 1
    while n < max:
        yield b
        a, b = b, a + b
        n += 1

    return 'done'


def draw_heart():
    print('\n'.join([''.join([(' happy-birthday'[(x - y) % 15] if ((x * 0.05) ** 2 + (y * 0.1) ** 2 - 1) ** 3 - (
                x * 0.05) ** 2 * (y * 0.1) ** 3 <= 0 else ' ') for x in range(-30, 30)]) for y in range(15, -15, -1)]))


if __name__ == "__main__":
    ger = fibo(10)
    for i in ger:
        print(i)