
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

type TCollectionItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCollectionItem
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCollectionItem() *TCollectionItem {
    c := new(TCollectionItem)
    c.instance = CollectionItem_Create()
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CollectionItemFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CollectionItemFromInst(inst uintptr) *TCollectionItem {
    c := new(TCollectionItem)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CollectionItemFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CollectionItemFromObj(obj IObject) *TCollectionItem {
    c := new(TCollectionItem)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CollectionItemFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CollectionItemFromUnsafePointer(ptr unsafe.Pointer) *TCollectionItem {
    c := new(TCollectionItem)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCollectionItem) Free() {
    if c.instance != 0 {
        CollectionItem_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCollectionItem) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCollectionItem) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCollectionItem) IsValid() bool {
    return c.instance != 0
}

// TCollectionItemClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCollectionItemClass() TClass {
    return CollectionItem_StaticClassType()
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCollectionItem) GetNamePath() string {
    return CollectionItem_GetNamePath(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCollectionItem) Assign(Source IObject) {
    CollectionItem_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCollectionItem) DisposeOf() {
    CollectionItem_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCollectionItem) ClassType() TClass {
    return CollectionItem_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCollectionItem) ClassName() string {
    return CollectionItem_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCollectionItem) InstanceSize() int32 {
    return CollectionItem_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCollectionItem) InheritsFrom(AClass TClass) bool {
    return CollectionItem_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCollectionItem) Equals(Obj IObject) bool {
    return CollectionItem_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCollectionItem) GetHashCode() int32 {
    return CollectionItem_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCollectionItem) ToString() string {
    return CollectionItem_ToString(c.instance)
}

// Collection
func (c *TCollectionItem) Collection() *TCollection {
    return CollectionFromInst(CollectionItem_GetCollection(c.instance))
}

// SetCollection
func (c *TCollectionItem) SetCollection(value *TCollection) {
    CollectionItem_SetCollection(c.instance, CheckPtr(value))
}

// Index
func (c *TCollectionItem) Index() int32 {
    return CollectionItem_GetIndex(c.instance)
}

// SetIndex
func (c *TCollectionItem) SetIndex(value int32) {
    CollectionItem_SetIndex(c.instance, value)
}

