
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

type TPaintBox struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPaintBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPaintBox(owner IComponent) *TPaintBox {
    p := new(TPaintBox)
    p.instance = PaintBox_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PaintBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PaintBoxFromInst(inst uintptr) *TPaintBox {
    p := new(TPaintBox)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PaintBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PaintBoxFromObj(obj IObject) *TPaintBox {
    p := new(TPaintBox)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PaintBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PaintBoxFromUnsafePointer(ptr unsafe.Pointer) *TPaintBox {
    p := new(TPaintBox)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPaintBox) Free() {
    if p.instance != 0 {
        PaintBox_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPaintBox) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPaintBox) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPaintBox) IsValid() bool {
    return p.instance != 0
}

// TPaintBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPaintBoxClass() TClass {
    return PaintBox_StaticClassType()
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (p *TPaintBox) BringToFront() {
    PaintBox_BringToFront(p.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (p *TPaintBox) ClientToScreen(Point TPoint) TPoint {
    return PaintBox_ClientToScreen(p.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (p *TPaintBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return PaintBox_ClientToParent(p.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (p *TPaintBox) Dragging() bool {
    return PaintBox_Dragging(p.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPaintBox) HasParent() bool {
    return PaintBox_HasParent(p.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (p *TPaintBox) Hide() {
    PaintBox_Hide(p.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (p *TPaintBox) Invalidate() {
    PaintBox_Invalidate(p.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (p *TPaintBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return PaintBox_Perform(p.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (p *TPaintBox) Refresh() {
    PaintBox_Refresh(p.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (p *TPaintBox) Repaint() {
    PaintBox_Repaint(p.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (p *TPaintBox) ScreenToClient(Point TPoint) TPoint {
    return PaintBox_ScreenToClient(p.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (p *TPaintBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return PaintBox_ParentToClient(p.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (p *TPaintBox) SendToBack() {
    PaintBox_SendToBack(p.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (p *TPaintBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    PaintBox_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (p *TPaintBox) Show() {
    PaintBox_Show(p.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (p *TPaintBox) Update() {
    PaintBox_Update(p.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (p *TPaintBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return PaintBox_GetTextBuf(p.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (p *TPaintBox) GetTextLen() int32 {
    return PaintBox_GetTextLen(p.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (p *TPaintBox) SetTextBuf(Buffer string) {
    PaintBox_SetTextBuf(p.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPaintBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PaintBox_FindComponent(p.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPaintBox) GetNamePath() string {
    return PaintBox_GetNamePath(p.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPaintBox) Assign(Source IObject) {
    PaintBox_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPaintBox) DisposeOf() {
    PaintBox_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPaintBox) ClassType() TClass {
    return PaintBox_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPaintBox) ClassName() string {
    return PaintBox_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPaintBox) InstanceSize() int32 {
    return PaintBox_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPaintBox) InheritsFrom(AClass TClass) bool {
    return PaintBox_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPaintBox) Equals(Obj IObject) bool {
    return PaintBox_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPaintBox) GetHashCode() int32 {
    return PaintBox_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPaintBox) ToString() string {
    return PaintBox_ToString(p.instance)
}

// Canvas
// CN: 获取画布。
// EN: .
func (p *TPaintBox) Canvas() *TCanvas {
    return CanvasFromInst(PaintBox_GetCanvas(p.instance))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (p *TPaintBox) Align() TAlign {
    return PaintBox_GetAlign(p.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (p *TPaintBox) SetAlign(value TAlign) {
    PaintBox_SetAlign(p.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (p *TPaintBox) Anchors() TAnchors {
    return PaintBox_GetAnchors(p.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (p *TPaintBox) SetAnchors(value TAnchors) {
    PaintBox_SetAnchors(p.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (p *TPaintBox) Color() TColor {
    return PaintBox_GetColor(p.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (p *TPaintBox) SetColor(value TColor) {
    PaintBox_SetColor(p.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (p *TPaintBox) DragCursor() TCursor {
    return PaintBox_GetDragCursor(p.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (p *TPaintBox) SetDragCursor(value TCursor) {
    PaintBox_SetDragCursor(p.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (p *TPaintBox) DragKind() TDragKind {
    return PaintBox_GetDragKind(p.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (p *TPaintBox) SetDragKind(value TDragKind) {
    PaintBox_SetDragKind(p.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (p *TPaintBox) DragMode() TDragMode {
    return PaintBox_GetDragMode(p.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (p *TPaintBox) SetDragMode(value TDragMode) {
    PaintBox_SetDragMode(p.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (p *TPaintBox) Enabled() bool {
    return PaintBox_GetEnabled(p.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (p *TPaintBox) SetEnabled(value bool) {
    PaintBox_SetEnabled(p.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (p *TPaintBox) Font() *TFont {
    return FontFromInst(PaintBox_GetFont(p.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (p *TPaintBox) SetFont(value *TFont) {
    PaintBox_SetFont(p.instance, CheckPtr(value))
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (p *TPaintBox) ParentColor() bool {
    return PaintBox_GetParentColor(p.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (p *TPaintBox) SetParentColor(value bool) {
    PaintBox_SetParentColor(p.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (p *TPaintBox) ParentFont() bool {
    return PaintBox_GetParentFont(p.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (p *TPaintBox) SetParentFont(value bool) {
    PaintBox_SetParentFont(p.instance, value)
}

// ParentShowHint
func (p *TPaintBox) ParentShowHint() bool {
    return PaintBox_GetParentShowHint(p.instance)
}

// SetParentShowHint
func (p *TPaintBox) SetParentShowHint(value bool) {
    PaintBox_SetParentShowHint(p.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (p *TPaintBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(PaintBox_GetPopupMenu(p.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (p *TPaintBox) SetPopupMenu(value IComponent) {
    PaintBox_SetPopupMenu(p.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (p *TPaintBox) ShowHint() bool {
    return PaintBox_GetShowHint(p.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (p *TPaintBox) SetShowHint(value bool) {
    PaintBox_SetShowHint(p.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (p *TPaintBox) Visible() bool {
    return PaintBox_GetVisible(p.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (p *TPaintBox) SetVisible(value bool) {
    PaintBox_SetVisible(p.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (p *TPaintBox) SetOnClick(fn TNotifyEvent) {
    PaintBox_SetOnClick(p.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (p *TPaintBox) SetOnContextPopup(fn TContextPopupEvent) {
    PaintBox_SetOnContextPopup(p.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (p *TPaintBox) SetOnDblClick(fn TNotifyEvent) {
    PaintBox_SetOnDblClick(p.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (p *TPaintBox) SetOnDragDrop(fn TDragDropEvent) {
    PaintBox_SetOnDragDrop(p.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (p *TPaintBox) SetOnDragOver(fn TDragOverEvent) {
    PaintBox_SetOnDragOver(p.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (p *TPaintBox) SetOnEndDock(fn TEndDragEvent) {
    PaintBox_SetOnEndDock(p.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (p *TPaintBox) SetOnEndDrag(fn TEndDragEvent) {
    PaintBox_SetOnEndDrag(p.instance, fn)
}

// SetOnGesture
func (p *TPaintBox) SetOnGesture(fn TGestureEvent) {
    PaintBox_SetOnGesture(p.instance, fn)
}

// SetOnMouseActivate
func (p *TPaintBox) SetOnMouseActivate(fn TMouseActivateEvent) {
    PaintBox_SetOnMouseActivate(p.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (p *TPaintBox) SetOnMouseDown(fn TMouseEvent) {
    PaintBox_SetOnMouseDown(p.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (p *TPaintBox) SetOnMouseEnter(fn TNotifyEvent) {
    PaintBox_SetOnMouseEnter(p.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (p *TPaintBox) SetOnMouseLeave(fn TNotifyEvent) {
    PaintBox_SetOnMouseLeave(p.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (p *TPaintBox) SetOnMouseMove(fn TMouseMoveEvent) {
    PaintBox_SetOnMouseMove(p.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (p *TPaintBox) SetOnMouseUp(fn TMouseEvent) {
    PaintBox_SetOnMouseUp(p.instance, fn)
}

// SetOnPaint
// CN: 设置绘画事件。
// EN: .
func (p *TPaintBox) SetOnPaint(fn TNotifyEvent) {
    PaintBox_SetOnPaint(p.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (p *TPaintBox) SetOnStartDock(fn TStartDockEvent) {
    PaintBox_SetOnStartDock(p.instance, fn)
}

// Action
func (p *TPaintBox) Action() *TAction {
    return ActionFromInst(PaintBox_GetAction(p.instance))
}

// SetAction
func (p *TPaintBox) SetAction(value IComponent) {
    PaintBox_SetAction(p.instance, CheckPtr(value))
}

// BiDiMode
func (p *TPaintBox) BiDiMode() TBiDiMode {
    return PaintBox_GetBiDiMode(p.instance)
}

// SetBiDiMode
func (p *TPaintBox) SetBiDiMode(value TBiDiMode) {
    PaintBox_SetBiDiMode(p.instance, value)
}

// BoundsRect
func (p *TPaintBox) BoundsRect() TRect {
    return PaintBox_GetBoundsRect(p.instance)
}

// SetBoundsRect
func (p *TPaintBox) SetBoundsRect(value TRect) {
    PaintBox_SetBoundsRect(p.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (p *TPaintBox) ClientHeight() int32 {
    return PaintBox_GetClientHeight(p.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (p *TPaintBox) SetClientHeight(value int32) {
    PaintBox_SetClientHeight(p.instance, value)
}

// ClientOrigin
func (p *TPaintBox) ClientOrigin() TPoint {
    return PaintBox_GetClientOrigin(p.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (p *TPaintBox) ClientRect() TRect {
    return PaintBox_GetClientRect(p.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (p *TPaintBox) ClientWidth() int32 {
    return PaintBox_GetClientWidth(p.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (p *TPaintBox) SetClientWidth(value int32) {
    PaintBox_SetClientWidth(p.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (p *TPaintBox) ControlState() TControlState {
    return PaintBox_GetControlState(p.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (p *TPaintBox) SetControlState(value TControlState) {
    PaintBox_SetControlState(p.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (p *TPaintBox) ControlStyle() TControlStyle {
    return PaintBox_GetControlStyle(p.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (p *TPaintBox) SetControlStyle(value TControlStyle) {
    PaintBox_SetControlStyle(p.instance, value)
}

// ExplicitLeft
func (p *TPaintBox) ExplicitLeft() int32 {
    return PaintBox_GetExplicitLeft(p.instance)
}

// ExplicitTop
func (p *TPaintBox) ExplicitTop() int32 {
    return PaintBox_GetExplicitTop(p.instance)
}

// ExplicitWidth
func (p *TPaintBox) ExplicitWidth() int32 {
    return PaintBox_GetExplicitWidth(p.instance)
}

// ExplicitHeight
func (p *TPaintBox) ExplicitHeight() int32 {
    return PaintBox_GetExplicitHeight(p.instance)
}

// Floating
func (p *TPaintBox) Floating() bool {
    return PaintBox_GetFloating(p.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (p *TPaintBox) Parent() *TWinControl {
    return WinControlFromInst(PaintBox_GetParent(p.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (p *TPaintBox) SetParent(value IWinControl) {
    PaintBox_SetParent(p.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (p *TPaintBox) StyleElements() TStyleElements {
    return PaintBox_GetStyleElements(p.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (p *TPaintBox) SetStyleElements(value TStyleElements) {
    PaintBox_SetStyleElements(p.instance, value)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (p *TPaintBox) AlignWithMargins() bool {
    return PaintBox_GetAlignWithMargins(p.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (p *TPaintBox) SetAlignWithMargins(value bool) {
    PaintBox_SetAlignWithMargins(p.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (p *TPaintBox) Left() int32 {
    return PaintBox_GetLeft(p.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (p *TPaintBox) SetLeft(value int32) {
    PaintBox_SetLeft(p.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (p *TPaintBox) Top() int32 {
    return PaintBox_GetTop(p.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (p *TPaintBox) SetTop(value int32) {
    PaintBox_SetTop(p.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (p *TPaintBox) Width() int32 {
    return PaintBox_GetWidth(p.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (p *TPaintBox) SetWidth(value int32) {
    PaintBox_SetWidth(p.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (p *TPaintBox) Height() int32 {
    return PaintBox_GetHeight(p.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (p *TPaintBox) SetHeight(value int32) {
    PaintBox_SetHeight(p.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (p *TPaintBox) Cursor() TCursor {
    return PaintBox_GetCursor(p.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (p *TPaintBox) SetCursor(value TCursor) {
    PaintBox_SetCursor(p.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (p *TPaintBox) Hint() string {
    return PaintBox_GetHint(p.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (p *TPaintBox) SetHint(value string) {
    PaintBox_SetHint(p.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (p *TPaintBox) Margins() *TMargins {
    return MarginsFromInst(PaintBox_GetMargins(p.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (p *TPaintBox) SetMargins(value *TMargins) {
    PaintBox_SetMargins(p.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (p *TPaintBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(PaintBox_GetCustomHint(p.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (p *TPaintBox) SetCustomHint(value IComponent) {
    PaintBox_SetCustomHint(p.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPaintBox) ComponentCount() int32 {
    return PaintBox_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPaintBox) ComponentIndex() int32 {
    return PaintBox_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPaintBox) SetComponentIndex(value int32) {
    PaintBox_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPaintBox) Owner() *TComponent {
    return ComponentFromInst(PaintBox_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPaintBox) Name() string {
    return PaintBox_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPaintBox) SetName(value string) {
    PaintBox_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPaintBox) Tag() int {
    return PaintBox_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPaintBox) SetTag(value int) {
    PaintBox_SetTag(p.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPaintBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PaintBox_GetComponents(p.instance, AIndex))
}

