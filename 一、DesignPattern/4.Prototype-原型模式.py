"""
4、原型模式
原型模式：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。

被实例化的类

原型模式包含以下主要角色。
1.抽象原型类：规定了具体原型对象必须实现的接口。
2.具体原型类：实现抽象原型类的 clone() 方法，它是可被复制的对象。
3.访问类：使用具体原型类中的 clone() 方法来复制新的对象。


"""
# 抽象原型类
class Prototype(object):
    value = 'default'

    def clone(self, **attrs):
        obj = self.__class__()
        obj.__dict__.update(attrs)
        return obj


# 具体原型类
class PrototypeDispatcher(object):
    def __init__(self):
        self._objects = {}

    def get_objects(self):
        return self._objects

    def register_object(self, name, obj):
        self._objects[name] = obj

    def unregister_object(self, name):
        del self._objects[name]


# 访问类


# class Prototype(object):
#
#     value = 'default'
#
#     def clone(self, **attrs):
#         """Clone a prototype and update inner attributes dictionary"""
#         # Python in Practice, Mark Summerfield
#         obj = self.__class__()
#         obj.__dict__.update(attrs)
#         return obj
#
#
# class PrototypeDispatcher(object):
#     def __init__(self):
#         self._objects = {}
#
#     def get_objects(self):
#         """Get all objects"""
#         return self._objects
#
#     def register_object(self, name, obj):
#         """Register an object"""
#         self._objects[name] = obj
#
#     def unregister_object(self, name):
#         """Unregister an object"""
#         del self._objects[name]
#
#
# def main():
#     dispatcher = PrototypeDispatcher()
#     prototype = Prototype()
#
#     d = prototype.clone()
#     a = prototype.clone(value='a-value', category='a')
#     b = prototype.clone(value='b-value', is_checked=True)
#     dispatcher.register_object('objecta', a)
#     dispatcher.register_object('objectb', b)
#     dispatcher.register_object('default', d)
#     print([{n: p.value} for n, p in dispatcher.get_objects().items()])
#
#
# if __name__ == '__main__':
#     main()
#
# ### OUTPUT ###
# # [{'objectb': 'b-value'}, {'default': 'default'}, {'objecta': 'a-value'}]