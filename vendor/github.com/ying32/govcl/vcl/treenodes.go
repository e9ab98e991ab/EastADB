
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

type TTreeNodes struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTreeNodes
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTreeNodes() *TTreeNodes {
    t := new(TTreeNodes)
    t.instance = TreeNodes_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TreeNodesFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TreeNodesFromInst(inst uintptr) *TTreeNodes {
    t := new(TTreeNodes)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TreeNodesFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TreeNodesFromObj(obj IObject) *TTreeNodes {
    t := new(TTreeNodes)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TreeNodesFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TreeNodesFromUnsafePointer(ptr unsafe.Pointer) *TTreeNodes {
    t := new(TTreeNodes)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTreeNodes) Free() {
    if t.instance != 0 {
        TreeNodes_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTreeNodes) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTreeNodes) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTreeNodes) IsValid() bool {
    return t.instance != 0
}

// TTreeNodesClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTreeNodesClass() TClass {
    return TreeNodes_StaticClassType()
}

// AddChildFirst
func (t *TTreeNodes) AddChildFirst(Parent *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddChildFirst(t.instance, CheckPtr(Parent), S))
}

// AddChild
func (t *TTreeNodes) AddChild(Parent *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddChild(t.instance, CheckPtr(Parent), S))
}

// AddChildObjectFirst
func (t *TTreeNodes) AddChildObjectFirst(Parent *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddChildObjectFirst(t.instance, CheckPtr(Parent), S , Ptr))
}

// AddChildObject
func (t *TTreeNodes) AddChildObject(Parent *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddChildObject(t.instance, CheckPtr(Parent), S , Ptr))
}

// AddObjectFirst
func (t *TTreeNodes) AddObjectFirst(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddObjectFirst(t.instance, CheckPtr(Sibling), S , Ptr))
}

// AddObject
func (t *TTreeNodes) AddObject(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddObject(t.instance, CheckPtr(Sibling), S , Ptr))
}

// AddNode
func (t *TTreeNodes) AddNode(Node *TTreeNode, Relative *TTreeNode, S string, Ptr uintptr, Method TNodeAttachMode) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddNode(t.instance, CheckPtr(Node), CheckPtr(Relative), S , Ptr , Method))
}

// AddFirst
func (t *TTreeNodes) AddFirst(Sibling *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_AddFirst(t.instance, CheckPtr(Sibling), S))
}

// Add
func (t *TTreeNodes) Add(Sibling *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_Add(t.instance, CheckPtr(Sibling), S))
}

// AlphaSort
func (t *TTreeNodes) AlphaSort(ARecurse bool) bool {
    return TreeNodes_AlphaSort(t.instance, ARecurse)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTreeNodes) Assign(Source IObject) {
    TreeNodes_Assign(t.instance, CheckPtr(Source))
}

// BeginUpdate
func (t *TTreeNodes) BeginUpdate() {
    TreeNodes_BeginUpdate(t.instance)
}

// Clear
// CN: 清除。
// EN: .
func (t *TTreeNodes) Clear() {
    TreeNodes_Clear(t.instance)
}

// Delete
func (t *TTreeNodes) Delete(Node *TTreeNode) {
    TreeNodes_Delete(t.instance, CheckPtr(Node))
}

// EndUpdate
func (t *TTreeNodes) EndUpdate() {
    TreeNodes_EndUpdate(t.instance)
}

// GetFirstNode
func (t *TTreeNodes) GetFirstNode() *TTreeNode {
    return TreeNodeFromInst(TreeNodes_GetFirstNode(t.instance))
}

// GetNode
func (t *TTreeNodes) GetNode(ItemId uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_GetNode(t.instance, ItemId))
}

// Insert
func (t *TTreeNodes) Insert(Sibling *TTreeNode, S string) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_Insert(t.instance, CheckPtr(Sibling), S))
}

// InsertObject
func (t *TTreeNodes) InsertObject(Sibling *TTreeNode, S string, Ptr uintptr) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_InsertObject(t.instance, CheckPtr(Sibling), S , Ptr))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTreeNodes) GetNamePath() string {
    return TreeNodes_GetNamePath(t.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTreeNodes) DisposeOf() {
    TreeNodes_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTreeNodes) ClassType() TClass {
    return TreeNodes_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTreeNodes) ClassName() string {
    return TreeNodes_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTreeNodes) InstanceSize() int32 {
    return TreeNodes_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTreeNodes) InheritsFrom(AClass TClass) bool {
    return TreeNodes_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTreeNodes) Equals(Obj IObject) bool {
    return TreeNodes_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTreeNodes) GetHashCode() int32 {
    return TreeNodes_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTreeNodes) ToString() string {
    return TreeNodes_ToString(t.instance)
}

// Count
func (t *TTreeNodes) Count() int32 {
    return TreeNodes_GetCount(t.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (t *TTreeNodes) Handle() HWND {
    return TreeNodes_GetHandle(t.instance)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTreeNodes) Owner() *TWinControl {
    return WinControlFromInst(TreeNodes_GetOwner(t.instance))
}

// Item
func (t *TTreeNodes) Item(Index int32) *TTreeNode {
    return TreeNodeFromInst(TreeNodes_GetItem(t.instance, Index))
}

