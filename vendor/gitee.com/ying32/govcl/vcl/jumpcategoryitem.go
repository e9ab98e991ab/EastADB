
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

type TJumpCategoryItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewJumpCategoryItem
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewJumpCategoryItem() *TJumpCategoryItem {
    j := new(TJumpCategoryItem)
    j.instance = JumpCategoryItem_Create()
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JumpCategoryItemFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func JumpCategoryItemFromInst(inst uintptr) *TJumpCategoryItem {
    j := new(TJumpCategoryItem)
    j.instance = inst
    j.ptr = unsafe.Pointer(inst)
    return j
}

// JumpCategoryItemFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func JumpCategoryItemFromObj(obj IObject) *TJumpCategoryItem {
    j := new(TJumpCategoryItem)
    j.instance = CheckPtr(obj)
    j.ptr = unsafe.Pointer(j.instance)
    return j
}

// JumpCategoryItemFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func JumpCategoryItemFromUnsafePointer(ptr unsafe.Pointer) *TJumpCategoryItem {
    j := new(TJumpCategoryItem)
    j.instance = uintptr(ptr)
    j.ptr = ptr
    return j
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (j *TJumpCategoryItem) Free() {
    if j.instance != 0 {
        JumpCategoryItem_Free(j.instance)
        j.instance = 0
        j.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (j *TJumpCategoryItem) Instance() uintptr {
    return j.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (j *TJumpCategoryItem) UnsafeAddr() unsafe.Pointer {
    return j.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (j *TJumpCategoryItem) IsValid() bool {
    return j.instance != 0
}

// TJumpCategoryItemClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TJumpCategoryItemClass() TClass {
    return JumpCategoryItem_StaticClassType()
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (j *TJumpCategoryItem) GetNamePath() string {
    return JumpCategoryItem_GetNamePath(j.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (j *TJumpCategoryItem) Assign(Source IObject) {
    JumpCategoryItem_Assign(j.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (j *TJumpCategoryItem) DisposeOf() {
    JumpCategoryItem_DisposeOf(j.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (j *TJumpCategoryItem) ClassType() TClass {
    return JumpCategoryItem_ClassType(j.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (j *TJumpCategoryItem) ClassName() string {
    return JumpCategoryItem_ClassName(j.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (j *TJumpCategoryItem) InstanceSize() int32 {
    return JumpCategoryItem_InstanceSize(j.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (j *TJumpCategoryItem) InheritsFrom(AClass TClass) bool {
    return JumpCategoryItem_InheritsFrom(j.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (j *TJumpCategoryItem) Equals(Obj IObject) bool {
    return JumpCategoryItem_Equals(j.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (j *TJumpCategoryItem) GetHashCode() int32 {
    return JumpCategoryItem_GetHashCode(j.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (j *TJumpCategoryItem) ToString() string {
    return JumpCategoryItem_ToString(j.instance)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (j *TJumpCategoryItem) Visible() bool {
    return JumpCategoryItem_GetVisible(j.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (j *TJumpCategoryItem) SetVisible(value bool) {
    JumpCategoryItem_SetVisible(j.instance, value)
}

// CategoryName
func (j *TJumpCategoryItem) CategoryName() string {
    return JumpCategoryItem_GetCategoryName(j.instance)
}

// SetCategoryName
func (j *TJumpCategoryItem) SetCategoryName(value string) {
    JumpCategoryItem_SetCategoryName(j.instance, value)
}

// Items
func (j *TJumpCategoryItem) Items() *TJumpListCollection {
    return JumpListCollectionFromInst(JumpCategoryItem_GetItems(j.instance))
}

// SetItems
func (j *TJumpCategoryItem) SetItems(value *TJumpListCollection) {
    JumpCategoryItem_SetItems(j.instance, CheckPtr(value))
}

// Collection
func (j *TJumpCategoryItem) Collection() *TCollection {
    return CollectionFromInst(JumpCategoryItem_GetCollection(j.instance))
}

// SetCollection
func (j *TJumpCategoryItem) SetCollection(value *TCollection) {
    JumpCategoryItem_SetCollection(j.instance, CheckPtr(value))
}

// Index
func (j *TJumpCategoryItem) Index() int32 {
    return JumpCategoryItem_GetIndex(j.instance)
}

// SetIndex
func (j *TJumpCategoryItem) SetIndex(value int32) {
    JumpCategoryItem_SetIndex(j.instance, value)
}

