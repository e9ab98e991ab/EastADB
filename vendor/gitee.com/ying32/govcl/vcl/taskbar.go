
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

type TTaskbar struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTaskbar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTaskbar(owner IComponent) *TTaskbar {
    t := new(TTaskbar)
    t.instance = Taskbar_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TaskbarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TaskbarFromInst(inst uintptr) *TTaskbar {
    t := new(TTaskbar)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TaskbarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TaskbarFromObj(obj IObject) *TTaskbar {
    t := new(TTaskbar)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TaskbarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TaskbarFromUnsafePointer(ptr unsafe.Pointer) *TTaskbar {
    t := new(TTaskbar)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTaskbar) Free() {
    if t.instance != 0 {
        Taskbar_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskbar) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskbar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskbar) IsValid() bool {
    return t.instance != 0
}

// TTaskbarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskbarClass() TClass {
    return Taskbar_StaticClassType()
}

// DoThumbButtonNotify
func (t *TTaskbar) DoThumbButtonNotify(ItemID uint16) {
    Taskbar_DoThumbButtonNotify(t.instance, ItemID)
}

// DoThumbPreviewRequest
func (t *TTaskbar) DoThumbPreviewRequest(APreviewHeight uint16, APreviewWidth uint16) {
    Taskbar_DoThumbPreviewRequest(t.instance, APreviewHeight , APreviewWidth)
}

// DoWindowPreviewRequest
func (t *TTaskbar) DoWindowPreviewRequest() {
    Taskbar_DoWindowPreviewRequest(t.instance)
}

// UnregisterTab
func (t *TTaskbar) UnregisterTab() {
    Taskbar_UnregisterTab(t.instance)
}

// ApplyButtonsChanges
func (t *TTaskbar) ApplyButtonsChanges() {
    Taskbar_ApplyButtonsChanges(t.instance)
}

// GetMainWindowHwnd
func (t *TTaskbar) GetMainWindowHwnd() HWND {
    return Taskbar_GetMainWindowHwnd(t.instance)
}

// GetOverlayHIcon
func (t *TTaskbar) GetOverlayHIcon() HICON {
    return Taskbar_GetOverlayHIcon(t.instance)
}

// ActivateTab
func (t *TTaskbar) ActivateTab() bool {
    return Taskbar_ActivateTab(t.instance)
}

// InvalidateThumbPreview
func (t *TTaskbar) InvalidateThumbPreview() {
    Taskbar_InvalidateThumbPreview(t.instance)
}

// UpdateTab
func (t *TTaskbar) UpdateTab() {
    Taskbar_UpdateTab(t.instance)
}

// CheckApplyChanges
func (t *TTaskbar) CheckApplyChanges() {
    Taskbar_CheckApplyChanges(t.instance)
}

// ApplyChanges
func (t *TTaskbar) ApplyChanges() {
    Taskbar_ApplyChanges(t.instance)
}

// ApplyOverlayChanges
func (t *TTaskbar) ApplyOverlayChanges() {
    Taskbar_ApplyOverlayChanges(t.instance)
}

// ApplyProgressChanges
func (t *TTaskbar) ApplyProgressChanges() {
    Taskbar_ApplyProgressChanges(t.instance)
}

// ApplyTabsChanges
func (t *TTaskbar) ApplyTabsChanges() {
    Taskbar_ApplyTabsChanges(t.instance)
}

// ApplyClipAreaChanges
func (t *TTaskbar) ApplyClipAreaChanges() {
    Taskbar_ApplyClipAreaChanges(t.instance)
}

// ClearClipArea
func (t *TTaskbar) ClearClipArea() {
    Taskbar_ClearClipArea(t.instance)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TTaskbar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Taskbar_FindComponent(t.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskbar) GetNamePath() string {
    return Taskbar_GetNamePath(t.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TTaskbar) HasParent() bool {
    return Taskbar_HasParent(t.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskbar) Assign(Source IObject) {
    Taskbar_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTaskbar) DisposeOf() {
    Taskbar_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskbar) ClassType() TClass {
    return Taskbar_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskbar) ClassName() string {
    return Taskbar_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskbar) InstanceSize() int32 {
    return Taskbar_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskbar) InheritsFrom(AClass TClass) bool {
    return Taskbar_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskbar) Equals(Obj IObject) bool {
    return Taskbar_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskbar) GetHashCode() int32 {
    return Taskbar_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTaskbar) ToString() string {
    return Taskbar_ToString(t.instance)
}

// TaskBarButtons
func (t *TTaskbar) TaskBarButtons() *TThumbBarButtonList {
    return ThumbBarButtonListFromInst(Taskbar_GetTaskBarButtons(t.instance))
}

// SetTaskBarButtons
func (t *TTaskbar) SetTaskBarButtons(value *TThumbBarButtonList) {
    Taskbar_SetTaskBarButtons(t.instance, CheckPtr(value))
}

// ProgressState
func (t *TTaskbar) ProgressState() TTaskBarProgressState {
    return Taskbar_GetProgressState(t.instance)
}

// SetProgressState
func (t *TTaskbar) SetProgressState(value TTaskBarProgressState) {
    Taskbar_SetProgressState(t.instance, value)
}

// ProgressMaxValue
func (t *TTaskbar) ProgressMaxValue() int64 {
    return Taskbar_GetProgressMaxValue(t.instance)
}

// SetProgressMaxValue
func (t *TTaskbar) SetProgressMaxValue(value int64) {
    Taskbar_SetProgressMaxValue(t.instance, value)
}

// ProgressValue
func (t *TTaskbar) ProgressValue() int64 {
    return Taskbar_GetProgressValue(t.instance)
}

// SetProgressValue
func (t *TTaskbar) SetProgressValue(value int64) {
    Taskbar_SetProgressValue(t.instance, value)
}

// OverlayIcon
func (t *TTaskbar) OverlayIcon() *TIcon {
    return IconFromInst(Taskbar_GetOverlayIcon(t.instance))
}

// SetOverlayIcon
func (t *TTaskbar) SetOverlayIcon(value *TIcon) {
    Taskbar_SetOverlayIcon(t.instance, CheckPtr(value))
}

// OverlayHint
func (t *TTaskbar) OverlayHint() string {
    return Taskbar_GetOverlayHint(t.instance)
}

// SetOverlayHint
func (t *TTaskbar) SetOverlayHint(value string) {
    Taskbar_SetOverlayHint(t.instance, value)
}

// PreviewClipRegion
func (t *TTaskbar) PreviewClipRegion() *TPreviewClipRegion {
    return PreviewClipRegionFromInst(Taskbar_GetPreviewClipRegion(t.instance))
}

// SetPreviewClipRegion
func (t *TTaskbar) SetPreviewClipRegion(value *TPreviewClipRegion) {
    Taskbar_SetPreviewClipRegion(t.instance, CheckPtr(value))
}

// TabProperties
func (t *TTaskbar) TabProperties() TThumbTabProperties {
    return Taskbar_GetTabProperties(t.instance)
}

// SetTabProperties
func (t *TTaskbar) SetTabProperties(value TThumbTabProperties) {
    Taskbar_SetTabProperties(t.instance, value)
}

// ToolTip
func (t *TTaskbar) ToolTip() string {
    return Taskbar_GetToolTip(t.instance)
}

// SetToolTip
func (t *TTaskbar) SetToolTip(value string) {
    Taskbar_SetToolTip(t.instance, value)
}

// SetOnThumbPreviewRequest
func (t *TTaskbar) SetOnThumbPreviewRequest(fn TThumbPreviewItemRequestEvent) {
    Taskbar_SetOnThumbPreviewRequest(t.instance, fn)
}

// SetOnWindowPreviewItemRequest
func (t *TTaskbar) SetOnWindowPreviewItemRequest(fn TWindowPreviewItemRequestEvent) {
    Taskbar_SetOnWindowPreviewItemRequest(t.instance, fn)
}

// SetOnThumbButtonClick
func (t *TTaskbar) SetOnThumbButtonClick(fn TThumbButtonNotifyEvent) {
    Taskbar_SetOnThumbButtonClick(t.instance, fn)
}

// TaskbarIsAvailable
func (t *TTaskbar) TaskbarIsAvailable() bool {
    return Taskbar_GetTaskbarIsAvailable(t.instance)
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTaskbar) ComponentCount() int32 {
    return Taskbar_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TTaskbar) ComponentIndex() int32 {
    return Taskbar_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TTaskbar) SetComponentIndex(value int32) {
    Taskbar_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTaskbar) Owner() *TComponent {
    return ComponentFromInst(Taskbar_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTaskbar) Name() string {
    return Taskbar_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTaskbar) SetName(value string) {
    Taskbar_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTaskbar) Tag() int {
    return Taskbar_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTaskbar) SetTag(value int) {
    Taskbar_SetTag(t.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTaskbar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Taskbar_GetComponents(t.instance, AIndex))
}

