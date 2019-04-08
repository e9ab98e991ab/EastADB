
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

type THeaderControl struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewHeaderControl
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewHeaderControl(owner IComponent) *THeaderControl {
    h := new(THeaderControl)
    h.instance = HeaderControl_Create(CheckPtr(owner))
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HeaderControlFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func HeaderControlFromInst(inst uintptr) *THeaderControl {
    h := new(THeaderControl)
    h.instance = inst
    h.ptr = unsafe.Pointer(inst)
    return h
}

// HeaderControlFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func HeaderControlFromObj(obj IObject) *THeaderControl {
    h := new(THeaderControl)
    h.instance = CheckPtr(obj)
    h.ptr = unsafe.Pointer(h.instance)
    return h
}

// HeaderControlFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func HeaderControlFromUnsafePointer(ptr unsafe.Pointer) *THeaderControl {
    h := new(THeaderControl)
    h.instance = uintptr(ptr)
    h.ptr = ptr
    return h
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (h *THeaderControl) Free() {
    if h.instance != 0 {
        HeaderControl_Free(h.instance)
        h.instance = 0
        h.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (h *THeaderControl) Instance() uintptr {
    return h.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (h *THeaderControl) UnsafeAddr() unsafe.Pointer {
    return h.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (h *THeaderControl) IsValid() bool {
    return h.instance != 0
}

// THeaderControlClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func THeaderControlClass() TClass {
    return HeaderControl_StaticClassType()
}

// FlipChildren
func (h *THeaderControl) FlipChildren(AllLevels bool) {
    HeaderControl_FlipChildren(h.instance, AllLevels)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (h *THeaderControl) CanFocus() bool {
    return HeaderControl_CanFocus(h.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (h *THeaderControl) ContainsControl(Control IControl) bool {
    return HeaderControl_ContainsControl(h.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (h *THeaderControl) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(HeaderControl_ControlAtPos(h.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (h *THeaderControl) DisableAlign() {
    HeaderControl_DisableAlign(h.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (h *THeaderControl) EnableAlign() {
    HeaderControl_EnableAlign(h.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (h *THeaderControl) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(HeaderControl_FindChildControl(h.instance, ControlName))
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (h *THeaderControl) Focused() bool {
    return HeaderControl_Focused(h.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (h *THeaderControl) HandleAllocated() bool {
    return HeaderControl_HandleAllocated(h.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (h *THeaderControl) InsertControl(AControl IControl) {
    HeaderControl_InsertControl(h.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (h *THeaderControl) Invalidate() {
    HeaderControl_Invalidate(h.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (h *THeaderControl) PaintTo(DC HDC, X int32, Y int32) {
    HeaderControl_PaintTo(h.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (h *THeaderControl) RemoveControl(AControl IControl) {
    HeaderControl_RemoveControl(h.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (h *THeaderControl) Realign() {
    HeaderControl_Realign(h.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (h *THeaderControl) Repaint() {
    HeaderControl_Repaint(h.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (h *THeaderControl) ScaleBy(M int32, D int32) {
    HeaderControl_ScaleBy(h.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (h *THeaderControl) ScrollBy(DeltaX int32, DeltaY int32) {
    HeaderControl_ScrollBy(h.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (h *THeaderControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    HeaderControl_SetBounds(h.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (h *THeaderControl) SetFocus() {
    HeaderControl_SetFocus(h.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (h *THeaderControl) Update() {
    HeaderControl_Update(h.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (h *THeaderControl) UpdateControlState() {
    HeaderControl_UpdateControlState(h.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (h *THeaderControl) BringToFront() {
    HeaderControl_BringToFront(h.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (h *THeaderControl) ClientToScreen(Point TPoint) TPoint {
    return HeaderControl_ClientToScreen(h.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (h *THeaderControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return HeaderControl_ClientToParent(h.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (h *THeaderControl) Dragging() bool {
    return HeaderControl_Dragging(h.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (h *THeaderControl) HasParent() bool {
    return HeaderControl_HasParent(h.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (h *THeaderControl) Hide() {
    HeaderControl_Hide(h.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (h *THeaderControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return HeaderControl_Perform(h.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (h *THeaderControl) Refresh() {
    HeaderControl_Refresh(h.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (h *THeaderControl) ScreenToClient(Point TPoint) TPoint {
    return HeaderControl_ScreenToClient(h.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (h *THeaderControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return HeaderControl_ParentToClient(h.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (h *THeaderControl) SendToBack() {
    HeaderControl_SendToBack(h.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (h *THeaderControl) Show() {
    HeaderControl_Show(h.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (h *THeaderControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return HeaderControl_GetTextBuf(h.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (h *THeaderControl) GetTextLen() int32 {
    return HeaderControl_GetTextLen(h.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (h *THeaderControl) SetTextBuf(Buffer string) {
    HeaderControl_SetTextBuf(h.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (h *THeaderControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(HeaderControl_FindComponent(h.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (h *THeaderControl) GetNamePath() string {
    return HeaderControl_GetNamePath(h.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (h *THeaderControl) Assign(Source IObject) {
    HeaderControl_Assign(h.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (h *THeaderControl) DisposeOf() {
    HeaderControl_DisposeOf(h.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (h *THeaderControl) ClassType() TClass {
    return HeaderControl_ClassType(h.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (h *THeaderControl) ClassName() string {
    return HeaderControl_ClassName(h.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (h *THeaderControl) InstanceSize() int32 {
    return HeaderControl_InstanceSize(h.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (h *THeaderControl) InheritsFrom(AClass TClass) bool {
    return HeaderControl_InheritsFrom(h.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (h *THeaderControl) Equals(Obj IObject) bool {
    return HeaderControl_Equals(h.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (h *THeaderControl) GetHashCode() int32 {
    return HeaderControl_GetHashCode(h.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (h *THeaderControl) ToString() string {
    return HeaderControl_ToString(h.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (h *THeaderControl) Align() TAlign {
    return HeaderControl_GetAlign(h.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (h *THeaderControl) SetAlign(value TAlign) {
    HeaderControl_SetAlign(h.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (h *THeaderControl) Anchors() TAnchors {
    return HeaderControl_GetAnchors(h.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (h *THeaderControl) SetAnchors(value TAnchors) {
    HeaderControl_SetAnchors(h.instance, value)
}

// BiDiMode
func (h *THeaderControl) BiDiMode() TBiDiMode {
    return HeaderControl_GetBiDiMode(h.instance)
}

// SetBiDiMode
func (h *THeaderControl) SetBiDiMode(value TBiDiMode) {
    HeaderControl_SetBiDiMode(h.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (h *THeaderControl) BorderWidth() int32 {
    return HeaderControl_GetBorderWidth(h.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (h *THeaderControl) SetBorderWidth(value int32) {
    HeaderControl_SetBorderWidth(h.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (h *THeaderControl) DoubleBuffered() bool {
    return HeaderControl_GetDoubleBuffered(h.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (h *THeaderControl) SetDoubleBuffered(value bool) {
    HeaderControl_SetDoubleBuffered(h.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (h *THeaderControl) DragCursor() TCursor {
    return HeaderControl_GetDragCursor(h.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (h *THeaderControl) SetDragCursor(value TCursor) {
    HeaderControl_SetDragCursor(h.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (h *THeaderControl) DragKind() TDragKind {
    return HeaderControl_GetDragKind(h.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (h *THeaderControl) SetDragKind(value TDragKind) {
    HeaderControl_SetDragKind(h.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (h *THeaderControl) DragMode() TDragMode {
    return HeaderControl_GetDragMode(h.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (h *THeaderControl) SetDragMode(value TDragMode) {
    HeaderControl_SetDragMode(h.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (h *THeaderControl) Enabled() bool {
    return HeaderControl_GetEnabled(h.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (h *THeaderControl) SetEnabled(value bool) {
    HeaderControl_SetEnabled(h.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (h *THeaderControl) Font() *TFont {
    return FontFromInst(HeaderControl_GetFont(h.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (h *THeaderControl) SetFont(value *TFont) {
    HeaderControl_SetFont(h.instance, CheckPtr(value))
}

// FullDrag
func (h *THeaderControl) FullDrag() bool {
    return HeaderControl_GetFullDrag(h.instance)
}

// SetFullDrag
func (h *THeaderControl) SetFullDrag(value bool) {
    HeaderControl_SetFullDrag(h.instance, value)
}

// HotTrack
func (h *THeaderControl) HotTrack() bool {
    return HeaderControl_GetHotTrack(h.instance)
}

// SetHotTrack
func (h *THeaderControl) SetHotTrack(value bool) {
    HeaderControl_SetHotTrack(h.instance, value)
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (h *THeaderControl) Images() *TImageList {
    return ImageListFromInst(HeaderControl_GetImages(h.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (h *THeaderControl) SetImages(value IComponent) {
    HeaderControl_SetImages(h.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (h *THeaderControl) ShowHint() bool {
    return HeaderControl_GetShowHint(h.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (h *THeaderControl) SetShowHint(value bool) {
    HeaderControl_SetShowHint(h.instance, value)
}

// Style
func (h *THeaderControl) Style() THeaderStyle {
    return HeaderControl_GetStyle(h.instance)
}

// SetStyle
func (h *THeaderControl) SetStyle(value THeaderStyle) {
    HeaderControl_SetStyle(h.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (h *THeaderControl) ParentDoubleBuffered() bool {
    return HeaderControl_GetParentDoubleBuffered(h.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (h *THeaderControl) SetParentDoubleBuffered(value bool) {
    HeaderControl_SetParentDoubleBuffered(h.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (h *THeaderControl) ParentFont() bool {
    return HeaderControl_GetParentFont(h.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (h *THeaderControl) SetParentFont(value bool) {
    HeaderControl_SetParentFont(h.instance, value)
}

// ParentShowHint
func (h *THeaderControl) ParentShowHint() bool {
    return HeaderControl_GetParentShowHint(h.instance)
}

// SetParentShowHint
func (h *THeaderControl) SetParentShowHint(value bool) {
    HeaderControl_SetParentShowHint(h.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (h *THeaderControl) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(HeaderControl_GetPopupMenu(h.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (h *THeaderControl) SetPopupMenu(value IComponent) {
    HeaderControl_SetPopupMenu(h.instance, CheckPtr(value))
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (h *THeaderControl) Visible() bool {
    return HeaderControl_GetVisible(h.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (h *THeaderControl) SetVisible(value bool) {
    HeaderControl_SetVisible(h.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (h *THeaderControl) StyleElements() TStyleElements {
    return HeaderControl_GetStyleElements(h.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (h *THeaderControl) SetStyleElements(value TStyleElements) {
    HeaderControl_SetStyleElements(h.instance, value)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (h *THeaderControl) SetOnContextPopup(fn TContextPopupEvent) {
    HeaderControl_SetOnContextPopup(h.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (h *THeaderControl) SetOnDragDrop(fn TDragDropEvent) {
    HeaderControl_SetOnDragDrop(h.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (h *THeaderControl) SetOnDragOver(fn TDragOverEvent) {
    HeaderControl_SetOnDragOver(h.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (h *THeaderControl) SetOnEndDock(fn TEndDragEvent) {
    HeaderControl_SetOnEndDock(h.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (h *THeaderControl) SetOnEndDrag(fn TEndDragEvent) {
    HeaderControl_SetOnEndDrag(h.instance, fn)
}

// SetOnGesture
func (h *THeaderControl) SetOnGesture(fn TGestureEvent) {
    HeaderControl_SetOnGesture(h.instance, fn)
}

// SetOnMouseActivate
func (h *THeaderControl) SetOnMouseActivate(fn TMouseActivateEvent) {
    HeaderControl_SetOnMouseActivate(h.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (h *THeaderControl) SetOnMouseDown(fn TMouseEvent) {
    HeaderControl_SetOnMouseDown(h.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (h *THeaderControl) SetOnMouseEnter(fn TNotifyEvent) {
    HeaderControl_SetOnMouseEnter(h.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (h *THeaderControl) SetOnMouseLeave(fn TNotifyEvent) {
    HeaderControl_SetOnMouseLeave(h.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (h *THeaderControl) SetOnMouseMove(fn TMouseMoveEvent) {
    HeaderControl_SetOnMouseMove(h.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (h *THeaderControl) SetOnMouseUp(fn TMouseEvent) {
    HeaderControl_SetOnMouseUp(h.instance, fn)
}

// SetOnResize
// CN: 设置大小被改变事件。
// EN: .
func (h *THeaderControl) SetOnResize(fn TNotifyEvent) {
    HeaderControl_SetOnResize(h.instance, fn)
}

// SetOnDrawSection
func (h *THeaderControl) SetOnDrawSection(fn TDrawSectionEvent) {
    HeaderControl_SetOnDrawSection(h.instance, fn)
}

// SetOnSectionClick
func (h *THeaderControl) SetOnSectionClick(fn TSectionNotifyEvent) {
    HeaderControl_SetOnSectionClick(h.instance, fn)
}

// SetOnSectionResize
func (h *THeaderControl) SetOnSectionResize(fn TSectionNotifyEvent) {
    HeaderControl_SetOnSectionResize(h.instance, fn)
}

// SetOnSectionTrack
func (h *THeaderControl) SetOnSectionTrack(fn TSectionTrackEvent) {
    HeaderControl_SetOnSectionTrack(h.instance, fn)
}

// SetOnSectionDrag
func (h *THeaderControl) SetOnSectionDrag(fn TSectionDragEvent) {
    HeaderControl_SetOnSectionDrag(h.instance, fn)
}

// SetOnSectionEndDrag
func (h *THeaderControl) SetOnSectionEndDrag(fn TNotifyEvent) {
    HeaderControl_SetOnSectionEndDrag(h.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (h *THeaderControl) SetOnStartDock(fn TStartDockEvent) {
    HeaderControl_SetOnStartDock(h.instance, fn)
}

// Canvas
// CN: 获取画布。
// EN: .
func (h *THeaderControl) Canvas() *TCanvas {
    return CanvasFromInst(HeaderControl_GetCanvas(h.instance))
}

// SetOnSectionCheck
func (h *THeaderControl) SetOnSectionCheck(fn TCustomSectionNotifyEvent) {
    HeaderControl_SetOnSectionCheck(h.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (h *THeaderControl) DockClientCount() int32 {
    return HeaderControl_GetDockClientCount(h.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (h *THeaderControl) DockSite() bool {
    return HeaderControl_GetDockSite(h.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (h *THeaderControl) SetDockSite(value bool) {
    HeaderControl_SetDockSite(h.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (h *THeaderControl) AlignDisabled() bool {
    return HeaderControl_GetAlignDisabled(h.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (h *THeaderControl) MouseInClient() bool {
    return HeaderControl_GetMouseInClient(h.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (h *THeaderControl) VisibleDockClientCount() int32 {
    return HeaderControl_GetVisibleDockClientCount(h.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (h *THeaderControl) Brush() *TBrush {
    return BrushFromInst(HeaderControl_GetBrush(h.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (h *THeaderControl) ControlCount() int32 {
    return HeaderControl_GetControlCount(h.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (h *THeaderControl) Handle() HWND {
    return HeaderControl_GetHandle(h.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (h *THeaderControl) ParentWindow() HWND {
    return HeaderControl_GetParentWindow(h.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (h *THeaderControl) SetParentWindow(value HWND) {
    HeaderControl_SetParentWindow(h.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (h *THeaderControl) TabOrder() TTabOrder {
    return HeaderControl_GetTabOrder(h.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (h *THeaderControl) SetTabOrder(value TTabOrder) {
    HeaderControl_SetTabOrder(h.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (h *THeaderControl) TabStop() bool {
    return HeaderControl_GetTabStop(h.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (h *THeaderControl) SetTabStop(value bool) {
    HeaderControl_SetTabStop(h.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (h *THeaderControl) UseDockManager() bool {
    return HeaderControl_GetUseDockManager(h.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (h *THeaderControl) SetUseDockManager(value bool) {
    HeaderControl_SetUseDockManager(h.instance, value)
}

// Action
func (h *THeaderControl) Action() *TAction {
    return ActionFromInst(HeaderControl_GetAction(h.instance))
}

// SetAction
func (h *THeaderControl) SetAction(value IComponent) {
    HeaderControl_SetAction(h.instance, CheckPtr(value))
}

// BoundsRect
func (h *THeaderControl) BoundsRect() TRect {
    return HeaderControl_GetBoundsRect(h.instance)
}

// SetBoundsRect
func (h *THeaderControl) SetBoundsRect(value TRect) {
    HeaderControl_SetBoundsRect(h.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (h *THeaderControl) ClientHeight() int32 {
    return HeaderControl_GetClientHeight(h.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (h *THeaderControl) SetClientHeight(value int32) {
    HeaderControl_SetClientHeight(h.instance, value)
}

// ClientOrigin
func (h *THeaderControl) ClientOrigin() TPoint {
    return HeaderControl_GetClientOrigin(h.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (h *THeaderControl) ClientRect() TRect {
    return HeaderControl_GetClientRect(h.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (h *THeaderControl) ClientWidth() int32 {
    return HeaderControl_GetClientWidth(h.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (h *THeaderControl) SetClientWidth(value int32) {
    HeaderControl_SetClientWidth(h.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (h *THeaderControl) ControlState() TControlState {
    return HeaderControl_GetControlState(h.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (h *THeaderControl) SetControlState(value TControlState) {
    HeaderControl_SetControlState(h.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (h *THeaderControl) ControlStyle() TControlStyle {
    return HeaderControl_GetControlStyle(h.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (h *THeaderControl) SetControlStyle(value TControlStyle) {
    HeaderControl_SetControlStyle(h.instance, value)
}

// ExplicitLeft
func (h *THeaderControl) ExplicitLeft() int32 {
    return HeaderControl_GetExplicitLeft(h.instance)
}

// ExplicitTop
func (h *THeaderControl) ExplicitTop() int32 {
    return HeaderControl_GetExplicitTop(h.instance)
}

// ExplicitWidth
func (h *THeaderControl) ExplicitWidth() int32 {
    return HeaderControl_GetExplicitWidth(h.instance)
}

// ExplicitHeight
func (h *THeaderControl) ExplicitHeight() int32 {
    return HeaderControl_GetExplicitHeight(h.instance)
}

// Floating
func (h *THeaderControl) Floating() bool {
    return HeaderControl_GetFloating(h.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (h *THeaderControl) Parent() *TWinControl {
    return WinControlFromInst(HeaderControl_GetParent(h.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (h *THeaderControl) SetParent(value IWinControl) {
    HeaderControl_SetParent(h.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (h *THeaderControl) AlignWithMargins() bool {
    return HeaderControl_GetAlignWithMargins(h.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (h *THeaderControl) SetAlignWithMargins(value bool) {
    HeaderControl_SetAlignWithMargins(h.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (h *THeaderControl) Left() int32 {
    return HeaderControl_GetLeft(h.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (h *THeaderControl) SetLeft(value int32) {
    HeaderControl_SetLeft(h.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (h *THeaderControl) Top() int32 {
    return HeaderControl_GetTop(h.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (h *THeaderControl) SetTop(value int32) {
    HeaderControl_SetTop(h.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (h *THeaderControl) Width() int32 {
    return HeaderControl_GetWidth(h.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (h *THeaderControl) SetWidth(value int32) {
    HeaderControl_SetWidth(h.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (h *THeaderControl) Height() int32 {
    return HeaderControl_GetHeight(h.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (h *THeaderControl) SetHeight(value int32) {
    HeaderControl_SetHeight(h.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (h *THeaderControl) Cursor() TCursor {
    return HeaderControl_GetCursor(h.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (h *THeaderControl) SetCursor(value TCursor) {
    HeaderControl_SetCursor(h.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (h *THeaderControl) Hint() string {
    return HeaderControl_GetHint(h.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (h *THeaderControl) SetHint(value string) {
    HeaderControl_SetHint(h.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (h *THeaderControl) Margins() *TMargins {
    return MarginsFromInst(HeaderControl_GetMargins(h.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (h *THeaderControl) SetMargins(value *TMargins) {
    HeaderControl_SetMargins(h.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (h *THeaderControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(HeaderControl_GetCustomHint(h.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (h *THeaderControl) SetCustomHint(value IComponent) {
    HeaderControl_SetCustomHint(h.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (h *THeaderControl) ComponentCount() int32 {
    return HeaderControl_GetComponentCount(h.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (h *THeaderControl) ComponentIndex() int32 {
    return HeaderControl_GetComponentIndex(h.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (h *THeaderControl) SetComponentIndex(value int32) {
    HeaderControl_SetComponentIndex(h.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (h *THeaderControl) Owner() *TComponent {
    return ComponentFromInst(HeaderControl_GetOwner(h.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (h *THeaderControl) Name() string {
    return HeaderControl_GetName(h.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (h *THeaderControl) SetName(value string) {
    HeaderControl_SetName(h.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (h *THeaderControl) Tag() int {
    return HeaderControl_GetTag(h.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (h *THeaderControl) SetTag(value int) {
    HeaderControl_SetTag(h.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (h *THeaderControl) DockClients(Index int32) *TControl {
    return ControlFromInst(HeaderControl_GetDockClients(h.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (h *THeaderControl) Controls(Index int32) *TControl {
    return ControlFromInst(HeaderControl_GetControls(h.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (h *THeaderControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(HeaderControl_GetComponents(h.instance, AIndex))
}

