# Rate Limit Checker Implementation

This project implements a rate limiting mechanism to control and monitor request rates in a Go application.

## Project Structure

```tree
.
├── cmd/          # Main application entry point
├── pkg/          # Package containing core implementation
└── tests/        # Test files
```

## Quick Start

### Running the Application

You can run the application using either Make or Go commands.

Using Make:

```bash
make gorun
```

Using Go directly:

```bash
go run ./cmd/app
```

### Running Tests

Run the test suite to verify the implementation.

Using Make:

```bash
make gotests
```

Using Go directly:

```bash
go test ./tests
```

### Building the Application

Build a binary of the application.

Using Make:

```bash
make gobuild
```

Using Go directly:

```bash
go build ./cmd/app
```

The compiled binary will be created in the `cmd/` directory.

## License

This project is licensed under the MIT License:

```text
MIT License

Copyright (c) 2025 fxmariojevta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Contributing

Contributions are welcome! Here's how you can help improve this project:

### Setting Up Development Environment

1. Fork the repository

2. Clone your fork:

   ```bash
   git clone https://github.com/YOUR-USERNAME/HitRateLimitCheckImplementation.git
   ```

3. Create a new branch for your feature:

   ```bash
   git checkout -b feature/your-feature-name
   ```

### Making Changes

1. Make your changes in the code

2. Run the tests to ensure everything works:

   ```bash
   make gotests
   ```

3. Commit your changes:

   ```bash
   git commit -m "Add: brief description of your changes"
   ```

4. Push to your fork:

   ```bash
   git push origin feature/your-feature-name
   ```

### Submitting Changes

1. Open a Pull Request (PR) from your fork to this repository
2. Ensure your PR description clearly describes the changes and their benefits
3. Link any related issues in your PR description

### Code Guidelines

- Write clear, documented code
- Follow Go best practices and conventions
- Include tests for new features
- Keep commits focused and atomic
- Update documentation as needed

### Bug Reports and Feature Requests

- Use the GitHub Issues tab to report bugs or request features
- Provide clear descriptions and steps to reproduce issues
- Include relevant code samples or error messages
