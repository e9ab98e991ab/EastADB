
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

type TPopupMenu struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPopupMenu
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPopupMenu(owner IComponent) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = PopupMenu_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PopupMenuFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PopupMenuFromInst(inst uintptr) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PopupMenuFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PopupMenuFromObj(obj IObject) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PopupMenuFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PopupMenuFromUnsafePointer(ptr unsafe.Pointer) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPopupMenu) Free() {
    if p.instance != 0 {
        PopupMenu_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPopupMenu) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPopupMenu) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPopupMenu) IsValid() bool {
    return p.instance != 0
}

// TPopupMenuClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPopupMenuClass() TClass {
    return PopupMenu_StaticClassType()
}

// CloseMenu
func (p *TPopupMenu) CloseMenu() {
    PopupMenu_CloseMenu(p.instance)
}

// Popup
func (p *TPopupMenu) Popup(X int32, Y int32) {
    PopupMenu_Popup(p.instance, X , Y)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPopupMenu) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PopupMenu_FindComponent(p.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPopupMenu) GetNamePath() string {
    return PopupMenu_GetNamePath(p.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPopupMenu) HasParent() bool {
    return PopupMenu_HasParent(p.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPopupMenu) Assign(Source IObject) {
    PopupMenu_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPopupMenu) DisposeOf() {
    PopupMenu_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPopupMenu) ClassType() TClass {
    return PopupMenu_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPopupMenu) ClassName() string {
    return PopupMenu_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPopupMenu) InstanceSize() int32 {
    return PopupMenu_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPopupMenu) InheritsFrom(AClass TClass) bool {
    return PopupMenu_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPopupMenu) Equals(Obj IObject) bool {
    return PopupMenu_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPopupMenu) GetHashCode() int32 {
    return PopupMenu_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPopupMenu) ToString() string {
    return PopupMenu_ToString(p.instance)
}

// PopupComponent
func (p *TPopupMenu) PopupComponent() *TComponent {
    return ComponentFromInst(PopupMenu_GetPopupComponent(p.instance))
}

// SetPopupComponent
func (p *TPopupMenu) SetPopupComponent(value IComponent) {
    PopupMenu_SetPopupComponent(p.instance, CheckPtr(value))
}

// PopupPoint
func (p *TPopupMenu) PopupPoint() TPoint {
    return PopupMenu_GetPopupPoint(p.instance)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (p *TPopupMenu) Alignment() TPopupAlignment {
    return PopupMenu_GetAlignment(p.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (p *TPopupMenu) SetAlignment(value TPopupAlignment) {
    PopupMenu_SetAlignment(p.instance, value)
}

// AutoHotkeys
func (p *TPopupMenu) AutoHotkeys() TMenuAutoFlag {
    return PopupMenu_GetAutoHotkeys(p.instance)
}

// SetAutoHotkeys
func (p *TPopupMenu) SetAutoHotkeys(value TMenuAutoFlag) {
    PopupMenu_SetAutoHotkeys(p.instance, value)
}

// BiDiMode
func (p *TPopupMenu) BiDiMode() TBiDiMode {
    return PopupMenu_GetBiDiMode(p.instance)
}

// SetBiDiMode
func (p *TPopupMenu) SetBiDiMode(value TBiDiMode) {
    PopupMenu_SetBiDiMode(p.instance, value)
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (p *TPopupMenu) Images() *TImageList {
    return ImageListFromInst(PopupMenu_GetImages(p.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (p *TPopupMenu) SetImages(value IComponent) {
    PopupMenu_SetImages(p.instance, CheckPtr(value))
}

// OwnerDraw
func (p *TPopupMenu) OwnerDraw() bool {
    return PopupMenu_GetOwnerDraw(p.instance)
}

// SetOwnerDraw
func (p *TPopupMenu) SetOwnerDraw(value bool) {
    PopupMenu_SetOwnerDraw(p.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (p *TPopupMenu) SetOnChange(fn TMenuChangeEvent) {
    PopupMenu_SetOnChange(p.instance, fn)
}

// SetOnPopup
func (p *TPopupMenu) SetOnPopup(fn TNotifyEvent) {
    PopupMenu_SetOnPopup(p.instance, fn)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPopupMenu) Handle() HMENU {
    return PopupMenu_GetHandle(p.instance)
}

// WindowHandle
func (p *TPopupMenu) WindowHandle() HWND {
    return PopupMenu_GetWindowHandle(p.instance)
}

// SetWindowHandle
func (p *TPopupMenu) SetWindowHandle(value HWND) {
    PopupMenu_SetWindowHandle(p.instance, value)
}

// Items
func (p *TPopupMenu) Items() *TMenuItem {
    return MenuItemFromInst(PopupMenu_GetItems(p.instance))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPopupMenu) ComponentCount() int32 {
    return PopupMenu_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPopupMenu) ComponentIndex() int32 {
    return PopupMenu_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPopupMenu) SetComponentIndex(value int32) {
    PopupMenu_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPopupMenu) Owner() *TComponent {
    return ComponentFromInst(PopupMenu_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPopupMenu) Name() string {
    return PopupMenu_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPopupMenu) SetName(value string) {
    PopupMenu_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPopupMenu) Tag() int {
    return PopupMenu_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPopupMenu) SetTag(value int) {
    PopupMenu_SetTag(p.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPopupMenu) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PopupMenu_GetComponents(p.instance, AIndex))
}

