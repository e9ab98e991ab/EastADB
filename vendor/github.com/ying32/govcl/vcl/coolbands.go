
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

type TCoolBands struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCoolBands
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCoolBands() *TCoolBands {
    c := new(TCoolBands)
    c.instance = CoolBands_Create()
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CoolBandsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CoolBandsFromInst(inst uintptr) *TCoolBands {
    c := new(TCoolBands)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CoolBandsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CoolBandsFromObj(obj IObject) *TCoolBands {
    c := new(TCoolBands)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CoolBandsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CoolBandsFromUnsafePointer(ptr unsafe.Pointer) *TCoolBands {
    c := new(TCoolBands)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCoolBands) Free() {
    if c.instance != 0 {
        CoolBands_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCoolBands) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCoolBands) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCoolBands) IsValid() bool {
    return c.instance != 0
}

// TCoolBandsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCoolBandsClass() TClass {
    return CoolBands_StaticClassType()
}

// Add
func (c *TCoolBands) Add() *TCoolBand {
    return CoolBandFromInst(CoolBands_Add(c.instance))
}

// FindBand
func (c *TCoolBands) FindBand(AControl IControl) *TCoolBand {
    return CoolBandFromInst(CoolBands_FindBand(c.instance, CheckPtr(AControl)))
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (c *TCoolBands) Owner() *TObject {
    return ObjectFromInst(CoolBands_Owner(c.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCoolBands) Assign(Source IObject) {
    CoolBands_Assign(c.instance, CheckPtr(Source))
}

// BeginUpdate
func (c *TCoolBands) BeginUpdate() {
    CoolBands_BeginUpdate(c.instance)
}

// Clear
// CN: 清除。
// EN: .
func (c *TCoolBands) Clear() {
    CoolBands_Clear(c.instance)
}

// ClearAndResetID
func (c *TCoolBands) ClearAndResetID() {
    CoolBands_ClearAndResetID(c.instance)
}

// Delete
func (c *TCoolBands) Delete(Index int32) {
    CoolBands_Delete(c.instance, Index)
}

// EndUpdate
func (c *TCoolBands) EndUpdate() {
    CoolBands_EndUpdate(c.instance)
}

// FindItemID
func (c *TCoolBands) FindItemID(ID int32) *TCollectionItem {
    return CollectionItemFromInst(CoolBands_FindItemID(c.instance, ID))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCoolBands) GetNamePath() string {
    return CoolBands_GetNamePath(c.instance)
}

// Insert
func (c *TCoolBands) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(CoolBands_Insert(c.instance, Index))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCoolBands) DisposeOf() {
    CoolBands_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCoolBands) ClassType() TClass {
    return CoolBands_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCoolBands) ClassName() string {
    return CoolBands_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCoolBands) InstanceSize() int32 {
    return CoolBands_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCoolBands) InheritsFrom(AClass TClass) bool {
    return CoolBands_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCoolBands) Equals(Obj IObject) bool {
    return CoolBands_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCoolBands) GetHashCode() int32 {
    return CoolBands_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCoolBands) ToString() string {
    return CoolBands_ToString(c.instance)
}

// CoolBar
func (c *TCoolBands) CoolBar() *TCoolBar {
    return CoolBarFromInst(CoolBands_GetCoolBar(c.instance))
}

// Capacity
func (c *TCoolBands) Capacity() int32 {
    return CoolBands_GetCapacity(c.instance)
}

// SetCapacity
func (c *TCoolBands) SetCapacity(value int32) {
    CoolBands_SetCapacity(c.instance, value)
}

// Count
func (c *TCoolBands) Count() int32 {
    return CoolBands_GetCount(c.instance)
}

// Items
func (c *TCoolBands) Items(Index int32) *TCoolBand {
    return CoolBandFromInst(CoolBands_GetItems(c.instance, Index))
}

// Items
func (c *TCoolBands) SetItems(Index int32, value *TCoolBand) {
    CoolBands_SetItems(c.instance, Index, CheckPtr(value))
}

