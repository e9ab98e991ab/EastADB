
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

type TIconOptions struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// IconOptionsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func IconOptionsFromInst(inst uintptr) *TIconOptions {
    i := new(TIconOptions)
    i.instance = inst
    i.ptr = unsafe.Pointer(inst)
    return i
}

// IconOptionsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func IconOptionsFromObj(obj IObject) *TIconOptions {
    i := new(TIconOptions)
    i.instance = CheckPtr(obj)
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// IconOptionsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func IconOptionsFromUnsafePointer(ptr unsafe.Pointer) *TIconOptions {
    i := new(TIconOptions)
    i.instance = uintptr(ptr)
    i.ptr = ptr
    return i
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (i *TIconOptions) Instance() uintptr {
    return i.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (i *TIconOptions) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (i *TIconOptions) IsValid() bool {
    return i.instance != 0
}

// TIconOptionsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TIconOptionsClass() TClass {
    return IconOptions_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (i *TIconOptions) Assign(Source IObject) {
    IconOptions_Assign(i.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (i *TIconOptions) GetNamePath() string {
    return IconOptions_GetNamePath(i.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (i *TIconOptions) DisposeOf() {
    IconOptions_DisposeOf(i.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (i *TIconOptions) ClassType() TClass {
    return IconOptions_ClassType(i.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (i *TIconOptions) ClassName() string {
    return IconOptions_ClassName(i.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (i *TIconOptions) InstanceSize() int32 {
    return IconOptions_InstanceSize(i.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (i *TIconOptions) InheritsFrom(AClass TClass) bool {
    return IconOptions_InheritsFrom(i.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (i *TIconOptions) Equals(Obj IObject) bool {
    return IconOptions_Equals(i.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (i *TIconOptions) GetHashCode() int32 {
    return IconOptions_GetHashCode(i.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (i *TIconOptions) ToString() string {
    return IconOptions_ToString(i.instance)
}

// Arrangement
func (i *TIconOptions) Arrangement() TIconArrangement {
    return IconOptions_GetArrangement(i.instance)
}

// SetArrangement
func (i *TIconOptions) SetArrangement(value TIconArrangement) {
    IconOptions_SetArrangement(i.instance, value)
}

// AutoArrange
func (i *TIconOptions) AutoArrange() bool {
    return IconOptions_GetAutoArrange(i.instance)
}

// SetAutoArrange
func (i *TIconOptions) SetAutoArrange(value bool) {
    IconOptions_SetAutoArrange(i.instance, value)
}

