package main

import (
    "fmt"
    "reflect"
)

// 定义一个示例结构体
type Person struct {
    Name string `json:"name" tag:"example"`
    Age  int    `json:"age"`
}

// 结构体方法
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

func main() {
    // 1. 基本类型反射
    var num float64 = 3.14
    fmt.Println("=== 基本类型反射 ===")
    fmt.Println("Type:", reflect.TypeOf(num))
    fmt.Println("Value:", reflect.ValueOf(num))
    fmt.Println()

    // 2. 结构体反射
    p := Person{"Alice", 30}
    t := reflect.TypeOf(p)
    v := reflect.ValueOf(p)

    fmt.Println("=== 结构体反射 ===")
    fmt.Println("Type:", t)
    fmt.Println("Kind:", t.Kind())
    fmt.Println("Value:", v)
    fmt.Println()

    // 3. 遍历结构体字段
    fmt.Println("=== 结构体字段 ===")
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        fmt.Printf("%d. %s (%s) = %v, tag:%v\n",
            i, field.Name, field.Type, value, field.Tag)
    }
    fmt.Println()

    // 4. 调用结构体方法
    fmt.Println("=== 方法调用 ===")
    method := v.MethodByName("Greet")
    if method.IsValid() {
        results := method.Call(nil)
        fmt.Println("Method result:", results[0])
    } else {
        fmt.Println("Method Greet not found")
    }
    fmt.Println()

    // 5. 修改值反射
    fmt.Println("=== 修改值 ===")
    // 注意: 要修改值必须获取指针的Value
    pv := reflect.ValueOf(&p).Elem()
    nameField := pv.FieldByName("Name")
    if nameField.CanSet() {
        nameField.SetString("Bob")
        fmt.Println("Modified name:", p.Name)
    }
    fmt.Println()

    // 6. 创建新实例
    fmt.Println("=== 创建实例 ===")
    newPersonValue := reflect.New(t)                  // This returns a Value representing a pointer to the new instance
    newPerson := newPersonValue.Interface().(*Person) // Proper type assertion
    fmt.Printf("New instance: %#v\n", *newPerson)     // Dereference the pointer to print the actual struct
}
