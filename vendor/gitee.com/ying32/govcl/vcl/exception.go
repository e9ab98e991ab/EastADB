
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type Exception struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// ExceptionFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ExceptionFromInst(inst uintptr) *Exception {
    e := new(Exception)
    e.instance = inst
    e.ptr = unsafe.Pointer(inst)
    return e
}

// ExceptionFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ExceptionFromObj(obj IObject) *Exception {
    e := new(Exception)
    e.instance = CheckPtr(obj)
    e.ptr = unsafe.Pointer(e.instance)
    return e
}

// ExceptionFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ExceptionFromUnsafePointer(ptr unsafe.Pointer) *Exception {
    e := new(Exception)
    e.instance = uintptr(ptr)
    e.ptr = ptr
    return e
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (e *Exception) Instance() uintptr {
    return e.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (e *Exception) UnsafeAddr() unsafe.Pointer {
    return e.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (e *Exception) IsValid() bool {
    return e.instance != 0
}

// ExceptionClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func ExceptionClass() TClass {
    return Exception_StaticClassType()
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (e *Exception) ToString() string {
    return Exception_ToString(e.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (e *Exception) DisposeOf() {
    Exception_DisposeOf(e.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (e *Exception) ClassType() TClass {
    return Exception_ClassType(e.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (e *Exception) ClassName() string {
    return Exception_ClassName(e.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (e *Exception) InstanceSize() int32 {
    return Exception_InstanceSize(e.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (e *Exception) InheritsFrom(AClass TClass) bool {
    return Exception_InheritsFrom(e.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (e *Exception) Equals(Obj IObject) bool {
    return Exception_Equals(e.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (e *Exception) GetHashCode() int32 {
    return Exception_GetHashCode(e.instance)
}

// BaseException
func (e *Exception) BaseException() *Exception {
    return ExceptionFromInst(Exception_GetBaseException(e.instance))
}

// InnerException
func (e *Exception) InnerException() *Exception {
    return ExceptionFromInst(Exception_GetInnerException(e.instance))
}

// Message
func (e *Exception) Message() string {
    return Exception_GetMessage(e.instance)
}

// SetMessage
func (e *Exception) SetMessage(value string) {
    Exception_SetMessage(e.instance, value)
}

// StackTrace
func (e *Exception) StackTrace() string {
    return Exception_GetStackTrace(e.instance)
}

// StackInfo
func (e *Exception) StackInfo() uintptr {
    return Exception_GetStackInfo(e.instance)
}

