// Copyright (c) 2018 Beta Kuang
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package glfw3

/*
#cgo pkg-config: glfw3
#include <stdlib.h>
#include <GLFW/glfw3.h>

// Gateway for error callback function in Go.
void _errorCallback(int, char*);

// Workaround due to that Go does not support const function params.
static void _errorCallbackConst(int err, const char* desc) {
	_errorCallback(err, desc);
}

// Sets an error callback.
static void goSetErrorCallback() {
	glfwSetErrorCallback(_errorCallbackConst);
}

// Removes the error callback.
static void goRemoveErrorCallback() {
	glfwSetErrorCallback(NULL);
}

// Gateway for monitor callback function in Go.
void _monitorCallback(GLFWmonitor*, int);

// Sets an monitor callback.
static void goSetMonitorCallback() {
	glfwSetMonitorCallback(_monitorCallback);
}

// Removes the monitor callback.
static void goRemoveMonitorCallback() {
	glfwSetMonitorCallback(NULL);
}
*/
import "C"
import "unsafe"

// GLFW version constants.
const (
	// VersionMajor : The major version number of the GLFW library.
	//
	// This is incremented when the API is changed in non-compatible ways.
	VersionMajor uint = 3
	// VersionMinor : The minor version number of the GLFW library.
	//
	// This is incremented when features are added to the API but it remains
	// backward-compatible.
	VersionMinor uint = 2
	// VersionRevision : The revision number of the GLFW library.
	//
	// This is incremented when a bug fix release is made that does not contain
	// any API changes.
	VersionRevision uint = 1
)

// Action is a key and button action.
type Action int

const (
	// Release : The key or mouse button was released.
	Release Action = 0
	// Press : The key or mouse button was pressed.
	Press Action = 1
	// Repeat : The key was held down until it repeated.
	Repeat Action = 2
)

// Key is a keyboard key.
//
// These key codes are inspired by the _USB HID Usage Tables v1.12_ (p. 53-60),
// but re-arranged to map to 7-bit ASCII for printable keys (function keys are
// put in the 256+ range).
//
// The naming of the key codes follow these rules:
//
// - The US keyboard layout is used
//
// - Names of printable alpha-numeric characters are used (e.g. "A", "R", "3",
// etc.)
//
// - For non-alphanumeric characters, Unicode:ish names are used (e.g. "Comma",
// "LeftSquareBracket", etc.). Note that some names do not correspond to the
// Unicode standard (usually for brevity)
//
// - Keys that lack a clear US mapping are named "Worldx"
//
// - For non-printable keys, custom names are used (e.g. "F4", "Backspace",
// etc.)
type Key int

// Keyboard keys.
const (
	// KeyUnknown : The unknown key.
	KeyUnknown Key = -1

	// Printable keys.

	KeySpace Key = 32
	// KeyApostrophe : "'".
	KeyApostrophe Key = 39
	// KeyComma : ",".
	KeyComma Key = 44
	// KeyMinus : "-".
	KeyMinus Key = 45
	// KeyPeriod : ".".
	KeyPeriod Key = 46
	// KeySlash : "/".
	KeySlash Key = 47
	Key0     Key = 48
	Key1     Key = 49
	Key2     Key = 50
	Key3     Key = 51
	Key4     Key = 52
	Key5     Key = 53
	Key6     Key = 54
	Key7     Key = 55
	Key8     Key = 56
	Key9     Key = 57
	// KeySemicolon : ";".
	KeySemicolon Key = 59
	// KeyEqual : "=".
	KeyEqual Key = 61
	KeyA     Key = 65
	KeyB     Key = 66
	KeyC     Key = 67
	KeyD     Key = 68
	KeyE     Key = 69
	KeyF     Key = 70
	KeyG     Key = 71
	KeyH     Key = 72
	KeyI     Key = 73
	KeyJ     Key = 74
	KeyK     Key = 75
	KeyL     Key = 76
	KeyM     Key = 77
	KeyN     Key = 78
	KeyO     Key = 79
	KeyP     Key = 80
	KeyQ     Key = 81
	KeyR     Key = 82
	KeyS     Key = 83
	KeyT     Key = 84
	KeyU     Key = 85
	KeyV     Key = 86
	KeyW     Key = 87
	KeyX     Key = 88
	KeyY     Key = 89
	KeyZ     Key = 90
	// KeyLeftBracket : "[".
	KeyLeftBracket Key = 91
	// KeyBackslash : "\".
	KeyBackslash Key = 92
	// KeyRightBracket : "]".
	KeyRightBracket Key = 93
	// KeyGraveAccent : "`".
	KeyGraveAccent Key = 96
	// KeyWorld1 : non-US #1.
	KeyWorld1 Key = 161
	// KeyWorld2 : non-US #2.
	KeyWorld2 Key = 162

	// Function keys.

	KeyEscape       Key = 256
	KeyEnter        Key = 257
	KeyTab          Key = 258
	KeyBackspace    Key = 259
	KeyInsert       Key = 260
	KeyDelete       Key = 261
	KeyRight        Key = 262
	KeyLeft         Key = 263
	KeyDown         Key = 264
	KeyUp           Key = 265
	KeyPageUp       Key = 266
	KeyPageDown     Key = 267
	KeyHome         Key = 268
	KeyEnd          Key = 269
	KeyCapsLock     Key = 280
	KeyScrollLock   Key = 281
	KeyNumLock      Key = 282
	KeyPrintScreen  Key = 283
	KeyPause        Key = 284
	KeyF1           Key = 290
	KeyF2           Key = 291
	KeyF3           Key = 292
	KeyF4           Key = 293
	KeyF5           Key = 294
	KeyF6           Key = 295
	KeyF7           Key = 296
	KeyF8           Key = 297
	KeyF9           Key = 298
	KeyF10          Key = 299
	KeyF11          Key = 300
	KeyF12          Key = 301
	KeyF13          Key = 302
	KeyF14          Key = 303
	KeyF15          Key = 304
	KeyF16          Key = 305
	KeyF17          Key = 306
	KeyF18          Key = 307
	KeyF19          Key = 308
	KeyF20          Key = 309
	KeyF21          Key = 310
	KeyF22          Key = 311
	KeyF23          Key = 312
	KeyF24          Key = 313
	KeyF25          Key = 314
	KeyKp0          Key = 320
	KeyKp1          Key = 321
	KeyKp2          Key = 322
	KeyKp3          Key = 323
	KeyKp4          Key = 324
	KeyKp5          Key = 325
	KeyKp6          Key = 326
	KeyKp7          Key = 327
	KeyKp8          Key = 328
	KeyKp9          Key = 329
	KeyKpDecimal    Key = 330
	KeyKpDivide     Key = 331
	KeyKpMultiply   Key = 332
	KeyKpSubtract   Key = 333
	KeyKpAdd        Key = 334
	KeyKpEnter      Key = 335
	KeyKpEqual      Key = 336
	KeyLeftShift    Key = 340
	KeyLeftControl  Key = 341
	KeyLeftAlt      Key = 342
	KeyLeftSuper    Key = 343
	KeyRightShift   Key = 344
	KeyRightControl Key = 345
	KeyRightAlt     Key = 346
	KeyRightSuper   Key = 347
	KeyMenu         Key = 348

	KeyLast Key = KeyMenu
)

// ModifierFlag is a modifier key flag.
type ModifierFlag int

const (
	// ModShift : If this bit is set one or more Shift keys were held down.
	ModShift ModifierFlag = 0x0001
	// ModControl : If this bit is set one or more Control keys were held down.
	ModControl ModifierFlag = 0x0002
	// ModAlt : If this bit is set one or more Alt keys were held down.
	ModAlt ModifierFlag = 0x0004
	// ModSuper : If this bit is set one or more Super keys were held down.
	ModSuper ModifierFlag = 0x0008
)

// Button is a mouse button.
type Button int

// Mouse buttons.
const (
	MouseButton1      Button = 0
	MouseButton2      Button = 1
	MouseButton3      Button = 2
	MouseButton4      Button = 3
	MouseButton5      Button = 4
	MouseButton6      Button = 5
	MouseButton7      Button = 6
	MouseButton8      Button = 7
	MouseButtonLast   Button = MouseButton8
	MouseButtonLeft   Button = MouseButton1
	MouseButtonRight  Button = MouseButton2
	MouseButtonMiddle Button = MouseButton3
)

// Joystick is a joystick.
type Joystick int

// Joysticks.
const (
	Joystick1    Joystick = 0
	Joystick2    Joystick = 1
	Joystick3    Joystick = 2
	Joystick4    Joystick = 3
	Joystick5    Joystick = 4
	Joystick6    Joystick = 5
	Joystick7    Joystick = 6
	Joystick8    Joystick = 7
	Joystick9    Joystick = 8
	Joystick10   Joystick = 9
	Joystick11   Joystick = 10
	Joystick12   Joystick = 11
	Joystick13   Joystick = 12
	Joystick14   Joystick = 13
	Joystick15   Joystick = 14
	Joystick16   Joystick = 15
	JoystickLast Joystick = Joystick16
)

// Error represents an error code.
type Error int

const (
	// NotInitialized : GLFW has not been initialized.
	//
	// This occurs if a GLFW function was called that must not be called unless
	// the library is initialized.
	//
	// Analysis
	//
	// This is an application programmer error. Initialize GLFW before calling
	// any function that requires initialization.
	NotInitialized Error = 0x00010001
	// NoCurrentContext : No context is current for this thread.
	//
	// This occurs if a GLFW function was called that needs and operates on the
	// current OpenGL or OpenGL ES contest but no context is current on the
	// calling thread. One such function is SwapInterval().
	//
	// Analysis
	//
	// This is an application programmer error. Ensure a context is current
	// before calling functions that require a current context.
	NoCurrentContext Error = 0x00010002
	// InvalidEnum : One or the arguments to the function was an invalid enum
	// value, for example requesting RedBits with GetWindowAttrib.
	//
	// Analysis
	//
	// This is an application programmer error. Fix the offending call.
	InvalidEnum Error = 0x00010003
	// InvalidValue : One of the arguments to the function was an invalid value,
	// for example requesting a non-existent OpenGL or OpenGL ES version like
	// 2.7.
	//
	// Requesting a valid but unavailable OpenGL or OpenGL ES version will
	// instead result in a VersionUnavailable error.
	//
	// Analysis
	//
	// This is an application programmer error. Fix the offending call.
	InvalidValue Error = 0x00010004
	// OutOfMemory : A memory allocation failed.
	//
	// Analysis
	//
	// This is a bug in GLFW3, GLFW or the underlying operating system. Report
	// the bug to our issue tracker (https://github.com/paperui/glfw3/issues) or
	// GLFW's (https://github.com/glfw/glfw/issues).
	OutOfMemory Error = 0x00010005
	// APIUnavailable : GLFW could not find support for the requested API on the
	// system.
	//
	// Analysis
	//
	// Thhe installed graphics driver does not support the requested API, or
	// does not support it via the chosen context creation backend. Below are a
	// few examples.
	//
	// Some pre-installed Windows graphics drivers do not support OpenGL. AMD
	// only supports OpenGL ES via EGL, while Nvidia and Intel only support it
	// via a WGL or GLX extension. OS X does not provide OpenGL ES at all. The
	// Mesa EGL, OpenGL and OpenGL ES libraries do not interface with the
	// Nvidia binary driver. Older graphics drivers do not support Vulkan.
	APIUnavailable Error = 0x00010006
	// VersionUnavailable : The requested OpenGL or OpenGL ES version (including
	// any requested context or framebuffer hints) is not available on this
	// machine.
	//
	// Analysis
	//
	// The machine does not support your requirements. If your application is
	// sufficiently flexible, downgrade your requirements and try again.
	// Otherwise, inform the user that their machine does not match your
	// requirements.
	//
	// Future invalid OpenGL and OpenGL ES versions, for example OpenGL 4.8 if
	// 5.0 comes out before the 4.x series gets that far, also fail with this
	// error and not InvalidValue, because GLFW cannot know what future versions
	// will exist.
	VersionUnavailable Error = 0x00010007
	// PlatformError : A platform-specific error occurred that does not match
	// any of the more specific categories.
	//
	// Analysis
	//
	// This is a bug or configuration error in GLFW3, GLFW, the underlying
	// operating system or its drivers, or a lack of required resources. Report
	// the issue to our issue tracker (https://github.com/paperui/glfw3/issues)
	// or GLFW's (https://github.com/glfw/glfw/issues).
	PlatformError Error = 0x00010008
	// FormatUnavailable : The requested format is not supported or available.
	//
	// If emitted during window creation, the requested pixel format is not
	// supported.
	//
	// If emitted when querying the clipboard, the contents of the clipboard
	// could not be converted to the requested format.
	//
	// Analysis
	//
	// If emmited during window creation, one or more hard constraints
	// (http://www.glfw.org/docs/latest/window_guide.html#window_hints_hard) did
	// not match any of the available pixel formats. If your application is
	// sufficiently flexible, downgrade your requirements and try again.
	// Otherwise, inform the user that their machine does not match your
	// requirements.
	//
	// If emitted when querying the clipboard, ignore the error or report it to
	// the user, as appropriate.
	FormatUnavailable Error = 0x00010009
	// NoWindowContext : A window that does not have an OpenGL or OpenGL ES was
	// passed to a function that requires it to have one.
	//
	// Analysis
	//
	// This is an application programmer error. Fix the offending call.
	NoWindowContext Error = 0x0001000A
)

// Hint is a bit field for creating windows and context.
type Hint int

// Hints.
const (
	// Window related hints.
	Focused     Hint = 0x00020001
	Iconified   Hint = 0x00020002
	Resizable   Hint = 0x00020003
	Visible     Hint = 0x00020004
	Decorated   Hint = 0x00020005
	AutoIconify Hint = 0x00020006
	Floating    Hint = 0x00020007
	Maximized   Hint = 0x00020008

	// Framebuffer related hints.
	RedBits        Hint = 0x00021001
	GreenBits      Hint = 0x00021002
	BlueBits       Hint = 0x00021003
	AlphaBits      Hint = 0x00021004
	DepthBits      Hint = 0x00021005
	StencilBits    Hint = 0x00021006
	AccumRedBits   Hint = 0x00021007
	AccumGreenBits Hint = 0x00021008
	AccumBlueBits  Hint = 0x00021009
	AccumAlphaBits Hint = 0x0002100A
	AuxBuffers     Hint = 0x0002100B
	Stereo         Hint = 0x0002100C
	Samples        Hint = 0x0002100D
	SRGBCapable    Hint = 0x0002100E
	Doublebuffer   Hint = 0x00021010

	// Monitor related hints.
	RefreshRate Hint = 0x0002100F

	// Context related hints.
	ClientAPI              Hint = 0x00022001
	ContextCreationAPI     Hint = 0x0002200B
	ContextVersionMajor    Hint = 0x00022002
	ContextVersionMinor    Hint = 0x00022003
	ContextRevision        Hint = 0x00022004
	ContextRobustness      Hint = 0x00022005
	OpenGLForwardCompat    Hint = 0x00022006
	OpenGLDebugContext     Hint = 0x00022007
	OpenGLProfile          Hint = 0x00022008
	ContextReleaseBehavior Hint = 0x00022009
	ContextNoError         Hint = 0x0002200A
)

// HintValue is the value for a hint.
type HintValue int

// Hint values.
const (
	// Common values.
	DontCare HintValue = -1

	// Values for ClientAPI.
	NoAPI       HintValue = 0
	OpenGLAPI   HintValue = 0x00030001
	OpenGLESAPI HintValue = 0x00030002

	// Values for ContextCreationAPI.
	NativeContextAPI HintValue = 0x00036001
	EGLContextAPI    HintValue = 0x00036002

	// Values for OpenGLProfile.
	OpenGLAnyProfile    HintValue = 0
	OpenGLCoreProfile   HintValue = 0x00032001
	OpenGLCompatProfile HintValue = 0x00032002

	// Values for ContextRobustness.
	NoRobustness        HintValue = 0
	NoResetNotification HintValue = 0x00031001
	LoseContextOnReset  HintValue = 0x00031002

	// Values for ContextReleaseBehavior.
	AnyReleaseBehavior   HintValue = 0
	ReleaseBehaviorFlush HintValue = 0x00035001
	ReleaseBehaviorNone  HintValue = 0x00035002
)

// InputMode specifies the input mode.
type InputMode int

// Input modes.
const (
	CursorMode             = 0x00033001
	StickyKeysMode         = 0x00033002
	StickyMouseButtonsMode = 0x00033003
)

// CursorModeValue specifies mode of the cursor when the input mode is
// CursorMode.
type CursorModeValue int

// Cursor mode values.
const (
	CursorNormal   CursorModeValue = 0x00034001
	CursorHidden   CursorModeValue = 0x00034002
	CursorDisabled CursorModeValue = 0x00034003
)

// ConnectionEvent represents whether a device is connected or disconnected
// from the system.
type ConnectionEvent int

// Connection events.
const (
	Connected    ConnectionEvent = 0x00040001
	Disconnected ConnectionEvent = 0x00040002
)

// Constants.
const (
	True  = 1
	False = 0
)

// CursorShape represent a standard cursor shape.
type CursorShape int

const (
	// ArrowCursor : The regular arrow cursor shape.
	ArrowCursor CursorShape = 0x00036001
	// IBeamCursor : The text input I-beam cursor shape.
	IBeamCursor CursorShape = 0x00036002
	// CrosshairCursor : The crosshair shape.
	CrosshairCursor CursorShape = 0x00036003
	// HandCursor : The hand shape.
	HandCursor CursorShape = 0x00036004
	// HResizeCursor : The horizontal resize arrow shape.
	HResizeCursor CursorShape = 0x00036005
	// VResizeCursor : The vertical resize arrow shape.
	VResizeCursor CursorShape = 0x00036006
)

// Monitor is an opaque monitor object.
type Monitor C.GLFWmonitor

// Window is an opaque window object.
type Window C.GLFWwindow

// Cursor is an opaque cursor object.
type Cursor *C.GLFWcursor

// ErrorCallback is a function type for error callbacks.
//
// err is an Error, and desc is a string describing the error.
type ErrorCallback func(err Error, desc string)

// WindowPosCallback is the function type for window position callbacks.
//
// win is the window that was moved. x and y are the new x- and y-coordinate,
// in screen coordinates, of the upper-left corner of the client area of the
// window.
type WindowPosCallback func(win *Window, x, y int)

// WindowSizeCallback is the function type for window resize callbacks.
//
// win is the window that was resized. width and height are the new width and
// height, in screen coordinates, of the window.
type WindowSizeCallback func(win *Window, width, height int)

// WindowCloseCallback is the function type for window close callbacks.
//
// win is the window that user attempted to close.
type WindowCloseCallback func(win *Window)

// WindowRefreshCallback is the function type for window content refresh
// callbacks.
//
// win is the window whose content needs to be refreshed.
type WindowRefreshCallback func(win *Window)

// WindowFocusCallback is the function type for window focus/defocus callbacks.
//
// win is the window that gained or lost input focus. focused is true if the
// window was given input focus, or false if it lost it.
type WindowFocusCallback func(win *Window, focused bool)

// WindowIconifyCallback is the function type for window iconify/restore
// callbacks.
//
// win is the window that was iconified or restored. iconified is true if the
// window was iconified, or false if it was restored.
type WindowIconifyCallback func(win *Window, iconified bool)

// FramebufferSizeCallback is the function type for framebuffer resize
// callbacks.
//
// win is the window whose framebuffer was resized. width and height are the new
// width and height, in pixels, of the framebuffer.
type FramebufferSizeCallback func(win *Window, width, height int)

// MouseButtonCallback is the function type for mouse button callbacks.
//
// win is the window that received the event. button is the mouse button that
// was pressed or released. actoin is one of Press or Release. mods is a bit
// field describing which modifier keys were held down.
type MouseButtonCallback func(win *Window, button Button, action Action, mods ModifierFlag)

// CursorPosCallback is the function type for cursor position callbacks.
//
// win is the window that received the event. x and y are the new x- and
// y-coordinate, relative to the left/top edge of the client area.
type CursorPosCallback func(win *Window, x, y float64)

// CursorEnterCallback is the function type for cursor enter/leave callbacks.
//
// win is the window that received the event. entered is true if the cursor
// entered the window's client area, or false if it left it.
type CursorEnterCallback func(win *Window, entered bool)

// ScrollCallback is the function type for scroll callbacks.
//
// win is the window that received the event. xOffset and yOffset are the scroll
// offsets along the x- and y-axis.
type ScrollCallback func(win *Window, xOffset, yOffset float64)

// KeyCallback is the function type for keyboard key callbacks.
//
// win is the window that received the event. key is the keyboard key that was
// pressed or released. scancode is the system-specific scancode of the key.
// action is Press, Release or Repeat. mods is a bit field describing which
// modifier keys were held down.
type KeyCallback func(win *Window, key Key, scancode int, action Action, mods ModifierFlag)

// CharCallback is the function type for Unicode character callbacks.
//
// win is the window that received the event. codepoint is the Unicode code
// point of the character.
type CharCallback func(win *Window, codepoint rune)

// CharModsCallback is the function type for Unicode character with modifiers
// callbacks. It is called for each input characters, regardless of what
// modifier keys are held down.
//
// win is the window that received the event. codepoint is the Unicode code
// point of the character. mods is a bit field describing which modifier keys
// were held down.
type CharModsCallback func(win *Window, codepoint rune, mods ModifierFlag)

// DropCallback is the function type for file drop callbacks.
//
// win is the window that received the event. paths is the UTF-8 encoded file
// and/or directory path names.
type DropCallback func(win *Window, paths []string)

// MonitorCallback is the function type for monitor configuration callbacks.
//
// monitor is the monitor that was connected or disconnected. event is one of
// Connected or Disconnected.
type MonitorCallback func(monitor *Monitor, event int)

// JoystickCallback is the function type for joystick configuration callbacks.
//
// joy is the joystick that was connected or disconnected. event is one of
// Connected or Disconnected.
type JoystickCallback func(joy Joystick, event int)

// VideoMode describes a single video mode.
type VideoMode struct {
	// Width : The width, in screen coordinates, of the video mode.
	Width int
	// Height : The height, in screen coordinates, of the video mode.
	Height int
	// RedBits : The bit depth of the red channel of the video mode.
	RedBits int
	// GreenBits : The bit depth of the green channel of the video mode.
	GreenBits int
	// BlueBits : The bit depth of the blue channel of the video mode.
	BlueBits int
	// RefreshRate : The refresh rate, in Hz, of the video mode.
	RefreshRate int
}

// GammaRamp describes the gamma ramp for a monitor.
type GammaRamp struct {
	// Red : An array of value describing the response of the red channel.
	Red []uint16
	// Green : An array of value describing the response of the green channel.
	Green []uint16
	// Blue : An array of value describing the response of the blue channel.
	Blue []uint16
}

// Image is the Image data.
type Image struct {
	// Width : The width, in pixels, of this image.
	Width int
	// Height : The height, in pixels, of this image.
	Height int
	// Pixels : The pixel data of this image, arranged left-to-right,
	// top-to-bottom.
	Pixels []uint8
}

// Context is an entry point of all GLFW APIs that require initialization.
type Context struct {
	errorCallback   ErrorCallback
	monitorCallback MonitorCallback
}

// Init initializes the GLFW library.
//
// This function initializes the GLFW library. Before most GLFW functions can be
// used, GLFW must be initialized, and before an application terminates GLFW
// should be terminated in order to free any resources allocated during or
// after initialization.
//
// If this function fails, it calls Terminate() before returning. If it
// succeeds, you should call Terminate() before the application exits.
//
// Additional calls to this function after successful initialization but before
// termination will return true immediately.
//
// Returns a Context if successful, or nil if an error occurred. Possible errors
// include PlatformError.
//
// On OS X this function will change the current directory of the application to
// the Contents/Resources subdirectory of the application's bundle, if present.
// This can be disabled with a compile-time option
// (http://www.glfw.org/docs/latest/compile_guide.html#compile_options_osx).
//
// This function must only be called from the main thread.
func Init() *Context {
	if int(C.glfwInit()) == 1 {
		currentContext = new(Context)
		return currentContext
	}
	return nil
}

// Terminate terminates the GLFW library.
//
// This function destroys all remaining windows and cursors, restores any
// modified gamma ramps and frees any other allocated resources. Once this
// function is called, you must again call Init() successfully before you will
// be able to use most GLFW functions.
//
// If GLFW has been successfully initialized, this function should be called
// before the application exits. If initialization fails, there is no need to
// call this function, as it is called by Init() before it returns failure.
//
// Possible errors include PlatformError.
//
// The contexts of any remaining windows must not be current on any other thread
// when this function is called.
//
// This function must not be called from a callback.
//
// This function must only be called from the main thread.
func (c *Context) Terminate() {
	C.glfwTerminate()
}

var currentContext *Context

// GetVersion retrieves the version of the GLFW library.
//
// This function retrieves the major, minor and revision numbers of the GLFW
// library. It is intended for when you are using GLFW as a shared library and
// want to ensure that you are using the minimum required version.
//
// This function may be called from any thread.
func GetVersion() (major, minor, rev int) {
	var cMajor, cMinor, cRev C.int
	C.glfwGetVersion(&cMajor, &cMinor, &cRev)
	major, minor, rev = int(cMajor), int(cMinor), int(cRev)
	return
}

// GetVersionString returns a string describing the compile-time configuration.
//
// This function returns the compile-time generated version string
// (http://www.glfw.org/docs/latest/intro_guide.html#intro_version_string) of
// the GLFW library binary. It describes the version, platform, compiler and any
// platform-specific compile-time options. It should not be confused with the
// OpenGL or OpenGL ES version string, queried with glGetString.
//
// Do not use the version string to parse the GLFW library version. The
// GetVersion() function provides the version of the running library binary in
// numerical format.
//
// This function may be called from any thread.
func GetVersionString() string {
	return C.GoString(C.glfwGetVersionString())
}

// SetErrorCallback sets the error callback.
//
// This function sets the error callback, which is called with an error code and
// a humen-readable description each time a GLFW error occurs.
//
// The error callback is called on the thread where the error occurred. If you
// are using GLFW from multiple thread, your error callback needs to be written
// accordingly.
//
// Once set, the error callback remains set even after the library has been
// terminated.
//
// callback is the new callback, or nil to remove the currently set callback.
//
// Returns the previously set callback, or nil if no callback was set.
//
// This function must only be called from the main thread.
func (c *Context) SetErrorCallback(callback ErrorCallback) ErrorCallback {
	previousCallback := c.errorCallback
	c.errorCallback = callback
	if callback != nil {
		C.goSetErrorCallback()
	} else {
		C.goRemoveErrorCallback()
	}
	return previousCallback
}

//export _errorCallback
func _errorCallback(cErr C.int, cDesc *C.char) {
	if currentContext != nil && currentContext.errorCallback != nil {
		err := Error(cErr)
		desc := C.GoString(cDesc)
		currentContext.errorCallback(err, desc)
	}
}

// GetMonitors returns the currently connected monitors.
//
// This function returns an array of handles for all currently connected
// monitors. The primary monitor is always first in the returned array. If no
// monitors were found or if an error occurred, this function returns nil.
//
// Possible errors include NotInitialized.
//
// This function must only be called from the main thread.
func (c *Context) GetMonitors() []*Monitor {
	var cCount C.int
	cMonitors := C.glfwGetMonitors(&cCount)
	if unsafe.Pointer(cMonitors) != C.NULL {
		count := int(cCount)
		monitors := make([]*Monitor, 0, count)
		for i := 0; i < count; i++ {
			offset := unsafe.Sizeof(unsafe.Pointer(*cMonitors)) * uintptr(i)
			monitor := *(**Monitor)(unsafe.Pointer(uintptr(unsafe.Pointer(cMonitors)) + offset))
			monitors = append(monitors, monitor)
		}
		return monitors
	}
	return nil
}

// GetPrimaryMonitor returns the primary monitor.
//
// This function returns the primary monitor. This is usually the monitor where
// elements like the task bar or global menu bar are located.
//
// The primary monitor, or nil if no monitors were found or if an error
// occurred.
//
// Possible errors include NotInitialized.
//
// This function must only be called from the main thread.
func (c *Context) GetPrimaryMonitor() *Monitor {
	cMonitor := C.glfwGetPrimaryMonitor()
	return (*Monitor)(cMonitor)
}

// GetPos returns the position of the monitor's viewport on the virtual screen.
//
// This function returns the position, in screen coordinates, of the upper-left
// corner of the specific monitor.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (monitor *Monitor) GetPos() (x, y int) {
	var cX, cY C.int
	C.glfwGetMonitorPos((*C.GLFWmonitor)(monitor), &cX, &cY)
	x, y = int(cX), int(cY)
	return
}

// GetPhysicalSize returns the physical size of the monitor.
//
// This function returns the size, in millimeters, of the display area of the
// specified monitor.
//
// Some systems do not provide accurate monitor size information, either because
// the monitor EDID
// (https://en.wikipedia.org/wiki/Extended_display_identification_data) data is
// incorrect or because the driver does not report it accurately.
//
// Possible errors include NotInitialized.
//
// On Windows, calculates the returned physical size from the current resolution
// and system DPI instead of querying the monitor EDID data.
//
// This function must only be called from the main thread.
func (monitor *Monitor) GetPhysicalSize() (widthMM, heightMM int) {
	var cWidth, cHeight C.int
	C.glfwGetMonitorPhysicalSize((*C.GLFWmonitor)(monitor), &cWidth, &cHeight)
	widthMM, heightMM = int(cWidth), int(cHeight)
	return
}

// GetName returns the name of the specified monitor.
//
// This function returns a human-readable name, encoded as UTF-8, of the
// specified monitor. The name typically reflects the make and model of the
// monitor and is not guaranteed to be unique among the connected monitors.
//
// Possible errors include NotInitialized.
//
// This function must only be called from the main thread.
func (monitor *Monitor) GetName() string {
	return C.GoString(C.glfwGetMonitorName((*C.GLFWmonitor)(monitor)))
}

// SetMonitorCallback sets the monitor configuration callback.
//
// This function sets the monitor configuration callback, or removes the
// currently set callback. This is called when a monitor is connected to or
// disconnected from the system.
//
// callback is the new callback, or nil to remove the currently set callback.
//
// Returns the previously set callback, or nil if no callback was set or the
// library had not been initialized.
//
// Possible errors include NotInitialized.
//
// This function must only be called from the main thread.
func (c *Context) SetMonitorCallback(callback MonitorCallback) MonitorCallback {
	previousCallback := c.monitorCallback
	c.monitorCallback = callback
	if callback != nil {
		C.goSetMonitorCallback()
	} else {
		C.goRemoveMonitorCallback()
	}
	return previousCallback
}

//export _monitorCallback
func _monitorCallback(cMonitor *C.GLFWmonitor, cEvent C.int) {
	if currentContext != nil && currentContext.errorCallback != nil {
		monitor := (*Monitor)(cMonitor)
		event := int(cEvent)
		currentContext.monitorCallback(monitor, event)
	}
}

// GetVideoModes returns an array of all video modes supported by monitor, or
// nil if an error occurred. The returned array is sorted in ascending order,
// first by color bit depth (the sum of all channel depths) and then by
// resolution area (the product of width and height).
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (monitor *Monitor) GetVideoModes() []*VideoMode {
	var cCount C.int
	cModes := C.glfwGetVideoModes((*C.GLFWmonitor)(monitor), &cCount)
	if unsafe.Pointer(cModes) != C.NULL {
		count := int(cCount)
		videoModes := make([]*VideoMode, 0, count)
		for i := 0; i < count; i++ {
			offset := unsafe.Sizeof(*cModes) * uintptr(i)
			cMode := (*C.GLFWvidmode)(unsafe.Pointer(uintptr(unsafe.Pointer(cModes)) + offset))
			videoModes = append(videoModes, &VideoMode{
				Width:       int(cMode.width),
				Height:      int(cMode.height),
				RedBits:     int(cMode.redBits),
				GreenBits:   int(cMode.greenBits),
				BlueBits:    int(cMode.blueBits),
				RefreshRate: int(cMode.refreshRate),
			})
		}
		return videoModes
	}
	return nil
}

// GetVideoMode returns the current video mode of monitor, or nil if an error
// occurred. If you have created a full screen window for monitor, the return
// value will depend on whether that window is iconified.
//
// Returns nil if an error occurred.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (monitor *Monitor) GetVideoMode() *VideoMode {
	cMode := C.glfwGetVideoMode((*C.GLFWmonitor)(monitor))
	if unsafe.Pointer(cMode) != C.NULL {
		return &VideoMode{
			Width:       int(cMode.width),
			Height:      int(cMode.height),
			RedBits:     int(cMode.redBits),
			GreenBits:   int(cMode.greenBits),
			BlueBits:    int(cMode.blueBits),
			RefreshRate: int(cMode.refreshRate),
		}
	}
	return nil
}

// SetGamma generates a 256-element gamma ramp from the specified exponent and
// then calls Monitor.SetGammaRamp() with it. The value must be a finite number
// greater than zero.
//
// Possible errors include NotInitialized, InvalidValue and PlatformError.
//
// This function must only be called from the main thread.
func (monitor *Monitor) SetGamma(gamma float32) {
	C.glfwSetGamma((*C.GLFWmonitor)(monitor), C.float(gamma))
}

// GetGammaRamp returns the current gamma ramp for monitor, or nil if an error
// occurred.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (monitor *Monitor) GetGammaRamp() *GammaRamp {
	cRamp := C.glfwGetGammaRamp((*C.GLFWmonitor)(monitor))
	if unsafe.Pointer(cRamp) != C.NULL {
		size := int(cRamp.size)
		ramp := &GammaRamp{
			Red:   make([]uint16, 0, size),
			Green: make([]uint16, 0, size),
			Blue:  make([]uint16, 0, size),
		}
		for i := 0; i < size; i++ {
			offset := unsafe.Sizeof(*cRamp.red) * uintptr(i)
			cRed := (*C.ushort)(unsafe.Pointer(uintptr(unsafe.Pointer(cRamp.red)) + offset))
			cGreen := (*C.ushort)(unsafe.Pointer(uintptr(unsafe.Pointer(cRamp.red)) + offset))
			cBlue := (*C.ushort)(unsafe.Pointer(uintptr(unsafe.Pointer(cRamp.red)) + offset))
			ramp.Red = append(ramp.Red, uint16(*cRed))
			ramp.Green = append(ramp.Green, uint16(*cGreen))
			ramp.Blue = append(ramp.Blue, uint16(*cBlue))
		}
		return ramp
	}
	return nil
}

// SetGammaRamp sets the current gamma ramp for monitor. The original gamma ramp
// for monitor is saved by GLFW the first time this function is called and is
// restored by Context.Terminate().
//
// Possible errors include NotInitialized and PlatformError.
//
// Gamma ramp sizes other than 256 are not supported by all platforms or
// graphics hardware.
//
// On Windows, the gamma ramp size must be 256.
//
// This function must only be called from the main thread.
func (monitor *Monitor) SetGammaRamp(ramp *GammaRamp) {
	size := len(ramp.Red)
	cRed := make([]C.ushort, 0, size)
	cGreen := make([]C.ushort, 0, size)
	cBlue := make([]C.ushort, 0, size)
	for i := 0; i < int(size); i++ {
		cRed = append(cRed, C.ushort(ramp.Red[i]))
		cGreen = append(cGreen, C.ushort(ramp.Green[i]))
		cBlue = append(cBlue, C.ushort(ramp.Blue[i]))
	}
	cRamp := C.GLFWgammaramp{
		red:   &cRed[0],
		green: &cGreen[0],
		blue:  &cBlue[0],
		size:  C.uint(size),
	}
	C.glfwSetGammaRamp((*C.GLFWmonitor)(monitor), &cRamp)
}

// DefaultWindowHints resets all window hints to their default values.
//
// Possible errors include NotInitialized.
//
// This function must only be called from the main thread.
func (c *Context) DefaultWindowHints() {
	C.glfwDefaultWindowHints()
}

// WindowHint sets the specified window hint to the desired value.
//
// This function sets hints for the next call to Context.CreateWindow(). The
// hints, once set, retain their values until changed by a call to
// Context.WindowHint() or Context.DefaultWindowHint(), or until the library is
// terminated.
//
// This function does not check whether the specified hint values are valid. If
// you set hints to invalid values this will instead be reported by the next
// call to Context.CreateWindow().
//
// Possible errors include NotInitialized and InvalidEnum.
//
// This function must only be called from the main thread.
func (c *Context) WindowHint(hint Hint, value HintValue) {
	C.glfwWindowHint(C.int(hint), C.int(value))
}

// WindowHintBool is a shorthand for setting window hints with boolean values.
func (c *Context) WindowHintBool(hint Hint, value bool) {
	if value {
		c.WindowHint(hint, HintValue(True))
	} else {
		c.WindowHint(hint, HintValue(False))
	}
}

// CreateWindow creates a window and its associated context.
//
// This function creates a window and its associated OpenGL or OpenGL ES
// context. Most of the options controlling how the window and its context
// should be created are specified with window hints
// (http://www.glfw.org/docs/latest/window_guide.html#window_hints).
//
// Successful creation does not change which context is current. Before you can
// use the newly created context, you need to make it current
// (http://www.glfw.org/docs/latest/context_guide.html#context_current). For
// information about the share parameter, see Context object sharing
// (http://www.glfw.org/docs/latest/context_guide.html#context_sharing).
//
// The created window, framebuffer and context may differ from what you
// requested, as not all parameters and hints are hard constraints
// (http://www.glfw.org/docs/latest/window_guide.html#window_hints_hard). This
// includes the size of the window, especially for full screen windows. To query
// the actual attributes of the created window, framebuffer and context, see
// Window.GetAttrib(), Window.GetSize() and Window.GetFramebufferSize().
//
// To create a full screen window, you need to specify the monitor the window
// will cover. If no monitor is specified, the window will be windowed mode.
// Unless you have a way for the user to choose a specific monitor, it is
// recommended that you pick the primary monitor. For more information on how
// to query connected monitors, see Retrieving monitors
// (http://www.glfw.org/docs/latest/monitor_guide.html#monitor_monitors).
//
// For full screen windows, the specified size becomes the resolution of the
// window's desired video mode. As long as a full screen window is not
// iconified, the supported video mode most closely matching the desired video
// mode is set for the specified monitor. For more information about full screen
// windows, including the creation of so called windowed full screen or
// borderless full screen windows, see "Windowed full screen" windows
// (http://www.glfw.org/docs/latest/window_guide.html#window_windowed_full_screen).
//
// Once you have created the window, you can switch it between windowed and full
// screen mode with Window.SetMonitor(). If the window has an OpenGL or OpenGL
// ES context, it will be unaffected.
//
// By default, newly created windows use the placement recommended by the window
// system. To create the window at a specific position, make it initially
// invisible using the Visible window hint, set its position
// (http://www.glfw.org/docs/latest/window_guide.html#window_pos) and then show
// (http://www.glfw.org/docs/latest/window_guide.html#window_hide) it.
//
// As long as at least one full screen window is not iconified, the screensaver
// is prohibited from starting.
//
// Window systems put limits on window sizes. Very large or very small window
// dimensions may be overridden by the window system on creation. Check the
// actual size (http://www.glfw.org/docs/latest/window_guide.html#window_size)
// after creation.
//
// The swap interval
// (http://www.glfw.org/docs/latest/window_guide.html#buffer_swap) is not set
// during window creation and the initial value may vary depending on driver
// settings and defaults.
//
// Returns the handle of the created window, or nil if an error occurred.
//
// Possible errors include NotInitialized, InvalidEnum, InvalidValue,
// APIUnavailable, VersionUnavailable, FormatUnavailable and PlatformError.
//
// On Windows, window creation will fail if the Microsoft GDI software OpenGL
// implementation is the only one available.
//
// On Windows, if the executable has an icon resource named GLFW_ICON, it will
// be set as the initial icon for the window. If no such icon is present, the
// IDI_WINLOGO icon will be used instead. To set a different icon, see
// Window.SetIcon().
//
// On Windows, the context to share resources with must not be current on any
// other thread.
//
// On OS X, the GLFW window has no icon, as it is not a document window, but the
// dock icon will be the same as the application bundle's icon. For more
// information on bundles, see the Bundle Programming Guide
// (https://developer.apple.com/library/mac/documentation/CoreFoundation/Conceptual/CFBundles/)
// in the Mac Developer Library.
//
// On OS X, the first time a window is created the menu bar is populated with
// common commands like Hide, Quit and About. The About entry opens a minimal
// about dialog with information from the application's bundle. The menu bar can
// be disabled with a compile-time option
// (http://www.glfw.org/docs/latest/compile_guide.html#compile_options_osx).
//
// On OS X 10.10 and later the window frame will not be rendered at full
// resolution on Retina displays unless the NSHighResolutionCapable key is
// enabled in the application bundle's Info.plist. For more information, see
// High Resolution Guidelines
// (https://developer.apple.com/library/mac/documentation/GraphicsAnimation/Conceptual/HighResolutionOSX/Explained/Explained.html)
// for OS X in the Mac Developer Library. The GLFW test and example programs use
// a custom Info.plist template for this, which can be found as
// CMake/MacOSXBundleInfo.plist.in in the source tree.
//
// On X11, some window managers will not respect the placement of initially
// hidden windows.
//
// On X11, due to the asynchronous nature of X11, it may take a moment for a
// window to reach its requested state. This means you may not be able to query
// the final size, position or other attributes directly after window creation.
//
// This function must not be called from a callback.
//
// This function must only be called from the main thread.
func (c *Context) CreateWindow(width, height int, title string, monitor *Monitor, share *Window) *Window {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cWindow := C.glfwCreateWindow(C.int(width), C.int(height), cTitle, (*C.GLFWmonitor)(monitor), (*C.GLFWwindow)(share))
	if unsafe.Pointer(cWindow) != C.NULL {
		return (*Window)(cWindow)
	}
	return nil
}

// Destroy destroys win and its context. On calling this function, no further
// callbacks will be called for that window.
//
// If the context of win is current on the main thread, it is detached before
// being destroyed.
//
// Possible errors include NotInitialized and PlatformError.
//
// The context of win must not be current on any other thread when this function
// is called.
//
// This function must not be called from a callback.
//
// This function must only be called from the main thread.
func (win *Window) Destroy() {
	C.glfwDestroyWindow((*C.GLFWwindow)(win))
}

// ShouldClose returns the value of the close flag of win.
//
// Possible errors include NotInitialized.
//
// This function may be called from any thread. Access is not synchronized.
func (win *Window) ShouldClose() bool {
	return int(C.glfwWindowShouldClose((*C.GLFWwindow)(win))) != 0
}

// SetShouldClose sets the value of the close flag of win. This can be used to
// override the user's attempt to close the window, or to signal that it should
// be closed.
//
// Possible errors include NotInitialized.
//
// This function may be called from any thread. Access is not synchronized.
func (win *Window) SetShouldClose(value bool) {
	var cValue C.int
	if value {
		cValue = 1
	}
	C.glfwSetWindowShouldClose((*C.GLFWwindow)(win), cValue)
}

// SetTitle sets the window title, encoded as UTF-8, of win.
//
// Possible errors include NotInitialized or PlatformError.
//
// On OS X, the window title will not be updated until the next time you process
// events.
//
// This function must only be called from the main thread.
func (win *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	C.glfwSetWindowTitle((*C.GLFWwindow)(win), cTitle)
}

// SetIcon sets the icon for win. If passed a slice of candidate images, those
// of or closest to the sizes desired by the system are selected. If no images
// are specified (nil or an empty slice), win reverts to its default icon.
//
// The desired image sizes varies depending on platform and system settings. The
// selected images will be rescaled as needed. Good sizes include 16x16, 32x32
// and 48x48.
//
// Possible errors include NotInitialized and PlatformError.
//
// On OS X, the GLFW window has no icon, as it is not a document window, so this
// function does nothing. The dock icon will be the same as the application
// bundle's icon. For more information on bundles, see the Bundle Programming
// Guide
// (https://developer.apple.com/library/mac/documentation/CoreFoundation/Conceptual/CFBundles/)
// in the Mac Developer Library.
//
// This function must only be called from the main thread.
func (win *Window) SetIcon(images []Image) {
	if images == nil || len(images) == 0 {
		C.glfwSetWindowIcon((*C.GLFWwindow)(win), 0, (*C.GLFWimage)(C.NULL))
		return
	}

	cImages := make([]C.GLFWimage, 0, len(images))
	for _, image := range images {
		cImage := C.GLFWimage{
			width:  C.int(image.Width),
			height: C.int(image.Height),
		}

		cPixels := make([]C.uchar, 0, len(image.Pixels))
		for _, pixel := range image.Pixels {
			cPixels = append(cPixels, C.uchar(pixel))
		}
		cImage.pixels = &cPixels[0]

		cImages = append(cImages, cImage)
	}

	C.glfwSetWindowIcon((*C.GLFWwindow)(win), C.int(len(images)), &cImages[0])
}

// GetPos retrieves the position, in screen coordinates, of the upper-left
// corner of the client area of win.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) GetPos() (x, y int) {
	var cX, cY C.int
	C.glfwGetWindowPos((*C.GLFWwindow)(win), &cX, &cY)
	x, y = int(cX), int(cY)
	return
}

// SetPos sets the position, in screen coordinates, of the upper-left corner of
// the client area of win, a windowed mode window. If the window is a full
// screen window, this function does nothing.
//
// Do not use this function to move an already visible window unless you have
// very good reasons for doing so, as it will confuse and annoy the user.
//
// The window manager may put limits on what positions are allowed. GLFW cannot
// and should not override these limits.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) SetPos(x, y int) {
	C.glfwSetWindowPos((*C.GLFWwindow)(win), C.int(x), C.int(y))
}

// GetSize retrieves the size, in screen coordinates, of the client area of win.
// If you wish to retrieve the size of the framebuffer of the window in pixels,
// see Window.GetFramebufferSize().
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) GetSize() (width, height int) {
	var cWidth, cHeight C.int
	C.glfwGetWindowSize((*C.GLFWwindow)(win), &cWidth, &cHeight)
	width, height = int(cWidth), int(cHeight)
	return
}

// SetSizeLimits sets the size limits of the client area of win. If win is full
// screen, the size limits only take effect once it is made windowed. If win is
// not resizable, this function does nothing.
//
// The size limits are applied immediately to a windowed mode window and may
// cause it to be resized.
//
// The maximum dimensions must be greater than or equal to the minimum
// dimensions and all must be greater than or equal to zero.
//
// Use DontCare to disable the size limit for a dimension.
//
// Possible errors include NotInitialized, InvalidValue and PlatformError.
//
// If you set size limits and an aspect ratio that conflict, the results are
// undefined.
//
// This function must only be called from the main thread.
func (win *Window) SetSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {
	C.glfwSetWindowSizeLimits((*C.GLFWwindow)(win), C.int(minWidth), C.int(minHeight), C.int(maxWidth), C.int(maxHeight))
}

// SetAspectRatio sets the required aspect ratio of the client area of win. If
// win is full screen, the aspect ratio only takes effect once it is made
// windowed. If win is not resizable, this function does nothing.
//
// The aspect ratio is specified as a numerator and a denominator and both
// values must be greater than zero. For example, the common 16:9 aspect ratio
// is specified as 16 and 9, respectively.
//
// If the numerator and denominator is set to DontCare then the aspect ratio
// limit is disabled.
//
// The aspect ratio is applied immediately to a windowed mode window and may
// cause it to be resized.
//
// Possible errors include NotInitialized, InvalidValue and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) SetAspectRatio(numer, denom int) {
	C.glfwSetWindowAspectRatio((*C.GLFWwindow)(win), C.int(numer), C.int(denom))
}

// SetSize sets the size, in screen coordinates, of the client area of win.
//
// For full screen windows, this function updates the resolution of its desired
// video mode and switches to the video mode closest to it, without affecting
// the window's context. As the context is unaffected, the bit depths of the
// framebuffer remain unchaged.
//
// If you wish to update the refresh rate of the desired video mode in addition
// to its resolution, see Window.SetMonitor().
//
// The window manager may put limits on what sizes are allowed. GLFW cannot and
// should not override these limits.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) SetSize(width, height int) {
	C.glfwSetWindowSize((*C.GLFWwindow)(win), C.int(width), C.int(height))
}

// GetFramebufferSize retrives the size, in pixels, of the framebuffer of win.
// If you wish to retrieve toe size of win in screen coordinates, see
// Window.GetSize().
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) GetFramebufferSize() (width, height int) {
	var cWidth, cHeight C.int
	C.glfwGetFramebufferSize((*C.GLFWwindow)(win), &cWidth, &cHeight)
	width, height = int(cWidth), int(cHeight)
	return
}

// GetFrameSize retrieves the size, in screen coordinates, of each edge of the
// frame of win. This size includes the title bar, if the window has one. The
// size of the frame may vary depending on the window-related hints
// (http://www.glfw.org/docs/latest/window_guide.html#window_hints_wnd) used to
// create it.
//
// Because this function retrieves the size of each window frame edge and not
// the offset along a particular coordinate axis, the retrieved values will
// always be zero or positive.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) GetFrameSize() (left, top, right, bottom int) {
	var cLeft, cTop, cRight, cBottom C.int
	C.glfwGetWindowFrameSize((*C.GLFWwindow)(win), &cLeft, &cTop, &cRight, &cBottom)
	left, top, right, bottom = int(cLeft), int(cTop), int(cRight), int(cBottom)
	return
}

// Iconify iconifies (minimizes) win if it was previously restored. If win is
// already iconified, this function does nothing.
//
// If win is a full screen window, the original monitor resolution is restored
// until the window is restored.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) Iconify() {
	C.glfwIconifyWindow((*C.GLFWwindow)(win))
}

// Restore restores win if it was previously iconified (minimized) or maximized.
// If win is already restored, this function does nothing.
//
// If win is a full screen window, the resolution chosen for win is restored
// on the selected monitor.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) Restore() {
	C.glfwRestoreWindow((*C.GLFWwindow)(win))
}

// Maximize maximizes win if it was previously not maximized. If win is already
// maximized, this function does nothing.
//
// If win is a full screen window, this function does nothing.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) Maximize() {
	C.glfwMaximizeWindow((*C.GLFWwindow)(win))
}

// Show makes win visible if it was previously hidden. If win is already visible
// or is in full screen mode, this function does nothing.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) Show() {
	C.glfwShowWindow((*C.GLFWwindow)(win))
}

// Hide hides win if it was previously visible. If win is already hidden of is
// in full screen mode, this function does nothing.
//
// Possible errors include NotInitialized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) Hide() {
	C.glfwHideWindow((*C.GLFWwindow)(win))
}

// Focus brings win to front and sets input focus. win should already be visible
// and not iconified.
//
// Be default, both windowed and full screen mode windows are focused when
// initially created. Set the Focused hint to disable this behavior.
//
// Do not use this function to steal focus from other applications unless you
// are certain that is what the user wants. Focus stealing can be extremely
// disruptive.
//
// Possible errors include NotInitailized and PlatformError.
//
// This function must only be called from the main thread.
func (win *Window) Focus() {
	C.glfwFocusWindow((*C.GLFWwindow)(win))
}
