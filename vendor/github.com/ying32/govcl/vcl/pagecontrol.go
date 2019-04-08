
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

type TPageControl struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPageControl
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPageControl(owner IComponent) *TPageControl {
    p := new(TPageControl)
    p.instance = PageControl_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PageControlFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PageControlFromInst(inst uintptr) *TPageControl {
    p := new(TPageControl)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PageControlFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PageControlFromObj(obj IObject) *TPageControl {
    p := new(TPageControl)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PageControlFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PageControlFromUnsafePointer(ptr unsafe.Pointer) *TPageControl {
    p := new(TPageControl)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPageControl) Free() {
    if p.instance != 0 {
        PageControl_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPageControl) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPageControl) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPageControl) IsValid() bool {
    return p.instance != 0
}

// TPageControlClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPageControlClass() TClass {
    return PageControl_StaticClassType()
}

// SelectNextPage
func (p *TPageControl) SelectNextPage(GoForward bool, CheckTabVisible bool) {
    PageControl_SelectNextPage(p.instance, GoForward , CheckTabVisible)
}

// TabRect
func (p *TPageControl) TabRect(Index int32) TRect {
    return PageControl_TabRect(p.instance, Index)
}

// RowCount
func (p *TPageControl) RowCount() int32 {
    return PageControl_RowCount(p.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (p *TPageControl) CanFocus() bool {
    return PageControl_CanFocus(p.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (p *TPageControl) ContainsControl(Control IControl) bool {
    return PageControl_ContainsControl(p.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (p *TPageControl) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(PageControl_ControlAtPos(p.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (p *TPageControl) DisableAlign() {
    PageControl_DisableAlign(p.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (p *TPageControl) EnableAlign() {
    PageControl_EnableAlign(p.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (p *TPageControl) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(PageControl_FindChildControl(p.instance, ControlName))
}

// FlipChildren
func (p *TPageControl) FlipChildren(AllLevels bool) {
    PageControl_FlipChildren(p.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (p *TPageControl) Focused() bool {
    return PageControl_Focused(p.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (p *TPageControl) HandleAllocated() bool {
    return PageControl_HandleAllocated(p.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (p *TPageControl) InsertControl(AControl IControl) {
    PageControl_InsertControl(p.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (p *TPageControl) Invalidate() {
    PageControl_Invalidate(p.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (p *TPageControl) PaintTo(DC HDC, X int32, Y int32) {
    PageControl_PaintTo(p.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (p *TPageControl) RemoveControl(AControl IControl) {
    PageControl_RemoveControl(p.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (p *TPageControl) Realign() {
    PageControl_Realign(p.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (p *TPageControl) Repaint() {
    PageControl_Repaint(p.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (p *TPageControl) ScaleBy(M int32, D int32) {
    PageControl_ScaleBy(p.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (p *TPageControl) ScrollBy(DeltaX int32, DeltaY int32) {
    PageControl_ScrollBy(p.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (p *TPageControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    PageControl_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (p *TPageControl) SetFocus() {
    PageControl_SetFocus(p.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (p *TPageControl) Update() {
    PageControl_Update(p.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (p *TPageControl) UpdateControlState() {
    PageControl_UpdateControlState(p.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (p *TPageControl) BringToFront() {
    PageControl_BringToFront(p.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (p *TPageControl) ClientToScreen(Point TPoint) TPoint {
    return PageControl_ClientToScreen(p.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (p *TPageControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return PageControl_ClientToParent(p.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (p *TPageControl) Dragging() bool {
    return PageControl_Dragging(p.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPageControl) HasParent() bool {
    return PageControl_HasParent(p.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (p *TPageControl) Hide() {
    PageControl_Hide(p.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (p *TPageControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return PageControl_Perform(p.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (p *TPageControl) Refresh() {
    PageControl_Refresh(p.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (p *TPageControl) ScreenToClient(Point TPoint) TPoint {
    return PageControl_ScreenToClient(p.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (p *TPageControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return PageControl_ParentToClient(p.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (p *TPageControl) SendToBack() {
    PageControl_SendToBack(p.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (p *TPageControl) Show() {
    PageControl_Show(p.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (p *TPageControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    return PageControl_GetTextBuf(p.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (p *TPageControl) GetTextLen() int32 {
    return PageControl_GetTextLen(p.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (p *TPageControl) SetTextBuf(Buffer string) {
    PageControl_SetTextBuf(p.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPageControl) FindComponent(AName string) *TComponent {
    return ComponentFromInst(PageControl_FindComponent(p.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPageControl) GetNamePath() string {
    return PageControl_GetNamePath(p.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPageControl) Assign(Source IObject) {
    PageControl_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPageControl) DisposeOf() {
    PageControl_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPageControl) ClassType() TClass {
    return PageControl_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPageControl) ClassName() string {
    return PageControl_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPageControl) InstanceSize() int32 {
    return PageControl_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPageControl) InheritsFrom(AClass TClass) bool {
    return PageControl_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPageControl) Equals(Obj IObject) bool {
    return PageControl_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPageControl) GetHashCode() int32 {
    return PageControl_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPageControl) ToString() string {
    return PageControl_ToString(p.instance)
}

// ActivePageIndex
func (p *TPageControl) ActivePageIndex() int32 {
    return PageControl_GetActivePageIndex(p.instance)
}

// SetActivePageIndex
func (p *TPageControl) SetActivePageIndex(value int32) {
    PageControl_SetActivePageIndex(p.instance, value)
}

// PageCount
func (p *TPageControl) PageCount() int32 {
    return PageControl_GetPageCount(p.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (p *TPageControl) Align() TAlign {
    return PageControl_GetAlign(p.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (p *TPageControl) SetAlign(value TAlign) {
    PageControl_SetAlign(p.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (p *TPageControl) Anchors() TAnchors {
    return PageControl_GetAnchors(p.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (p *TPageControl) SetAnchors(value TAnchors) {
    PageControl_SetAnchors(p.instance, value)
}

// BiDiMode
func (p *TPageControl) BiDiMode() TBiDiMode {
    return PageControl_GetBiDiMode(p.instance)
}

// SetBiDiMode
func (p *TPageControl) SetBiDiMode(value TBiDiMode) {
    PageControl_SetBiDiMode(p.instance, value)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (p *TPageControl) DockSite() bool {
    return PageControl_GetDockSite(p.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (p *TPageControl) SetDockSite(value bool) {
    PageControl_SetDockSite(p.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (p *TPageControl) DoubleBuffered() bool {
    return PageControl_GetDoubleBuffered(p.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (p *TPageControl) SetDoubleBuffered(value bool) {
    PageControl_SetDoubleBuffered(p.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (p *TPageControl) DragCursor() TCursor {
    return PageControl_GetDragCursor(p.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (p *TPageControl) SetDragCursor(value TCursor) {
    PageControl_SetDragCursor(p.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (p *TPageControl) DragKind() TDragKind {
    return PageControl_GetDragKind(p.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (p *TPageControl) SetDragKind(value TDragKind) {
    PageControl_SetDragKind(p.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (p *TPageControl) DragMode() TDragMode {
    return PageControl_GetDragMode(p.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (p *TPageControl) SetDragMode(value TDragMode) {
    PageControl_SetDragMode(p.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (p *TPageControl) Enabled() bool {
    return PageControl_GetEnabled(p.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (p *TPageControl) SetEnabled(value bool) {
    PageControl_SetEnabled(p.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (p *TPageControl) Font() *TFont {
    return FontFromInst(PageControl_GetFont(p.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (p *TPageControl) SetFont(value *TFont) {
    PageControl_SetFont(p.instance, CheckPtr(value))
}

// HotTrack
func (p *TPageControl) HotTrack() bool {
    return PageControl_GetHotTrack(p.instance)
}

// SetHotTrack
func (p *TPageControl) SetHotTrack(value bool) {
    PageControl_SetHotTrack(p.instance, value)
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (p *TPageControl) Images() *TImageList {
    return ImageListFromInst(PageControl_GetImages(p.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (p *TPageControl) SetImages(value IComponent) {
    PageControl_SetImages(p.instance, CheckPtr(value))
}

// MultiLine
func (p *TPageControl) MultiLine() bool {
    return PageControl_GetMultiLine(p.instance)
}

// SetMultiLine
func (p *TPageControl) SetMultiLine(value bool) {
    PageControl_SetMultiLine(p.instance, value)
}

// OwnerDraw
func (p *TPageControl) OwnerDraw() bool {
    return PageControl_GetOwnerDraw(p.instance)
}

// SetOwnerDraw
func (p *TPageControl) SetOwnerDraw(value bool) {
    PageControl_SetOwnerDraw(p.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (p *TPageControl) ParentDoubleBuffered() bool {
    return PageControl_GetParentDoubleBuffered(p.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (p *TPageControl) SetParentDoubleBuffered(value bool) {
    PageControl_SetParentDoubleBuffered(p.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (p *TPageControl) ParentFont() bool {
    return PageControl_GetParentFont(p.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (p *TPageControl) SetParentFont(value bool) {
    PageControl_SetParentFont(p.instance, value)
}

// ParentShowHint
func (p *TPageControl) ParentShowHint() bool {
    return PageControl_GetParentShowHint(p.instance)
}

// SetParentShowHint
func (p *TPageControl) SetParentShowHint(value bool) {
    PageControl_SetParentShowHint(p.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (p *TPageControl) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(PageControl_GetPopupMenu(p.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (p *TPageControl) SetPopupMenu(value IComponent) {
    PageControl_SetPopupMenu(p.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (p *TPageControl) ShowHint() bool {
    return PageControl_GetShowHint(p.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (p *TPageControl) SetShowHint(value bool) {
    PageControl_SetShowHint(p.instance, value)
}

// Style
func (p *TPageControl) Style() TTabStyle {
    return PageControl_GetStyle(p.instance)
}

// SetStyle
func (p *TPageControl) SetStyle(value TTabStyle) {
    PageControl_SetStyle(p.instance, value)
}

// TabHeight
func (p *TPageControl) TabHeight() int16 {
    return PageControl_GetTabHeight(p.instance)
}

// SetTabHeight
func (p *TPageControl) SetTabHeight(value int16) {
    PageControl_SetTabHeight(p.instance, value)
}

// TabIndex
func (p *TPageControl) TabIndex() int32 {
    return PageControl_GetTabIndex(p.instance)
}

// SetTabIndex
func (p *TPageControl) SetTabIndex(value int32) {
    PageControl_SetTabIndex(p.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (p *TPageControl) TabOrder() TTabOrder {
    return PageControl_GetTabOrder(p.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (p *TPageControl) SetTabOrder(value TTabOrder) {
    PageControl_SetTabOrder(p.instance, value)
}

// TabPosition
func (p *TPageControl) TabPosition() TTabPosition {
    return PageControl_GetTabPosition(p.instance)
}

// SetTabPosition
func (p *TPageControl) SetTabPosition(value TTabPosition) {
    PageControl_SetTabPosition(p.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (p *TPageControl) TabStop() bool {
    return PageControl_GetTabStop(p.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (p *TPageControl) SetTabStop(value bool) {
    PageControl_SetTabStop(p.instance, value)
}

// TabWidth
func (p *TPageControl) TabWidth() int16 {
    return PageControl_GetTabWidth(p.instance)
}

// SetTabWidth
func (p *TPageControl) SetTabWidth(value int16) {
    PageControl_SetTabWidth(p.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (p *TPageControl) Visible() bool {
    return PageControl_GetVisible(p.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (p *TPageControl) SetVisible(value bool) {
    PageControl_SetVisible(p.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (p *TPageControl) StyleElements() TStyleElements {
    return PageControl_GetStyleElements(p.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (p *TPageControl) SetStyleElements(value TStyleElements) {
    PageControl_SetStyleElements(p.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (p *TPageControl) SetOnChange(fn TNotifyEvent) {
    PageControl_SetOnChange(p.instance, fn)
}

// SetOnChanging
func (p *TPageControl) SetOnChanging(fn TTabChangingEvent) {
    PageControl_SetOnChanging(p.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (p *TPageControl) SetOnContextPopup(fn TContextPopupEvent) {
    PageControl_SetOnContextPopup(p.instance, fn)
}

// SetOnDockDrop
func (p *TPageControl) SetOnDockDrop(fn TDockDropEvent) {
    PageControl_SetOnDockDrop(p.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (p *TPageControl) SetOnDragDrop(fn TDragDropEvent) {
    PageControl_SetOnDragDrop(p.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (p *TPageControl) SetOnDragOver(fn TDragOverEvent) {
    PageControl_SetOnDragOver(p.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (p *TPageControl) SetOnEndDock(fn TEndDragEvent) {
    PageControl_SetOnEndDock(p.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (p *TPageControl) SetOnEndDrag(fn TEndDragEvent) {
    PageControl_SetOnEndDrag(p.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (p *TPageControl) SetOnEnter(fn TNotifyEvent) {
    PageControl_SetOnEnter(p.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (p *TPageControl) SetOnExit(fn TNotifyEvent) {
    PageControl_SetOnExit(p.instance, fn)
}

// SetOnGesture
func (p *TPageControl) SetOnGesture(fn TGestureEvent) {
    PageControl_SetOnGesture(p.instance, fn)
}

// SetOnGetImageIndex
func (p *TPageControl) SetOnGetImageIndex(fn TTabGetImageEvent) {
    PageControl_SetOnGetImageIndex(p.instance, fn)
}

// SetOnGetSiteInfo
func (p *TPageControl) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    PageControl_SetOnGetSiteInfo(p.instance, fn)
}

// SetOnMouseActivate
func (p *TPageControl) SetOnMouseActivate(fn TMouseActivateEvent) {
    PageControl_SetOnMouseActivate(p.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (p *TPageControl) SetOnMouseDown(fn TMouseEvent) {
    PageControl_SetOnMouseDown(p.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (p *TPageControl) SetOnMouseEnter(fn TNotifyEvent) {
    PageControl_SetOnMouseEnter(p.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (p *TPageControl) SetOnMouseLeave(fn TNotifyEvent) {
    PageControl_SetOnMouseLeave(p.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (p *TPageControl) SetOnMouseMove(fn TMouseMoveEvent) {
    PageControl_SetOnMouseMove(p.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (p *TPageControl) SetOnMouseUp(fn TMouseEvent) {
    PageControl_SetOnMouseUp(p.instance, fn)
}

// SetOnResize
// CN: 设置大小被改变事件。
// EN: .
func (p *TPageControl) SetOnResize(fn TNotifyEvent) {
    PageControl_SetOnResize(p.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (p *TPageControl) SetOnStartDock(fn TStartDockEvent) {
    PageControl_SetOnStartDock(p.instance, fn)
}

// SetOnUnDock
func (p *TPageControl) SetOnUnDock(fn TUnDockEvent) {
    PageControl_SetOnUnDock(p.instance, fn)
}

// Canvas
// CN: 获取画布。
// EN: .
func (p *TPageControl) Canvas() *TCanvas {
    return CanvasFromInst(PageControl_GetCanvas(p.instance))
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (p *TPageControl) DockClientCount() int32 {
    return PageControl_GetDockClientCount(p.instance)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (p *TPageControl) AlignDisabled() bool {
    return PageControl_GetAlignDisabled(p.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (p *TPageControl) MouseInClient() bool {
    return PageControl_GetMouseInClient(p.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (p *TPageControl) VisibleDockClientCount() int32 {
    return PageControl_GetVisibleDockClientCount(p.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (p *TPageControl) Brush() *TBrush {
    return BrushFromInst(PageControl_GetBrush(p.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (p *TPageControl) ControlCount() int32 {
    return PageControl_GetControlCount(p.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPageControl) Handle() HWND {
    return PageControl_GetHandle(p.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (p *TPageControl) ParentWindow() HWND {
    return PageControl_GetParentWindow(p.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (p *TPageControl) SetParentWindow(value HWND) {
    PageControl_SetParentWindow(p.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (p *TPageControl) UseDockManager() bool {
    return PageControl_GetUseDockManager(p.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (p *TPageControl) SetUseDockManager(value bool) {
    PageControl_SetUseDockManager(p.instance, value)
}

// Action
func (p *TPageControl) Action() *TAction {
    return ActionFromInst(PageControl_GetAction(p.instance))
}

// SetAction
func (p *TPageControl) SetAction(value IComponent) {
    PageControl_SetAction(p.instance, CheckPtr(value))
}

// BoundsRect
func (p *TPageControl) BoundsRect() TRect {
    return PageControl_GetBoundsRect(p.instance)
}

// SetBoundsRect
func (p *TPageControl) SetBoundsRect(value TRect) {
    PageControl_SetBoundsRect(p.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (p *TPageControl) ClientHeight() int32 {
    return PageControl_GetClientHeight(p.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (p *TPageControl) SetClientHeight(value int32) {
    PageControl_SetClientHeight(p.instance, value)
}

// ClientOrigin
func (p *TPageControl) ClientOrigin() TPoint {
    return PageControl_GetClientOrigin(p.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (p *TPageControl) ClientRect() TRect {
    return PageControl_GetClientRect(p.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (p *TPageControl) ClientWidth() int32 {
    return PageControl_GetClientWidth(p.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (p *TPageControl) SetClientWidth(value int32) {
    PageControl_SetClientWidth(p.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (p *TPageControl) ControlState() TControlState {
    return PageControl_GetControlState(p.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (p *TPageControl) SetControlState(value TControlState) {
    PageControl_SetControlState(p.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (p *TPageControl) ControlStyle() TControlStyle {
    return PageControl_GetControlStyle(p.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (p *TPageControl) SetControlStyle(value TControlStyle) {
    PageControl_SetControlStyle(p.instance, value)
}

// ExplicitLeft
func (p *TPageControl) ExplicitLeft() int32 {
    return PageControl_GetExplicitLeft(p.instance)
}

// ExplicitTop
func (p *TPageControl) ExplicitTop() int32 {
    return PageControl_GetExplicitTop(p.instance)
}

// ExplicitWidth
func (p *TPageControl) ExplicitWidth() int32 {
    return PageControl_GetExplicitWidth(p.instance)
}

// ExplicitHeight
func (p *TPageControl) ExplicitHeight() int32 {
    return PageControl_GetExplicitHeight(p.instance)
}

// Floating
func (p *TPageControl) Floating() bool {
    return PageControl_GetFloating(p.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (p *TPageControl) Parent() *TWinControl {
    return WinControlFromInst(PageControl_GetParent(p.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (p *TPageControl) SetParent(value IWinControl) {
    PageControl_SetParent(p.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (p *TPageControl) AlignWithMargins() bool {
    return PageControl_GetAlignWithMargins(p.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (p *TPageControl) SetAlignWithMargins(value bool) {
    PageControl_SetAlignWithMargins(p.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (p *TPageControl) Left() int32 {
    return PageControl_GetLeft(p.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (p *TPageControl) SetLeft(value int32) {
    PageControl_SetLeft(p.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (p *TPageControl) Top() int32 {
    return PageControl_GetTop(p.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (p *TPageControl) SetTop(value int32) {
    PageControl_SetTop(p.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (p *TPageControl) Width() int32 {
    return PageControl_GetWidth(p.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (p *TPageControl) SetWidth(value int32) {
    PageControl_SetWidth(p.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (p *TPageControl) Height() int32 {
    return PageControl_GetHeight(p.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (p *TPageControl) SetHeight(value int32) {
    PageControl_SetHeight(p.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (p *TPageControl) Cursor() TCursor {
    return PageControl_GetCursor(p.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (p *TPageControl) SetCursor(value TCursor) {
    PageControl_SetCursor(p.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (p *TPageControl) Hint() string {
    return PageControl_GetHint(p.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (p *TPageControl) SetHint(value string) {
    PageControl_SetHint(p.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (p *TPageControl) Margins() *TMargins {
    return MarginsFromInst(PageControl_GetMargins(p.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (p *TPageControl) SetMargins(value *TMargins) {
    PageControl_SetMargins(p.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (p *TPageControl) CustomHint() *TCustomHint {
    return CustomHintFromInst(PageControl_GetCustomHint(p.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (p *TPageControl) SetCustomHint(value IComponent) {
    PageControl_SetCustomHint(p.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPageControl) ComponentCount() int32 {
    return PageControl_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPageControl) ComponentIndex() int32 {
    return PageControl_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPageControl) SetComponentIndex(value int32) {
    PageControl_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPageControl) Owner() *TComponent {
    return ComponentFromInst(PageControl_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPageControl) Name() string {
    return PageControl_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPageControl) SetName(value string) {
    PageControl_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPageControl) Tag() int {
    return PageControl_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPageControl) SetTag(value int) {
    PageControl_SetTag(p.instance, value)
}

// Pages
func (p *TPageControl) Pages(Index int32) *TTabSheet {
    return TabSheetFromInst(PageControl_GetPages(p.instance, Index))
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (p *TPageControl) DockClients(Index int32) *TControl {
    return ControlFromInst(PageControl_GetDockClients(p.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (p *TPageControl) Controls(Index int32) *TControl {
    return ControlFromInst(PageControl_GetControls(p.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPageControl) Components(AIndex int32) *TComponent {
    return ComponentFromInst(PageControl_GetComponents(p.instance, AIndex))
}

