
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

type TGroupBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewGroupBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGroupBox(owner IComponent) *TGroupBox {
    g := new(TGroupBox)
    g.instance = GroupBox_Create(CheckPtr(owner))
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GroupBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func GroupBoxFromInst(inst uintptr) *TGroupBox {
    g := new(TGroupBox)
    g.instance = inst
    g.ptr = unsafe.Pointer(inst)
    return g
}

// GroupBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func GroupBoxFromObj(obj IObject) *TGroupBox {
    g := new(TGroupBox)
    g.instance = CheckPtr(obj)
    g.ptr = unsafe.Pointer(g.instance)
    return g
}

// GroupBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func GroupBoxFromUnsafePointer(ptr unsafe.Pointer) *TGroupBox {
    g := new(TGroupBox)
    g.instance = uintptr(ptr)
    g.ptr = ptr
    return g
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (g *TGroupBox) Free() {
    if g.instance != 0 {
        GroupBox_Free(g.instance)
        g.instance = 0
        g.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGroupBox) Instance() uintptr {
    return g.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGroupBox) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGroupBox) IsValid() bool {
    return g.instance != 0
}

// TGroupBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGroupBoxClass() TClass {
    return GroupBox_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (g *TGroupBox) CanFocus() bool {
    return GroupBox_CanFocus(g.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (g *TGroupBox) ContainsControl(Control IControl) bool {
    return GroupBox_ContainsControl(g.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (g *TGroupBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(GroupBox_ControlAtPos(g.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (g *TGroupBox) DisableAlign() {
    GroupBox_DisableAlign(g.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (g *TGroupBox) EnableAlign() {
    GroupBox_EnableAlign(g.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (g *TGroupBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(GroupBox_FindChildControl(g.instance, ControlName))
}

// FlipChildren
func (g *TGroupBox) FlipChildren(AllLevels bool) {
    GroupBox_FlipChildren(g.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (g *TGroupBox) Focused() bool {
    return GroupBox_Focused(g.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (g *TGroupBox) HandleAllocated() bool {
    return GroupBox_HandleAllocated(g.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (g *TGroupBox) InsertControl(AControl IControl) {
    GroupBox_InsertControl(g.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (g *TGroupBox) Invalidate() {
    GroupBox_Invalidate(g.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (g *TGroupBox) PaintTo(DC HDC, X int32, Y int32) {
    GroupBox_PaintTo(g.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (g *TGroupBox) RemoveControl(AControl IControl) {
    GroupBox_RemoveControl(g.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (g *TGroupBox) Realign() {
    GroupBox_Realign(g.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (g *TGroupBox) Repaint() {
    GroupBox_Repaint(g.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (g *TGroupBox) ScaleBy(M int32, D int32) {
    GroupBox_ScaleBy(g.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (g *TGroupBox) ScrollBy(DeltaX int32, DeltaY int32) {
    GroupBox_ScrollBy(g.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (g *TGroupBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    GroupBox_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (g *TGroupBox) SetFocus() {
    GroupBox_SetFocus(g.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (g *TGroupBox) Update() {
    GroupBox_Update(g.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (g *TGroupBox) UpdateControlState() {
    GroupBox_UpdateControlState(g.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (g *TGroupBox) BringToFront() {
    GroupBox_BringToFront(g.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (g *TGroupBox) ClientToScreen(Point TPoint) TPoint {
    return GroupBox_ClientToScreen(g.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (g *TGroupBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return GroupBox_ClientToParent(g.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (g *TGroupBox) Dragging() bool {
    return GroupBox_Dragging(g.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (g *TGroupBox) HasParent() bool {
    return GroupBox_HasParent(g.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (g *TGroupBox) Hide() {
    GroupBox_Hide(g.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (g *TGroupBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return GroupBox_Perform(g.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (g *TGroupBox) Refresh() {
    GroupBox_Refresh(g.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (g *TGroupBox) ScreenToClient(Point TPoint) TPoint {
    return GroupBox_ScreenToClient(g.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (g *TGroupBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return GroupBox_ParentToClient(g.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (g *TGroupBox) SendToBack() {
    GroupBox_SendToBack(g.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (g *TGroupBox) Show() {
    GroupBox_Show(g.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (g *TGroupBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return GroupBox_GetTextBuf(g.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (g *TGroupBox) GetTextLen() int32 {
    return GroupBox_GetTextLen(g.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (g *TGroupBox) SetTextBuf(Buffer string) {
    GroupBox_SetTextBuf(g.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (g *TGroupBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(GroupBox_FindComponent(g.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (g *TGroupBox) GetNamePath() string {
    return GroupBox_GetNamePath(g.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (g *TGroupBox) Assign(Source IObject) {
    GroupBox_Assign(g.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (g *TGroupBox) DisposeOf() {
    GroupBox_DisposeOf(g.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGroupBox) ClassType() TClass {
    return GroupBox_ClassType(g.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGroupBox) ClassName() string {
    return GroupBox_ClassName(g.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGroupBox) InstanceSize() int32 {
    return GroupBox_InstanceSize(g.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGroupBox) InheritsFrom(AClass TClass) bool {
    return GroupBox_InheritsFrom(g.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGroupBox) Equals(Obj IObject) bool {
    return GroupBox_Equals(g.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGroupBox) GetHashCode() int32 {
    return GroupBox_GetHashCode(g.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (g *TGroupBox) ToString() string {
    return GroupBox_ToString(g.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (g *TGroupBox) Align() TAlign {
    return GroupBox_GetAlign(g.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (g *TGroupBox) SetAlign(value TAlign) {
    GroupBox_SetAlign(g.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (g *TGroupBox) Anchors() TAnchors {
    return GroupBox_GetAnchors(g.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (g *TGroupBox) SetAnchors(value TAnchors) {
    GroupBox_SetAnchors(g.instance, value)
}

// BiDiMode
func (g *TGroupBox) BiDiMode() TBiDiMode {
    return GroupBox_GetBiDiMode(g.instance)
}

// SetBiDiMode
func (g *TGroupBox) SetBiDiMode(value TBiDiMode) {
    GroupBox_SetBiDiMode(g.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (g *TGroupBox) Caption() string {
    return GroupBox_GetCaption(g.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (g *TGroupBox) SetCaption(value string) {
    GroupBox_SetCaption(g.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (g *TGroupBox) Color() TColor {
    return GroupBox_GetColor(g.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (g *TGroupBox) SetColor(value TColor) {
    GroupBox_SetColor(g.instance, value)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (g *TGroupBox) DockSite() bool {
    return GroupBox_GetDockSite(g.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (g *TGroupBox) SetDockSite(value bool) {
    GroupBox_SetDockSite(g.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (g *TGroupBox) DoubleBuffered() bool {
    return GroupBox_GetDoubleBuffered(g.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (g *TGroupBox) SetDoubleBuffered(value bool) {
    GroupBox_SetDoubleBuffered(g.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (g *TGroupBox) DragCursor() TCursor {
    return GroupBox_GetDragCursor(g.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (g *TGroupBox) SetDragCursor(value TCursor) {
    GroupBox_SetDragCursor(g.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (g *TGroupBox) DragKind() TDragKind {
    return GroupBox_GetDragKind(g.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (g *TGroupBox) SetDragKind(value TDragKind) {
    GroupBox_SetDragKind(g.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (g *TGroupBox) DragMode() TDragMode {
    return GroupBox_GetDragMode(g.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (g *TGroupBox) SetDragMode(value TDragMode) {
    GroupBox_SetDragMode(g.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (g *TGroupBox) Enabled() bool {
    return GroupBox_GetEnabled(g.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (g *TGroupBox) SetEnabled(value bool) {
    GroupBox_SetEnabled(g.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (g *TGroupBox) Font() *TFont {
    return FontFromInst(GroupBox_GetFont(g.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (g *TGroupBox) SetFont(value *TFont) {
    GroupBox_SetFont(g.instance, CheckPtr(value))
}

// ParentBackground
func (g *TGroupBox) ParentBackground() bool {
    return GroupBox_GetParentBackground(g.instance)
}

// SetParentBackground
func (g *TGroupBox) SetParentBackground(value bool) {
    GroupBox_SetParentBackground(g.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (g *TGroupBox) ParentColor() bool {
    return GroupBox_GetParentColor(g.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (g *TGroupBox) SetParentColor(value bool) {
    GroupBox_SetParentColor(g.instance, value)
}

// ParentCtl3D
func (g *TGroupBox) ParentCtl3D() bool {
    return GroupBox_GetParentCtl3D(g.instance)
}

// SetParentCtl3D
func (g *TGroupBox) SetParentCtl3D(value bool) {
    GroupBox_SetParentCtl3D(g.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (g *TGroupBox) ParentDoubleBuffered() bool {
    return GroupBox_GetParentDoubleBuffered(g.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (g *TGroupBox) SetParentDoubleBuffered(value bool) {
    GroupBox_SetParentDoubleBuffered(g.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (g *TGroupBox) ParentFont() bool {
    return GroupBox_GetParentFont(g.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (g *TGroupBox) SetParentFont(value bool) {
    GroupBox_SetParentFont(g.instance, value)
}

// ParentShowHint
func (g *TGroupBox) ParentShowHint() bool {
    return GroupBox_GetParentShowHint(g.instance)
}

// SetParentShowHint
func (g *TGroupBox) SetParentShowHint(value bool) {
    GroupBox_SetParentShowHint(g.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (g *TGroupBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(GroupBox_GetPopupMenu(g.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (g *TGroupBox) SetPopupMenu(value IComponent) {
    GroupBox_SetPopupMenu(g.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (g *TGroupBox) ShowHint() bool {
    return GroupBox_GetShowHint(g.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (g *TGroupBox) SetShowHint(value bool) {
    GroupBox_SetShowHint(g.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (g *TGroupBox) TabOrder() TTabOrder {
    return GroupBox_GetTabOrder(g.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (g *TGroupBox) SetTabOrder(value TTabOrder) {
    GroupBox_SetTabOrder(g.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (g *TGroupBox) TabStop() bool {
    return GroupBox_GetTabStop(g.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (g *TGroupBox) SetTabStop(value bool) {
    GroupBox_SetTabStop(g.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (g *TGroupBox) Visible() bool {
    return GroupBox_GetVisible(g.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (g *TGroupBox) SetVisible(value bool) {
    GroupBox_SetVisible(g.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (g *TGroupBox) StyleElements() TStyleElements {
    return GroupBox_GetStyleElements(g.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (g *TGroupBox) SetStyleElements(value TStyleElements) {
    GroupBox_SetStyleElements(g.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (g *TGroupBox) SetOnClick(fn TNotifyEvent) {
    GroupBox_SetOnClick(g.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (g *TGroupBox) SetOnContextPopup(fn TContextPopupEvent) {
    GroupBox_SetOnContextPopup(g.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (g *TGroupBox) SetOnDblClick(fn TNotifyEvent) {
    GroupBox_SetOnDblClick(g.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (g *TGroupBox) SetOnDragDrop(fn TDragDropEvent) {
    GroupBox_SetOnDragDrop(g.instance, fn)
}

// SetOnDockDrop
func (g *TGroupBox) SetOnDockDrop(fn TDockDropEvent) {
    GroupBox_SetOnDockDrop(g.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (g *TGroupBox) SetOnDragOver(fn TDragOverEvent) {
    GroupBox_SetOnDragOver(g.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (g *TGroupBox) SetOnEndDock(fn TEndDragEvent) {
    GroupBox_SetOnEndDock(g.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (g *TGroupBox) SetOnEndDrag(fn TEndDragEvent) {
    GroupBox_SetOnEndDrag(g.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (g *TGroupBox) SetOnEnter(fn TNotifyEvent) {
    GroupBox_SetOnEnter(g.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (g *TGroupBox) SetOnExit(fn TNotifyEvent) {
    GroupBox_SetOnExit(g.instance, fn)
}

// SetOnGesture
func (g *TGroupBox) SetOnGesture(fn TGestureEvent) {
    GroupBox_SetOnGesture(g.instance, fn)
}

// SetOnGetSiteInfo
func (g *TGroupBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    GroupBox_SetOnGetSiteInfo(g.instance, fn)
}

// SetOnMouseActivate
func (g *TGroupBox) SetOnMouseActivate(fn TMouseActivateEvent) {
    GroupBox_SetOnMouseActivate(g.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (g *TGroupBox) SetOnMouseDown(fn TMouseEvent) {
    GroupBox_SetOnMouseDown(g.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (g *TGroupBox) SetOnMouseEnter(fn TNotifyEvent) {
    GroupBox_SetOnMouseEnter(g.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (g *TGroupBox) SetOnMouseLeave(fn TNotifyEvent) {
    GroupBox_SetOnMouseLeave(g.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (g *TGroupBox) SetOnMouseMove(fn TMouseMoveEvent) {
    GroupBox_SetOnMouseMove(g.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (g *TGroupBox) SetOnMouseUp(fn TMouseEvent) {
    GroupBox_SetOnMouseUp(g.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (g *TGroupBox) SetOnStartDock(fn TStartDockEvent) {
    GroupBox_SetOnStartDock(g.instance, fn)
}

// SetOnUnDock
func (g *TGroupBox) SetOnUnDock(fn TUnDockEvent) {
    GroupBox_SetOnUnDock(g.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (g *TGroupBox) DockClientCount() int32 {
    return GroupBox_GetDockClientCount(g.instance)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (g *TGroupBox) AlignDisabled() bool {
    return GroupBox_GetAlignDisabled(g.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (g *TGroupBox) MouseInClient() bool {
    return GroupBox_GetMouseInClient(g.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (g *TGroupBox) VisibleDockClientCount() int32 {
    return GroupBox_GetVisibleDockClientCount(g.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (g *TGroupBox) Brush() *TBrush {
    return BrushFromInst(GroupBox_GetBrush(g.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (g *TGroupBox) ControlCount() int32 {
    return GroupBox_GetControlCount(g.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (g *TGroupBox) Handle() HWND {
    return GroupBox_GetHandle(g.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (g *TGroupBox) ParentWindow() HWND {
    return GroupBox_GetParentWindow(g.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (g *TGroupBox) SetParentWindow(value HWND) {
    GroupBox_SetParentWindow(g.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (g *TGroupBox) UseDockManager() bool {
    return GroupBox_GetUseDockManager(g.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (g *TGroupBox) SetUseDockManager(value bool) {
    GroupBox_SetUseDockManager(g.instance, value)
}

// Action
func (g *TGroupBox) Action() *TAction {
    return ActionFromInst(GroupBox_GetAction(g.instance))
}

// SetAction
func (g *TGroupBox) SetAction(value IComponent) {
    GroupBox_SetAction(g.instance, CheckPtr(value))
}

// BoundsRect
func (g *TGroupBox) BoundsRect() TRect {
    return GroupBox_GetBoundsRect(g.instance)
}

// SetBoundsRect
func (g *TGroupBox) SetBoundsRect(value TRect) {
    GroupBox_SetBoundsRect(g.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (g *TGroupBox) ClientHeight() int32 {
    return GroupBox_GetClientHeight(g.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (g *TGroupBox) SetClientHeight(value int32) {
    GroupBox_SetClientHeight(g.instance, value)
}

// ClientOrigin
func (g *TGroupBox) ClientOrigin() TPoint {
    return GroupBox_GetClientOrigin(g.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (g *TGroupBox) ClientRect() TRect {
    return GroupBox_GetClientRect(g.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (g *TGroupBox) ClientWidth() int32 {
    return GroupBox_GetClientWidth(g.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (g *TGroupBox) SetClientWidth(value int32) {
    GroupBox_SetClientWidth(g.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (g *TGroupBox) ControlState() TControlState {
    return GroupBox_GetControlState(g.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (g *TGroupBox) SetControlState(value TControlState) {
    GroupBox_SetControlState(g.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (g *TGroupBox) ControlStyle() TControlStyle {
    return GroupBox_GetControlStyle(g.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (g *TGroupBox) SetControlStyle(value TControlStyle) {
    GroupBox_SetControlStyle(g.instance, value)
}

// ExplicitLeft
func (g *TGroupBox) ExplicitLeft() int32 {
    return GroupBox_GetExplicitLeft(g.instance)
}

// ExplicitTop
func (g *TGroupBox) ExplicitTop() int32 {
    return GroupBox_GetExplicitTop(g.instance)
}

// ExplicitWidth
func (g *TGroupBox) ExplicitWidth() int32 {
    return GroupBox_GetExplicitWidth(g.instance)
}

// ExplicitHeight
func (g *TGroupBox) ExplicitHeight() int32 {
    return GroupBox_GetExplicitHeight(g.instance)
}

// Floating
func (g *TGroupBox) Floating() bool {
    return GroupBox_GetFloating(g.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (g *TGroupBox) Parent() *TWinControl {
    return WinControlFromInst(GroupBox_GetParent(g.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (g *TGroupBox) SetParent(value IWinControl) {
    GroupBox_SetParent(g.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (g *TGroupBox) AlignWithMargins() bool {
    return GroupBox_GetAlignWithMargins(g.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (g *TGroupBox) SetAlignWithMargins(value bool) {
    GroupBox_SetAlignWithMargins(g.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (g *TGroupBox) Left() int32 {
    return GroupBox_GetLeft(g.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (g *TGroupBox) SetLeft(value int32) {
    GroupBox_SetLeft(g.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (g *TGroupBox) Top() int32 {
    return GroupBox_GetTop(g.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (g *TGroupBox) SetTop(value int32) {
    GroupBox_SetTop(g.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (g *TGroupBox) Width() int32 {
    return GroupBox_GetWidth(g.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (g *TGroupBox) SetWidth(value int32) {
    GroupBox_SetWidth(g.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (g *TGroupBox) Height() int32 {
    return GroupBox_GetHeight(g.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (g *TGroupBox) SetHeight(value int32) {
    GroupBox_SetHeight(g.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (g *TGroupBox) Cursor() TCursor {
    return GroupBox_GetCursor(g.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (g *TGroupBox) SetCursor(value TCursor) {
    GroupBox_SetCursor(g.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (g *TGroupBox) Hint() string {
    return GroupBox_GetHint(g.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (g *TGroupBox) SetHint(value string) {
    GroupBox_SetHint(g.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (g *TGroupBox) Margins() *TMargins {
    return MarginsFromInst(GroupBox_GetMargins(g.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (g *TGroupBox) SetMargins(value *TMargins) {
    GroupBox_SetMargins(g.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (g *TGroupBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(GroupBox_GetCustomHint(g.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (g *TGroupBox) SetCustomHint(value IComponent) {
    GroupBox_SetCustomHint(g.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (g *TGroupBox) ComponentCount() int32 {
    return GroupBox_GetComponentCount(g.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (g *TGroupBox) ComponentIndex() int32 {
    return GroupBox_GetComponentIndex(g.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (g *TGroupBox) SetComponentIndex(value int32) {
    GroupBox_SetComponentIndex(g.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (g *TGroupBox) Owner() *TComponent {
    return ComponentFromInst(GroupBox_GetOwner(g.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (g *TGroupBox) Name() string {
    return GroupBox_GetName(g.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (g *TGroupBox) SetName(value string) {
    GroupBox_SetName(g.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (g *TGroupBox) Tag() int {
    return GroupBox_GetTag(g.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (g *TGroupBox) SetTag(value int) {
    GroupBox_SetTag(g.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (g *TGroupBox) DockClients(Index int32) *TControl {
    return ControlFromInst(GroupBox_GetDockClients(g.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (g *TGroupBox) Controls(Index int32) *TControl {
    return ControlFromInst(GroupBox_GetControls(g.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (g *TGroupBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(GroupBox_GetComponents(g.instance, AIndex))
}

