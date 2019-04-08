
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

type TBevel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewBevel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBevel(owner IComponent) *TBevel {
    b := new(TBevel)
    b.instance = Bevel_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BevelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func BevelFromInst(inst uintptr) *TBevel {
    b := new(TBevel)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// BevelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func BevelFromObj(obj IObject) *TBevel {
    b := new(TBevel)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BevelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func BevelFromUnsafePointer(ptr unsafe.Pointer) *TBevel {
    b := new(TBevel)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TBevel) Free() {
    if b.instance != 0 {
        Bevel_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBevel) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBevel) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBevel) IsValid() bool {
    return b.instance != 0
}

// TBevelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBevelClass() TClass {
    return Bevel_StaticClassType()
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (b *TBevel) BringToFront() {
    Bevel_BringToFront(b.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (b *TBevel) ClientToScreen(Point TPoint) TPoint {
    return Bevel_ClientToScreen(b.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (b *TBevel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (b *TBevel) Dragging() bool {
    return Bevel_Dragging(b.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (b *TBevel) HasParent() bool {
    return Bevel_HasParent(b.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (b *TBevel) Hide() {
    Bevel_Hide(b.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (b *TBevel) Invalidate() {
    Bevel_Invalidate(b.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (b *TBevel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Bevel_Perform(b.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (b *TBevel) Refresh() {
    Bevel_Refresh(b.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (b *TBevel) Repaint() {
    Bevel_Repaint(b.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (b *TBevel) ScreenToClient(Point TPoint) TPoint {
    return Bevel_ScreenToClient(b.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (b *TBevel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Bevel_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (b *TBevel) SendToBack() {
    Bevel_SendToBack(b.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (b *TBevel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Bevel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (b *TBevel) Show() {
    Bevel_Show(b.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (b *TBevel) Update() {
    Bevel_Update(b.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (b *TBevel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Bevel_GetTextBuf(b.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (b *TBevel) GetTextLen() int32 {
    return Bevel_GetTextLen(b.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (b *TBevel) SetTextBuf(Buffer string) {
    Bevel_SetTextBuf(b.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (b *TBevel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Bevel_FindComponent(b.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBevel) GetNamePath() string {
    return Bevel_GetNamePath(b.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBevel) Assign(Source IObject) {
    Bevel_Assign(b.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBevel) DisposeOf() {
    Bevel_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBevel) ClassType() TClass {
    return Bevel_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBevel) ClassName() string {
    return Bevel_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBevel) InstanceSize() int32 {
    return Bevel_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBevel) InheritsFrom(AClass TClass) bool {
    return Bevel_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBevel) Equals(Obj IObject) bool {
    return Bevel_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBevel) GetHashCode() int32 {
    return Bevel_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TBevel) ToString() string {
    return Bevel_ToString(b.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (b *TBevel) Align() TAlign {
    return Bevel_GetAlign(b.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (b *TBevel) SetAlign(value TAlign) {
    Bevel_SetAlign(b.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (b *TBevel) Anchors() TAnchors {
    return Bevel_GetAnchors(b.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (b *TBevel) SetAnchors(value TAnchors) {
    Bevel_SetAnchors(b.instance, value)
}

// ParentShowHint
func (b *TBevel) ParentShowHint() bool {
    return Bevel_GetParentShowHint(b.instance)
}

// SetParentShowHint
func (b *TBevel) SetParentShowHint(value bool) {
    Bevel_SetParentShowHint(b.instance, value)
}

// Shape
func (b *TBevel) Shape() TBevelShape {
    return Bevel_GetShape(b.instance)
}

// SetShape
func (b *TBevel) SetShape(value TBevelShape) {
    Bevel_SetShape(b.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (b *TBevel) ShowHint() bool {
    return Bevel_GetShowHint(b.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (b *TBevel) SetShowHint(value bool) {
    Bevel_SetShowHint(b.instance, value)
}

// Style
func (b *TBevel) Style() TBevelStyle {
    return Bevel_GetStyle(b.instance)
}

// SetStyle
func (b *TBevel) SetStyle(value TBevelStyle) {
    Bevel_SetStyle(b.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBevel) Visible() bool {
    return Bevel_GetVisible(b.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBevel) SetVisible(value bool) {
    Bevel_SetVisible(b.instance, value)
}

// SetOnGesture
func (b *TBevel) SetOnGesture(fn TGestureEvent) {
    Bevel_SetOnGesture(b.instance, fn)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBevel) Enabled() bool {
    return Bevel_GetEnabled(b.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBevel) SetEnabled(value bool) {
    Bevel_SetEnabled(b.instance, value)
}

// Action
func (b *TBevel) Action() *TAction {
    return ActionFromInst(Bevel_GetAction(b.instance))
}

// SetAction
func (b *TBevel) SetAction(value IComponent) {
    Bevel_SetAction(b.instance, CheckPtr(value))
}

// BiDiMode
func (b *TBevel) BiDiMode() TBiDiMode {
    return Bevel_GetBiDiMode(b.instance)
}

// SetBiDiMode
func (b *TBevel) SetBiDiMode(value TBiDiMode) {
    Bevel_SetBiDiMode(b.instance, value)
}

// BoundsRect
func (b *TBevel) BoundsRect() TRect {
    return Bevel_GetBoundsRect(b.instance)
}

// SetBoundsRect
func (b *TBevel) SetBoundsRect(value TRect) {
    Bevel_SetBoundsRect(b.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (b *TBevel) ClientHeight() int32 {
    return Bevel_GetClientHeight(b.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (b *TBevel) SetClientHeight(value int32) {
    Bevel_SetClientHeight(b.instance, value)
}

// ClientOrigin
func (b *TBevel) ClientOrigin() TPoint {
    return Bevel_GetClientOrigin(b.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (b *TBevel) ClientRect() TRect {
    return Bevel_GetClientRect(b.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (b *TBevel) ClientWidth() int32 {
    return Bevel_GetClientWidth(b.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (b *TBevel) SetClientWidth(value int32) {
    Bevel_SetClientWidth(b.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (b *TBevel) ControlState() TControlState {
    return Bevel_GetControlState(b.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (b *TBevel) SetControlState(value TControlState) {
    Bevel_SetControlState(b.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (b *TBevel) ControlStyle() TControlStyle {
    return Bevel_GetControlStyle(b.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (b *TBevel) SetControlStyle(value TControlStyle) {
    Bevel_SetControlStyle(b.instance, value)
}

// ExplicitLeft
func (b *TBevel) ExplicitLeft() int32 {
    return Bevel_GetExplicitLeft(b.instance)
}

// ExplicitTop
func (b *TBevel) ExplicitTop() int32 {
    return Bevel_GetExplicitTop(b.instance)
}

// ExplicitWidth
func (b *TBevel) ExplicitWidth() int32 {
    return Bevel_GetExplicitWidth(b.instance)
}

// ExplicitHeight
func (b *TBevel) ExplicitHeight() int32 {
    return Bevel_GetExplicitHeight(b.instance)
}

// Floating
func (b *TBevel) Floating() bool {
    return Bevel_GetFloating(b.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBevel) Parent() *TWinControl {
    return WinControlFromInst(Bevel_GetParent(b.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBevel) SetParent(value IWinControl) {
    Bevel_SetParent(b.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (b *TBevel) StyleElements() TStyleElements {
    return Bevel_GetStyleElements(b.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (b *TBevel) SetStyleElements(value TStyleElements) {
    Bevel_SetStyleElements(b.instance, value)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (b *TBevel) AlignWithMargins() bool {
    return Bevel_GetAlignWithMargins(b.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (b *TBevel) SetAlignWithMargins(value bool) {
    Bevel_SetAlignWithMargins(b.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (b *TBevel) Left() int32 {
    return Bevel_GetLeft(b.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (b *TBevel) SetLeft(value int32) {
    Bevel_SetLeft(b.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (b *TBevel) Top() int32 {
    return Bevel_GetTop(b.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (b *TBevel) SetTop(value int32) {
    Bevel_SetTop(b.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (b *TBevel) Width() int32 {
    return Bevel_GetWidth(b.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (b *TBevel) SetWidth(value int32) {
    Bevel_SetWidth(b.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (b *TBevel) Height() int32 {
    return Bevel_GetHeight(b.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (b *TBevel) SetHeight(value int32) {
    Bevel_SetHeight(b.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBevel) Cursor() TCursor {
    return Bevel_GetCursor(b.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBevel) SetCursor(value TCursor) {
    Bevel_SetCursor(b.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (b *TBevel) Hint() string {
    return Bevel_GetHint(b.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (b *TBevel) SetHint(value string) {
    Bevel_SetHint(b.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (b *TBevel) Margins() *TMargins {
    return MarginsFromInst(Bevel_GetMargins(b.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (b *TBevel) SetMargins(value *TMargins) {
    Bevel_SetMargins(b.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (b *TBevel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Bevel_GetCustomHint(b.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (b *TBevel) SetCustomHint(value IComponent) {
    Bevel_SetCustomHint(b.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBevel) ComponentCount() int32 {
    return Bevel_GetComponentCount(b.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (b *TBevel) ComponentIndex() int32 {
    return Bevel_GetComponentIndex(b.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (b *TBevel) SetComponentIndex(value int32) {
    Bevel_SetComponentIndex(b.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBevel) Owner() *TComponent {
    return ComponentFromInst(Bevel_GetOwner(b.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBevel) Name() string {
    return Bevel_GetName(b.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBevel) SetName(value string) {
    Bevel_SetName(b.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBevel) Tag() int {
    return Bevel_GetTag(b.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBevel) SetTag(value int) {
    Bevel_SetTag(b.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBevel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Bevel_GetComponents(b.instance, AIndex))
}

