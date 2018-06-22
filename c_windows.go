// Copyright (c) 2018 Beta Kuang
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build windows

package glfw

/*
#include "glfw/src/win32_init.c"
#include "glfw/src/win32_joystick.c"
#include "glfw/src/win32_monitor.c"
#include "glfw/src/win32_time.c"
#include "glfw/src/win32_tls.c"
#include "glfw/src/win32_window.c"
#include "glfw/src/wgl_context.c"
#include "glfw/src/egl_context.c"
*/
import "C"
