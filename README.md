### The README

Here is a clean, developer-focused `README.md` that highlights the raw memory performance and the TinyGo/Wasm architecture.

```markdown
# Conway's Game of Life (Go + WebAssembly)

A blazingly fast, infinite-grid implementation of Conway's Game of Life. Built with Go, compiled to WebAssembly via TinyGo, and rendered on a zero-dependency HTML5 Canvas.

🌍 **Live Demo:** [game-of-life.infinage.space](https://game-of-life.infinage.space)

## Architecture

Instead of relying on heavy DOM manipulation or expensive JSON serialization across the JS/Wasm bridge, this engine uses contiguous memory blocks. The Go backend tracks a sparse map of live cells and passes a flat array of 32-bit integers across the WebAssembly boundary using raw byte copying (`js.CopyBytesToJS`). The vanilla JavaScript frontend directly translates these logical coordinates into screen-space Canvas pixels.

### Features
* **Infinite Universe:** Coordinate bounds are dictated by standard `int32` limits, allowing patterns like Gliders to travel practically forever without crashing.
* **Tiny Footprint:** Compiled via TinyGo for a drastically reduced `.wasm` binary size compared to the standard Go compiler.
* **Responsive Canvas:** Fluid 60FPS panning and zooming using standard mouse and touch inputs.
* **Retro UI:** Pure Tailwind CSS and raw SVGs—no heavy frontend frameworks.

## Tech Stack
* **Engine:** Go 1.26
* **Compiler:** TinyGo
* **Frontend:** Vanilla JavaScript, HTML5 Canvas
* **Styling:** Tailwind CSS (via CDN)
* **Deployment:** Docker (BusyBox)

## Building from Source

### Prerequisites
* [Go](https://go.dev/doc/install) (1.26+)
* [TinyGo](https://tinygo.org/getting-started/install/)

### 1. Compile to WebAssembly
To compile the Go engine into a Wasm binary, you must use the `wasm_exec.js` provided specifically by TinyGo (not the standard Go one).

```bash
# Copy the TinyGo bridge script to the web directory
cp /opt/tinygo/targets/wasm_exec.js web/wasm_exec.js

# Build the Wasm binary
tinygo build -o web/gol.wasm -target wasm -no-debug web/main.go

```

### 2. Run Locally (Docker)

The included `Dockerfile` uses BusyBox to spin up a tiny, static HTTP server that correctly maps the `application/wasm` MIME type.

```bash
# Build the image
docker build -t gol-wasm .

# Run the container on port 8080
docker run -p 8080:80 gol-wasm

```

Navigate to `http://localhost:8080` in your browser.

## The Rules of Life

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, live or dead. Every cell interacts with its eight neighbors:

1. **Underpopulation:** Any live cell with fewer than two live neighbors dies.
2. **Survival:** Any live cell with two or three live neighbors lives on to the next generation.
3. **Overpopulation:** Any live cell with more than three live neighbors dies.
4. **Reproduction:** Any dead cell with exactly three live neighbors becomes a live cell.

## License

MIT
