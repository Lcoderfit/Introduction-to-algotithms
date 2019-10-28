def simple_corouting():
    print('-> corouting started')
    x = yield
    print('-> corouting received:', x)


if __name__ == "__main__":
    my_coro = simple_corouting()
    print(my_coro)
    print(next(my_coro))
    my_coro.send(42)
