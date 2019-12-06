import time


if __name__ == "__main__":
    string = "—\|/"
    n = 45
    while n > 0 :
        for i in range(0, len(string)):
            print(string[i], end="")
            time.sleep(0.3)
            print("\b \b", end="")
