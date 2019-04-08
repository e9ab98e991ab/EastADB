
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

type TSaveTextFileDialog struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewSaveTextFileDialog
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewSaveTextFileDialog(owner IComponent) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = SaveTextFileDialog_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SaveTextFileDialogFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func SaveTextFileDialogFromInst(inst uintptr) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// SaveTextFileDialogFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func SaveTextFileDialogFromObj(obj IObject) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// SaveTextFileDialogFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func SaveTextFileDialogFromUnsafePointer(ptr unsafe.Pointer) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TSaveTextFileDialog) Free() {
    if s.instance != 0 {
        SaveTextFileDialog_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TSaveTextFileDialog) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TSaveTextFileDialog) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TSaveTextFileDialog) IsValid() bool {
    return s.instance != 0
}

// TSaveTextFileDialogClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TSaveTextFileDialogClass() TClass {
    return SaveTextFileDialog_StaticClassType()
}

// Execute
// CN: 执行。
// EN: .
func (s *TSaveTextFileDialog) Execute() bool {
    return SaveTextFileDialog_Execute(s.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TSaveTextFileDialog) FindComponent(AName string) *TComponent {
    return ComponentFromInst(SaveTextFileDialog_FindComponent(s.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TSaveTextFileDialog) GetNamePath() string {
    return SaveTextFileDialog_GetNamePath(s.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TSaveTextFileDialog) HasParent() bool {
    return SaveTextFileDialog_HasParent(s.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TSaveTextFileDialog) Assign(Source IObject) {
    SaveTextFileDialog_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TSaveTextFileDialog) DisposeOf() {
    SaveTextFileDialog_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TSaveTextFileDialog) ClassType() TClass {
    return SaveTextFileDialog_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TSaveTextFileDialog) ClassName() string {
    return SaveTextFileDialog_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TSaveTextFileDialog) InstanceSize() int32 {
    return SaveTextFileDialog_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TSaveTextFileDialog) InheritsFrom(AClass TClass) bool {
    return SaveTextFileDialog_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TSaveTextFileDialog) Equals(Obj IObject) bool {
    return SaveTextFileDialog_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TSaveTextFileDialog) GetHashCode() int32 {
    return SaveTextFileDialog_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TSaveTextFileDialog) ToString() string {
    return SaveTextFileDialog_ToString(s.instance)
}

// Files
func (s *TSaveTextFileDialog) Files() *TStrings {
    return StringsFromInst(SaveTextFileDialog_GetFiles(s.instance))
}

// DefaultExt
func (s *TSaveTextFileDialog) DefaultExt() string {
    return SaveTextFileDialog_GetDefaultExt(s.instance)
}

// SetDefaultExt
func (s *TSaveTextFileDialog) SetDefaultExt(value string) {
    SaveTextFileDialog_SetDefaultExt(s.instance, value)
}

// FileName
func (s *TSaveTextFileDialog) FileName() string {
    return SaveTextFileDialog_GetFileName(s.instance)
}

// SetFileName
func (s *TSaveTextFileDialog) SetFileName(value string) {
    SaveTextFileDialog_SetFileName(s.instance, value)
}

// Filter
func (s *TSaveTextFileDialog) Filter() string {
    return SaveTextFileDialog_GetFilter(s.instance)
}

// SetFilter
func (s *TSaveTextFileDialog) SetFilter(value string) {
    SaveTextFileDialog_SetFilter(s.instance, value)
}

// FilterIndex
func (s *TSaveTextFileDialog) FilterIndex() int32 {
    return SaveTextFileDialog_GetFilterIndex(s.instance)
}

// SetFilterIndex
func (s *TSaveTextFileDialog) SetFilterIndex(value int32) {
    SaveTextFileDialog_SetFilterIndex(s.instance, value)
}

// InitialDir
func (s *TSaveTextFileDialog) InitialDir() string {
    return SaveTextFileDialog_GetInitialDir(s.instance)
}

// SetInitialDir
func (s *TSaveTextFileDialog) SetInitialDir(value string) {
    SaveTextFileDialog_SetInitialDir(s.instance, value)
}

// Options
func (s *TSaveTextFileDialog) Options() TOpenOptions {
    return SaveTextFileDialog_GetOptions(s.instance)
}

// SetOptions
func (s *TSaveTextFileDialog) SetOptions(value TOpenOptions) {
    SaveTextFileDialog_SetOptions(s.instance, value)
}

// OptionsEx
func (s *TSaveTextFileDialog) OptionsEx() TOpenOptionsEx {
    return SaveTextFileDialog_GetOptionsEx(s.instance)
}

// SetOptionsEx
func (s *TSaveTextFileDialog) SetOptionsEx(value TOpenOptionsEx) {
    SaveTextFileDialog_SetOptionsEx(s.instance, value)
}

// Title
func (s *TSaveTextFileDialog) Title() string {
    return SaveTextFileDialog_GetTitle(s.instance)
}

// SetTitle
func (s *TSaveTextFileDialog) SetTitle(value string) {
    SaveTextFileDialog_SetTitle(s.instance, value)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TSaveTextFileDialog) Handle() HWND {
    return SaveTextFileDialog_GetHandle(s.instance)
}

// SetOnClose
func (s *TSaveTextFileDialog) SetOnClose(fn TNotifyEvent) {
    SaveTextFileDialog_SetOnClose(s.instance, fn)
}

// SetOnShow
// CN: 设置显示事件。
// EN: .
func (s *TSaveTextFileDialog) SetOnShow(fn TNotifyEvent) {
    SaveTextFileDialog_SetOnShow(s.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TSaveTextFileDialog) ComponentCount() int32 {
    return SaveTextFileDialog_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TSaveTextFileDialog) ComponentIndex() int32 {
    return SaveTextFileDialog_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TSaveTextFileDialog) SetComponentIndex(value int32) {
    SaveTextFileDialog_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TSaveTextFileDialog) Owner() *TComponent {
    return ComponentFromInst(SaveTextFileDialog_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TSaveTextFileDialog) Name() string {
    return SaveTextFileDialog_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TSaveTextFileDialog) SetName(value string) {
    SaveTextFileDialog_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TSaveTextFileDialog) Tag() int {
    return SaveTextFileDialog_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TSaveTextFileDialog) SetTag(value int) {
    SaveTextFileDialog_SetTag(s.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TSaveTextFileDialog) Components(AIndex int32) *TComponent {
    return ComponentFromInst(SaveTextFileDialog_GetComponents(s.instance, AIndex))
}

