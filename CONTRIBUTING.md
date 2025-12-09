# Contributing to Foggo

Thanks for contributing! 🎉

## Quick Start

1. Fork the repository
2. Create a branch (`git checkout -b feature/your-feature`)
3. Make your changes
4. Push and create a PR
5. **CI will run tests and linters automatically** - no need to run them locally!

## Pull Request Guidelines

- **Don't worry about local testing** - GitHub Actions CI will handle it
- Just make sure your code builds: `go build ./...`
- CI will check:
  - Tests (`go test`)
  - Formatting (`gofmt`, `goimports`)
  - Linting (`golangci-lint`)
  - Build verification

## Commit Format

Use simple prefixes:
- `feat:` New features
- `fix:` Bug fixes
- `docs:` Documentation
- `test:` Tests
- `chore:` Maintenance

Example: `feat: add custom template support`

## Need Help?

- Check existing issues
- Ask questions in issues with "question" label
- CI failures? Check the Actions tab for details

## License

Contributions are under the MIT License.