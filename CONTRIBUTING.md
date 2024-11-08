# Contributing to Melodica

Thank you for your interest in contributing to Melodica! Your contributions help make this project better. Here’s a guide to help you get started.

## How to Contribute

### 1. Report Issues

If you encounter any bugs or have feature requests, please open an [issue](https://github.com/zombocoder/melodica/issues). Make sure to provide enough context and details for others to understand and potentially reproduce the issue.

### 2. Fork the Repository

To contribute code, start by forking the repository. This will create a copy of the repository under your GitHub account.

1. Click the **Fork** button at the top of this repository.
2. Clone the forked repository to your local machine:
   ```bash
   git clone https://github.com/zombocoder/melodica.git
   cd melodica
   ```

### 3. Create a Branch

Create a new branch for each feature or bug fix:
```bash
git checkout -b feature/your-feature-branch-name
```

### 4. Make Changes

Make your changes in the new branch. Ensure your code is clean and well-documented. Follow the project’s coding style, and keep changes focused on the scope of the feature or bug fix.

### 5. Run Tests

Make sure all tests pass by running:
```bash
make test
```
If you’re adding a new feature or fixing a bug, consider adding new tests to verify your changes.

### 6. Commit Changes

Commit your changes with a clear and concise message:
```bash
git add .
git commit -m "Add feature: brief description of your feature"
```

### 7. Push Changes

Push your branch to your forked repository:
```bash
git push origin feature/your-feature-branch-name
```

### 8. Open a Pull Request

Go to the original Melodica repository, and open a pull request from your branch. In your pull request, please provide:

- A clear title and description of the change
- Reference to any issues it addresses, if applicable (e.g., `Closes #issue-number`)
- Any additional information that could help reviewers understand your change

## Code Style

- Follow Go conventions for formatting and naming.
- Use comments to document complex or important sections of code.
- Write tests for new features and bug fixes where applicable.

## Development Workflow

1. **Build the Project**: Use `make build` to compile the application.
2. **Run the Application**: Use `make run PLAYLIST=playlist.txt` to run the app with a specified playlist.
3. **Run Tests**: Use `make test` to ensure that tests pass before submitting a pull request.
4. **Clean Up**: Use `make clean` to remove generated files when needed.

## Additional Notes

- **Be respectful**: Please follow the [Code of Conduct](CODE_OF_CONDUCT.md) and treat all contributors with respect.
- **Be patient**: Reviewers may need time to review your changes.

Thank you for contributing to Melodica and helping make it better for everyone! We appreciate your support and look forward to collaborating with you.
