
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

type TCoolBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCoolBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCoolBar(owner IComponent) *TCoolBar {
    c := new(TCoolBar)
    c.instance = CoolBar_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CoolBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CoolBarFromInst(inst uintptr) *TCoolBar {
    c := new(TCoolBar)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CoolBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CoolBarFromObj(obj IObject) *TCoolBar {
    c := new(TCoolBar)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CoolBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CoolBarFromUnsafePointer(ptr unsafe.Pointer) *TCoolBar {
    c := new(TCoolBar)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCoolBar) Free() {
    if c.instance != 0 {
        CoolBar_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCoolBar) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCoolBar) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCoolBar) IsValid() bool {
    return c.instance != 0
}

// TCoolBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCoolBarClass() TClass {
    return CoolBar_StaticClassType()
}

// FlipChildren
func (c *TCoolBar) FlipChildren(AllLevels bool) {
    CoolBar_FlipChildren(c.instance, AllLevels)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (c *TCoolBar) CanFocus() bool {
    return CoolBar_CanFocus(c.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TCoolBar) ContainsControl(Control IControl) bool {
    return CoolBar_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TCoolBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(CoolBar_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TCoolBar) DisableAlign() {
    CoolBar_DisableAlign(c.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TCoolBar) EnableAlign() {
    CoolBar_EnableAlign(c.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (c *TCoolBar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(CoolBar_FindChildControl(c.instance, ControlName))
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TCoolBar) Focused() bool {
    return CoolBar_Focused(c.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TCoolBar) HandleAllocated() bool {
    return CoolBar_HandleAllocated(c.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (c *TCoolBar) InsertControl(AControl IControl) {
    CoolBar_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (c *TCoolBar) Invalidate() {
    CoolBar_Invalidate(c.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (c *TCoolBar) PaintTo(DC HDC, X int32, Y int32) {
    CoolBar_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (c *TCoolBar) RemoveControl(AControl IControl) {
    CoolBar_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (c *TCoolBar) Realign() {
    CoolBar_Realign(c.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (c *TCoolBar) Repaint() {
    CoolBar_Repaint(c.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (c *TCoolBar) ScaleBy(M int32, D int32) {
    CoolBar_ScaleBy(c.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TCoolBar) ScrollBy(DeltaX int32, DeltaY int32) {
    CoolBar_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TCoolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CoolBar_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TCoolBar) SetFocus() {
    CoolBar_SetFocus(c.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (c *TCoolBar) Update() {
    CoolBar_Update(c.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (c *TCoolBar) UpdateControlState() {
    CoolBar_UpdateControlState(c.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TCoolBar) BringToFront() {
    CoolBar_BringToFront(c.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TCoolBar) ClientToScreen(Point TPoint) TPoint {
    return CoolBar_ClientToScreen(c.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TCoolBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CoolBar_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TCoolBar) Dragging() bool {
    return CoolBar_Dragging(c.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TCoolBar) HasParent() bool {
    return CoolBar_HasParent(c.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (c *TCoolBar) Hide() {
    CoolBar_Hide(c.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (c *TCoolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CoolBar_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (c *TCoolBar) Refresh() {
    CoolBar_Refresh(c.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TCoolBar) ScreenToClient(Point TPoint) TPoint {
    return CoolBar_ScreenToClient(c.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TCoolBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CoolBar_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TCoolBar) SendToBack() {
    CoolBar_SendToBack(c.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (c *TCoolBar) Show() {
    CoolBar_Show(c.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TCoolBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CoolBar_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TCoolBar) GetTextLen() int32 {
    return CoolBar_GetTextLen(c.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TCoolBar) SetTextBuf(Buffer string) {
    CoolBar_SetTextBuf(c.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TCoolBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CoolBar_FindComponent(c.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCoolBar) GetNamePath() string {
    return CoolBar_GetNamePath(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCoolBar) Assign(Source IObject) {
    CoolBar_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCoolBar) DisposeOf() {
    CoolBar_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCoolBar) ClassType() TClass {
    return CoolBar_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCoolBar) ClassName() string {
    return CoolBar_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCoolBar) InstanceSize() int32 {
    return CoolBar_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCoolBar) InheritsFrom(AClass TClass) bool {
    return CoolBar_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCoolBar) Equals(Obj IObject) bool {
    return CoolBar_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCoolBar) GetHashCode() int32 {
    return CoolBar_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCoolBar) ToString() string {
    return CoolBar_ToString(c.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TCoolBar) Align() TAlign {
    return CoolBar_GetAlign(c.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TCoolBar) SetAlign(value TAlign) {
    CoolBar_SetAlign(c.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (c *TCoolBar) Anchors() TAnchors {
    return CoolBar_GetAnchors(c.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (c *TCoolBar) SetAnchors(value TAnchors) {
    CoolBar_SetAnchors(c.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (c *TCoolBar) AutoSize() bool {
    return CoolBar_GetAutoSize(c.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (c *TCoolBar) SetAutoSize(value bool) {
    CoolBar_SetAutoSize(c.instance, value)
}

// BandBorderStyle
func (c *TCoolBar) BandBorderStyle() TBorderStyle {
    return CoolBar_GetBandBorderStyle(c.instance)
}

// SetBandBorderStyle
func (c *TCoolBar) SetBandBorderStyle(value TBorderStyle) {
    CoolBar_SetBandBorderStyle(c.instance, value)
}

// BandMaximize
func (c *TCoolBar) BandMaximize() TCoolBandMaximize {
    return CoolBar_GetBandMaximize(c.instance)
}

// SetBandMaximize
func (c *TCoolBar) SetBandMaximize(value TCoolBandMaximize) {
    CoolBar_SetBandMaximize(c.instance, value)
}

// Bands
func (c *TCoolBar) Bands() *TCoolBands {
    return CoolBandsFromInst(CoolBar_GetBands(c.instance))
}

// SetBands
func (c *TCoolBar) SetBands(value *TCoolBands) {
    CoolBar_SetBands(c.instance, CheckPtr(value))
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (c *TCoolBar) BorderWidth() int32 {
    return CoolBar_GetBorderWidth(c.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (c *TCoolBar) SetBorderWidth(value int32) {
    CoolBar_SetBorderWidth(c.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (c *TCoolBar) Color() TColor {
    return CoolBar_GetColor(c.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (c *TCoolBar) SetColor(value TColor) {
    CoolBar_SetColor(c.instance, value)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TCoolBar) DockSite() bool {
    return CoolBar_GetDockSite(c.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TCoolBar) SetDockSite(value bool) {
    CoolBar_SetDockSite(c.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TCoolBar) DoubleBuffered() bool {
    return CoolBar_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TCoolBar) SetDoubleBuffered(value bool) {
    CoolBar_SetDoubleBuffered(c.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (c *TCoolBar) DragCursor() TCursor {
    return CoolBar_GetDragCursor(c.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (c *TCoolBar) SetDragCursor(value TCursor) {
    CoolBar_SetDragCursor(c.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (c *TCoolBar) DragKind() TDragKind {
    return CoolBar_GetDragKind(c.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (c *TCoolBar) SetDragKind(value TDragKind) {
    CoolBar_SetDragKind(c.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (c *TCoolBar) DragMode() TDragMode {
    return CoolBar_GetDragMode(c.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (c *TCoolBar) SetDragMode(value TDragMode) {
    CoolBar_SetDragMode(c.instance, value)
}

// EdgeBorders
func (c *TCoolBar) EdgeBorders() TEdgeBorders {
    return CoolBar_GetEdgeBorders(c.instance)
}

// SetEdgeBorders
func (c *TCoolBar) SetEdgeBorders(value TEdgeBorders) {
    CoolBar_SetEdgeBorders(c.instance, value)
}

// EdgeInner
func (c *TCoolBar) EdgeInner() TEdgeStyle {
    return CoolBar_GetEdgeInner(c.instance)
}

// SetEdgeInner
func (c *TCoolBar) SetEdgeInner(value TEdgeStyle) {
    CoolBar_SetEdgeInner(c.instance, value)
}

// EdgeOuter
func (c *TCoolBar) EdgeOuter() TEdgeStyle {
    return CoolBar_GetEdgeOuter(c.instance)
}

// SetEdgeOuter
func (c *TCoolBar) SetEdgeOuter(value TEdgeStyle) {
    CoolBar_SetEdgeOuter(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCoolBar) Enabled() bool {
    return CoolBar_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCoolBar) SetEnabled(value bool) {
    CoolBar_SetEnabled(c.instance, value)
}

// FixedSize
func (c *TCoolBar) FixedSize() bool {
    return CoolBar_GetFixedSize(c.instance)
}

// SetFixedSize
func (c *TCoolBar) SetFixedSize(value bool) {
    CoolBar_SetFixedSize(c.instance, value)
}

// FixedOrder
func (c *TCoolBar) FixedOrder() bool {
    return CoolBar_GetFixedOrder(c.instance)
}

// SetFixedOrder
func (c *TCoolBar) SetFixedOrder(value bool) {
    CoolBar_SetFixedOrder(c.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (c *TCoolBar) Font() *TFont {
    return FontFromInst(CoolBar_GetFont(c.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (c *TCoolBar) SetFont(value *TFont) {
    CoolBar_SetFont(c.instance, CheckPtr(value))
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (c *TCoolBar) Images() *TImageList {
    return ImageListFromInst(CoolBar_GetImages(c.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (c *TCoolBar) SetImages(value IComponent) {
    CoolBar_SetImages(c.instance, CheckPtr(value))
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TCoolBar) ParentColor() bool {
    return CoolBar_GetParentColor(c.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TCoolBar) SetParentColor(value bool) {
    CoolBar_SetParentColor(c.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TCoolBar) ParentDoubleBuffered() bool {
    return CoolBar_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TCoolBar) SetParentDoubleBuffered(value bool) {
    CoolBar_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (c *TCoolBar) ParentFont() bool {
    return CoolBar_GetParentFont(c.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (c *TCoolBar) SetParentFont(value bool) {
    CoolBar_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TCoolBar) ParentShowHint() bool {
    return CoolBar_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TCoolBar) SetParentShowHint(value bool) {
    CoolBar_SetParentShowHint(c.instance, value)
}

// Bitmap
func (c *TCoolBar) Bitmap() *TBitmap {
    return BitmapFromInst(CoolBar_GetBitmap(c.instance))
}

// SetBitmap
func (c *TCoolBar) SetBitmap(value *TBitmap) {
    CoolBar_SetBitmap(c.instance, CheckPtr(value))
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TCoolBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CoolBar_GetPopupMenu(c.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TCoolBar) SetPopupMenu(value IComponent) {
    CoolBar_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TCoolBar) ShowHint() bool {
    return CoolBar_GetShowHint(c.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TCoolBar) SetShowHint(value bool) {
    CoolBar_SetShowHint(c.instance, value)
}

// ShowText
func (c *TCoolBar) ShowText() bool {
    return CoolBar_GetShowText(c.instance)
}

// SetShowText
func (c *TCoolBar) SetShowText(value bool) {
    CoolBar_SetShowText(c.instance, value)
}

// Vertical
func (c *TCoolBar) Vertical() bool {
    return CoolBar_GetVertical(c.instance)
}

// SetVertical
func (c *TCoolBar) SetVertical(value bool) {
    CoolBar_SetVertical(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCoolBar) Visible() bool {
    return CoolBar_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCoolBar) SetVisible(value bool) {
    CoolBar_SetVisible(c.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (c *TCoolBar) StyleElements() TStyleElements {
    return CoolBar_GetStyleElements(c.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (c *TCoolBar) SetStyleElements(value TStyleElements) {
    CoolBar_SetStyleElements(c.instance, value)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (c *TCoolBar) SetOnChange(fn TNotifyEvent) {
    CoolBar_SetOnChange(c.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCoolBar) SetOnClick(fn TNotifyEvent) {
    CoolBar_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TCoolBar) SetOnContextPopup(fn TContextPopupEvent) {
    CoolBar_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (c *TCoolBar) SetOnDblClick(fn TNotifyEvent) {
    CoolBar_SetOnDblClick(c.instance, fn)
}

// SetOnDockDrop
func (c *TCoolBar) SetOnDockDrop(fn TDockDropEvent) {
    CoolBar_SetOnDockDrop(c.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TCoolBar) SetOnDragDrop(fn TDragDropEvent) {
    CoolBar_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TCoolBar) SetOnDragOver(fn TDragOverEvent) {
    CoolBar_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (c *TCoolBar) SetOnEndDock(fn TEndDragEvent) {
    CoolBar_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TCoolBar) SetOnEndDrag(fn TEndDragEvent) {
    CoolBar_SetOnEndDrag(c.instance, fn)
}

// SetOnGesture
func (c *TCoolBar) SetOnGesture(fn TGestureEvent) {
    CoolBar_SetOnGesture(c.instance, fn)
}

// SetOnGetSiteInfo
func (c *TCoolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    CoolBar_SetOnGetSiteInfo(c.instance, fn)
}

// SetOnMouseActivate
func (c *TCoolBar) SetOnMouseActivate(fn TMouseActivateEvent) {
    CoolBar_SetOnMouseActivate(c.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (c *TCoolBar) SetOnMouseDown(fn TMouseEvent) {
    CoolBar_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (c *TCoolBar) SetOnMouseEnter(fn TNotifyEvent) {
    CoolBar_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (c *TCoolBar) SetOnMouseLeave(fn TNotifyEvent) {
    CoolBar_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (c *TCoolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    CoolBar_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (c *TCoolBar) SetOnMouseUp(fn TMouseEvent) {
    CoolBar_SetOnMouseUp(c.instance, fn)
}

// SetOnResize
// CN: 设置大小被改变事件。
// EN: .
func (c *TCoolBar) SetOnResize(fn TNotifyEvent) {
    CoolBar_SetOnResize(c.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (c *TCoolBar) SetOnStartDock(fn TStartDockEvent) {
    CoolBar_SetOnStartDock(c.instance, fn)
}

// SetOnUnDock
func (c *TCoolBar) SetOnUnDock(fn TUnDockEvent) {
    CoolBar_SetOnUnDock(c.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (c *TCoolBar) DockClientCount() int32 {
    return CoolBar_GetDockClientCount(c.instance)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (c *TCoolBar) AlignDisabled() bool {
    return CoolBar_GetAlignDisabled(c.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TCoolBar) MouseInClient() bool {
    return CoolBar_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TCoolBar) VisibleDockClientCount() int32 {
    return CoolBar_GetVisibleDockClientCount(c.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TCoolBar) Brush() *TBrush {
    return BrushFromInst(CoolBar_GetBrush(c.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TCoolBar) ControlCount() int32 {
    return CoolBar_GetControlCount(c.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TCoolBar) Handle() HWND {
    return CoolBar_GetHandle(c.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TCoolBar) ParentWindow() HWND {
    return CoolBar_GetParentWindow(c.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TCoolBar) SetParentWindow(value HWND) {
    CoolBar_SetParentWindow(c.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TCoolBar) TabOrder() TTabOrder {
    return CoolBar_GetTabOrder(c.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TCoolBar) SetTabOrder(value TTabOrder) {
    CoolBar_SetTabOrder(c.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TCoolBar) TabStop() bool {
    return CoolBar_GetTabStop(c.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TCoolBar) SetTabStop(value bool) {
    CoolBar_SetTabStop(c.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (c *TCoolBar) UseDockManager() bool {
    return CoolBar_GetUseDockManager(c.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (c *TCoolBar) SetUseDockManager(value bool) {
    CoolBar_SetUseDockManager(c.instance, value)
}

// Action
func (c *TCoolBar) Action() *TAction {
    return ActionFromInst(CoolBar_GetAction(c.instance))
}

// SetAction
func (c *TCoolBar) SetAction(value IComponent) {
    CoolBar_SetAction(c.instance, CheckPtr(value))
}

// BiDiMode
func (c *TCoolBar) BiDiMode() TBiDiMode {
    return CoolBar_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TCoolBar) SetBiDiMode(value TBiDiMode) {
    CoolBar_SetBiDiMode(c.instance, value)
}

// BoundsRect
func (c *TCoolBar) BoundsRect() TRect {
    return CoolBar_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TCoolBar) SetBoundsRect(value TRect) {
    CoolBar_SetBoundsRect(c.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (c *TCoolBar) ClientHeight() int32 {
    return CoolBar_GetClientHeight(c.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (c *TCoolBar) SetClientHeight(value int32) {
    CoolBar_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TCoolBar) ClientOrigin() TPoint {
    return CoolBar_GetClientOrigin(c.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TCoolBar) ClientRect() TRect {
    return CoolBar_GetClientRect(c.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TCoolBar) ClientWidth() int32 {
    return CoolBar_GetClientWidth(c.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TCoolBar) SetClientWidth(value int32) {
    CoolBar_SetClientWidth(c.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (c *TCoolBar) ControlState() TControlState {
    return CoolBar_GetControlState(c.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (c *TCoolBar) SetControlState(value TControlState) {
    CoolBar_SetControlState(c.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (c *TCoolBar) ControlStyle() TControlStyle {
    return CoolBar_GetControlStyle(c.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (c *TCoolBar) SetControlStyle(value TControlStyle) {
    CoolBar_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TCoolBar) ExplicitLeft() int32 {
    return CoolBar_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TCoolBar) ExplicitTop() int32 {
    return CoolBar_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TCoolBar) ExplicitWidth() int32 {
    return CoolBar_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TCoolBar) ExplicitHeight() int32 {
    return CoolBar_GetExplicitHeight(c.instance)
}

// Floating
func (c *TCoolBar) Floating() bool {
    return CoolBar_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCoolBar) Parent() *TWinControl {
    return WinControlFromInst(CoolBar_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCoolBar) SetParent(value IWinControl) {
    CoolBar_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (c *TCoolBar) AlignWithMargins() bool {
    return CoolBar_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (c *TCoolBar) SetAlignWithMargins(value bool) {
    CoolBar_SetAlignWithMargins(c.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (c *TCoolBar) Left() int32 {
    return CoolBar_GetLeft(c.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (c *TCoolBar) SetLeft(value int32) {
    CoolBar_SetLeft(c.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TCoolBar) Top() int32 {
    return CoolBar_GetTop(c.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TCoolBar) SetTop(value int32) {
    CoolBar_SetTop(c.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (c *TCoolBar) Width() int32 {
    return CoolBar_GetWidth(c.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (c *TCoolBar) SetWidth(value int32) {
    CoolBar_SetWidth(c.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (c *TCoolBar) Height() int32 {
    return CoolBar_GetHeight(c.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (c *TCoolBar) SetHeight(value int32) {
    CoolBar_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCoolBar) Cursor() TCursor {
    return CoolBar_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCoolBar) SetCursor(value TCursor) {
    CoolBar_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TCoolBar) Hint() string {
    return CoolBar_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TCoolBar) SetHint(value string) {
    CoolBar_SetHint(c.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (c *TCoolBar) Margins() *TMargins {
    return MarginsFromInst(CoolBar_GetMargins(c.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (c *TCoolBar) SetMargins(value *TMargins) {
    CoolBar_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (c *TCoolBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(CoolBar_GetCustomHint(c.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (c *TCoolBar) SetCustomHint(value IComponent) {
    CoolBar_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCoolBar) ComponentCount() int32 {
    return CoolBar_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCoolBar) ComponentIndex() int32 {
    return CoolBar_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCoolBar) SetComponentIndex(value int32) {
    CoolBar_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCoolBar) Owner() *TComponent {
    return ComponentFromInst(CoolBar_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCoolBar) Name() string {
    return CoolBar_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCoolBar) SetName(value string) {
    CoolBar_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCoolBar) Tag() int {
    return CoolBar_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCoolBar) SetTag(value int) {
    CoolBar_SetTag(c.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TCoolBar) DockClients(Index int32) *TControl {
    return ControlFromInst(CoolBar_GetDockClients(c.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (c *TCoolBar) Controls(Index int32) *TControl {
    return ControlFromInst(CoolBar_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCoolBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CoolBar_GetComponents(c.instance, AIndex))
}

