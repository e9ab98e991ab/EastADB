
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

type TParaAttributes struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// ParaAttributesFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ParaAttributesFromInst(inst uintptr) *TParaAttributes {
    p := new(TParaAttributes)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// ParaAttributesFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ParaAttributesFromObj(obj IObject) *TParaAttributes {
    p := new(TParaAttributes)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// ParaAttributesFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ParaAttributesFromUnsafePointer(ptr unsafe.Pointer) *TParaAttributes {
    p := new(TParaAttributes)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TParaAttributes) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TParaAttributes) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TParaAttributes) IsValid() bool {
    return p.instance != 0
}

// TParaAttributesClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TParaAttributesClass() TClass {
    return ParaAttributes_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TParaAttributes) Assign(Source IObject) {
    ParaAttributes_Assign(p.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TParaAttributes) GetNamePath() string {
    return ParaAttributes_GetNamePath(p.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TParaAttributes) DisposeOf() {
    ParaAttributes_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TParaAttributes) ClassType() TClass {
    return ParaAttributes_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TParaAttributes) ClassName() string {
    return ParaAttributes_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TParaAttributes) InstanceSize() int32 {
    return ParaAttributes_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TParaAttributes) InheritsFrom(AClass TClass) bool {
    return ParaAttributes_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TParaAttributes) Equals(Obj IObject) bool {
    return ParaAttributes_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TParaAttributes) GetHashCode() int32 {
    return ParaAttributes_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TParaAttributes) ToString() string {
    return ParaAttributes_ToString(p.instance)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (p *TParaAttributes) Alignment() TAlignment {
    return ParaAttributes_GetAlignment(p.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (p *TParaAttributes) SetAlignment(value TAlignment) {
    ParaAttributes_SetAlignment(p.instance, value)
}

// FirstIndent
func (p *TParaAttributes) FirstIndent() int32 {
    return ParaAttributes_GetFirstIndent(p.instance)
}

// SetFirstIndent
func (p *TParaAttributes) SetFirstIndent(value int32) {
    ParaAttributes_SetFirstIndent(p.instance, value)
}

// LeftIndent
func (p *TParaAttributes) LeftIndent() int32 {
    return ParaAttributes_GetLeftIndent(p.instance)
}

// SetLeftIndent
func (p *TParaAttributes) SetLeftIndent(value int32) {
    ParaAttributes_SetLeftIndent(p.instance, value)
}

// Numbering
func (p *TParaAttributes) Numbering() TNumberingStyle {
    return ParaAttributes_GetNumbering(p.instance)
}

// SetNumbering
func (p *TParaAttributes) SetNumbering(value TNumberingStyle) {
    ParaAttributes_SetNumbering(p.instance, value)
}

// RightIndent
func (p *TParaAttributes) RightIndent() int32 {
    return ParaAttributes_GetRightIndent(p.instance)
}

// SetRightIndent
func (p *TParaAttributes) SetRightIndent(value int32) {
    ParaAttributes_SetRightIndent(p.instance, value)
}

// TabCount
func (p *TParaAttributes) TabCount() int32 {
    return ParaAttributes_GetTabCount(p.instance)
}

// SetTabCount
func (p *TParaAttributes) SetTabCount(value int32) {
    ParaAttributes_SetTabCount(p.instance, value)
}

// Tab
func (p *TParaAttributes) Tab(Index uint8) int32 {
    return ParaAttributes_GetTab(p.instance, Index)
}

// Tab
func (p *TParaAttributes) SetTab(Index uint8, value int32) {
    ParaAttributes_SetTab(p.instance, Index, value)
}

