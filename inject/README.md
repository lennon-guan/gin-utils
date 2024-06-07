# inject

[中文版](README_cn.md)

inject is a simple dependency injection tool for gin.HandlerFunc. By registering specific types of injection functions using the inject.AddInjector function before registering the HandlerFunc, you can inject the desired values through the parameters of the HandlerFunc when binding routes.


Additionally, if the injected type implements io.Closer or has its own .Close() method without a return value, the Close() method will be automatically executed when the HandlerFunc exits.


For specific usage, you can refer to the code in the examples directory.
