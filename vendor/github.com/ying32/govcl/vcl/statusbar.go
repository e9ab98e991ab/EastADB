
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

type TStatusBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewStatusBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewStatusBar(owner IComponent) *TStatusBar {
    s := new(TStatusBar)
    s.instance = StatusBar_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StatusBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func StatusBarFromInst(inst uintptr) *TStatusBar {
    s := new(TStatusBar)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// StatusBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func StatusBarFromObj(obj IObject) *TStatusBar {
    s := new(TStatusBar)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// StatusBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func StatusBarFromUnsafePointer(ptr unsafe.Pointer) *TStatusBar {
    s := new(TStatusBar)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TStatusBar) Free() {
    if s.instance != 0 {
        StatusBar_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TStatusBar) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TStatusBar) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TStatusBar) IsValid() bool {
    return s.instance != 0
}

// TStatusBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TStatusBarClass() TClass {
    return StatusBar_StaticClassType()
}

// FlipChildren
func (s *TStatusBar) FlipChildren(AllLevels bool) {
    StatusBar_FlipChildren(s.instance, AllLevels)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TStatusBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    StatusBar_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (s *TStatusBar) CanFocus() bool {
    return StatusBar_CanFocus(s.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (s *TStatusBar) ContainsControl(Control IControl) bool {
    return StatusBar_ContainsControl(s.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (s *TStatusBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(StatusBar_ControlAtPos(s.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (s *TStatusBar) DisableAlign() {
    StatusBar_DisableAlign(s.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (s *TStatusBar) EnableAlign() {
    StatusBar_EnableAlign(s.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (s *TStatusBar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(StatusBar_FindChildControl(s.instance, ControlName))
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (s *TStatusBar) Focused() bool {
    return StatusBar_Focused(s.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (s *TStatusBar) HandleAllocated() bool {
    return StatusBar_HandleAllocated(s.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (s *TStatusBar) InsertControl(AControl IControl) {
    StatusBar_InsertControl(s.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (s *TStatusBar) Invalidate() {
    StatusBar_Invalidate(s.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (s *TStatusBar) PaintTo(DC HDC, X int32, Y int32) {
    StatusBar_PaintTo(s.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (s *TStatusBar) RemoveControl(AControl IControl) {
    StatusBar_RemoveControl(s.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (s *TStatusBar) Realign() {
    StatusBar_Realign(s.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (s *TStatusBar) Repaint() {
    StatusBar_Repaint(s.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (s *TStatusBar) ScaleBy(M int32, D int32) {
    StatusBar_ScaleBy(s.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (s *TStatusBar) ScrollBy(DeltaX int32, DeltaY int32) {
    StatusBar_ScrollBy(s.instance, DeltaX , DeltaY)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (s *TStatusBar) SetFocus() {
    StatusBar_SetFocus(s.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (s *TStatusBar) Update() {
    StatusBar_Update(s.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (s *TStatusBar) UpdateControlState() {
    StatusBar_UpdateControlState(s.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TStatusBar) BringToFront() {
    StatusBar_BringToFront(s.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TStatusBar) ClientToScreen(Point TPoint) TPoint {
    return StatusBar_ClientToScreen(s.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TStatusBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return StatusBar_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TStatusBar) Dragging() bool {
    return StatusBar_Dragging(s.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TStatusBar) HasParent() bool {
    return StatusBar_HasParent(s.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (s *TStatusBar) Hide() {
    StatusBar_Hide(s.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (s *TStatusBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return StatusBar_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (s *TStatusBar) Refresh() {
    StatusBar_Refresh(s.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TStatusBar) ScreenToClient(Point TPoint) TPoint {
    return StatusBar_ScreenToClient(s.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TStatusBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return StatusBar_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TStatusBar) SendToBack() {
    StatusBar_SendToBack(s.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (s *TStatusBar) Show() {
    StatusBar_Show(s.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TStatusBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return StatusBar_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TStatusBar) GetTextLen() int32 {
    return StatusBar_GetTextLen(s.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TStatusBar) SetTextBuf(Buffer string) {
    StatusBar_SetTextBuf(s.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TStatusBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(StatusBar_FindComponent(s.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TStatusBar) GetNamePath() string {
    return StatusBar_GetNamePath(s.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TStatusBar) Assign(Source IObject) {
    StatusBar_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TStatusBar) DisposeOf() {
    StatusBar_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TStatusBar) ClassType() TClass {
    return StatusBar_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TStatusBar) ClassName() string {
    return StatusBar_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TStatusBar) InstanceSize() int32 {
    return StatusBar_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TStatusBar) InheritsFrom(AClass TClass) bool {
    return StatusBar_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TStatusBar) Equals(Obj IObject) bool {
    return StatusBar_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TStatusBar) GetHashCode() int32 {
    return StatusBar_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TStatusBar) ToString() string {
    return StatusBar_ToString(s.instance)
}

// Action
func (s *TStatusBar) Action() *TAction {
    return ActionFromInst(StatusBar_GetAction(s.instance))
}

// SetAction
func (s *TStatusBar) SetAction(value IComponent) {
    StatusBar_SetAction(s.instance, CheckPtr(value))
}

// AutoHint
func (s *TStatusBar) AutoHint() bool {
    return StatusBar_GetAutoHint(s.instance)
}

// SetAutoHint
func (s *TStatusBar) SetAutoHint(value bool) {
    StatusBar_SetAutoHint(s.instance, value)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TStatusBar) Align() TAlign {
    return StatusBar_GetAlign(s.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TStatusBar) SetAlign(value TAlign) {
    StatusBar_SetAlign(s.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (s *TStatusBar) Anchors() TAnchors {
    return StatusBar_GetAnchors(s.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (s *TStatusBar) SetAnchors(value TAnchors) {
    StatusBar_SetAnchors(s.instance, value)
}

// BiDiMode
func (s *TStatusBar) BiDiMode() TBiDiMode {
    return StatusBar_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TStatusBar) SetBiDiMode(value TBiDiMode) {
    StatusBar_SetBiDiMode(s.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (s *TStatusBar) BorderWidth() int32 {
    return StatusBar_GetBorderWidth(s.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (s *TStatusBar) SetBorderWidth(value int32) {
    StatusBar_SetBorderWidth(s.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (s *TStatusBar) Color() TColor {
    return StatusBar_GetColor(s.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (s *TStatusBar) SetColor(value TColor) {
    StatusBar_SetColor(s.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (s *TStatusBar) DoubleBuffered() bool {
    return StatusBar_GetDoubleBuffered(s.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (s *TStatusBar) SetDoubleBuffered(value bool) {
    StatusBar_SetDoubleBuffered(s.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (s *TStatusBar) DragCursor() TCursor {
    return StatusBar_GetDragCursor(s.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (s *TStatusBar) SetDragCursor(value TCursor) {
    StatusBar_SetDragCursor(s.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (s *TStatusBar) DragKind() TDragKind {
    return StatusBar_GetDragKind(s.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (s *TStatusBar) SetDragKind(value TDragKind) {
    StatusBar_SetDragKind(s.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (s *TStatusBar) DragMode() TDragMode {
    return StatusBar_GetDragMode(s.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (s *TStatusBar) SetDragMode(value TDragMode) {
    StatusBar_SetDragMode(s.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TStatusBar) Enabled() bool {
    return StatusBar_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TStatusBar) SetEnabled(value bool) {
    StatusBar_SetEnabled(s.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (s *TStatusBar) Font() *TFont {
    return FontFromInst(StatusBar_GetFont(s.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (s *TStatusBar) SetFont(value *TFont) {
    StatusBar_SetFont(s.instance, CheckPtr(value))
}

// Panels
func (s *TStatusBar) Panels() *TStatusPanels {
    return StatusPanelsFromInst(StatusBar_GetPanels(s.instance))
}

// SetPanels
func (s *TStatusBar) SetPanels(value *TStatusPanels) {
    StatusBar_SetPanels(s.instance, CheckPtr(value))
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (s *TStatusBar) ParentColor() bool {
    return StatusBar_GetParentColor(s.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (s *TStatusBar) SetParentColor(value bool) {
    StatusBar_SetParentColor(s.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (s *TStatusBar) ParentDoubleBuffered() bool {
    return StatusBar_GetParentDoubleBuffered(s.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (s *TStatusBar) SetParentDoubleBuffered(value bool) {
    StatusBar_SetParentDoubleBuffered(s.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (s *TStatusBar) ParentFont() bool {
    return StatusBar_GetParentFont(s.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (s *TStatusBar) SetParentFont(value bool) {
    StatusBar_SetParentFont(s.instance, value)
}

// ParentShowHint
func (s *TStatusBar) ParentShowHint() bool {
    return StatusBar_GetParentShowHint(s.instance)
}

// SetParentShowHint
func (s *TStatusBar) SetParentShowHint(value bool) {
    StatusBar_SetParentShowHint(s.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (s *TStatusBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(StatusBar_GetPopupMenu(s.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (s *TStatusBar) SetPopupMenu(value IComponent) {
    StatusBar_SetPopupMenu(s.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TStatusBar) ShowHint() bool {
    return StatusBar_GetShowHint(s.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TStatusBar) SetShowHint(value bool) {
    StatusBar_SetShowHint(s.instance, value)
}

// SimplePanel
func (s *TStatusBar) SimplePanel() bool {
    return StatusBar_GetSimplePanel(s.instance)
}

// SetSimplePanel
func (s *TStatusBar) SetSimplePanel(value bool) {
    StatusBar_SetSimplePanel(s.instance, value)
}

// SimpleText
func (s *TStatusBar) SimpleText() string {
    return StatusBar_GetSimpleText(s.instance)
}

// SetSimpleText
func (s *TStatusBar) SetSimpleText(value string) {
    StatusBar_SetSimpleText(s.instance, value)
}

// SizeGrip
func (s *TStatusBar) SizeGrip() bool {
    return StatusBar_GetSizeGrip(s.instance)
}

// SetSizeGrip
func (s *TStatusBar) SetSizeGrip(value bool) {
    StatusBar_SetSizeGrip(s.instance, value)
}

// UseSystemFont
func (s *TStatusBar) UseSystemFont() bool {
    return StatusBar_GetUseSystemFont(s.instance)
}

// SetUseSystemFont
func (s *TStatusBar) SetUseSystemFont(value bool) {
    StatusBar_SetUseSystemFont(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TStatusBar) Visible() bool {
    return StatusBar_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TStatusBar) SetVisible(value bool) {
    StatusBar_SetVisible(s.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (s *TStatusBar) StyleElements() TStyleElements {
    return StatusBar_GetStyleElements(s.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (s *TStatusBar) SetStyleElements(value TStyleElements) {
    StatusBar_SetStyleElements(s.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TStatusBar) SetOnClick(fn TNotifyEvent) {
    StatusBar_SetOnClick(s.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (s *TStatusBar) SetOnContextPopup(fn TContextPopupEvent) {
    StatusBar_SetOnContextPopup(s.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (s *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
    StatusBar_SetOnDblClick(s.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (s *TStatusBar) SetOnDragDrop(fn TDragDropEvent) {
    StatusBar_SetOnDragDrop(s.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (s *TStatusBar) SetOnDragOver(fn TDragOverEvent) {
    StatusBar_SetOnDragOver(s.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (s *TStatusBar) SetOnEndDock(fn TEndDragEvent) {
    StatusBar_SetOnEndDock(s.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (s *TStatusBar) SetOnEndDrag(fn TEndDragEvent) {
    StatusBar_SetOnEndDrag(s.instance, fn)
}

// SetOnGesture
func (s *TStatusBar) SetOnGesture(fn TGestureEvent) {
    StatusBar_SetOnGesture(s.instance, fn)
}

// SetOnHint
// CN: 设置鼠标悬停提示事件。
// EN: .
func (s *TStatusBar) SetOnHint(fn TNotifyEvent) {
    StatusBar_SetOnHint(s.instance, fn)
}

// SetOnMouseActivate
func (s *TStatusBar) SetOnMouseActivate(fn TMouseActivateEvent) {
    StatusBar_SetOnMouseActivate(s.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (s *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
    StatusBar_SetOnMouseDown(s.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (s *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
    StatusBar_SetOnMouseEnter(s.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (s *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
    StatusBar_SetOnMouseLeave(s.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (s *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
    StatusBar_SetOnMouseMove(s.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (s *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
    StatusBar_SetOnMouseUp(s.instance, fn)
}

// SetOnResize
// CN: 设置大小被改变事件。
// EN: .
func (s *TStatusBar) SetOnResize(fn TNotifyEvent) {
    StatusBar_SetOnResize(s.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (s *TStatusBar) SetOnStartDock(fn TStartDockEvent) {
    StatusBar_SetOnStartDock(s.instance, fn)
}

// Canvas
// CN: 获取画布。
// EN: .
func (s *TStatusBar) Canvas() *TCanvas {
    return CanvasFromInst(StatusBar_GetCanvas(s.instance))
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (s *TStatusBar) DockClientCount() int32 {
    return StatusBar_GetDockClientCount(s.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (s *TStatusBar) DockSite() bool {
    return StatusBar_GetDockSite(s.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (s *TStatusBar) SetDockSite(value bool) {
    StatusBar_SetDockSite(s.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (s *TStatusBar) AlignDisabled() bool {
    return StatusBar_GetAlignDisabled(s.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (s *TStatusBar) MouseInClient() bool {
    return StatusBar_GetMouseInClient(s.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (s *TStatusBar) VisibleDockClientCount() int32 {
    return StatusBar_GetVisibleDockClientCount(s.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (s *TStatusBar) Brush() *TBrush {
    return BrushFromInst(StatusBar_GetBrush(s.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (s *TStatusBar) ControlCount() int32 {
    return StatusBar_GetControlCount(s.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TStatusBar) Handle() HWND {
    return StatusBar_GetHandle(s.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (s *TStatusBar) ParentWindow() HWND {
    return StatusBar_GetParentWindow(s.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (s *TStatusBar) SetParentWindow(value HWND) {
    StatusBar_SetParentWindow(s.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (s *TStatusBar) TabOrder() TTabOrder {
    return StatusBar_GetTabOrder(s.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (s *TStatusBar) SetTabOrder(value TTabOrder) {
    StatusBar_SetTabOrder(s.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (s *TStatusBar) TabStop() bool {
    return StatusBar_GetTabStop(s.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (s *TStatusBar) SetTabStop(value bool) {
    StatusBar_SetTabStop(s.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (s *TStatusBar) UseDockManager() bool {
    return StatusBar_GetUseDockManager(s.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (s *TStatusBar) SetUseDockManager(value bool) {
    StatusBar_SetUseDockManager(s.instance, value)
}

// BoundsRect
func (s *TStatusBar) BoundsRect() TRect {
    return StatusBar_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TStatusBar) SetBoundsRect(value TRect) {
    StatusBar_SetBoundsRect(s.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (s *TStatusBar) ClientHeight() int32 {
    return StatusBar_GetClientHeight(s.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (s *TStatusBar) SetClientHeight(value int32) {
    StatusBar_SetClientHeight(s.instance, value)
}

// ClientOrigin
func (s *TStatusBar) ClientOrigin() TPoint {
    return StatusBar_GetClientOrigin(s.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TStatusBar) ClientRect() TRect {
    return StatusBar_GetClientRect(s.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TStatusBar) ClientWidth() int32 {
    return StatusBar_GetClientWidth(s.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TStatusBar) SetClientWidth(value int32) {
    StatusBar_SetClientWidth(s.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (s *TStatusBar) ControlState() TControlState {
    return StatusBar_GetControlState(s.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (s *TStatusBar) SetControlState(value TControlState) {
    StatusBar_SetControlState(s.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (s *TStatusBar) ControlStyle() TControlStyle {
    return StatusBar_GetControlStyle(s.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (s *TStatusBar) SetControlStyle(value TControlStyle) {
    StatusBar_SetControlStyle(s.instance, value)
}

// ExplicitLeft
func (s *TStatusBar) ExplicitLeft() int32 {
    return StatusBar_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TStatusBar) ExplicitTop() int32 {
    return StatusBar_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TStatusBar) ExplicitWidth() int32 {
    return StatusBar_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TStatusBar) ExplicitHeight() int32 {
    return StatusBar_GetExplicitHeight(s.instance)
}

// Floating
func (s *TStatusBar) Floating() bool {
    return StatusBar_GetFloating(s.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TStatusBar) Parent() *TWinControl {
    return WinControlFromInst(StatusBar_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TStatusBar) SetParent(value IWinControl) {
    StatusBar_SetParent(s.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (s *TStatusBar) AlignWithMargins() bool {
    return StatusBar_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (s *TStatusBar) SetAlignWithMargins(value bool) {
    StatusBar_SetAlignWithMargins(s.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (s *TStatusBar) Left() int32 {
    return StatusBar_GetLeft(s.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (s *TStatusBar) SetLeft(value int32) {
    StatusBar_SetLeft(s.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TStatusBar) Top() int32 {
    return StatusBar_GetTop(s.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TStatusBar) SetTop(value int32) {
    StatusBar_SetTop(s.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (s *TStatusBar) Width() int32 {
    return StatusBar_GetWidth(s.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (s *TStatusBar) SetWidth(value int32) {
    StatusBar_SetWidth(s.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (s *TStatusBar) Height() int32 {
    return StatusBar_GetHeight(s.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (s *TStatusBar) SetHeight(value int32) {
    StatusBar_SetHeight(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TStatusBar) Cursor() TCursor {
    return StatusBar_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TStatusBar) SetCursor(value TCursor) {
    StatusBar_SetCursor(s.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (s *TStatusBar) Hint() string {
    return StatusBar_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (s *TStatusBar) SetHint(value string) {
    StatusBar_SetHint(s.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (s *TStatusBar) Margins() *TMargins {
    return MarginsFromInst(StatusBar_GetMargins(s.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (s *TStatusBar) SetMargins(value *TMargins) {
    StatusBar_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (s *TStatusBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(StatusBar_GetCustomHint(s.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (s *TStatusBar) SetCustomHint(value IComponent) {
    StatusBar_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TStatusBar) ComponentCount() int32 {
    return StatusBar_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TStatusBar) ComponentIndex() int32 {
    return StatusBar_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TStatusBar) SetComponentIndex(value int32) {
    StatusBar_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TStatusBar) Owner() *TComponent {
    return ComponentFromInst(StatusBar_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TStatusBar) Name() string {
    return StatusBar_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TStatusBar) SetName(value string) {
    StatusBar_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TStatusBar) Tag() int {
    return StatusBar_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TStatusBar) SetTag(value int) {
    StatusBar_SetTag(s.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (s *TStatusBar) DockClients(Index int32) *TControl {
    return ControlFromInst(StatusBar_GetDockClients(s.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (s *TStatusBar) Controls(Index int32) *TControl {
    return ControlFromInst(StatusBar_GetControls(s.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TStatusBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(StatusBar_GetComponents(s.instance, AIndex))
}

