
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

type TGauge struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewGauge
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGauge(owner IComponent) *TGauge {
    g := new(TGauge)
    g.instance = Gauge_Create(CheckPtr(owner))
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GaugeFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func GaugeFromInst(inst uintptr) *TGauge {
    g := new(TGauge)
    g.instance = inst
    g.ptr = unsafe.Pointer(inst)
    return g
}

// GaugeFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func GaugeFromObj(obj IObject) *TGauge {
    g := new(TGauge)
    g.instance = CheckPtr(obj)
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GaugeFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func GaugeFromUnsafePointer(ptr unsafe.Pointer) *TGauge {
    g := new(TGauge)
    g.instance = uintptr(ptr)
    g.ptr = ptr
    return g
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (g *TGauge) Free() {
    if g.instance != 0 {
        Gauge_Free(g.instance)
        g.instance = 0
        g.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGauge) Instance() uintptr {
    return g.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGauge) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGauge) IsValid() bool {
    return g.instance != 0
}

// TGaugeClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGaugeClass() TClass {
    return Gauge_StaticClassType()
}

// AddProgress
func (g *TGauge) AddProgress(Value int32) {
    Gauge_AddProgress(g.instance, Value)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (g *TGauge) BringToFront() {
    Gauge_BringToFront(g.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (g *TGauge) ClientToScreen(Point TPoint) TPoint {
    return Gauge_ClientToScreen(g.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (g *TGauge) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ClientToParent(g.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (g *TGauge) Dragging() bool {
    return Gauge_Dragging(g.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (g *TGauge) HasParent() bool {
    return Gauge_HasParent(g.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (g *TGauge) Hide() {
    Gauge_Hide(g.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (g *TGauge) Invalidate() {
    Gauge_Invalidate(g.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (g *TGauge) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Gauge_Perform(g.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (g *TGauge) Refresh() {
    Gauge_Refresh(g.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (g *TGauge) Repaint() {
    Gauge_Repaint(g.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (g *TGauge) ScreenToClient(Point TPoint) TPoint {
    return Gauge_ScreenToClient(g.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (g *TGauge) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ParentToClient(g.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (g *TGauge) SendToBack() {
    Gauge_SendToBack(g.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (g *TGauge) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Gauge_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (g *TGauge) Show() {
    Gauge_Show(g.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (g *TGauge) Update() {
    Gauge_Update(g.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (g *TGauge) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Gauge_GetTextBuf(g.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (g *TGauge) GetTextLen() int32 {
    return Gauge_GetTextLen(g.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (g *TGauge) SetTextBuf(Buffer string) {
    Gauge_SetTextBuf(g.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (g *TGauge) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Gauge_FindComponent(g.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (g *TGauge) GetNamePath() string {
    return Gauge_GetNamePath(g.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (g *TGauge) Assign(Source IObject) {
    Gauge_Assign(g.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (g *TGauge) DisposeOf() {
    Gauge_DisposeOf(g.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGauge) ClassType() TClass {
    return Gauge_ClassType(g.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGauge) ClassName() string {
    return Gauge_ClassName(g.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGauge) InstanceSize() int32 {
    return Gauge_InstanceSize(g.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGauge) InheritsFrom(AClass TClass) bool {
    return Gauge_InheritsFrom(g.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGauge) Equals(Obj IObject) bool {
    return Gauge_Equals(g.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGauge) GetHashCode() int32 {
    return Gauge_GetHashCode(g.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (g *TGauge) ToString() string {
    return Gauge_ToString(g.instance)
}

// PercentDone
func (g *TGauge) PercentDone() int32 {
    return Gauge_GetPercentDone(g.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (g *TGauge) Align() TAlign {
    return Gauge_GetAlign(g.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (g *TGauge) SetAlign(value TAlign) {
    Gauge_SetAlign(g.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (g *TGauge) Anchors() TAnchors {
    return Gauge_GetAnchors(g.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (g *TGauge) SetAnchors(value TAnchors) {
    Gauge_SetAnchors(g.instance, value)
}

// BackColor
func (g *TGauge) BackColor() TColor {
    return Gauge_GetBackColor(g.instance)
}

// SetBackColor
func (g *TGauge) SetBackColor(value TColor) {
    Gauge_SetBackColor(g.instance, value)
}

// BorderStyle
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (g *TGauge) BorderStyle() TBorderStyle {
    return Gauge_GetBorderStyle(g.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (g *TGauge) SetBorderStyle(value TBorderStyle) {
    Gauge_SetBorderStyle(g.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (g *TGauge) Color() TColor {
    return Gauge_GetColor(g.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (g *TGauge) SetColor(value TColor) {
    Gauge_SetColor(g.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (g *TGauge) Enabled() bool {
    return Gauge_GetEnabled(g.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (g *TGauge) SetEnabled(value bool) {
    Gauge_SetEnabled(g.instance, value)
}

// ForeColor
func (g *TGauge) ForeColor() TColor {
    return Gauge_GetForeColor(g.instance)
}

// SetForeColor
func (g *TGauge) SetForeColor(value TColor) {
    Gauge_SetForeColor(g.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (g *TGauge) Font() *TFont {
    return FontFromInst(Gauge_GetFont(g.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (g *TGauge) SetFont(value *TFont) {
    Gauge_SetFont(g.instance, CheckPtr(value))
}

// Kind
func (g *TGauge) Kind() TGaugeKind {
    return Gauge_GetKind(g.instance)
}

// SetKind
func (g *TGauge) SetKind(value TGaugeKind) {
    Gauge_SetKind(g.instance, value)
}

// MinValue
func (g *TGauge) MinValue() int32 {
    return Gauge_GetMinValue(g.instance)
}

// SetMinValue
func (g *TGauge) SetMinValue(value int32) {
    Gauge_SetMinValue(g.instance, value)
}

// MaxValue
func (g *TGauge) MaxValue() int32 {
    return Gauge_GetMaxValue(g.instance)
}

// SetMaxValue
func (g *TGauge) SetMaxValue(value int32) {
    Gauge_SetMaxValue(g.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (g *TGauge) ParentColor() bool {
    return Gauge_GetParentColor(g.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (g *TGauge) SetParentColor(value bool) {
    Gauge_SetParentColor(g.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (g *TGauge) ParentFont() bool {
    return Gauge_GetParentFont(g.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (g *TGauge) SetParentFont(value bool) {
    Gauge_SetParentFont(g.instance, value)
}

// ParentShowHint
func (g *TGauge) ParentShowHint() bool {
    return Gauge_GetParentShowHint(g.instance)
}

// SetParentShowHint
func (g *TGauge) SetParentShowHint(value bool) {
    Gauge_SetParentShowHint(g.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (g *TGauge) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Gauge_GetPopupMenu(g.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (g *TGauge) SetPopupMenu(value IComponent) {
    Gauge_SetPopupMenu(g.instance, CheckPtr(value))
}

// Progress
func (g *TGauge) Progress() int32 {
    return Gauge_GetProgress(g.instance)
}

// SetProgress
func (g *TGauge) SetProgress(value int32) {
    Gauge_SetProgress(g.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (g *TGauge) ShowHint() bool {
    return Gauge_GetShowHint(g.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (g *TGauge) SetShowHint(value bool) {
    Gauge_SetShowHint(g.instance, value)
}

// ShowText
func (g *TGauge) ShowText() bool {
    return Gauge_GetShowText(g.instance)
}

// SetShowText
func (g *TGauge) SetShowText(value bool) {
    Gauge_SetShowText(g.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (g *TGauge) Visible() bool {
    return Gauge_GetVisible(g.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (g *TGauge) SetVisible(value bool) {
    Gauge_SetVisible(g.instance, value)
}

// Action
func (g *TGauge) Action() *TAction {
    return ActionFromInst(Gauge_GetAction(g.instance))
}

// SetAction
func (g *TGauge) SetAction(value IComponent) {
    Gauge_SetAction(g.instance, CheckPtr(value))
}

// BiDiMode
func (g *TGauge) BiDiMode() TBiDiMode {
    return Gauge_GetBiDiMode(g.instance)
}

// SetBiDiMode
func (g *TGauge) SetBiDiMode(value TBiDiMode) {
    Gauge_SetBiDiMode(g.instance, value)
}

// BoundsRect
func (g *TGauge) BoundsRect() TRect {
    return Gauge_GetBoundsRect(g.instance)
}

// SetBoundsRect
func (g *TGauge) SetBoundsRect(value TRect) {
    Gauge_SetBoundsRect(g.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (g *TGauge) ClientHeight() int32 {
    return Gauge_GetClientHeight(g.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (g *TGauge) SetClientHeight(value int32) {
    Gauge_SetClientHeight(g.instance, value)
}

// ClientOrigin
func (g *TGauge) ClientOrigin() TPoint {
    return Gauge_GetClientOrigin(g.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (g *TGauge) ClientRect() TRect {
    return Gauge_GetClientRect(g.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (g *TGauge) ClientWidth() int32 {
    return Gauge_GetClientWidth(g.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (g *TGauge) SetClientWidth(value int32) {
    Gauge_SetClientWidth(g.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (g *TGauge) ControlState() TControlState {
    return Gauge_GetControlState(g.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (g *TGauge) SetControlState(value TControlState) {
    Gauge_SetControlState(g.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (g *TGauge) ControlStyle() TControlStyle {
    return Gauge_GetControlStyle(g.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (g *TGauge) SetControlStyle(value TControlStyle) {
    Gauge_SetControlStyle(g.instance, value)
}

// ExplicitLeft
func (g *TGauge) ExplicitLeft() int32 {
    return Gauge_GetExplicitLeft(g.instance)
}

// ExplicitTop
func (g *TGauge) ExplicitTop() int32 {
    return Gauge_GetExplicitTop(g.instance)
}

// ExplicitWidth
func (g *TGauge) ExplicitWidth() int32 {
    return Gauge_GetExplicitWidth(g.instance)
}

// ExplicitHeight
func (g *TGauge) ExplicitHeight() int32 {
    return Gauge_GetExplicitHeight(g.instance)
}

// Floating
func (g *TGauge) Floating() bool {
    return Gauge_GetFloating(g.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (g *TGauge) Parent() *TWinControl {
    return WinControlFromInst(Gauge_GetParent(g.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (g *TGauge) SetParent(value IWinControl) {
    Gauge_SetParent(g.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (g *TGauge) StyleElements() TStyleElements {
    return Gauge_GetStyleElements(g.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (g *TGauge) SetStyleElements(value TStyleElements) {
    Gauge_SetStyleElements(g.instance, value)
}

// SetOnGesture
func (g *TGauge) SetOnGesture(fn TGestureEvent) {
    Gauge_SetOnGesture(g.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (g *TGauge) AlignWithMargins() bool {
    return Gauge_GetAlignWithMargins(g.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (g *TGauge) SetAlignWithMargins(value bool) {
    Gauge_SetAlignWithMargins(g.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (g *TGauge) Left() int32 {
    return Gauge_GetLeft(g.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (g *TGauge) SetLeft(value int32) {
    Gauge_SetLeft(g.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (g *TGauge) Top() int32 {
    return Gauge_GetTop(g.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (g *TGauge) SetTop(value int32) {
    Gauge_SetTop(g.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (g *TGauge) Width() int32 {
    return Gauge_GetWidth(g.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (g *TGauge) SetWidth(value int32) {
    Gauge_SetWidth(g.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (g *TGauge) Height() int32 {
    return Gauge_GetHeight(g.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (g *TGauge) SetHeight(value int32) {
    Gauge_SetHeight(g.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (g *TGauge) Cursor() TCursor {
    return Gauge_GetCursor(g.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (g *TGauge) SetCursor(value TCursor) {
    Gauge_SetCursor(g.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (g *TGauge) Hint() string {
    return Gauge_GetHint(g.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (g *TGauge) SetHint(value string) {
    Gauge_SetHint(g.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (g *TGauge) Margins() *TMargins {
    return MarginsFromInst(Gauge_GetMargins(g.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (g *TGauge) SetMargins(value *TMargins) {
    Gauge_SetMargins(g.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (g *TGauge) CustomHint() *TCustomHint {
    return CustomHintFromInst(Gauge_GetCustomHint(g.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (g *TGauge) SetCustomHint(value IComponent) {
    Gauge_SetCustomHint(g.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (g *TGauge) ComponentCount() int32 {
    return Gauge_GetComponentCount(g.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (g *TGauge) ComponentIndex() int32 {
    return Gauge_GetComponentIndex(g.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (g *TGauge) SetComponentIndex(value int32) {
    Gauge_SetComponentIndex(g.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (g *TGauge) Owner() *TComponent {
    return ComponentFromInst(Gauge_GetOwner(g.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (g *TGauge) Name() string {
    return Gauge_GetName(g.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (g *TGauge) SetName(value string) {
    Gauge_SetName(g.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (g *TGauge) Tag() int {
    return Gauge_GetTag(g.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (g *TGauge) SetTag(value int) {
    Gauge_SetTag(g.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (g *TGauge) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Gauge_GetComponents(g.instance, AIndex))
}

