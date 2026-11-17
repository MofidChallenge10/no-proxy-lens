<p align="center"><img src="assets/banner.svg" alt="NO Proxy Lens — Explain whether a URL matches NO_PROXY and which rule caused the bypass." width="100%"></p>

# NO Proxy Lens ✨

> Explain whether a URL matches NO_PROXY and which rule caused the bypass.

[![CI](https://github.com/MofidChallenge10/no-proxy-lens/actions/workflows/ci.yml/badge.svg)](https://github.com/MofidChallenge10/no-proxy-lens/actions/workflows/ci.yml) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE) ![Stack](https://img.shields.io/badge/Stack-Go-informational)

## 🎯 Why this project?

Explain whether a URL matches NO_PROXY and which rule caused the bypass. It focuses on a narrow operational problem and keeps its output inspectable, deterministic, and suitable for scripts or CI.

## 🚀 Features

- 🧭 Solves the documented use case without a hosted service.
- 🧪 Includes repeatable tests and a runnable example.
- 📄 Supports concise human-readable output; machine-readable output is available where documented.
- 🔒 Avoids changing the target system during inspection.

## ⚡ Quick Start

```bash
go run ./cmd/no-proxy-lens --env 'localhost,.internal.example,10.0.0.0/8' https://api.internal.example:8443/health
```

## 🖥️ Example

```text
BYPASS  https://api.internal.example:8443/health
Reason: domain suffix rule ".internal.example"
```

## 🗺️ Roadmap

- Add an optional `--explain-all` mode showing every evaluated rule.
- Add conformance fixtures for additional client-specific NO_PROXY dialects.
- Keep current matching behavior documented and stable.

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Bug reports should include the command, a minimal input fixture, and the observed output. Please do not attach credentials or private production data.

## 📜 License

MIT — see [LICENSE](LICENSE).
