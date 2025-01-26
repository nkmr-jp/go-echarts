package opts

import "github.com/go-echarts/go-echarts/v2/types"

// DataZoom is the option set for a zoom component.
// dataZoom component is used for zooming a specific area, which enables user to
// investigate data in detail, or get an overview of the data, or get rid of outlier points.
// https://echarts.apache.org/en/option.html#dataZoom
type DataZoom struct {
	// Data zoom component of inside type, Options: "inside", "slider"
	Type string `json:"type" default:"inside"`

	// The start percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 0
	Start float32 `json:"start,omitempty"`

	// The end percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 100
	End float32 `json:"end,omitempty"`

	// Specify whether the layout of dataZoom component is horizontal or vertical. What's more, it indicates whether the horizontal axis or vertical axis is controlled by default in catesian coordinate system.
	//
	// Valid values:
	// 'horizontal': horizontal.
	// 'vertical': vertical.
	Orient string `json:"orient,omitempty"`

	// Specify the frame rate of views refreshing, with unit millisecond (ms).
	// If animation set as true and animationDurationUpdate set as bigger than 0,
	// you can keep throttle as the default value 100 (or set it as a value bigger than 0),
	// otherwise it might be not smooth when dragging.
	// If animation set as false or animationDurationUpdate set as 0, and data size is not very large,
	// and it seems to be not smooth when dragging, you can set throttle as 0 to improve that.
	Throttle float32 `json:"throttle,omitempty"`

	// Specify which xAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first xAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'horizontal'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`

	// Specify which yAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first yAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'vertical'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`

	// LabelFormatter is the formatter tool for the label.
	//
	// If it is a string, it will be a template. For instance, aaaa{value}bbbb, where {value} will be replaced by the value of actual data value.
	// It can also be a callback function. For example:
	//
	// /** @param {*} value If axis.type is 'category', `value` is the index of axis.data.
	//  *                   else `value` is current value.
	//  * @param {string} valueStr Inner formatted string.
	//  * @return {string} Returns the label formatted.
	//  labelFormatter: function (value, valueStr) {
	//     return 'aaa' + value + 'bbb';
	// }
	LabelFormatter string `json:"labelFormatter,omitempty"`

	// FilterMode Generally dataZoom component zoom or roam coordinate system
	// https://echarts.apache.org/en/option.html#dataZoom-inside.filterMode
	// through data filtering and set the windows of axes internally.
	// Possible values:
	//'filter': data that outside the window will be filtered, which may lead to some changes of windows of other axes. For each data item, it will be filtered if one of the relevant dimensions is out of the window.
	//'weakFilter': data that outside the window will be filtered, which may lead to some changes of windows of other axes. For each data item, it will be filtered only if all of the relevant dimensions are out of the same side of the window.
	//'empty': data that outside the window will be set to NaN, which will not lead to changes of windows of other axes.
	//'none': Do not filter data.
	FilterMode string `json:"filterMode,omitempty"`

	// Distance between toolbox component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between toolbox component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between toolbox component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`

	// Distance between toolbox component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// Width of text block
	Width float32 `json:"width,omitempty"`

	// Height of text block
	Height float32 `json:"height,omitempty"`

	// How to trigger zoom. Optional values:
	//
	// true：Mouse wheel triggers zoom.
	// false：Mouse wheel can not triggers zoom.
	// 'shift'：Holding shift and mouse wheel triggers zoom.
	// 'ctrl'：Holding ctrl and mouse wheel triggers zoom.
	// 'alt'：Holding alt and mouse wheel triggers zoom.
	ZoomOnMouseWheel any `json:"zoomOnMouseWheel,omitempty"`

	// How to trigger data window move. Optional values:
	//
	// true：Mouse move triggers data window move.
	// false：Mouse move can not triggers data window move.
	// 'shift'：Holding shift and mouse move triggers data window move.
	// 'ctrl'：Holding ctrl and mouse move triggers data window move.
	// 'alt'：Holding alt and mouse move triggers data window move.
	MoveOnMouseMove any `json:"moveOnMouseMove,omitempty"`

	// How to trigger data window move. Optional values:
	//
	// true：Mouse wheel triggers data window move.
	// false：Mouse wheel can not triggers data window move.
	// 'shift'：Holding shift and mouse wheel triggers data window move.
	// 'ctrl'：Holding ctrl and mouse wheel triggers data window move.
	// 'alt'：Holding alt and mouse wheel triggers data window move.
	MoveOnMouseWheel any `json:"moveOnMouseWheel,omitempty"`

	// Whether to prevent default behavior of mouse move event.
	PreventDefaultMouseMove types.Bool `json:"preventDefaultMouseMove,omitempty"`
}
