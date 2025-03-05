# Commitsmith

Commitsmith is a command-line tool designed to streamline and simplify the process of writing Git commit messages following conventional commit standards. It provides a user-friendly TUI (Terminal User Interface) built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and styles powered by [Lip Gloss](https://github.com/charmbracelet/lipgloss). Commitsmith helps developers select commit types, specify scopes, handle breaking changes, and craft detailed descriptions for their commits efficiently.

---

## Features
- **Guided Commit Message Creation**: An interactive form for writing Git commit messages.
- **Conventional Commit Formatting**: Generates commit messages using standard templates.
- **Integrated Git Support**: Automatically stages and commits changes.
- **Dry-run Mode**: Preview your commit message without actually making any changes.
- **Support for Breaking Changes**: Easily mark commits as introducing breaking changes.
- **Customizable Body and Scope**: Add additional details to your commit messages.

---

## Installation
To use Commitsmith, either clone this repository and build the binary for the latest release, or download a pre-built 
binary from the [Releases](https://github.com/metaldrummer610/commitsmith/releases):

1. Clone the repository:
    ```shell script
    git clone https://github.com/metaldrummer610/commitsmith.git
    ```

2. Navigate into the directory:
    ```shell script
    cd commitsmith
    ```

3. Build the project:
    ```shell script
    make
    ```

Now, you can use the `commitsmith` binary to run the tool. If typing out `commitsmith` is a bit much, feel free to alias
the command to something like `cs`. You'll also want to copy the binary to somewhere on your `$PATH` so it can be executed
from anywhere.

---

## Usage

Run Commitsmith in your terminal to generate and create commit messages for Git repositories. Here are some examples of how you can use it:

### Output the Commit Message Without Applying It (Dry Run)
```shell script
commitsmith --dry-run -m "Initial setup for commitsmith"
```

### Generate and Commit the Message
```shell script
commitsmith -m "Fix a bug in authentication module"
```

### Run in Interactive Mode Without Supplying a Message
If the `-m` flag is not specified, Commitsmith will open the interactive commit editor:
```shell script
commitsmith
```

### Display Help
```shell script
commitsmith -h
```

---

## Examples

### Dry Run Example
If you wish to preview the commit message without making a commit, you can use the `--dry-run` flag. For example:
```shell script
commitsmith --dry-run
```

This will display an output similar to:

```
Commit message:
feat(auth): Add OAuth2 support

This commit introduces support for OAuth2 authentication for our backend API.
```

### Interactive Commit Creation
Running `commitsmith` without the `-m` flag opens an interactive TUI. You'll navigate through prompts to select the commit type, add a scope, mark breaking changes if needed, and input detailed descriptions.

Example flow:
1. Select commit type (`feat`, `fix`, etc.).
2. (Optional) Add a scope (e.g., `auth` for authentication-related changes).
3. Specify if it's a breaking change.
4. Add a short description and optionally a body.

---

## Contributing

We welcome contributions to Commitsmith! Here's how you can contribute:

### Prerequisites
- **Go Toolchain**: Ensure you have Go 1.24 or newer installed. You can download it [here](https://golang.org/dl/).
- **Git**: Required for cloning the repository and using the tool itself.

### Steps to Contribute
1. **Fork the Repository**:
   Navigate to the [Commitsmith repository](https://github.com/metaldrummer610/commitsmith) and click "Fork."

2. **Clone Your Fork**:
   Clone your forked repository to your local machine:
    ```shell script
    git clone https://github.com/yourusername/commitsmith.git
    ```

3. **Create a New Branch**:
   Create a new branch to work on your feature or bug fix:
    ```shell script
    git checkout -b feature/your-feature-name
    ```

4. **Make Your Changes**:
   Modify the code and include tests for any new functionality in the `*_test.go` files.

5. **Run Tests**:
   Before submitting your changes, ensure all tests pass:
    ```shell script
    make test
    ```

6. **Commit Your Changes**:
   Write a descriptive commit message using commitsmith:
    ```shell script
    commitsmith
    ```

7. **Push Your Changes**:
   Push your branch to your forked repository:
    ```shell script
    git push origin feature/your-feature-name
    ```

8. **Submit a Pull Request**:
   Open a pull request from your branch into the `main` branch of the original repository.

---

## Development

### Running the Project in Development
You can run the project directly from the source for development purposes:
```shell script
make
./commitsmith
```

### Testing
Run test cases with the following command:
```shell script
make test
```

Make sure to write new test cases for additional functionality or bug fixes in the appropriate `*_test.go` files.

---

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.

---

We hope you enjoy using Commitsmith! If you encounter any issues, feel free to open an issue on our GitHub repository. Happy coding! 🚀