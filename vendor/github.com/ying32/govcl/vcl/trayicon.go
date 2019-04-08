
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

type TTrayIcon struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTrayIcon
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTrayIcon(owner IComponent) *TTrayIcon {
    t := new(TTrayIcon)
    t.instance = TrayIcon_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TrayIconFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TrayIconFromInst(inst uintptr) *TTrayIcon {
    t := new(TTrayIcon)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TrayIconFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TrayIconFromObj(obj IObject) *TTrayIcon {
    t := new(TTrayIcon)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TrayIconFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TrayIconFromUnsafePointer(ptr unsafe.Pointer) *TTrayIcon {
    t := new(TTrayIcon)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTrayIcon) Free() {
    if t.instance != 0 {
        TrayIcon_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTrayIcon) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTrayIcon) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTrayIcon) IsValid() bool {
    return t.instance != 0
}

// TTrayIconClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTrayIconClass() TClass {
    return TrayIcon_StaticClassType()
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (t *TTrayIcon) Refresh() {
    TrayIcon_Refresh(t.instance)
}

// SetDefaultIcon
func (t *TTrayIcon) SetDefaultIcon() {
    TrayIcon_SetDefaultIcon(t.instance)
}

// ShowBalloonHint
func (t *TTrayIcon) ShowBalloonHint() {
    TrayIcon_ShowBalloonHint(t.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TTrayIcon) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TrayIcon_FindComponent(t.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTrayIcon) GetNamePath() string {
    return TrayIcon_GetNamePath(t.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TTrayIcon) HasParent() bool {
    return TrayIcon_HasParent(t.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTrayIcon) Assign(Source IObject) {
    TrayIcon_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTrayIcon) DisposeOf() {
    TrayIcon_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTrayIcon) ClassType() TClass {
    return TrayIcon_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTrayIcon) ClassName() string {
    return TrayIcon_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTrayIcon) InstanceSize() int32 {
    return TrayIcon_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTrayIcon) InheritsFrom(AClass TClass) bool {
    return TrayIcon_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTrayIcon) Equals(Obj IObject) bool {
    return TrayIcon_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTrayIcon) GetHashCode() int32 {
    return TrayIcon_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTrayIcon) ToString() string {
    return TrayIcon_ToString(t.instance)
}

// Animate
func (t *TTrayIcon) Animate() bool {
    return TrayIcon_GetAnimate(t.instance)
}

// SetAnimate
func (t *TTrayIcon) SetAnimate(value bool) {
    TrayIcon_SetAnimate(t.instance, value)
}

// AnimateInterval
func (t *TTrayIcon) AnimateInterval() uint32 {
    return TrayIcon_GetAnimateInterval(t.instance)
}

// SetAnimateInterval
func (t *TTrayIcon) SetAnimateInterval(value uint32) {
    TrayIcon_SetAnimateInterval(t.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (t *TTrayIcon) Hint() string {
    return TrayIcon_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (t *TTrayIcon) SetHint(value string) {
    TrayIcon_SetHint(t.instance, value)
}

// BalloonHint
func (t *TTrayIcon) BalloonHint() string {
    return TrayIcon_GetBalloonHint(t.instance)
}

// SetBalloonHint
func (t *TTrayIcon) SetBalloonHint(value string) {
    TrayIcon_SetBalloonHint(t.instance, value)
}

// BalloonTitle
func (t *TTrayIcon) BalloonTitle() string {
    return TrayIcon_GetBalloonTitle(t.instance)
}

// SetBalloonTitle
func (t *TTrayIcon) SetBalloonTitle(value string) {
    TrayIcon_SetBalloonTitle(t.instance, value)
}

// BalloonTimeout
func (t *TTrayIcon) BalloonTimeout() int32 {
    return TrayIcon_GetBalloonTimeout(t.instance)
}

// SetBalloonTimeout
func (t *TTrayIcon) SetBalloonTimeout(value int32) {
    TrayIcon_SetBalloonTimeout(t.instance, value)
}

// BalloonFlags
func (t *TTrayIcon) BalloonFlags() TBalloonFlags {
    return TrayIcon_GetBalloonFlags(t.instance)
}

// SetBalloonFlags
func (t *TTrayIcon) SetBalloonFlags(value TBalloonFlags) {
    TrayIcon_SetBalloonFlags(t.instance, value)
}

// Icon
// CN: 获取图标。
// EN: Get icon.
func (t *TTrayIcon) Icon() *TIcon {
    return IconFromInst(TrayIcon_GetIcon(t.instance))
}

// SetIcon
// CN: 设置图标。
// EN: Set icon.
func (t *TTrayIcon) SetIcon(value *TIcon) {
    TrayIcon_SetIcon(t.instance, CheckPtr(value))
}

// IconIndex
func (t *TTrayIcon) IconIndex() int32 {
    return TrayIcon_GetIconIndex(t.instance)
}

// SetIconIndex
func (t *TTrayIcon) SetIconIndex(value int32) {
    TrayIcon_SetIconIndex(t.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (t *TTrayIcon) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TrayIcon_GetPopupMenu(t.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (t *TTrayIcon) SetPopupMenu(value IComponent) {
    TrayIcon_SetPopupMenu(t.instance, CheckPtr(value))
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TTrayIcon) Visible() bool {
    return TrayIcon_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TTrayIcon) SetVisible(value bool) {
    TrayIcon_SetVisible(t.instance, value)
}

// SetOnBalloonClick
func (t *TTrayIcon) SetOnBalloonClick(fn TNotifyEvent) {
    TrayIcon_SetOnBalloonClick(t.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (t *TTrayIcon) SetOnClick(fn TNotifyEvent) {
    TrayIcon_SetOnClick(t.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (t *TTrayIcon) SetOnDblClick(fn TNotifyEvent) {
    TrayIcon_SetOnDblClick(t.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (t *TTrayIcon) SetOnMouseMove(fn TMouseMoveEvent) {
    TrayIcon_SetOnMouseMove(t.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (t *TTrayIcon) SetOnMouseUp(fn TMouseEvent) {
    TrayIcon_SetOnMouseUp(t.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (t *TTrayIcon) SetOnMouseDown(fn TMouseEvent) {
    TrayIcon_SetOnMouseDown(t.instance, fn)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTrayIcon) ComponentCount() int32 {
    return TrayIcon_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TTrayIcon) ComponentIndex() int32 {
    return TrayIcon_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TTrayIcon) SetComponentIndex(value int32) {
    TrayIcon_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTrayIcon) Owner() *TComponent {
    return ComponentFromInst(TrayIcon_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTrayIcon) Name() string {
    return TrayIcon_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTrayIcon) SetName(value string) {
    TrayIcon_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTrayIcon) Tag() int {
    return TrayIcon_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTrayIcon) SetTag(value int) {
    TrayIcon_SetTag(t.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTrayIcon) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TrayIcon_GetComponents(t.instance, AIndex))
}

