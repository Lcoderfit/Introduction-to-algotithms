"""
1、抽象工厂模式
抽象工厂模式(Abstract Factory Pattern):提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们的类
"""
import random


class PetShop:
    """宠物商店"""
    def __init__(self, animal_factory=None):
        """定义一个宠物工厂作为抽象工厂"""
        self.pet_factory = animal_factory

    def show_pet(self):
        """使用抽象工厂创建并显示一个宠物"""
        pet = self.pet_factory.get_pet()
        print("我们有一个可爱的 {}".format(pet))
        print("它说{}".format(pet.speak()))
        print("我们还有{}".format(self.pet_factory.get_food()))


# 工厂生产的事物
class Dog:
    def speak(self):
        return "汪"

    def __str__(self):
        return "Dog"


class Cat:
    def speak(self):
        return "喵"

    def __str__(self):
        return "Cat"


# 工厂类
class DogFactory:
    def get_pet(self):
        return Dog()

    def get_food(self):
        return "狗食"


class CatFactory:
    def get_pet(self):
        return Cat()

    def get_food(self):
        return "猫食"


def get_factory():
    return random.choice([DogFactory, CatFactory])()


if __name__ == "__main__":
    for i in range(4):
        shop = PetShop(get_factory())
        shop.show_pet()
        print("=" * 20)
