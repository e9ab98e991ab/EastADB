
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

type TPanel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewPanel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewPanel(owner IComponent) *TPanel {
    p := new(TPanel)
    p.instance = Panel_Create(CheckPtr(owner))
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PanelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func PanelFromInst(inst uintptr) *TPanel {
    p := new(TPanel)
    p.instance = inst
    p.ptr = unsafe.Pointer(inst)
    return p
}

// PanelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func PanelFromObj(obj IObject) *TPanel {
    p := new(TPanel)
    p.instance = CheckPtr(obj)
    p.ptr = unsafe.Pointer(p.instance)
    return p
}

// PanelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func PanelFromUnsafePointer(ptr unsafe.Pointer) *TPanel {
    p := new(TPanel)
    p.instance = uintptr(ptr)
    p.ptr = ptr
    return p
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (p *TPanel) Free() {
    if p.instance != 0 {
        Panel_Free(p.instance)
        p.instance = 0
        p.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (p *TPanel) Instance() uintptr {
    return p.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (p *TPanel) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (p *TPanel) IsValid() bool {
    return p.instance != 0
}

// TPanelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TPanelClass() TClass {
    return Panel_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (p *TPanel) CanFocus() bool {
    return Panel_CanFocus(p.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (p *TPanel) ContainsControl(Control IControl) bool {
    return Panel_ContainsControl(p.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (p *TPanel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(Panel_ControlAtPos(p.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (p *TPanel) DisableAlign() {
    Panel_DisableAlign(p.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (p *TPanel) EnableAlign() {
    Panel_EnableAlign(p.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (p *TPanel) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(Panel_FindChildControl(p.instance, ControlName))
}

// FlipChildren
func (p *TPanel) FlipChildren(AllLevels bool) {
    Panel_FlipChildren(p.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (p *TPanel) Focused() bool {
    return Panel_Focused(p.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (p *TPanel) HandleAllocated() bool {
    return Panel_HandleAllocated(p.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (p *TPanel) InsertControl(AControl IControl) {
    Panel_InsertControl(p.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (p *TPanel) Invalidate() {
    Panel_Invalidate(p.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (p *TPanel) PaintTo(DC HDC, X int32, Y int32) {
    Panel_PaintTo(p.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (p *TPanel) RemoveControl(AControl IControl) {
    Panel_RemoveControl(p.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (p *TPanel) Realign() {
    Panel_Realign(p.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (p *TPanel) Repaint() {
    Panel_Repaint(p.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (p *TPanel) ScaleBy(M int32, D int32) {
    Panel_ScaleBy(p.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (p *TPanel) ScrollBy(DeltaX int32, DeltaY int32) {
    Panel_ScrollBy(p.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (p *TPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Panel_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (p *TPanel) SetFocus() {
    Panel_SetFocus(p.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (p *TPanel) Update() {
    Panel_Update(p.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (p *TPanel) UpdateControlState() {
    Panel_UpdateControlState(p.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (p *TPanel) BringToFront() {
    Panel_BringToFront(p.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (p *TPanel) ClientToScreen(Point TPoint) TPoint {
    return Panel_ClientToScreen(p.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (p *TPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Panel_ClientToParent(p.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (p *TPanel) Dragging() bool {
    return Panel_Dragging(p.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (p *TPanel) HasParent() bool {
    return Panel_HasParent(p.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (p *TPanel) Hide() {
    Panel_Hide(p.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (p *TPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Panel_Perform(p.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (p *TPanel) Refresh() {
    Panel_Refresh(p.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (p *TPanel) ScreenToClient(Point TPoint) TPoint {
    return Panel_ScreenToClient(p.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (p *TPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Panel_ParentToClient(p.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (p *TPanel) SendToBack() {
    Panel_SendToBack(p.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (p *TPanel) Show() {
    Panel_Show(p.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (p *TPanel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Panel_GetTextBuf(p.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (p *TPanel) GetTextLen() int32 {
    return Panel_GetTextLen(p.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (p *TPanel) SetTextBuf(Buffer string) {
    Panel_SetTextBuf(p.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (p *TPanel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Panel_FindComponent(p.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (p *TPanel) GetNamePath() string {
    return Panel_GetNamePath(p.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (p *TPanel) Assign(Source IObject) {
    Panel_Assign(p.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (p *TPanel) DisposeOf() {
    Panel_DisposeOf(p.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (p *TPanel) ClassType() TClass {
    return Panel_ClassType(p.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (p *TPanel) ClassName() string {
    return Panel_ClassName(p.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (p *TPanel) InstanceSize() int32 {
    return Panel_InstanceSize(p.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (p *TPanel) InheritsFrom(AClass TClass) bool {
    return Panel_InheritsFrom(p.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (p *TPanel) Equals(Obj IObject) bool {
    return Panel_Equals(p.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (p *TPanel) GetHashCode() int32 {
    return Panel_GetHashCode(p.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (p *TPanel) ToString() string {
    return Panel_ToString(p.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (p *TPanel) Align() TAlign {
    return Panel_GetAlign(p.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (p *TPanel) SetAlign(value TAlign) {
    Panel_SetAlign(p.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (p *TPanel) Alignment() TAlignment {
    return Panel_GetAlignment(p.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (p *TPanel) SetAlignment(value TAlignment) {
    Panel_SetAlignment(p.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (p *TPanel) Anchors() TAnchors {
    return Panel_GetAnchors(p.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (p *TPanel) SetAnchors(value TAnchors) {
    Panel_SetAnchors(p.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (p *TPanel) AutoSize() bool {
    return Panel_GetAutoSize(p.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (p *TPanel) SetAutoSize(value bool) {
    Panel_SetAutoSize(p.instance, value)
}

// BevelEdges
func (p *TPanel) BevelEdges() TBevelEdges {
    return Panel_GetBevelEdges(p.instance)
}

// SetBevelEdges
func (p *TPanel) SetBevelEdges(value TBevelEdges) {
    Panel_SetBevelEdges(p.instance, value)
}

// BevelInner
func (p *TPanel) BevelInner() TBevelCut {
    return Panel_GetBevelInner(p.instance)
}

// SetBevelInner
func (p *TPanel) SetBevelInner(value TBevelCut) {
    Panel_SetBevelInner(p.instance, value)
}

// BevelKind
func (p *TPanel) BevelKind() TBevelKind {
    return Panel_GetBevelKind(p.instance)
}

// SetBevelKind
func (p *TPanel) SetBevelKind(value TBevelKind) {
    Panel_SetBevelKind(p.instance, value)
}

// BevelOuter
func (p *TPanel) BevelOuter() TBevelCut {
    return Panel_GetBevelOuter(p.instance)
}

// SetBevelOuter
func (p *TPanel) SetBevelOuter(value TBevelCut) {
    Panel_SetBevelOuter(p.instance, value)
}

// BiDiMode
func (p *TPanel) BiDiMode() TBiDiMode {
    return Panel_GetBiDiMode(p.instance)
}

// SetBiDiMode
func (p *TPanel) SetBiDiMode(value TBiDiMode) {
    Panel_SetBiDiMode(p.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (p *TPanel) BorderWidth() int32 {
    return Panel_GetBorderWidth(p.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (p *TPanel) SetBorderWidth(value int32) {
    Panel_SetBorderWidth(p.instance, value)
}

// BorderStyle
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (p *TPanel) BorderStyle() TBorderStyle {
    return Panel_GetBorderStyle(p.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (p *TPanel) SetBorderStyle(value TBorderStyle) {
    Panel_SetBorderStyle(p.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (p *TPanel) Caption() string {
    return Panel_GetCaption(p.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (p *TPanel) SetCaption(value string) {
    Panel_SetCaption(p.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (p *TPanel) Color() TColor {
    return Panel_GetColor(p.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (p *TPanel) SetColor(value TColor) {
    Panel_SetColor(p.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (p *TPanel) UseDockManager() bool {
    return Panel_GetUseDockManager(p.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (p *TPanel) SetUseDockManager(value bool) {
    Panel_SetUseDockManager(p.instance, value)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (p *TPanel) DockSite() bool {
    return Panel_GetDockSite(p.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (p *TPanel) SetDockSite(value bool) {
    Panel_SetDockSite(p.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (p *TPanel) DoubleBuffered() bool {
    return Panel_GetDoubleBuffered(p.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (p *TPanel) SetDoubleBuffered(value bool) {
    Panel_SetDoubleBuffered(p.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (p *TPanel) DragCursor() TCursor {
    return Panel_GetDragCursor(p.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (p *TPanel) SetDragCursor(value TCursor) {
    Panel_SetDragCursor(p.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (p *TPanel) DragKind() TDragKind {
    return Panel_GetDragKind(p.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (p *TPanel) SetDragKind(value TDragKind) {
    Panel_SetDragKind(p.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (p *TPanel) DragMode() TDragMode {
    return Panel_GetDragMode(p.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (p *TPanel) SetDragMode(value TDragMode) {
    Panel_SetDragMode(p.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (p *TPanel) Enabled() bool {
    return Panel_GetEnabled(p.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (p *TPanel) SetEnabled(value bool) {
    Panel_SetEnabled(p.instance, value)
}

// FullRepaint
func (p *TPanel) FullRepaint() bool {
    return Panel_GetFullRepaint(p.instance)
}

// SetFullRepaint
func (p *TPanel) SetFullRepaint(value bool) {
    Panel_SetFullRepaint(p.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (p *TPanel) Font() *TFont {
    return FontFromInst(Panel_GetFont(p.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (p *TPanel) SetFont(value *TFont) {
    Panel_SetFont(p.instance, CheckPtr(value))
}

// Locked
func (p *TPanel) Locked() bool {
    return Panel_GetLocked(p.instance)
}

// SetLocked
func (p *TPanel) SetLocked(value bool) {
    Panel_SetLocked(p.instance, value)
}

// ParentBackground
func (p *TPanel) ParentBackground() bool {
    return Panel_GetParentBackground(p.instance)
}

// SetParentBackground
func (p *TPanel) SetParentBackground(value bool) {
    Panel_SetParentBackground(p.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (p *TPanel) ParentColor() bool {
    return Panel_GetParentColor(p.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (p *TPanel) SetParentColor(value bool) {
    Panel_SetParentColor(p.instance, value)
}

// ParentCtl3D
func (p *TPanel) ParentCtl3D() bool {
    return Panel_GetParentCtl3D(p.instance)
}

// SetParentCtl3D
func (p *TPanel) SetParentCtl3D(value bool) {
    Panel_SetParentCtl3D(p.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (p *TPanel) ParentDoubleBuffered() bool {
    return Panel_GetParentDoubleBuffered(p.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (p *TPanel) SetParentDoubleBuffered(value bool) {
    Panel_SetParentDoubleBuffered(p.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (p *TPanel) ParentFont() bool {
    return Panel_GetParentFont(p.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (p *TPanel) SetParentFont(value bool) {
    Panel_SetParentFont(p.instance, value)
}

// ParentShowHint
func (p *TPanel) ParentShowHint() bool {
    return Panel_GetParentShowHint(p.instance)
}

// SetParentShowHint
func (p *TPanel) SetParentShowHint(value bool) {
    Panel_SetParentShowHint(p.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (p *TPanel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Panel_GetPopupMenu(p.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (p *TPanel) SetPopupMenu(value IComponent) {
    Panel_SetPopupMenu(p.instance, CheckPtr(value))
}

// ShowCaption
func (p *TPanel) ShowCaption() bool {
    return Panel_GetShowCaption(p.instance)
}

// SetShowCaption
func (p *TPanel) SetShowCaption(value bool) {
    Panel_SetShowCaption(p.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (p *TPanel) ShowHint() bool {
    return Panel_GetShowHint(p.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (p *TPanel) SetShowHint(value bool) {
    Panel_SetShowHint(p.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (p *TPanel) TabOrder() TTabOrder {
    return Panel_GetTabOrder(p.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (p *TPanel) SetTabOrder(value TTabOrder) {
    Panel_SetTabOrder(p.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (p *TPanel) TabStop() bool {
    return Panel_GetTabStop(p.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (p *TPanel) SetTabStop(value bool) {
    Panel_SetTabStop(p.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (p *TPanel) Visible() bool {
    return Panel_GetVisible(p.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (p *TPanel) SetVisible(value bool) {
    Panel_SetVisible(p.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (p *TPanel) StyleElements() TStyleElements {
    return Panel_GetStyleElements(p.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (p *TPanel) SetStyleElements(value TStyleElements) {
    Panel_SetStyleElements(p.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (p *TPanel) SetOnClick(fn TNotifyEvent) {
    Panel_SetOnClick(p.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (p *TPanel) SetOnContextPopup(fn TContextPopupEvent) {
    Panel_SetOnContextPopup(p.instance, fn)
}

// SetOnDockDrop
func (p *TPanel) SetOnDockDrop(fn TDockDropEvent) {
    Panel_SetOnDockDrop(p.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (p *TPanel) SetOnDblClick(fn TNotifyEvent) {
    Panel_SetOnDblClick(p.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (p *TPanel) SetOnDragDrop(fn TDragDropEvent) {
    Panel_SetOnDragDrop(p.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (p *TPanel) SetOnDragOver(fn TDragOverEvent) {
    Panel_SetOnDragOver(p.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (p *TPanel) SetOnEndDock(fn TEndDragEvent) {
    Panel_SetOnEndDock(p.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (p *TPanel) SetOnEndDrag(fn TEndDragEvent) {
    Panel_SetOnEndDrag(p.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (p *TPanel) SetOnEnter(fn TNotifyEvent) {
    Panel_SetOnEnter(p.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (p *TPanel) SetOnExit(fn TNotifyEvent) {
    Panel_SetOnExit(p.instance, fn)
}

// SetOnGesture
func (p *TPanel) SetOnGesture(fn TGestureEvent) {
    Panel_SetOnGesture(p.instance, fn)
}

// SetOnGetSiteInfo
func (p *TPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    Panel_SetOnGetSiteInfo(p.instance, fn)
}

// SetOnMouseActivate
func (p *TPanel) SetOnMouseActivate(fn TMouseActivateEvent) {
    Panel_SetOnMouseActivate(p.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (p *TPanel) SetOnMouseDown(fn TMouseEvent) {
    Panel_SetOnMouseDown(p.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (p *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
    Panel_SetOnMouseEnter(p.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (p *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
    Panel_SetOnMouseLeave(p.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (p *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    Panel_SetOnMouseMove(p.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (p *TPanel) SetOnMouseUp(fn TMouseEvent) {
    Panel_SetOnMouseUp(p.instance, fn)
}

// SetOnResize
// CN: 设置大小被改变事件。
// EN: .
func (p *TPanel) SetOnResize(fn TNotifyEvent) {
    Panel_SetOnResize(p.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (p *TPanel) SetOnStartDock(fn TStartDockEvent) {
    Panel_SetOnStartDock(p.instance, fn)
}

// SetOnUnDock
func (p *TPanel) SetOnUnDock(fn TUnDockEvent) {
    Panel_SetOnUnDock(p.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (p *TPanel) DockClientCount() int32 {
    return Panel_GetDockClientCount(p.instance)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (p *TPanel) AlignDisabled() bool {
    return Panel_GetAlignDisabled(p.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (p *TPanel) MouseInClient() bool {
    return Panel_GetMouseInClient(p.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (p *TPanel) VisibleDockClientCount() int32 {
    return Panel_GetVisibleDockClientCount(p.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (p *TPanel) Brush() *TBrush {
    return BrushFromInst(Panel_GetBrush(p.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (p *TPanel) ControlCount() int32 {
    return Panel_GetControlCount(p.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (p *TPanel) Handle() HWND {
    return Panel_GetHandle(p.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (p *TPanel) ParentWindow() HWND {
    return Panel_GetParentWindow(p.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (p *TPanel) SetParentWindow(value HWND) {
    Panel_SetParentWindow(p.instance, value)
}

// Action
func (p *TPanel) Action() *TAction {
    return ActionFromInst(Panel_GetAction(p.instance))
}

// SetAction
func (p *TPanel) SetAction(value IComponent) {
    Panel_SetAction(p.instance, CheckPtr(value))
}

// BoundsRect
func (p *TPanel) BoundsRect() TRect {
    return Panel_GetBoundsRect(p.instance)
}

// SetBoundsRect
func (p *TPanel) SetBoundsRect(value TRect) {
    Panel_SetBoundsRect(p.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (p *TPanel) ClientHeight() int32 {
    return Panel_GetClientHeight(p.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (p *TPanel) SetClientHeight(value int32) {
    Panel_SetClientHeight(p.instance, value)
}

// ClientOrigin
func (p *TPanel) ClientOrigin() TPoint {
    return Panel_GetClientOrigin(p.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (p *TPanel) ClientRect() TRect {
    return Panel_GetClientRect(p.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (p *TPanel) ClientWidth() int32 {
    return Panel_GetClientWidth(p.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (p *TPanel) SetClientWidth(value int32) {
    Panel_SetClientWidth(p.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (p *TPanel) ControlState() TControlState {
    return Panel_GetControlState(p.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (p *TPanel) SetControlState(value TControlState) {
    Panel_SetControlState(p.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (p *TPanel) ControlStyle() TControlStyle {
    return Panel_GetControlStyle(p.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (p *TPanel) SetControlStyle(value TControlStyle) {
    Panel_SetControlStyle(p.instance, value)
}

// ExplicitLeft
func (p *TPanel) ExplicitLeft() int32 {
    return Panel_GetExplicitLeft(p.instance)
}

// ExplicitTop
func (p *TPanel) ExplicitTop() int32 {
    return Panel_GetExplicitTop(p.instance)
}

// ExplicitWidth
func (p *TPanel) ExplicitWidth() int32 {
    return Panel_GetExplicitWidth(p.instance)
}

// ExplicitHeight
func (p *TPanel) ExplicitHeight() int32 {
    return Panel_GetExplicitHeight(p.instance)
}

// Floating
func (p *TPanel) Floating() bool {
    return Panel_GetFloating(p.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (p *TPanel) Parent() *TWinControl {
    return WinControlFromInst(Panel_GetParent(p.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (p *TPanel) SetParent(value IWinControl) {
    Panel_SetParent(p.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (p *TPanel) AlignWithMargins() bool {
    return Panel_GetAlignWithMargins(p.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (p *TPanel) SetAlignWithMargins(value bool) {
    Panel_SetAlignWithMargins(p.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (p *TPanel) Left() int32 {
    return Panel_GetLeft(p.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (p *TPanel) SetLeft(value int32) {
    Panel_SetLeft(p.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (p *TPanel) Top() int32 {
    return Panel_GetTop(p.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (p *TPanel) SetTop(value int32) {
    Panel_SetTop(p.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (p *TPanel) Width() int32 {
    return Panel_GetWidth(p.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (p *TPanel) SetWidth(value int32) {
    Panel_SetWidth(p.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (p *TPanel) Height() int32 {
    return Panel_GetHeight(p.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (p *TPanel) SetHeight(value int32) {
    Panel_SetHeight(p.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (p *TPanel) Cursor() TCursor {
    return Panel_GetCursor(p.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (p *TPanel) SetCursor(value TCursor) {
    Panel_SetCursor(p.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (p *TPanel) Hint() string {
    return Panel_GetHint(p.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (p *TPanel) SetHint(value string) {
    Panel_SetHint(p.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (p *TPanel) Margins() *TMargins {
    return MarginsFromInst(Panel_GetMargins(p.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (p *TPanel) SetMargins(value *TMargins) {
    Panel_SetMargins(p.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (p *TPanel) CustomHint() *TCustomHint {
    return CustomHintFromInst(Panel_GetCustomHint(p.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (p *TPanel) SetCustomHint(value IComponent) {
    Panel_SetCustomHint(p.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (p *TPanel) ComponentCount() int32 {
    return Panel_GetComponentCount(p.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (p *TPanel) ComponentIndex() int32 {
    return Panel_GetComponentIndex(p.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (p *TPanel) SetComponentIndex(value int32) {
    Panel_SetComponentIndex(p.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (p *TPanel) Owner() *TComponent {
    return ComponentFromInst(Panel_GetOwner(p.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (p *TPanel) Name() string {
    return Panel_GetName(p.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (p *TPanel) SetName(value string) {
    Panel_SetName(p.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (p *TPanel) Tag() int {
    return Panel_GetTag(p.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (p *TPanel) SetTag(value int) {
    Panel_SetTag(p.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (p *TPanel) DockClients(Index int32) *TControl {
    return ControlFromInst(Panel_GetDockClients(p.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (p *TPanel) Controls(Index int32) *TControl {
    return ControlFromInst(Panel_GetControls(p.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (p *TPanel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Panel_GetComponents(p.instance, AIndex))
}

