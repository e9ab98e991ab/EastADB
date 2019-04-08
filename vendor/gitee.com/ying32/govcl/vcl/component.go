
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

type TComponent struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewComponent
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewComponent(owner IComponent) *TComponent {
    c := new(TComponent)
    c.instance = Component_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ComponentFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ComponentFromInst(inst uintptr) *TComponent {
    c := new(TComponent)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ComponentFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ComponentFromObj(obj IObject) *TComponent {
    c := new(TComponent)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ComponentFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ComponentFromUnsafePointer(ptr unsafe.Pointer) *TComponent {
    c := new(TComponent)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TComponent) Free() {
    if c.instance != 0 {
        Component_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TComponent) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TComponent) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TComponent) IsValid() bool {
    return c.instance != 0
}

// TComponentClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TComponentClass() TClass {
    return Component_StaticClassType()
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TComponent) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Component_FindComponent(c.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TComponent) GetNamePath() string {
    return Component_GetNamePath(c.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TComponent) HasParent() bool {
    return Component_HasParent(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TComponent) Assign(Source IObject) {
    Component_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TComponent) DisposeOf() {
    Component_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TComponent) ClassType() TClass {
    return Component_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TComponent) ClassName() string {
    return Component_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TComponent) InstanceSize() int32 {
    return Component_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TComponent) InheritsFrom(AClass TClass) bool {
    return Component_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TComponent) Equals(Obj IObject) bool {
    return Component_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TComponent) GetHashCode() int32 {
    return Component_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TComponent) ToString() string {
    return Component_ToString(c.instance)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TComponent) ComponentCount() int32 {
    return Component_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TComponent) ComponentIndex() int32 {
    return Component_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TComponent) SetComponentIndex(value int32) {
    Component_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TComponent) Owner() *TComponent {
    return ComponentFromInst(Component_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TComponent) Name() string {
    return Component_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TComponent) SetName(value string) {
    Component_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TComponent) Tag() int {
    return Component_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TComponent) SetTag(value int) {
    Component_SetTag(c.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TComponent) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Component_GetComponents(c.instance, AIndex))
}

