
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

type TTreeView struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTreeView
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTreeView(owner IComponent) *TTreeView {
    t := new(TTreeView)
    t.instance = TreeView_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TreeViewFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TreeViewFromInst(inst uintptr) *TTreeView {
    t := new(TTreeView)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TreeViewFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TreeViewFromObj(obj IObject) *TTreeView {
    t := new(TTreeView)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TreeViewFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TreeViewFromUnsafePointer(ptr unsafe.Pointer) *TTreeView {
    t := new(TTreeView)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTreeView) Free() {
    if t.instance != 0 {
        TreeView_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTreeView) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTreeView) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTreeView) IsValid() bool {
    return t.instance != 0
}

// TTreeViewClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTreeViewClass() TClass {
    return TreeView_StaticClassType()
}

// AlphaSort
func (t *TTreeView) AlphaSort(ARecurse bool) bool {
    return TreeView_AlphaSort(t.instance, ARecurse)
}

// FullCollapse
func (t *TTreeView) FullCollapse() {
    TreeView_FullCollapse(t.instance)
}

// FullExpand
func (t *TTreeView) FullExpand() {
    TreeView_FullExpand(t.instance)
}

// GetNodeAt
func (t *TTreeView) GetNodeAt(X int32, Y int32) *TTreeNode {
    return TreeNodeFromInst(TreeView_GetNodeAt(t.instance, X , Y))
}

// IsEditing
func (t *TTreeView) IsEditing() bool {
    return TreeView_IsEditing(t.instance)
}

// LoadFromFile
// CN: 从文件加载。
// EN: .
func (t *TTreeView) LoadFromFile(FileName string) {
    TreeView_LoadFromFile(t.instance, FileName)
}

// LoadFromStream
// CN: 文件流加载。
// EN: .
func (t *TTreeView) LoadFromStream(Stream IObject) {
    TreeView_LoadFromStream(t.instance, CheckPtr(Stream))
}

// SaveToFile
// CN: 保存至文件。
// EN: .
func (t *TTreeView) SaveToFile(FileName string) {
    TreeView_SaveToFile(t.instance, FileName)
}

// SaveToStream
// CN: 保存至流。
// EN: .
func (t *TTreeView) SaveToStream(Stream IObject) {
    TreeView_SaveToStream(t.instance, CheckPtr(Stream))
}

// Deselect
func (t *TTreeView) Deselect(Node *TTreeNode) {
    TreeView_Deselect(t.instance, CheckPtr(Node))
}

// Subselect
func (t *TTreeView) Subselect(Node *TTreeNode, Validate bool) {
    TreeView_Subselect(t.instance, CheckPtr(Node), Validate)
}

// ClearSelection
// CN: 清除选择。
// EN: .
func (t *TTreeView) ClearSelection(KeepPrimary bool) {
    TreeView_ClearSelection(t.instance, KeepPrimary)
}

// FindNextToSelect
func (t *TTreeView) FindNextToSelect() *TTreeNode {
    return TreeNodeFromInst(TreeView_FindNextToSelect(t.instance))
}

// CustomSort
func (t *TTreeView) CustomSort(SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
    return TreeView_CustomSort(t.instance, SortProc , Data , ARecurse)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (t *TTreeView) CanFocus() bool {
    return TreeView_CanFocus(t.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (t *TTreeView) ContainsControl(Control IControl) bool {
    return TreeView_ContainsControl(t.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (t *TTreeView) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(TreeView_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (t *TTreeView) DisableAlign() {
    TreeView_DisableAlign(t.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (t *TTreeView) EnableAlign() {
    TreeView_EnableAlign(t.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (t *TTreeView) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(TreeView_FindChildControl(t.instance, ControlName))
}

// FlipChildren
func (t *TTreeView) FlipChildren(AllLevels bool) {
    TreeView_FlipChildren(t.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (t *TTreeView) Focused() bool {
    return TreeView_Focused(t.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (t *TTreeView) HandleAllocated() bool {
    return TreeView_HandleAllocated(t.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (t *TTreeView) InsertControl(AControl IControl) {
    TreeView_InsertControl(t.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (t *TTreeView) Invalidate() {
    TreeView_Invalidate(t.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (t *TTreeView) PaintTo(DC HDC, X int32, Y int32) {
    TreeView_PaintTo(t.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (t *TTreeView) RemoveControl(AControl IControl) {
    TreeView_RemoveControl(t.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (t *TTreeView) Realign() {
    TreeView_Realign(t.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (t *TTreeView) Repaint() {
    TreeView_Repaint(t.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (t *TTreeView) ScaleBy(M int32, D int32) {
    TreeView_ScaleBy(t.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (t *TTreeView) ScrollBy(DeltaX int32, DeltaY int32) {
    TreeView_ScrollBy(t.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (t *TTreeView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TreeView_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (t *TTreeView) SetFocus() {
    TreeView_SetFocus(t.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (t *TTreeView) Update() {
    TreeView_Update(t.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (t *TTreeView) UpdateControlState() {
    TreeView_UpdateControlState(t.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (t *TTreeView) BringToFront() {
    TreeView_BringToFront(t.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (t *TTreeView) ClientToScreen(Point TPoint) TPoint {
    return TreeView_ClientToScreen(t.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (t *TTreeView) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return TreeView_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (t *TTreeView) Dragging() bool {
    return TreeView_Dragging(t.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TTreeView) HasParent() bool {
    return TreeView_HasParent(t.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (t *TTreeView) Hide() {
    TreeView_Hide(t.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (t *TTreeView) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TreeView_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (t *TTreeView) Refresh() {
    TreeView_Refresh(t.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (t *TTreeView) ScreenToClient(Point TPoint) TPoint {
    return TreeView_ScreenToClient(t.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (t *TTreeView) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return TreeView_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (t *TTreeView) SendToBack() {
    TreeView_SendToBack(t.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (t *TTreeView) Show() {
    TreeView_Show(t.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (t *TTreeView) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TreeView_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (t *TTreeView) GetTextLen() int32 {
    return TreeView_GetTextLen(t.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (t *TTreeView) SetTextBuf(Buffer string) {
    TreeView_SetTextBuf(t.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TTreeView) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TreeView_FindComponent(t.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTreeView) GetNamePath() string {
    return TreeView_GetNamePath(t.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTreeView) Assign(Source IObject) {
    TreeView_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTreeView) DisposeOf() {
    TreeView_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTreeView) ClassType() TClass {
    return TreeView_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTreeView) ClassName() string {
    return TreeView_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTreeView) InstanceSize() int32 {
    return TreeView_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTreeView) InheritsFrom(AClass TClass) bool {
    return TreeView_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTreeView) Equals(Obj IObject) bool {
    return TreeView_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTreeView) GetHashCode() int32 {
    return TreeView_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTreeView) ToString() string {
    return TreeView_ToString(t.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (t *TTreeView) Align() TAlign {
    return TreeView_GetAlign(t.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (t *TTreeView) SetAlign(value TAlign) {
    TreeView_SetAlign(t.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (t *TTreeView) Anchors() TAnchors {
    return TreeView_GetAnchors(t.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (t *TTreeView) SetAnchors(value TAnchors) {
    TreeView_SetAnchors(t.instance, value)
}

// AutoExpand
func (t *TTreeView) AutoExpand() bool {
    return TreeView_GetAutoExpand(t.instance)
}

// SetAutoExpand
func (t *TTreeView) SetAutoExpand(value bool) {
    TreeView_SetAutoExpand(t.instance, value)
}

// BevelEdges
func (t *TTreeView) BevelEdges() TBevelEdges {
    return TreeView_GetBevelEdges(t.instance)
}

// SetBevelEdges
func (t *TTreeView) SetBevelEdges(value TBevelEdges) {
    TreeView_SetBevelEdges(t.instance, value)
}

// BevelInner
func (t *TTreeView) BevelInner() TBevelCut {
    return TreeView_GetBevelInner(t.instance)
}

// SetBevelInner
func (t *TTreeView) SetBevelInner(value TBevelCut) {
    TreeView_SetBevelInner(t.instance, value)
}

// BevelOuter
func (t *TTreeView) BevelOuter() TBevelCut {
    return TreeView_GetBevelOuter(t.instance)
}

// SetBevelOuter
func (t *TTreeView) SetBevelOuter(value TBevelCut) {
    TreeView_SetBevelOuter(t.instance, value)
}

// BevelKind
func (t *TTreeView) BevelKind() TBevelKind {
    return TreeView_GetBevelKind(t.instance)
}

// SetBevelKind
func (t *TTreeView) SetBevelKind(value TBevelKind) {
    TreeView_SetBevelKind(t.instance, value)
}

// BiDiMode
func (t *TTreeView) BiDiMode() TBiDiMode {
    return TreeView_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TTreeView) SetBiDiMode(value TBiDiMode) {
    TreeView_SetBiDiMode(t.instance, value)
}

// BorderStyle
// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (t *TTreeView) BorderStyle() TBorderStyle {
    return TreeView_GetBorderStyle(t.instance)
}

// SetBorderStyle
// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (t *TTreeView) SetBorderStyle(value TBorderStyle) {
    TreeView_SetBorderStyle(t.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (t *TTreeView) BorderWidth() int32 {
    return TreeView_GetBorderWidth(t.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (t *TTreeView) SetBorderWidth(value int32) {
    TreeView_SetBorderWidth(t.instance, value)
}

// ChangeDelay
func (t *TTreeView) ChangeDelay() int32 {
    return TreeView_GetChangeDelay(t.instance)
}

// SetChangeDelay
func (t *TTreeView) SetChangeDelay(value int32) {
    TreeView_SetChangeDelay(t.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (t *TTreeView) Color() TColor {
    return TreeView_GetColor(t.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (t *TTreeView) SetColor(value TColor) {
    TreeView_SetColor(t.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (t *TTreeView) DoubleBuffered() bool {
    return TreeView_GetDoubleBuffered(t.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (t *TTreeView) SetDoubleBuffered(value bool) {
    TreeView_SetDoubleBuffered(t.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (t *TTreeView) DragKind() TDragKind {
    return TreeView_GetDragKind(t.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (t *TTreeView) SetDragKind(value TDragKind) {
    TreeView_SetDragKind(t.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (t *TTreeView) DragCursor() TCursor {
    return TreeView_GetDragCursor(t.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (t *TTreeView) SetDragCursor(value TCursor) {
    TreeView_SetDragCursor(t.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (t *TTreeView) DragMode() TDragMode {
    return TreeView_GetDragMode(t.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (t *TTreeView) SetDragMode(value TDragMode) {
    TreeView_SetDragMode(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTreeView) Enabled() bool {
    return TreeView_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTreeView) SetEnabled(value bool) {
    TreeView_SetEnabled(t.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (t *TTreeView) Font() *TFont {
    return FontFromInst(TreeView_GetFont(t.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (t *TTreeView) SetFont(value *TFont) {
    TreeView_SetFont(t.instance, CheckPtr(value))
}

// HideSelection
// CN: 获取隐藏选择。
// EN: .
func (t *TTreeView) HideSelection() bool {
    return TreeView_GetHideSelection(t.instance)
}

// SetHideSelection
// CN: 设置隐藏选择。
// EN: .
func (t *TTreeView) SetHideSelection(value bool) {
    TreeView_SetHideSelection(t.instance, value)
}

// HotTrack
func (t *TTreeView) HotTrack() bool {
    return TreeView_GetHotTrack(t.instance)
}

// SetHotTrack
func (t *TTreeView) SetHotTrack(value bool) {
    TreeView_SetHotTrack(t.instance, value)
}

// Images
// CN: 获取图标索引列表对象。
// EN: .
func (t *TTreeView) Images() *TImageList {
    return ImageListFromInst(TreeView_GetImages(t.instance))
}

// SetImages
// CN: 设置图标索引列表对象。
// EN: .
func (t *TTreeView) SetImages(value IComponent) {
    TreeView_SetImages(t.instance, CheckPtr(value))
}

// Indent
func (t *TTreeView) Indent() int32 {
    return TreeView_GetIndent(t.instance)
}

// SetIndent
func (t *TTreeView) SetIndent(value int32) {
    TreeView_SetIndent(t.instance, value)
}

// MultiSelect
func (t *TTreeView) MultiSelect() bool {
    return TreeView_GetMultiSelect(t.instance)
}

// SetMultiSelect
func (t *TTreeView) SetMultiSelect(value bool) {
    TreeView_SetMultiSelect(t.instance, value)
}

// MultiSelectStyle
func (t *TTreeView) MultiSelectStyle() TMultiSelectStyle {
    return TreeView_GetMultiSelectStyle(t.instance)
}

// SetMultiSelectStyle
func (t *TTreeView) SetMultiSelectStyle(value TMultiSelectStyle) {
    TreeView_SetMultiSelectStyle(t.instance, value)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (t *TTreeView) ParentColor() bool {
    return TreeView_GetParentColor(t.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (t *TTreeView) SetParentColor(value bool) {
    TreeView_SetParentColor(t.instance, value)
}

// ParentCtl3D
func (t *TTreeView) ParentCtl3D() bool {
    return TreeView_GetParentCtl3D(t.instance)
}

// SetParentCtl3D
func (t *TTreeView) SetParentCtl3D(value bool) {
    TreeView_SetParentCtl3D(t.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (t *TTreeView) ParentDoubleBuffered() bool {
    return TreeView_GetParentDoubleBuffered(t.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (t *TTreeView) SetParentDoubleBuffered(value bool) {
    TreeView_SetParentDoubleBuffered(t.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (t *TTreeView) ParentFont() bool {
    return TreeView_GetParentFont(t.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (t *TTreeView) SetParentFont(value bool) {
    TreeView_SetParentFont(t.instance, value)
}

// ParentShowHint
func (t *TTreeView) ParentShowHint() bool {
    return TreeView_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TTreeView) SetParentShowHint(value bool) {
    TreeView_SetParentShowHint(t.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (t *TTreeView) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TreeView_GetPopupMenu(t.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (t *TTreeView) SetPopupMenu(value IComponent) {
    TreeView_SetPopupMenu(t.instance, CheckPtr(value))
}

// ReadOnly
// CN: 获取只读。
// EN: .
func (t *TTreeView) ReadOnly() bool {
    return TreeView_GetReadOnly(t.instance)
}

// SetReadOnly
// CN: 设置只读。
// EN: .
func (t *TTreeView) SetReadOnly(value bool) {
    TreeView_SetReadOnly(t.instance, value)
}

// RightClickSelect
func (t *TTreeView) RightClickSelect() bool {
    return TreeView_GetRightClickSelect(t.instance)
}

// SetRightClickSelect
func (t *TTreeView) SetRightClickSelect(value bool) {
    TreeView_SetRightClickSelect(t.instance, value)
}

// RowSelect
func (t *TTreeView) RowSelect() bool {
    return TreeView_GetRowSelect(t.instance)
}

// SetRowSelect
func (t *TTreeView) SetRowSelect(value bool) {
    TreeView_SetRowSelect(t.instance, value)
}

// ShowButtons
func (t *TTreeView) ShowButtons() bool {
    return TreeView_GetShowButtons(t.instance)
}

// SetShowButtons
func (t *TTreeView) SetShowButtons(value bool) {
    TreeView_SetShowButtons(t.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (t *TTreeView) ShowHint() bool {
    return TreeView_GetShowHint(t.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (t *TTreeView) SetShowHint(value bool) {
    TreeView_SetShowHint(t.instance, value)
}

// ShowLines
func (t *TTreeView) ShowLines() bool {
    return TreeView_GetShowLines(t.instance)
}

// SetShowLines
func (t *TTreeView) SetShowLines(value bool) {
    TreeView_SetShowLines(t.instance, value)
}

// ShowRoot
func (t *TTreeView) ShowRoot() bool {
    return TreeView_GetShowRoot(t.instance)
}

// SetShowRoot
func (t *TTreeView) SetShowRoot(value bool) {
    TreeView_SetShowRoot(t.instance, value)
}

// SortType
func (t *TTreeView) SortType() TSortType {
    return TreeView_GetSortType(t.instance)
}

// SetSortType
func (t *TTreeView) SetSortType(value TSortType) {
    TreeView_SetSortType(t.instance, value)
}

// StateImages
func (t *TTreeView) StateImages() *TImageList {
    return ImageListFromInst(TreeView_GetStateImages(t.instance))
}

// SetStateImages
func (t *TTreeView) SetStateImages(value IComponent) {
    TreeView_SetStateImages(t.instance, CheckPtr(value))
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (t *TTreeView) TabOrder() TTabOrder {
    return TreeView_GetTabOrder(t.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (t *TTreeView) SetTabOrder(value TTabOrder) {
    TreeView_SetTabOrder(t.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (t *TTreeView) TabStop() bool {
    return TreeView_GetTabStop(t.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (t *TTreeView) SetTabStop(value bool) {
    TreeView_SetTabStop(t.instance, value)
}

// ToolTips
func (t *TTreeView) ToolTips() bool {
    return TreeView_GetToolTips(t.instance)
}

// SetToolTips
func (t *TTreeView) SetToolTips(value bool) {
    TreeView_SetToolTips(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TTreeView) Visible() bool {
    return TreeView_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TTreeView) SetVisible(value bool) {
    TreeView_SetVisible(t.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (t *TTreeView) StyleElements() TStyleElements {
    return TreeView_GetStyleElements(t.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (t *TTreeView) SetStyleElements(value TStyleElements) {
    TreeView_SetStyleElements(t.instance, value)
}

// SetOnAddition
func (t *TTreeView) SetOnAddition(fn TTVExpandedEvent) {
    TreeView_SetOnAddition(t.instance, fn)
}

// SetOnAdvancedCustomDraw
func (t *TTreeView) SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent) {
    TreeView_SetOnAdvancedCustomDraw(t.instance, fn)
}

// SetOnAdvancedCustomDrawItem
func (t *TTreeView) SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) {
    TreeView_SetOnAdvancedCustomDrawItem(t.instance, fn)
}

// SetOnCancelEdit
func (t *TTreeView) SetOnCancelEdit(fn TTVChangedEvent) {
    TreeView_SetOnCancelEdit(t.instance, fn)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (t *TTreeView) SetOnChange(fn TTVChangedEvent) {
    TreeView_SetOnChange(t.instance, fn)
}

// SetOnChanging
func (t *TTreeView) SetOnChanging(fn TTVChangingEvent) {
    TreeView_SetOnChanging(t.instance, fn)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (t *TTreeView) SetOnClick(fn TNotifyEvent) {
    TreeView_SetOnClick(t.instance, fn)
}

// SetOnCollapsed
func (t *TTreeView) SetOnCollapsed(fn TTVExpandedEvent) {
    TreeView_SetOnCollapsed(t.instance, fn)
}

// SetOnCollapsing
func (t *TTreeView) SetOnCollapsing(fn TTVCollapsingEvent) {
    TreeView_SetOnCollapsing(t.instance, fn)
}

// SetOnCompare
func (t *TTreeView) SetOnCompare(fn TTVCompareEvent) {
    TreeView_SetOnCompare(t.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (t *TTreeView) SetOnContextPopup(fn TContextPopupEvent) {
    TreeView_SetOnContextPopup(t.instance, fn)
}

// SetOnCustomDraw
func (t *TTreeView) SetOnCustomDraw(fn TTVCustomDrawEvent) {
    TreeView_SetOnCustomDraw(t.instance, fn)
}

// SetOnCustomDrawItem
func (t *TTreeView) SetOnCustomDrawItem(fn TTVCustomDrawItemEvent) {
    TreeView_SetOnCustomDrawItem(t.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (t *TTreeView) SetOnDblClick(fn TNotifyEvent) {
    TreeView_SetOnDblClick(t.instance, fn)
}

// SetOnDeletion
func (t *TTreeView) SetOnDeletion(fn TTVExpandedEvent) {
    TreeView_SetOnDeletion(t.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (t *TTreeView) SetOnDragDrop(fn TDragDropEvent) {
    TreeView_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (t *TTreeView) SetOnDragOver(fn TDragOverEvent) {
    TreeView_SetOnDragOver(t.instance, fn)
}

// SetOnEdited
func (t *TTreeView) SetOnEdited(fn TTVEditedEvent) {
    TreeView_SetOnEdited(t.instance, fn)
}

// SetOnEditing
func (t *TTreeView) SetOnEditing(fn TTVEditingEvent) {
    TreeView_SetOnEditing(t.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (t *TTreeView) SetOnEndDock(fn TEndDragEvent) {
    TreeView_SetOnEndDock(t.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (t *TTreeView) SetOnEndDrag(fn TEndDragEvent) {
    TreeView_SetOnEndDrag(t.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (t *TTreeView) SetOnEnter(fn TNotifyEvent) {
    TreeView_SetOnEnter(t.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (t *TTreeView) SetOnExit(fn TNotifyEvent) {
    TreeView_SetOnExit(t.instance, fn)
}

// SetOnExpanding
func (t *TTreeView) SetOnExpanding(fn TTVExpandingEvent) {
    TreeView_SetOnExpanding(t.instance, fn)
}

// SetOnExpanded
func (t *TTreeView) SetOnExpanded(fn TTVExpandedEvent) {
    TreeView_SetOnExpanded(t.instance, fn)
}

// SetOnGesture
func (t *TTreeView) SetOnGesture(fn TGestureEvent) {
    TreeView_SetOnGesture(t.instance, fn)
}

// SetOnGetImageIndex
func (t *TTreeView) SetOnGetImageIndex(fn TTVExpandedEvent) {
    TreeView_SetOnGetImageIndex(t.instance, fn)
}

// SetOnGetSelectedIndex
func (t *TTreeView) SetOnGetSelectedIndex(fn TTVExpandedEvent) {
    TreeView_SetOnGetSelectedIndex(t.instance, fn)
}

// SetOnHint
// CN: 设置鼠标悬停提示事件。
// EN: .
func (t *TTreeView) SetOnHint(fn TTVHintEvent) {
    TreeView_SetOnHint(t.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (t *TTreeView) SetOnKeyDown(fn TKeyEvent) {
    TreeView_SetOnKeyDown(t.instance, fn)
}

// SetOnKeyPress
func (t *TTreeView) SetOnKeyPress(fn TKeyPressEvent) {
    TreeView_SetOnKeyPress(t.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (t *TTreeView) SetOnKeyUp(fn TKeyEvent) {
    TreeView_SetOnKeyUp(t.instance, fn)
}

// SetOnMouseActivate
func (t *TTreeView) SetOnMouseActivate(fn TMouseActivateEvent) {
    TreeView_SetOnMouseActivate(t.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (t *TTreeView) SetOnMouseDown(fn TMouseEvent) {
    TreeView_SetOnMouseDown(t.instance, fn)
}

// SetOnMouseEnter
// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (t *TTreeView) SetOnMouseEnter(fn TNotifyEvent) {
    TreeView_SetOnMouseEnter(t.instance, fn)
}

// SetOnMouseLeave
// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (t *TTreeView) SetOnMouseLeave(fn TNotifyEvent) {
    TreeView_SetOnMouseLeave(t.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (t *TTreeView) SetOnMouseMove(fn TMouseMoveEvent) {
    TreeView_SetOnMouseMove(t.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (t *TTreeView) SetOnMouseUp(fn TMouseEvent) {
    TreeView_SetOnMouseUp(t.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (t *TTreeView) SetOnStartDock(fn TStartDockEvent) {
    TreeView_SetOnStartDock(t.instance, fn)
}

// Items
func (t *TTreeView) Items() *TTreeNodes {
    return TreeNodesFromInst(TreeView_GetItems(t.instance))
}

// SetItems
func (t *TTreeView) SetItems(value *TTreeNodes) {
    TreeView_SetItems(t.instance, CheckPtr(value))
}

// Canvas
// CN: 获取画布。
// EN: .
func (t *TTreeView) Canvas() *TCanvas {
    return CanvasFromInst(TreeView_GetCanvas(t.instance))
}

// DropTarget
func (t *TTreeView) DropTarget() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetDropTarget(t.instance))
}

// SetDropTarget
func (t *TTreeView) SetDropTarget(value *TTreeNode) {
    TreeView_SetDropTarget(t.instance, CheckPtr(value))
}

// Selected
func (t *TTreeView) Selected() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetSelected(t.instance))
}

// SetSelected
func (t *TTreeView) SetSelected(value *TTreeNode) {
    TreeView_SetSelected(t.instance, CheckPtr(value))
}

// TopItem
func (t *TTreeView) TopItem() *TTreeNode {
    return TreeNodeFromInst(TreeView_GetTopItem(t.instance))
}

// SetTopItem
func (t *TTreeView) SetTopItem(value *TTreeNode) {
    TreeView_SetTopItem(t.instance, CheckPtr(value))
}

// SelectionCount
func (t *TTreeView) SelectionCount() uint32 {
    return TreeView_GetSelectionCount(t.instance)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (t *TTreeView) DockClientCount() int32 {
    return TreeView_GetDockClientCount(t.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (t *TTreeView) DockSite() bool {
    return TreeView_GetDockSite(t.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (t *TTreeView) SetDockSite(value bool) {
    TreeView_SetDockSite(t.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (t *TTreeView) AlignDisabled() bool {
    return TreeView_GetAlignDisabled(t.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (t *TTreeView) MouseInClient() bool {
    return TreeView_GetMouseInClient(t.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (t *TTreeView) VisibleDockClientCount() int32 {
    return TreeView_GetVisibleDockClientCount(t.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (t *TTreeView) Brush() *TBrush {
    return BrushFromInst(TreeView_GetBrush(t.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (t *TTreeView) ControlCount() int32 {
    return TreeView_GetControlCount(t.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (t *TTreeView) Handle() HWND {
    return TreeView_GetHandle(t.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (t *TTreeView) ParentWindow() HWND {
    return TreeView_GetParentWindow(t.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (t *TTreeView) SetParentWindow(value HWND) {
    TreeView_SetParentWindow(t.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (t *TTreeView) UseDockManager() bool {
    return TreeView_GetUseDockManager(t.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (t *TTreeView) SetUseDockManager(value bool) {
    TreeView_SetUseDockManager(t.instance, value)
}

// Action
func (t *TTreeView) Action() *TAction {
    return ActionFromInst(TreeView_GetAction(t.instance))
}

// SetAction
func (t *TTreeView) SetAction(value IComponent) {
    TreeView_SetAction(t.instance, CheckPtr(value))
}

// BoundsRect
func (t *TTreeView) BoundsRect() TRect {
    return TreeView_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TTreeView) SetBoundsRect(value TRect) {
    TreeView_SetBoundsRect(t.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (t *TTreeView) ClientHeight() int32 {
    return TreeView_GetClientHeight(t.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (t *TTreeView) SetClientHeight(value int32) {
    TreeView_SetClientHeight(t.instance, value)
}

// ClientOrigin
func (t *TTreeView) ClientOrigin() TPoint {
    return TreeView_GetClientOrigin(t.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (t *TTreeView) ClientRect() TRect {
    return TreeView_GetClientRect(t.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (t *TTreeView) ClientWidth() int32 {
    return TreeView_GetClientWidth(t.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (t *TTreeView) SetClientWidth(value int32) {
    TreeView_SetClientWidth(t.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (t *TTreeView) ControlState() TControlState {
    return TreeView_GetControlState(t.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (t *TTreeView) SetControlState(value TControlState) {
    TreeView_SetControlState(t.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (t *TTreeView) ControlStyle() TControlStyle {
    return TreeView_GetControlStyle(t.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (t *TTreeView) SetControlStyle(value TControlStyle) {
    TreeView_SetControlStyle(t.instance, value)
}

// ExplicitLeft
func (t *TTreeView) ExplicitLeft() int32 {
    return TreeView_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TTreeView) ExplicitTop() int32 {
    return TreeView_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TTreeView) ExplicitWidth() int32 {
    return TreeView_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TTreeView) ExplicitHeight() int32 {
    return TreeView_GetExplicitHeight(t.instance)
}

// Floating
func (t *TTreeView) Floating() bool {
    return TreeView_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TTreeView) Parent() *TWinControl {
    return WinControlFromInst(TreeView_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TTreeView) SetParent(value IWinControl) {
    TreeView_SetParent(t.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (t *TTreeView) AlignWithMargins() bool {
    return TreeView_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (t *TTreeView) SetAlignWithMargins(value bool) {
    TreeView_SetAlignWithMargins(t.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (t *TTreeView) Left() int32 {
    return TreeView_GetLeft(t.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (t *TTreeView) SetLeft(value int32) {
    TreeView_SetLeft(t.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (t *TTreeView) Top() int32 {
    return TreeView_GetTop(t.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (t *TTreeView) SetTop(value int32) {
    TreeView_SetTop(t.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (t *TTreeView) Width() int32 {
    return TreeView_GetWidth(t.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (t *TTreeView) SetWidth(value int32) {
    TreeView_SetWidth(t.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (t *TTreeView) Height() int32 {
    return TreeView_GetHeight(t.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (t *TTreeView) SetHeight(value int32) {
    TreeView_SetHeight(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TTreeView) Cursor() TCursor {
    return TreeView_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TTreeView) SetCursor(value TCursor) {
    TreeView_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (t *TTreeView) Hint() string {
    return TreeView_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (t *TTreeView) SetHint(value string) {
    TreeView_SetHint(t.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (t *TTreeView) Margins() *TMargins {
    return MarginsFromInst(TreeView_GetMargins(t.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (t *TTreeView) SetMargins(value *TMargins) {
    TreeView_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (t *TTreeView) CustomHint() *TCustomHint {
    return CustomHintFromInst(TreeView_GetCustomHint(t.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (t *TTreeView) SetCustomHint(value IComponent) {
    TreeView_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTreeView) ComponentCount() int32 {
    return TreeView_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TTreeView) ComponentIndex() int32 {
    return TreeView_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TTreeView) SetComponentIndex(value int32) {
    TreeView_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTreeView) Owner() *TComponent {
    return ComponentFromInst(TreeView_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTreeView) Name() string {
    return TreeView_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTreeView) SetName(value string) {
    TreeView_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTreeView) Tag() int {
    return TreeView_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTreeView) SetTag(value int) {
    TreeView_SetTag(t.instance, value)
}

// Selections
func (t *TTreeView) Selections(Index int32) *TTreeNode {
    return TreeNodeFromInst(TreeView_GetSelections(t.instance, Index))
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (t *TTreeView) DockClients(Index int32) *TControl {
    return ControlFromInst(TreeView_GetDockClients(t.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (t *TTreeView) Controls(Index int32) *TControl {
    return ControlFromInst(TreeView_GetControls(t.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTreeView) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TreeView_GetComponents(t.instance, AIndex))
}

