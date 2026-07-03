# ASCII Art Web Stylize

## Description

ASCII Art Web Stylize is a web-based application that converts user input text into ASCII art using different banner styles. It provides a graphical user interface (GUI) built with HTML and CSS, and a backend server written in Go.

The application allows users to:
- Enter any text they want to convert
- Choose from three different ASCII art banner styles: **standard**, **shadow**, and **thinkertoy**
- Generate and view the ASCII art result directly in the browser
- Experience a styled, visually appealing web interface with custom CSS theming

This project extends the original [ascii-art](https://github.com/01-edu/public) CLI project by making it accessible through a stylized web interface.

---

## Authors

- **Chisom Nnamani** (@cnnamani)
- **Samuel Momoh** (@GRITTY)
- **Stephen Chima** (@teeschima)

---

## Usage

### Prerequisites

Before running the application, ensure you have:
- **Go 1.22 or higher** installed on your system
- The project files downloaded to your local machine

### Installation

1. **Clone or download the project:**
   ```bash
   git clone https://github.com/cnnamani/ascii-art-web-stylize.git
   cd ascii-art-web-stylize
   ```

2. **Verify the folder structure:**
   ```
   ascii-art-web-stylize/
   ├── ascii_art.go
   ├── ascii_art_test.go
   ├── design/
   │   └── Inspo-Image.png
   ├── go.mod
   ├── main.go
   ├── static/
   │   ├── assets/
   │   │   ├── banners/
   │   │   │   ├── shadow.txt
   │   │   │   ├── standard.txt
   │   │   │   └── thinkertoy.txt
   │   │   └── images/
   │   │       ├── aag_logo.png
   │   │       └── arrow_down.png
   │   └── styles/
   │       └── index.css
   ├── templates/
   │   └── index.html
   └── README.md
   ```

### Running the Server

1. **Navigate to the project directory:**
   ```bash
   cd ascii-art-web-stylize
   ```

2. **Start the server:**
   ```bash
   go run .
   ```

3. **Expected output:**
   ```
   Server is running on port: http://localhost:5500
   ```

4. **Open your browser:**
   - Navigate to: `http://localhost:5500`
   - The home page with the ASCII art generator form will load

### Using the Application

1. **Enter text:** Type any text in the textarea field (do not leave it empty)
2. **Select a banner:** Click one of the three radio buttons to choose a style:
   - Standard
   - Shadow
   - Thinkertoy
3. **Generate:** Click the "GENERATE ART" button
4. **View result:** The ASCII art will be displayed below in a scrollable `<pre>` block

> **Note:** Radio buttons auto-submit the form on change, so you can switch banner styles and see results instantly without clicking the generate button.

### Example

**Input:**
- Text: "Hedlo"
- Banner: "standard"

**Output:**
```
 _    _          _ _       
| |  | |        | | |      
| |__| | ___  __| | | ___  
|  __  |/ _ \/ _` | |/ _ \ 
| |  | |  __/ (_| | | (_) |
|_|  |_|\___|\__,_|_|\___/ 
```

---

## Implementation Details

### Architecture

The application follows a simple monolithic architecture with all logic in a single package:

```
main.go (Router + Handler)
    ├── Serves templates/index.html (GET /)
    ├── Processes form submission (POST /)
    └── ascii_art.go (Business Logic)
```

### How It Works

#### 1. **Frontend (HTML Template + CSS)**
- Located in: `templates/index.html`
- Styled with: `static/styles/index.css`
- Contains a form with:
  - Banner selection radio buttons (auto-submit on change)
  - Text input field (`<textarea>`)
  - Submit button
  - Result display area (`<pre>`) in a scrollable container
- Uses a custom color scheme with orange background, red accents, and dark textarea

#### 2. **Router (main.go)**
- Registers a single route:
  - `GET /` → Serves the home page with the form
  - `POST /` → Processes form submission and generates ASCII art
- Serves static files from `./static` directory
- Starts the HTTP server on port 5500

#### 3. **HomePage Handler**
- Handles both GET and POST requests to `/`
- **GET requests:** Renders the empty form with default banner (standard.txt)
- **POST requests:**
  - Parses form data (text input and banner type)
  - Calls `GenerateAsciiArtText()` to generate the result
  - Detects emoji input and displays a custom error message
  - Renders the template with the ASCII art result

#### 4. **Core Logic (ascii_art.go)**

**GenerateAsciiArtText Function**
```go
func GenerateAsciiArtText(text, bannerType string) (string, error)
```

**Parameters:**
- `text` (string): The input text to convert to ASCII art
- `bannerType` (string): The banner filename ("standard.txt", "shadow.txt", or "thinkertoy.txt")

**Returns:**
- `result` (string): The generated ASCII art
- `error`: Any error that occurred during processing

**Algorithm:**

1. **Normalize Input:** Replace Windows newlines (`\r\n`) with Unix newlines (`\n`)
2. **Read Banner File:** Load the appropriate `.txt` file from `static/assets/banners/`
3. **Parse Banner Data:** Split the file into lines for character mapping
4. **Process Input Text:** Split text by newlines for multi-line support
5. **Build ASCII Art:**
   - Loop through each line of input
   - For each character, calculate its ASCII index in the banner file
   - Each character has 8 lines of ASCII art (height)
   - Each character takes 9 positions width (including spacing)
   - Index formula: `(int(char) - 32) * 9 + 1 + current_height`
6. **Validate Characters:** Ensure only printable ASCII characters (32-126) are processed
7. **Return Result:** Return the complete ASCII art string

**Emoji Detection:**
- If a non-ASCII character (e.g., emoji) is detected, an ASCII art error message is returned:
  ```
   /\
   /  \
   / || \
   /  ||  \
   /   ..   \
   /__________\
     E R R O R

   NO EMOJI ALLOWED IN THE TEXT AREA
  ```
- The handler displays this error in red text in the result container

#### 5. **Banner Files**

Three `.txt` files contain ASCII art character definitions:
- `standard.txt` - Simple, clean ASCII art
- `shadow.txt` - Shadowed ASCII art style
- `thinkertoy.txt` - Decorative ASCII art style

Each file contains ASCII art for all printable ASCII characters (space to tilde), organized in lines.

#### 6. **Static Assets**

- **CSS:** `static/styles/index.css` - Custom styling with a cohesive color palette
- **Images:**
  - `aag_logo.png` - Application logo displayed in the header
  - `arrow_down.png` - Decorative arrow pointing to the output area

### Data Flow

```
User Input (HTML Form)
        ↓
POST /
        ↓
homePageHandler
        ├── Parse form data
        ├── Validate inputs
        └── Call GenerateAsciiArtText()
                ↓
        GenerateAsciiArtText Function
        ├── Read banner file
        ├── Parse banner data
        ├── Detect emoji/non-ASCII input
        ├── Process input text
        └── Build ASCII art
                ↓
        Return result to handler
        ↓
Execute template with result
        ↓
Send HTML response to browser
        ↓
Display ASCII art to user
```

### HTTP Status Codes

| Status Code | Meaning | When Used |
|-------------|---------|-----------|
| 200 OK | Success | Valid request processed successfully |
| 404 Not Found | Not found | Invalid path, missing template file, missing banner file |
| 405 Method Not Allowed | Method error | Request method other than GET or POST |
| 500 Internal Server Error | Server error | Template parsing failure, invalid banner file |

---

## Testing

The project includes unit tests in `ascii_art_test.go`.

### Test Categories

1. **Valid Text with Missing Banner** - Tests error handling when a non-existent banner file is specified
2. **Valid Text with Valid Banner** - Tests correct ASCII art generation with the standard banner

### Running Tests

**Run all tests:**
```bash
go test -v
```

**Run specific test:**
```bash
go test -run Test_With_Valid_Text_And_Valid_Banner -v
```

**Run tests with coverage:**
```bash
go test -cover
```

**Generate coverage report:**
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Run tests before committing (team workflow):**
```bash
go test -v
git add .
git commit -m "Your message"
git push
```

---

## Technologies Used

- **Language:** Go (Golang) 1.22.2
- **Web Framework:** Go standard library (`net/http`)
- **Templates:** Go `text/template` package
- **Testing:** Go `testing` package
- **Frontend:** HTML5, CSS3 with custom styling

---

## Folder Structure

```
ascii-art-web-stylize/
│
├── main.go                 # Application entry point, routes, HTTP handler
├── ascii_art.go            # Core ASCII art generation logic
├── ascii_art_test.go       # Unit tests
├── go.mod                  # Go module file
│
├── templates/              # HTML templates
│   └── index.html          # Main page form with result display
│
├── static/                 # Static assets served to the browser
│   ├── assets/
│   │   ├── banners/        # ASCII art banner definition files
│   │   │   ├── standard.txt
│   │   │   ├── shadow.txt
│   │   │   └── thinkertoy.txt
│   │   └── images/         # Site images and icons
│   │       ├── aag_logo.png
│   │       └── arrow_down.png
│   └── styles/
│       └── index.css       # Custom stylesheet
│
├── design/                 # Design assets
│   └── Inspo-Image.png     # Design inspiration reference
│
└── README.md               # This file
```

---

## Styling Details

The application features a custom-designed UI with:

- **Color Palette:** Orange (`#fca311`) background, dark blue (`#003049`) banner selector, red (`#c1121f`) accents, black textarea
- **Layout:** Centered content with flexbox header alignment
- **Banner Selection:** Hidden radio buttons with styled labels that highlight on selection
- **Result Display:** Scrollable output container with monospace font for ASCII art
- **Hover Effects:** Submit button color change on hover
- **Responsive Elements:** Logo aligned with title, arrow indicator pointing to output

---

## Good Practices Implemented

✅ **Error Handling**
- Proper HTTP status codes for all scenarios
- Meaningful error messages to users and developers
- Input validation before processing (emoji detection)
- Custom ASCII art for emoji error feedback

✅ **Code Organization**
- Clear separation of concerns between handler and business logic
- Static assets organized in dedicated directories
- Consistent naming conventions

✅ **Documentation**
- README with detailed implementation details
- Code is self-documenting with clear variable names

✅ **Testing**
- Unit tests for core ASCII art generation
- Valid input and error scenarios covered

✅ **Go Best Practices**
- Using standard library only (no external dependencies)
- Proper error handling with `if err != nil`
- Clean naming conventions
- Proper use of HTTP methods and status codes

---

## Troubleshooting

### Problem: Port 5500 already in use

**Solution:**
- Change the port in `main.go` from `:5500` to another port like `:8080`:
  ```go
  err := http.ListenAndServe(":8080", mux)
  ```
- Then access `http://localhost:8080`

### Problem: Template file not found error

**Ensure:**
- `templates/index.html` exists in the project root
- You're running the server from the project root directory
- Use `go run .` to run from the root

### Problem: Banner file not found error

**Ensure:**
- Banner files exist in `static/assets/banners/`
- You're running the server from the project root directory

### Problem: Static assets not loading (CSS, images)

**Ensure:**
- The `static/` directory is in the project root
- File paths in HTML templates match the actual directory structure

### Problem: Tests are failing

**Solution:**
1. Run tests with verbose output: `go test -v`
2. Check which test is failing
3. Read the error message carefully
4. Verify your code implements the requirements

---

## Contributing

When contributing to this project:

1. Run all tests before committing:
   ```bash
   go test -v
   ```

2. Follow the code style and structure
3. Add tests for any new features
4. Update this README if you add new functionality
5. Use meaningful commit messages

---

## License

This project is part of the 01-Edu curriculum.

---

## References

- [Go Documentation](https://golang.org/doc/)
- [Go HTTP Package](https://golang.org/pkg/net/http/)
- [Go Templates](https://golang.org/pkg/text/template/)
- [ASCII Art Example](http://patorjk.com/software/taag/)

---

## Contact & Support

For issues or questions:
- Check existing GitHub issues
- Create a new GitHub issue with detailed description
- Contact team members directly

---

**Last Updated:** July 2026
**Project Version:** 1.0
