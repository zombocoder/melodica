# Melodica

![Build](https://github.com/zombocoder/melodica/actions/workflows/build.yml/badge.svg)
![Release](https://github.com/zombocoder/melodica/actions/workflows/release.yml/badge.svg)

Melodica is a console-based audio player built with Go. It supports playback of MP3 files from a playlist and includes basic controls like play, stop, next, previous, volume control, and more. The application is designed to work across multiple operating systems, including Linux, macOS, and Windows.

## Screenshot

<img width="1505" alt="image" src="https://github.com/user-attachments/assets/d6e6e43b-73ad-4cab-a998-d6e8855a3f4a">

## Features

- **Cross-platform**: Supports Linux, macOS (including Apple Silicon), and Windows.
- **Console Interface**: Simple and lightweight user interface using the terminal.
- **Controls**:
  - `Enter`: Play selected track
  - `s`: Stop playback
  - `n`: Next track
  - `p`: Previous track
  - `x`: Pause/Resume playback
  - `>`: Volume up
  - `<`: Volume down
  - `Esc`: Quit application

## Installation

To install Melodica, ensure you have Go installed (version 1.20 or later). You can install the latest version of Melodica directly from the GitHub repository using the following command:

```bash
go install github.com/zombocoder/melodica/cmd/melodica@latest
```

This command will download and install Melodica into your `$GOPATH/bin` directory.

## Usage

1. **Download a Sample Playlist**: To get started quickly, download a sample `playlist.txt` file from GitHub:

   ```bash
   curl -o playlist.txt https://raw.githubusercontent.com/zombocoder/melodica/main/playlist.txt
   ```

2. **Run Melodica**:

   ```bash
   melodica playlist.txt
   ```

3. **Controls**: Use the following key bindings within the application:
   - `Enter`: Play the selected track from the playlist
   - `s`: Stop the current track
   - `n`: Skip to the next track
   - `p`: Go back to the previous track
   - `x`: Pause or resume playback
   - `>`: Increase volume
   - `<`: Decrease volume
   - `Esc`: Quit the application

## Makefile Commands

The `Makefile` provides a convenient way to build and run the application, as well as clean up generated files.

- **Build**: Compile the application for your OS.

  ```bash
  make build
  ```

- **Run**: Run the application with a specified playlist file.

  ```bash
  make run PLAYLIST=playlist.txt
  ```

- **Test**: Run all tests.

  ```bash
  make test
  ```

- **Clean**: Remove the compiled binary and other generated files.
  ```bash
  make clean
  ```

## Disclaimer

The audio files used in Melodica are sourced from [Lofi Girl's](https://lofigirl.com/) website and are subject to their [licensing guidelines](https://form.lofigirl.com/CommercialLicense).

Please ensure compliance with Lofi Girl's licensing terms if you plan to use Melodica in a commercial context.

---

## Development

To contribute or run the application locally:

1. Clone the repository:

   ```bash
   git clone https://github.com/zombocoder/melodica.git
   cd melodica
   ```

2. Build the application using `Makefile`:

   ```bash
   make build
   ```

3. Run the application with your playlist:
   ```bash
   go run cmd/melodica/main.go playlist.txt
   ```

## Building for Other Platforms

This project is configured with a GitHub Actions workflow that builds and releases artifacts for Linux, macOS (including Apple Silicon), and Windows on each tagged release.

## License

This project is licensed under the MIT License.
