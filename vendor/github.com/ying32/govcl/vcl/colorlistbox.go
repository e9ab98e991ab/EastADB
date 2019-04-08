
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

type TColorListBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewColorListBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewColorListBox(owner IComponent) *TColorListBox {
    c := new(TColorListBox)
    c.instance = ColorListBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorListBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ColorListBoxFromInst(inst uintptr) *TColorListBox {
    c := new(TColorListBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// ColorListBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ColorListBoxFromObj(obj IObject) *TColorListBox {
    c := new(TColorListBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// ColorListBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ColorListBoxFromUnsafePointer(ptr unsafe.Pointer) *TColorListBox {
    c := new(TColorListBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TColorListBox) Free() {
    if c.instance != 0 {
        ColorListBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TColorListBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TColorListBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TColorListBox) IsValid() bool {
    return c.instance != 0
}

// TColorListBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TColorListBoxClass() TClass {
    return ColorListBox_StaticClassType()
}

// AddItem
func (c *TColorListBox) AddItem(Item string, AObject IObject) {
    ColorListBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

// Clear
// CN: 清除。
// EN: .
func (c *TColorListBox) Clear() {
    ColorListBox_Clear(c.instance)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (c *TColorListBox) ClearSelection() {
    ColorListBox_ClearSelection(c.instance)
}

// DeleteSelected
func (c *TColorListBox) DeleteSelected() {
    ColorListBox_DeleteSelected(c.instance)
}

// SelectAll
// CN: 全选。
// EN: .
func (c *TColorListBox) SelectAll() {
    ColorListBox_SelectAll(c.instance)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (c *TColorListBox) CanFocus() bool {
    return ColorListBox_CanFocus(c.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TColorListBox) ContainsControl(Control IControl) bool {
    return ColorListBox_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TColorListBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(ColorListBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TColorListBox) DisableAlign() {
    ColorListBox_DisableAlign(c.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TColorListBox) EnableAlign() {
    ColorListBox_EnableAlign(c.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (c *TColorListBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(ColorListBox_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TColorListBox) FlipChildren(AllLevels bool) {
    ColorListBox_FlipChildren(c.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TColorListBox) Focused() bool {
    return ColorListBox_Focused(c.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TColorListBox) HandleAllocated() bool {
    return ColorListBox_HandleAllocated(c.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (c *TColorListBox) InsertControl(AControl IControl) {
    ColorListBox_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (c *TColorListBox) Invalidate() {
    ColorListBox_Invalidate(c.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (c *TColorListBox) PaintTo(DC HDC, X int32, Y int32) {
    ColorListBox_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (c *TColorListBox) RemoveControl(AControl IControl) {
    ColorListBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (c *TColorListBox) Realign() {
    ColorListBox_Realign(c.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (c *TColorListBox) Repaint() {
    ColorListBox_Repaint(c.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (c *TColorListBox) ScaleBy(M int32, D int32) {
    ColorListBox_ScaleBy(c.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TColorListBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ColorListBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TColorListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ColorListBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TColorListBox) SetFocus() {
    ColorListBox_SetFocus(c.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (c *TColorListBox) Update() {
    ColorListBox_Update(c.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (c *TColorListBox) UpdateControlState() {
    ColorListBox_UpdateControlState(c.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TColorListBox) BringToFront() {
    ColorListBox_BringToFront(c.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TColorListBox) ClientToScreen(Point TPoint) TPoint {
    return ColorListBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TColorListBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ColorListBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TColorListBox) Dragging() bool {
    return ColorListBox_Dragging(c.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TColorListBox) HasParent() bool {
    return ColorListBox_HasParent(c.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (c *TColorListBox) Hide() {
    ColorListBox_Hide(c.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (c *TColorListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ColorListBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (c *TColorListBox) Refresh() {
    ColorListBox_Refresh(c.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TColorListBox) ScreenToClient(Point TPoint) TPoint {
    return ColorListBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TColorListBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ColorListBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TColorListBox) SendToBack() {
    ColorListBox_SendToBack(c.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (c *TColorListBox) Show() {
    ColorListBox_Show(c.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TColorListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ColorListBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TColorListBox) GetTextLen() int32 {
    return ColorListBox_GetTextLen(c.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TColorListBox) SetTextBuf(Buffer string) {
    ColorListBox_SetTextBuf(c.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TColorListBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ColorListBox_FindComponent(c.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TColorListBox) GetNamePath() string {
    return ColorListBox_GetNamePath(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TColorListBox) Assign(Source IObject) {
    ColorListBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TColorListBox) DisposeOf() {
    ColorListBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TColorListBox) ClassType() TClass {
    return ColorListBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TColorListBox) ClassName() string {
    return ColorListBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TColorListBox) InstanceSize() int32 {
    return ColorListBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TColorListBox) InheritsFrom(AClass TClass) bool {
    return ColorListBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TColorListBox) Equals(Obj IObject) bool {
    return ColorListBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TColorListBox) GetHashCode() int32 {
    return ColorListBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TColorListBox) ToString() string {
    return ColorListBox_ToString(c.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TColorListBox) Align() TAlign {
    return ColorListBox_GetAlign(c.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TColorListBox) SetAlign(value TAlign) {
    ColorListBox_SetAlign(c.instance, value)
}

// AutoComplete
func (c *TColorListBox) AutoComplete() bool {
    return ColorListBox_GetAutoComplete(c.instance)
}

// SetAutoComplete
func (c *TColorListBox) SetAutoComplete(value bool) {
    ColorListBox_SetAutoComplete(c.instance, value)
}

// DefaultColorColor
func (c *TColorListBox) DefaultColorColor() TColor {
    return ColorListBox_GetDefaultColorColor(c.instance)
}

// SetDefaultColorColor
func (c *TColorListBox) SetDefaultColorColor(value TColor) {
    ColorListBox_SetDefaultColorColor(c.instance, value)
}

// NoneColorColor
func (c *TColorListBox) NoneColorColor() TColor {
    return ColorListBox_GetNoneColorColor(c.instance)
}

// SetNoneColorColor
func (c *TColorListBox) SetNoneColorColor(value TColor) {
    ColorListBox_SetNoneColorColor(c.instance, value)
}

// Selected
func (c *TColorListBox) Selected() TColor {
    return ColorListBox_GetSelected(c.instance)
}

// SetSelected
func (c *TColorListBox) SetSelected(value TColor) {
    ColorListBox_SetSelected(c.instance, value)
}

// Style
func (c *TColorListBox) Style() TColorBoxStyle {
    return ColorListBox_GetStyle(c.instance)
}

// SetStyle
func (c *TColorListBox) SetStyle(value TColorBoxStyle) {
    ColorListBox_SetStyle(c.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (c *TColorListBox) Anchors() TAnchors {
    return ColorListBox_GetAnchors(c.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (c *TColorListBox) SetAnchors(value TAnchors) {
    ColorListBox_SetAnchors(c.instance, value)
}

// BevelEdges
func (c *TColorListBox) BevelEdges() TBevelEdges {
    return ColorListBox_GetBevelEdges(c.instance)
}

// SetBevelEdges
func (c *TColorListBox) SetBevelEdges(value TBevelEdges) {
    ColorListBox_SetBevelEdges(c.instance, value)
}

// BevelInner
func (c *TColorListBox) BevelInner() TBevelCut {
    return ColorListBox_GetBevelInner(c.instance)
}

// SetBevelInner
func (c *TColorListBox) SetBevelInner(value TBevelCut) {
    ColorListBox_SetBevelInner(c.instance, value)
}

// BevelKind
func (c *TColorListBox) BevelKind() TBevelKind {
    return ColorListBox_GetBevelKind(c.instance)
}

// SetBevelKind
func (c *TColorListBox) SetBevelKind(value TBevelKind) {
    ColorListBox_SetBevelKind(c.instance, value)
}

// BevelOuter
func (c *TColorListBox) BevelOuter() TBevelCut {
    return ColorListBox_GetBevelOuter(c.instance)
}

// SetBevelOuter
func (c *TColorListBox) SetBevelOuter(value TBevelCut) {
    ColorListBox_SetBevelOuter(c.instance, value)
}

// BiDiMode
func (c *TColorListBox) BiDiMode() TBiDiMode {
    return ColorListBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TColorListBox) SetBiDiMode(value TBiDiMode) {
    ColorListBox_SetBiDiMode(c.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (c *TColorListBox) Color() TColor {
    return ColorListBox_GetColor(c.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (c *TColorListBox) SetColor(value TColor) {
    ColorListBox_SetColor(c.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TColorListBox) DoubleBuffered() bool {
    return ColorListBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TColorListBox) SetDoubleBuffered(value bool) {
    ColorListBox_SetDoubleBuffered(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TColorListBox) Enabled() bool {
    return ColorListBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TColorListBox) SetEnabled(value bool) {
    ColorListBox_SetEnabled(c.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (c *TColorListBox) Font() *TFont {
    return FontFromInst(ColorListBox_GetFont(c.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (c *TColorListBox) SetFont(value *TFont) {
    ColorListBox_SetFont(c.instance, CheckPtr(value))
}

// ItemHeight
func (c *TColorListBox) ItemHeight() int32 {
    return ColorListBox_GetItemHeight(c.instance)
}

// SetItemHeight
func (c *TColorListBox) SetItemHeight(value int32) {
    ColorListBox_SetItemHeight(c.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TColorListBox) ParentColor() bool {
    return ColorListBox_GetParentColor(c.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TColorListBox) SetParentColor(value bool) {
    ColorListBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TColorListBox) ParentCtl3D() bool {
    return ColorListBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TColorListBox) SetParentCtl3D(value bool) {
    ColorListBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TColorListBox) ParentDoubleBuffered() bool {
    return ColorListBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TColorListBox) SetParentDoubleBuffered(value bool) {
    ColorListBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (c *TColorListBox) ParentFont() bool {
    return ColorListBox_GetParentFont(c.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (c *TColorListBox) SetParentFont(value bool) {
    ColorListBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TColorListBox) ParentShowHint() bool {
    return ColorListBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TColorListBox) SetParentShowHint(value bool) {
    ColorListBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TColorListBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ColorListBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TColorListBox) SetPopupMenu(value IComponent) {
    ColorListBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TColorListBox) ShowHint() bool {
    return ColorListBox_GetShowHint(c.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TColorListBox) SetShowHint(value bool) {
    ColorListBox_SetShowHint(c.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TColorListBox) TabOrder() TTabOrder {
    return ColorListBox_GetTabOrder(c.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TColorListBox) SetTabOrder(value TTabOrder) {
    ColorListBox_SetTabOrder(c.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TColorListBox) TabStop() bool {
    return ColorListBox_GetTabStop(c.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TColorListBox) SetTabStop(value bool) {
    ColorListBox_SetTabStop(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TColorListBox) Visible() bool {
    return ColorListBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TColorListBox) SetVisible(value bool) {
    ColorListBox_SetVisible(c.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (c *TColorListBox) StyleElements() TStyleElements {
    return ColorListBox_GetStyleElements(c.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (c *TColorListBox) SetStyleElements(value TStyleElements) {
    ColorListBox_SetStyleElements(c.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TColorListBox) SetOnClick(fn TNotifyEvent) {
    ColorListBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TColorListBox) SetOnContextPopup(fn TContextPopupEvent) {
    ColorListBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (c *TColorListBox) SetOnDblClick(fn TNotifyEvent) {
    ColorListBox_SetOnDblClick(c.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TColorListBox) SetOnDragDrop(fn TDragDropEvent) {
    ColorListBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TColorListBox) SetOnDragOver(fn TDragOverEvent) {
    ColorListBox_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (c *TColorListBox) SetOnEndDock(fn TEndDragEvent) {
    ColorListBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TColorListBox) SetOnEndDrag(fn TEndDragEvent) {
    ColorListBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (c *TColorListBox) SetOnEnter(fn TNotifyEvent) {
    ColorListBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (c *TColorListBox) SetOnExit(fn TNotifyEvent) {
    ColorListBox_SetOnExit(c.instance, fn)
}

// SetOnGesture
func (c *TColorListBox) SetOnGesture(fn TGestureEvent) {
    ColorListBox_SetOnGesture(c.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (c *TColorListBox) SetOnKeyDown(fn TKeyEvent) {
    ColorListBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TColorListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorListBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (c *TColorListBox) SetOnKeyUp(fn TKeyEvent) {
    ColorListBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseActivate
func (c *TColorListBox) SetOnMouseActivate(fn TMouseActivateEvent) {
    ColorListBox_SetOnMouseActivate(c.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (c *TColorListBox) SetOnMouseDown(fn TMouseEvent) {
    ColorListBox_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (c *TColorListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorListBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (c *TColorListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorListBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (c *TColorListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ColorListBox_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (c *TColorListBox) SetOnMouseUp(fn TMouseEvent) {
    ColorListBox_SetOnMouseUp(c.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (c *TColorListBox) SetOnStartDock(fn TStartDockEvent) {
    ColorListBox_SetOnStartDock(c.instance, fn)
}

// AutoCompleteDelay
func (c *TColorListBox) AutoCompleteDelay() uint32 {
    return ColorListBox_GetAutoCompleteDelay(c.instance)
}

// SetAutoCompleteDelay
func (c *TColorListBox) SetAutoCompleteDelay(value uint32) {
    ColorListBox_SetAutoCompleteDelay(c.instance, value)
}

// Canvas
// CN: 获取画布。
// EN: .
func (c *TColorListBox) Canvas() *TCanvas {
    return CanvasFromInst(ColorListBox_GetCanvas(c.instance))
}

// Count
func (c *TColorListBox) Count() int32 {
    return ColorListBox_GetCount(c.instance)
}

// SetCount
func (c *TColorListBox) SetCount(value int32) {
    ColorListBox_SetCount(c.instance, value)
}

// Items
func (c *TColorListBox) Items() *TStrings {
    return StringsFromInst(ColorListBox_GetItems(c.instance))
}

// SetItems
func (c *TColorListBox) SetItems(value IObject) {
    ColorListBox_SetItems(c.instance, CheckPtr(value))
}

// MultiSelect
func (c *TColorListBox) MultiSelect() bool {
    return ColorListBox_GetMultiSelect(c.instance)
}

// SetMultiSelect
func (c *TColorListBox) SetMultiSelect(value bool) {
    ColorListBox_SetMultiSelect(c.instance, value)
}

// SelCount
func (c *TColorListBox) SelCount() int32 {
    return ColorListBox_GetSelCount(c.instance)
}

// ItemIndex
func (c *TColorListBox) ItemIndex() int32 {
    return ColorListBox_GetItemIndex(c.instance)
}

// SetItemIndex
func (c *TColorListBox) SetItemIndex(value int32) {
    ColorListBox_SetItemIndex(c.instance, value)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (c *TColorListBox) DockClientCount() int32 {
    return ColorListBox_GetDockClientCount(c.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TColorListBox) DockSite() bool {
    return ColorListBox_GetDockSite(c.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TColorListBox) SetDockSite(value bool) {
    ColorListBox_SetDockSite(c.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (c *TColorListBox) AlignDisabled() bool {
    return ColorListBox_GetAlignDisabled(c.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TColorListBox) MouseInClient() bool {
    return ColorListBox_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TColorListBox) VisibleDockClientCount() int32 {
    return ColorListBox_GetVisibleDockClientCount(c.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TColorListBox) Brush() *TBrush {
    return BrushFromInst(ColorListBox_GetBrush(c.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TColorListBox) ControlCount() int32 {
    return ColorListBox_GetControlCount(c.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TColorListBox) Handle() HWND {
    return ColorListBox_GetHandle(c.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TColorListBox) ParentWindow() HWND {
    return ColorListBox_GetParentWindow(c.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TColorListBox) SetParentWindow(value HWND) {
    ColorListBox_SetParentWindow(c.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (c *TColorListBox) UseDockManager() bool {
    return ColorListBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (c *TColorListBox) SetUseDockManager(value bool) {
    ColorListBox_SetUseDockManager(c.instance, value)
}

// Action
func (c *TColorListBox) Action() *TAction {
    return ActionFromInst(ColorListBox_GetAction(c.instance))
}

// SetAction
func (c *TColorListBox) SetAction(value IComponent) {
    ColorListBox_SetAction(c.instance, CheckPtr(value))
}

// BoundsRect
func (c *TColorListBox) BoundsRect() TRect {
    return ColorListBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TColorListBox) SetBoundsRect(value TRect) {
    ColorListBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (c *TColorListBox) ClientHeight() int32 {
    return ColorListBox_GetClientHeight(c.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (c *TColorListBox) SetClientHeight(value int32) {
    ColorListBox_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TColorListBox) ClientOrigin() TPoint {
    return ColorListBox_GetClientOrigin(c.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TColorListBox) ClientRect() TRect {
    return ColorListBox_GetClientRect(c.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TColorListBox) ClientWidth() int32 {
    return ColorListBox_GetClientWidth(c.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TColorListBox) SetClientWidth(value int32) {
    ColorListBox_SetClientWidth(c.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (c *TColorListBox) ControlState() TControlState {
    return ColorListBox_GetControlState(c.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (c *TColorListBox) SetControlState(value TControlState) {
    ColorListBox_SetControlState(c.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (c *TColorListBox) ControlStyle() TControlStyle {
    return ColorListBox_GetControlStyle(c.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (c *TColorListBox) SetControlStyle(value TControlStyle) {
    ColorListBox_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TColorListBox) ExplicitLeft() int32 {
    return ColorListBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TColorListBox) ExplicitTop() int32 {
    return ColorListBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TColorListBox) ExplicitWidth() int32 {
    return ColorListBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TColorListBox) ExplicitHeight() int32 {
    return ColorListBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TColorListBox) Floating() bool {
    return ColorListBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TColorListBox) Parent() *TWinControl {
    return WinControlFromInst(ColorListBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TColorListBox) SetParent(value IWinControl) {
    ColorListBox_SetParent(c.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (c *TColorListBox) AlignWithMargins() bool {
    return ColorListBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (c *TColorListBox) SetAlignWithMargins(value bool) {
    ColorListBox_SetAlignWithMargins(c.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (c *TColorListBox) Left() int32 {
    return ColorListBox_GetLeft(c.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (c *TColorListBox) SetLeft(value int32) {
    ColorListBox_SetLeft(c.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TColorListBox) Top() int32 {
    return ColorListBox_GetTop(c.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TColorListBox) SetTop(value int32) {
    ColorListBox_SetTop(c.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (c *TColorListBox) Width() int32 {
    return ColorListBox_GetWidth(c.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (c *TColorListBox) SetWidth(value int32) {
    ColorListBox_SetWidth(c.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (c *TColorListBox) Height() int32 {
    return ColorListBox_GetHeight(c.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (c *TColorListBox) SetHeight(value int32) {
    ColorListBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TColorListBox) Cursor() TCursor {
    return ColorListBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TColorListBox) SetCursor(value TCursor) {
    ColorListBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TColorListBox) Hint() string {
    return ColorListBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TColorListBox) SetHint(value string) {
    ColorListBox_SetHint(c.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (c *TColorListBox) Margins() *TMargins {
    return MarginsFromInst(ColorListBox_GetMargins(c.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (c *TColorListBox) SetMargins(value *TMargins) {
    ColorListBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (c *TColorListBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(ColorListBox_GetCustomHint(c.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (c *TColorListBox) SetCustomHint(value IComponent) {
    ColorListBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TColorListBox) ComponentCount() int32 {
    return ColorListBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TColorListBox) ComponentIndex() int32 {
    return ColorListBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TColorListBox) SetComponentIndex(value int32) {
    ColorListBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TColorListBox) Owner() *TComponent {
    return ComponentFromInst(ColorListBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TColorListBox) Name() string {
    return ColorListBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TColorListBox) SetName(value string) {
    ColorListBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TColorListBox) Tag() int {
    return ColorListBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TColorListBox) SetTag(value int) {
    ColorListBox_SetTag(c.instance, value)
}

// Colors
func (c *TColorListBox) Colors(Index int32) TColor {
    return ColorListBox_GetColors(c.instance, Index)
}

// ColorNames
func (c *TColorListBox) ColorNames(Index int32) string {
    return ColorListBox_GetColorNames(c.instance, Index)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TColorListBox) DockClients(Index int32) *TControl {
    return ControlFromInst(ColorListBox_GetDockClients(c.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (c *TColorListBox) Controls(Index int32) *TControl {
    return ControlFromInst(ColorListBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TColorListBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ColorListBox_GetComponents(c.instance, AIndex))
}

