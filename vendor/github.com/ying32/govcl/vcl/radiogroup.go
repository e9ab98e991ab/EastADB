
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

type TRadioGroup struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewRadioGroup
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewRadioGroup(owner IComponent) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = RadioGroup_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RadioGroupFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func RadioGroupFromInst(inst uintptr) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = inst
    r.ptr = unsafe.Pointer(inst)
    return r
}

// RadioGroupFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func RadioGroupFromObj(obj IObject) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = CheckPtr(obj)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RadioGroupFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func RadioGroupFromUnsafePointer(ptr unsafe.Pointer) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = uintptr(ptr)
    r.ptr = ptr
    return r
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (r *TRadioGroup) Free() {
    if r.instance != 0 {
        RadioGroup_Free(r.instance)
        r.instance = 0
        r.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (r *TRadioGroup) Instance() uintptr {
    return r.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (r *TRadioGroup) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (r *TRadioGroup) IsValid() bool {
    return r.instance != 0
}

// TRadioGroupClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TRadioGroupClass() TClass {
    return RadioGroup_StaticClassType()
}

// FlipChildren
func (r *TRadioGroup) FlipChildren(AllLevels bool) {
    RadioGroup_FlipChildren(r.instance, AllLevels)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (r *TRadioGroup) CanFocus() bool {
    return RadioGroup_CanFocus(r.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (r *TRadioGroup) ContainsControl(Control IControl) bool {
    return RadioGroup_ContainsControl(r.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (r *TRadioGroup) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(RadioGroup_ControlAtPos(r.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (r *TRadioGroup) DisableAlign() {
    RadioGroup_DisableAlign(r.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (r *TRadioGroup) EnableAlign() {
    RadioGroup_EnableAlign(r.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (r *TRadioGroup) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(RadioGroup_FindChildControl(r.instance, ControlName))
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (r *TRadioGroup) Focused() bool {
    return RadioGroup_Focused(r.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (r *TRadioGroup) HandleAllocated() bool {
    return RadioGroup_HandleAllocated(r.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (r *TRadioGroup) InsertControl(AControl IControl) {
    RadioGroup_InsertControl(r.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (r *TRadioGroup) Invalidate() {
    RadioGroup_Invalidate(r.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (r *TRadioGroup) PaintTo(DC HDC, X int32, Y int32) {
    RadioGroup_PaintTo(r.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (r *TRadioGroup) RemoveControl(AControl IControl) {
    RadioGroup_RemoveControl(r.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (r *TRadioGroup) Realign() {
    RadioGroup_Realign(r.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (r *TRadioGroup) Repaint() {
    RadioGroup_Repaint(r.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (r *TRadioGroup) ScaleBy(M int32, D int32) {
    RadioGroup_ScaleBy(r.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (r *TRadioGroup) ScrollBy(DeltaX int32, DeltaY int32) {
    RadioGroup_ScrollBy(r.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (r *TRadioGroup) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RadioGroup_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (r *TRadioGroup) SetFocus() {
    RadioGroup_SetFocus(r.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (r *TRadioGroup) Update() {
    RadioGroup_Update(r.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (r *TRadioGroup) UpdateControlState() {
    RadioGroup_UpdateControlState(r.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (r *TRadioGroup) BringToFront() {
    RadioGroup_BringToFront(r.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (r *TRadioGroup) ClientToScreen(Point TPoint) TPoint {
    return RadioGroup_ClientToScreen(r.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (r *TRadioGroup) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RadioGroup_ClientToParent(r.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (r *TRadioGroup) Dragging() bool {
    return RadioGroup_Dragging(r.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (r *TRadioGroup) HasParent() bool {
    return RadioGroup_HasParent(r.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (r *TRadioGroup) Hide() {
    RadioGroup_Hide(r.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (r *TRadioGroup) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RadioGroup_Perform(r.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (r *TRadioGroup) Refresh() {
    RadioGroup_Refresh(r.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (r *TRadioGroup) ScreenToClient(Point TPoint) TPoint {
    return RadioGroup_ScreenToClient(r.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (r *TRadioGroup) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RadioGroup_ParentToClient(r.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (r *TRadioGroup) SendToBack() {
    RadioGroup_SendToBack(r.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (r *TRadioGroup) Show() {
    RadioGroup_Show(r.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (r *TRadioGroup) GetTextBuf(Buffer string, BufSize int32) int32 {
    return RadioGroup_GetTextBuf(r.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (r *TRadioGroup) GetTextLen() int32 {
    return RadioGroup_GetTextLen(r.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (r *TRadioGroup) SetTextBuf(Buffer string) {
    RadioGroup_SetTextBuf(r.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (r *TRadioGroup) FindComponent(AName string) *TComponent {
    return ComponentFromInst(RadioGroup_FindComponent(r.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (r *TRadioGroup) GetNamePath() string {
    return RadioGroup_GetNamePath(r.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (r *TRadioGroup) Assign(Source IObject) {
    RadioGroup_Assign(r.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (r *TRadioGroup) DisposeOf() {
    RadioGroup_DisposeOf(r.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (r *TRadioGroup) ClassType() TClass {
    return RadioGroup_ClassType(r.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (r *TRadioGroup) ClassName() string {
    return RadioGroup_ClassName(r.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (r *TRadioGroup) InstanceSize() int32 {
    return RadioGroup_InstanceSize(r.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (r *TRadioGroup) InheritsFrom(AClass TClass) bool {
    return RadioGroup_InheritsFrom(r.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (r *TRadioGroup) Equals(Obj IObject) bool {
    return RadioGroup_Equals(r.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (r *TRadioGroup) GetHashCode() int32 {
    return RadioGroup_GetHashCode(r.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (r *TRadioGroup) ToString() string {
    return RadioGroup_ToString(r.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (r *TRadioGroup) Align() TAlign {
    return RadioGroup_GetAlign(r.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (r *TRadioGroup) SetAlign(value TAlign) {
    RadioGroup_SetAlign(r.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (r *TRadioGroup) Anchors() TAnchors {
    return RadioGroup_GetAnchors(r.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (r *TRadioGroup) SetAnchors(value TAnchors) {
    RadioGroup_SetAnchors(r.instance, value)
}

// BiDiMode
func (r *TRadioGroup) BiDiMode() TBiDiMode {
    return RadioGroup_GetBiDiMode(r.instance)
}

// SetBiDiMode
func (r *TRadioGroup) SetBiDiMode(value TBiDiMode) {
    RadioGroup_SetBiDiMode(r.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (r *TRadioGroup) Caption() string {
    return RadioGroup_GetCaption(r.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (r *TRadioGroup) SetCaption(value string) {
    RadioGroup_SetCaption(r.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (r *TRadioGroup) Color() TColor {
    return RadioGroup_GetColor(r.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (r *TRadioGroup) SetColor(value TColor) {
    RadioGroup_SetColor(r.instance, value)
}

// Columns
func (r *TRadioGroup) Columns() int32 {
    return RadioGroup_GetColumns(r.instance)
}

// SetColumns
func (r *TRadioGroup) SetColumns(value int32) {
    RadioGroup_SetColumns(r.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (r *TRadioGroup) DoubleBuffered() bool {
    return RadioGroup_GetDoubleBuffered(r.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (r *TRadioGroup) SetDoubleBuffered(value bool) {
    RadioGroup_SetDoubleBuffered(r.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (r *TRadioGroup) DragCursor() TCursor {
    return RadioGroup_GetDragCursor(r.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (r *TRadioGroup) SetDragCursor(value TCursor) {
    RadioGroup_SetDragCursor(r.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (r *TRadioGroup) DragKind() TDragKind {
    return RadioGroup_GetDragKind(r.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (r *TRadioGroup) SetDragKind(value TDragKind) {
    RadioGroup_SetDragKind(r.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (r *TRadioGroup) DragMode() TDragMode {
    return RadioGroup_GetDragMode(r.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (r *TRadioGroup) SetDragMode(value TDragMode) {
    RadioGroup_SetDragMode(r.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (r *TRadioGroup) Enabled() bool {
    return RadioGroup_GetEnabled(r.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (r *TRadioGroup) SetEnabled(value bool) {
    RadioGroup_SetEnabled(r.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (r *TRadioGroup) Font() *TFont {
    return FontFromInst(RadioGroup_GetFont(r.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (r *TRadioGroup) SetFont(value *TFont) {
    RadioGroup_SetFont(r.instance, CheckPtr(value))
}

// ItemIndex
func (r *TRadioGroup) ItemIndex() int32 {
    return RadioGroup_GetItemIndex(r.instance)
}

// SetItemIndex
func (r *TRadioGroup) SetItemIndex(value int32) {
    RadioGroup_SetItemIndex(r.instance, value)
}

// Items
func (r *TRadioGroup) Items() *TStrings {
    return StringsFromInst(RadioGroup_GetItems(r.instance))
}

// SetItems
func (r *TRadioGroup) SetItems(value IObject) {
    RadioGroup_SetItems(r.instance, CheckPtr(value))
}

// ParentBackground
func (r *TRadioGroup) ParentBackground() bool {
    return RadioGroup_GetParentBackground(r.instance)
}

// SetParentBackground
func (r *TRadioGroup) SetParentBackground(value bool) {
    RadioGroup_SetParentBackground(r.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (r *TRadioGroup) ParentColor() bool {
    return RadioGroup_GetParentColor(r.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (r *TRadioGroup) SetParentColor(value bool) {
    RadioGroup_SetParentColor(r.instance, value)
}

// ParentCtl3D
func (r *TRadioGroup) ParentCtl3D() bool {
    return RadioGroup_GetParentCtl3D(r.instance)
}

// SetParentCtl3D
func (r *TRadioGroup) SetParentCtl3D(value bool) {
    RadioGroup_SetParentCtl3D(r.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (r *TRadioGroup) ParentDoubleBuffered() bool {
    return RadioGroup_GetParentDoubleBuffered(r.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (r *TRadioGroup) SetParentDoubleBuffered(value bool) {
    RadioGroup_SetParentDoubleBuffered(r.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (r *TRadioGroup) ParentFont() bool {
    return RadioGroup_GetParentFont(r.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (r *TRadioGroup) SetParentFont(value bool) {
    RadioGroup_SetParentFont(r.instance, value)
}

// ParentShowHint
func (r *TRadioGroup) ParentShowHint() bool {
    return RadioGroup_GetParentShowHint(r.instance)
}

// SetParentShowHint
func (r *TRadioGroup) SetParentShowHint(value bool) {
    RadioGroup_SetParentShowHint(r.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (r *TRadioGroup) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(RadioGroup_GetPopupMenu(r.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (r *TRadioGroup) SetPopupMenu(value IComponent) {
    RadioGroup_SetPopupMenu(r.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (r *TRadioGroup) ShowHint() bool {
    return RadioGroup_GetShowHint(r.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (r *TRadioGroup) SetShowHint(value bool) {
    RadioGroup_SetShowHint(r.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (r *TRadioGroup) TabOrder() TTabOrder {
    return RadioGroup_GetTabOrder(r.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (r *TRadioGroup) SetTabOrder(value TTabOrder) {
    RadioGroup_SetTabOrder(r.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (r *TRadioGroup) TabStop() bool {
    return RadioGroup_GetTabStop(r.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (r *TRadioGroup) SetTabStop(value bool) {
    RadioGroup_SetTabStop(r.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (r *TRadioGroup) Visible() bool {
    return RadioGroup_GetVisible(r.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (r *TRadioGroup) SetVisible(value bool) {
    RadioGroup_SetVisible(r.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (r *TRadioGroup) StyleElements() TStyleElements {
    return RadioGroup_GetStyleElements(r.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (r *TRadioGroup) SetStyleElements(value TStyleElements) {
    RadioGroup_SetStyleElements(r.instance, value)
}

// WordWrap
// CN: 获取自动换行。
// EN: Get Automatic line break.
func (r *TRadioGroup) WordWrap() bool {
    return RadioGroup_GetWordWrap(r.instance)
}

// SetWordWrap
// CN: 设置自动换行。
// EN: Set Automatic line break.
func (r *TRadioGroup) SetWordWrap(value bool) {
    RadioGroup_SetWordWrap(r.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (r *TRadioGroup) SetOnClick(fn TNotifyEvent) {
    RadioGroup_SetOnClick(r.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (r *TRadioGroup) SetOnContextPopup(fn TContextPopupEvent) {
    RadioGroup_SetOnContextPopup(r.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (r *TRadioGroup) SetOnDragDrop(fn TDragDropEvent) {
    RadioGroup_SetOnDragDrop(r.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (r *TRadioGroup) SetOnDragOver(fn TDragOverEvent) {
    RadioGroup_SetOnDragOver(r.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (r *TRadioGroup) SetOnEndDock(fn TEndDragEvent) {
    RadioGroup_SetOnEndDock(r.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (r *TRadioGroup) SetOnEndDrag(fn TEndDragEvent) {
    RadioGroup_SetOnEndDrag(r.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (r *TRadioGroup) SetOnEnter(fn TNotifyEvent) {
    RadioGroup_SetOnEnter(r.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (r *TRadioGroup) SetOnExit(fn TNotifyEvent) {
    RadioGroup_SetOnExit(r.instance, fn)
}

// SetOnGesture
func (r *TRadioGroup) SetOnGesture(fn TGestureEvent) {
    RadioGroup_SetOnGesture(r.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (r *TRadioGroup) SetOnStartDock(fn TStartDockEvent) {
    RadioGroup_SetOnStartDock(r.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (r *TRadioGroup) DockClientCount() int32 {
    return RadioGroup_GetDockClientCount(r.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (r *TRadioGroup) DockSite() bool {
    return RadioGroup_GetDockSite(r.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (r *TRadioGroup) SetDockSite(value bool) {
    RadioGroup_SetDockSite(r.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (r *TRadioGroup) AlignDisabled() bool {
    return RadioGroup_GetAlignDisabled(r.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (r *TRadioGroup) MouseInClient() bool {
    return RadioGroup_GetMouseInClient(r.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (r *TRadioGroup) VisibleDockClientCount() int32 {
    return RadioGroup_GetVisibleDockClientCount(r.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (r *TRadioGroup) Brush() *TBrush {
    return BrushFromInst(RadioGroup_GetBrush(r.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (r *TRadioGroup) ControlCount() int32 {
    return RadioGroup_GetControlCount(r.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (r *TRadioGroup) Handle() HWND {
    return RadioGroup_GetHandle(r.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (r *TRadioGroup) ParentWindow() HWND {
    return RadioGroup_GetParentWindow(r.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (r *TRadioGroup) SetParentWindow(value HWND) {
    RadioGroup_SetParentWindow(r.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (r *TRadioGroup) UseDockManager() bool {
    return RadioGroup_GetUseDockManager(r.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (r *TRadioGroup) SetUseDockManager(value bool) {
    RadioGroup_SetUseDockManager(r.instance, value)
}

// Action
func (r *TRadioGroup) Action() *TAction {
    return ActionFromInst(RadioGroup_GetAction(r.instance))
}

// SetAction
func (r *TRadioGroup) SetAction(value IComponent) {
    RadioGroup_SetAction(r.instance, CheckPtr(value))
}

// BoundsRect
func (r *TRadioGroup) BoundsRect() TRect {
    return RadioGroup_GetBoundsRect(r.instance)
}

// SetBoundsRect
func (r *TRadioGroup) SetBoundsRect(value TRect) {
    RadioGroup_SetBoundsRect(r.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (r *TRadioGroup) ClientHeight() int32 {
    return RadioGroup_GetClientHeight(r.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (r *TRadioGroup) SetClientHeight(value int32) {
    RadioGroup_SetClientHeight(r.instance, value)
}

// ClientOrigin
func (r *TRadioGroup) ClientOrigin() TPoint {
    return RadioGroup_GetClientOrigin(r.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (r *TRadioGroup) ClientRect() TRect {
    return RadioGroup_GetClientRect(r.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (r *TRadioGroup) ClientWidth() int32 {
    return RadioGroup_GetClientWidth(r.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (r *TRadioGroup) SetClientWidth(value int32) {
    RadioGroup_SetClientWidth(r.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (r *TRadioGroup) ControlState() TControlState {
    return RadioGroup_GetControlState(r.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (r *TRadioGroup) SetControlState(value TControlState) {
    RadioGroup_SetControlState(r.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (r *TRadioGroup) ControlStyle() TControlStyle {
    return RadioGroup_GetControlStyle(r.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (r *TRadioGroup) SetControlStyle(value TControlStyle) {
    RadioGroup_SetControlStyle(r.instance, value)
}

// ExplicitLeft
func (r *TRadioGroup) ExplicitLeft() int32 {
    return RadioGroup_GetExplicitLeft(r.instance)
}

// ExplicitTop
func (r *TRadioGroup) ExplicitTop() int32 {
    return RadioGroup_GetExplicitTop(r.instance)
}

// ExplicitWidth
func (r *TRadioGroup) ExplicitWidth() int32 {
    return RadioGroup_GetExplicitWidth(r.instance)
}

// ExplicitHeight
func (r *TRadioGroup) ExplicitHeight() int32 {
    return RadioGroup_GetExplicitHeight(r.instance)
}

// Floating
func (r *TRadioGroup) Floating() bool {
    return RadioGroup_GetFloating(r.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (r *TRadioGroup) Parent() *TWinControl {
    return WinControlFromInst(RadioGroup_GetParent(r.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (r *TRadioGroup) SetParent(value IWinControl) {
    RadioGroup_SetParent(r.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (r *TRadioGroup) AlignWithMargins() bool {
    return RadioGroup_GetAlignWithMargins(r.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (r *TRadioGroup) SetAlignWithMargins(value bool) {
    RadioGroup_SetAlignWithMargins(r.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (r *TRadioGroup) Left() int32 {
    return RadioGroup_GetLeft(r.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (r *TRadioGroup) SetLeft(value int32) {
    RadioGroup_SetLeft(r.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (r *TRadioGroup) Top() int32 {
    return RadioGroup_GetTop(r.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (r *TRadioGroup) SetTop(value int32) {
    RadioGroup_SetTop(r.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (r *TRadioGroup) Width() int32 {
    return RadioGroup_GetWidth(r.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (r *TRadioGroup) SetWidth(value int32) {
    RadioGroup_SetWidth(r.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (r *TRadioGroup) Height() int32 {
    return RadioGroup_GetHeight(r.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (r *TRadioGroup) SetHeight(value int32) {
    RadioGroup_SetHeight(r.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (r *TRadioGroup) Cursor() TCursor {
    return RadioGroup_GetCursor(r.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (r *TRadioGroup) SetCursor(value TCursor) {
    RadioGroup_SetCursor(r.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (r *TRadioGroup) Hint() string {
    return RadioGroup_GetHint(r.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (r *TRadioGroup) SetHint(value string) {
    RadioGroup_SetHint(r.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (r *TRadioGroup) Margins() *TMargins {
    return MarginsFromInst(RadioGroup_GetMargins(r.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (r *TRadioGroup) SetMargins(value *TMargins) {
    RadioGroup_SetMargins(r.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (r *TRadioGroup) CustomHint() *TCustomHint {
    return CustomHintFromInst(RadioGroup_GetCustomHint(r.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (r *TRadioGroup) SetCustomHint(value IComponent) {
    RadioGroup_SetCustomHint(r.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (r *TRadioGroup) ComponentCount() int32 {
    return RadioGroup_GetComponentCount(r.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (r *TRadioGroup) ComponentIndex() int32 {
    return RadioGroup_GetComponentIndex(r.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (r *TRadioGroup) SetComponentIndex(value int32) {
    RadioGroup_SetComponentIndex(r.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (r *TRadioGroup) Owner() *TComponent {
    return ComponentFromInst(RadioGroup_GetOwner(r.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (r *TRadioGroup) Name() string {
    return RadioGroup_GetName(r.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (r *TRadioGroup) SetName(value string) {
    RadioGroup_SetName(r.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (r *TRadioGroup) Tag() int {
    return RadioGroup_GetTag(r.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (r *TRadioGroup) SetTag(value int) {
    RadioGroup_SetTag(r.instance, value)
}

// Buttons
func (r *TRadioGroup) Buttons(Index int32) *TRadioButton {
    return RadioButtonFromInst(RadioGroup_GetButtons(r.instance, Index))
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (r *TRadioGroup) DockClients(Index int32) *TControl {
    return ControlFromInst(RadioGroup_GetDockClients(r.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (r *TRadioGroup) Controls(Index int32) *TControl {
    return ControlFromInst(RadioGroup_GetControls(r.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (r *TRadioGroup) Components(AIndex int32) *TComponent {
    return ComponentFromInst(RadioGroup_GetComponents(r.instance, AIndex))
}

