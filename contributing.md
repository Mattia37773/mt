# Contributing Guidelines

Thank you for considering a contribution to the MT project. We appreciate your interest in making this tool more effective and robust for the developer community.

By participating in this project, you agree to abide by the terms of the MIT License.

## How to Contribute

We welcome contributions in the form of bug reports, feature requests, and code improvements.

### 1. Reporting Issues

Before opening a new issue, please check the existing issues to ensure it has not already been reported. When creating a new report, include:

- A clear, concise title.
- A detailed description of the expected behavior versus the actual behavior.
- Steps to reproduce the issue.
- Your environment details (OS, Docker version, etc.).

### 2. Suggesting Features

If you have a feature request, open an issue labeled as an enhancement. Describe the problem you are trying to solve and why this feature would be a valuable addition to the CLI.

### 3. Pull Requests

We follow a standard Git workflow:

1. **Fork the repository** to your own account.
2. **Create a feature branch** from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Make your changes**. Please ensure your code follows the project's existing style and includes relevant documentation or tests if applicable.
4. **Commit your changes**. Use clear, imperative commit messages (e.g., "Add support for MariaDB" instead of "Fixed stuff").
5. **Push your branch**:
   ```bash
   git push origin feature/your-feature-name
   ```
6. **Open a Pull Request** against the `main` branch. Provide a summary of the changes and reference any related issues.

## Development Setup

To work on the project locally:

1. Clone your fork:
   ```bash
   git clone https://github.com/mattia37773/mt.git
   cd mt
   ```
2. Build the project:
   ```bash
   go build -o mt .
   ```
3. Verify your changes by running the binary locally.

## Code Style

- Use idiomatic Go code.
- Ensure all exported functions and types are documented.
- Run `go mod tidy` before committing to ensure dependencies are managed correctly.

Thank you for your help in improving the MT CLI.
