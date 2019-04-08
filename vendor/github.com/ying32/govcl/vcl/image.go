
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

type TImage struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewImage
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewImage(owner IComponent) *TImage {
    i := new(TImage)
    i.instance = Image_Create(CheckPtr(owner))
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// ImageFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ImageFromInst(inst uintptr) *TImage {
    i := new(TImage)
    i.instance = inst
    i.ptr = unsafe.Pointer(inst)
    return i
}

// ImageFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ImageFromObj(obj IObject) *TImage {
    i := new(TImage)
    i.instance = CheckPtr(obj)
    i.ptr = unsafe.Pointer(i.instance)
    return i
}

// ImageFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ImageFromUnsafePointer(ptr unsafe.Pointer) *TImage {
    i := new(TImage)
    i.instance = uintptr(ptr)
    i.ptr = ptr
    return i
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (i *TImage) Free() {
    if i.instance != 0 {
        Image_Free(i.instance)
        i.instance = 0
        i.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (i *TImage) Instance() uintptr {
    return i.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (i *TImage) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (i *TImage) IsValid() bool {
    return i.instance != 0
}

// TImageClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TImageClass() TClass {
    return Image_StaticClassType()
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (i *TImage) BringToFront() {
    Image_BringToFront(i.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (i *TImage) ClientToScreen(Point TPoint) TPoint {
    return Image_ClientToScreen(i.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (i *TImage) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Image_ClientToParent(i.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (i *TImage) Dragging() bool {
    return Image_Dragging(i.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (i *TImage) HasParent() bool {
    return Image_HasParent(i.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (i *TImage) Hide() {
    Image_Hide(i.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (i *TImage) Invalidate() {
    Image_Invalidate(i.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (i *TImage) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Image_Perform(i.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (i *TImage) Refresh() {
    Image_Refresh(i.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (i *TImage) Repaint() {
    Image_Repaint(i.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (i *TImage) ScreenToClient(Point TPoint) TPoint {
    return Image_ScreenToClient(i.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (i *TImage) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Image_ParentToClient(i.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (i *TImage) SendToBack() {
    Image_SendToBack(i.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (i *TImage) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Image_SetBounds(i.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (i *TImage) Show() {
    Image_Show(i.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (i *TImage) Update() {
    Image_Update(i.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (i *TImage) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Image_GetTextBuf(i.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (i *TImage) GetTextLen() int32 {
    return Image_GetTextLen(i.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (i *TImage) SetTextBuf(Buffer string) {
    Image_SetTextBuf(i.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (i *TImage) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Image_FindComponent(i.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (i *TImage) GetNamePath() string {
    return Image_GetNamePath(i.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (i *TImage) Assign(Source IObject) {
    Image_Assign(i.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (i *TImage) DisposeOf() {
    Image_DisposeOf(i.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (i *TImage) ClassType() TClass {
    return Image_ClassType(i.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (i *TImage) ClassName() string {
    return Image_ClassName(i.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (i *TImage) InstanceSize() int32 {
    return Image_InstanceSize(i.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (i *TImage) InheritsFrom(AClass TClass) bool {
    return Image_InheritsFrom(i.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (i *TImage) Equals(Obj IObject) bool {
    return Image_Equals(i.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (i *TImage) GetHashCode() int32 {
    return Image_GetHashCode(i.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (i *TImage) ToString() string {
    return Image_ToString(i.instance)
}

// Canvas
// CN: 获取画布。
// EN: .
func (i *TImage) Canvas() *TCanvas {
    return CanvasFromInst(Image_GetCanvas(i.instance))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (i *TImage) Align() TAlign {
    return Image_GetAlign(i.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (i *TImage) SetAlign(value TAlign) {
    Image_SetAlign(i.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (i *TImage) Anchors() TAnchors {
    return Image_GetAnchors(i.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (i *TImage) SetAnchors(value TAnchors) {
    Image_SetAnchors(i.instance, value)
}

// AutoSize
// CN: 获取自动调整大小。
// EN: .
func (i *TImage) AutoSize() bool {
    return Image_GetAutoSize(i.instance)
}

// SetAutoSize
// CN: 设置自动调整大小。
// EN: .
func (i *TImage) SetAutoSize(value bool) {
    Image_SetAutoSize(i.instance, value)
}

// Center
func (i *TImage) Center() bool {
    return Image_GetCenter(i.instance)
}

// SetCenter
func (i *TImage) SetCenter(value bool) {
    Image_SetCenter(i.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (i *TImage) DragCursor() TCursor {
    return Image_GetDragCursor(i.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (i *TImage) SetDragCursor(value TCursor) {
    Image_SetDragCursor(i.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (i *TImage) DragKind() TDragKind {
    return Image_GetDragKind(i.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (i *TImage) SetDragKind(value TDragKind) {
    Image_SetDragKind(i.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (i *TImage) DragMode() TDragMode {
    return Image_GetDragMode(i.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (i *TImage) SetDragMode(value TDragMode) {
    Image_SetDragMode(i.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (i *TImage) Enabled() bool {
    return Image_GetEnabled(i.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (i *TImage) SetEnabled(value bool) {
    Image_SetEnabled(i.instance, value)
}

// IncrementalDisplay
func (i *TImage) IncrementalDisplay() bool {
    return Image_GetIncrementalDisplay(i.instance)
}

// SetIncrementalDisplay
func (i *TImage) SetIncrementalDisplay(value bool) {
    Image_SetIncrementalDisplay(i.instance, value)
}

// ParentShowHint
func (i *TImage) ParentShowHint() bool {
    return Image_GetParentShowHint(i.instance)
}

// SetParentShowHint
func (i *TImage) SetParentShowHint(value bool) {
    Image_SetParentShowHint(i.instance, value)
}

// Picture
func (i *TImage) Picture() *TPicture {
    return PictureFromInst(Image_GetPicture(i.instance))
}

// SetPicture
func (i *TImage) SetPicture(value *TPicture) {
    Image_SetPicture(i.instance, CheckPtr(value))
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (i *TImage) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(Image_GetPopupMenu(i.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (i *TImage) SetPopupMenu(value IComponent) {
    Image_SetPopupMenu(i.instance, CheckPtr(value))
}

// Proportional
func (i *TImage) Proportional() bool {
    return Image_GetProportional(i.instance)
}

// SetProportional
func (i *TImage) SetProportional(value bool) {
    Image_SetProportional(i.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (i *TImage) ShowHint() bool {
    return Image_GetShowHint(i.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (i *TImage) SetShowHint(value bool) {
    Image_SetShowHint(i.instance, value)
}

// Stretch
func (i *TImage) Stretch() bool {
    return Image_GetStretch(i.instance)
}

// SetStretch
func (i *TImage) SetStretch(value bool) {
    Image_SetStretch(i.instance, value)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (i *TImage) Transparent() bool {
    return Image_GetTransparent(i.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (i *TImage) SetTransparent(value bool) {
    Image_SetTransparent(i.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (i *TImage) Visible() bool {
    return Image_GetVisible(i.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (i *TImage) SetVisible(value bool) {
    Image_SetVisible(i.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (i *TImage) SetOnClick(fn TNotifyEvent) {
    Image_SetOnClick(i.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (i *TImage) SetOnContextPopup(fn TContextPopupEvent) {
    Image_SetOnContextPopup(i.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (i *TImage) SetOnDblClick(fn TNotifyEvent) {
    Image_SetOnDblClick(i.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (i *TImage) SetOnDragDrop(fn TDragDropEvent) {
    Image_SetOnDragDrop(i.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (i *TImage) SetOnDragOver(fn TDragOverEvent) {
    Image_SetOnDragOver(i.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (i *TImage) SetOnEndDock(fn TEndDragEvent) {
    Image_SetOnEndDock(i.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (i *TImage) SetOnEndDrag(fn TEndDragEvent) {
    Image_SetOnEndDrag(i.instance, fn)
}

// SetOnGesture
func (i *TImage) SetOnGesture(fn TGestureEvent) {
    Image_SetOnGesture(i.instance, fn)
}

// SetOnMouseActivate
func (i *TImage) SetOnMouseActivate(fn TMouseActivateEvent) {
    Image_SetOnMouseActivate(i.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (i *TImage) SetOnMouseDown(fn TMouseEvent) {
    Image_SetOnMouseDown(i.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (i *TImage) SetOnMouseEnter(fn TNotifyEvent) {
    Image_SetOnMouseEnter(i.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (i *TImage) SetOnMouseLeave(fn TNotifyEvent) {
    Image_SetOnMouseLeave(i.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (i *TImage) SetOnMouseMove(fn TMouseMoveEvent) {
    Image_SetOnMouseMove(i.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (i *TImage) SetOnMouseUp(fn TMouseEvent) {
    Image_SetOnMouseUp(i.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (i *TImage) SetOnStartDock(fn TStartDockEvent) {
    Image_SetOnStartDock(i.instance, fn)
}

// Action
func (i *TImage) Action() *TAction {
    return ActionFromInst(Image_GetAction(i.instance))
}

// SetAction
func (i *TImage) SetAction(value IComponent) {
    Image_SetAction(i.instance, CheckPtr(value))
}

// BiDiMode
func (i *TImage) BiDiMode() TBiDiMode {
    return Image_GetBiDiMode(i.instance)
}

// SetBiDiMode
func (i *TImage) SetBiDiMode(value TBiDiMode) {
    Image_SetBiDiMode(i.instance, value)
}

// BoundsRect
func (i *TImage) BoundsRect() TRect {
    return Image_GetBoundsRect(i.instance)
}

// SetBoundsRect
func (i *TImage) SetBoundsRect(value TRect) {
    Image_SetBoundsRect(i.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (i *TImage) ClientHeight() int32 {
    return Image_GetClientHeight(i.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (i *TImage) SetClientHeight(value int32) {
    Image_SetClientHeight(i.instance, value)
}

// ClientOrigin
func (i *TImage) ClientOrigin() TPoint {
    return Image_GetClientOrigin(i.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (i *TImage) ClientRect() TRect {
    return Image_GetClientRect(i.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (i *TImage) ClientWidth() int32 {
    return Image_GetClientWidth(i.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (i *TImage) SetClientWidth(value int32) {
    Image_SetClientWidth(i.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (i *TImage) ControlState() TControlState {
    return Image_GetControlState(i.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (i *TImage) SetControlState(value TControlState) {
    Image_SetControlState(i.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (i *TImage) ControlStyle() TControlStyle {
    return Image_GetControlStyle(i.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (i *TImage) SetControlStyle(value TControlStyle) {
    Image_SetControlStyle(i.instance, value)
}

// ExplicitLeft
func (i *TImage) ExplicitLeft() int32 {
    return Image_GetExplicitLeft(i.instance)
}

// ExplicitTop
func (i *TImage) ExplicitTop() int32 {
    return Image_GetExplicitTop(i.instance)
}

// ExplicitWidth
func (i *TImage) ExplicitWidth() int32 {
    return Image_GetExplicitWidth(i.instance)
}

// ExplicitHeight
func (i *TImage) ExplicitHeight() int32 {
    return Image_GetExplicitHeight(i.instance)
}

// Floating
func (i *TImage) Floating() bool {
    return Image_GetFloating(i.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (i *TImage) Parent() *TWinControl {
    return WinControlFromInst(Image_GetParent(i.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (i *TImage) SetParent(value IWinControl) {
    Image_SetParent(i.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (i *TImage) StyleElements() TStyleElements {
    return Image_GetStyleElements(i.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (i *TImage) SetStyleElements(value TStyleElements) {
    Image_SetStyleElements(i.instance, value)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (i *TImage) AlignWithMargins() bool {
    return Image_GetAlignWithMargins(i.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (i *TImage) SetAlignWithMargins(value bool) {
    Image_SetAlignWithMargins(i.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (i *TImage) Left() int32 {
    return Image_GetLeft(i.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (i *TImage) SetLeft(value int32) {
    Image_SetLeft(i.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (i *TImage) Top() int32 {
    return Image_GetTop(i.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (i *TImage) SetTop(value int32) {
    Image_SetTop(i.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (i *TImage) Width() int32 {
    return Image_GetWidth(i.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (i *TImage) SetWidth(value int32) {
    Image_SetWidth(i.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (i *TImage) Height() int32 {
    return Image_GetHeight(i.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (i *TImage) SetHeight(value int32) {
    Image_SetHeight(i.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (i *TImage) Cursor() TCursor {
    return Image_GetCursor(i.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (i *TImage) SetCursor(value TCursor) {
    Image_SetCursor(i.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (i *TImage) Hint() string {
    return Image_GetHint(i.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (i *TImage) SetHint(value string) {
    Image_SetHint(i.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (i *TImage) Margins() *TMargins {
    return MarginsFromInst(Image_GetMargins(i.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (i *TImage) SetMargins(value *TMargins) {
    Image_SetMargins(i.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (i *TImage) CustomHint() *TCustomHint {
    return CustomHintFromInst(Image_GetCustomHint(i.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (i *TImage) SetCustomHint(value IComponent) {
    Image_SetCustomHint(i.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (i *TImage) ComponentCount() int32 {
    return Image_GetComponentCount(i.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (i *TImage) ComponentIndex() int32 {
    return Image_GetComponentIndex(i.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (i *TImage) SetComponentIndex(value int32) {
    Image_SetComponentIndex(i.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (i *TImage) Owner() *TComponent {
    return ComponentFromInst(Image_GetOwner(i.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (i *TImage) Name() string {
    return Image_GetName(i.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (i *TImage) SetName(value string) {
    Image_SetName(i.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (i *TImage) Tag() int {
    return Image_GetTag(i.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (i *TImage) SetTag(value int) {
    Image_SetTag(i.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (i *TImage) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Image_GetComponents(i.instance, AIndex))
}

