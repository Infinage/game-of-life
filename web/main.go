package main

import (
	"syscall/js"
	"unsafe"

	"github.com/infinage/game-of-life/pkg/gol"
)

// Input is fed as [x1, y1, x2, y2, ...]
// Output as [Cell{X1, Y1}, Cell{X2, Y2}, ...]
func fromPairs(arr []int32) gol.Grid {
	grid := make(gol.Grid)
	for i := 0; i < len(arr); i += 2 {
		cell := gol.Cell{X: int(arr[i]), Y: int(arr[i + 1])}
		grid[cell] = nil
	}
	return grid
}

// Input is fed as [Cell{X1, Y1}, Cell{X2, Y2}, ...]
// Output as [x1, y1, x2, y2, ...]
func toPairs(grid gol.Grid) []int32 {
	var res []int32	
	for cell := range grid {
		res = append(res, int32(cell.X), int32(cell.Y))
	}
	return res
}

func main() {
	// Object is passed in as a int32array [x1, y1, x2, y2, ..]
	js.Global().Set("GOLNext", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return map[string]any{"success": false, "data": nil, "error": "Missing input for Grid"}
		}

		// Capture JS Int32array from function args
		jsArr := args[0]
		length := jsArr.Get("length").Int()
		if length % 2 == 1 {
			return map[string]any{"success": false, "data": nil, "error": "Expected even number of elements in Grid"}
		}

		// On zero length input, return empty result
		if length == 0 {
			return map[string]any{"success": true, "data": js.Global().Get("Uint8Array").New(0), "error": ""}
		}

		// Access the underlying Uint8Array representation in JS
		jsUint8Buffer := js.Global().Get("Uint8Array").New(
			jsArr.Get("buffer"), jsArr.Get("byteOffset"), 
			length * 4, // int32 = 4 bytes
		)

		// Allocate a raw byte slice in Go and copy memory
		goByteSlice := make([]byte, length * 4)
		js.CopyBytesToGo(goByteSlice, jsUint8Buffer)

		// Reinterpret the byte slice as an int32 slice
		goSlice := unsafe.Slice((*int32)(unsafe.Pointer(&goByteSlice[0])), length)

		// Parse & perform one game of conway iteration
		grid := fromPairs(goSlice)
		grid.Next()

		// Write back as []int32
		if goSlice = toPairs(grid); len(goSlice) == 0 {
			return map[string]any{"success": true, "data": js.Global().Get("Uint8Array").New(0), "error": ""}
		}

		// Type cast back to uint8 buffer
		goUint8Slice := unsafe.Slice((*uint8)(unsafe.Pointer(&goSlice[0])), len(goSlice) * 4)
		jsUint8Buffer = js.Global().Get("Uint8Array").New(len(goUint8Slice))
		js.CopyBytesToJS(jsUint8Buffer, goUint8Slice)

		return map[string]any{"success": true, "data": jsUint8Buffer, "error": ""}
	}))

	// Keep alive
	select {}
}
