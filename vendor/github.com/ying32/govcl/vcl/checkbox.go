
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

type TCheckBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewCheckBox
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewCheckBox(owner IComponent) *TCheckBox {
    c := new(TCheckBox)
    c.instance = CheckBox_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CheckBoxFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func CheckBoxFromInst(inst uintptr) *TCheckBox {
    c := new(TCheckBox)
    c.instance = inst
    c.ptr = unsafe.Pointer(inst)
    return c
}

// CheckBoxFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func CheckBoxFromObj(obj IObject) *TCheckBox {
    c := new(TCheckBox)
    c.instance = CheckPtr(obj)
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// CheckBoxFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func CheckBoxFromUnsafePointer(ptr unsafe.Pointer) *TCheckBox {
    c := new(TCheckBox)
    c.instance = uintptr(ptr)
    c.ptr = ptr
    return c
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (c *TCheckBox) Free() {
    if c.instance != 0 {
        CheckBox_Free(c.instance)
        c.instance = 0
        c.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TCheckBox) Instance() uintptr {
    return c.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TCheckBox) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TCheckBox) IsValid() bool {
    return c.instance != 0
}

// TCheckBoxClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TCheckBoxClass() TClass {
    return CheckBox_StaticClassType()
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (c *TCheckBox) CanFocus() bool {
    return CheckBox_CanFocus(c.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (c *TCheckBox) ContainsControl(Control IControl) bool {
    return CheckBox_ContainsControl(c.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (c *TCheckBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(CheckBox_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (c *TCheckBox) DisableAlign() {
    CheckBox_DisableAlign(c.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (c *TCheckBox) EnableAlign() {
    CheckBox_EnableAlign(c.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (c *TCheckBox) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(CheckBox_FindChildControl(c.instance, ControlName))
}

// FlipChildren
func (c *TCheckBox) FlipChildren(AllLevels bool) {
    CheckBox_FlipChildren(c.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (c *TCheckBox) Focused() bool {
    return CheckBox_Focused(c.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (c *TCheckBox) HandleAllocated() bool {
    return CheckBox_HandleAllocated(c.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (c *TCheckBox) InsertControl(AControl IControl) {
    CheckBox_InsertControl(c.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (c *TCheckBox) Invalidate() {
    CheckBox_Invalidate(c.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (c *TCheckBox) PaintTo(DC HDC, X int32, Y int32) {
    CheckBox_PaintTo(c.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (c *TCheckBox) RemoveControl(AControl IControl) {
    CheckBox_RemoveControl(c.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (c *TCheckBox) Realign() {
    CheckBox_Realign(c.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (c *TCheckBox) Repaint() {
    CheckBox_Repaint(c.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (c *TCheckBox) ScaleBy(M int32, D int32) {
    CheckBox_ScaleBy(c.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (c *TCheckBox) ScrollBy(DeltaX int32, DeltaY int32) {
    CheckBox_ScrollBy(c.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (c *TCheckBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    CheckBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (c *TCheckBox) SetFocus() {
    CheckBox_SetFocus(c.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (c *TCheckBox) Update() {
    CheckBox_Update(c.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (c *TCheckBox) UpdateControlState() {
    CheckBox_UpdateControlState(c.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (c *TCheckBox) BringToFront() {
    CheckBox_BringToFront(c.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (c *TCheckBox) ClientToScreen(Point TPoint) TPoint {
    return CheckBox_ClientToScreen(c.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (c *TCheckBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return CheckBox_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (c *TCheckBox) Dragging() bool {
    return CheckBox_Dragging(c.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (c *TCheckBox) HasParent() bool {
    return CheckBox_HasParent(c.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (c *TCheckBox) Hide() {
    CheckBox_Hide(c.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (c *TCheckBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return CheckBox_Perform(c.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (c *TCheckBox) Refresh() {
    CheckBox_Refresh(c.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (c *TCheckBox) ScreenToClient(Point TPoint) TPoint {
    return CheckBox_ScreenToClient(c.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (c *TCheckBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return CheckBox_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (c *TCheckBox) SendToBack() {
    CheckBox_SendToBack(c.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (c *TCheckBox) Show() {
    CheckBox_Show(c.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (c *TCheckBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    return CheckBox_GetTextBuf(c.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (c *TCheckBox) GetTextLen() int32 {
    return CheckBox_GetTextLen(c.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (c *TCheckBox) SetTextBuf(Buffer string) {
    CheckBox_SetTextBuf(c.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (c *TCheckBox) FindComponent(AName string) *TComponent {
    return ComponentFromInst(CheckBox_FindComponent(c.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TCheckBox) GetNamePath() string {
    return CheckBox_GetNamePath(c.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TCheckBox) Assign(Source IObject) {
    CheckBox_Assign(c.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TCheckBox) DisposeOf() {
    CheckBox_DisposeOf(c.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TCheckBox) ClassType() TClass {
    return CheckBox_ClassType(c.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TCheckBox) ClassName() string {
    return CheckBox_ClassName(c.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TCheckBox) InstanceSize() int32 {
    return CheckBox_InstanceSize(c.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TCheckBox) InheritsFrom(AClass TClass) bool {
    return CheckBox_InheritsFrom(c.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TCheckBox) Equals(Obj IObject) bool {
    return CheckBox_Equals(c.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TCheckBox) GetHashCode() int32 {
    return CheckBox_GetHashCode(c.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (c *TCheckBox) ToString() string {
    return CheckBox_ToString(c.instance)
}

// Action
func (c *TCheckBox) Action() *TAction {
    return ActionFromInst(CheckBox_GetAction(c.instance))
}

// SetAction
func (c *TCheckBox) SetAction(value IComponent) {
    CheckBox_SetAction(c.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (c *TCheckBox) Align() TAlign {
    return CheckBox_GetAlign(c.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (c *TCheckBox) SetAlign(value TAlign) {
    CheckBox_SetAlign(c.instance, value)
}

// Alignment
// CN: 获取文字对齐。
// EN: Get Text alignment.
func (c *TCheckBox) Alignment() TLeftRight {
    return CheckBox_GetAlignment(c.instance)
}

// SetAlignment
// CN: 设置文字对齐。
// EN: Set Text alignment.
func (c *TCheckBox) SetAlignment(value TLeftRight) {
    CheckBox_SetAlignment(c.instance, value)
}

// AllowGrayed
func (c *TCheckBox) AllowGrayed() bool {
    return CheckBox_GetAllowGrayed(c.instance)
}

// SetAllowGrayed
func (c *TCheckBox) SetAllowGrayed(value bool) {
    CheckBox_SetAllowGrayed(c.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (c *TCheckBox) Anchors() TAnchors {
    return CheckBox_GetAnchors(c.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (c *TCheckBox) SetAnchors(value TAnchors) {
    CheckBox_SetAnchors(c.instance, value)
}

// BiDiMode
func (c *TCheckBox) BiDiMode() TBiDiMode {
    return CheckBox_GetBiDiMode(c.instance)
}

// SetBiDiMode
func (c *TCheckBox) SetBiDiMode(value TBiDiMode) {
    CheckBox_SetBiDiMode(c.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (c *TCheckBox) Caption() string {
    return CheckBox_GetCaption(c.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (c *TCheckBox) SetCaption(value string) {
    CheckBox_SetCaption(c.instance, value)
}

// Checked
// CN: 获取是否选中。
// EN: .
func (c *TCheckBox) Checked() bool {
    return CheckBox_GetChecked(c.instance)
}

// SetChecked
// CN: 设置是否选中。
// EN: .
func (c *TCheckBox) SetChecked(value bool) {
    CheckBox_SetChecked(c.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (c *TCheckBox) Color() TColor {
    return CheckBox_GetColor(c.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (c *TCheckBox) SetColor(value TColor) {
    CheckBox_SetColor(c.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (c *TCheckBox) DoubleBuffered() bool {
    return CheckBox_GetDoubleBuffered(c.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (c *TCheckBox) SetDoubleBuffered(value bool) {
    CheckBox_SetDoubleBuffered(c.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (c *TCheckBox) DragCursor() TCursor {
    return CheckBox_GetDragCursor(c.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (c *TCheckBox) SetDragCursor(value TCursor) {
    CheckBox_SetDragCursor(c.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (c *TCheckBox) DragKind() TDragKind {
    return CheckBox_GetDragKind(c.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (c *TCheckBox) SetDragKind(value TDragKind) {
    CheckBox_SetDragKind(c.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (c *TCheckBox) DragMode() TDragMode {
    return CheckBox_GetDragMode(c.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (c *TCheckBox) SetDragMode(value TDragMode) {
    CheckBox_SetDragMode(c.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (c *TCheckBox) Enabled() bool {
    return CheckBox_GetEnabled(c.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (c *TCheckBox) SetEnabled(value bool) {
    CheckBox_SetEnabled(c.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (c *TCheckBox) Font() *TFont {
    return FontFromInst(CheckBox_GetFont(c.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (c *TCheckBox) SetFont(value *TFont) {
    CheckBox_SetFont(c.instance, CheckPtr(value))
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (c *TCheckBox) ParentColor() bool {
    return CheckBox_GetParentColor(c.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (c *TCheckBox) SetParentColor(value bool) {
    CheckBox_SetParentColor(c.instance, value)
}

// ParentCtl3D
func (c *TCheckBox) ParentCtl3D() bool {
    return CheckBox_GetParentCtl3D(c.instance)
}

// SetParentCtl3D
func (c *TCheckBox) SetParentCtl3D(value bool) {
    CheckBox_SetParentCtl3D(c.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (c *TCheckBox) ParentDoubleBuffered() bool {
    return CheckBox_GetParentDoubleBuffered(c.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (c *TCheckBox) SetParentDoubleBuffered(value bool) {
    CheckBox_SetParentDoubleBuffered(c.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (c *TCheckBox) ParentFont() bool {
    return CheckBox_GetParentFont(c.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (c *TCheckBox) SetParentFont(value bool) {
    CheckBox_SetParentFont(c.instance, value)
}

// ParentShowHint
func (c *TCheckBox) ParentShowHint() bool {
    return CheckBox_GetParentShowHint(c.instance)
}

// SetParentShowHint
func (c *TCheckBox) SetParentShowHint(value bool) {
    CheckBox_SetParentShowHint(c.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (c *TCheckBox) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(CheckBox_GetPopupMenu(c.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (c *TCheckBox) SetPopupMenu(value IComponent) {
    CheckBox_SetPopupMenu(c.instance, CheckPtr(value))
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (c *TCheckBox) ShowHint() bool {
    return CheckBox_GetShowHint(c.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (c *TCheckBox) SetShowHint(value bool) {
    CheckBox_SetShowHint(c.instance, value)
}

// State
func (c *TCheckBox) State() TCheckBoxState {
    return CheckBox_GetState(c.instance)
}

// SetState
func (c *TCheckBox) SetState(value TCheckBoxState) {
    CheckBox_SetState(c.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (c *TCheckBox) TabOrder() TTabOrder {
    return CheckBox_GetTabOrder(c.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (c *TCheckBox) SetTabOrder(value TTabOrder) {
    CheckBox_SetTabOrder(c.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (c *TCheckBox) TabStop() bool {
    return CheckBox_GetTabStop(c.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (c *TCheckBox) SetTabStop(value bool) {
    CheckBox_SetTabStop(c.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (c *TCheckBox) Visible() bool {
    return CheckBox_GetVisible(c.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (c *TCheckBox) SetVisible(value bool) {
    CheckBox_SetVisible(c.instance, value)
}

// WordWrap
// CN: 获取自动换行。
// EN: Get Automatic line break.
func (c *TCheckBox) WordWrap() bool {
    return CheckBox_GetWordWrap(c.instance)
}

// SetWordWrap
// CN: 设置自动换行。
// EN: Set Automatic line break.
func (c *TCheckBox) SetWordWrap(value bool) {
    CheckBox_SetWordWrap(c.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (c *TCheckBox) StyleElements() TStyleElements {
    return CheckBox_GetStyleElements(c.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (c *TCheckBox) SetStyleElements(value TStyleElements) {
    CheckBox_SetStyleElements(c.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (c *TCheckBox) SetOnClick(fn TNotifyEvent) {
    CheckBox_SetOnClick(c.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (c *TCheckBox) SetOnContextPopup(fn TContextPopupEvent) {
    CheckBox_SetOnContextPopup(c.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (c *TCheckBox) SetOnDragDrop(fn TDragDropEvent) {
    CheckBox_SetOnDragDrop(c.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (c *TCheckBox) SetOnDragOver(fn TDragOverEvent) {
    CheckBox_SetOnDragOver(c.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (c *TCheckBox) SetOnEndDock(fn TEndDragEvent) {
    CheckBox_SetOnEndDock(c.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (c *TCheckBox) SetOnEndDrag(fn TEndDragEvent) {
    CheckBox_SetOnEndDrag(c.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (c *TCheckBox) SetOnEnter(fn TNotifyEvent) {
    CheckBox_SetOnEnter(c.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (c *TCheckBox) SetOnExit(fn TNotifyEvent) {
    CheckBox_SetOnExit(c.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (c *TCheckBox) SetOnKeyDown(fn TKeyEvent) {
    CheckBox_SetOnKeyDown(c.instance, fn)
}

// SetOnKeyPress
func (c *TCheckBox) SetOnKeyPress(fn TKeyPressEvent) {
    CheckBox_SetOnKeyPress(c.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (c *TCheckBox) SetOnKeyUp(fn TKeyEvent) {
    CheckBox_SetOnKeyUp(c.instance, fn)
}

// SetOnMouseActivate
func (c *TCheckBox) SetOnMouseActivate(fn TMouseActivateEvent) {
    CheckBox_SetOnMouseActivate(c.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (c *TCheckBox) SetOnMouseDown(fn TMouseEvent) {
    CheckBox_SetOnMouseDown(c.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (c *TCheckBox) SetOnMouseEnter(fn TNotifyEvent) {
    CheckBox_SetOnMouseEnter(c.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (c *TCheckBox) SetOnMouseLeave(fn TNotifyEvent) {
    CheckBox_SetOnMouseLeave(c.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (c *TCheckBox) SetOnMouseMove(fn TMouseMoveEvent) {
    CheckBox_SetOnMouseMove(c.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (c *TCheckBox) SetOnMouseUp(fn TMouseEvent) {
    CheckBox_SetOnMouseUp(c.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (c *TCheckBox) SetOnStartDock(fn TStartDockEvent) {
    CheckBox_SetOnStartDock(c.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (c *TCheckBox) DockClientCount() int32 {
    return CheckBox_GetDockClientCount(c.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (c *TCheckBox) DockSite() bool {
    return CheckBox_GetDockSite(c.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (c *TCheckBox) SetDockSite(value bool) {
    CheckBox_SetDockSite(c.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (c *TCheckBox) AlignDisabled() bool {
    return CheckBox_GetAlignDisabled(c.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (c *TCheckBox) MouseInClient() bool {
    return CheckBox_GetMouseInClient(c.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (c *TCheckBox) VisibleDockClientCount() int32 {
    return CheckBox_GetVisibleDockClientCount(c.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (c *TCheckBox) Brush() *TBrush {
    return BrushFromInst(CheckBox_GetBrush(c.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (c *TCheckBox) ControlCount() int32 {
    return CheckBox_GetControlCount(c.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (c *TCheckBox) Handle() HWND {
    return CheckBox_GetHandle(c.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (c *TCheckBox) ParentWindow() HWND {
    return CheckBox_GetParentWindow(c.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (c *TCheckBox) SetParentWindow(value HWND) {
    CheckBox_SetParentWindow(c.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (c *TCheckBox) UseDockManager() bool {
    return CheckBox_GetUseDockManager(c.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (c *TCheckBox) SetUseDockManager(value bool) {
    CheckBox_SetUseDockManager(c.instance, value)
}

// BoundsRect
func (c *TCheckBox) BoundsRect() TRect {
    return CheckBox_GetBoundsRect(c.instance)
}

// SetBoundsRect
func (c *TCheckBox) SetBoundsRect(value TRect) {
    CheckBox_SetBoundsRect(c.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (c *TCheckBox) ClientHeight() int32 {
    return CheckBox_GetClientHeight(c.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (c *TCheckBox) SetClientHeight(value int32) {
    CheckBox_SetClientHeight(c.instance, value)
}

// ClientOrigin
func (c *TCheckBox) ClientOrigin() TPoint {
    return CheckBox_GetClientOrigin(c.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (c *TCheckBox) ClientRect() TRect {
    return CheckBox_GetClientRect(c.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (c *TCheckBox) ClientWidth() int32 {
    return CheckBox_GetClientWidth(c.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (c *TCheckBox) SetClientWidth(value int32) {
    CheckBox_SetClientWidth(c.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (c *TCheckBox) ControlState() TControlState {
    return CheckBox_GetControlState(c.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (c *TCheckBox) SetControlState(value TControlState) {
    CheckBox_SetControlState(c.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (c *TCheckBox) ControlStyle() TControlStyle {
    return CheckBox_GetControlStyle(c.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (c *TCheckBox) SetControlStyle(value TControlStyle) {
    CheckBox_SetControlStyle(c.instance, value)
}

// ExplicitLeft
func (c *TCheckBox) ExplicitLeft() int32 {
    return CheckBox_GetExplicitLeft(c.instance)
}

// ExplicitTop
func (c *TCheckBox) ExplicitTop() int32 {
    return CheckBox_GetExplicitTop(c.instance)
}

// ExplicitWidth
func (c *TCheckBox) ExplicitWidth() int32 {
    return CheckBox_GetExplicitWidth(c.instance)
}

// ExplicitHeight
func (c *TCheckBox) ExplicitHeight() int32 {
    return CheckBox_GetExplicitHeight(c.instance)
}

// Floating
func (c *TCheckBox) Floating() bool {
    return CheckBox_GetFloating(c.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (c *TCheckBox) Parent() *TWinControl {
    return WinControlFromInst(CheckBox_GetParent(c.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (c *TCheckBox) SetParent(value IWinControl) {
    CheckBox_SetParent(c.instance, CheckPtr(value))
}

// SetOnGesture
func (c *TCheckBox) SetOnGesture(fn TGestureEvent) {
    CheckBox_SetOnGesture(c.instance, fn)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (c *TCheckBox) AlignWithMargins() bool {
    return CheckBox_GetAlignWithMargins(c.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (c *TCheckBox) SetAlignWithMargins(value bool) {
    CheckBox_SetAlignWithMargins(c.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (c *TCheckBox) Left() int32 {
    return CheckBox_GetLeft(c.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (c *TCheckBox) SetLeft(value int32) {
    CheckBox_SetLeft(c.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (c *TCheckBox) Top() int32 {
    return CheckBox_GetTop(c.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (c *TCheckBox) SetTop(value int32) {
    CheckBox_SetTop(c.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (c *TCheckBox) Width() int32 {
    return CheckBox_GetWidth(c.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (c *TCheckBox) SetWidth(value int32) {
    CheckBox_SetWidth(c.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (c *TCheckBox) Height() int32 {
    return CheckBox_GetHeight(c.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (c *TCheckBox) SetHeight(value int32) {
    CheckBox_SetHeight(c.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (c *TCheckBox) Cursor() TCursor {
    return CheckBox_GetCursor(c.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (c *TCheckBox) SetCursor(value TCursor) {
    CheckBox_SetCursor(c.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (c *TCheckBox) Hint() string {
    return CheckBox_GetHint(c.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (c *TCheckBox) SetHint(value string) {
    CheckBox_SetHint(c.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (c *TCheckBox) Margins() *TMargins {
    return MarginsFromInst(CheckBox_GetMargins(c.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (c *TCheckBox) SetMargins(value *TMargins) {
    CheckBox_SetMargins(c.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (c *TCheckBox) CustomHint() *TCustomHint {
    return CustomHintFromInst(CheckBox_GetCustomHint(c.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (c *TCheckBox) SetCustomHint(value IComponent) {
    CheckBox_SetCustomHint(c.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (c *TCheckBox) ComponentCount() int32 {
    return CheckBox_GetComponentCount(c.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (c *TCheckBox) ComponentIndex() int32 {
    return CheckBox_GetComponentIndex(c.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (c *TCheckBox) SetComponentIndex(value int32) {
    CheckBox_SetComponentIndex(c.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (c *TCheckBox) Owner() *TComponent {
    return ComponentFromInst(CheckBox_GetOwner(c.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (c *TCheckBox) Name() string {
    return CheckBox_GetName(c.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (c *TCheckBox) SetName(value string) {
    CheckBox_SetName(c.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (c *TCheckBox) Tag() int {
    return CheckBox_GetTag(c.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (c *TCheckBox) SetTag(value int) {
    CheckBox_SetTag(c.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (c *TCheckBox) DockClients(Index int32) *TControl {
    return ControlFromInst(CheckBox_GetDockClients(c.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (c *TCheckBox) Controls(Index int32) *TControl {
    return ControlFromInst(CheckBox_GetControls(c.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (c *TCheckBox) Components(AIndex int32) *TComponent {
    return ComponentFromInst(CheckBox_GetComponents(c.instance, AIndex))
}

