
def parent_func(n):
    def decorator_func(func):
        def child_func():
            print(f" {n} before argument func")
            func()
            print("after argument func")
            return 22 
        return child_func
    return decorator_func

@parent_func(20)
def arg_func():
    print("middle of the func")

print(arg_func())


