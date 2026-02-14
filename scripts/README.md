# Git Hooks

This directory contains Git hooks used in the Hauk-TUI project.

## Pre-Commit Hook

**Purpose:** Automatically verify that code compiles before allowing commits.

### Installation

After cloning the repository, install the pre-commit hook:

```bash
cp scripts/pre-commit .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
```

### What It Does

1. Checks if Go is installed
2. Runs `go build -o bin/hauk ./cmd/hauk/main.go`
3. Allows commit if build succeeds
4. Blocks commit if build fails

### Example Output

**Success:**
```
🔨 Running pre-commit checks...

📦 Building application...
✅ Build successful!

[commit proceeds normally]
```

**Failure:**
```
🔨 Running pre-commit checks...

📦 Building application...
# compilation errors appear here

❌ Build failed!

Please fix the compilation errors before committing.
[commit blocked]
```

### Optional: Enable Test Verification

Edit `.git/hooks/pre-commit` and uncomment lines 27-35 to also run tests before commits.

### Bypassing (Emergency Only)

If you absolutely must commit without running checks:

```bash
git commit --no-verify -m "emergency fix"
```

**⚠️ Warning:** Only use `--no-verify` in emergencies. The hook is there to prevent broken code from entering the repository.

## Why Git Hooks Aren't Tracked

Git hooks live in `.git/hooks/` which is not tracked by version control (for security reasons). That's why we keep a copy in `scripts/` and document the installation process.

## Future Hooks

Potential additions:
- **pre-push**: Run full test suite before pushing
- **commit-msg**: Enforce conventional commit format
- **post-merge**: Auto-install dependencies after pulling
